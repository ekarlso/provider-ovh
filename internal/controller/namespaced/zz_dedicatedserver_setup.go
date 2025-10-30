/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	servernetworking "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/servernetworking"
	serverreboottask "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/serverreboottask"
	serverupdate "github.com/edixos/provider-ovh/internal/controller/namespaced/dedicatedserver/serverupdate"
)

// Setup_dedicatedserver creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dedicatedserver(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		servernetworking.Setup,
		serverreboottask.Setup,
		serverupdate.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dedicatedserver creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dedicatedserver(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		servernetworking.SetupGated,
		serverreboottask.SetupGated,
		serverupdate.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
