/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	group "github.com/edixos/provider-ovh/internal/controller/cluster/me/group"
	oauth2client "github.com/edixos/provider-ovh/internal/controller/cluster/me/oauth2client"
	user "github.com/edixos/provider-ovh/internal/controller/cluster/me/user"
)

// Setup_me creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_me(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.Setup,
		oauth2client.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_me creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_me(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		group.SetupGated,
		oauth2client.SetupGated,
		user.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
