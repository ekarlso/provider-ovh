// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

// nolint:misspell
var (
	// provider-azure uses the OVH service names extracted
	// from Terraform import statements, e.g.:
	// this table holds only overrides for the resources.
	resourceAPIGroupMap = map[string]string{}

	// this table holds overrides of OVH provider API groups
	// for example API group extracted from OVH provider "DocumentDB"
	// is overridden as "cosmodb".
	apiGroupOverrides = map[string]string{}
)
