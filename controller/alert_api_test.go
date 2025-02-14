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
	"time"

	"github.com/mobiledgex/edge-cloud/cloudcommon"
	"github.com/mobiledgex/edge-cloud/cloudcommon/node"
	influxq "github.com/mobiledgex/edge-cloud/controller/influxq_client"
	dme "github.com/mobiledgex/edge-cloud/d-match-engine/dme-proto"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/integration/process"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/mobiledgex/edge-cloud/objstore"
	"github.com/mobiledgex/edge-cloud/rediscache"
	"github.com/mobiledgex/edge-cloud/testutil"
	"github.com/mobiledgex/edge-cloud/vault"
	"github.com/stretchr/testify/require"
)

func TestAlertApi(t *testing.T) {
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

	for _, alert := range testutil.AlertData {
		apis.alertApi.Update(ctx, &alert, 0)
	}
	testutil.InternalAlertTest(t, "show", apis.alertApi, testutil.AlertData)

	cloudletData := testutil.CloudletData()
	testutil.InternalFlavorCreate(t, apis.flavorApi, testutil.FlavorData)
	testutil.InternalGPUDriverCreate(t, apis.gpuDriverApi, testutil.GPUDriverData)
	testutil.InternalResTagTableCreate(t, apis.resTagTableApi, testutil.ResTagTableData)
	testutil.InternalCloudletCreate(t, apis.cloudletApi, cloudletData)
	testCloudlet := cloudletData[0]
	testCloudlet.Key.Name = "testcloudlet"
	testutil.InternalCloudletCreate(t, apis.cloudletApi, []edgeproto.Cloudlet{testCloudlet})
	testCloudletInfo := testutil.CloudletInfoData[0]
	testCloudletInfo.Key.Name = testCloudlet.Key.Name
	insertCloudletInfo(ctx, apis, []edgeproto.CloudletInfo{testCloudletInfo})
	getAlertsCount := func() (int, int) {
		count := 0
		totalCount := 0
		for _, data := range apis.alertApi.cache.Objs {
			val := data.Obj
			totalCount++
			if cloudletName, found := val.Labels[edgeproto.CloudletKeyTagName]; !found ||
				cloudletName != testCloudlet.Key.Name {
				continue
			}
			if cloudletOrg, found := val.Labels[edgeproto.CloudletKeyTagOrganization]; !found ||
				cloudletOrg != testCloudlet.Key.Organization {
				continue
			}
			count++
		}
		return count, totalCount
	}
	cloudletCount, totalCount := getAlertsCount()
	require.Greater(t, cloudletCount, 0, "cloudlet alerts exists")
	require.Greater(t, totalCount, 0, "alerts exists")
	err := apis.cloudletApi.DeleteCloudlet(&testCloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err, "delete cloudlet")
	expectedTotalCount := totalCount - cloudletCount
	cloudletCount, totalCount = getAlertsCount()
	require.Equal(t, cloudletCount, 0, "cloudlet alerts should not exist")
	require.Equal(t, totalCount, expectedTotalCount, "expected alerts should exist")

	dummy.Stop()
}

