/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cluster "github.com/edixos/provider-ovh/internal/controller/cluster/kube/cluster"
	iprestriction "github.com/edixos/provider-ovh/internal/controller/cluster/kube/iprestriction"
	nodepool "github.com/edixos/provider-ovh/internal/controller/cluster/kube/nodepool"
	oidcconfiguration "github.com/edixos/provider-ovh/internal/controller/cluster/kube/oidcconfiguration"
)

// Setup_kube creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_kube(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		iprestriction.Setup,
		nodepool.Setup,
		oidcconfiguration.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_kube creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_kube(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.SetupGated,
		iprestriction.SetupGated,
		nodepool.SetupGated,
		oidcconfiguration.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
