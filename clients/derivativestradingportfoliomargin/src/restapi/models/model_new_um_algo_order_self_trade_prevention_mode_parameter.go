/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderSelfTradePreventionModeParameter the model 'NewUmAlgoOrderSelfTradePreventionModeParameter'
type NewUmAlgoOrderSelfTradePreventionModeParameter string

// List of newUmAlgoOrder_selfTradePreventionMode_parameter
const (
	NewUmAlgoOrderSelfTradePreventionModeParameterNone        NewUmAlgoOrderSelfTradePreventionModeParameter = "NONE"
	NewUmAlgoOrderSelfTradePreventionModeParameterExpireTaker NewUmAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	NewUmAlgoOrderSelfTradePreventionModeParameterExpireMaker NewUmAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
	NewUmAlgoOrderSelfTradePreventionModeParameterExpireBoth  NewUmAlgoOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
)

// All allowed values of NewUmAlgoOrderSelfTradePreventionModeParameter enum
var AllowedNewUmAlgoOrderSelfTradePreventionModeParameterEnumValues = []NewUmAlgoOrderSelfTradePreventionModeParameter{
	"NONE",
	"EXPIRE_TAKER",
	"EXPIRE_MAKER",
	"EXPIRE_BOTH",
}

func (v *NewUmAlgoOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderSelfTradePreventionModeParameter", value)
}

// NewNewUmAlgoOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid NewUmAlgoOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderSelfTradePreventionModeParameterFromValue(v string) (*NewUmAlgoOrderSelfTradePreventionModeParameter, error) {
	ev := NewUmAlgoOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedNewUmAlgoOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_selfTradePreventionMode_parameter value
func (v NewUmAlgoOrderSelfTradePreventionModeParameter) Ptr() *NewUmAlgoOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableNewUmAlgoOrderSelfTradePreventionModeParameter struct {
	value *NewUmAlgoOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderSelfTradePreventionModeParameter) Get() *NewUmAlgoOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderSelfTradePreventionModeParameter) Set(val *NewUmAlgoOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderSelfTradePreventionModeParameter(val *NewUmAlgoOrderSelfTradePreventionModeParameter) *NullableNewUmAlgoOrderSelfTradePreventionModeParameter {
	return &NullableNewUmAlgoOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
