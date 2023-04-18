package certificate


type CertificateHttpMemcachedChallenge struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/vancluever/acme/2.13.1/docs/resources/certificate#hosts Certificate#hosts}.
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
}

