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

package testutil

// Stubs for DummyServer.
// Revisit as needed for unit tests.
import (
	"context"
	"io"

	"github.com/mobiledgex/edge-cloud/edgeproto"
	"google.golang.org/grpc"
)

func (s *DummyServer) AddCloudletResMapping(ctx context.Context, in *edgeproto.CloudletResMap) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) RemoveCloudletResMapping(ctx context.Context, in *edgeproto.CloudletResMap) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) AddCloudletAllianceOrg(ctx context.Context, in *edgeproto.CloudletAllianceOrg) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) RemoveCloudletAllianceOrg(ctx context.Context, in *edgeproto.CloudletAllianceOrg) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) FindFlavorMatch(ctx context.Context, in *edgeproto.FlavorMatch) (*edgeproto.FlavorMatch, error) {
	return in, nil
}

func (s *DummyServer) GetCloudletProps(ctx context.Context, in *edgeproto.CloudletProps) (*edgeproto.CloudletProps, error) {
	return in, nil
}

func (s *DummyServer) RevokeAccessKey(ctx context.Context, key *edgeproto.CloudletKey) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) GenerateAccessKey(ctx context.Context, key *edgeproto.CloudletKey) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

func (s *DummyServer) UpgradeAccessKey(stream edgeproto.CloudletAccessKeyApi_UpgradeAccessKeyServer) error {
	return nil
}

func (s *DummyServer) GetCloudletManifest(ctx context.Context, key *edgeproto.CloudletKey) (*edgeproto.CloudletManifest, error) {
	return &edgeproto.CloudletManifest{}, nil
}

func (s *DummyServer) GetCloudletResourceUsage(ctx context.Context, usage *edgeproto.CloudletResourceUsage) (*edgeproto.CloudletResourceUsage, error) {
	return &edgeproto.CloudletResourceUsage{}, nil
}

func (s *DummyServer) GetCloudletResourceQuotaProps(ctx context.Context, in *edgeproto.CloudletResourceQuotaProps) (*edgeproto.CloudletResourceQuotaProps, error) {
	return &edgeproto.CloudletResourceQuotaProps{}, nil
}

func (s *DummyServer) GetOrganizationsOnCloudlet(in *edgeproto.CloudletKey, cb edgeproto.CloudletApi_GetOrganizationsOnCloudletServer) error {
	orgs := s.OrgsOnCloudlet[*in]
	for _, org := range orgs {
		eorg := edgeproto.Organization{
			Name: org,
		}
		cb.Send(&eorg)
	}
	return nil
}

func (s *DummyServer) GetCloudletGPUDriverLicenseConfig(ctx context.Context, in *edgeproto.CloudletKey) (*edgeproto.Result, error) {
	return &edgeproto.Result{}, nil
}

// minimal bits not currently generated for flavorkey.proto to stream flavorKey objs
// for ShowFlavorsForCloudlet cli
type ShowFlavorsForCloudlet struct {
	Data map[string]edgeproto.FlavorKey
	grpc.ServerStream
	Ctx context.Context
}

func (x *ShowFlavorsForCloudlet) Init() {
	x.Data = make(map[string]edgeproto.FlavorKey)
}

func (x *ShowFlavorsForCloudlet) Send(m *edgeproto.FlavorKey) error {
	x.Data[m.Name] = *m
	return nil
}

func (x *ShowFlavorsForCloudlet) Context() context.Context {
	return x.Ctx
}

var ShowFlavorsForCloudletExtraCount = 0

func (x *ShowFlavorsForCloudlet) ReadStream(stream edgeproto.CloudletApi_ShowFlavorsForCloudletClient, err error) {

	x.Data = make(map[string]edgeproto.FlavorKey)
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
		x.Data[obj.Name] = *obj
	}
}

func (x *CloudletCommonApi) ShowFlavorsForCloudlet(ctx context.Context, filter *edgeproto.CloudletKey, showData *ShowFlavorsForCloudlet) error {

	if x.internal_api != nil {
		showData.Ctx = ctx
		return x.internal_api.ShowFlavorsForCloudlet(filter, showData)
	} else {

		stream, err := x.client_api.ShowFlavorsForCloudlet(ctx, filter)
		showData.ReadStream(stream, err)
		return err
	}
}
