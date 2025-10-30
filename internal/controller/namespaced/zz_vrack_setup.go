/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	cloudproject "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/cloudproject"
	dedicatedserver "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/dedicatedserver"
	dedicatedserverinterface "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/dedicatedserverinterface"
	ip "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/ip"
	iploadbalancing "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/iploadbalancing"
	vrack "github.com/edixos/provider-ovh/internal/controller/namespaced/vrack/vrack"
)

// Setup_vrack creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vrack(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudproject.Setup,
		dedicatedserver.Setup,
		dedicatedserverinterface.Setup,
		ip.Setup,
		iploadbalancing.Setup,
		vrack.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_vrack creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_vrack(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cloudproject.SetupGated,
		dedicatedserver.SetupGated,
		dedicatedserverinterface.SetupGated,
		ip.SetupGated,
		iploadbalancing.SetupGated,
		vrack.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
