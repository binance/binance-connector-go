/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOcoAboveTypeParameter the model 'OrderListOcoAboveTypeParameter'
type OrderListOcoAboveTypeParameter string

// List of orderListOco_aboveType_parameter
const (
	OrderListOcoAboveTypeParameterStopLossLimit   OrderListOcoAboveTypeParameter = "STOP_LOSS_LIMIT"
	OrderListOcoAboveTypeParameterStopLoss        OrderListOcoAboveTypeParameter = "STOP_LOSS"
	OrderListOcoAboveTypeParameterLimitMaker      OrderListOcoAboveTypeParameter = "LIMIT_MAKER"
	OrderListOcoAboveTypeParameterTakeProfit      OrderListOcoAboveTypeParameter = "TAKE_PROFIT"
	OrderListOcoAboveTypeParameterTakeProfitLimit OrderListOcoAboveTypeParameter = "TAKE_PROFIT_LIMIT"
)

// All allowed values of OrderListOcoAboveTypeParameter enum
var AllowedOrderListOcoAboveTypeParameterEnumValues = []OrderListOcoAboveTypeParameter{
	"STOP_LOSS_LIMIT",
	"STOP_LOSS",
	"LIMIT_MAKER",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
}

func (v *OrderListOcoAboveTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOcoAboveTypeParameter(value)
	for _, existing := range AllowedOrderListOcoAboveTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOcoAboveTypeParameter", value)
}

// NewOrderListOcoAboveTypeParameterFromValue returns a pointer to a valid OrderListOcoAboveTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOcoAboveTypeParameterFromValue(v string) (*OrderListOcoAboveTypeParameter, error) {
	ev := OrderListOcoAboveTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOcoAboveTypeParameter: valid values are %v", v, AllowedOrderListOcoAboveTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOcoAboveTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOcoAboveTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOco_aboveType_parameter value
func (v OrderListOcoAboveTypeParameter) Ptr() *OrderListOcoAboveTypeParameter {
	return &v
}

type NullableOrderListOcoAboveTypeParameter struct {
	value *OrderListOcoAboveTypeParameter
	isSet bool
}

func (v NullableOrderListOcoAboveTypeParameter) Get() *OrderListOcoAboveTypeParameter {
	return v.value
}

func (v *NullableOrderListOcoAboveTypeParameter) Set(val *OrderListOcoAboveTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOcoAboveTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOcoAboveTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOcoAboveTypeParameter(val *OrderListOcoAboveTypeParameter) *NullableOrderListOcoAboveTypeParameter {
	return &NullableOrderListOcoAboveTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOcoAboveTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOcoAboveTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
