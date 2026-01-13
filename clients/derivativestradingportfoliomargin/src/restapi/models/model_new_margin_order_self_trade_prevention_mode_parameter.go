/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewMarginOrderSelfTradePreventionModeParameter the model 'NewMarginOrderSelfTradePreventionModeParameter'
type NewMarginOrderSelfTradePreventionModeParameter string

// List of newMarginOrder_selfTradePreventionMode_parameter
const (
	NewMarginOrderSelfTradePreventionModeParameterNone        NewMarginOrderSelfTradePreventionModeParameter = "NONE"
	NewMarginOrderSelfTradePreventionModeParameterExpireTaker NewMarginOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	NewMarginOrderSelfTradePreventionModeParameterExpireBoth  NewMarginOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	NewMarginOrderSelfTradePreventionModeParameterExpireMaker NewMarginOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
)

// All allowed values of NewMarginOrderSelfTradePreventionModeParameter enum
var AllowedNewMarginOrderSelfTradePreventionModeParameterEnumValues = []NewMarginOrderSelfTradePreventionModeParameter{
	"NONE",
	"EXPIRE_TAKER",
	"EXPIRE_BOTH",
	"EXPIRE_MAKER",
}

func (v *NewMarginOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewMarginOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedNewMarginOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewMarginOrderSelfTradePreventionModeParameter", value)
}

// NewNewMarginOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid NewMarginOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewMarginOrderSelfTradePreventionModeParameterFromValue(v string) (*NewMarginOrderSelfTradePreventionModeParameter, error) {
	ev := NewMarginOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewMarginOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedNewMarginOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewMarginOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedNewMarginOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newMarginOrder_selfTradePreventionMode_parameter value
func (v NewMarginOrderSelfTradePreventionModeParameter) Ptr() *NewMarginOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableNewMarginOrderSelfTradePreventionModeParameter struct {
	value *NewMarginOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableNewMarginOrderSelfTradePreventionModeParameter) Get() *NewMarginOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableNewMarginOrderSelfTradePreventionModeParameter) Set(val *NewMarginOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderSelfTradePreventionModeParameter(val *NewMarginOrderSelfTradePreventionModeParameter) *NullableNewMarginOrderSelfTradePreventionModeParameter {
	return &NullableNewMarginOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableNewMarginOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
