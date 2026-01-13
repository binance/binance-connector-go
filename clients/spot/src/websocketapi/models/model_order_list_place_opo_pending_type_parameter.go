/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListPlaceOpoPendingTypeParameter the model 'OrderListPlaceOpoPendingTypeParameter'
type OrderListPlaceOpoPendingTypeParameter string

// List of orderListPlaceOpo_pendingType_parameter
const (
	OrderListPlaceOpoPendingTypeParameterLimit           OrderListPlaceOpoPendingTypeParameter = "LIMIT"
	OrderListPlaceOpoPendingTypeParameterMarket          OrderListPlaceOpoPendingTypeParameter = "MARKET"
	OrderListPlaceOpoPendingTypeParameterStopLoss        OrderListPlaceOpoPendingTypeParameter = "STOP_LOSS"
	OrderListPlaceOpoPendingTypeParameterStopLossLimit   OrderListPlaceOpoPendingTypeParameter = "STOP_LOSS_LIMIT"
	OrderListPlaceOpoPendingTypeParameterTakeProfit      OrderListPlaceOpoPendingTypeParameter = "TAKE_PROFIT"
	OrderListPlaceOpoPendingTypeParameterTakeProfitLimit OrderListPlaceOpoPendingTypeParameter = "TAKE_PROFIT_LIMIT"
	OrderListPlaceOpoPendingTypeParameterLimitMaker      OrderListPlaceOpoPendingTypeParameter = "LIMIT_MAKER"
)

// All allowed values of OrderListPlaceOpoPendingTypeParameter enum
var AllowedOrderListPlaceOpoPendingTypeParameterEnumValues = []OrderListPlaceOpoPendingTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
	"LIMIT_MAKER",
}

func (v *OrderListPlaceOpoPendingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOpoPendingTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOpoPendingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOpoPendingTypeParameter", value)
}

// NewOrderListPlaceOpoPendingTypeParameterFromValue returns a pointer to a valid OrderListPlaceOpoPendingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOpoPendingTypeParameterFromValue(v string) (*OrderListPlaceOpoPendingTypeParameter, error) {
	ev := OrderListPlaceOpoPendingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOpoPendingTypeParameter: valid values are %v", v, AllowedOrderListPlaceOpoPendingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOpoPendingTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOpoPendingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOpo_pendingType_parameter value
func (v OrderListPlaceOpoPendingTypeParameter) Ptr() *OrderListPlaceOpoPendingTypeParameter {
	return &v
}

type NullableOrderListPlaceOpoPendingTypeParameter struct {
	value *OrderListPlaceOpoPendingTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOpoPendingTypeParameter) Get() *OrderListPlaceOpoPendingTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOpoPendingTypeParameter) Set(val *OrderListPlaceOpoPendingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOpoPendingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOpoPendingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOpoPendingTypeParameter(val *OrderListPlaceOpoPendingTypeParameter) *NullableOrderListPlaceOpoPendingTypeParameter {
	return &NullableOrderListPlaceOpoPendingTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOpoPendingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOpoPendingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
