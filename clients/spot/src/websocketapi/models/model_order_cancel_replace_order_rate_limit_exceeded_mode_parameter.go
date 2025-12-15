/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelReplaceOrderRateLimitExceededModeParameter the model 'OrderCancelReplaceOrderRateLimitExceededModeParameter'
type OrderCancelReplaceOrderRateLimitExceededModeParameter string

// List of orderCancelReplace_orderRateLimitExceededMode_parameter
const (
	OrderCancelReplaceOrderRateLimitExceededModeParameterDoNothing  OrderCancelReplaceOrderRateLimitExceededModeParameter = "DO_NOTHING"
	OrderCancelReplaceOrderRateLimitExceededModeParameterCancelOnly OrderCancelReplaceOrderRateLimitExceededModeParameter = "CANCEL_ONLY"
)

// All allowed values of OrderCancelReplaceOrderRateLimitExceededModeParameter enum
var AllowedOrderCancelReplaceOrderRateLimitExceededModeParameterEnumValues = []OrderCancelReplaceOrderRateLimitExceededModeParameter{
	"DO_NOTHING",
	"CANCEL_ONLY",
}

func (v *OrderCancelReplaceOrderRateLimitExceededModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceOrderRateLimitExceededModeParameter(value)
	for _, existing := range AllowedOrderCancelReplaceOrderRateLimitExceededModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceOrderRateLimitExceededModeParameter", value)
}

// NewOrderCancelReplaceOrderRateLimitExceededModeParameterFromValue returns a pointer to a valid OrderCancelReplaceOrderRateLimitExceededModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceOrderRateLimitExceededModeParameterFromValue(v string) (*OrderCancelReplaceOrderRateLimitExceededModeParameter, error) {
	ev := OrderCancelReplaceOrderRateLimitExceededModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceOrderRateLimitExceededModeParameter: valid values are %v", v, AllowedOrderCancelReplaceOrderRateLimitExceededModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceOrderRateLimitExceededModeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceOrderRateLimitExceededModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_orderRateLimitExceededMode_parameter value
func (v OrderCancelReplaceOrderRateLimitExceededModeParameter) Ptr() *OrderCancelReplaceOrderRateLimitExceededModeParameter {
	return &v
}

type NullableOrderCancelReplaceOrderRateLimitExceededModeParameter struct {
	value *OrderCancelReplaceOrderRateLimitExceededModeParameter
	isSet bool
}

func (v NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) Get() *OrderCancelReplaceOrderRateLimitExceededModeParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) Set(val *OrderCancelReplaceOrderRateLimitExceededModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceOrderRateLimitExceededModeParameter(val *OrderCancelReplaceOrderRateLimitExceededModeParameter) *NullableOrderCancelReplaceOrderRateLimitExceededModeParameter {
	return &NullableOrderCancelReplaceOrderRateLimitExceededModeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceOrderRateLimitExceededModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
