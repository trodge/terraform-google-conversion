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

package dialogflowcx

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DialogflowCXAgentAssetType string = "{{location}}-dialogflow.googleapis.com/Agent"

func ResourceConverterDialogflowCXAgent() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: DialogflowCXAgentAssetType,
		Convert:   GetDialogflowCXAgentCaiObject,
	}
}

func GetDialogflowCXAgentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//{{location}}-dialogflow.googleapis.com/projects/{{project}}/locations/{{location}}/agents/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetDialogflowCXAgentApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: DialogflowCXAgentAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-dialogflow/v3/rest",
				DiscoveryName:        "Agent",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetDialogflowCXAgentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXAgentDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	defaultLanguageCodeProp, err := expandDialogflowCXAgentDefaultLanguageCode(d.Get("default_language_code"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_language_code"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultLanguageCodeProp)) && (ok || !reflect.DeepEqual(v, defaultLanguageCodeProp)) {
		obj["defaultLanguageCode"] = defaultLanguageCodeProp
	}
	supportedLanguageCodesProp, err := expandDialogflowCXAgentSupportedLanguageCodes(d.Get("supported_language_codes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("supported_language_codes"); !tpgresource.IsEmptyValue(reflect.ValueOf(supportedLanguageCodesProp)) && (ok || !reflect.DeepEqual(v, supportedLanguageCodesProp)) {
		obj["supportedLanguageCodes"] = supportedLanguageCodesProp
	}
	timeZoneProp, err := expandDialogflowCXAgentTimeZone(d.Get("time_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("time_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(timeZoneProp)) && (ok || !reflect.DeepEqual(v, timeZoneProp)) {
		obj["timeZone"] = timeZoneProp
	}
	descriptionProp, err := expandDialogflowCXAgentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	avatarUriProp, err := expandDialogflowCXAgentAvatarUri(d.Get("avatar_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("avatar_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(avatarUriProp)) && (ok || !reflect.DeepEqual(v, avatarUriProp)) {
		obj["avatarUri"] = avatarUriProp
	}
	speechToTextSettingsProp, err := expandDialogflowCXAgentSpeechToTextSettings(d.Get("speech_to_text_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("speech_to_text_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(speechToTextSettingsProp)) && (ok || !reflect.DeepEqual(v, speechToTextSettingsProp)) {
		obj["speechToTextSettings"] = speechToTextSettingsProp
	}
	securitySettingsProp, err := expandDialogflowCXAgentSecuritySettings(d.Get("security_settings"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("security_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(securitySettingsProp)) && (ok || !reflect.DeepEqual(v, securitySettingsProp)) {
		obj["securitySettings"] = securitySettingsProp
	}
	enableStackdriverLoggingProp, err := expandDialogflowCXAgentEnableStackdriverLogging(d.Get("enable_stackdriver_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_stackdriver_logging"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableStackdriverLoggingProp)) && (ok || !reflect.DeepEqual(v, enableStackdriverLoggingProp)) {
		obj["enableStackdriverLogging"] = enableStackdriverLoggingProp
	}
	enableSpellCorrectionProp, err := expandDialogflowCXAgentEnableSpellCorrection(d.Get("enable_spell_correction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_spell_correction"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableSpellCorrectionProp)) && (ok || !reflect.DeepEqual(v, enableSpellCorrectionProp)) {
		obj["enableSpellCorrection"] = enableSpellCorrectionProp
	}

	return obj, nil
}

func expandDialogflowCXAgentDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSupportedLanguageCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentAvatarUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSpeechToTextSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableSpeechAdaptation, err := expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(original["enable_speech_adaptation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableSpeechAdaptation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableSpeechAdaptation"] = transformedEnableSpeechAdaptation
	}

	return transformed, nil
}

func expandDialogflowCXAgentSpeechToTextSettingsEnableSpeechAdaptation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentSecuritySettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableStackdriverLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXAgentEnableSpellCorrection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}