/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderTimeInForceParameter the model 'NewAlgoOrderTimeInForceParameter'
type NewAlgoOrderTimeInForceParameter string

// List of newAlgoOrder_timeInForce_parameter
const (
	NewAlgoOrderTimeInForceParameterGtc NewAlgoOrderTimeInForceParameter = "GTC"
	NewAlgoOrderTimeInForceParameterIoc NewAlgoOrderTimeInForceParameter = "IOC"
	NewAlgoOrderTimeInForceParameterFok NewAlgoOrderTimeInForceParameter = "FOK"
	NewAlgoOrderTimeInForceParameterGtx NewAlgoOrderTimeInForceParameter = "GTX"
	NewAlgoOrderTimeInForceParameterGtd NewAlgoOrderTimeInForceParameter = "GTD"
	NewAlgoOrderTimeInForceParameterRpi NewAlgoOrderTimeInForceParameter = "RPI"
)

// All allowed values of NewAlgoOrderTimeInForceParameter enum
var AllowedNewAlgoOrderTimeInForceParameterEnumValues = []NewAlgoOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
	"GTD",
	"RPI",
}

func (v *NewAlgoOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewAlgoOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderTimeInForceParameter", value)
}

// NewNewAlgoOrderTimeInForceParameterFromValue returns a pointer to a valid NewAlgoOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderTimeInForceParameterFromValue(v string) (*NewAlgoOrderTimeInForceParameter, error) {
	ev := NewAlgoOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderTimeInForceParameter: valid values are %v", v, AllowedNewAlgoOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_timeInForce_parameter value
func (v NewAlgoOrderTimeInForceParameter) Ptr() *NewAlgoOrderTimeInForceParameter {
	return &v
}

type NullableNewAlgoOrderTimeInForceParameter struct {
	value *NewAlgoOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewAlgoOrderTimeInForceParameter) Get() *NewAlgoOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewAlgoOrderTimeInForceParameter) Set(val *NewAlgoOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderTimeInForceParameter(val *NewAlgoOrderTimeInForceParameter) *NullableNewAlgoOrderTimeInForceParameter {
	return &NullableNewAlgoOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
