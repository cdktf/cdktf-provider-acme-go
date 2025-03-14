// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate


type CertificateTlsChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.30.2/docs/resources/certificate#port Certificate#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

