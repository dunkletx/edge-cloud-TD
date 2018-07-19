package main

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mobiledgex/edge-cloud/edgeproto"
)

type ClusterApi struct {
	sync  *Sync
	store edgeproto.ClusterStore
	cache edgeproto.ClusterCache
}

var clusterApi = ClusterApi{}

const ClusterAutoPrefix = "AutoCluster"
const ClusterAutoNodes = 3

var ClusterAutoPrefixErr = fmt.Sprintf("Cluster name prefix \"%s\" is reserved",
	ClusterAutoPrefix)

func InitClusterApi(sync *Sync) {
	clusterApi.sync = sync
	clusterApi.store = edgeproto.NewClusterStore(sync.store)
	edgeproto.InitClusterCache(&clusterApi.cache)
	sync.RegisterCache(&clusterApi.cache)
}

func (s *ClusterApi) UsesFlavor(key *edgeproto.FlavorKey) bool {
	s.cache.Mux.Lock()
	defer s.cache.Mux.Unlock()
	for _, cluster := range s.cache.Objs {
		if cluster.Flavor.Matches(key) {
			return true
		}
	}
	return false
}

func (s *ClusterApi) HasKey(key *edgeproto.ClusterKey) bool {
	return s.cache.HasKey(key)
}

func (s *ClusterApi) Get(key *edgeproto.ClusterKey, buf *edgeproto.Cluster) bool {
	return s.cache.Get(key, buf)
}

func (s *ClusterApi) CreateCluster(ctx context.Context, in *edgeproto.Cluster) (*edgeproto.Result, error) {
	if strings.HasPrefix(in.Key.Name, ClusterAutoPrefix) {
		return &edgeproto.Result{}, errors.New(ClusterAutoPrefixErr)
	}
	return s.createClusterInternal(in)
}

func (s *ClusterApi) createClusterInternal(in *edgeproto.Cluster) (*edgeproto.Result, error) {
	return s.store.Create(in, s.sync.syncWait)
}

func (s *ClusterApi) UpdateCluster(ctx context.Context, in *edgeproto.Cluster) (*edgeproto.Result, error) {
	// Unsupported for now
	return &edgeproto.Result{}, errors.New("Update cluster not supported")
	//return s.store.Update(in, s.sync.syncWait)
}

func (s *ClusterApi) DeleteCluster(ctx context.Context, in *edgeproto.Cluster) (*edgeproto.Result, error) {
	return s.deleteClusterInternal(in)
}

func (s *ClusterApi) deleteClusterInternal(in *edgeproto.Cluster) (*edgeproto.Result, error) {
	if appApi.UsesCluster(&in.Key) {
		return &edgeproto.Result{}, errors.New("Cluster in use by Application")
	}
	return s.store.Delete(in, s.sync.syncWait)
}

func (s *ClusterApi) ShowCluster(in *edgeproto.Cluster, cb edgeproto.ClusterApi_ShowClusterServer) error {
	err := s.cache.Show(in, func(obj *edgeproto.Cluster) error {
		err := cb.Send(obj)
		return err
	})
	return err
}
