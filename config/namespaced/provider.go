/*
Copyright 2021 Upbound Inc.
*/

package namespaced

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/edixos/provider-ovh/config/namespaced/gateway"
	"github.com/edixos/provider-ovh/config/namespaced/vps"

	"github.com/edixos/provider-ovh/config/namespaced/cloud"
	"github.com/edixos/provider-ovh/config/namespaced/iam"
	"github.com/edixos/provider-ovh/config/namespaced/me"

	"github.com/edixos/provider-ovh/config/namespaced/vrack"

	"github.com/edixos/provider-ovh/config/namespaced/web_cloud_private_sql"

	"github.com/edixos/provider-ovh/config/namespaced/object_storage"

	"github.com/edixos/provider-ovh/config/namespaced/nas"

	"github.com/edixos/provider-ovh/config/namespaced/registry"

	"github.com/edixos/provider-ovh/config/namespaced/kube"

	"github.com/edixos/provider-ovh/config/namespaced/databases"

	"github.com/edixos/provider-ovh/config/namespaced/logs"

	"github.com/edixos/provider-ovh/config/namespaced/lb"

	"github.com/edixos/provider-ovh/config/namespaced/dns"

	"github.com/edixos/provider-ovh/config/namespaced/dedicated_server"

	"github.com/edixos/provider-ovh/config/namespaced/cloud_disk_array"

	"github.com/edixos/provider-ovh/config/namespaced/additional_ip"
	"github.com/edixos/provider-ovh/config/namespaced/vm_instances"

	"github.com/edixos/provider-ovh/config/namespaced/public_cloud_network"
)

func init() {
	ProviderConfiguration.AddConfig(gateway.Configure)
	ProviderConfiguration.AddConfig(vps.Configure)
	ProviderConfiguration.AddConfig(cloud.Configure)
	ProviderConfiguration.AddConfig(iam.Configure)
	ProviderConfiguration.AddConfig(me.Configure)
	ProviderConfiguration.AddConfig(vrack.Configure)
	ProviderConfiguration.AddConfig(web_cloud_private_sql.Configure)
	ProviderConfiguration.AddConfig(object_storage.Configure)
	ProviderConfiguration.AddConfig(nas.Configure)
	ProviderConfiguration.AddConfig(registry.Configure)
	ProviderConfiguration.AddConfig(kube.Configure)
	ProviderConfiguration.AddConfig(databases.Configure)
	ProviderConfiguration.AddConfig(logs.Configure)
	ProviderConfiguration.AddConfig(lb.Configure)
	ProviderConfiguration.AddConfig(dns.Configure)
	ProviderConfiguration.AddConfig(dedicated_server.Configure)
	ProviderConfiguration.AddConfig(cloud_disk_array.Configure)
	ProviderConfiguration.AddConfig(additional_ip.Configure)
	ProviderConfiguration.AddConfig(vm_instances.Configure)
	ProviderConfiguration.AddConfig(public_cloud_network.Configure)
}
