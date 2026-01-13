/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelReplacePegOffsetTypeParameter the model 'OrderCancelReplacePegOffsetTypeParameter'
type OrderCancelReplacePegOffsetTypeParameter string

// List of orderCancelReplace_pegOffsetType_parameter
const (
	OrderCancelReplacePegOffsetTypeParameterPriceLevel       OrderCancelReplacePegOffsetTypeParameter = "PRICE_LEVEL"
	OrderCancelReplacePegOffsetTypeParameterNonRepresentable OrderCancelReplacePegOffsetTypeParameter = "NON_REPRESENTABLE"
)

// All allowed values of OrderCancelReplacePegOffsetTypeParameter enum
var AllowedOrderCancelReplacePegOffsetTypeParameterEnumValues = []OrderCancelReplacePegOffsetTypeParameter{
	"PRICE_LEVEL",
	"NON_REPRESENTABLE",
}

func (v *OrderCancelReplacePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplacePegOffsetTypeParameter(value)
	for _, existing := range AllowedOrderCancelReplacePegOffsetTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplacePegOffsetTypeParameter", value)
}

// NewOrderCancelReplacePegOffsetTypeParameterFromValue returns a pointer to a valid OrderCancelReplacePegOffsetTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplacePegOffsetTypeParameterFromValue(v string) (*OrderCancelReplacePegOffsetTypeParameter, error) {
	ev := OrderCancelReplacePegOffsetTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplacePegOffsetTypeParameter: valid values are %v", v, AllowedOrderCancelReplacePegOffsetTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplacePegOffsetTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplacePegOffsetTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_pegOffsetType_parameter value
func (v OrderCancelReplacePegOffsetTypeParameter) Ptr() *OrderCancelReplacePegOffsetTypeParameter {
	return &v
}

type NullableOrderCancelReplacePegOffsetTypeParameter struct {
	value *OrderCancelReplacePegOffsetTypeParameter
	isSet bool
}

func (v NullableOrderCancelReplacePegOffsetTypeParameter) Get() *OrderCancelReplacePegOffsetTypeParameter {
	return v.value
}

func (v *NullableOrderCancelReplacePegOffsetTypeParameter) Set(val *OrderCancelReplacePegOffsetTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplacePegOffsetTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplacePegOffsetTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplacePegOffsetTypeParameter(val *OrderCancelReplacePegOffsetTypeParameter) *NullableOrderCancelReplacePegOffsetTypeParameter {
	return &NullableOrderCancelReplacePegOffsetTypeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplacePegOffsetTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplacePegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
