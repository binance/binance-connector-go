/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderNewOrderRespTypeParameter the model 'MarginAccountNewOrderNewOrderRespTypeParameter'
type MarginAccountNewOrderNewOrderRespTypeParameter string

// List of marginAccountNewOrder_newOrderRespType_parameter
const (
	MarginAccountNewOrderNewOrderRespTypeParameterAck    MarginAccountNewOrderNewOrderRespTypeParameter = "ACK"
	MarginAccountNewOrderNewOrderRespTypeParameterResult MarginAccountNewOrderNewOrderRespTypeParameter = "RESULT"
	MarginAccountNewOrderNewOrderRespTypeParameterFull   MarginAccountNewOrderNewOrderRespTypeParameter = "FULL"
)

// All allowed values of MarginAccountNewOrderNewOrderRespTypeParameter enum
var AllowedMarginAccountNewOrderNewOrderRespTypeParameterEnumValues = []MarginAccountNewOrderNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
	"FULL",
}

func (v *MarginAccountNewOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderNewOrderRespTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderNewOrderRespTypeParameter", value)
}

// NewMarginAccountNewOrderNewOrderRespTypeParameterFromValue returns a pointer to a valid MarginAccountNewOrderNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderNewOrderRespTypeParameterFromValue(v string) (*MarginAccountNewOrderNewOrderRespTypeParameter, error) {
	ev := MarginAccountNewOrderNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderNewOrderRespTypeParameter: valid values are %v", v, AllowedMarginAccountNewOrderNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_newOrderRespType_parameter value
func (v MarginAccountNewOrderNewOrderRespTypeParameter) Ptr() *MarginAccountNewOrderNewOrderRespTypeParameter {
	return &v
}

type NullableMarginAccountNewOrderNewOrderRespTypeParameter struct {
	value *MarginAccountNewOrderNewOrderRespTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderNewOrderRespTypeParameter) Get() *MarginAccountNewOrderNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderNewOrderRespTypeParameter) Set(val *MarginAccountNewOrderNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderNewOrderRespTypeParameter(val *MarginAccountNewOrderNewOrderRespTypeParameter) *NullableMarginAccountNewOrderNewOrderRespTypeParameter {
	return &NullableMarginAccountNewOrderNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
