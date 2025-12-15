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

// MarginAccountNewOcoStopLimitTimeInForceParameter the model 'MarginAccountNewOcoStopLimitTimeInForceParameter'
type MarginAccountNewOcoStopLimitTimeInForceParameter string

// List of marginAccountNewOco_stopLimitTimeInForce_parameter
const (
	MarginAccountNewOcoStopLimitTimeInForceParameterGtc MarginAccountNewOcoStopLimitTimeInForceParameter = "GTC"
	MarginAccountNewOcoStopLimitTimeInForceParameterIoc MarginAccountNewOcoStopLimitTimeInForceParameter = "IOC"
	MarginAccountNewOcoStopLimitTimeInForceParameterFok MarginAccountNewOcoStopLimitTimeInForceParameter = "FOK"
)

// All allowed values of MarginAccountNewOcoStopLimitTimeInForceParameter enum
var AllowedMarginAccountNewOcoStopLimitTimeInForceParameterEnumValues = []MarginAccountNewOcoStopLimitTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *MarginAccountNewOcoStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOcoStopLimitTimeInForceParameter(value)
	for _, existing := range AllowedMarginAccountNewOcoStopLimitTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOcoStopLimitTimeInForceParameter", value)
}

// NewMarginAccountNewOcoStopLimitTimeInForceParameterFromValue returns a pointer to a valid MarginAccountNewOcoStopLimitTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOcoStopLimitTimeInForceParameterFromValue(v string) (*MarginAccountNewOcoStopLimitTimeInForceParameter, error) {
	ev := MarginAccountNewOcoStopLimitTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOcoStopLimitTimeInForceParameter: valid values are %v", v, AllowedMarginAccountNewOcoStopLimitTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOcoStopLimitTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOcoStopLimitTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOco_stopLimitTimeInForce_parameter value
func (v MarginAccountNewOcoStopLimitTimeInForceParameter) Ptr() *MarginAccountNewOcoStopLimitTimeInForceParameter {
	return &v
}

type NullableMarginAccountNewOcoStopLimitTimeInForceParameter struct {
	value *MarginAccountNewOcoStopLimitTimeInForceParameter
	isSet bool
}

func (v NullableMarginAccountNewOcoStopLimitTimeInForceParameter) Get() *MarginAccountNewOcoStopLimitTimeInForceParameter {
	return v.value
}

func (v *NullableMarginAccountNewOcoStopLimitTimeInForceParameter) Set(val *MarginAccountNewOcoStopLimitTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOcoStopLimitTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOcoStopLimitTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOcoStopLimitTimeInForceParameter(val *MarginAccountNewOcoStopLimitTimeInForceParameter) *NullableMarginAccountNewOcoStopLimitTimeInForceParameter {
	return &NullableMarginAccountNewOcoStopLimitTimeInForceParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOcoStopLimitTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOcoStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
