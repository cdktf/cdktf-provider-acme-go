package certificate


type CertificateHttpWebrootChallenge struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#directory Certificate#directory}.
	Directory *string `field:"required" json:"directory" yaml:"directory"`
}

