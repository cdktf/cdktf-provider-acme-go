// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registration


type RegistrationExternalAccountBinding struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.34.0/docs/resources/registration#hmac_base64 Registration#hmac_base64}.
	HmacBase64 *string `field:"required" json:"hmacBase64" yaml:"hmacBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.34.0/docs/resources/registration#key_id Registration#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

