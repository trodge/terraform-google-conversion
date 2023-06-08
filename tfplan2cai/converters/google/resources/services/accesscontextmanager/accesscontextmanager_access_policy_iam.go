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

package accesscontextmanager

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const AccessContextManagerAccessPolicyIAMAssetType string = "accesscontextmanager.googleapis.com/AccessPolicy"

func ResourceConverterAccessContextManagerAccessPolicyIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamPolicyCaiObject,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamPolicy,
	}
}

func ResourceConverterAccessContextManagerAccessPolicyIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamBindingCaiObject,
		FetchFullResource: FetchAccessContextManagerAccessPolicyIamPolicy,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamBinding,
		MergeDelete:       MergeAccessContextManagerAccessPolicyIamBindingDelete,
	}
}

func ResourceConverterAccessContextManagerAccessPolicyIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamMemberCaiObject,
		FetchFullResource: FetchAccessContextManagerAccessPolicyIamPolicy,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamMember,
		MergeDelete:       MergeAccessContextManagerAccessPolicyIamMemberDelete,
	}
}

func GetAccessContextManagerAccessPolicyIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetAccessContextManagerAccessPolicyIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetAccessContextManagerAccessPolicyIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeAccessContextManagerAccessPolicyIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeAccessContextManagerAccessPolicyIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeAccessContextManagerAccessPolicyIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeAccessContextManagerAccessPolicyIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeAccessContextManagerAccessPolicyIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newAccessContextManagerAccessPolicyIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//accesscontextmanager.googleapis.com/accessPolicies/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: AccessContextManagerAccessPolicyIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchAccessContextManagerAccessPolicyIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		AccessContextManagerAccessPolicyIamUpdaterProducer,
		d,
		config,
		"//accesscontextmanager.googleapis.com/accessPolicies/{{name}}",
		AccessContextManagerAccessPolicyIAMAssetType,
	)
}