/*
Wallet REST API

Query balances, manage assets, and perform wallet operations via the Binance Wallet API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryUserDelegationHistoryTypeParameter the model 'QueryUserDelegationHistoryTypeParameter'
type QueryUserDelegationHistoryTypeParameter string

// List of queryUserDelegationHistory_type_parameter
const (
	QueryUserDelegationHistoryTypeParameterDelegate   QueryUserDelegationHistoryTypeParameter = "DELEGATE"
	QueryUserDelegationHistoryTypeParameterUndelegate QueryUserDelegationHistoryTypeParameter = "UNDELEGATE"
)

// All allowed values of QueryUserDelegationHistoryTypeParameter enum
var AllowedQueryUserDelegationHistoryTypeParameterEnumValues = []QueryUserDelegationHistoryTypeParameter{
	"DELEGATE",
	"UNDELEGATE",
}

func (v *QueryUserDelegationHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryUserDelegationHistoryTypeParameter(value)
	for _, existing := range AllowedQueryUserDelegationHistoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryUserDelegationHistoryTypeParameter", value)
}

// NewQueryUserDelegationHistoryTypeParameterFromValue returns a pointer to a valid QueryUserDelegationHistoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryUserDelegationHistoryTypeParameterFromValue(v string) (*QueryUserDelegationHistoryTypeParameter, error) {
	ev := QueryUserDelegationHistoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryUserDelegationHistoryTypeParameter: valid values are %v", v, AllowedQueryUserDelegationHistoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryUserDelegationHistoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryUserDelegationHistoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryUserDelegationHistory_type_parameter value
func (v QueryUserDelegationHistoryTypeParameter) Ptr() *QueryUserDelegationHistoryTypeParameter {
	return &v
}

type NullableQueryUserDelegationHistoryTypeParameter struct {
	value *QueryUserDelegationHistoryTypeParameter
	isSet bool
}

func (v NullableQueryUserDelegationHistoryTypeParameter) Get() *QueryUserDelegationHistoryTypeParameter {
	return v.value
}

func (v *NullableQueryUserDelegationHistoryTypeParameter) Set(val *QueryUserDelegationHistoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserDelegationHistoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserDelegationHistoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserDelegationHistoryTypeParameter(val *QueryUserDelegationHistoryTypeParameter) *NullableQueryUserDelegationHistoryTypeParameter {
	return &NullableQueryUserDelegationHistoryTypeParameter{value: val, isSet: true}
}

func (v NullableQueryUserDelegationHistoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserDelegationHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
