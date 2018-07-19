// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: clusterinst.proto

package testutil

import "google.golang.org/grpc"
import "github.com/mobiledgex/edge-cloud/edgeproto"
import "io"
import "testing"
import "context"
import "time"
import "github.com/stretchr/testify/assert"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/mobiledgex/edge-cloud/protogen"
import _ "github.com/mobiledgex/edge-cloud/protoc-gen-cmd/protocmd"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

type ShowClusterInst struct {
	Data map[string]edgeproto.ClusterInst
	grpc.ServerStream
}

func (x *ShowClusterInst) Init() {
	x.Data = make(map[string]edgeproto.ClusterInst)
}

func (x *ShowClusterInst) Send(m *edgeproto.ClusterInst) error {
	x.Data[m.Key.GetKeyString()] = *m
	return nil
}

func (x *ShowClusterInst) ReadStream(stream edgeproto.ClusterInstApi_ShowClusterInstClient, err error) {
	x.Data = make(map[string]edgeproto.ClusterInst)
	if err != nil {
		return
	}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		x.Data[obj.Key.GetKeyString()] = *obj
	}
}

func (x *ShowClusterInst) CheckFound(obj *edgeproto.ClusterInst) bool {
	_, found := x.Data[obj.Key.GetKeyString()]
	return found
}

func (x *ShowClusterInst) AssertFound(t *testing.T, obj *edgeproto.ClusterInst) {
	check, found := x.Data[obj.Key.GetKeyString()]
	assert.True(t, found, "find ClusterInst %s", obj.Key.GetKeyString())
	if found && !check.MatchesIgnoreBackend(obj) {
		assert.Equal(t, *obj, check, "ClusterInst are equal")
	}
}

func (x *ShowClusterInst) AssertNotFound(t *testing.T, obj *edgeproto.ClusterInst) {
	_, found := x.Data[obj.Key.GetKeyString()]
	assert.False(t, found, "do not find ClusterInst %s", obj.Key.GetKeyString())
}

func WaitAssertFoundClusterInst(t *testing.T, api edgeproto.ClusterInstApiClient, obj *edgeproto.ClusterInst, count int, retry time.Duration) {
	show := ShowClusterInst{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowClusterInst(ctx, obj)
		show.ReadStream(stream, err)
		cancel()
		if show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertFound(t, obj)
}

func WaitAssertNotFoundClusterInst(t *testing.T, api edgeproto.ClusterInstApiClient, obj *edgeproto.ClusterInst, count int, retry time.Duration) {
	show := ShowClusterInst{}
	filterNone := edgeproto.ClusterInst{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowClusterInst(ctx, &filterNone)
		show.ReadStream(stream, err)
		cancel()
		if !show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertNotFound(t, obj)
}

// Wrap the api with a common interface
type ClusterInstCommonApi struct {
	internal_api edgeproto.ClusterInstApiServer
	client_api   edgeproto.ClusterInstApiClient
}

func (x *ClusterInstCommonApi) CreateClusterInst(ctx context.Context, in *edgeproto.ClusterInst) (*edgeproto.Result, error) {
	if x.internal_api != nil {
		return x.internal_api.CreateClusterInst(ctx, in)
	} else {
		return x.client_api.CreateClusterInst(ctx, in)
	}
}

func (x *ClusterInstCommonApi) UpdateClusterInst(ctx context.Context, in *edgeproto.ClusterInst) (*edgeproto.Result, error) {
	if x.internal_api != nil {
		return x.internal_api.UpdateClusterInst(ctx, in)
	} else {
		return x.client_api.UpdateClusterInst(ctx, in)
	}
}

func (x *ClusterInstCommonApi) DeleteClusterInst(ctx context.Context, in *edgeproto.ClusterInst) (*edgeproto.Result, error) {
	if x.internal_api != nil {
		return x.internal_api.DeleteClusterInst(ctx, in)
	} else {
		return x.client_api.DeleteClusterInst(ctx, in)
	}
}

func (x *ClusterInstCommonApi) ShowClusterInst(ctx context.Context, filter *edgeproto.ClusterInst, showData *ShowClusterInst) error {
	if x.internal_api != nil {
		return x.internal_api.ShowClusterInst(filter, showData)
	} else {
		stream, err := x.client_api.ShowClusterInst(ctx, filter)
		showData.ReadStream(stream, err)
		return err
	}
}

func NewInternalClusterInstApi(api edgeproto.ClusterInstApiServer) *ClusterInstCommonApi {
	apiWrap := ClusterInstCommonApi{}
	apiWrap.internal_api = api
	return &apiWrap
}

func NewClientClusterInstApi(api edgeproto.ClusterInstApiClient) *ClusterInstCommonApi {
	apiWrap := ClusterInstCommonApi{}
	apiWrap.client_api = api
	return &apiWrap
}
func InternalClusterInstCudTest(t *testing.T, api edgeproto.ClusterInstApiServer, testData []edgeproto.ClusterInst) {
	basicClusterInstCudTest(t, NewInternalClusterInstApi(api), testData)
}

func ClientClusterInstCudTest(t *testing.T, api edgeproto.ClusterInstApiClient, testData []edgeproto.ClusterInst) {
	basicClusterInstCudTest(t, NewClientClusterInstApi(api), testData)
}

func basicClusterInstCudTest(t *testing.T, api *ClusterInstCommonApi, testData []edgeproto.ClusterInst) {
	var err error
	ctx := context.TODO()

	if len(testData) < 3 {
		assert.True(t, false, "Need at least 3 test data objects")
		return
	}

	// test create
	for _, obj := range testData {
		_, err = api.CreateClusterInst(ctx, &obj)
		assert.Nil(t, err, "Create ClusterInst %s", obj.Key.GetKeyString())
	}
	_, err = api.CreateClusterInst(ctx, &testData[0])
	assert.NotNil(t, err, "Create duplicate ClusterInst")

	// test show all items
	show := ShowClusterInst{}
	show.Init()
	filterNone := edgeproto.ClusterInst{}
	err = api.ShowClusterInst(ctx, &filterNone, &show)
	assert.Nil(t, err, "show data")
	for _, obj := range testData {
		show.AssertFound(t, &obj)
	}
	assert.Equal(t, len(testData), len(show.Data), "Show count")

	// test delete
	_, err = api.DeleteClusterInst(ctx, &testData[0])
	assert.Nil(t, err, "delete ClusterInst %s", testData[0].Key.GetKeyString())
	show.Init()
	err = api.ShowClusterInst(ctx, &filterNone, &show)
	assert.Nil(t, err, "show data")
	assert.Equal(t, len(testData)-1, len(show.Data), "Show count")
	show.AssertNotFound(t, &testData[0])
	// test update of missing object
	_, err = api.UpdateClusterInst(ctx, &testData[0])
	assert.NotNil(t, err, "Update missing object")
	// create it back
	_, err = api.CreateClusterInst(ctx, &testData[0])
	assert.Nil(t, err, "Create ClusterInst %s", testData[0].Key.GetKeyString())

	// test invalid keys
	bad := edgeproto.ClusterInst{}
	_, err = api.CreateClusterInst(ctx, &bad)
	assert.NotNil(t, err, "Create ClusterInst with no key info")

}
