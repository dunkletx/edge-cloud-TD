// Copyright 2022 MobiledgeX, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8smgmt

import (
	"fmt"
	"strings"

	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/mobiledgex/edge-cloud/util"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type KubeNames struct {
	AppName                    string
	AppVersion                 string
	AppOrg                     string
	HelmAppName                string
	AppURI                     string
	AppImage                   string
	AppRevision                string
	AppInstRevision            string
	ClusterName                string
	K8sNodeNameSuffix          string
	OperatorName               string
	ServiceNames               []string
	DeveloperDefinedNamespaces []string // namespaces included by developer in manifest
	KconfName                  string
	KconfEnv                   string
	DeploymentType             string
	ImagePullSecrets           []string
	ImagePaths                 []string
	IsUriIPAddr                bool
	MultitenantNamespace       string // for apps launched in a multi-tenant cluster
	BaseKconfName              string
}

type KubeNamesOp func(k *KubeNames) error

func GetKconfName(clusterInst *edgeproto.ClusterInst) string {
	return fmt.Sprintf("%s.%s.kubeconfig",
		clusterInst.Key.ClusterKey.Name,
		clusterInst.Key.CloudletKey.Organization)
}

func GetK8sNodeNameSuffix(clusterInstKey *edgeproto.ClusterInstKey) string {
	cloudletName := clusterInstKey.CloudletKey.Name
	clusterName := clusterInstKey.ClusterKey.Name
	devName := clusterInstKey.Organization
	if devName != "" {
		return NormalizeName(cloudletName + "-" + clusterName + "-" + devName)
	}
	return NormalizeName(cloudletName + "-" + clusterName)
}

// GetCloudletClusterName return the name of the cluster including cloudlet
func GetCloudletClusterName(clusterKey *edgeproto.ClusterInstKey) string {
	return GetK8sNodeNameSuffix(clusterKey)
}

func GetNamespace(appInstKey *edgeproto.AppInstKey) string {
	// Note that we use the virtual cluster name, not the real cluster name
	return util.NamespaceSanitize(fmt.Sprintf("%s-%s-%s-%s", appInstKey.AppKey.Organization, appInstKey.AppKey.Name, appInstKey.AppKey.Version, appInstKey.ClusterInstKey.ClusterKey.Name))
}

func NormalizeName(name string) string {
	return util.K8SSanitize(name)
}

// FixImagePath removes localhost and adds Docker Hub as needed.  For example,
// networkstatic/iperf3 becomes docker.io/networkstatic/iperf3
func FixImagePath(origImagePath string) string {
	newImagePath := origImagePath
	parts := strings.Split(origImagePath, "/")
	if parts[0] == "localhost" {
		newImagePath = strings.Replace(origImagePath, "localhost/", "", -1)
	} else {
		// Append default registry address for internal image paths
		if len(parts) < 2 || !strings.Contains(parts[0], ".") {
			newImagePath = cloudcommon.DockerHub + "/" + origImagePath
		}
	}
	return newImagePath
}

func GetNormalizedClusterName(clusterInst *edgeproto.ClusterInst) string {
	return NormalizeName(clusterInst.Key.ClusterKey.Name + clusterInst.Key.Organization)
}

