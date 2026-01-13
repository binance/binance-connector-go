/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmConditionalOrderPositionSideParameter the model 'NewCmConditionalOrderPositionSideParameter'
type NewCmConditionalOrderPositionSideParameter string

// List of newCmConditionalOrder_positionSide_parameter
const (
	NewCmConditionalOrderPositionSideParameterBoth  NewCmConditionalOrderPositionSideParameter = "BOTH"
	NewCmConditionalOrderPositionSideParameterLong  NewCmConditionalOrderPositionSideParameter = "LONG"
	NewCmConditionalOrderPositionSideParameterShort NewCmConditionalOrderPositionSideParameter = "SHORT"
)

// All allowed values of NewCmConditionalOrderPositionSideParameter enum
var AllowedNewCmConditionalOrderPositionSideParameterEnumValues = []NewCmConditionalOrderPositionSideParameter{
	"BOTH",
	"LONG",
	"SHORT",
}

func (v *NewCmConditionalOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmConditionalOrderPositionSideParameter(value)
	for _, existing := range AllowedNewCmConditionalOrderPositionSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmConditionalOrderPositionSideParameter", value)
}

// NewNewCmConditionalOrderPositionSideParameterFromValue returns a pointer to a valid NewCmConditionalOrderPositionSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmConditionalOrderPositionSideParameterFromValue(v string) (*NewCmConditionalOrderPositionSideParameter, error) {
	ev := NewCmConditionalOrderPositionSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmConditionalOrderPositionSideParameter: valid values are %v", v, AllowedNewCmConditionalOrderPositionSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmConditionalOrderPositionSideParameter) IsValid() bool {
	for _, existing := range AllowedNewCmConditionalOrderPositionSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmConditionalOrder_positionSide_parameter value
func (v NewCmConditionalOrderPositionSideParameter) Ptr() *NewCmConditionalOrderPositionSideParameter {
	return &v
}

type NullableNewCmConditionalOrderPositionSideParameter struct {
	value *NewCmConditionalOrderPositionSideParameter
	isSet bool
}

func (v NullableNewCmConditionalOrderPositionSideParameter) Get() *NewCmConditionalOrderPositionSideParameter {
	return v.value
}

func (v *NullableNewCmConditionalOrderPositionSideParameter) Set(val *NewCmConditionalOrderPositionSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmConditionalOrderPositionSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmConditionalOrderPositionSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmConditionalOrderPositionSideParameter(val *NewCmConditionalOrderPositionSideParameter) *NullableNewCmConditionalOrderPositionSideParameter {
	return &NullableNewCmConditionalOrderPositionSideParameter{value: val, isSet: true}
}

func (v NullableNewCmConditionalOrderPositionSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmConditionalOrderPositionSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
