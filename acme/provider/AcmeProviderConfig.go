// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AcmeProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.28.1/docs#server_url AcmeProvider#server_url}.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.28.1/docs#alias AcmeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

