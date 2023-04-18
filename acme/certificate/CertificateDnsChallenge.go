package certificate


type CertificateDnsChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.13.1/docs/resources/certificate#provider Certificate#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.13.1/docs/resources/certificate#config Certificate#config}.
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
}

