/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewCmConditionalOrderSideParameter the model 'NewCmConditionalOrderSideParameter'
type NewCmConditionalOrderSideParameter string

// List of newCmConditionalOrder_side_parameter
const (
	NewCmConditionalOrderSideParameterBuy  NewCmConditionalOrderSideParameter = "BUY"
	NewCmConditionalOrderSideParameterSell NewCmConditionalOrderSideParameter = "SELL"
)

// All allowed values of NewCmConditionalOrderSideParameter enum
var AllowedNewCmConditionalOrderSideParameterEnumValues = []NewCmConditionalOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *NewCmConditionalOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewCmConditionalOrderSideParameter(value)
	for _, existing := range AllowedNewCmConditionalOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewCmConditionalOrderSideParameter", value)
}

// NewNewCmConditionalOrderSideParameterFromValue returns a pointer to a valid NewCmConditionalOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewCmConditionalOrderSideParameterFromValue(v string) (*NewCmConditionalOrderSideParameter, error) {
	ev := NewCmConditionalOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewCmConditionalOrderSideParameter: valid values are %v", v, AllowedNewCmConditionalOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewCmConditionalOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedNewCmConditionalOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newCmConditionalOrder_side_parameter value
func (v NewCmConditionalOrderSideParameter) Ptr() *NewCmConditionalOrderSideParameter {
	return &v
}

type NullableNewCmConditionalOrderSideParameter struct {
	value *NewCmConditionalOrderSideParameter
	isSet bool
}

func (v NullableNewCmConditionalOrderSideParameter) Get() *NewCmConditionalOrderSideParameter {
	return v.value
}

func (v *NullableNewCmConditionalOrderSideParameter) Set(val *NewCmConditionalOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCmConditionalOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCmConditionalOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCmConditionalOrderSideParameter(val *NewCmConditionalOrderSideParameter) *NullableNewCmConditionalOrderSideParameter {
	return &NullableNewCmConditionalOrderSideParameter{value: val, isSet: true}
}

func (v NullableNewCmConditionalOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCmConditionalOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
