/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderTimeInForceParameter the model 'NewOrderTimeInForceParameter'
type NewOrderTimeInForceParameter string

// List of newOrder_timeInForce_parameter
const (
	NewOrderTimeInForceParameterGtc NewOrderTimeInForceParameter = "GTC"
	NewOrderTimeInForceParameterIoc NewOrderTimeInForceParameter = "IOC"
	NewOrderTimeInForceParameterFok NewOrderTimeInForceParameter = "FOK"
	NewOrderTimeInForceParameterGtx NewOrderTimeInForceParameter = "GTX"
)

// All allowed values of NewOrderTimeInForceParameter enum
var AllowedNewOrderTimeInForceParameterEnumValues = []NewOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"GTX",
}

func (v *NewOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderTimeInForceParameter", value)
}

// NewNewOrderTimeInForceParameterFromValue returns a pointer to a valid NewOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderTimeInForceParameterFromValue(v string) (*NewOrderTimeInForceParameter, error) {
	ev := NewOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderTimeInForceParameter: valid values are %v", v, AllowedNewOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_timeInForce_parameter value
func (v NewOrderTimeInForceParameter) Ptr() *NewOrderTimeInForceParameter {
	return &v
}

type NullableNewOrderTimeInForceParameter struct {
	value *NewOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewOrderTimeInForceParameter) Get() *NewOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewOrderTimeInForceParameter) Set(val *NewOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderTimeInForceParameter(val *NewOrderTimeInForceParameter) *NullableNewOrderTimeInForceParameter {
	return &NullableNewOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
