/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewAlgoOrderSideParameter the model 'NewAlgoOrderSideParameter'
type NewAlgoOrderSideParameter string

// List of newAlgoOrder_side_parameter
const (
	NewAlgoOrderSideParameterBuy  NewAlgoOrderSideParameter = "BUY"
	NewAlgoOrderSideParameterSell NewAlgoOrderSideParameter = "SELL"
)

// All allowed values of NewAlgoOrderSideParameter enum
var AllowedNewAlgoOrderSideParameterEnumValues = []NewAlgoOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *NewAlgoOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewAlgoOrderSideParameter(value)
	for _, existing := range AllowedNewAlgoOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewAlgoOrderSideParameter", value)
}

// NewNewAlgoOrderSideParameterFromValue returns a pointer to a valid NewAlgoOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewAlgoOrderSideParameterFromValue(v string) (*NewAlgoOrderSideParameter, error) {
	ev := NewAlgoOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewAlgoOrderSideParameter: valid values are %v", v, AllowedNewAlgoOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewAlgoOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedNewAlgoOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newAlgoOrder_side_parameter value
func (v NewAlgoOrderSideParameter) Ptr() *NewAlgoOrderSideParameter {
	return &v
}

type NullableNewAlgoOrderSideParameter struct {
	value *NewAlgoOrderSideParameter
	isSet bool
}

func (v NullableNewAlgoOrderSideParameter) Get() *NewAlgoOrderSideParameter {
	return v.value
}

func (v *NullableNewAlgoOrderSideParameter) Set(val *NewAlgoOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAlgoOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAlgoOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAlgoOrderSideParameter(val *NewAlgoOrderSideParameter) *NullableNewAlgoOrderSideParameter {
	return &NullableNewAlgoOrderSideParameter{value: val, isSet: true}
}

func (v NullableNewAlgoOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAlgoOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
