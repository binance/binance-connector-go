/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmConditionalOrderTimeInForceParameter the model 'NewUmConditionalOrderTimeInForceParameter'
type NewUmConditionalOrderTimeInForceParameter string

// List of newUmConditionalOrder_timeInForce_parameter
const (
	NewUmConditionalOrderTimeInForceParameterGtc NewUmConditionalOrderTimeInForceParameter = "GTC"
	NewUmConditionalOrderTimeInForceParameterIoc NewUmConditionalOrderTimeInForceParameter = "IOC"
	NewUmConditionalOrderTimeInForceParameterFok NewUmConditionalOrderTimeInForceParameter = "FOK"
	NewUmConditionalOrderTimeInForceParameterGtx NewUmConditionalOrderTimeInForceParameter = "GTX"
	NewUmConditionalOrderTimeInForceParameterGtd NewUmConditionalOrderTimeInForceParameter = "GTD"
)

// All allowed values of NewUmConditionalOrderTimeInForceParameter enum
var AllowedNewUmConditionalOrderTimeInForceParameterEnumValues = []NewUmConditionalOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
	"GTD",
}

func (v *NewUmConditionalOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmConditionalOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewUmConditionalOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmConditionalOrderTimeInForceParameter", value)
}

// NewNewUmConditionalOrderTimeInForceParameterFromValue returns a pointer to a valid NewUmConditionalOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmConditionalOrderTimeInForceParameterFromValue(v string) (*NewUmConditionalOrderTimeInForceParameter, error) {
	ev := NewUmConditionalOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmConditionalOrderTimeInForceParameter: valid values are %v", v, AllowedNewUmConditionalOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmConditionalOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewUmConditionalOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmConditionalOrder_timeInForce_parameter value
func (v NewUmConditionalOrderTimeInForceParameter) Ptr() *NewUmConditionalOrderTimeInForceParameter {
	return &v
}

type NullableNewUmConditionalOrderTimeInForceParameter struct {
	value *NewUmConditionalOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewUmConditionalOrderTimeInForceParameter) Get() *NewUmConditionalOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewUmConditionalOrderTimeInForceParameter) Set(val *NewUmConditionalOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmConditionalOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmConditionalOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmConditionalOrderTimeInForceParameter(val *NewUmConditionalOrderTimeInForceParameter) *NullableNewUmConditionalOrderTimeInForceParameter {
	return &NullableNewUmConditionalOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewUmConditionalOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmConditionalOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
