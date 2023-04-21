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

const GKEBackupBackupPlanAssetType string = "gkebackup.googleapis.com/BackupPlan"

func resourceConverterGKEBackupBackupPlan() ResourceConverter {
	return ResourceConverter{
		AssetType: GKEBackupBackupPlanAssetType,
		Convert:   GetGKEBackupBackupPlanCaiObject,
	}
}

func GetGKEBackupBackupPlanCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetGKEBackupBackupPlanApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: GKEBackupBackupPlanAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkebackup/v1/rest",
				DiscoveryName:        "BackupPlan",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetGKEBackupBackupPlanApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandGKEBackupBackupPlanName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandGKEBackupBackupPlanDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	clusterProp, err := expandGKEBackupBackupPlanCluster(d.Get("cluster"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cluster"); !isEmptyValue(reflect.ValueOf(clusterProp)) && (ok || !reflect.DeepEqual(v, clusterProp)) {
		obj["cluster"] = clusterProp
	}
	retentionPolicyProp, err := expandGKEBackupBackupPlanRetentionPolicy(d.Get("retention_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention_policy"); !isEmptyValue(reflect.ValueOf(retentionPolicyProp)) && (ok || !reflect.DeepEqual(v, retentionPolicyProp)) {
		obj["retentionPolicy"] = retentionPolicyProp
	}
	labelsProp, err := expandGKEBackupBackupPlanLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	backupScheduleProp, err := expandGKEBackupBackupPlanBackupSchedule(d.Get("backup_schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_schedule"); !isEmptyValue(reflect.ValueOf(backupScheduleProp)) && (ok || !reflect.DeepEqual(v, backupScheduleProp)) {
		obj["backupSchedule"] = backupScheduleProp
	}
	deactivatedProp, err := expandGKEBackupBackupPlanDeactivated(d.Get("deactivated"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deactivated"); !isEmptyValue(reflect.ValueOf(deactivatedProp)) && (ok || !reflect.DeepEqual(v, deactivatedProp)) {
		obj["deactivated"] = deactivatedProp
	}
	backupConfigProp, err := expandGKEBackupBackupPlanBackupConfig(d.Get("backup_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_config"); !isEmptyValue(reflect.ValueOf(backupConfigProp)) && (ok || !reflect.DeepEqual(v, backupConfigProp)) {
		obj["backupConfig"] = backupConfigProp
	}

	return obj, nil
}

func expandGKEBackupBackupPlanName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
}

func expandGKEBackupBackupPlanDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanCluster(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicy(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackupDeleteLockDays, err := expandGKEBackupBackupPlanRetentionPolicyBackupDeleteLockDays(original["backup_delete_lock_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupDeleteLockDays); val.IsValid() && !isEmptyValue(val) {
		transformed["backupDeleteLockDays"] = transformedBackupDeleteLockDays
	}

	transformedBackupRetainDays, err := expandGKEBackupBackupPlanRetentionPolicyBackupRetainDays(original["backup_retain_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupRetainDays); val.IsValid() && !isEmptyValue(val) {
		transformed["backupRetainDays"] = transformedBackupRetainDays
	}

	transformedLocked, err := expandGKEBackupBackupPlanRetentionPolicyLocked(original["locked"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocked); val.IsValid() && !isEmptyValue(val) {
		transformed["locked"] = transformedLocked
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanRetentionPolicyBackupDeleteLockDays(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicyBackupRetainDays(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicyLocked(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanLabels(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGKEBackupBackupPlanBackupSchedule(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCronSchedule, err := expandGKEBackupBackupPlanBackupScheduleCronSchedule(original["cron_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCronSchedule); val.IsValid() && !isEmptyValue(val) {
		transformed["cronSchedule"] = transformedCronSchedule
	}

	transformedPaused, err := expandGKEBackupBackupPlanBackupSchedulePaused(original["paused"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPaused); val.IsValid() && !isEmptyValue(val) {
		transformed["paused"] = transformedPaused
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupScheduleCronSchedule(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupSchedulePaused(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanDeactivated(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIncludeVolumeData, err := expandGKEBackupBackupPlanBackupConfigIncludeVolumeData(original["include_volume_data"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeVolumeData); val.IsValid() && !isEmptyValue(val) {
		transformed["includeVolumeData"] = transformedIncludeVolumeData
	}

	transformedIncludeSecrets, err := expandGKEBackupBackupPlanBackupConfigIncludeSecrets(original["include_secrets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeSecrets); val.IsValid() && !isEmptyValue(val) {
		transformed["includeSecrets"] = transformedIncludeSecrets
	}

	transformedEncryptionKey, err := expandGKEBackupBackupPlanBackupConfigEncryptionKey(original["encryption_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncryptionKey); val.IsValid() && !isEmptyValue(val) {
		transformed["encryptionKey"] = transformedEncryptionKey
	}

	transformedAllNamespaces, err := expandGKEBackupBackupPlanBackupConfigAllNamespaces(original["all_namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllNamespaces); val.IsValid() && !isEmptyValue(val) {
		transformed["allNamespaces"] = transformedAllNamespaces
	}

	transformedSelectedNamespaces, err := expandGKEBackupBackupPlanBackupConfigSelectedNamespaces(original["selected_namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelectedNamespaces); val.IsValid() && !isEmptyValue(val) {
		transformed["selectedNamespaces"] = transformedSelectedNamespaces
	}

	transformedSelectedApplications, err := expandGKEBackupBackupPlanBackupConfigSelectedApplications(original["selected_applications"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelectedApplications); val.IsValid() && !isEmptyValue(val) {
		transformed["selectedApplications"] = transformedSelectedApplications
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigIncludeVolumeData(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigIncludeSecrets(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigEncryptionKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcpKmsEncryptionKey, err := expandGKEBackupBackupPlanBackupConfigEncryptionKeyGcpKmsEncryptionKey(original["gcp_kms_encryption_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpKmsEncryptionKey); val.IsValid() && !isEmptyValue(val) {
		transformed["gcpKmsEncryptionKey"] = transformedGcpKmsEncryptionKey
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigEncryptionKeyGcpKmsEncryptionKey(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigAllNamespaces(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedNamespaces(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespaces, err := expandGKEBackupBackupPlanBackupConfigSelectedNamespacesNamespaces(original["namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespaces); val.IsValid() && !isEmptyValue(val) {
		transformed["namespaces"] = transformedNamespaces
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedNamespacesNamespaces(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplications(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespacedNames, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames(original["namespaced_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespacedNames); val.IsValid() && !isEmptyValue(val) {
		transformed["namespacedNames"] = transformedNamespacedNames
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamespace, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesNamespace(original["namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !isEmptyValue(val) {
			transformed["namespace"] = transformedNamespace
		}

		transformedName, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesNamespace(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
