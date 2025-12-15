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

// OrderCancelReplacePegPriceTypeParameter the model 'OrderCancelReplacePegPriceTypeParameter'
type OrderCancelReplacePegPriceTypeParameter string

// List of orderCancelReplace_pegPriceType_parameter
const (
	OrderCancelReplacePegPriceTypeParameterPrimaryPeg       OrderCancelReplacePegPriceTypeParameter = "PRIMARY_PEG"
	OrderCancelReplacePegPriceTypeParameterMarketPeg        OrderCancelReplacePegPriceTypeParameter = "MARKET_PEG"
	OrderCancelReplacePegPriceTypeParameterNonRepresentable OrderCancelReplacePegPriceTypeParameter = "NON_REPRESENTABLE"
)

// All allowed values of OrderCancelReplacePegPriceTypeParameter enum
var AllowedOrderCancelReplacePegPriceTypeParameterEnumValues = []OrderCancelReplacePegPriceTypeParameter{
	"PRIMARY_PEG",
	"MARKET_PEG",
	"NON_REPRESENTABLE",
}

func (v *OrderCancelReplacePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplacePegPriceTypeParameter(value)
	for _, existing := range AllowedOrderCancelReplacePegPriceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplacePegPriceTypeParameter", value)
}

// NewOrderCancelReplacePegPriceTypeParameterFromValue returns a pointer to a valid OrderCancelReplacePegPriceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplacePegPriceTypeParameterFromValue(v string) (*OrderCancelReplacePegPriceTypeParameter, error) {
	ev := OrderCancelReplacePegPriceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplacePegPriceTypeParameter: valid values are %v", v, AllowedOrderCancelReplacePegPriceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplacePegPriceTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplacePegPriceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_pegPriceType_parameter value
func (v OrderCancelReplacePegPriceTypeParameter) Ptr() *OrderCancelReplacePegPriceTypeParameter {
	return &v
}

type NullableOrderCancelReplacePegPriceTypeParameter struct {
	value *OrderCancelReplacePegPriceTypeParameter
	isSet bool
}

func (v NullableOrderCancelReplacePegPriceTypeParameter) Get() *OrderCancelReplacePegPriceTypeParameter {
	return v.value
}

func (v *NullableOrderCancelReplacePegPriceTypeParameter) Set(val *OrderCancelReplacePegPriceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplacePegPriceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplacePegPriceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplacePegPriceTypeParameter(val *OrderCancelReplacePegPriceTypeParameter) *NullableOrderCancelReplacePegPriceTypeParameter {
	return &NullableOrderCancelReplacePegPriceTypeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplacePegPriceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplacePegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
