/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmOrderSelfTradePreventionModeParameter the model 'NewUmOrderSelfTradePreventionModeParameter'
type NewUmOrderSelfTradePreventionModeParameter string

// List of newUmOrder_selfTradePreventionMode_parameter
const (
	NewUmOrderSelfTradePreventionModeParameterExpireTaker NewUmOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	NewUmOrderSelfTradePreventionModeParameterExpireBoth  NewUmOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	NewUmOrderSelfTradePreventionModeParameterExpireMaker NewUmOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
)

// All allowed values of NewUmOrderSelfTradePreventionModeParameter enum
var AllowedNewUmOrderSelfTradePreventionModeParameterEnumValues = []NewUmOrderSelfTradePreventionModeParameter{
	"EXPIRE_TAKER",
	"EXPIRE_BOTH",
	"EXPIRE_MAKER",
}

func (v *NewUmOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedNewUmOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmOrderSelfTradePreventionModeParameter", value)
}

// NewNewUmOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid NewUmOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmOrderSelfTradePreventionModeParameterFromValue(v string) (*NewUmOrderSelfTradePreventionModeParameter, error) {
	ev := NewUmOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedNewUmOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedNewUmOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmOrder_selfTradePreventionMode_parameter value
func (v NewUmOrderSelfTradePreventionModeParameter) Ptr() *NewUmOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableNewUmOrderSelfTradePreventionModeParameter struct {
	value *NewUmOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableNewUmOrderSelfTradePreventionModeParameter) Get() *NewUmOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableNewUmOrderSelfTradePreventionModeParameter) Set(val *NewUmOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmOrderSelfTradePreventionModeParameter(val *NewUmOrderSelfTradePreventionModeParameter) *NullableNewUmOrderSelfTradePreventionModeParameter {
	return &NullableNewUmOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableNewUmOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
