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
const VertexAIFeaturestoreEntitytypeIAMAssetType string = "{{region}}-aiplatform.googleapis.com/FeaturestoreEntitytype"

func resourceConverterVertexAIFeaturestoreEntitytypeIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamPolicyCaiObject,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamPolicy,
	}
}

func resourceConverterVertexAIFeaturestoreEntitytypeIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamBindingCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreEntitytypeIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamBinding,
		MergeDelete:       MergeVertexAIFeaturestoreEntitytypeIamBindingDelete,
	}
}

func resourceConverterVertexAIFeaturestoreEntitytypeIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         VertexAIFeaturestoreEntitytypeIAMAssetType,
		Convert:           GetVertexAIFeaturestoreEntitytypeIamMemberCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreEntitytypeIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreEntitytypeIamMember,
		MergeDelete:       MergeVertexAIFeaturestoreEntitytypeIamMemberDelete,
	}
}

func GetVertexAIFeaturestoreEntitytypeIamPolicyCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, expandIamPolicyBindings)
}

func GetVertexAIFeaturestoreEntitytypeIamBindingCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, expandIamRoleBindings)
}

func GetVertexAIFeaturestoreEntitytypeIamMemberCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newVertexAIFeaturestoreEntitytypeIamAsset(d, config, expandIamMemberBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeVertexAIFeaturestoreEntitytypeIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeVertexAIFeaturestoreEntitytypeIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newVertexAIFeaturestoreEntitytypeIamAsset(
	d TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//{{region}}-aiplatform.googleapis.com/{{featurestore}}/entityTypes/{{entitytype}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: VertexAIFeaturestoreEntitytypeIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchVertexAIFeaturestoreEntitytypeIamPolicy(d TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("featurestore"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("entitytype"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		VertexAIFeaturestoreEntitytypeIamUpdaterProducer,
		d,
		config,
		"//{{region}}-aiplatform.googleapis.com/{{featurestore}}/entityTypes/{{entitytype}}",
		VertexAIFeaturestoreEntitytypeIAMAssetType,
	)
}
