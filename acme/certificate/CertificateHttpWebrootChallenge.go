// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate


type CertificateHttpWebrootChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.34.0/docs/resources/certificate#directory Certificate#directory}.
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

