package certificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CertificateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#account_key_pem Certificate#account_key_pem}.
	AccountKeyPem *string `field:"required" json:"accountKeyPem" yaml:"accountKeyPem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#certificate_p12_password Certificate#certificate_p12_password}.
	CertificateP12Password *string `field:"optional" json:"certificateP12Password" yaml:"certificateP12Password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#certificate_request_pem Certificate#certificate_request_pem}.
	CertificateRequestPem *string `field:"optional" json:"certificateRequestPem" yaml:"certificateRequestPem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#common_name Certificate#common_name}.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#disable_complete_propagation Certificate#disable_complete_propagation}.
	DisableCompletePropagation interface{} `field:"optional" json:"disableCompletePropagation" yaml:"disableCompletePropagation"`
	// dns_challenge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#dns_challenge Certificate#dns_challenge}
	DnsChallenge interface{} `field:"optional" json:"dnsChallenge" yaml:"dnsChallenge"`
	// http_challenge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#http_challenge Certificate#http_challenge}
	HttpChallenge *CertificateHttpChallenge `field:"optional" json:"httpChallenge" yaml:"httpChallenge"`
	// http_memcached_challenge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#http_memcached_challenge Certificate#http_memcached_challenge}
	HttpMemcachedChallenge *CertificateHttpMemcachedChallenge `field:"optional" json:"httpMemcachedChallenge" yaml:"httpMemcachedChallenge"`
	// http_webroot_challenge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#http_webroot_challenge Certificate#http_webroot_challenge}
	HttpWebrootChallenge *CertificateHttpWebrootChallenge `field:"optional" json:"httpWebrootChallenge" yaml:"httpWebrootChallenge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#id Certificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#key_type Certificate#key_type}.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#min_days_remaining Certificate#min_days_remaining}.
	MinDaysRemaining *float64 `field:"optional" json:"minDaysRemaining" yaml:"minDaysRemaining"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#must_staple Certificate#must_staple}.
	MustStaple interface{} `field:"optional" json:"mustStaple" yaml:"mustStaple"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#pre_check_delay Certificate#pre_check_delay}.
	PreCheckDelay *float64 `field:"optional" json:"preCheckDelay" yaml:"preCheckDelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#preferred_chain Certificate#preferred_chain}.
	PreferredChain *string `field:"optional" json:"preferredChain" yaml:"preferredChain"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#recursive_nameservers Certificate#recursive_nameservers}.
	RecursiveNameservers *[]*string `field:"optional" json:"recursiveNameservers" yaml:"recursiveNameservers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#revoke_certificate_on_destroy Certificate#revoke_certificate_on_destroy}.
	RevokeCertificateOnDestroy interface{} `field:"optional" json:"revokeCertificateOnDestroy" yaml:"revokeCertificateOnDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#subject_alternative_names Certificate#subject_alternative_names}.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// tls_challenge block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/acme/r/certificate#tls_challenge Certificate#tls_challenge}
	TlsChallenge *CertificateTlsChallenge `field:"optional" json:"tlsChallenge" yaml:"tlsChallenge"`
}

