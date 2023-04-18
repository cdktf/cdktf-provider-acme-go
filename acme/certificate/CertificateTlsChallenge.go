package certificate


type CertificateTlsChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.13.1/docs/resources/certificate#port Certificate#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

