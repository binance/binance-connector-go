/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// SubscribeFlexibleProductSourceAccountParameter the model 'SubscribeFlexibleProductSourceAccountParameter'
type SubscribeFlexibleProductSourceAccountParameter string

// List of subscribeFlexibleProduct_sourceAccount_parameter
const (
	SubscribeFlexibleProductSourceAccountParameterSpot SubscribeFlexibleProductSourceAccountParameter = "SPOT"
	SubscribeFlexibleProductSourceAccountParameterFund SubscribeFlexibleProductSourceAccountParameter = "FUND"
	SubscribeFlexibleProductSourceAccountParameterAll  SubscribeFlexibleProductSourceAccountParameter = "ALL"
)

// All allowed values of SubscribeFlexibleProductSourceAccountParameter enum
var AllowedSubscribeFlexibleProductSourceAccountParameterEnumValues = []SubscribeFlexibleProductSourceAccountParameter{
	"SPOT",
	"FUND",
	"ALL",
}

func (v *SubscribeFlexibleProductSourceAccountParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscribeFlexibleProductSourceAccountParameter(value)
	for _, existing := range AllowedSubscribeFlexibleProductSourceAccountParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscribeFlexibleProductSourceAccountParameter", value)
}

// NewSubscribeFlexibleProductSourceAccountParameterFromValue returns a pointer to a valid SubscribeFlexibleProductSourceAccountParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribeFlexibleProductSourceAccountParameterFromValue(v string) (*SubscribeFlexibleProductSourceAccountParameter, error) {
	ev := SubscribeFlexibleProductSourceAccountParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribeFlexibleProductSourceAccountParameter: valid values are %v", v, AllowedSubscribeFlexibleProductSourceAccountParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribeFlexibleProductSourceAccountParameter) IsValid() bool {
	for _, existing := range AllowedSubscribeFlexibleProductSourceAccountParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to subscribeFlexibleProduct_sourceAccount_parameter value
func (v SubscribeFlexibleProductSourceAccountParameter) Ptr() *SubscribeFlexibleProductSourceAccountParameter {
	return &v
}

type NullableSubscribeFlexibleProductSourceAccountParameter struct {
	value *SubscribeFlexibleProductSourceAccountParameter
	isSet bool
}

func (v NullableSubscribeFlexibleProductSourceAccountParameter) Get() *SubscribeFlexibleProductSourceAccountParameter {
	return v.value
}

func (v *NullableSubscribeFlexibleProductSourceAccountParameter) Set(val *SubscribeFlexibleProductSourceAccountParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeFlexibleProductSourceAccountParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeFlexibleProductSourceAccountParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeFlexibleProductSourceAccountParameter(val *SubscribeFlexibleProductSourceAccountParameter) *NullableSubscribeFlexibleProductSourceAccountParameter {
	return &NullableSubscribeFlexibleProductSourceAccountParameter{value: val, isSet: true}
}

func (v NullableSubscribeFlexibleProductSourceAccountParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeFlexibleProductSourceAccountParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
