// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/config/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterImagePolicyLister helps list ClusterImagePolicies.
// All objects returned here must be treated as read-only.
type ClusterImagePolicyLister interface {
	// List lists all ClusterImagePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterImagePolicy, err error)
	// Get retrieves the ClusterImagePolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterImagePolicy, error)
	ClusterImagePolicyListerExpansion
}

// clusterImagePolicyLister implements the ClusterImagePolicyLister interface.
type clusterImagePolicyLister struct {
	listers.ResourceIndexer[*v1alpha1.ClusterImagePolicy]
}

// NewClusterImagePolicyLister returns a new ClusterImagePolicyLister.
func NewClusterImagePolicyLister(indexer cache.Indexer) ClusterImagePolicyLister {
	return &clusterImagePolicyLister{listers.New[*v1alpha1.ClusterImagePolicy](indexer, v1alpha1.Resource("clusterimagepolicy"))}
}