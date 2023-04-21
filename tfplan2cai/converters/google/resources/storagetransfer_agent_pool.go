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

const StorageTransferAgentPoolAssetType string = "storagetransfer.googleapis.com/AgentPool"

func resourceConverterStorageTransferAgentPool() ResourceConverter {
	return ResourceConverter{
		AssetType: StorageTransferAgentPoolAssetType,
		Convert:   GetStorageTransferAgentPoolCaiObject,
	}
}

func GetStorageTransferAgentPoolCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//storagetransfer.googleapis.com/projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetStorageTransferAgentPoolApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: StorageTransferAgentPoolAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storagetransfer/v1/rest",
				DiscoveryName:        "AgentPool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetStorageTransferAgentPoolApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandStorageTransferAgentPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bandwidthLimitProp, err := expandStorageTransferAgentPoolBandwidthLimit(d.Get("bandwidth_limit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bandwidth_limit"); !isEmptyValue(reflect.ValueOf(bandwidthLimitProp)) && (ok || !reflect.DeepEqual(v, bandwidthLimitProp)) {
		obj["bandwidthLimit"] = bandwidthLimitProp
	}

	return obj, nil
}

func expandStorageTransferAgentPoolDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageTransferAgentPoolBandwidthLimit(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLimitMbps, err := expandStorageTransferAgentPoolBandwidthLimitLimitMbps(original["limit_mbps"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimitMbps); val.IsValid() && !isEmptyValue(val) {
		transformed["limitMbps"] = transformedLimitMbps
	}

	return transformed, nil
}

func expandStorageTransferAgentPoolBandwidthLimitLimitMbps(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
