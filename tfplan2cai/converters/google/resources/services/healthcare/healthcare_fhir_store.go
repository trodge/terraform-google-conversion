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

package healthcare

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const HealthcareFhirStoreAssetType string = "healthcare.googleapis.com/FhirStore"

func ResourceConverterHealthcareFhirStore() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: HealthcareFhirStoreAssetType,
		Convert:   GetHealthcareFhirStoreCaiObject,
	}
}

func GetHealthcareFhirStoreCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/fhirStores/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetHealthcareFhirStoreApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: HealthcareFhirStoreAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/healthcare/v1/rest",
				DiscoveryName:        "FhirStore",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetHealthcareFhirStoreApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandHealthcareFhirStoreName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	versionProp, err := expandHealthcareFhirStoreVersion(d.Get("version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("version"); !tpgresource.IsEmptyValue(reflect.ValueOf(versionProp)) && (ok || !reflect.DeepEqual(v, versionProp)) {
		obj["version"] = versionProp
	}
	enableUpdateCreateProp, err := expandHealthcareFhirStoreEnableUpdateCreate(d.Get("enable_update_create"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_update_create"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableUpdateCreateProp)) && (ok || !reflect.DeepEqual(v, enableUpdateCreateProp)) {
		obj["enableUpdateCreate"] = enableUpdateCreateProp
	}
	disableReferentialIntegrityProp, err := expandHealthcareFhirStoreDisableReferentialIntegrity(d.Get("disable_referential_integrity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disable_referential_integrity"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableReferentialIntegrityProp)) && (ok || !reflect.DeepEqual(v, disableReferentialIntegrityProp)) {
		obj["disableReferentialIntegrity"] = disableReferentialIntegrityProp
	}
	disableResourceVersioningProp, err := expandHealthcareFhirStoreDisableResourceVersioning(d.Get("disable_resource_versioning"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disable_resource_versioning"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableResourceVersioningProp)) && (ok || !reflect.DeepEqual(v, disableResourceVersioningProp)) {
		obj["disableResourceVersioning"] = disableResourceVersioningProp
	}
	enableHistoryImportProp, err := expandHealthcareFhirStoreEnableHistoryImport(d.Get("enable_history_import"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_history_import"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableHistoryImportProp)) && (ok || !reflect.DeepEqual(v, enableHistoryImportProp)) {
		obj["enableHistoryImport"] = enableHistoryImportProp
	}
	labelsProp, err := expandHealthcareFhirStoreLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	notificationConfigProp, err := expandHealthcareFhirStoreNotificationConfig(d.Get("notification_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationConfigProp)) && (ok || !reflect.DeepEqual(v, notificationConfigProp)) {
		obj["notificationConfig"] = notificationConfigProp
	}
	streamConfigsProp, err := expandHealthcareFhirStoreStreamConfigs(d.Get("stream_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("stream_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(streamConfigsProp)) && (ok || !reflect.DeepEqual(v, streamConfigsProp)) {
		obj["streamConfigs"] = streamConfigsProp
	}

	return obj, nil
}

func expandHealthcareFhirStoreName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableUpdateCreate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableReferentialIntegrity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreDisableResourceVersioning(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreEnableHistoryImport(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandHealthcareFhirStoreNotificationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPubsubTopic, err := expandHealthcareFhirStoreNotificationConfigPubsubTopic(original["pubsub_topic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPubsubTopic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["pubsubTopic"] = transformedPubsubTopic
	}

	return transformed, nil
}

func expandHealthcareFhirStoreNotificationConfigPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreStreamConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedResourceTypes, err := expandHealthcareFhirStoreStreamConfigsResourceTypes(original["resource_types"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedResourceTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["resourceTypes"] = transformedResourceTypes
		}

		transformedBigqueryDestination, err := expandHealthcareFhirStoreStreamConfigsBigqueryDestination(original["bigquery_destination"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedBigqueryDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["bigqueryDestination"] = transformedBigqueryDestination
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandHealthcareFhirStoreStreamConfigsResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreStreamConfigsBigqueryDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDatasetUri, err := expandHealthcareFhirStoreStreamConfigsBigqueryDestinationDatasetUri(original["dataset_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDatasetUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["datasetUri"] = transformedDatasetUri
	}

	transformedSchemaConfig, err := expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(original["schema_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schemaConfig"] = transformedSchemaConfig
	}

	return transformed, nil
}

func expandHealthcareFhirStoreStreamConfigsBigqueryDestinationDatasetUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchemaType, err := expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaType(original["schema_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchemaType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schemaType"] = transformedSchemaType
	}

	transformedRecursiveStructureDepth, err := expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigRecursiveStructureDepth(original["recursive_structure_depth"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecursiveStructureDepth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recursiveStructureDepth"] = transformedRecursiveStructureDepth
	}

	return transformed, nil
}

func expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigSchemaType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigRecursiveStructureDepth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}