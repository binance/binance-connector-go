/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOcoBelowTypeParameter the model 'OrderListOcoBelowTypeParameter'
type OrderListOcoBelowTypeParameter string

// List of orderListOco_belowType_parameter
const (
	OrderListOcoBelowTypeParameterStopLoss        OrderListOcoBelowTypeParameter = "STOP_LOSS"
	OrderListOcoBelowTypeParameterStopLossLimit   OrderListOcoBelowTypeParameter = "STOP_LOSS_LIMIT"
	OrderListOcoBelowTypeParameterTakeProfit      OrderListOcoBelowTypeParameter = "TAKE_PROFIT"
	OrderListOcoBelowTypeParameterTakeProfitLimit OrderListOcoBelowTypeParameter = "TAKE_PROFIT_LIMIT"
)

// All allowed values of OrderListOcoBelowTypeParameter enum
var AllowedOrderListOcoBelowTypeParameterEnumValues = []OrderListOcoBelowTypeParameter{
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
	"TAKE_PROFIT",
	"TAKE_PROFIT_LIMIT",
}

func (v *OrderListOcoBelowTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOcoBelowTypeParameter(value)
	for _, existing := range AllowedOrderListOcoBelowTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOcoBelowTypeParameter", value)
}

// NewOrderListOcoBelowTypeParameterFromValue returns a pointer to a valid OrderListOcoBelowTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOcoBelowTypeParameterFromValue(v string) (*OrderListOcoBelowTypeParameter, error) {
	ev := OrderListOcoBelowTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOcoBelowTypeParameter: valid values are %v", v, AllowedOrderListOcoBelowTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOcoBelowTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOcoBelowTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOco_belowType_parameter value
func (v OrderListOcoBelowTypeParameter) Ptr() *OrderListOcoBelowTypeParameter {
	return &v
}

type NullableOrderListOcoBelowTypeParameter struct {
	value *OrderListOcoBelowTypeParameter
	isSet bool
}

func (v NullableOrderListOcoBelowTypeParameter) Get() *OrderListOcoBelowTypeParameter {
	return v.value
}

func (v *NullableOrderListOcoBelowTypeParameter) Set(val *OrderListOcoBelowTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOcoBelowTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOcoBelowTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOcoBelowTypeParameter(val *OrderListOcoBelowTypeParameter) *NullableOrderListOcoBelowTypeParameter {
	return &NullableOrderListOcoBelowTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOcoBelowTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOcoBelowTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
