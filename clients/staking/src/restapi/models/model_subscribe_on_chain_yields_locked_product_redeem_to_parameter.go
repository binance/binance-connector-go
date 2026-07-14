/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SubscribeOnChainYieldsLockedProductRedeemToParameter the model 'SubscribeOnChainYieldsLockedProductRedeemToParameter'
type SubscribeOnChainYieldsLockedProductRedeemToParameter string

// List of subscribeOnChainYieldsLockedProduct_redeemTo_parameter
const (
	SubscribeOnChainYieldsLockedProductRedeemToParameterSpot     SubscribeOnChainYieldsLockedProductRedeemToParameter = "SPOT"
	SubscribeOnChainYieldsLockedProductRedeemToParameterFlexible SubscribeOnChainYieldsLockedProductRedeemToParameter = "FLEXIBLE"
)

// All allowed values of SubscribeOnChainYieldsLockedProductRedeemToParameter enum
var AllowedSubscribeOnChainYieldsLockedProductRedeemToParameterEnumValues = []SubscribeOnChainYieldsLockedProductRedeemToParameter{
	"SPOT",
	"FLEXIBLE",
}

func (v *SubscribeOnChainYieldsLockedProductRedeemToParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscribeOnChainYieldsLockedProductRedeemToParameter(value)
	for _, existing := range AllowedSubscribeOnChainYieldsLockedProductRedeemToParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscribeOnChainYieldsLockedProductRedeemToParameter", value)
}

// NewSubscribeOnChainYieldsLockedProductRedeemToParameterFromValue returns a pointer to a valid SubscribeOnChainYieldsLockedProductRedeemToParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeOnChainYieldsLockedProductRedeemToParameterFromValue(v string) (*SubscribeOnChainYieldsLockedProductRedeemToParameter, error) {
	ev := SubscribeOnChainYieldsLockedProductRedeemToParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeOnChainYieldsLockedProductRedeemToParameter: valid values are %v", v, AllowedSubscribeOnChainYieldsLockedProductRedeemToParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeOnChainYieldsLockedProductRedeemToParameter) IsValid() bool {
	for _, existing := range AllowedSubscribeOnChainYieldsLockedProductRedeemToParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribeOnChainYieldsLockedProduct_redeemTo_parameter value
func (v SubscribeOnChainYieldsLockedProductRedeemToParameter) Ptr() *SubscribeOnChainYieldsLockedProductRedeemToParameter {
	return &v
}

type NullableSubscribeOnChainYieldsLockedProductRedeemToParameter struct {
	value *SubscribeOnChainYieldsLockedProductRedeemToParameter
	isSet bool
}

func (v NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) Get() *SubscribeOnChainYieldsLockedProductRedeemToParameter {
	return v.value
}

func (v *NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) Set(val *SubscribeOnChainYieldsLockedProductRedeemToParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeOnChainYieldsLockedProductRedeemToParameter(val *SubscribeOnChainYieldsLockedProductRedeemToParameter) *NullableSubscribeOnChainYieldsLockedProductRedeemToParameter {
	return &NullableSubscribeOnChainYieldsLockedProductRedeemToParameter{value: val, isSet: true}
}

func (v NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeOnChainYieldsLockedProductRedeemToParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
