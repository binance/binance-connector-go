/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryManagedSubAccountSnapshotTypeParameter the model 'QueryManagedSubAccountSnapshotTypeParameter'
type QueryManagedSubAccountSnapshotTypeParameter string

// List of queryManagedSubAccountSnapshot_type_parameter
const (
	QueryManagedSubAccountSnapshotTypeParameterSpot    QueryManagedSubAccountSnapshotTypeParameter = "SPOT"
	QueryManagedSubAccountSnapshotTypeParameterMargin  QueryManagedSubAccountSnapshotTypeParameter = "MARGIN"
	QueryManagedSubAccountSnapshotTypeParameterFutures QueryManagedSubAccountSnapshotTypeParameter = "FUTURES"
)

// All allowed values of QueryManagedSubAccountSnapshotTypeParameter enum
var AllowedQueryManagedSubAccountSnapshotTypeParameterEnumValues = []QueryManagedSubAccountSnapshotTypeParameter{
	"SPOT",
	"MARGIN",
	"FUTURES",
}

func (v *QueryManagedSubAccountSnapshotTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryManagedSubAccountSnapshotTypeParameter(value)
	for _, existing := range AllowedQueryManagedSubAccountSnapshotTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryManagedSubAccountSnapshotTypeParameter", value)
}

// NewQueryManagedSubAccountSnapshotTypeParameterFromValue returns a pointer to a valid QueryManagedSubAccountSnapshotTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryManagedSubAccountSnapshotTypeParameterFromValue(v string) (*QueryManagedSubAccountSnapshotTypeParameter, error) {
	ev := QueryManagedSubAccountSnapshotTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryManagedSubAccountSnapshotTypeParameter: valid values are %v", v, AllowedQueryManagedSubAccountSnapshotTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryManagedSubAccountSnapshotTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryManagedSubAccountSnapshotTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryManagedSubAccountSnapshot_type_parameter value
func (v QueryManagedSubAccountSnapshotTypeParameter) Ptr() *QueryManagedSubAccountSnapshotTypeParameter {
	return &v
}

type NullableQueryManagedSubAccountSnapshotTypeParameter struct {
	value *QueryManagedSubAccountSnapshotTypeParameter
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotTypeParameter) Get() *QueryManagedSubAccountSnapshotTypeParameter {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotTypeParameter) Set(val *QueryManagedSubAccountSnapshotTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotTypeParameter(val *QueryManagedSubAccountSnapshotTypeParameter) *NullableQueryManagedSubAccountSnapshotTypeParameter {
	return &NullableQueryManagedSubAccountSnapshotTypeParameter{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
