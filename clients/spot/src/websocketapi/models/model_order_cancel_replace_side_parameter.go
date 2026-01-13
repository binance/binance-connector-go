/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelReplaceSideParameter the model 'OrderCancelReplaceSideParameter'
type OrderCancelReplaceSideParameter string

// List of orderCancelReplace_side_parameter
const (
	OrderCancelReplaceSideParameterBuy  OrderCancelReplaceSideParameter = "BUY"
	OrderCancelReplaceSideParameterSell OrderCancelReplaceSideParameter = "SELL"
)

// All allowed values of OrderCancelReplaceSideParameter enum
var AllowedOrderCancelReplaceSideParameterEnumValues = []OrderCancelReplaceSideParameter{
	"BUY",
	"SELL",
}

func (v *OrderCancelReplaceSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceSideParameter(value)
	for _, existing := range AllowedOrderCancelReplaceSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceSideParameter", value)
}

// NewOrderCancelReplaceSideParameterFromValue returns a pointer to a valid OrderCancelReplaceSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceSideParameterFromValue(v string) (*OrderCancelReplaceSideParameter, error) {
	ev := OrderCancelReplaceSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceSideParameter: valid values are %v", v, AllowedOrderCancelReplaceSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceSideParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_side_parameter value
func (v OrderCancelReplaceSideParameter) Ptr() *OrderCancelReplaceSideParameter {
	return &v
}

type NullableOrderCancelReplaceSideParameter struct {
	value *OrderCancelReplaceSideParameter
	isSet bool
}

func (v NullableOrderCancelReplaceSideParameter) Get() *OrderCancelReplaceSideParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceSideParameter) Set(val *OrderCancelReplaceSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceSideParameter(val *OrderCancelReplaceSideParameter) *NullableOrderCancelReplaceSideParameter {
	return &NullableOrderCancelReplaceSideParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
