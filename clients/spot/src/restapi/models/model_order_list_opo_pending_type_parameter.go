/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOpoPendingTypeParameter the model 'OrderListOpoPendingTypeParameter'
type OrderListOpoPendingTypeParameter string

// List of orderListOpo_pendingType_parameter
const (
	OrderListOpoPendingTypeParameterLimit           OrderListOpoPendingTypeParameter = "LIMIT"
	OrderListOpoPendingTypeParameterMarket          OrderListOpoPendingTypeParameter = "MARKET"
	OrderListOpoPendingTypeParameterStopLoss        OrderListOpoPendingTypeParameter = "STOP_LOSS"
	OrderListOpoPendingTypeParameterStopLossLimit   OrderListOpoPendingTypeParameter = "STOP_LOSS_LIMIT"
	OrderListOpoPendingTypeParameterTakeProfit      OrderListOpoPendingTypeParameter = "TAKE_PROFIT"
	OrderListOpoPendingTypeParameterTakeProfitLimit OrderListOpoPendingTypeParameter = "TAKE_PROFIT_LIMIT"
	OrderListOpoPendingTypeParameterLimitMaker      OrderListOpoPendingTypeParameter = "LIMIT_MAKER"
)

// All allowed values of OrderListOpoPendingTypeParameter enum
var AllowedOrderListOpoPendingTypeParameterEnumValues = []OrderListOpoPendingTypeParameter{
	"LIMIT",
	"MARKET",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
	"LIMIT_MAKER",
}

func (v *OrderListOpoPendingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOpoPendingTypeParameter(value)
	for _, existing := range AllowedOrderListOpoPendingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOpoPendingTypeParameter", value)
}

// NewOrderListOpoPendingTypeParameterFromValue returns a pointer to a valid OrderListOpoPendingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOpoPendingTypeParameterFromValue(v string) (*OrderListOpoPendingTypeParameter, error) {
	ev := OrderListOpoPendingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOpoPendingTypeParameter: valid values are %v", v, AllowedOrderListOpoPendingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOpoPendingTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOpoPendingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOpo_pendingType_parameter value
func (v OrderListOpoPendingTypeParameter) Ptr() *OrderListOpoPendingTypeParameter {
	return &v
}

type NullableOrderListOpoPendingTypeParameter struct {
	value *OrderListOpoPendingTypeParameter
	isSet bool
}

func (v NullableOrderListOpoPendingTypeParameter) Get() *OrderListOpoPendingTypeParameter {
	return v.value
}

func (v *NullableOrderListOpoPendingTypeParameter) Set(val *OrderListOpoPendingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOpoPendingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOpoPendingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOpoPendingTypeParameter(val *OrderListOpoPendingTypeParameter) *NullableOrderListOpoPendingTypeParameter {
	return &NullableOrderListOpoPendingTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOpoPendingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOpoPendingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
