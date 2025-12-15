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

// NewOrderSideParameter the model 'NewOrderSideParameter'
type NewOrderSideParameter string

// List of newOrder_side_parameter
const (
	NewOrderSideParameterBuy  NewOrderSideParameter = "BUY"
	NewOrderSideParameterSell NewOrderSideParameter = "SELL"
)

// All allowed values of NewOrderSideParameter enum
var AllowedNewOrderSideParameterEnumValues = []NewOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *NewOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderSideParameter(value)
	for _, existing := range AllowedNewOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderSideParameter", value)
}

// NewNewOrderSideParameterFromValue returns a pointer to a valid NewOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderSideParameterFromValue(v string) (*NewOrderSideParameter, error) {
	ev := NewOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderSideParameter: valid values are %v", v, AllowedNewOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_side_parameter value
func (v NewOrderSideParameter) Ptr() *NewOrderSideParameter {
	return &v
}

type NullableNewOrderSideParameter struct {
	value *NewOrderSideParameter
	isSet bool
}

func (v NullableNewOrderSideParameter) Get() *NewOrderSideParameter {
	return v.value
}

func (v *NullableNewOrderSideParameter) Set(val *NewOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderSideParameter(val *NewOrderSideParameter) *NullableNewOrderSideParameter {
	return &NullableNewOrderSideParameter{value: val, isSet: true}
}

func (v NullableNewOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
