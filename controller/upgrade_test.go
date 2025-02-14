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
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mobiledgex/edge-cloud/cloudcommon/node"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/mobiledgex/edge-cloud/objstore"
	"github.com/stretchr/testify/require"
)

var upgradeTestFileLocation = "./upgrade_testfiles"
var upgradeTestFilePreSuffix = "_pre.etcd"
var upgradeTestFilePostSuffix = "_post.etcd"

// Walk testutils data and populate objStore
func buildDbFromTestData(objStore objstore.KVStore, funcName string) error {
	var key, val string
	ctx := context.Background()

	filename := upgradeTestFileLocation + "/" + funcName + upgradeTestFilePreSuffix
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Unable to find preupgrade testdata file at %s", filename)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for {
		if !scanner.Scan() {
			break
		}
		key = scanner.Text()
		if !scanner.Scan() {
			return fmt.Errorf("Improper formatted preupgrade .etcd file - Unmatched key, without a value.")
		}
		val = scanner.Text()
		if _, err := objStore.Put(ctx, key, val); err != nil {
			return err
		}
	}
	return nil
}

// walk testutils data and see if the entries exist in the objstore
func compareDbToExpected(objStore objstore.KVStore, funcName string) error {
	var dbObjCount, fileObjCount int

	var key, val string

	filename := upgradeTestFileLocation + "/" + funcName + upgradeTestFilePostSuffix
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Unable to find postupgrade testdata file at %s", filename)
	}
	defer file.Close()

	fileExpected, err := os.Create(upgradeTestFileLocation + "/" + funcName + "_expected.etcd")
	if err != nil {
		return err
	}
	defer fileExpected.Close()
	writtenKeys := make(map[string]struct{})

	var compareErr error
	scanner := bufio.NewScanner(file)
	for {
		if !scanner.Scan() {
			break
		}
		key = scanner.Text()
		if !scanner.Scan() {
			return fmt.Errorf("Improper formatted postupgrade .etcd file - Unmatched key, without a value.")
		}
		val = scanner.Text()
		dbVal, _, _, err := objStore.Get(key)
		if err != nil {
			return fmt.Errorf("Unable to get value for key[%s], %v", key, err)
		}
		// data may be in json format or non-json string
		compareDone, err := compareJson(funcName, key, val, string(dbVal))
		if !compareDone {
			err = compareString(funcName, key, val, string(dbVal))
		}
		if err != nil && compareErr == nil {
			// continue writing to expected file
			compareErr = err
		}
		fileExpected.WriteString(string(key) + "\n")
		fileExpected.WriteString(string(dbVal) + "\n")
		writtenKeys[string(key)] = struct{}{}
		dbObjCount++
		fileObjCount++
	}
	err = objStore.List("", func(key, val []byte, rev, modRev int64) error {
		if _, found := writtenKeys[string(key)]; found {
			return nil
		}
		fileExpected.WriteString(string(key) + "\n")
		fileExpected.WriteString(string(val) + "\n")
		dbObjCount++
		return nil
	})
	if compareErr != nil {
		return compareErr
	}
	if err != nil {
		return err
	}
	if fileObjCount != dbObjCount {
		return fmt.Errorf("Number of objects in the etcd db[%d] doesn't match the number of expected objects[%d]\n",
			dbObjCount, fileObjCount)
	}
	return nil
}

func compareJson(funcName, key, expected, actual string) (bool, error) {
	expectedMap := make(map[string]interface{})
	actualMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(expected), &expectedMap)
	if err != nil {
		return false, fmt.Errorf("Unmarshal failed, %v, %s\n", err, expected)
	}
	err = json.Unmarshal([]byte(actual), &actualMap)
	if err != nil {
		return false, fmt.Errorf("Unmarshal failed, %v, %s\n", err, actual)
	}
	if !cmp.Equal(expectedMap, actualMap) {
		fmt.Printf("[%s] comparsion fail for key: %s\n", funcName, key)
		fmt.Printf("expected vs actual:\n")
		fmt.Printf(cmp.Diff(expectedMap, actualMap))
		return true, fmt.Errorf("Values don't match for the key, upgradeFunc: %s", funcName)
	}
	return true, nil
}

func compareString(funcName, key, expected, actual string) error {
	if expected != actual {
		fmt.Printf("[%s] values don't match for the key: %s\n", funcName, key)
		fmt.Printf("[%s] expected: \n%s\n", funcName, expected)
		fmt.Printf("[%s] actual: \n%s\n", funcName, actual)
		return fmt.Errorf("Values don't match for the key, upgradeFunc: %s", funcName)
	}
	return nil
}

// Run each upgrade function after populating dummy etcd with test data.
// Verify that the resulting content in etcd matches expected
func TestAllUpgradeFuncs(t *testing.T) {
	log.SetDebugLevel(log.DebugLevelUpgrade | log.DebugLevelApi)
	objStore := dummyEtcd{}
	objstore.InitRegion(1)
	log.InitTracer(nil)
	defer log.FinishTracer()

	// There are timestamp fields which are stored in RFC3339 format.
	// Hence, fix a timezone for consistent comparison
	time.Local = time.UTC

	cplookup := &node.CloudletPoolCache{}
	cplookup.Init()
	nodeMgr.CloudletPoolLookup = cplookup
	cloudletLookup := &node.CloudletCache{}
	cloudletLookup.Init()
	nodeMgr.CloudletLookup = cloudletLookup

	sync := InitSync(&objStore)
	apis := NewAllApis(sync)

	ctx := log.StartTestSpan(context.Background())
	for ii, fn := range VersionHash_UpgradeFuncs {
		if fn == nil {
			continue
		}
		objStore.Start()
		err := buildDbFromTestData(&objStore, VersionHash_UpgradeFuncNames[ii])
		require.Nil(t, err, "Unable to build db from testData")
		err = RunSingleUpgrade(ctx, &objStore, apis, fn)
		require.Nil(t, err, "Upgrade failed")
		err = compareDbToExpected(&objStore, VersionHash_UpgradeFuncNames[ii])
		require.Nil(t, err, "Unexpected result from upgrade function(%s)", VersionHash_UpgradeFuncNames[ii])
		// Run the upgrade again to make sure it's idempotent
		err = RunSingleUpgrade(ctx, &objStore, apis, fn)
		require.Nil(t, err, "Upgrade second run failed")
		err = compareDbToExpected(&objStore, VersionHash_UpgradeFuncNames[ii])
		require.Nil(t, err, "Unexpected result from upgrade function second run (idempotency check) (%s)", VersionHash_UpgradeFuncNames[ii])
		// Stop it, so it's re-created again
		objStore.Stop()
	}
	//manually test a failure of checkHttpPorts upgrade
	objStore.Start()
	err := buildDbFromTestData(&objStore, "CheckForHttpPortsFail")
	require.Nil(t, err, "Unable to build db from testData")
	err = RunSingleUpgrade(ctx, &objStore, apis, CheckForHttpPorts)
	require.NotNil(t, err, "Upgrade did not fail")
	objStore.Stop()
}
