/*
Copyright 2021 Upbound Inc.
*/

package cluster

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/edixos/provider-ovh/config/cluster/gateway"
	"github.com/edixos/provider-ovh/config/cluster/vps"

	"github.com/edixos/provider-ovh/config/cluster/cloud"
	"github.com/edixos/provider-ovh/config/cluster/iam"
	"github.com/edixos/provider-ovh/config/cluster/me"

	"github.com/edixos/provider-ovh/config/cluster/vrack"

	"github.com/edixos/provider-ovh/config/cluster/web_cloud_private_sql"

	"github.com/edixos/provider-ovh/config/cluster/object_storage"

	"github.com/edixos/provider-ovh/config/cluster/nas"

	"github.com/edixos/provider-ovh/config/cluster/registry"

	"github.com/edixos/provider-ovh/config/cluster/kube"

	"github.com/edixos/provider-ovh/config/cluster/databases"

	"github.com/edixos/provider-ovh/config/cluster/logs"

	"github.com/edixos/provider-ovh/config/cluster/lb"

	"github.com/edixos/provider-ovh/config/cluster/dns"

	"github.com/edixos/provider-ovh/config/cluster/dedicated_server"

	"github.com/edixos/provider-ovh/config/cluster/cloud_disk_array"

	"github.com/edixos/provider-ovh/config/cluster/additional_ip"
	"github.com/edixos/provider-ovh/config/cluster/vm_instances"

	"github.com/edixos/provider-ovh/config/cluster/public_cloud_network"
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
