/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	project "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/project"
	projectregionloadbalancerlogsubscription "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/projectregionloadbalancerlogsubscription"
	s3credentials "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/s3credentials"
	s3policy "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/s3policy"
	user "github.com/edixos/provider-ovh/internal/controller/namespaced/cloud/user"
)

// Setup_cloud creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cloud(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.Setup,
		projectregionloadbalancerlogsubscription.Setup,
		s3credentials.Setup,
		s3policy.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cloud creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cloud(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		project.SetupGated,
		projectregionloadbalancerlogsubscription.SetupGated,
		s3credentials.SetupGated,
		s3policy.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
