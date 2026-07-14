/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SubscribeRwusdAssetParameter the model 'SubscribeRwusdAssetParameter'
type SubscribeRwusdAssetParameter string

// List of subscribeRwusd_asset_parameter
const (
	SubscribeRwusdAssetParameterUsdt SubscribeRwusdAssetParameter = "USDT"
	SubscribeRwusdAssetParameterUsdc SubscribeRwusdAssetParameter = "USDC"
)

// All allowed values of SubscribeRwusdAssetParameter enum
var AllowedSubscribeRwusdAssetParameterEnumValues = []SubscribeRwusdAssetParameter{
	"USDT",
	"USDC",
}

func (v *SubscribeRwusdAssetParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscribeRwusdAssetParameter(value)
	for _, existing := range AllowedSubscribeRwusdAssetParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscribeRwusdAssetParameter", value)
}

// NewSubscribeRwusdAssetParameterFromValue returns a pointer to a valid SubscribeRwusdAssetParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeRwusdAssetParameterFromValue(v string) (*SubscribeRwusdAssetParameter, error) {
	ev := SubscribeRwusdAssetParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeRwusdAssetParameter: valid values are %v", v, AllowedSubscribeRwusdAssetParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeRwusdAssetParameter) IsValid() bool {
	for _, existing := range AllowedSubscribeRwusdAssetParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribeRwusd_asset_parameter value
func (v SubscribeRwusdAssetParameter) Ptr() *SubscribeRwusdAssetParameter {
	return &v
}

type NullableSubscribeRwusdAssetParameter struct {
	value *SubscribeRwusdAssetParameter
	isSet bool
}

func (v NullableSubscribeRwusdAssetParameter) Get() *SubscribeRwusdAssetParameter {
	return v.value
}

func (v *NullableSubscribeRwusdAssetParameter) Set(val *SubscribeRwusdAssetParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeRwusdAssetParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeRwusdAssetParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeRwusdAssetParameter(val *SubscribeRwusdAssetParameter) *NullableSubscribeRwusdAssetParameter {
	return &NullableSubscribeRwusdAssetParameter{value: val, isSet: true}
}

func (v NullableSubscribeRwusdAssetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeRwusdAssetParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
