/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetBfusdSubscriptionHistoryAssetParameter the model 'GetBfusdSubscriptionHistoryAssetParameter'
type GetBfusdSubscriptionHistoryAssetParameter string

// List of getBfusdSubscriptionHistory_asset_parameter
const (
	GetBfusdSubscriptionHistoryAssetParameterUsdc GetBfusdSubscriptionHistoryAssetParameter = "USDC"
	GetBfusdSubscriptionHistoryAssetParameterUsdt GetBfusdSubscriptionHistoryAssetParameter = "USDT"
)

// All allowed values of GetBfusdSubscriptionHistoryAssetParameter enum
var AllowedGetBfusdSubscriptionHistoryAssetParameterEnumValues = []GetBfusdSubscriptionHistoryAssetParameter{
	"USDC",
	"USDT",
}

func (v *GetBfusdSubscriptionHistoryAssetParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetBfusdSubscriptionHistoryAssetParameter(value)
	for _, existing := range AllowedGetBfusdSubscriptionHistoryAssetParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetBfusdSubscriptionHistoryAssetParameter", value)
}

// NewGetBfusdSubscriptionHistoryAssetParameterFromValue returns a pointer to a valid GetBfusdSubscriptionHistoryAssetParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetBfusdSubscriptionHistoryAssetParameterFromValue(v string) (*GetBfusdSubscriptionHistoryAssetParameter, error) {
	ev := GetBfusdSubscriptionHistoryAssetParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetBfusdSubscriptionHistoryAssetParameter: valid values are %v", v, AllowedGetBfusdSubscriptionHistoryAssetParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetBfusdSubscriptionHistoryAssetParameter) IsValid() bool {
	for _, existing := range AllowedGetBfusdSubscriptionHistoryAssetParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getBfusdSubscriptionHistory_asset_parameter value
func (v GetBfusdSubscriptionHistoryAssetParameter) Ptr() *GetBfusdSubscriptionHistoryAssetParameter {
	return &v
}

type NullableGetBfusdSubscriptionHistoryAssetParameter struct {
	value *GetBfusdSubscriptionHistoryAssetParameter
	isSet bool
}

func (v NullableGetBfusdSubscriptionHistoryAssetParameter) Get() *GetBfusdSubscriptionHistoryAssetParameter {
	return v.value
}

func (v *NullableGetBfusdSubscriptionHistoryAssetParameter) Set(val *GetBfusdSubscriptionHistoryAssetParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBfusdSubscriptionHistoryAssetParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBfusdSubscriptionHistoryAssetParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBfusdSubscriptionHistoryAssetParameter(val *GetBfusdSubscriptionHistoryAssetParameter) *NullableGetBfusdSubscriptionHistoryAssetParameter {
	return &NullableGetBfusdSubscriptionHistoryAssetParameter{value: val, isSet: true}
}

func (v NullableGetBfusdSubscriptionHistoryAssetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBfusdSubscriptionHistoryAssetParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
