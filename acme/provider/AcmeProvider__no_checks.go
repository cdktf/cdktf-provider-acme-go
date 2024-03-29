// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AcmeProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_AcmeProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateAcmeProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateAcmeProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAcmeProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateAcmeProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func validateNewAcmeProviderParameters(scope constructs.Construct, id *string, config *AcmeProviderConfig) error {
	return nil
}