// GetKubeNames udpates kubeNames with normalized strings for the included clusterinst, app, and appisnt
func GetKubeNames(clusterInst *edgeproto.ClusterInst, app *edgeproto.App, appInst *edgeproto.AppInst, opts ...KubeNamesOp) (*KubeNames, error) {
	if clusterInst == nil {
		return nil, fmt.Errorf("nil cluster inst")
	}
	if app == nil {
		return nil, fmt.Errorf("nil app")
	}
	if appInst == nil {
		return nil, fmt.Errorf("nil app inst")
	}
	kubeNames := KubeNames{}
	for _, op := range opts {
		if err := op(&kubeNames); err != nil {
			return nil, err
		}
	}
	kubeNames.ClusterName = GetNormalizedClusterName(clusterInst)
	kubeNames.K8sNodeNameSuffix = GetK8sNodeNameSuffix(&clusterInst.Key)
	kubeNames.AppName = NormalizeName(app.Key.Name)
	kubeNames.AppVersion = NormalizeName(app.Key.Version)
	kubeNames.AppOrg = NormalizeName(app.Key.Organization)
	// Helm app name has to conform to DNS naming standards
	kubeNames.HelmAppName = util.DNSSanitize(app.Key.Name + "v" + app.Key.Version)
	kubeNames.AppURI = appInst.Uri
	kubeNames.AppRevision = app.Revision
	kubeNames.AppInstRevision = appInst.Revision
	kubeNames.AppImage = NormalizeName(app.ImagePath)
	kubeNames.OperatorName = NormalizeName(clusterInst.Key.CloudletKey.Organization)
	kubeNames.KconfName = GetKconfName(clusterInst)
	// if clusterInst is multi-tenant and AppInst is specified, use namespaced config
	if clusterInst.MultiTenant && appInst.Key.AppKey.Name != "" && !cloudcommon.IsSideCarApp(app) {
		kubeNames.MultitenantNamespace = GetNamespace(&appInst.Key)
		kubeNames.BaseKconfName = kubeNames.KconfName
		baseName := strings.TrimSuffix(kubeNames.KconfName, ".kubeconfig")
		kubeNames.KconfName = fmt.Sprintf("%s.%s.kubeconfig", baseName, kubeNames.MultitenantNamespace)
	}
	kubeNames.KconfEnv = "KUBECONFIG=" + kubeNames.KconfName
	kubeNames.DeploymentType = app.Deployment
	if app.ImagePath != "" {
		kubeNames.ImagePaths = append(kubeNames.ImagePaths, app.ImagePath)
	}
	//get service names from the yaml
	if app.Deployment == cloudcommon.DeploymentTypeKubernetes {
		objs, _, err := cloudcommon.DecodeK8SYaml(app.DeploymentManifest)
		if err != nil {
			return nil, fmt.Errorf("invalid kubernetes deployment yaml, %s", err.Error())
		}
		var template *v1.PodTemplateSpec
		for _, o := range objs {
			log.DebugLog(log.DebugLevelInfra, "k8s obj", "obj", o)
			template = nil
			switch obj := o.(type) {
			case *v1.Service:
				svcName := obj.ObjectMeta.Name
				kubeNames.ServiceNames = append(kubeNames.ServiceNames, svcName)
			case *appsv1.Deployment:
				template = &obj.Spec.Template
			case *appsv1.DaemonSet:
				template = &obj.Spec.Template
			case *appsv1.StatefulSet:
				template = &obj.Spec.Template
			case *v1.Namespace:
				// if this is not a multi tenant case, any additional namespaces are from a developer manifest
				if kubeNames.MultitenantNamespace == "" {
					kubeNames.DeveloperDefinedNamespaces = append(kubeNames.DeveloperDefinedNamespaces, obj.Name)
				}
			}
			if template == nil {
				continue
			}
			containers := []v1.Container{}
			containers = append(containers, template.Spec.InitContainers...)
			containers = append(containers, template.Spec.Containers...)
			for _, cont := range containers {
				if cont.Image == "" {
					continue
				}
				kubeNames.ImagePaths = append(kubeNames.ImagePaths, cont.Image)
			}
		}
	} else if app.Deployment == cloudcommon.DeploymentTypeHelm {
		// for helm chart just make sure it's the same prefix
		kubeNames.ServiceNames = append(kubeNames.ServiceNames, kubeNames.AppName)
	} else if app.Deployment == cloudcommon.DeploymentTypeDocker {
		// for docker use the app name
		kubeNames.ServiceNames = append(kubeNames.ServiceNames, kubeNames.AppName)
		if app.DeploymentManifest != "" && !strings.HasSuffix(app.DeploymentManifest, ".zip") {
			containers, err := cloudcommon.DecodeDockerComposeYaml(app.DeploymentManifest)
			if err != nil {
				return nil, fmt.Errorf("invalid docker compose yaml, %s", err.Error())
			}
			for _, cont := range containers {
				kubeNames.ImagePaths = append(kubeNames.ImagePaths, FixImagePath(cont.Image))
			}
		}
	}
	return &kubeNames, nil
}

func (k *KubeNames) ContainsService(svc string) bool {
	for _, s := range k.ServiceNames {
		if strings.HasPrefix(svc, s) {
			return true
		}
	}
	return false
}
