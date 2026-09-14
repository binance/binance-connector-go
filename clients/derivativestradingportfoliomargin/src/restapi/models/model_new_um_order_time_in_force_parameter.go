/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmOrderTimeInForceParameter the model 'NewUmOrderTimeInForceParameter'
type NewUmOrderTimeInForceParameter string

// List of newUmOrder_timeInForce_parameter
const (
	NewUmOrderTimeInForceParameterGtc NewUmOrderTimeInForceParameter = "GTC"
	NewUmOrderTimeInForceParameterIoc NewUmOrderTimeInForceParameter = "IOC"
	NewUmOrderTimeInForceParameterFok NewUmOrderTimeInForceParameter = "FOK"
	NewUmOrderTimeInForceParameterGtx NewUmOrderTimeInForceParameter = "GTX"
	NewUmOrderTimeInForceParameterGtd NewUmOrderTimeInForceParameter = "GTD"
)

// All allowed values of NewUmOrderTimeInForceParameter enum
var AllowedNewUmOrderTimeInForceParameterEnumValues = []NewUmOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
	"GTD",
}

func (v *NewUmOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewUmOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmOrderTimeInForceParameter", value)
}

// NewNewUmOrderTimeInForceParameterFromValue returns a pointer to a valid NewUmOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmOrderTimeInForceParameterFromValue(v string) (*NewUmOrderTimeInForceParameter, error) {
	ev := NewUmOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmOrderTimeInForceParameter: valid values are %v", v, AllowedNewUmOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewUmOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmOrder_timeInForce_parameter value
func (v NewUmOrderTimeInForceParameter) Ptr() *NewUmOrderTimeInForceParameter {
	return &v
}

type NullableNewUmOrderTimeInForceParameter struct {
	value *NewUmOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewUmOrderTimeInForceParameter) Get() *NewUmOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewUmOrderTimeInForceParameter) Set(val *NewUmOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmOrderTimeInForceParameter(val *NewUmOrderTimeInForceParameter) *NullableNewUmOrderTimeInForceParameter {
	return &NullableNewUmOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewUmOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
