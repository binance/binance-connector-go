/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// RedeemBfusdTypeParameter the model 'RedeemBfusdTypeParameter'
type RedeemBfusdTypeParameter string

// List of redeemBfusd_type_parameter
const (
	RedeemBfusdTypeParameterFast     RedeemBfusdTypeParameter = "FAST"
	RedeemBfusdTypeParameterStandard RedeemBfusdTypeParameter = "STANDARD"
)

// All allowed values of RedeemBfusdTypeParameter enum
var AllowedRedeemBfusdTypeParameterEnumValues = []RedeemBfusdTypeParameter{
	"FAST",
	"STANDARD",
}

func (v *RedeemBfusdTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedeemBfusdTypeParameter(value)
	for _, existing := range AllowedRedeemBfusdTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedeemBfusdTypeParameter", value)
}

// NewRedeemBfusdTypeParameterFromValue returns a pointer to a valid RedeemBfusdTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedeemBfusdTypeParameterFromValue(v string) (*RedeemBfusdTypeParameter, error) {
	ev := RedeemBfusdTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedeemBfusdTypeParameter: valid values are %v", v, AllowedRedeemBfusdTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedeemBfusdTypeParameter) IsValid() bool {
	for _, existing := range AllowedRedeemBfusdTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to redeemBfusd_type_parameter value
func (v RedeemBfusdTypeParameter) Ptr() *RedeemBfusdTypeParameter {
	return &v
}

type NullableRedeemBfusdTypeParameter struct {
	value *RedeemBfusdTypeParameter
	isSet bool
}

func (v NullableRedeemBfusdTypeParameter) Get() *RedeemBfusdTypeParameter {
	return v.value
}

func (v *NullableRedeemBfusdTypeParameter) Set(val *RedeemBfusdTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemBfusdTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemBfusdTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemBfusdTypeParameter(val *RedeemBfusdTypeParameter) *NullableRedeemBfusdTypeParameter {
	return &NullableRedeemBfusdTypeParameter{value: val, isSet: true}
}

func (v NullableRedeemBfusdTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemBfusdTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
