/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cephacl "github.com/edixos/provider-ovh/internal/controller/cluster/clouddiskarray/cephacl"
)

// Setup_clouddiskarray creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_clouddiskarray(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cephacl.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_clouddiskarray creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_clouddiskarray(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cephacl.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
