/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListPlaceOcoAbovePegPriceTypeParameter the model 'OrderListPlaceOcoAbovePegPriceTypeParameter'
type OrderListPlaceOcoAbovePegPriceTypeParameter string

// List of orderListPlaceOco_abovePegPriceType_parameter
const (
	OrderListPlaceOcoAbovePegPriceTypeParameterPrimaryPeg OrderListPlaceOcoAbovePegPriceTypeParameter = "PRIMARY_PEG"
	OrderListPlaceOcoAbovePegPriceTypeParameterMarketPeg  OrderListPlaceOcoAbovePegPriceTypeParameter = "MARKET_PEG"
)

// All allowed values of OrderListPlaceOcoAbovePegPriceTypeParameter enum
var AllowedOrderListPlaceOcoAbovePegPriceTypeParameterEnumValues = []OrderListPlaceOcoAbovePegPriceTypeParameter{
	"PRIMARY_PEG",
	"MARKET_PEG",
}

func (v *OrderListPlaceOcoAbovePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOcoAbovePegPriceTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOcoAbovePegPriceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOcoAbovePegPriceTypeParameter", value)
}

// NewOrderListPlaceOcoAbovePegPriceTypeParameterFromValue returns a pointer to a valid OrderListPlaceOcoAbovePegPriceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOcoAbovePegPriceTypeParameterFromValue(v string) (*OrderListPlaceOcoAbovePegPriceTypeParameter, error) {
	ev := OrderListPlaceOcoAbovePegPriceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOcoAbovePegPriceTypeParameter: valid values are %v", v, AllowedOrderListPlaceOcoAbovePegPriceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOcoAbovePegPriceTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOcoAbovePegPriceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOco_abovePegPriceType_parameter value
func (v OrderListPlaceOcoAbovePegPriceTypeParameter) Ptr() *OrderListPlaceOcoAbovePegPriceTypeParameter {
	return &v
}

type NullableOrderListPlaceOcoAbovePegPriceTypeParameter struct {
	value *OrderListPlaceOcoAbovePegPriceTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOcoAbovePegPriceTypeParameter) Get() *OrderListPlaceOcoAbovePegPriceTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOcoAbovePegPriceTypeParameter) Set(val *OrderListPlaceOcoAbovePegPriceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoAbovePegPriceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoAbovePegPriceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoAbovePegPriceTypeParameter(val *OrderListPlaceOcoAbovePegPriceTypeParameter) *NullableOrderListPlaceOcoAbovePegPriceTypeParameter {
	return &NullableOrderListPlaceOcoAbovePegPriceTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoAbovePegPriceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoAbovePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
