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

const KMSCryptoKeyVersionAssetType string = "cloudkms.googleapis.com/CryptoKeyVersion"

func resourceConverterKMSCryptoKeyVersion() ResourceConverter {
	return ResourceConverter{
		AssetType: KMSCryptoKeyVersionAssetType,
		Convert:   GetKMSCryptoKeyVersionCaiObject,
	}
}

func GetKMSCryptoKeyVersionCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//cloudkms.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetKMSCryptoKeyVersionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: KMSCryptoKeyVersionAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "CryptoKeyVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetKMSCryptoKeyVersionApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	stateProp, err := expandKMSCryptoKeyVersionState(d.Get("state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("state"); !isEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}

	return obj, nil
}

func expandKMSCryptoKeyVersionState(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
