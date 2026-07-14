/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SubscribeOnChainYieldsLockedProductSourceAccountParameter the model 'SubscribeOnChainYieldsLockedProductSourceAccountParameter'
type SubscribeOnChainYieldsLockedProductSourceAccountParameter string

// List of subscribeOnChainYieldsLockedProduct_sourceAccount_parameter
const (
	SubscribeOnChainYieldsLockedProductSourceAccountParameterSpot SubscribeOnChainYieldsLockedProductSourceAccountParameter = "SPOT"
	SubscribeOnChainYieldsLockedProductSourceAccountParameterFund SubscribeOnChainYieldsLockedProductSourceAccountParameter = "FUND"
	SubscribeOnChainYieldsLockedProductSourceAccountParameterAll  SubscribeOnChainYieldsLockedProductSourceAccountParameter = "ALL"
)

// All allowed values of SubscribeOnChainYieldsLockedProductSourceAccountParameter enum
var AllowedSubscribeOnChainYieldsLockedProductSourceAccountParameterEnumValues = []SubscribeOnChainYieldsLockedProductSourceAccountParameter{
	"SPOT",
	"FUND",
	"ALL",
}

func (v *SubscribeOnChainYieldsLockedProductSourceAccountParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscribeOnChainYieldsLockedProductSourceAccountParameter(value)
	for _, existing := range AllowedSubscribeOnChainYieldsLockedProductSourceAccountParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscribeOnChainYieldsLockedProductSourceAccountParameter", value)
}

// NewSubscribeOnChainYieldsLockedProductSourceAccountParameterFromValue returns a pointer to a valid SubscribeOnChainYieldsLockedProductSourceAccountParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeOnChainYieldsLockedProductSourceAccountParameterFromValue(v string) (*SubscribeOnChainYieldsLockedProductSourceAccountParameter, error) {
	ev := SubscribeOnChainYieldsLockedProductSourceAccountParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeOnChainYieldsLockedProductSourceAccountParameter: valid values are %v", v, AllowedSubscribeOnChainYieldsLockedProductSourceAccountParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeOnChainYieldsLockedProductSourceAccountParameter) IsValid() bool {
	for _, existing := range AllowedSubscribeOnChainYieldsLockedProductSourceAccountParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribeOnChainYieldsLockedProduct_sourceAccount_parameter value
func (v SubscribeOnChainYieldsLockedProductSourceAccountParameter) Ptr() *SubscribeOnChainYieldsLockedProductSourceAccountParameter {
	return &v
}

type NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter struct {
	value *SubscribeOnChainYieldsLockedProductSourceAccountParameter
	isSet bool
}

func (v NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) Get() *SubscribeOnChainYieldsLockedProductSourceAccountParameter {
	return v.value
}

func (v *NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) Set(val *SubscribeOnChainYieldsLockedProductSourceAccountParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeOnChainYieldsLockedProductSourceAccountParameter(val *SubscribeOnChainYieldsLockedProductSourceAccountParameter) *NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter {
	return &NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter{value: val, isSet: true}
}

func (v NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeOnChainYieldsLockedProductSourceAccountParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
