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

import "fmt"

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ArtifactRegistryRepositoryIAMAssetType string = "artifactregistry.googleapis.com/Repository"

func resourceConverterArtifactRegistryRepositoryIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         ArtifactRegistryRepositoryIAMAssetType,
		Convert:           GetArtifactRegistryRepositoryIamPolicyCaiObject,
		MergeCreateUpdate: MergeArtifactRegistryRepositoryIamPolicy,
	}
}

func resourceConverterArtifactRegistryRepositoryIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         ArtifactRegistryRepositoryIAMAssetType,
		Convert:           GetArtifactRegistryRepositoryIamBindingCaiObject,
		FetchFullResource: FetchArtifactRegistryRepositoryIamPolicy,
		MergeCreateUpdate: MergeArtifactRegistryRepositoryIamBinding,
		MergeDelete:       MergeArtifactRegistryRepositoryIamBindingDelete,
	}
}

func resourceConverterArtifactRegistryRepositoryIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         ArtifactRegistryRepositoryIAMAssetType,
		Convert:           GetArtifactRegistryRepositoryIamMemberCaiObject,
		FetchFullResource: FetchArtifactRegistryRepositoryIamPolicy,
		MergeCreateUpdate: MergeArtifactRegistryRepositoryIamMember,
		MergeDelete:       MergeArtifactRegistryRepositoryIamMemberDelete,
	}
}

func GetArtifactRegistryRepositoryIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newArtifactRegistryRepositoryIamAsset(d, config, expandIamPolicyBindings)
}

func GetArtifactRegistryRepositoryIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newArtifactRegistryRepositoryIamAsset(d, config, expandIamRoleBindings)
}

func GetArtifactRegistryRepositoryIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newArtifactRegistryRepositoryIamAsset(d, config, expandIamMemberBindings)
}

func MergeArtifactRegistryRepositoryIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeArtifactRegistryRepositoryIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeArtifactRegistryRepositoryIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeArtifactRegistryRepositoryIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeArtifactRegistryRepositoryIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newArtifactRegistryRepositoryIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//artifactregistry.googleapis.com/projects/{{project}}/locations/{{location}}/repositories/{{repository}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: ArtifactRegistryRepositoryIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchArtifactRegistryRepositoryIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("repository"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		ArtifactRegistryRepositoryIamUpdaterProducer,
		d,
		config,
		"//artifactregistry.googleapis.com/projects/{{project}}/locations/{{location}}/repositories/{{repository}}",
		ArtifactRegistryRepositoryIAMAssetType,
	)
}
