package certificate


type CertificateDnsChallenge struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#provider Certificate#provider}.
	Provider *string `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#config Certificate#config}.
	Config *map[string]*string `field:"optional" json:"config" yaml:"config"`
}

