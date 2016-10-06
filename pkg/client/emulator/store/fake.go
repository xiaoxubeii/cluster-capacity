package store

import (
	"fmt"
	"reflect"

	ccapi "github.com/ingvagabund/cluster-capacity/pkg/api"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	"k8s.io/kubernetes/pkg/client/cache"
)

type FakeResourceStore struct {
	PodsData                   func() []api.Pod
	ServicesData               func() []api.Service
	ReplicationControllersData func() []api.ReplicationController
	NodesData                  func() []api.Node
	PersistentVolumesData      func() []api.PersistentVolume
	PersistentVolumeClaimsData func() []api.PersistentVolumeClaim
	// TODO(jchaloup): fill missing resource functions
}

func (s *FakeResourceStore) Add(resource string, obj interface{}) error {
	return nil
}

func (s *FakeResourceStore) Update(resource string, obj interface{}) error {
	return nil
}

func (s *FakeResourceStore) Delete(resource string, obj interface{}) error {
	return nil
}

func resourcesToItems(objs interface{}) []interface{} {
	objsSlice := reflect.ValueOf(objs)
	items := make([]interface{}, 0, objsSlice.Len())
	for i := 0; i < objsSlice.Len(); i++ {
		items = append(items, objsSlice.Index(i).Interface())
	}
	return items
}

func findResource(obj interface{}, objs interface{}) (item interface{}, exists bool, err error) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		return nil, false, err
	}

	var obj_key string
	var key_err error
	objsSlice := reflect.ValueOf(objs)
	for i := 0; i < objsSlice.Len(); i++ {
		item := objsSlice.Index(i).Interface()
		// TODO(jchaloup): make this resource type independent
		switch item.(type) {
		case api.Pod:
			value := item.(api.Pod)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		case api.Service:
			value := item.(api.Service)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		case api.ReplicationController:
			value := item.(api.ReplicationController)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		case api.Node:
			value := item.(api.Node)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		case api.PersistentVolume:
			value := item.(api.PersistentVolume)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		case api.PersistentVolumeClaim:
			value := item.(api.PersistentVolumeClaim)
			obj_key, key_err = cache.MetaNamespaceKeyFunc(meta.Object(&value))
		}
		if key_err != nil {
			return nil, false, key_err
		}
		if obj_key == key {
			return item, true, nil
		}
	}
	return nil, false, fmt.Errorf("Resource obj not found")
}

func (s *FakeResourceStore) List(resource string) []interface{} {
	switch resource {
	case ccapi.Pods:
		return resourcesToItems(s.PodsData())
	case ccapi.Services:
		return resourcesToItems(s.ServicesData())
	case ccapi.ReplicationControllers:
		return resourcesToItems(s.ReplicationControllersData())
	case ccapi.Nodes:
		return resourcesToItems(s.NodesData())
	case ccapi.PersistentVolumes:
		return resourcesToItems(s.PersistentVolumesData())
	case ccapi.PersistentVolumeClaims:
		return resourcesToItems(s.PersistentVolumeClaimsData())
		//case "replicasets":
		//	return testReplicaSetsData().Items
	}
	return nil
}

func (s *FakeResourceStore) Get(resource string, obj interface{}) (item interface{}, exists bool, err error) {
	switch resource {
	case ccapi.Pods:
		return findResource(obj, s.PodsData())
	case ccapi.Services:
		return findResource(obj, s.ServicesData())
	case ccapi.ReplicationControllers:
		return findResource(obj, s.ReplicationControllersData())
	case ccapi.Nodes:
		return findResource(obj, s.NodesData())
	case ccapi.PersistentVolumes:
		return findResource(obj, s.PersistentVolumesData())
	case ccapi.PersistentVolumeClaims:
		return findResource(obj, s.PersistentVolumeClaimsData())
		//case "replicasets":
		//	return testReplicaSetsData().Items
	}
	return nil, false, nil
}

func (s *FakeResourceStore) GetByKey(key string) (item interface{}, exists bool, err error) {
	return nil, false, nil
}

func (s *FakeResourceStore) RegisterEventHandler(resource string, handler cache.ResourceEventHandler) error {
	return nil
}

func (s *FakeResourceStore) Replace(resource string, items []interface{}, resourceVersion string) error {
	return nil
}

func (s *FakeResourceStore) Resources() []string {
	return []string{ccapi.Pods, ccapi.Services, ccapi.ReplicationControllers, ccapi.Nodes, ccapi.PersistentVolumes, ccapi.PersistentVolumeClaims}
}
