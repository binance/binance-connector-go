/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SetLockedProductRedeemOptionRedeemToParameter the model 'SetLockedProductRedeemOptionRedeemToParameter'
type SetLockedProductRedeemOptionRedeemToParameter string

// List of setLockedProductRedeemOption_redeemTo_parameter
const (
	SetLockedProductRedeemOptionRedeemToParameterSpot     SetLockedProductRedeemOptionRedeemToParameter = "SPOT"
	SetLockedProductRedeemOptionRedeemToParameterFlexible SetLockedProductRedeemOptionRedeemToParameter = "FLEXIBLE"
)

// All allowed values of SetLockedProductRedeemOptionRedeemToParameter enum
var AllowedSetLockedProductRedeemOptionRedeemToParameterEnumValues = []SetLockedProductRedeemOptionRedeemToParameter{
	"SPOT",
	"FLEXIBLE",
}

func (v *SetLockedProductRedeemOptionRedeemToParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SetLockedProductRedeemOptionRedeemToParameter(value)
	for _, existing := range AllowedSetLockedProductRedeemOptionRedeemToParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SetLockedProductRedeemOptionRedeemToParameter", value)
}

// NewSetLockedProductRedeemOptionRedeemToParameterFromValue returns a pointer to a valid SetLockedProductRedeemOptionRedeemToParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSetLockedProductRedeemOptionRedeemToParameterFromValue(v string) (*SetLockedProductRedeemOptionRedeemToParameter, error) {
	ev := SetLockedProductRedeemOptionRedeemToParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SetLockedProductRedeemOptionRedeemToParameter: valid values are %v", v, AllowedSetLockedProductRedeemOptionRedeemToParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SetLockedProductRedeemOptionRedeemToParameter) IsValid() bool {
	for _, existing := range AllowedSetLockedProductRedeemOptionRedeemToParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to setLockedProductRedeemOption_redeemTo_parameter value
func (v SetLockedProductRedeemOptionRedeemToParameter) Ptr() *SetLockedProductRedeemOptionRedeemToParameter {
	return &v
}

type NullableSetLockedProductRedeemOptionRedeemToParameter struct {
	value *SetLockedProductRedeemOptionRedeemToParameter
	isSet bool
}

func (v NullableSetLockedProductRedeemOptionRedeemToParameter) Get() *SetLockedProductRedeemOptionRedeemToParameter {
	return v.value
}

func (v *NullableSetLockedProductRedeemOptionRedeemToParameter) Set(val *SetLockedProductRedeemOptionRedeemToParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSetLockedProductRedeemOptionRedeemToParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSetLockedProductRedeemOptionRedeemToParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetLockedProductRedeemOptionRedeemToParameter(val *SetLockedProductRedeemOptionRedeemToParameter) *NullableSetLockedProductRedeemOptionRedeemToParameter {
	return &NullableSetLockedProductRedeemOptionRedeemToParameter{value: val, isSet: true}
}

func (v NullableSetLockedProductRedeemOptionRedeemToParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetLockedProductRedeemOptionRedeemToParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
