/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderReduceOnlyParameter the model 'NewOrderReduceOnlyParameter'
type NewOrderReduceOnlyParameter string

// List of newOrder_reduceOnly_parameter
const (
	NewOrderReduceOnlyParameterTrue  NewOrderReduceOnlyParameter = "true"
	NewOrderReduceOnlyParameterFalse NewOrderReduceOnlyParameter = "false"
)

// All allowed values of NewOrderReduceOnlyParameter enum
var AllowedNewOrderReduceOnlyParameterEnumValues = []NewOrderReduceOnlyParameter{
	"true",
	"false",
}

func (v *NewOrderReduceOnlyParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderReduceOnlyParameter(value)
	for _, existing := range AllowedNewOrderReduceOnlyParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderReduceOnlyParameter", value)
}

// NewNewOrderReduceOnlyParameterFromValue returns a pointer to a valid NewOrderReduceOnlyParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderReduceOnlyParameterFromValue(v string) (*NewOrderReduceOnlyParameter, error) {
	ev := NewOrderReduceOnlyParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderReduceOnlyParameter: valid values are %v", v, AllowedNewOrderReduceOnlyParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderReduceOnlyParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderReduceOnlyParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_reduceOnly_parameter value
func (v NewOrderReduceOnlyParameter) Ptr() *NewOrderReduceOnlyParameter {
	return &v
}

type NullableNewOrderReduceOnlyParameter struct {
	value *NewOrderReduceOnlyParameter
	isSet bool
}

func (v NullableNewOrderReduceOnlyParameter) Get() *NewOrderReduceOnlyParameter {
	return v.value
}

func (v *NullableNewOrderReduceOnlyParameter) Set(val *NewOrderReduceOnlyParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderReduceOnlyParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderReduceOnlyParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderReduceOnlyParameter(val *NewOrderReduceOnlyParameter) *NullableNewOrderReduceOnlyParameter {
	return &NullableNewOrderReduceOnlyParameter{value: val, isSet: true}
}

func (v NullableNewOrderReduceOnlyParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderReduceOnlyParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
