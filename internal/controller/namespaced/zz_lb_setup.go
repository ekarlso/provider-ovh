/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	httpfarm "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfarm"
	httpfarmserver "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfarmserver"
	httpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httpfrontend"
	httproute "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httproute"
	httprouterule "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/httprouterule"
	iploadbalancing "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/iploadbalancing"
	refresh "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/refresh"
	tcpfarm "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfarm"
	tcpfarmserver "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfarmserver"
	tcpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcpfrontend"
	tcproute "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcproute"
	tcprouterule "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/tcprouterule"
	udpfrontend "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/udpfrontend"
	vracknetwork "github.com/edixos/provider-ovh/internal/controller/namespaced/lb/vracknetwork"
)

// Setup_lb creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_lb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		httpfarm.Setup,
		httpfarmserver.Setup,
		httpfrontend.Setup,
		httproute.Setup,
		httprouterule.Setup,
		iploadbalancing.Setup,
		refresh.Setup,
		tcpfarm.Setup,
		tcpfarmserver.Setup,
		tcpfrontend.Setup,
		tcproute.Setup,
		tcprouterule.Setup,
		udpfrontend.Setup,
		vracknetwork.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_lb creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_lb(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		httpfarm.SetupGated,
		httpfarmserver.SetupGated,
		httpfrontend.SetupGated,
		httproute.SetupGated,
		httprouterule.SetupGated,
		iploadbalancing.SetupGated,
		refresh.SetupGated,
		tcpfarm.SetupGated,
		tcpfarmserver.SetupGated,
		tcpfrontend.SetupGated,
		tcproute.SetupGated,
		tcprouterule.SetupGated,
		udpfrontend.SetupGated,
		vracknetwork.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