func TestAppInstDownAlert(t *testing.T) {
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
	// Create a reservable clusterInst
	cinst := testutil.ClusterInstData[7]
	streamOut := testutil.NewCudStreamoutAppInst(ctx)
	appinst := edgeproto.AppInst{}
	appinst.Key.AppKey = testutil.AppData[0].Key
	appinst.Key.ClusterInstKey = *cinst.Key.Virtual("")
	err := apis.appInstApi.CreateAppInst(&appinst, streamOut)
	require.Nil(t, err, "create AppInst")
	// Inject AppInst info check that all appInsts are Healthy
	for ii, _ := range testutil.AppInstInfoData {
		in := &testutil.AppInstInfoData[ii]
		apis.appInstInfoApi.Update(ctx, in, 0)
	}
	for _, val := range apis.appInstApi.cache.Objs {
		require.Equal(t, dme.HealthCheck_HEALTH_CHECK_OK, val.Obj.HealthCheck)
	}
	// Trigger Alerts
	for _, alert := range testutil.AlertData {
		apis.alertApi.Update(ctx, &alert, 0)
	}
	// Check reservable cluster

	found := apis.appInstApi.Get(&appinst.Key, &appinst)
	require.True(t, found)
	require.Equal(t, dme.HealthCheck_HEALTH_CHECK_ROOTLB_OFFLINE, appinst.HealthCheck)
	// check other appInstances
	for ii, testData := range testutil.CreatedAppInstData() {
		found = apis.appInstApi.Get(&testData.Key, &appinst)
		require.True(t, found)
		if ii == 0 {
			require.Equal(t, dme.HealthCheck_HEALTH_CHECK_SERVER_FAIL, appinst.HealthCheck)
		} else {
			require.Equal(t, dme.HealthCheck_HEALTH_CHECK_OK, appinst.HealthCheck)
		}
	}

	dummy.Stop()
}

type testServices struct {
	DummyRedisSrv *rediscache.DummyRedis
	RedisLocalSrv *process.RedisCache
}

type TestOptions struct {
	// Start local redis server
	LocalRedis bool
}

type TestOp func(op *TestOptions)

func WithLocalRedis() TestOp {
	return func(op *TestOptions) { op.LocalRedis = true }
}

// Set up globals for API unit tests
func testinit(ctx context.Context, t *testing.T, opts ...TestOp) *testServices {
	options := TestOptions{}
	for _, op := range opts {
		op(&options)
	}
	svcs := &testServices{}
	objstore.InitRegion(1)
	tMode := true
	testMode = &tMode
	dockerRegistry := "docker.mobiledgex.net"
	registryFQDN = &dockerRegistry
	vaultConfig, _ = vault.BestConfig("")
	services.events = influxq.NewInfluxQ("events", "user", "pass")
	services.cloudletResourcesInfluxQ = influxq.NewInfluxQ(cloudcommon.CloudletResourceUsageDbName, "user", "pass")
	cleanupCloudletInfoTimeout = 100 * time.Millisecond
	RequireAppInstPortConsistency = true
	cplookup := &node.CloudletPoolCache{}
	cplookup.Init()
	nodeMgr.CloudletPoolLookup = cplookup
	cloudletLookup := &node.CloudletCache{}
	cloudletLookup.Init()
	nodeMgr.CloudletLookup = cloudletLookup
	if options.LocalRedis {
		// Since it is a single node, config file is not required
		procOpts := []process.StartOp{process.WithNoConfig()}
		redisLocal, err := StartLocalRedisServer(procOpts...)
		require.Nil(t, err, "start redis server")
		svcs.RedisLocalSrv = redisLocal
		redisCfg = rediscache.RedisConfig{
			StandaloneAddr: rediscache.DefaultRedisStandaloneAddr,
		}
		redisClient, err = rediscache.NewClient(ctx, &redisCfg)
		require.Nil(t, err, "setup redis client")
	} else {
		redisServer, err := rediscache.NewMockRedisServer()
		require.Nil(t, err, "start mock redis server")
		svcs.DummyRedisSrv = redisServer
		redisCfg = rediscache.RedisConfig{
			SentinelAddrs: redisServer.GetSentinelAddr(),
		}
		redisClient, err = rediscache.NewClient(ctx, &redisCfg)
		require.Nil(t, err, "setup redis client")
	}

	return svcs
}

func testfinish(s *testServices) {
	if redisClient != nil {
		redisClient.Close()
		redisClient = nil
	}
	if s.DummyRedisSrv != nil {
		s.DummyRedisSrv.Close()
		s.DummyRedisSrv = nil
	}
	if s.RedisLocalSrv != nil {
		s.RedisLocalSrv.StopLocal()
		s.RedisLocalSrv = nil
	}
	services = Services{}
}
