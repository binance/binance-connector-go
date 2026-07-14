/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewBlockTradeOrderLegsParameterInnerType Order type
type NewBlockTradeOrderLegsParameterInnerType string

// List of newBlockTradeOrder_legs_parameter_inner_type
const (
	NewBlockTradeOrderLegsParameterInnerTypeLimit NewBlockTradeOrderLegsParameterInnerType = "LIMIT"
)

// All allowed values of NewBlockTradeOrderLegsParameterInnerType enum
var AllowedNewBlockTradeOrderLegsParameterInnerTypeEnumValues = []NewBlockTradeOrderLegsParameterInnerType{
	"LIMIT",
}

func (v *NewBlockTradeOrderLegsParameterInnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewBlockTradeOrderLegsParameterInnerType(value)
	for _, existing := range AllowedNewBlockTradeOrderLegsParameterInnerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewBlockTradeOrderLegsParameterInnerType", value)
}

// NewNewBlockTradeOrderLegsParameterInnerTypeFromValue returns a pointer to a valid NewBlockTradeOrderLegsParameterInnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewBlockTradeOrderLegsParameterInnerTypeFromValue(v string) (*NewBlockTradeOrderLegsParameterInnerType, error) {
	ev := NewBlockTradeOrderLegsParameterInnerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewBlockTradeOrderLegsParameterInnerType: valid values are %v", v, AllowedNewBlockTradeOrderLegsParameterInnerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewBlockTradeOrderLegsParameterInnerType) IsValid() bool {
	for _, existing := range AllowedNewBlockTradeOrderLegsParameterInnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newBlockTradeOrder_legs_parameter_inner_type value
func (v NewBlockTradeOrderLegsParameterInnerType) Ptr() *NewBlockTradeOrderLegsParameterInnerType {
	return &v
}

type NullableNewBlockTradeOrderLegsParameterInnerType struct {
	value *NewBlockTradeOrderLegsParameterInnerType
	isSet bool
}

func (v NullableNewBlockTradeOrderLegsParameterInnerType) Get() *NewBlockTradeOrderLegsParameterInnerType {
	return v.value
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerType) Set(val *NewBlockTradeOrderLegsParameterInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullableNewBlockTradeOrderLegsParameterInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewBlockTradeOrderLegsParameterInnerType(val *NewBlockTradeOrderLegsParameterInnerType) *NullableNewBlockTradeOrderLegsParameterInnerType {
	return &NullableNewBlockTradeOrderLegsParameterInnerType{value: val, isSet: true}
}

func (v NullableNewBlockTradeOrderLegsParameterInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewBlockTradeOrderLegsParameterInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
