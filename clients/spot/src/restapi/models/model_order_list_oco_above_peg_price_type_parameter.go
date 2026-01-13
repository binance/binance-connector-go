/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOcoAbovePegPriceTypeParameter the model 'OrderListOcoAbovePegPriceTypeParameter'
type OrderListOcoAbovePegPriceTypeParameter string

// List of orderListOco_abovePegPriceType_parameter
const (
	OrderListOcoAbovePegPriceTypeParameterPrimaryPeg OrderListOcoAbovePegPriceTypeParameter = "PRIMARY_PEG"
	OrderListOcoAbovePegPriceTypeParameterMarketPeg  OrderListOcoAbovePegPriceTypeParameter = "MARKET_PEG"
)

// All allowed values of OrderListOcoAbovePegPriceTypeParameter enum
var AllowedOrderListOcoAbovePegPriceTypeParameterEnumValues = []OrderListOcoAbovePegPriceTypeParameter{
	"PRIMARY_PEG",
	"MARKET_PEG",
}

func (v *OrderListOcoAbovePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOcoAbovePegPriceTypeParameter(value)
	for _, existing := range AllowedOrderListOcoAbovePegPriceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOcoAbovePegPriceTypeParameter", value)
}

// NewOrderListOcoAbovePegPriceTypeParameterFromValue returns a pointer to a valid OrderListOcoAbovePegPriceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOcoAbovePegPriceTypeParameterFromValue(v string) (*OrderListOcoAbovePegPriceTypeParameter, error) {
	ev := OrderListOcoAbovePegPriceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOcoAbovePegPriceTypeParameter: valid values are %v", v, AllowedOrderListOcoAbovePegPriceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOcoAbovePegPriceTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOcoAbovePegPriceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOco_abovePegPriceType_parameter value
func (v OrderListOcoAbovePegPriceTypeParameter) Ptr() *OrderListOcoAbovePegPriceTypeParameter {
	return &v
}

type NullableOrderListOcoAbovePegPriceTypeParameter struct {
	value *OrderListOcoAbovePegPriceTypeParameter
	isSet bool
}

func (v NullableOrderListOcoAbovePegPriceTypeParameter) Get() *OrderListOcoAbovePegPriceTypeParameter {
	return v.value
}

func (v *NullableOrderListOcoAbovePegPriceTypeParameter) Set(val *OrderListOcoAbovePegPriceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOcoAbovePegPriceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOcoAbovePegPriceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOcoAbovePegPriceTypeParameter(val *OrderListOcoAbovePegPriceTypeParameter) *NullableOrderListOcoAbovePegPriceTypeParameter {
	return &NullableOrderListOcoAbovePegPriceTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOcoAbovePegPriceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOcoAbovePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
