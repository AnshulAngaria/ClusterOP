package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterOP struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec ClusterSpec
}

type ClusterSpec struct {
	Name      string
	Version   string
	Region    string
	NodePools []NodePool
}

type NodePool struct {
	Size  string
	Name  string
	Count int
}

type ClusterList struct {
}
