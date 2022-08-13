// Prebuilt acme Provider for Terraform CDK (cdktf)
package acme


type RegistrationExternalAccountBinding struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/registration#hmac_base64 Registration#hmac_base64}.
	HmacBase64 *string `field:"required" json:"hmacBase64" yaml:"hmacBase64"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/registration#key_id Registration#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

