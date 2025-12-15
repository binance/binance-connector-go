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

// OrderOcoStopLimitTimeInForceParameter the model 'OrderOcoStopLimitTimeInForceParameter'
type OrderOcoStopLimitTimeInForceParameter string

// List of orderOco_stopLimitTimeInForce_parameter
const (
	OrderOcoStopLimitTimeInForceParameterGtc OrderOcoStopLimitTimeInForceParameter = "GTC"
	OrderOcoStopLimitTimeInForceParameterIoc OrderOcoStopLimitTimeInForceParameter = "IOC"
	OrderOcoStopLimitTimeInForceParameterFok OrderOcoStopLimitTimeInForceParameter = "FOK"
)

// All allowed values of OrderOcoStopLimitTimeInForceParameter enum
var AllowedOrderOcoStopLimitTimeInForceParameterEnumValues = []OrderOcoStopLimitTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *OrderOcoStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderOcoStopLimitTimeInForceParameter(value)
	for _, existing := range AllowedOrderOcoStopLimitTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderOcoStopLimitTimeInForceParameter", value)
}

// NewOrderOcoStopLimitTimeInForceParameterFromValue returns a pointer to a valid OrderOcoStopLimitTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderOcoStopLimitTimeInForceParameterFromValue(v string) (*OrderOcoStopLimitTimeInForceParameter, error) {
	ev := OrderOcoStopLimitTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderOcoStopLimitTimeInForceParameter: valid values are %v", v, AllowedOrderOcoStopLimitTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderOcoStopLimitTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedOrderOcoStopLimitTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderOco_stopLimitTimeInForce_parameter value
func (v OrderOcoStopLimitTimeInForceParameter) Ptr() *OrderOcoStopLimitTimeInForceParameter {
	return &v
}

type NullableOrderOcoStopLimitTimeInForceParameter struct {
	value *OrderOcoStopLimitTimeInForceParameter
	isSet bool
}

func (v NullableOrderOcoStopLimitTimeInForceParameter) Get() *OrderOcoStopLimitTimeInForceParameter {
	return v.value
}

func (v *NullableOrderOcoStopLimitTimeInForceParameter) Set(val *OrderOcoStopLimitTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderOcoStopLimitTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderOcoStopLimitTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderOcoStopLimitTimeInForceParameter(val *OrderOcoStopLimitTimeInForceParameter) *NullableOrderOcoStopLimitTimeInForceParameter {
	return &NullableOrderOcoStopLimitTimeInForceParameter{value: val, isSet: true}
}

func (v NullableOrderOcoStopLimitTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderOcoStopLimitTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
