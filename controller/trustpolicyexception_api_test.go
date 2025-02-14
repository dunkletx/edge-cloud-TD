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
	"context"
	"testing"

	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/mobiledgex/edge-cloud/testutil"
	"github.com/stretchr/testify/require"
)

func TestTrustPolicyExceptionApi(t *testing.T) {
	log.SetDebugLevel(log.DebugLevelEtcd | log.DebugLevelApi)
	log.InitTracer(nil)
	defer log.FinishTracer()
	ctx := log.StartTestSpan(context.Background())
	testSvcs := testinit(ctx, t)
	defer testfinish(testSvcs)

	dummy := dummyEtcd{}
	dummy.Start()

	sync := InitSync(&dummy)
	apis := NewAllApis(sync)
	sync.Start()
	defer sync.Done()

	dummyResponder := DummyInfoResponder{
		AppInstCache:        &apis.appInstApi.cache,
		ClusterInstCache:    &apis.clusterInstApi.cache,
		RecvAppInstInfo:     apis.appInstInfoApi,
		RecvClusterInstInfo: apis.clusterInstInfoApi,
	}
	dummyResponder.InitDummyInfoResponder()

	// create supporting data
	testutil.InternalFlavorCreate(t, apis.flavorApi, testutil.FlavorData)
	testutil.InternalGPUDriverCreate(t, apis.gpuDriverApi, testutil.GPUDriverData)
	testutil.InternalResTagTableCreate(t, apis.resTagTableApi, testutil.ResTagTableData)
	testutil.InternalCloudletCreate(t, apis.cloudletApi, testutil.CloudletData())
	insertCloudletInfo(ctx, apis, testutil.CloudletInfoData)
	testutil.InternalAutoProvPolicyCreate(t, apis.autoProvPolicyApi, testutil.AutoProvPolicyData)
	testutil.InternalAutoScalePolicyCreate(t, apis.autoScalePolicyApi, testutil.AutoScalePolicyData)
	testutil.InternalAppCreate(t, apis.appApi, testutil.AppData)
	testutil.InternalClusterInstCreate(t, apis.clusterInstApi, testutil.ClusterInstData)
	testutil.InternalAppInstCreate(t, apis.appInstApi, testutil.AppInstData)
	testutil.InternalCloudletPoolTest(t, "cud", apis.cloudletPoolApi, testutil.CloudletPoolData)

	// CUD for Trust Policy Exception
	testutil.InternalTrustPolicyExceptionTest(t, "cud", apis.trustPolicyExceptionApi, testutil.TrustPolicyExceptionData)

	// Basic error case - when TPE already exists
	_, err := apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &testutil.TrustPolicyExceptionData[0])
	require.NotNil(t, err)
	require.Contains(t, err.Error(), " already exists")

	// This is an app for which Trusted is not set
	tpeDataFail := edgeproto.TrustPolicyException{
		Key: edgeproto.TrustPolicyExceptionKey{
			AppKey: edgeproto.AppKey{
				Organization: testutil.DevData[0],
				Name:         "Pillimo Go!",
				Version:      "1.0.1",
			},
			CloudletPoolKey: edgeproto.CloudletPoolKey{
				Organization: testutil.OperatorData[2],
				Name:         "test-and-dev",
			},
			Name: "someapp-tpe2",
		},
		State: edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_APPROVAL_REQUESTED,
		OutboundSecurityRules: []edgeproto.SecurityRule{
			edgeproto.SecurityRule{
				Protocol:     "tcp",
				RemoteCidr:   "10.1.0.0/16",
				PortRangeMin: 201,
				PortRangeMax: 210,
			},
		},
	}
	_, err = apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &tpeDataFail)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Non trusted app")

	// Without any security rules
	tpeDataFail2 := edgeproto.TrustPolicyException{
		Key: edgeproto.TrustPolicyExceptionKey{
			AppKey: edgeproto.AppKey{
				Organization: testutil.DevData[0],
				Name:         "Pillimo Go!",
				Version:      "1.0.0",
			},
			CloudletPoolKey: edgeproto.CloudletPoolKey{
				Organization: testutil.OperatorData[2],
				Name:         "test-and-dev",
			},
			Name: "someapp-tpe2",
		},
		State: edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_APPROVAL_REQUESTED,
	}
	_, err = apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &tpeDataFail2)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Security rules must be specified")

	tpeData := edgeproto.TrustPolicyException{
		Key: edgeproto.TrustPolicyExceptionKey{
			AppKey: edgeproto.AppKey{
				Organization: testutil.DevData[0],
				Name:         "Pillimo Go!",
				Version:      "1.0.0",
			},
			CloudletPoolKey: edgeproto.CloudletPoolKey{
				Organization: testutil.OperatorData[2],
				Name:         "test-and-dev",
			},
			Name: "someapp-tpe2",
		},
		State: edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_APPROVAL_REQUESTED,
		OutboundSecurityRules: []edgeproto.SecurityRule{
			edgeproto.SecurityRule{
				Protocol:     "tcp",
				RemoteCidr:   "10.1.0.0/16",
				PortRangeMin: 201,
				PortRangeMax: 210,
			},
		},
	}
	_, err = apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &tpeData)
	require.Nil(t, err)

	// State related tests - begin
	tpeData.Fields = []string{edgeproto.TrustPolicyExceptionFieldState}

	// test that TPE update state to STATE_ACTIVE, passes
	tpeData.State = edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_ACTIVE
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.Nil(t, err)

	// test that TPE update state to STATE_REJECTED
	tpeData.State = edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_REJECTED
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.Nil(t, err)

	// test that TPE update state to STATE_APPROVAL_REQUESTED, fails
	tpeData.State = edgeproto.TrustPolicyExceptionState_TRUST_POLICY_EXCEPTION_STATE_APPROVAL_REQUESTED
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "New state must be either Active or Rejected")

	// test that TPE update with no security rules, does not give any error
	savedSecurityRules := tpeData.OutboundSecurityRules
	tpeData.OutboundSecurityRules = []edgeproto.SecurityRule{}
	tpeData.Fields = []string{edgeproto.TrustPolicyExceptionFieldOutboundSecurityRules}
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.Nil(t, err)
	tpeData.OutboundSecurityRules = savedSecurityRules

	// test that TPE update with non-existent CloudletPoolKey Organization, fails
	tpeData.Fields = []string{
		edgeproto.TrustPolicyExceptionFieldKeyCloudletPoolKeyOrganization,
		edgeproto.TrustPolicyExceptionFieldKeyCloudletPoolKey}
	tpeData.Key.CloudletPoolKey.Organization = "MarsCloudletPoolOrg"
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.NotNil(t, err)
	strCloudOrgErr := "TrustPolicyException key {\"app_key\":{\"organization\":\"AtlanticInc\",\"name\":\"Pillimo Go!\",\"version\":\"1.0.0\"},\"cloudlet_pool_key\":{\"organization\":\"MarsCloudletPoolOrg\",\"name\":\"test-and-dev\"},\"name\":\"someapp-tpe2\"} not found"
	require.Equal(t, err.Error(), strCloudOrgErr)

	// test that TPE update with non-existent AppKey Organization, fails
	tpeData.Fields = []string{
		edgeproto.TrustPolicyExceptionFieldKeyAppKey,
		edgeproto.TrustPolicyExceptionFieldKeyAppKeyOrganization}
	tpeData.Key.AppKey.Organization = "MarsAppOrg"
	_, err = apis.trustPolicyExceptionApi.UpdateTrustPolicyException(ctx, &tpeData)
	require.NotNil(t, err)
	strAppOrgErr := "TrustPolicyException key {\"app_key\":{\"organization\":\"MarsAppOrg\",\"name\":\"Pillimo Go!\",\"version\":\"1.0.0\"},\"cloudlet_pool_key\":{\"organization\":\"MarsCloudletPoolOrg\",\"name\":\"test-and-dev\"},\"name\":\"someapp-tpe2\"} not found"
	require.Equal(t, err.Error(), strAppOrgErr)

	// State related tests - end, restore everything
	tpeData.Key.AppKey.Organization = testutil.DevData[0]
	tpeData.Fields = []string{}
	// test that TPE create when specified CloudletPool does not exist, fails
	tpeData.Key.CloudletPoolKey.Organization = "Mission Mars"
	_, err = apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &tpeData)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), tpeData.Key.CloudletPoolKey.NotFoundError().Error())
	// Restore tpeData Key to original values
	tpeData.Key.CloudletPoolKey.Organization = testutil.OperatorData[2]

	// test that TPE create when specified App does not exist, fails
	tpeData.Key.AppKey.Organization = testutil.DevData[2]
	_, err = apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, &tpeData)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), tpeData.Key.AppKey.NotFoundError().Error())
	// Restore tpeData Key to original values
	tpeData.Key.AppKey.Organization = testutil.DevData[0]

	testutil.InternalAppInstDelete(t, apis.appInstApi, testutil.AppInstData)

	// test that App delete fails if TPE exists that refers to it
	app0 := testutil.AppData[0]
	_, err = apis.appApi.DeleteApp(ctx, &app0)
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "Application in use by Trust Policy Exception")

	// Success : Delete
	_, err = apis.trustPolicyExceptionApi.DeleteTrustPolicyException(ctx, &tpeData)
	require.Nil(t, err)

	// error cases for Create Trust Policy Exception
	expectCreatePolicyExceptionError(t, ctx, apis, &testutil.TrustPolicyExceptionErrorData[0], "cannot be higher than max")
	expectCreatePolicyExceptionError(t, ctx, apis, &testutil.TrustPolicyExceptionErrorData[1], "invalid CIDR")
	expectCreatePolicyExceptionError(t, ctx, apis, &testutil.TrustPolicyExceptionErrorData[2], "Invalid min port")
	expectCreatePolicyExceptionError(t, ctx, apis, &testutil.TrustPolicyExceptionErrorData[3],
		testutil.TrustPolicyExceptionErrorData[3].Key.AppKey.NotFoundError().Error())
	expectCreatePolicyExceptionError(t, ctx, apis, &testutil.TrustPolicyExceptionErrorData[4],
		testutil.TrustPolicyExceptionErrorData[4].Key.CloudletPoolKey.NotFoundError().Error())

	dummy.Stop()
}

func expectCreatePolicyExceptionError(t *testing.T, ctx context.Context, apis *AllApis, in *edgeproto.TrustPolicyException, msg string) {
	_, err := apis.trustPolicyExceptionApi.CreateTrustPolicyException(ctx, in)
	require.NotNil(t, err, "create %v", in)
	require.Contains(t, err.Error(), msg, "error %v contains %s", err, msg)
}
