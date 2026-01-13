/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderSelfTradePreventionModeParameter the model 'NewOrderSelfTradePreventionModeParameter'
type NewOrderSelfTradePreventionModeParameter string

// List of newOrder_selfTradePreventionMode_parameter
const (
	NewOrderSelfTradePreventionModeParameterNone             NewOrderSelfTradePreventionModeParameter = "NONE"
	NewOrderSelfTradePreventionModeParameterExpireTaker      NewOrderSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	NewOrderSelfTradePreventionModeParameterExpireMaker      NewOrderSelfTradePreventionModeParameter = "EXPIRE_MAKER"
	NewOrderSelfTradePreventionModeParameterExpireBoth       NewOrderSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	NewOrderSelfTradePreventionModeParameterDecrement        NewOrderSelfTradePreventionModeParameter = "DECREMENT"
	NewOrderSelfTradePreventionModeParameterNonRepresentable NewOrderSelfTradePreventionModeParameter = "NON_REPRESENTABLE"
)

// All allowed values of NewOrderSelfTradePreventionModeParameter enum
var AllowedNewOrderSelfTradePreventionModeParameterEnumValues = []NewOrderSelfTradePreventionModeParameter{
	"NONE",
	"EXPIRE_TAKER",
	"EXPIRE_MAKER",
	"EXPIRE_BOTH",
	"DECREMENT",
	"NON_REPRESENTABLE",
}

func (v *NewOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedNewOrderSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderSelfTradePreventionModeParameter", value)
}

// NewNewOrderSelfTradePreventionModeParameterFromValue returns a pointer to a valid NewOrderSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderSelfTradePreventionModeParameterFromValue(v string) (*NewOrderSelfTradePreventionModeParameter, error) {
	ev := NewOrderSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderSelfTradePreventionModeParameter: valid values are %v", v, AllowedNewOrderSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_selfTradePreventionMode_parameter value
func (v NewOrderSelfTradePreventionModeParameter) Ptr() *NewOrderSelfTradePreventionModeParameter {
	return &v
}

type NullableNewOrderSelfTradePreventionModeParameter struct {
	value *NewOrderSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableNewOrderSelfTradePreventionModeParameter) Get() *NewOrderSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableNewOrderSelfTradePreventionModeParameter) Set(val *NewOrderSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderSelfTradePreventionModeParameter(val *NewOrderSelfTradePreventionModeParameter) *NullableNewOrderSelfTradePreventionModeParameter {
	return &NullableNewOrderSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableNewOrderSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
