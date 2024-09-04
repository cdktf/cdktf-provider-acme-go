// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate


type CertificateHttpChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.26.0/docs/resources/certificate#port Certificate#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.26.0/docs/resources/certificate#proxy_header Certificate#proxy_header}.
	ProxyHeader *string `field:"optional" json:"proxyHeader" yaml:"proxyHeader"`
}

