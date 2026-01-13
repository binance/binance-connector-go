/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ChangeMarginTypeMarginTypeParameter the model 'ChangeMarginTypeMarginTypeParameter'
type ChangeMarginTypeMarginTypeParameter string

// List of changeMarginType_marginType_parameter
const (
	ChangeMarginTypeMarginTypeParameterIsolated ChangeMarginTypeMarginTypeParameter = "ISOLATED"
	ChangeMarginTypeMarginTypeParameterCrossed  ChangeMarginTypeMarginTypeParameter = "CROSSED"
)

// All allowed values of ChangeMarginTypeMarginTypeParameter enum
var AllowedChangeMarginTypeMarginTypeParameterEnumValues = []ChangeMarginTypeMarginTypeParameter{
	"ISOLATED",
	"CROSSED",
}

func (v *ChangeMarginTypeMarginTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChangeMarginTypeMarginTypeParameter(value)
	for _, existing := range AllowedChangeMarginTypeMarginTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChangeMarginTypeMarginTypeParameter", value)
}

// NewChangeMarginTypeMarginTypeParameterFromValue returns a pointer to a valid ChangeMarginTypeMarginTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChangeMarginTypeMarginTypeParameterFromValue(v string) (*ChangeMarginTypeMarginTypeParameter, error) {
	ev := ChangeMarginTypeMarginTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChangeMarginTypeMarginTypeParameter: valid values are %v", v, AllowedChangeMarginTypeMarginTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChangeMarginTypeMarginTypeParameter) IsValid() bool {
	for _, existing := range AllowedChangeMarginTypeMarginTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to changeMarginType_marginType_parameter value
func (v ChangeMarginTypeMarginTypeParameter) Ptr() *ChangeMarginTypeMarginTypeParameter {
	return &v
}

type NullableChangeMarginTypeMarginTypeParameter struct {
	value *ChangeMarginTypeMarginTypeParameter
	isSet bool
}

func (v NullableChangeMarginTypeMarginTypeParameter) Get() *ChangeMarginTypeMarginTypeParameter {
	return v.value
}

func (v *NullableChangeMarginTypeMarginTypeParameter) Set(val *ChangeMarginTypeMarginTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeMarginTypeMarginTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeMarginTypeMarginTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeMarginTypeMarginTypeParameter(val *ChangeMarginTypeMarginTypeParameter) *NullableChangeMarginTypeMarginTypeParameter {
	return &NullableChangeMarginTypeMarginTypeParameter{value: val, isSet: true}
}

func (v NullableChangeMarginTypeMarginTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeMarginTypeMarginTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
