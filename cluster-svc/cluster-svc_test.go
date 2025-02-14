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

package main

import (
	"testing"
	"time"

	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/stretchr/testify/require"
)

var oldPrometheusControllerApp = edgeproto.App{
	Key:           MEXPrometheusAppKey,
	ImagePath:     "https://prometheus-community.github.io/helm-charts:prometheus-community/kube-prometheus-stack",
	Deployment:    cloudcommon.DeploymentTypeHelm,
	DelOpt:        edgeproto.DeleteType_AUTO_DELETE,
	InternalPorts: true,
	Trusted:       true,
}
var durationShort = "5s"
var durationLong = "45s"

func TestInfra(t *testing.T) {
	// Test duration to string conversion
	oldInterval, _ := time.ParseDuration(durationShort)
	require.Equal(t, durationShort, scrapeIntervalInSeconds(oldInterval), "scrapeIntervalInSeconds test")

	// Test fillConfig scrape interval 5 sec
	err := fillAppConfigs(&oldPrometheusControllerApp, oldInterval)
	require.Nil(t, err, "fillAppConfigs failed")
	// Should be a single Config there
	require.Equal(t, 1, len(oldPrometheusControllerApp.Configs), "Number of configs in app is wrong")
	*scrapeInterval, _ = time.ParseDuration(durationLong)
	newApp, err := getAppFromClusterSvc(&MEXPrometheusAppKey)
	require.Nil(t, err, "getPrometheusAppFromController failed")
	// Check that the fields that are different are correct
	setAppDiffFields(&oldPrometheusControllerApp, newApp)
	require.Equal(t, 4, len(newApp.Fields), "Incorrect number of different fields")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigs, "Missing edgeproto.AppFieldConfigs")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigsKind, "Missing edgeproto.AppFieldConfigsKind")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigsConfig, "Missing edgeproto.AppFieldConfigsConfig")

	// Change the image path and check that image path only gets set
	newApp2 := oldPrometheusControllerApp
	newApp2.ImagePath = "newImagePath"
	setAppDiffFields(&oldPrometheusControllerApp, &newApp2)
	require.Equal(t, 1, len(newApp2.Fields), "Incorrect number of different fields")
	require.Contains(t, newApp2.Fields, edgeproto.AppFieldImagePath, "Missing edgeproto.AppFieldImagePath")

	// If the apps have both the image and imagePath, make sure everything gets set
	*scrapeInterval, _ = time.ParseDuration(durationLong)
	newApp, err = getAppFromClusterSvc(&MEXPrometheusAppKey)
	require.Nil(t, err, "getAppFromClusterSvc failed")
	newApp.ImagePath = "newImagePath"
	// Check that the fields that are different are correct
	setAppDiffFields(&oldPrometheusControllerApp, newApp)
	require.Equal(t, 5, len(newApp.Fields), "Incorrect number of different fields")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigs, "Missing edgeproto.AppFieldConfigs")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigsKind, "Missing edgeproto.AppFieldConfigsKind")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldConfigsConfig, "Missing edgeproto.AppFieldConfigsConfig")
	require.Contains(t, newApp.Fields, edgeproto.AppFieldImagePath, "Missing edgeproto.AppFieldImagePath")
}
