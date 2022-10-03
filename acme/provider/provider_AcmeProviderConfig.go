package provider


type AcmeProviderConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme#server_url AcmeProvider#server_url}.
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme#alias AcmeProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

