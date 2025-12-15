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

// OrderListPlaceOcoBelowTypeParameter the model 'OrderListPlaceOcoBelowTypeParameter'
type OrderListPlaceOcoBelowTypeParameter string

// List of orderListPlaceOco_belowType_parameter
const (
	OrderListPlaceOcoBelowTypeParameterStopLoss        OrderListPlaceOcoBelowTypeParameter = "STOP_LOSS"
	OrderListPlaceOcoBelowTypeParameterStopLossLimit   OrderListPlaceOcoBelowTypeParameter = "STOP_LOSS_LIMIT"
	OrderListPlaceOcoBelowTypeParameterTakeProfit      OrderListPlaceOcoBelowTypeParameter = "TAKE_PROFIT"
	OrderListPlaceOcoBelowTypeParameterTakeProfitLimit OrderListPlaceOcoBelowTypeParameter = "TAKE_PROFIT_LIMIT"
)

// All allowed values of OrderListPlaceOcoBelowTypeParameter enum
var AllowedOrderListPlaceOcoBelowTypeParameterEnumValues = []OrderListPlaceOcoBelowTypeParameter{
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
}

func (v *OrderListPlaceOcoBelowTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOcoBelowTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOcoBelowTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOcoBelowTypeParameter", value)
}

// NewOrderListPlaceOcoBelowTypeParameterFromValue returns a pointer to a valid OrderListPlaceOcoBelowTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOcoBelowTypeParameterFromValue(v string) (*OrderListPlaceOcoBelowTypeParameter, error) {
	ev := OrderListPlaceOcoBelowTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOcoBelowTypeParameter: valid values are %v", v, AllowedOrderListPlaceOcoBelowTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOcoBelowTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOcoBelowTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOco_belowType_parameter value
func (v OrderListPlaceOcoBelowTypeParameter) Ptr() *OrderListPlaceOcoBelowTypeParameter {
	return &v
}

type NullableOrderListPlaceOcoBelowTypeParameter struct {
	value *OrderListPlaceOcoBelowTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOcoBelowTypeParameter) Get() *OrderListPlaceOcoBelowTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOcoBelowTypeParameter) Set(val *OrderListPlaceOcoBelowTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoBelowTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoBelowTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoBelowTypeParameter(val *OrderListPlaceOcoBelowTypeParameter) *NullableOrderListPlaceOcoBelowTypeParameter {
	return &NullableOrderListPlaceOcoBelowTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoBelowTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoBelowTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
