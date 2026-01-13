/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListPlaceOcoAboveTypeParameter the model 'OrderListPlaceOcoAboveTypeParameter'
type OrderListPlaceOcoAboveTypeParameter string

// List of orderListPlaceOco_aboveType_parameter
const (
	OrderListPlaceOcoAboveTypeParameterStopLossLimit   OrderListPlaceOcoAboveTypeParameter = "STOP_LOSS_LIMIT"
	OrderListPlaceOcoAboveTypeParameterStopLoss        OrderListPlaceOcoAboveTypeParameter = "STOP_LOSS"
	OrderListPlaceOcoAboveTypeParameterLimitMaker      OrderListPlaceOcoAboveTypeParameter = "LIMIT_MAKER"
	OrderListPlaceOcoAboveTypeParameterTakeProfit      OrderListPlaceOcoAboveTypeParameter = "TAKE_PROFIT"
	OrderListPlaceOcoAboveTypeParameterTakeProfitLimit OrderListPlaceOcoAboveTypeParameter = "TAKE_PROFIT_LIMIT"
)

// All allowed values of OrderListPlaceOcoAboveTypeParameter enum
var AllowedOrderListPlaceOcoAboveTypeParameterEnumValues = []OrderListPlaceOcoAboveTypeParameter{
	"STOP_LOSS_LIMIT",
	"STOP_LOSS",
	"LIMIT_MAKER",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
}

func (v *OrderListPlaceOcoAboveTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOcoAboveTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOcoAboveTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOcoAboveTypeParameter", value)
}

// NewOrderListPlaceOcoAboveTypeParameterFromValue returns a pointer to a valid OrderListPlaceOcoAboveTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOcoAboveTypeParameterFromValue(v string) (*OrderListPlaceOcoAboveTypeParameter, error) {
	ev := OrderListPlaceOcoAboveTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOcoAboveTypeParameter: valid values are %v", v, AllowedOrderListPlaceOcoAboveTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOcoAboveTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOcoAboveTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOco_aboveType_parameter value
func (v OrderListPlaceOcoAboveTypeParameter) Ptr() *OrderListPlaceOcoAboveTypeParameter {
	return &v
}

type NullableOrderListPlaceOcoAboveTypeParameter struct {
	value *OrderListPlaceOcoAboveTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOcoAboveTypeParameter) Get() *OrderListPlaceOcoAboveTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOcoAboveTypeParameter) Set(val *OrderListPlaceOcoAboveTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoAboveTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoAboveTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoAboveTypeParameter(val *OrderListPlaceOcoAboveTypeParameter) *NullableOrderListPlaceOcoAboveTypeParameter {
	return &NullableOrderListPlaceOcoAboveTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoAboveTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoAboveTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
