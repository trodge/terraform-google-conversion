// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

import (
	"fmt"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const PrivatecaCertificateTemplateIAMAssetType string = "privateca.googleapis.com/CertificateTemplate"

func resourceConverterPrivatecaCertificateTemplateIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCertificateTemplateIAMAssetType,
		Convert:           GetPrivatecaCertificateTemplateIamPolicyCaiObject,
		MergeCreateUpdate: MergePrivatecaCertificateTemplateIamPolicy,
	}
}

func resourceConverterPrivatecaCertificateTemplateIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCertificateTemplateIAMAssetType,
		Convert:           GetPrivatecaCertificateTemplateIamBindingCaiObject,
		FetchFullResource: FetchPrivatecaCertificateTemplateIamPolicy,
		MergeCreateUpdate: MergePrivatecaCertificateTemplateIamBinding,
		MergeDelete:       MergePrivatecaCertificateTemplateIamBindingDelete,
	}
}

func resourceConverterPrivatecaCertificateTemplateIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         PrivatecaCertificateTemplateIAMAssetType,
		Convert:           GetPrivatecaCertificateTemplateIamMemberCaiObject,
		FetchFullResource: FetchPrivatecaCertificateTemplateIamPolicy,
		MergeCreateUpdate: MergePrivatecaCertificateTemplateIamMember,
		MergeDelete:       MergePrivatecaCertificateTemplateIamMemberDelete,
	}
}

func GetPrivatecaCertificateTemplateIamPolicyCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPrivatecaCertificateTemplateIamAsset(d, config, expandIamPolicyBindings)
}

func GetPrivatecaCertificateTemplateIamBindingCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPrivatecaCertificateTemplateIamAsset(d, config, expandIamRoleBindings)
}

func GetPrivatecaCertificateTemplateIamMemberCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newPrivatecaCertificateTemplateIamAsset(d, config, expandIamMemberBindings)
}

func MergePrivatecaCertificateTemplateIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergePrivatecaCertificateTemplateIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergePrivatecaCertificateTemplateIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergePrivatecaCertificateTemplateIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergePrivatecaCertificateTemplateIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newPrivatecaCertificateTemplateIamAsset(
	d TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}/{{certificate_template}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: PrivatecaCertificateTemplateIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchPrivatecaCertificateTemplateIamPolicy(d TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("certificate_template"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		PrivatecaCertificateTemplateIamUpdaterProducer,
		d,
		config,
		"//privateca.googleapis.com/projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}/{{certificate_template}}",
		PrivatecaCertificateTemplateIAMAssetType,
	)
}
