// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-acme-go/acme/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-acme-go/acme/v10/certificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/vancluever/acme/2.19.0/docs/resources/certificate acme_certificate}.
type Certificate interface {
	cdktf.TerraformResource
	AccountKeyPem() *string
	SetAccountKeyPem(val *string)
	AccountKeyPemInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateDomain() *string
	CertificateNotAfter() *string
	CertificateP12() *string
	CertificateP12Password() *string
	SetCertificateP12Password(val *string)
	CertificateP12PasswordInput() *string
	CertificatePem() *string
	CertificateRequestPem() *string
	SetCertificateRequestPem(val *string)
	CertificateRequestPemInput() *string
	CertificateUrl() *string
	CertTimeout() *float64
	SetCertTimeout(val *float64)
	CertTimeoutInput() *float64
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisableCompletePropagation() interface{}
	SetDisableCompletePropagation(val interface{})
	DisableCompletePropagationInput() interface{}
	DnsChallenge() CertificateDnsChallengeList
	DnsChallengeInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpChallenge() CertificateHttpChallengeOutputReference
	HttpChallengeInput() *CertificateHttpChallenge
	HttpMemcachedChallenge() CertificateHttpMemcachedChallengeOutputReference
	HttpMemcachedChallengeInput() *CertificateHttpMemcachedChallenge
	HttpS3Challenge() CertificateHttpS3ChallengeOutputReference
	HttpS3ChallengeInput() *CertificateHttpS3Challenge
	HttpWebrootChallenge() CertificateHttpWebrootChallengeOutputReference
	HttpWebrootChallengeInput() *CertificateHttpWebrootChallenge
	Id() *string
	SetId(val *string)
	IdInput() *string
	IssuerPem() *string
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MinDaysRemaining() *float64
	SetMinDaysRemaining(val *float64)
	MinDaysRemainingInput() *float64
	MustStaple() interface{}
	SetMustStaple(val interface{})
	MustStapleInput() interface{}
	// The tree node.
	Node() constructs.Node
	PreCheckDelay() *float64
	SetPreCheckDelay(val *float64)
	PreCheckDelayInput() *float64
	PreferredChain() *string
	SetPreferredChain(val *string)
	PreferredChainInput() *string
	PrivateKeyPem() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RecursiveNameservers() *[]*string
	SetRecursiveNameservers(val *[]*string)
	RecursiveNameserversInput() *[]*string
	RevokeCertificateOnDestroy() interface{}
	SetRevokeCertificateOnDestroy(val interface{})
	RevokeCertificateOnDestroyInput() interface{}
	RevokeCertificateReason() *string
	SetRevokeCertificateReason(val *string)
	RevokeCertificateReasonInput() *string
	SubjectAlternativeNames() *[]*string
	SetSubjectAlternativeNames(val *[]*string)
	SubjectAlternativeNamesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsChallenge() CertificateTlsChallengeOutputReference
	TlsChallengeInput() *CertificateTlsChallenge
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDnsChallenge(value interface{})
	PutHttpChallenge(value *CertificateHttpChallenge)
	PutHttpMemcachedChallenge(value *CertificateHttpMemcachedChallenge)
	PutHttpS3Challenge(value *CertificateHttpS3Challenge)
	PutHttpWebrootChallenge(value *CertificateHttpWebrootChallenge)
	PutTlsChallenge(value *CertificateTlsChallenge)
	ResetCertificateP12Password()
	ResetCertificateRequestPem()
	ResetCertTimeout()
	ResetCommonName()
	ResetDisableCompletePropagation()
	ResetDnsChallenge()
	ResetHttpChallenge()
	ResetHttpMemcachedChallenge()
	ResetHttpS3Challenge()
	ResetHttpWebrootChallenge()
	ResetId()
	ResetKeyType()
	ResetMinDaysRemaining()
	ResetMustStaple()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreCheckDelay()
	ResetPreferredChain()
	ResetRecursiveNameservers()
	ResetRevokeCertificateOnDestroy()
	ResetRevokeCertificateReason()
	ResetSubjectAlternativeNames()
	ResetTlsChallenge()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Certificate
type jsiiProxy_Certificate struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Certificate) AccountKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) AccountKeyPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateNotAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNotAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateP12() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateP12",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateP12Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateP12Password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateP12PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateP12PasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificatePem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateRequestPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateRequestPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateRequestPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateRequestPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertificateUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CertTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) DisableCompletePropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCompletePropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) DisableCompletePropagationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCompletePropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) DnsChallenge() CertificateDnsChallengeList {
	var returns CertificateDnsChallengeList
	_jsii_.Get(
		j,
		"dnsChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) DnsChallengeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpChallenge() CertificateHttpChallengeOutputReference {
	var returns CertificateHttpChallengeOutputReference
	_jsii_.Get(
		j,
		"httpChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpChallengeInput() *CertificateHttpChallenge {
	var returns *CertificateHttpChallenge
	_jsii_.Get(
		j,
		"httpChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpMemcachedChallenge() CertificateHttpMemcachedChallengeOutputReference {
	var returns CertificateHttpMemcachedChallengeOutputReference
	_jsii_.Get(
		j,
		"httpMemcachedChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpMemcachedChallengeInput() *CertificateHttpMemcachedChallenge {
	var returns *CertificateHttpMemcachedChallenge
	_jsii_.Get(
		j,
		"httpMemcachedChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpS3Challenge() CertificateHttpS3ChallengeOutputReference {
	var returns CertificateHttpS3ChallengeOutputReference
	_jsii_.Get(
		j,
		"httpS3Challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpS3ChallengeInput() *CertificateHttpS3Challenge {
	var returns *CertificateHttpS3Challenge
	_jsii_.Get(
		j,
		"httpS3ChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpWebrootChallenge() CertificateHttpWebrootChallengeOutputReference {
	var returns CertificateHttpWebrootChallengeOutputReference
	_jsii_.Get(
		j,
		"httpWebrootChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) HttpWebrootChallengeInput() *CertificateHttpWebrootChallenge {
	var returns *CertificateHttpWebrootChallenge
	_jsii_.Get(
		j,
		"httpWebrootChallengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) IssuerPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) MinDaysRemaining() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDaysRemaining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) MinDaysRemainingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minDaysRemainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) MustStaple() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustStaple",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) MustStapleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mustStapleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PreCheckDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preCheckDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PreCheckDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preCheckDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PreferredChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PreferredChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) PrivateKeyPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RecursiveNameservers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recursiveNameservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RecursiveNameserversInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"recursiveNameserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RevokeCertificateOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeCertificateOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RevokeCertificateOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"revokeCertificateOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RevokeCertificateReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokeCertificateReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) RevokeCertificateReasonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokeCertificateReasonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) SubjectAlternativeNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) SubjectAlternativeNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) TlsChallenge() CertificateTlsChallengeOutputReference {
	var returns CertificateTlsChallengeOutputReference
	_jsii_.Get(
		j,
		"tlsChallenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Certificate) TlsChallengeInput() *CertificateTlsChallenge {
	var returns *CertificateTlsChallenge
	_jsii_.Get(
		j,
		"tlsChallengeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/vancluever/acme/2.19.0/docs/resources/certificate acme_certificate} Resource.
func NewCertificate(scope constructs.Construct, id *string, config *CertificateConfig) Certificate {
	_init_.Initialize()

	if err := validateNewCertificateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Certificate{}

	_jsii_.Create(
		"@cdktf/provider-acme.certificate.Certificate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/vancluever/acme/2.19.0/docs/resources/certificate acme_certificate} Resource.
func NewCertificate_Override(c Certificate, scope constructs.Construct, id *string, config *CertificateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-acme.certificate.Certificate",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Certificate)SetAccountKeyPem(val *string) {
	if err := j.validateSetAccountKeyPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKeyPem",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetCertificateP12Password(val *string) {
	if err := j.validateSetCertificateP12PasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateP12Password",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetCertificateRequestPem(val *string) {
	if err := j.validateSetCertificateRequestPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateRequestPem",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetCertTimeout(val *float64) {
	if err := j.validateSetCertTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certTimeout",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetCommonName(val *string) {
	if err := j.validateSetCommonNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetDisableCompletePropagation(val interface{}) {
	if err := j.validateSetDisableCompletePropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCompletePropagation",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetKeyType(val *string) {
	if err := j.validateSetKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetMinDaysRemaining(val *float64) {
	if err := j.validateSetMinDaysRemainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minDaysRemaining",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetMustStaple(val interface{}) {
	if err := j.validateSetMustStapleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mustStaple",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetPreCheckDelay(val *float64) {
	if err := j.validateSetPreCheckDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preCheckDelay",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetPreferredChain(val *string) {
	if err := j.validateSetPreferredChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredChain",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetRecursiveNameservers(val *[]*string) {
	if err := j.validateSetRecursiveNameserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recursiveNameservers",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetRevokeCertificateOnDestroy(val interface{}) {
	if err := j.validateSetRevokeCertificateOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revokeCertificateOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetRevokeCertificateReason(val *string) {
	if err := j.validateSetRevokeCertificateReasonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revokeCertificateReason",
		val,
	)
}

func (j *jsiiProxy_Certificate)SetSubjectAlternativeNames(val *[]*string) {
	if err := j.validateSetSubjectAlternativeNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subjectAlternativeNames",
		val,
	)
}

// Generates CDKTF code for importing a Certificate resource upon running "cdktf plan <stack-name>".
func Certificate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCertificate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-acme.certificate.Certificate",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Certificate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-acme.certificate.Certificate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Certificate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-acme.certificate.Certificate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Certificate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCertificate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-acme.certificate.Certificate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Certificate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-acme.certificate.Certificate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Certificate) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Certificate) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Certificate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Certificate) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Certificate) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Certificate) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Certificate) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Certificate) PutDnsChallenge(value interface{}) {
	if err := c.validatePutDnsChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDnsChallenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) PutHttpChallenge(value *CertificateHttpChallenge) {
	if err := c.validatePutHttpChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpChallenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) PutHttpMemcachedChallenge(value *CertificateHttpMemcachedChallenge) {
	if err := c.validatePutHttpMemcachedChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpMemcachedChallenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) PutHttpS3Challenge(value *CertificateHttpS3Challenge) {
	if err := c.validatePutHttpS3ChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpS3Challenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) PutHttpWebrootChallenge(value *CertificateHttpWebrootChallenge) {
	if err := c.validatePutHttpWebrootChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpWebrootChallenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) PutTlsChallenge(value *CertificateTlsChallenge) {
	if err := c.validatePutTlsChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTlsChallenge",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Certificate) ResetCertificateP12Password() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificateP12Password",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetCertificateRequestPem() {
	_jsii_.InvokeVoid(
		c,
		"resetCertificateRequestPem",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetCertTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetCertTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetCommonName() {
	_jsii_.InvokeVoid(
		c,
		"resetCommonName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetDisableCompletePropagation() {
	_jsii_.InvokeVoid(
		c,
		"resetDisableCompletePropagation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetDnsChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetHttpChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetHttpMemcachedChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpMemcachedChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetHttpS3Challenge() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpS3Challenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetHttpWebrootChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpWebrootChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetKeyType() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetMinDaysRemaining() {
	_jsii_.InvokeVoid(
		c,
		"resetMinDaysRemaining",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetMustStaple() {
	_jsii_.InvokeVoid(
		c,
		"resetMustStaple",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetPreCheckDelay() {
	_jsii_.InvokeVoid(
		c,
		"resetPreCheckDelay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetPreferredChain() {
	_jsii_.InvokeVoid(
		c,
		"resetPreferredChain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetRecursiveNameservers() {
	_jsii_.InvokeVoid(
		c,
		"resetRecursiveNameservers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetRevokeCertificateOnDestroy() {
	_jsii_.InvokeVoid(
		c,
		"resetRevokeCertificateOnDestroy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetRevokeCertificateReason() {
	_jsii_.InvokeVoid(
		c,
		"resetRevokeCertificateReason",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		c,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) ResetTlsChallenge() {
	_jsii_.InvokeVoid(
		c,
		"resetTlsChallenge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Certificate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Certificate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

