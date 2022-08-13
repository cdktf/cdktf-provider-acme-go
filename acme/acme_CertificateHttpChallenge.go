// Prebuilt acme Provider for Terraform CDK (cdktf)
package acme


type CertificateHttpChallenge struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#port Certificate#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#proxy_header Certificate#proxy_header}.
	ProxyHeader *string `field:"optional" json:"proxyHeader" yaml:"proxyHeader"`
}

