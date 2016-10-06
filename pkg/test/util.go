package test

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	apitesting "k8s.io/kubernetes/pkg/api/testing"
)

func NodeExample(name string) api.Node {
	return api.Node{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: "test", ResourceVersion: "123"},
		Spec: api.NodeSpec{
			ExternalID: "ext",
		},
	}
}

func PodExample(name string) api.Pod {
	return api.Pod{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: "test", ResourceVersion: "10"},
		Spec:       apitesting.DeepEqualSafePodSpec(),
	}
}

func ServiceExample(name string) api.Service {
	return api.Service{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: "test", ResourceVersion: "12"},
		Spec: api.ServiceSpec{
			SessionAffinity: "None",
			Type:            api.ServiceTypeClusterIP,
		},
	}
}

func ReplicationControllerExample(name string) api.ReplicationController {
	return api.ReplicationController{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: "test", ResourceVersion: "18"},
		Spec: api.ReplicationControllerSpec{
			Replicas: 1,
		},
	}
}
func PersistentVolumeExample(name string) api.PersistentVolume {
	return api.PersistentVolume{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: name, ResourceVersion: "123"},
		Spec: api.PersistentVolumeSpec{
			Capacity: api.ResourceList{
				api.ResourceName(api.ResourceStorage): resource.MustParse("10G"),
			},
			PersistentVolumeSource: api.PersistentVolumeSource{
				HostPath: &api.HostPathVolumeSource{Path: "/foo"},
			},
			PersistentVolumeReclaimPolicy: "Retain",
		},
		Status: api.PersistentVolumeStatus{
			Phase: api.PersistentVolumePhase("Pending"),
		},
	}
}

func PersistentVolumeClaimExample(name string) api.PersistentVolumeClaim {
	return api.PersistentVolumeClaim{
		ObjectMeta: api.ObjectMeta{Name: name, Namespace: "test", ResourceVersion: "123"},
		Spec: api.PersistentVolumeClaimSpec{
			VolumeName: "volume",
		},
		Status: api.PersistentVolumeClaimStatus{
			Phase: api.PersistentVolumeClaimPhase("Pending"),
		},
	}
}
