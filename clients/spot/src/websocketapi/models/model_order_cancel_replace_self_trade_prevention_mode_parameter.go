/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelReplaceSelfTradePreventionModeParameter the model 'OrderCancelReplaceSelfTradePreventionModeParameter'
type OrderCancelReplaceSelfTradePreventionModeParameter string

// List of orderCancelReplace_selfTradePreventionMode_parameter
const (
	OrderCancelReplaceSelfTradePreventionModeParameterNone             OrderCancelReplaceSelfTradePreventionModeParameter = "NONE"
	OrderCancelReplaceSelfTradePreventionModeParameterExpireTaker      OrderCancelReplaceSelfTradePreventionModeParameter = "EXPIRE_TAKER"
	OrderCancelReplaceSelfTradePreventionModeParameterExpireMaker      OrderCancelReplaceSelfTradePreventionModeParameter = "EXPIRE_MAKER"
	OrderCancelReplaceSelfTradePreventionModeParameterExpireBoth       OrderCancelReplaceSelfTradePreventionModeParameter = "EXPIRE_BOTH"
	OrderCancelReplaceSelfTradePreventionModeParameterDecrement        OrderCancelReplaceSelfTradePreventionModeParameter = "DECREMENT"
	OrderCancelReplaceSelfTradePreventionModeParameterNonRepresentable OrderCancelReplaceSelfTradePreventionModeParameter = "NON_REPRESENTABLE"
)

// All allowed values of OrderCancelReplaceSelfTradePreventionModeParameter enum
var AllowedOrderCancelReplaceSelfTradePreventionModeParameterEnumValues = []OrderCancelReplaceSelfTradePreventionModeParameter{
	"NONE",
	"EXPIRE_TAKER",
	"EXPIRE_MAKER",
	"EXPIRE_BOTH",
	"DECREMENT",
	"NON_REPRESENTABLE",
}

func (v *OrderCancelReplaceSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceSelfTradePreventionModeParameter(value)
	for _, existing := range AllowedOrderCancelReplaceSelfTradePreventionModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceSelfTradePreventionModeParameter", value)
}

// NewOrderCancelReplaceSelfTradePreventionModeParameterFromValue returns a pointer to a valid OrderCancelReplaceSelfTradePreventionModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceSelfTradePreventionModeParameterFromValue(v string) (*OrderCancelReplaceSelfTradePreventionModeParameter, error) {
	ev := OrderCancelReplaceSelfTradePreventionModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceSelfTradePreventionModeParameter: valid values are %v", v, AllowedOrderCancelReplaceSelfTradePreventionModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceSelfTradePreventionModeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceSelfTradePreventionModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_selfTradePreventionMode_parameter value
func (v OrderCancelReplaceSelfTradePreventionModeParameter) Ptr() *OrderCancelReplaceSelfTradePreventionModeParameter {
	return &v
}

type NullableOrderCancelReplaceSelfTradePreventionModeParameter struct {
	value *OrderCancelReplaceSelfTradePreventionModeParameter
	isSet bool
}

func (v NullableOrderCancelReplaceSelfTradePreventionModeParameter) Get() *OrderCancelReplaceSelfTradePreventionModeParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceSelfTradePreventionModeParameter) Set(val *OrderCancelReplaceSelfTradePreventionModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceSelfTradePreventionModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceSelfTradePreventionModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceSelfTradePreventionModeParameter(val *OrderCancelReplaceSelfTradePreventionModeParameter) *NullableOrderCancelReplaceSelfTradePreventionModeParameter {
	return &NullableOrderCancelReplaceSelfTradePreventionModeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceSelfTradePreventionModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceSelfTradePreventionModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
