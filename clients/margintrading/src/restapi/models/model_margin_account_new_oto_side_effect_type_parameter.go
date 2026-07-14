/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOtoSideEffectTypeParameter the model 'MarginAccountNewOtoSideEffectTypeParameter'
type MarginAccountNewOtoSideEffectTypeParameter string

// List of marginAccountNewOto_sideEffectType_parameter
const (
	MarginAccountNewOtoSideEffectTypeParameterNoSideEffect MarginAccountNewOtoSideEffectTypeParameter = "NO_SIDE_EFFECT"
	MarginAccountNewOtoSideEffectTypeParameterMarginBuy    MarginAccountNewOtoSideEffectTypeParameter = "MARGIN_BUY"
)

// All allowed values of MarginAccountNewOtoSideEffectTypeParameter enum
var AllowedMarginAccountNewOtoSideEffectTypeParameterEnumValues = []MarginAccountNewOtoSideEffectTypeParameter{
	"NO_SIDE_EFFECT",
	"MARGIN_BUY",
}

func (v *MarginAccountNewOtoSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOtoSideEffectTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOtoSideEffectTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOtoSideEffectTypeParameter", value)
}

// NewMarginAccountNewOtoSideEffectTypeParameterFromValue returns a pointer to a valid MarginAccountNewOtoSideEffectTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOtoSideEffectTypeParameterFromValue(v string) (*MarginAccountNewOtoSideEffectTypeParameter, error) {
	ev := MarginAccountNewOtoSideEffectTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOtoSideEffectTypeParameter: valid values are %v", v, AllowedMarginAccountNewOtoSideEffectTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOtoSideEffectTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOtoSideEffectTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOto_sideEffectType_parameter value
func (v MarginAccountNewOtoSideEffectTypeParameter) Ptr() *MarginAccountNewOtoSideEffectTypeParameter {
	return &v
}

type NullableMarginAccountNewOtoSideEffectTypeParameter struct {
	value *MarginAccountNewOtoSideEffectTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOtoSideEffectTypeParameter) Get() *MarginAccountNewOtoSideEffectTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOtoSideEffectTypeParameter) Set(val *MarginAccountNewOtoSideEffectTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtoSideEffectTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtoSideEffectTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtoSideEffectTypeParameter(val *MarginAccountNewOtoSideEffectTypeParameter) *NullableMarginAccountNewOtoSideEffectTypeParameter {
	return &NullableMarginAccountNewOtoSideEffectTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtoSideEffectTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtoSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
