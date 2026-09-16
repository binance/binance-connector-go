/*
Futures (USDⓈ-M) WebSocket API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ModifyOrderReduceOnlyParameter the model 'ModifyOrderReduceOnlyParameter'
type ModifyOrderReduceOnlyParameter string

// List of modifyOrder_reduceOnly_parameter
const (
	ModifyOrderReduceOnlyParameterTrue  ModifyOrderReduceOnlyParameter = "true"
	ModifyOrderReduceOnlyParameterFalse ModifyOrderReduceOnlyParameter = "false"
)

// All allowed values of ModifyOrderReduceOnlyParameter enum
var AllowedModifyOrderReduceOnlyParameterEnumValues = []ModifyOrderReduceOnlyParameter{
	"true",
	"false",
}

func (v *ModifyOrderReduceOnlyParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModifyOrderReduceOnlyParameter(value)
	for _, existing := range AllowedModifyOrderReduceOnlyParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModifyOrderReduceOnlyParameter", value)
}

// NewModifyOrderReduceOnlyParameterFromValue returns a pointer to a valid ModifyOrderReduceOnlyParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModifyOrderReduceOnlyParameterFromValue(v string) (*ModifyOrderReduceOnlyParameter, error) {
	ev := ModifyOrderReduceOnlyParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModifyOrderReduceOnlyParameter: valid values are %v", v, AllowedModifyOrderReduceOnlyParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModifyOrderReduceOnlyParameter) IsValid() bool {
	for _, existing := range AllowedModifyOrderReduceOnlyParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to modifyOrder_reduceOnly_parameter value
func (v ModifyOrderReduceOnlyParameter) Ptr() *ModifyOrderReduceOnlyParameter {
	return &v
}

type NullableModifyOrderReduceOnlyParameter struct {
	value *ModifyOrderReduceOnlyParameter
	isSet bool
}

func (v NullableModifyOrderReduceOnlyParameter) Get() *ModifyOrderReduceOnlyParameter {
	return v.value
}

func (v *NullableModifyOrderReduceOnlyParameter) Set(val *ModifyOrderReduceOnlyParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyOrderReduceOnlyParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyOrderReduceOnlyParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyOrderReduceOnlyParameter(val *ModifyOrderReduceOnlyParameter) *NullableModifyOrderReduceOnlyParameter {
	return &NullableModifyOrderReduceOnlyParameter{value: val, isSet: true}
}

func (v NullableModifyOrderReduceOnlyParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyOrderReduceOnlyParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
