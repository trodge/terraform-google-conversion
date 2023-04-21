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
const DataFusionInstanceIAMAssetType string = "datafusion.googleapis.com/Instance"

func resourceConverterDataFusionInstanceIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataFusionInstanceIAMAssetType,
		Convert:           GetDataFusionInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataFusionInstanceIamPolicy,
	}
}

func resourceConverterDataFusionInstanceIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataFusionInstanceIAMAssetType,
		Convert:           GetDataFusionInstanceIamBindingCaiObject,
		FetchFullResource: FetchDataFusionInstanceIamPolicy,
		MergeCreateUpdate: MergeDataFusionInstanceIamBinding,
		MergeDelete:       MergeDataFusionInstanceIamBindingDelete,
	}
}

func resourceConverterDataFusionInstanceIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         DataFusionInstanceIAMAssetType,
		Convert:           GetDataFusionInstanceIamMemberCaiObject,
		FetchFullResource: FetchDataFusionInstanceIamPolicy,
		MergeCreateUpdate: MergeDataFusionInstanceIamMember,
		MergeDelete:       MergeDataFusionInstanceIamMemberDelete,
	}
}

func GetDataFusionInstanceIamPolicyCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataFusionInstanceIamAsset(d, config, expandIamPolicyBindings)
}

func GetDataFusionInstanceIamBindingCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataFusionInstanceIamAsset(d, config, expandIamRoleBindings)
}

func GetDataFusionInstanceIamMemberCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newDataFusionInstanceIamAsset(d, config, expandIamMemberBindings)
}

func MergeDataFusionInstanceIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataFusionInstanceIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeDataFusionInstanceIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeDataFusionInstanceIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeDataFusionInstanceIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newDataFusionInstanceIamAsset(
	d TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//datafusion.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: DataFusionInstanceIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataFusionInstanceIamPolicy(d TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		DataFusionInstanceIamUpdaterProducer,
		d,
		config,
		"//datafusion.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}",
		DataFusionInstanceIAMAssetType,
	)
}
