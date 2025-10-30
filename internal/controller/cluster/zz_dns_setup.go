/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	zone "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zone"
	zonednssec "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonednssec"
	zonerecord "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zonerecord"
	zoneredirection "github.com/edixos/provider-ovh/internal/controller/cluster/dns/zoneredirection"
)

// Setup_dns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		zone.Setup,
		zonednssec.Setup,
		zonerecord.Setup,
		zoneredirection.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dns creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		zone.SetupGated,
		zonednssec.SetupGated,
		zonerecord.SetupGated,
		zoneredirection.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
