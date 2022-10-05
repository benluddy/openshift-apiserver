package etcd

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	authorizationclient "k8s.io/client-go/kubernetes/typed/authorization/v1"
	"k8s.io/kubernetes/pkg/printers"
	printerstorage "k8s.io/kubernetes/pkg/printers/storage"

	"github.com/openshift/api/template"

	templateapi "github.com/openshift/openshift-apiserver/pkg/template/apis/template"
	"github.com/openshift/openshift-apiserver/pkg/template/apiserver/registry/templateinstance"
	templateprinters "github.com/openshift/openshift-apiserver/pkg/template/printers/internalversion"
)

// REST implements a RESTStorage for templateinstances against etcd
type REST struct {
	*registry.Store
}

var _ rest.StandardStorage = &REST{}

// NewREST returns a RESTStorage object that will work against templateinstances.
func NewREST(optsGetter generic.RESTOptionsGetter, authorizationClient authorizationclient.AuthorizationV1Interface) (*REST, *StatusREST, error) {
	strategy := templateinstance.NewStrategy(authorizationClient)

	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &templateapi.TemplateInstance{} },
		NewListFunc:              func() runtime.Object { return &templateapi.TemplateInstanceList{} },
		DefaultQualifiedResource: template.Resource("templateinstances"),

		TableConvertor: printerstorage.TableConvertor{TableGenerator: printers.NewTableGenerator().With(templateprinters.AddTemplateOpenShiftHandlers)},

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}

	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, nil, err
	}

	statusStore := *store
	statusStore.UpdateStrategy = templateinstance.StatusStrategy

	return &REST{store}, &StatusREST{&statusStore}, nil
}

// StatusREST implements the REST endpoint for changing the status of a templateInstance.
type StatusREST struct {
	*registry.Store
}

// StatusREST implements Patcher
var _ = rest.Patcher(&StatusREST{})

// New creates a new templateInstance resource
func (r *StatusREST) New() runtime.Object {
	return &templateapi.TemplateInstance{}
}
