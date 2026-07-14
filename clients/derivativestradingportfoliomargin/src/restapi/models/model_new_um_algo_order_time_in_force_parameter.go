/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderTimeInForceParameter the model 'NewUmAlgoOrderTimeInForceParameter'
type NewUmAlgoOrderTimeInForceParameter string

// List of newUmAlgoOrder_timeInForce_parameter
const (
	NewUmAlgoOrderTimeInForceParameterIoc NewUmAlgoOrderTimeInForceParameter = "IOC"
	NewUmAlgoOrderTimeInForceParameterGtc NewUmAlgoOrderTimeInForceParameter = "GTC"
	NewUmAlgoOrderTimeInForceParameterFok NewUmAlgoOrderTimeInForceParameter = "FOK"
	NewUmAlgoOrderTimeInForceParameterGtx NewUmAlgoOrderTimeInForceParameter = "GTX"
	NewUmAlgoOrderTimeInForceParameterGtd NewUmAlgoOrderTimeInForceParameter = "GTD"
)

// All allowed values of NewUmAlgoOrderTimeInForceParameter enum
var AllowedNewUmAlgoOrderTimeInForceParameterEnumValues = []NewUmAlgoOrderTimeInForceParameter{
	"IOC",
	"GTC",
	"FOK",
	"GTX",
	"GTD",
}

func (v *NewUmAlgoOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderTimeInForceParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderTimeInForceParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderTimeInForceParameter", value)
}

// NewNewUmAlgoOrderTimeInForceParameterFromValue returns a pointer to a valid NewUmAlgoOrderTimeInForceParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderTimeInForceParameterFromValue(v string) (*NewUmAlgoOrderTimeInForceParameter, error) {
	ev := NewUmAlgoOrderTimeInForceParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderTimeInForceParameter: valid values are %v", v, AllowedNewUmAlgoOrderTimeInForceParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderTimeInForceParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderTimeInForceParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_timeInForce_parameter value
func (v NewUmAlgoOrderTimeInForceParameter) Ptr() *NewUmAlgoOrderTimeInForceParameter {
	return &v
}

type NullableNewUmAlgoOrderTimeInForceParameter struct {
	value *NewUmAlgoOrderTimeInForceParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderTimeInForceParameter) Get() *NewUmAlgoOrderTimeInForceParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderTimeInForceParameter) Set(val *NewUmAlgoOrderTimeInForceParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderTimeInForceParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderTimeInForceParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderTimeInForceParameter(val *NewUmAlgoOrderTimeInForceParameter) *NullableNewUmAlgoOrderTimeInForceParameter {
	return &NullableNewUmAlgoOrderTimeInForceParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderTimeInForceParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderTimeInForceParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
