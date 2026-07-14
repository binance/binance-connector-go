/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOrderSideEffectTypeParameter the model 'MarginAccountNewOrderSideEffectTypeParameter'
type MarginAccountNewOrderSideEffectTypeParameter string

// List of marginAccountNewOrder_sideEffectType_parameter
const (
	MarginAccountNewOrderSideEffectTypeParameterNoSideEffect    MarginAccountNewOrderSideEffectTypeParameter = "NO_SIDE_EFFECT"
	MarginAccountNewOrderSideEffectTypeParameterMarginBuy       MarginAccountNewOrderSideEffectTypeParameter = "MARGIN_BUY"
	MarginAccountNewOrderSideEffectTypeParameterAutoRepay       MarginAccountNewOrderSideEffectTypeParameter = "AUTO_REPAY"
	MarginAccountNewOrderSideEffectTypeParameterAutoBorrowRepay MarginAccountNewOrderSideEffectTypeParameter = "AUTO_BORROW_REPAY"
)

// All allowed values of MarginAccountNewOrderSideEffectTypeParameter enum
var AllowedMarginAccountNewOrderSideEffectTypeParameterEnumValues = []MarginAccountNewOrderSideEffectTypeParameter{
	"NO_SIDE_EFFECT",
	"MARGIN_BUY",
	"AUTO_REPAY",
	"AUTO_BORROW_REPAY",
}

func (v *MarginAccountNewOrderSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOrderSideEffectTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOrderSideEffectTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOrderSideEffectTypeParameter", value)
}

// NewMarginAccountNewOrderSideEffectTypeParameterFromValue returns a pointer to a valid MarginAccountNewOrderSideEffectTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOrderSideEffectTypeParameterFromValue(v string) (*MarginAccountNewOrderSideEffectTypeParameter, error) {
	ev := MarginAccountNewOrderSideEffectTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOrderSideEffectTypeParameter: valid values are %v", v, AllowedMarginAccountNewOrderSideEffectTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOrderSideEffectTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOrderSideEffectTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOrder_sideEffectType_parameter value
func (v MarginAccountNewOrderSideEffectTypeParameter) Ptr() *MarginAccountNewOrderSideEffectTypeParameter {
	return &v
}

type NullableMarginAccountNewOrderSideEffectTypeParameter struct {
	value *MarginAccountNewOrderSideEffectTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOrderSideEffectTypeParameter) Get() *MarginAccountNewOrderSideEffectTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOrderSideEffectTypeParameter) Set(val *MarginAccountNewOrderSideEffectTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOrderSideEffectTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOrderSideEffectTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOrderSideEffectTypeParameter(val *MarginAccountNewOrderSideEffectTypeParameter) *NullableMarginAccountNewOrderSideEffectTypeParameter {
	return &NullableMarginAccountNewOrderSideEffectTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOrderSideEffectTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOrderSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
