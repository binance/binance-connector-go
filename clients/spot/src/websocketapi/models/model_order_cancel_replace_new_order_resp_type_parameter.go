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

// OrderCancelReplaceNewOrderRespTypeParameter the model 'OrderCancelReplaceNewOrderRespTypeParameter'
type OrderCancelReplaceNewOrderRespTypeParameter string

// List of orderCancelReplace_newOrderRespType_parameter
const (
	OrderCancelReplaceNewOrderRespTypeParameterAck    OrderCancelReplaceNewOrderRespTypeParameter = "ACK"
	OrderCancelReplaceNewOrderRespTypeParameterResult OrderCancelReplaceNewOrderRespTypeParameter = "RESULT"
	OrderCancelReplaceNewOrderRespTypeParameterFull   OrderCancelReplaceNewOrderRespTypeParameter = "FULL"
	OrderCancelReplaceNewOrderRespTypeParameterMarket OrderCancelReplaceNewOrderRespTypeParameter = "MARKET"
	OrderCancelReplaceNewOrderRespTypeParameterLimit  OrderCancelReplaceNewOrderRespTypeParameter = "LIMIT"
)

// All allowed values of OrderCancelReplaceNewOrderRespTypeParameter enum
var AllowedOrderCancelReplaceNewOrderRespTypeParameterEnumValues = []OrderCancelReplaceNewOrderRespTypeParameter{
	"ACK",
	"RESULT",
	"FULL",
	"MARKET",
	"LIMIT",
}

func (v *OrderCancelReplaceNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceNewOrderRespTypeParameter(value)
	for _, existing := range AllowedOrderCancelReplaceNewOrderRespTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceNewOrderRespTypeParameter", value)
}

// NewOrderCancelReplaceNewOrderRespTypeParameterFromValue returns a pointer to a valid OrderCancelReplaceNewOrderRespTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceNewOrderRespTypeParameterFromValue(v string) (*OrderCancelReplaceNewOrderRespTypeParameter, error) {
	ev := OrderCancelReplaceNewOrderRespTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceNewOrderRespTypeParameter: valid values are %v", v, AllowedOrderCancelReplaceNewOrderRespTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceNewOrderRespTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceNewOrderRespTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_newOrderRespType_parameter value
func (v OrderCancelReplaceNewOrderRespTypeParameter) Ptr() *OrderCancelReplaceNewOrderRespTypeParameter {
	return &v
}

type NullableOrderCancelReplaceNewOrderRespTypeParameter struct {
	value *OrderCancelReplaceNewOrderRespTypeParameter
	isSet bool
}

func (v NullableOrderCancelReplaceNewOrderRespTypeParameter) Get() *OrderCancelReplaceNewOrderRespTypeParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceNewOrderRespTypeParameter) Set(val *OrderCancelReplaceNewOrderRespTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceNewOrderRespTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceNewOrderRespTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceNewOrderRespTypeParameter(val *OrderCancelReplaceNewOrderRespTypeParameter) *NullableOrderCancelReplaceNewOrderRespTypeParameter {
	return &NullableOrderCancelReplaceNewOrderRespTypeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceNewOrderRespTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceNewOrderRespTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
