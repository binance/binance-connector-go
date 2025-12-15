/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderSelfTradePreventionModeParameter the model 'NewAlgoOrderSelfTradePreventionModeParameter'
type NewAlgoOrderSelfTradePreventionModeParameter string

// List of newAlgoOrder_selfTradePreventionMode_parameter
const (
	NewAlgoOrderSelfTradePreventionModeParameterExpireTaker NewAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	NewAlgoOrderSelfTradePreventionModeParameterExpireBoth  NewAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	NewAlgoOrderSelfTradePreventionModeParameterExpireMaker NewAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
)

// All allowed values of NewAlgoOrderSelfTradePreventionModeParameter enum
var AllowedNewAlgoOrderSelfTradePreventionModeParameterEnumValues = []NewAlgoOrderSelfTradePreventionModeParameter{
	"EXPIRE_TAKER",
	"EXPIRE_BOTH",
	"EXPIRE_MAKER",
}

func (v *NewAlgoOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedNewAlgoOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderSelfTradePreventionModeParameter", value)
}

// NewNewAlgoOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid NewAlgoOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderSelfTradePreventionModeParameterFromValue(v string) (*NewAlgoOrderSelfTradePreventionModeParameter, error) {
	ev := NewAlgoOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedNewAlgoOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_selfTradePreventionMode_parameter value
func (v NewAlgoOrderSelfTradePreventionModeParameter) Ptr() *NewAlgoOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableNewAlgoOrderSelfTradePreventionModeParameter struct {
	value *NewAlgoOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableNewAlgoOrderSelfTradePreventionModeParameter) Get() *NewAlgoOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableNewAlgoOrderSelfTradePreventionModeParameter) Set(val *NewAlgoOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderSelfTradePreventionModeParameter(val *NewAlgoOrderSelfTradePreventionModeParameter) *NullableNewAlgoOrderSelfTradePreventionModeParameter {
	return &NullableNewAlgoOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
