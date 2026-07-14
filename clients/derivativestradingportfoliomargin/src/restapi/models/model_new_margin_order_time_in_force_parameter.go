/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewMarginOrderTimeInForceParameter the model 'NewMarginOrderTimeInForceParameter'
type NewMarginOrderTimeInForceParameter string

// List of newMarginOrder_timeInForce_parameter
const (
	NewMarginOrderTimeInForceParameterGtc NewMarginOrderTimeInForceParameter = "GTC"
	NewMarginOrderTimeInForceParameterIoc NewMarginOrderTimeInForceParameter = "IOC"
	NewMarginOrderTimeInForceParameterFok NewMarginOrderTimeInForceParameter = "FOK"
)

// All allowed values of NewMarginOrderTimeInForceParameter enum
var AllowedNewMarginOrderTimeInForceParameterEnumValues = []NewMarginOrderTimeInForceParameter{
	"GTC",
	"IOC",
	"FOK",
}

func (v *NewMarginOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewMarginOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewMarginOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewMarginOrderTimeInForceParameter", value)
}

// NewNewMarginOrderTimeInForceParameterFromValue returns a pointer to a valid NewMarginOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewMarginOrderTimeInForceParameterFromValue(v string) (*NewMarginOrderTimeInForceParameter, error) {
	ev := NewMarginOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewMarginOrderTimeInForceParameter: valid values are %v", v, AllowedNewMarginOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewMarginOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewMarginOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newMarginOrder_timeInForce_parameter value
func (v NewMarginOrderTimeInForceParameter) Ptr() *NewMarginOrderTimeInForceParameter {
	return &v
}

type NullableNewMarginOrderTimeInForceParameter struct {
	value *NewMarginOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewMarginOrderTimeInForceParameter) Get() *NewMarginOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewMarginOrderTimeInForceParameter) Set(val *NewMarginOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewMarginOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewMarginOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewMarginOrderTimeInForceParameter(val *NewMarginOrderTimeInForceParameter) *NullableNewMarginOrderTimeInForceParameter {
	return &NullableNewMarginOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewMarginOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewMarginOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
