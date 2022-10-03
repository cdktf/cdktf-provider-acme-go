package certificate


type CertificateHttpMemcachedChallenge struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#hosts Certificate#hosts}.
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
}

