/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewUmAlgoOrderSideParameter the model 'NewUmAlgoOrderSideParameter'
type NewUmAlgoOrderSideParameter string

// List of newUmAlgoOrder_side_parameter
const (
	NewUmAlgoOrderSideParameterBuy  NewUmAlgoOrderSideParameter = "BUY"
	NewUmAlgoOrderSideParameterSell NewUmAlgoOrderSideParameter = "SELL"
)

// All allowed values of NewUmAlgoOrderSideParameter enum
var AllowedNewUmAlgoOrderSideParameterEnumValues = []NewUmAlgoOrderSideParameter{
	"BUY",
	"SELL",
}

func (v *NewUmAlgoOrderSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewUmAlgoOrderSideParameter(value)
	for _, existing := range AllowedNewUmAlgoOrderSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewUmAlgoOrderSideParameter", value)
}

// NewNewUmAlgoOrderSideParameterFromValue returns a pointer to a valid NewUmAlgoOrderSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewUmAlgoOrderSideParameterFromValue(v string) (*NewUmAlgoOrderSideParameter, error) {
	ev := NewUmAlgoOrderSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewUmAlgoOrderSideParameter: valid values are %v", v, AllowedNewUmAlgoOrderSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewUmAlgoOrderSideParameter) IsValid() bool {
	for _, existing := range AllowedNewUmAlgoOrderSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newUmAlgoOrder_side_parameter value
func (v NewUmAlgoOrderSideParameter) Ptr() *NewUmAlgoOrderSideParameter {
	return &v
}

type NullableNewUmAlgoOrderSideParameter struct {
	value *NewUmAlgoOrderSideParameter
	isSet bool
}

func (v NullableNewUmAlgoOrderSideParameter) Get() *NewUmAlgoOrderSideParameter {
	return v.value
}

func (v *NullableNewUmAlgoOrderSideParameter) Set(val *NewUmAlgoOrderSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmAlgoOrderSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmAlgoOrderSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmAlgoOrderSideParameter(val *NewUmAlgoOrderSideParameter) *NullableNewUmAlgoOrderSideParameter {
	return &NullableNewUmAlgoOrderSideParameter{value: val, isSet: true}
}

func (v NullableNewUmAlgoOrderSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmAlgoOrderSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
