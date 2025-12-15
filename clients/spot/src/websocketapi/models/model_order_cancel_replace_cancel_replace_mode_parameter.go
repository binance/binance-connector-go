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

// OrderCancelReplaceCancelReplaceModeParameter the model 'OrderCancelReplaceCancelReplaceModeParameter'
type OrderCancelReplaceCancelReplaceModeParameter string

// List of orderCancelReplace_cancelReplaceMode_parameter
const (
	OrderCancelReplaceCancelReplaceModeParameterStopOnFailure OrderCancelReplaceCancelReplaceModeParameter = "STOP_ON_FAILURE"
	OrderCancelReplaceCancelReplaceModeParameterAllowFailure  OrderCancelReplaceCancelReplaceModeParameter = "ALLOW_FAILURE"
)

// All allowed values of OrderCancelReplaceCancelReplaceModeParameter enum
var AllowedOrderCancelReplaceCancelReplaceModeParameterEnumValues = []OrderCancelReplaceCancelReplaceModeParameter{
	"STOP_ON_FAILURE",
	"ALLOW_FAILURE",
}

func (v *OrderCancelReplaceCancelReplaceModeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderCancelReplaceCancelReplaceModeParameter(value)
	for _, existing := range AllowedOrderCancelReplaceCancelReplaceModeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderCancelReplaceCancelReplaceModeParameter", value)
}

// NewOrderCancelReplaceCancelReplaceModeParameterFromValue returns a pointer to a valid OrderCancelReplaceCancelReplaceModeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderCancelReplaceCancelReplaceModeParameterFromValue(v string) (*OrderCancelReplaceCancelReplaceModeParameter, error) {
	ev := OrderCancelReplaceCancelReplaceModeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderCancelReplaceCancelReplaceModeParameter: valid values are %v", v, AllowedOrderCancelReplaceCancelReplaceModeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderCancelReplaceCancelReplaceModeParameter) IsValid() bool {
	for _, existing := range AllowedOrderCancelReplaceCancelReplaceModeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderCancelReplace_cancelReplaceMode_parameter value
func (v OrderCancelReplaceCancelReplaceModeParameter) Ptr() *OrderCancelReplaceCancelReplaceModeParameter {
	return &v
}

type NullableOrderCancelReplaceCancelReplaceModeParameter struct {
	value *OrderCancelReplaceCancelReplaceModeParameter
	isSet bool
}

func (v NullableOrderCancelReplaceCancelReplaceModeParameter) Get() *OrderCancelReplaceCancelReplaceModeParameter {
	return v.value
}

func (v *NullableOrderCancelReplaceCancelReplaceModeParameter) Set(val *OrderCancelReplaceCancelReplaceModeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelReplaceCancelReplaceModeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelReplaceCancelReplaceModeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelReplaceCancelReplaceModeParameter(val *OrderCancelReplaceCancelReplaceModeParameter) *NullableOrderCancelReplaceCancelReplaceModeParameter {
	return &NullableOrderCancelReplaceCancelReplaceModeParameter{value: val, isSet: true}
}

func (v NullableOrderCancelReplaceCancelReplaceModeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelReplaceCancelReplaceModeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
