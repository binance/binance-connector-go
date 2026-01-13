/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderSideParameter the model 'MarginAccountNewOrderSideParameter'
type MarginAccountNewOrderSideParameter string

// List of marginAccountNewOrder_side_parameter
const (
	MarginAccountNewOrderSideParameterBuy  MarginAccountNewOrderSideParameter = "BUY"
	MarginAccountNewOrderSideParameterSell MarginAccountNewOrderSideParameter = "SELL"
)

// All allowed values of MarginAccountNewOrderSideParameter enum
var AllowedMarginAccountNewOrderSideParameterEnumValues = []MarginAccountNewOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *MarginAccountNewOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderSideParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderSideParameter", value)
}

// NewMarginAccountNewOrderSideParameterFromValue returns a pointer to a valid MarginAccountNewOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderSideParameterFromValue(v string) (*MarginAccountNewOrderSideParameter, error) {
	ev := MarginAccountNewOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderSideParameter: valid values are %v", v, AllowedMarginAccountNewOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_side_parameter value
func (v MarginAccountNewOrderSideParameter) Ptr() *MarginAccountNewOrderSideParameter {
	return &v
}

type NullableMarginAccountNewOrderSideParameter struct {
	value *MarginAccountNewOrderSideParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderSideParameter) Get() *MarginAccountNewOrderSideParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderSideParameter) Set(val *MarginAccountNewOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderSideParameter(val *MarginAccountNewOrderSideParameter) *NullableMarginAccountNewOrderSideParameter {
	return &NullableMarginAccountNewOrderSideParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
