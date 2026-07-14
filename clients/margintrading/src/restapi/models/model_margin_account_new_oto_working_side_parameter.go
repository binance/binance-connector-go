/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOtoWorkingSideParameter the model 'MarginAccountNewOtoWorkingSideParameter'
type MarginAccountNewOtoWorkingSideParameter string

// List of marginAccountNewOto_workingSide_parameter
const (
	MarginAccountNewOtoWorkingSideParameterBuy  MarginAccountNewOtoWorkingSideParameter = "BUY"
	MarginAccountNewOtoWorkingSideParameterSell MarginAccountNewOtoWorkingSideParameter = "SELL"
)

// All allowed values of MarginAccountNewOtoWorkingSideParameter enum
var AllowedMarginAccountNewOtoWorkingSideParameterEnumValues = []MarginAccountNewOtoWorkingSideParameter{
	"BUY",
	"SELL",
}

func (v *MarginAccountNewOtoWorkingSideParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOtoWorkingSideParameter(value)
	for _, existing := range AllowedMarginAccountNewOtoWorkingSideParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOtoWorkingSideParameter", value)
}

// NewMarginAccountNewOtoWorkingSideParameterFromValue returns a pointer to a valid MarginAccountNewOtoWorkingSideParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOtoWorkingSideParameterFromValue(v string) (*MarginAccountNewOtoWorkingSideParameter, error) {
	ev := MarginAccountNewOtoWorkingSideParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOtoWorkingSideParameter: valid values are %v", v, AllowedMarginAccountNewOtoWorkingSideParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOtoWorkingSideParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOtoWorkingSideParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOto_workingSide_parameter value
func (v MarginAccountNewOtoWorkingSideParameter) Ptr() *MarginAccountNewOtoWorkingSideParameter {
	return &v
}

type NullableMarginAccountNewOtoWorkingSideParameter struct {
	value *MarginAccountNewOtoWorkingSideParameter
	isSet bool
}

func (v NullableMarginAccountNewOtoWorkingSideParameter) Get() *MarginAccountNewOtoWorkingSideParameter {
	return v.value
}

func (v *NullableMarginAccountNewOtoWorkingSideParameter) Set(val *MarginAccountNewOtoWorkingSideParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtoWorkingSideParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtoWorkingSideParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtoWorkingSideParameter(val *MarginAccountNewOtoWorkingSideParameter) *NullableMarginAccountNewOtoWorkingSideParameter {
	return &NullableMarginAccountNewOtoWorkingSideParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtoWorkingSideParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtoWorkingSideParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
