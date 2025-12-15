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

// OrderListPlaceOcoAbovePegOffsetTypeParameter the model 'OrderListPlaceOcoAbovePegOffsetTypeParameter'
type OrderListPlaceOcoAbovePegOffsetTypeParameter string

// List of orderListPlaceOco_abovePegOffsetType_parameter
const (
	OrderListPlaceOcoAbovePegOffsetTypeParameterPriceLevel OrderListPlaceOcoAbovePegOffsetTypeParameter = "PRICE_LEVEL"
)

// All allowed values of OrderListPlaceOcoAbovePegOffsetTypeParameter enum
var AllowedOrderListPlaceOcoAbovePegOffsetTypeParameterEnumValues = []OrderListPlaceOcoAbovePegOffsetTypeParameter{
	"PRICE_LEVEL",
}

func (v *OrderListPlaceOcoAbovePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOcoAbovePegOffsetTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOcoAbovePegOffsetTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOcoAbovePegOffsetTypeParameter", value)
}

// NewOrderListPlaceOcoAbovePegOffsetTypeParameterFromValue returns a pointer to a valid OrderListPlaceOcoAbovePegOffsetTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOcoAbovePegOffsetTypeParameterFromValue(v string) (*OrderListPlaceOcoAbovePegOffsetTypeParameter, error) {
	ev := OrderListPlaceOcoAbovePegOffsetTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOcoAbovePegOffsetTypeParameter: valid values are %v", v, AllowedOrderListPlaceOcoAbovePegOffsetTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOcoAbovePegOffsetTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOcoAbovePegOffsetTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOco_abovePegOffsetType_parameter value
func (v OrderListPlaceOcoAbovePegOffsetTypeParameter) Ptr() *OrderListPlaceOcoAbovePegOffsetTypeParameter {
	return &v
}

type NullableOrderListPlaceOcoAbovePegOffsetTypeParameter struct {
	value *OrderListPlaceOcoAbovePegOffsetTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) Get() *OrderListPlaceOcoAbovePegOffsetTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) Set(val *OrderListPlaceOcoAbovePegOffsetTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoAbovePegOffsetTypeParameter(val *OrderListPlaceOcoAbovePegOffsetTypeParameter) *NullableOrderListPlaceOcoAbovePegOffsetTypeParameter {
	return &NullableOrderListPlaceOcoAbovePegOffsetTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoAbovePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
