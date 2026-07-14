/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter the model 'SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter'
type SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter string

// List of setOnChainYieldsLockedProductRedeemOption_redeemTo_parameter
const (
	SetOnChainYieldsLockedProductRedeemOptionRedeemToParameterSpot     SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter = "SPOT"
	SetOnChainYieldsLockedProductRedeemOptionRedeemToParameterFlexible SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter = "FLEXIBLE"
)

// All allowed values of SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter enum
var AllowedSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterEnumValues = []SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter{
	"SPOT",
	"FLEXIBLE",
}

func (v *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter(value)
	for _, existing := range AllowedSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter", value)
}

// NewSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterFromValue returns a pointer to a valid SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterFromValue(v string) (*SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter, error) {
	ev := SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter: valid values are %v", v, AllowedSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) IsValid() bool {
	for _, existing := range AllowedSetOnChainYieldsLockedProductRedeemOptionRedeemToParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to setOnChainYieldsLockedProductRedeemOption_redeemTo_parameter value
func (v SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) Ptr() *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter {
	return &v
}

type NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter struct {
	value *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter
	isSet bool
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) Get() *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter {
	return v.value
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) Set(val *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter(val *SetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) *NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter {
	return &NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter{value: val, isSet: true}
}

func (v NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOnChainYieldsLockedProductRedeemOptionRedeemToParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
