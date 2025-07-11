// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package registration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RegistrationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#email_address Registration#email_address}.
	EmailAddress *string `field:"required" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#account_key_algorithm Registration#account_key_algorithm}.
	AccountKeyAlgorithm *string `field:"optional" json:"accountKeyAlgorithm" yaml:"accountKeyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#account_key_ecdsa_curve Registration#account_key_ecdsa_curve}.
	AccountKeyEcdsaCurve *string `field:"optional" json:"accountKeyEcdsaCurve" yaml:"accountKeyEcdsaCurve"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#account_key_pem Registration#account_key_pem}.
	AccountKeyPem *string `field:"optional" json:"accountKeyPem" yaml:"accountKeyPem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#account_key_rsa_bits Registration#account_key_rsa_bits}.
	AccountKeyRsaBits *float64 `field:"optional" json:"accountKeyRsaBits" yaml:"accountKeyRsaBits"`
	// external_account_binding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#external_account_binding Registration#external_account_binding}
	ExternalAccountBinding *RegistrationExternalAccountBinding `field:"optional" json:"externalAccountBinding" yaml:"externalAccountBinding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.33.0/docs/resources/registration#id Registration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

