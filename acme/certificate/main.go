// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package certificate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.Certificate",
		reflect.TypeOf((*Certificate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountKeyPem", GoGetter: "AccountKeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "accountKeyPemInput", GoGetter: "AccountKeyPemInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certificateDomain", GoGetter: "CertificateDomain"},
			_jsii_.MemberProperty{JsiiProperty: "certificateNotAfter", GoGetter: "CertificateNotAfter"},
			_jsii_.MemberProperty{JsiiProperty: "certificateP12", GoGetter: "CertificateP12"},
			_jsii_.MemberProperty{JsiiProperty: "certificateP12Password", GoGetter: "CertificateP12Password"},
			_jsii_.MemberProperty{JsiiProperty: "certificateP12PasswordInput", GoGetter: "CertificateP12PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificatePem", GoGetter: "CertificatePem"},
			_jsii_.MemberProperty{JsiiProperty: "certificateRequestPem", GoGetter: "CertificateRequestPem"},
			_jsii_.MemberProperty{JsiiProperty: "certificateRequestPemInput", GoGetter: "CertificateRequestPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateUrl", GoGetter: "CertificateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "certTimeout", GoGetter: "CertTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "certTimeoutInput", GoGetter: "CertTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "disableCompletePropagation", GoGetter: "DisableCompletePropagation"},
			_jsii_.MemberProperty{JsiiProperty: "disableCompletePropagationInput", GoGetter: "DisableCompletePropagationInput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsChallenge", GoGetter: "DnsChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "dnsChallengeInput", GoGetter: "DnsChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpChallenge", GoGetter: "HttpChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "httpChallengeInput", GoGetter: "HttpChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpMemcachedChallenge", GoGetter: "HttpMemcachedChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "httpMemcachedChallengeInput", GoGetter: "HttpMemcachedChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpS3Challenge", GoGetter: "HttpS3Challenge"},
			_jsii_.MemberProperty{JsiiProperty: "httpS3ChallengeInput", GoGetter: "HttpS3ChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpWebrootChallenge", GoGetter: "HttpWebrootChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "httpWebrootChallengeInput", GoGetter: "HttpWebrootChallengeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerPem", GoGetter: "IssuerPem"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minDaysRemaining", GoGetter: "MinDaysRemaining"},
			_jsii_.MemberProperty{JsiiProperty: "minDaysRemainingInput", GoGetter: "MinDaysRemainingInput"},
			_jsii_.MemberProperty{JsiiProperty: "mustStaple", GoGetter: "MustStaple"},
			_jsii_.MemberProperty{JsiiProperty: "mustStapleInput", GoGetter: "MustStapleInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "preCheckDelay", GoGetter: "PreCheckDelay"},
			_jsii_.MemberProperty{JsiiProperty: "preCheckDelayInput", GoGetter: "PreCheckDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChain", GoGetter: "PreferredChain"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChainInput", GoGetter: "PreferredChainInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPem", GoGetter: "PrivateKeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDnsChallenge", GoMethod: "PutDnsChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpChallenge", GoMethod: "PutHttpChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpMemcachedChallenge", GoMethod: "PutHttpMemcachedChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpS3Challenge", GoMethod: "PutHttpS3Challenge"},
			_jsii_.MemberMethod{JsiiMethod: "putHttpWebrootChallenge", GoMethod: "PutHttpWebrootChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "putTlsChallenge", GoMethod: "PutTlsChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recursiveNameservers", GoGetter: "RecursiveNameservers"},
			_jsii_.MemberProperty{JsiiProperty: "recursiveNameserversInput", GoGetter: "RecursiveNameserversInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateP12Password", GoMethod: "ResetCertificateP12Password"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateRequestPem", GoMethod: "ResetCertificateRequestPem"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertTimeout", GoMethod: "ResetCertTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonName", GoMethod: "ResetCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableCompletePropagation", GoMethod: "ResetDisableCompletePropagation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsChallenge", GoMethod: "ResetDnsChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpChallenge", GoMethod: "ResetHttpChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpMemcachedChallenge", GoMethod: "ResetHttpMemcachedChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpS3Challenge", GoMethod: "ResetHttpS3Challenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpWebrootChallenge", GoMethod: "ResetHttpWebrootChallenge"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyType", GoMethod: "ResetKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinDaysRemaining", GoMethod: "ResetMinDaysRemaining"},
			_jsii_.MemberMethod{JsiiMethod: "resetMustStaple", GoMethod: "ResetMustStaple"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreCheckDelay", GoMethod: "ResetPreCheckDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredChain", GoMethod: "ResetPreferredChain"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecursiveNameservers", GoMethod: "ResetRecursiveNameservers"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevokeCertificateOnDestroy", GoMethod: "ResetRevokeCertificateOnDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsChallenge", GoMethod: "ResetTlsChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "revokeCertificateOnDestroy", GoGetter: "RevokeCertificateOnDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "revokeCertificateOnDestroyInput", GoGetter: "RevokeCertificateOnDestroyInput"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsChallenge", GoGetter: "TlsChallenge"},
			_jsii_.MemberProperty{JsiiProperty: "tlsChallengeInput", GoGetter: "TlsChallengeInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Certificate{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateConfig",
		reflect.TypeOf((*CertificateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateDnsChallenge",
		reflect.TypeOf((*CertificateDnsChallenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateDnsChallengeList",
		reflect.TypeOf((*CertificateDnsChallengeList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateDnsChallengeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateDnsChallengeOutputReference",
		reflect.TypeOf((*CertificateDnsChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerInput", GoGetter: "ProviderInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfig", GoMethod: "ResetConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateDnsChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateHttpChallenge",
		reflect.TypeOf((*CertificateHttpChallenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateHttpChallengeOutputReference",
		reflect.TypeOf((*CertificateHttpChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "proxyHeader", GoGetter: "ProxyHeader"},
			_jsii_.MemberProperty{JsiiProperty: "proxyHeaderInput", GoGetter: "ProxyHeaderInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetProxyHeader", GoMethod: "ResetProxyHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateHttpChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateHttpMemcachedChallenge",
		reflect.TypeOf((*CertificateHttpMemcachedChallenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateHttpMemcachedChallengeOutputReference",
		reflect.TypeOf((*CertificateHttpMemcachedChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hosts", GoGetter: "Hosts"},
			_jsii_.MemberProperty{JsiiProperty: "hostsInput", GoGetter: "HostsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateHttpMemcachedChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateHttpS3Challenge",
		reflect.TypeOf((*CertificateHttpS3Challenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateHttpS3ChallengeOutputReference",
		reflect.TypeOf((*CertificateHttpS3ChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "s3Bucket", GoGetter: "S3Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "s3BucketInput", GoGetter: "S3BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateHttpS3ChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateHttpWebrootChallenge",
		reflect.TypeOf((*CertificateHttpWebrootChallenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateHttpWebrootChallengeOutputReference",
		reflect.TypeOf((*CertificateHttpWebrootChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "directory", GoGetter: "Directory"},
			_jsii_.MemberProperty{JsiiProperty: "directoryInput", GoGetter: "DirectoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateHttpWebrootChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-acme.certificate.CertificateTlsChallenge",
		reflect.TypeOf((*CertificateTlsChallenge)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-acme.certificate.CertificateTlsChallengeOutputReference",
		reflect.TypeOf((*CertificateTlsChallengeOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPort", GoMethod: "ResetPort"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CertificateTlsChallengeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
