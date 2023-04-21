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
	"reflect"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DataCatalogPolicyTagAssetType string = "datacatalog.googleapis.com/PolicyTag"

func resourceConverterDataCatalogPolicyTag() ResourceConverter {
	return ResourceConverter{
		AssetType: DataCatalogPolicyTagAssetType,
		Convert:   GetDataCatalogPolicyTagCaiObject,
	}
}

func GetDataCatalogPolicyTagCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//datacatalog.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDataCatalogPolicyTagApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DataCatalogPolicyTagAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datacatalog/v1/rest",
				DiscoveryName:        "PolicyTag",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDataCatalogPolicyTagApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogPolicyTagDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogPolicyTagDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentPolicyTagProp, err := expandDataCatalogPolicyTagParentPolicyTag(d.Get("parent_policy_tag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_policy_tag"); !isEmptyValue(reflect.ValueOf(parentPolicyTagProp)) && (ok || !reflect.DeepEqual(v, parentPolicyTagProp)) {
		obj["parentPolicyTag"] = parentPolicyTagProp
	}

	return obj, nil
}

func expandDataCatalogPolicyTagDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagParentPolicyTag(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
