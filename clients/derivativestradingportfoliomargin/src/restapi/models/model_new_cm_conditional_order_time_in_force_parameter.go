/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmConditionalOrderTimeInForceParameter the model 'NewCmConditionalOrderTimeInForceParameter'
type NewCmConditionalOrderTimeInForceParameter string

// List of newCmConditionalOrder_timeInForce_parameter
const (
	NewCmConditionalOrderTimeInForceParameterGtc NewCmConditionalOrderTimeInForceParameter = "GTC"
	NewCmConditionalOrderTimeInForceParameterIoc NewCmConditionalOrderTimeInForceParameter = "IOC"
	NewCmConditionalOrderTimeInForceParameterFok NewCmConditionalOrderTimeInForceParameter = "FOK"
	NewCmConditionalOrderTimeInForceParameterGtx NewCmConditionalOrderTimeInForceParameter = "GTX"
)

// All allowed values of NewCmConditionalOrderTimeInForceParameter enum
var AllowedNewCmConditionalOrderTimeInForceParameterEnumValues = []NewCmConditionalOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
}

func (v *NewCmConditionalOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmConditionalOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewCmConditionalOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmConditionalOrderTimeInForceParameter", value)
}

// NewNewCmConditionalOrderTimeInForceParameterFromValue returns a pointer to a valid NewCmConditionalOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmConditionalOrderTimeInForceParameterFromValue(v string) (*NewCmConditionalOrderTimeInForceParameter, error) {
	ev := NewCmConditionalOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmConditionalOrderTimeInForceParameter: valid values are %v", v, AllowedNewCmConditionalOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmConditionalOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewCmConditionalOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmConditionalOrder_timeInForce_parameter value
func (v NewCmConditionalOrderTimeInForceParameter) Ptr() *NewCmConditionalOrderTimeInForceParameter {
	return &v
}

type NullableNewCmConditionalOrderTimeInForceParameter struct {
	value *NewCmConditionalOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewCmConditionalOrderTimeInForceParameter) Get() *NewCmConditionalOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewCmConditionalOrderTimeInForceParameter) Set(val *NewCmConditionalOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmConditionalOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmConditionalOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmConditionalOrderTimeInForceParameter(val *NewCmConditionalOrderTimeInForceParameter) *NullableNewCmConditionalOrderTimeInForceParameter {
	return &NullableNewCmConditionalOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewCmConditionalOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmConditionalOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
