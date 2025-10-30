/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	iampermissionsgroup "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iampermissionsgroup"
	iampolicy "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iampolicy"
	iamresourcegroup "github.com/edixos/provider-ovh/internal/controller/namespaced/iam/iamresourcegroup"
)

// Setup_iam creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_iam(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		iampermissionsgroup.Setup,
		iampolicy.Setup,
		iamresourcegroup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_iam creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_iam(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		iampermissionsgroup.SetupGated,
		iampolicy.SetupGated,
		iamresourcegroup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
