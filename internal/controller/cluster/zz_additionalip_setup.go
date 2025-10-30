/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	firewall "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/firewall"
	firewallrule "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/firewallrule"
	mitigation "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/mitigation"
	move "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/move"
	projectfailoveripattach "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/projectfailoveripattach"
	reverse "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/reverse"
	service "github.com/edixos/provider-ovh/internal/controller/cluster/additionalip/service"
)

// Setup_additionalip creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_additionalip(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.Setup,
		firewallrule.Setup,
		mitigation.Setup,
		move.Setup,
		projectfailoveripattach.Setup,
		reverse.Setup,
		service.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_additionalip creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_additionalip(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		firewall.SetupGated,
		firewallrule.SetupGated,
		mitigation.SetupGated,
		move.SetupGated,
		projectfailoveripattach.SetupGated,
		reverse.SetupGated,
		service.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
