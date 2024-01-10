// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package certificate

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificateDnsChallengeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CertificateDnsChallengeList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificateDnsChallengeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificateDnsChallengeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CertificateDnsChallengeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificateDnsChallengeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificateDnsChallengeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificateDnsChallengeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

