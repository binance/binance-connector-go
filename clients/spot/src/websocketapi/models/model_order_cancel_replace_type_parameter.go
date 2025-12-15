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

// OrderCancelReplaceTypeParameter the model 'OrderCancelReplaceTypeParameter'
type OrderCancelReplaceTypeParameter string

// List of orderCancelReplace_type_parameter
const (
	OrderCancelReplaceTypeParameterMarket           OrderCancelReplaceTypeParameter = "MARKET"
	OrderCancelReplaceTypeParameterLimit            OrderCancelReplaceTypeParameter = "LIMIT"
	OrderCancelReplaceTypeParameterStopLoss         OrderCancelReplaceTypeParameter = "STOP_LOSS"
	OrderCancelReplaceTypeParameterStopLossLimit    OrderCancelReplaceTypeParameter = "STOP_LOSS_LIMIT"
	OrderCancelReplaceTypeParameterTakeProfit       OrderCancelReplaceTypeParameter = "TAKE_PROFIT"
	OrderCancelReplaceTypeParameterTakeProfitLimit  OrderCancelReplaceTypeParameter = "TAKE_PROFIT_LIMIT"
	OrderCancelReplaceTypeParameterLimitMaker       OrderCancelReplaceTypeParameter = "LIMIT_MAKER"
	OrderCancelReplaceTypeParameterNonRepresentable OrderCancelReplaceTypeParameter = "NON_REPRESENTABLE"
)

// All allowed values of OrderCancelReplaceTypeParameter enum
var AllowedOrderCancelReplaceTypeParameterEnumValues = []OrderCancelReplaceTypeParameter{
	"MARKET",
	"LIMIT",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
	"LIMIT_MAKER",
	"NON_REPRESENTABLE",
}

func (v *OrderCancelReplaceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceTypeParameter(value)
	for _, existing := range AllowedOrderCancelReplaceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceTypeParameter", value)
}

// NewOrderCancelReplaceTypeParameterFromValue returns a pointer to a valid OrderCancelReplaceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceTypeParameterFromValue(v string) (*OrderCancelReplaceTypeParameter, error) {
	ev := OrderCancelReplaceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceTypeParameter: valid values are %v", v, AllowedOrderCancelReplaceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_type_parameter value
func (v OrderCancelReplaceTypeParameter) Ptr() *OrderCancelReplaceTypeParameter {
	return &v
}

type NullableOrderCancelReplaceTypeParameter struct {
	value *OrderCancelReplaceTypeParameter
	isSet bool
}

func (v NullableOrderCancelReplaceTypeParameter) Get() *OrderCancelReplaceTypeParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceTypeParameter) Set(val *OrderCancelReplaceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceTypeParameter(val *OrderCancelReplaceTypeParameter) *NullableOrderCancelReplaceTypeParameter {
	return &NullableOrderCancelReplaceTypeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
