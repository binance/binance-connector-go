/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOcoSideEffectTypeParameter the model 'MarginAccountNewOcoSideEffectTypeParameter'
type MarginAccountNewOcoSideEffectTypeParameter string

// List of marginAccountNewOco_sideEffectType_parameter
const (
	MarginAccountNewOcoSideEffectTypeParameterNoSideEffect MarginAccountNewOcoSideEffectTypeParameter = "NO_SIDE_EFFECT"
	MarginAccountNewOcoSideEffectTypeParameterMarginBuy    MarginAccountNewOcoSideEffectTypeParameter = "MARGIN_BUY"
	MarginAccountNewOcoSideEffectTypeParameterAutoRepay    MarginAccountNewOcoSideEffectTypeParameter = "AUTO_REPAY"
)

// All allowed values of MarginAccountNewOcoSideEffectTypeParameter enum
var AllowedMarginAccountNewOcoSideEffectTypeParameterEnumValues = []MarginAccountNewOcoSideEffectTypeParameter{
	"NO_SIDE_EFFECT",
	"MARGIN_BUY",
	"AUTO_REPAY",
}

func (v *MarginAccountNewOcoSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOcoSideEffectTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOcoSideEffectTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOcoSideEffectTypeParameter", value)
}

// NewMarginAccountNewOcoSideEffectTypeParameterFromValue returns a pointer to a valid MarginAccountNewOcoSideEffectTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOcoSideEffectTypeParameterFromValue(v string) (*MarginAccountNewOcoSideEffectTypeParameter, error) {
	ev := MarginAccountNewOcoSideEffectTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOcoSideEffectTypeParameter: valid values are %v", v, AllowedMarginAccountNewOcoSideEffectTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOcoSideEffectTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOcoSideEffectTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOco_sideEffectType_parameter value
func (v MarginAccountNewOcoSideEffectTypeParameter) Ptr() *MarginAccountNewOcoSideEffectTypeParameter {
	return &v
}

type NullableMarginAccountNewOcoSideEffectTypeParameter struct {
	value *MarginAccountNewOcoSideEffectTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOcoSideEffectTypeParameter) Get() *MarginAccountNewOcoSideEffectTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOcoSideEffectTypeParameter) Set(val *MarginAccountNewOcoSideEffectTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOcoSideEffectTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOcoSideEffectTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOcoSideEffectTypeParameter(val *MarginAccountNewOcoSideEffectTypeParameter) *NullableMarginAccountNewOcoSideEffectTypeParameter {
	return &NullableMarginAccountNewOcoSideEffectTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOcoSideEffectTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOcoSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
