// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate


type CertificateHttpMemcachedChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.32.1/docs/resources/certificate#hosts Certificate#hosts}.
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
}

