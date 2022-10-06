//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package registration

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Registration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Registration) validatePutExternalAccountBindingParameters(value *RegistrationExternalAccountBinding) error {
	return nil
}

func validateRegistration_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetAccountKeyPemParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetEmailAddressParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Registration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewRegistrationParameters(scope constructs.Construct, id *string, config *RegistrationConfig) error {
	return nil
}
