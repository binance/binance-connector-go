/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListPlaceStopLimitTimeInForceParameter the model 'OrderListPlaceStopLimitTimeInForceParameter'
type OrderListPlaceStopLimitTimeInForceParameter string

// List of orderListPlace_stopLimitTimeInForce_parameter
const (
	OrderListPlaceStopLimitTimeInForceParameterGtc OrderListPlaceStopLimitTimeInForceParameter = "GTC"
	OrderListPlaceStopLimitTimeInForceParameterIoc OrderListPlaceStopLimitTimeInForceParameter = "IOC"
	OrderListPlaceStopLimitTimeInForceParameterFok OrderListPlaceStopLimitTimeInForceParameter = "FOK"
)

// All allowed values of OrderListPlaceStopLimitTimeInForceParameter enum
var AllowedOrderListPlaceStopLimitTimeInForceParameterEnumValues = []OrderListPlaceStopLimitTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *OrderListPlaceStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceStopLimitTimeInForceParameter(value)
	for _, existing := range AllowedOrderListPlaceStopLimitTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceStopLimitTimeInForceParameter", value)
}

// NewOrderListPlaceStopLimitTimeInForceParameterFromValue returns a pointer to a valid OrderListPlaceStopLimitTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceStopLimitTimeInForceParameterFromValue(v string) (*OrderListPlaceStopLimitTimeInForceParameter, error) {
	ev := OrderListPlaceStopLimitTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceStopLimitTimeInForceParameter: valid values are %v", v, AllowedOrderListPlaceStopLimitTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceStopLimitTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceStopLimitTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlace_stopLimitTimeInForce_parameter value
func (v OrderListPlaceStopLimitTimeInForceParameter) Ptr() *OrderListPlaceStopLimitTimeInForceParameter {
	return &v
}

type NullableOrderListPlaceStopLimitTimeInForceParameter struct {
	value *OrderListPlaceStopLimitTimeInForceParameter
	isSet bool
}

func (v NullableOrderListPlaceStopLimitTimeInForceParameter) Get() *OrderListPlaceStopLimitTimeInForceParameter {
	return v.value
}

func (v *NullableOrderListPlaceStopLimitTimeInForceParameter) Set(val *OrderListPlaceStopLimitTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceStopLimitTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceStopLimitTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceStopLimitTimeInForceParameter(val *OrderListPlaceStopLimitTimeInForceParameter) *NullableOrderListPlaceStopLimitTimeInForceParameter {
	return &NullableOrderListPlaceStopLimitTimeInForceParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceStopLimitTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
