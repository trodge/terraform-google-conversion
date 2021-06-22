// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "fmt"

func GetIapAppEngineVersionIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapAppEngineVersionIamAsset(d, config, expandIamPolicyBindings)
}

func GetIapAppEngineVersionIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapAppEngineVersionIamAsset(d, config, expandIamRoleBindings)
}

func GetIapAppEngineVersionIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newIapAppEngineVersionIamAsset(d, config, expandIamMemberBindings)
}

func MergeIapAppEngineVersionIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapAppEngineVersionIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeIapAppEngineVersionIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeIapAppEngineVersionIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeIapAppEngineVersionIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newIapAppEngineVersionIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//iap.googleapis.com/{{appengineversion}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "iap.googleapis.com/AppEngineVersion",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapAppEngineVersionIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{appengineversion}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		IapAppEngineVersionIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/{{appengineversion}}",
		"iap.googleapis.com/AppEngineVersion",
	)
}