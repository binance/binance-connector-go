/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderTimeInForceParameter the model 'MarginAccountNewOrderTimeInForceParameter'
type MarginAccountNewOrderTimeInForceParameter string

// List of marginAccountNewOrder_timeInForce_parameter
const (
	MarginAccountNewOrderTimeInForceParameterGtc MarginAccountNewOrderTimeInForceParameter = "GTC"
	MarginAccountNewOrderTimeInForceParameterIoc MarginAccountNewOrderTimeInForceParameter = "IOC"
	MarginAccountNewOrderTimeInForceParameterFok MarginAccountNewOrderTimeInForceParameter = "FOK"
)

// All allowed values of MarginAccountNewOrderTimeInForceParameter enum
var AllowedMarginAccountNewOrderTimeInForceParameterEnumValues = []MarginAccountNewOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *MarginAccountNewOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderTimeInForceParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderTimeInForceParameter", value)
}

// NewMarginAccountNewOrderTimeInForceParameterFromValue returns a pointer to a valid MarginAccountNewOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderTimeInForceParameterFromValue(v string) (*MarginAccountNewOrderTimeInForceParameter, error) {
	ev := MarginAccountNewOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderTimeInForceParameter: valid values are %v", v, AllowedMarginAccountNewOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_timeInForce_parameter value
func (v MarginAccountNewOrderTimeInForceParameter) Ptr() *MarginAccountNewOrderTimeInForceParameter {
	return &v
}

type NullableMarginAccountNewOrderTimeInForceParameter struct {
	value *MarginAccountNewOrderTimeInForceParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderTimeInForceParameter) Get() *MarginAccountNewOrderTimeInForceParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderTimeInForceParameter) Set(val *MarginAccountNewOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderTimeInForceParameter(val *MarginAccountNewOrderTimeInForceParameter) *NullableMarginAccountNewOrderTimeInForceParameter {
	return &NullableMarginAccountNewOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
