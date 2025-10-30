/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	projectworkflowbackup "github.com/edixos/provider-ovh/internal/controller/cluster/vminstances/projectworkflowbackup"
)

// Setup_vminstances creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vminstances(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectworkflowbackup.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_vminstances creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_vminstances(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectworkflowbackup.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
