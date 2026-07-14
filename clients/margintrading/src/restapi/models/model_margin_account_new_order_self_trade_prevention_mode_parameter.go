/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderSelfTradePreventionModeParameter the model 'MarginAccountNewOrderSelfTradePreventionModeParameter'
type MarginAccountNewOrderSelfTradePreventionModeParameter string

// List of marginAccountNewOrder_selfTradePreventionMode_parameter
const (
	MarginAccountNewOrderSelfTradePreventionModeParameterExpireTaker MarginAccountNewOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	MarginAccountNewOrderSelfTradePreventionModeParameterExpireMaker MarginAccountNewOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
	MarginAccountNewOrderSelfTradePreventionModeParameterExpireBoth  MarginAccountNewOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	MarginAccountNewOrderSelfTradePreventionModeParameterNone        MarginAccountNewOrderSelfTradePreventionModeParameter = "NONE"
)

// All allowed values of MarginAccountNewOrderSelfTradePreventionModeParameter enum
var AllowedMarginAccountNewOrderSelfTradePreventionModeParameterEnumValues = []MarginAccountNewOrderSelfTradePreventionModeParameter{
	"EXPIRE_TAKER",
	"EXPIRE_MAKER",
	"EXPIRE_BOTH",
	"NONE",
}

func (v *MarginAccountNewOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderSelfTradePreventionModeParameter", value)
}

// NewMarginAccountNewOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid MarginAccountNewOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderSelfTradePreventionModeParameterFromValue(v string) (*MarginAccountNewOrderSelfTradePreventionModeParameter, error) {
	ev := MarginAccountNewOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedMarginAccountNewOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_selfTradePreventionMode_parameter value
func (v MarginAccountNewOrderSelfTradePreventionModeParameter) Ptr() *MarginAccountNewOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableMarginAccountNewOrderSelfTradePreventionModeParameter struct {
	value *MarginAccountNewOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderSelfTradePreventionModeParameter) Get() *MarginAccountNewOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderSelfTradePreventionModeParameter) Set(val *MarginAccountNewOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderSelfTradePreventionModeParameter(val *MarginAccountNewOrderSelfTradePreventionModeParameter) *NullableMarginAccountNewOrderSelfTradePreventionModeParameter {
	return &NullableMarginAccountNewOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
