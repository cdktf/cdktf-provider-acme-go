// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certificate

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Certificate) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutDnsChallengeParameters(value interface{}) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutHttpChallengeParameters(value *CertificateHttpChallenge) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutHttpMemcachedChallengeParameters(value *CertificateHttpMemcachedChallenge) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutHttpS3ChallengeParameters(value *CertificateHttpS3Challenge) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutHttpWebrootChallengeParameters(value *CertificateHttpWebrootChallenge) error {
	return nil
}

func (c *jsiiProxy_Certificate) validatePutTlsChallengeParameters(value *CertificateTlsChallenge) error {
	return nil
}

func validateCertificate_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCertificate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCertificate_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCertificate_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetAccountKeyPemParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCertificateP12PasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCertificateRequestPemParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCertTimeoutParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCommonNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetDisableCompletePropagationParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetKeyTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetMinDaysRemainingParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetMustStapleParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetPreCheckDelayParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetPreferredChainParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetProfileParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetRecursiveNameserversParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetRevokeCertificateOnDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetRevokeCertificateReasonParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Certificate) validateSetSubjectAlternativeNamesParameters(val *[]*string) error {
	return nil
}

func validateNewCertificateParameters(scope constructs.Construct, id *string, config *CertificateConfig) error {
	return nil
}

