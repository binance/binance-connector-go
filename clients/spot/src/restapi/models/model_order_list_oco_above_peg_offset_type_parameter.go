/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOcoAbovePegOffsetTypeParameter the model 'OrderListOcoAbovePegOffsetTypeParameter'
type OrderListOcoAbovePegOffsetTypeParameter string

// List of orderListOco_abovePegOffsetType_parameter
const (
	OrderListOcoAbovePegOffsetTypeParameterPriceLevel OrderListOcoAbovePegOffsetTypeParameter = "PRICE_LEVEL"
)

// All allowed values of OrderListOcoAbovePegOffsetTypeParameter enum
var AllowedOrderListOcoAbovePegOffsetTypeParameterEnumValues = []OrderListOcoAbovePegOffsetTypeParameter{
	"PRICE_LEVEL",
}

func (v *OrderListOcoAbovePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOcoAbovePegOffsetTypeParameter(value)
	for _, existing := range AllowedOrderListOcoAbovePegOffsetTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOcoAbovePegOffsetTypeParameter", value)
}

// NewOrderListOcoAbovePegOffsetTypeParameterFromValue returns a pointer to a valid OrderListOcoAbovePegOffsetTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOcoAbovePegOffsetTypeParameterFromValue(v string) (*OrderListOcoAbovePegOffsetTypeParameter, error) {
	ev := OrderListOcoAbovePegOffsetTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOcoAbovePegOffsetTypeParameter: valid values are %v", v, AllowedOrderListOcoAbovePegOffsetTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOcoAbovePegOffsetTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOcoAbovePegOffsetTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOco_abovePegOffsetType_parameter value
func (v OrderListOcoAbovePegOffsetTypeParameter) Ptr() *OrderListOcoAbovePegOffsetTypeParameter {
	return &v
}

type NullableOrderListOcoAbovePegOffsetTypeParameter struct {
	value *OrderListOcoAbovePegOffsetTypeParameter
	isSet bool
}

func (v NullableOrderListOcoAbovePegOffsetTypeParameter) Get() *OrderListOcoAbovePegOffsetTypeParameter {
	return v.value
}

func (v *NullableOrderListOcoAbovePegOffsetTypeParameter) Set(val *OrderListOcoAbovePegOffsetTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOcoAbovePegOffsetTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOcoAbovePegOffsetTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOcoAbovePegOffsetTypeParameter(val *OrderListOcoAbovePegOffsetTypeParameter) *NullableOrderListOcoAbovePegOffsetTypeParameter {
	return &NullableOrderListOcoAbovePegOffsetTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOcoAbovePegOffsetTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOcoAbovePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
