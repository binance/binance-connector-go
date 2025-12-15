/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmConditionalOrderWorkingTypeParameter the model 'NewCmConditionalOrderWorkingTypeParameter'
type NewCmConditionalOrderWorkingTypeParameter string

// List of newCmConditionalOrder_workingType_parameter
const (
	NewCmConditionalOrderWorkingTypeParameterMarkPrice NewCmConditionalOrderWorkingTypeParameter = "MARK_PRICE"
)

// All allowed values of NewCmConditionalOrderWorkingTypeParameter enum
var AllowedNewCmConditionalOrderWorkingTypeParameterEnumValues = []NewCmConditionalOrderWorkingTypeParameter{
	"MARK_PRICE",
}

func (v *NewCmConditionalOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmConditionalOrderWorkingTypeParameter(value)
	for _, existing := range AllowedNewCmConditionalOrderWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmConditionalOrderWorkingTypeParameter", value)
}

// NewNewCmConditionalOrderWorkingTypeParameterFromValue returns a pointer to a valid NewCmConditionalOrderWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmConditionalOrderWorkingTypeParameterFromValue(v string) (*NewCmConditionalOrderWorkingTypeParameter, error) {
	ev := NewCmConditionalOrderWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmConditionalOrderWorkingTypeParameter: valid values are %v", v, AllowedNewCmConditionalOrderWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmConditionalOrderWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewCmConditionalOrderWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmConditionalOrder_workingType_parameter value
func (v NewCmConditionalOrderWorkingTypeParameter) Ptr() *NewCmConditionalOrderWorkingTypeParameter {
	return &v
}

type NullableNewCmConditionalOrderWorkingTypeParameter struct {
	value *NewCmConditionalOrderWorkingTypeParameter
	isSet bool
}

func (v NullableNewCmConditionalOrderWorkingTypeParameter) Get() *NewCmConditionalOrderWorkingTypeParameter {
	return v.value
}

func (v *NullableNewCmConditionalOrderWorkingTypeParameter) Set(val *NewCmConditionalOrderWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmConditionalOrderWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmConditionalOrderWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmConditionalOrderWorkingTypeParameter(val *NewCmConditionalOrderWorkingTypeParameter) *NullableNewCmConditionalOrderWorkingTypeParameter {
	return &NullableNewCmConditionalOrderWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableNewCmConditionalOrderWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmConditionalOrderWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
