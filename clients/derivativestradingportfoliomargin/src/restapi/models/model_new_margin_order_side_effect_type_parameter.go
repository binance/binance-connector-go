/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewMarginOrderSideEffectTypeParameter the model 'NewMarginOrderSideEffectTypeParameter'
type NewMarginOrderSideEffectTypeParameter string

// List of newMarginOrder_sideEffectType_parameter
const (
	NewMarginOrderSideEffectTypeParameterNoSideEffect NewMarginOrderSideEffectTypeParameter = "NO_SIDE_EFFECT"
	NewMarginOrderSideEffectTypeParameterMarginBuy    NewMarginOrderSideEffectTypeParameter = "MARGIN_BUY"
	NewMarginOrderSideEffectTypeParameterAutoRepay    NewMarginOrderSideEffectTypeParameter = "AUTO_REPAY"
)

// All allowed values of NewMarginOrderSideEffectTypeParameter enum
var AllowedNewMarginOrderSideEffectTypeParameterEnumValues = []NewMarginOrderSideEffectTypeParameter{
	"NO_SIDE_EFFECT",
	"MARGIN_BUY",
	"AUTO_REPAY",
}

func (v *NewMarginOrderSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewMarginOrderSideEffectTypeParameter(value)
	for _, existing := range AllowedNewMarginOrderSideEffectTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewMarginOrderSideEffectTypeParameter", value)
}

// NewNewMarginOrderSideEffectTypeParameterFromValue returns a pointer to a valid NewMarginOrderSideEffectTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewMarginOrderSideEffectTypeParameterFromValue(v string) (*NewMarginOrderSideEffectTypeParameter, error) {
	ev := NewMarginOrderSideEffectTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewMarginOrderSideEffectTypeParameter: valid values are %v", v, AllowedNewMarginOrderSideEffectTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewMarginOrderSideEffectTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewMarginOrderSideEffectTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newMarginOrder_sideEffectType_parameter value
func (v NewMarginOrderSideEffectTypeParameter) Ptr() *NewMarginOrderSideEffectTypeParameter {
	return &v
}

type NullableNewMarginOrderSideEffectTypeParameter struct {
	value *NewMarginOrderSideEffectTypeParameter
	isSet bool
}

func (v NullableNewMarginOrderSideEffectTypeParameter) Get() *NewMarginOrderSideEffectTypeParameter {
	return v.value
}

func (v *NullableNewMarginOrderSideEffectTypeParameter) Set(val *NewMarginOrderSideEffectTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderSideEffectTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderSideEffectTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderSideEffectTypeParameter(val *NewMarginOrderSideEffectTypeParameter) *NullableNewMarginOrderSideEffectTypeParameter {
	return &NullableNewMarginOrderSideEffectTypeParameter{value: val, isSet: true}
}

func (v NullableNewMarginOrderSideEffectTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderSideEffectTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
