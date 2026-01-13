/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmOrderTypeParameter the model 'NewCmOrderTypeParameter'
type NewCmOrderTypeParameter string

// List of newCmOrder_type_parameter
const (
	NewCmOrderTypeParameterLimit  NewCmOrderTypeParameter = "LIMIT"
	NewCmOrderTypeParameterMarket NewCmOrderTypeParameter = "MARKET"
)

// All allowed values of NewCmOrderTypeParameter enum
var AllowedNewCmOrderTypeParameterEnumValues = []NewCmOrderTypeParameter{
	"LIMIT",
	"MARKET",
}

func (v *NewCmOrderTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmOrderTypeParameter(value)
	for _, existing := range AllowedNewCmOrderTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmOrderTypeParameter", value)
}

// NewNewCmOrderTypeParameterFromValue returns a pointer to a valid NewCmOrderTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmOrderTypeParameterFromValue(v string) (*NewCmOrderTypeParameter, error) {
	ev := NewCmOrderTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmOrderTypeParameter: valid values are %v", v, AllowedNewCmOrderTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmOrderTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewCmOrderTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmOrder_type_parameter value
func (v NewCmOrderTypeParameter) Ptr() *NewCmOrderTypeParameter {
	return &v
}

type NullableNewCmOrderTypeParameter struct {
	value *NewCmOrderTypeParameter
	isSet bool
}

func (v NullableNewCmOrderTypeParameter) Get() *NewCmOrderTypeParameter {
	return v.value
}

func (v *NullableNewCmOrderTypeParameter) Set(val *NewCmOrderTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmOrderTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmOrderTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmOrderTypeParameter(val *NewCmOrderTypeParameter) *NullableNewCmOrderTypeParameter {
	return &NullableNewCmOrderTypeParameter{value: val, isSet: true}
}

func (v NullableNewCmOrderTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmOrderTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
