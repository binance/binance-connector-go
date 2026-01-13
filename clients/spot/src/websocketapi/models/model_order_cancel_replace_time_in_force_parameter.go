/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderCancelReplaceTimeInForceParameter the model 'OrderCancelReplaceTimeInForceParameter'
type OrderCancelReplaceTimeInForceParameter string

// List of orderCancelReplace_timeInForce_parameter
const (
	OrderCancelReplaceTimeInForceParameterGtc              OrderCancelReplaceTimeInForceParameter = "GTC"
	OrderCancelReplaceTimeInForceParameterIoc              OrderCancelReplaceTimeInForceParameter = "IOC"
	OrderCancelReplaceTimeInForceParameterFok              OrderCancelReplaceTimeInForceParameter = "FOK"
	OrderCancelReplaceTimeInForceParameterNonRepresentable OrderCancelReplaceTimeInForceParameter = "NON_REPRESENTABLE"
)

// All allowed values of OrderCancelReplaceTimeInForceParameter enum
var AllowedOrderCancelReplaceTimeInForceParameterEnumValues = []OrderCancelReplaceTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
	"NON_REPRESENTABLE",
}

func (v *OrderCancelReplaceTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceTimeInForceParameter(value)
	for _, existing := range AllowedOrderCancelReplaceTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceTimeInForceParameter", value)
}

// NewOrderCancelReplaceTimeInForceParameterFromValue returns a pointer to a valid OrderCancelReplaceTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceTimeInForceParameterFromValue(v string) (*OrderCancelReplaceTimeInForceParameter, error) {
	ev := OrderCancelReplaceTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceTimeInForceParameter: valid values are %v", v, AllowedOrderCancelReplaceTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_timeInForce_parameter value
func (v OrderCancelReplaceTimeInForceParameter) Ptr() *OrderCancelReplaceTimeInForceParameter {
	return &v
}

type NullableOrderCancelReplaceTimeInForceParameter struct {
	value *OrderCancelReplaceTimeInForceParameter
	isSet bool
}

func (v NullableOrderCancelReplaceTimeInForceParameter) Get() *OrderCancelReplaceTimeInForceParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceTimeInForceParameter) Set(val *OrderCancelReplaceTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceTimeInForceParameter(val *OrderCancelReplaceTimeInForceParameter) *NullableOrderCancelReplaceTimeInForceParameter {
	return &NullableOrderCancelReplaceTimeInForceParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
