/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	logscluster "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logscluster"
	logsinput "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logsinput"
	logstoken "github.com/edixos/provider-ovh/internal/controller/cluster/logs/logstoken"
)

// Setup_logs creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_logs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		logscluster.Setup,
		logsinput.Setup,
		logstoken.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_logs creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_logs(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		logscluster.SetupGated,
		logsinput.SetupGated,
		logstoken.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
