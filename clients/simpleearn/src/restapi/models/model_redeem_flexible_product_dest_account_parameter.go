/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// RedeemFlexibleProductDestAccountParameter the model 'RedeemFlexibleProductDestAccountParameter'
type RedeemFlexibleProductDestAccountParameter string

// List of redeemFlexibleProduct_destAccount_parameter
const (
	RedeemFlexibleProductDestAccountParameterSpot RedeemFlexibleProductDestAccountParameter = "SPOT"
	RedeemFlexibleProductDestAccountParameterFund RedeemFlexibleProductDestAccountParameter = "FUND"
)

// All allowed values of RedeemFlexibleProductDestAccountParameter enum
var AllowedRedeemFlexibleProductDestAccountParameterEnumValues = []RedeemFlexibleProductDestAccountParameter{
	"SPOT",
	"FUND",
}

func (v *RedeemFlexibleProductDestAccountParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedeemFlexibleProductDestAccountParameter(value)
	for _, existing := range AllowedRedeemFlexibleProductDestAccountParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedeemFlexibleProductDestAccountParameter", value)
}

// NewRedeemFlexibleProductDestAccountParameterFromValue returns a pointer to a valid RedeemFlexibleProductDestAccountParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedeemFlexibleProductDestAccountParameterFromValue(v string) (*RedeemFlexibleProductDestAccountParameter, error) {
	ev := RedeemFlexibleProductDestAccountParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedeemFlexibleProductDestAccountParameter: valid values are %v", v, AllowedRedeemFlexibleProductDestAccountParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedeemFlexibleProductDestAccountParameter) IsValid() bool {
	for _, existing := range AllowedRedeemFlexibleProductDestAccountParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to redeemFlexibleProduct_destAccount_parameter value
func (v RedeemFlexibleProductDestAccountParameter) Ptr() *RedeemFlexibleProductDestAccountParameter {
	return &v
}

type NullableRedeemFlexibleProductDestAccountParameter struct {
	value *RedeemFlexibleProductDestAccountParameter
	isSet bool
}

func (v NullableRedeemFlexibleProductDestAccountParameter) Get() *RedeemFlexibleProductDestAccountParameter {
	return v.value
}

func (v *NullableRedeemFlexibleProductDestAccountParameter) Set(val *RedeemFlexibleProductDestAccountParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemFlexibleProductDestAccountParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemFlexibleProductDestAccountParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemFlexibleProductDestAccountParameter(val *RedeemFlexibleProductDestAccountParameter) *NullableRedeemFlexibleProductDestAccountParameter {
	return &NullableRedeemFlexibleProductDestAccountParameter{value: val, isSet: true}
}

func (v NullableRedeemFlexibleProductDestAccountParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemFlexibleProductDestAccountParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
