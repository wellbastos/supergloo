// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionInstallFunc func(original, desired *Install) (bool, error)

type InstallReconciler interface {
	Reconcile(namespace string, desiredResources InstallList, transition TransitionInstallFunc, opts clients.ListOpts) error
}

func installsToResources(list InstallList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, install := range list {
		resourceList = append(resourceList, install)
	}
	return resourceList
}

func NewInstallReconciler(client InstallClient) InstallReconciler {
	return &installReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type installReconciler struct {
	base reconcile.Reconciler
}

func (r *installReconciler) Reconcile(namespace string, desiredResources InstallList, transition TransitionInstallFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "install_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*Install), desired.(*Install))
		}
	}
	return r.base.Reconcile(namespace, installsToResources(desiredResources), transitionResources, opts)
}
