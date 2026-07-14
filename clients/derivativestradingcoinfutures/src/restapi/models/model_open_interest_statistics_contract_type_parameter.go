/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OpenInterestStatisticsContractTypeParameter the model 'OpenInterestStatisticsContractTypeParameter'
type OpenInterestStatisticsContractTypeParameter string

// List of openInterestStatistics_contractType_parameter
const (
	OpenInterestStatisticsContractTypeParameterAll            OpenInterestStatisticsContractTypeParameter = "ALL"
	OpenInterestStatisticsContractTypeParameterPerpetual      OpenInterestStatisticsContractTypeParameter = "PERPETUAL"
	OpenInterestStatisticsContractTypeParameterCurrentQuarter OpenInterestStatisticsContractTypeParameter = "CURRENT_QUARTER"
	OpenInterestStatisticsContractTypeParameterNextQuarter    OpenInterestStatisticsContractTypeParameter = "NEXT_QUARTER"
)

// All allowed values of OpenInterestStatisticsContractTypeParameter enum
var AllowedOpenInterestStatisticsContractTypeParameterEnumValues = []OpenInterestStatisticsContractTypeParameter{
	"ALL",
	"PERPETUAL",
	"CURRENT_QUARTER",
	"NEXT_QUARTER",
}

func (v *OpenInterestStatisticsContractTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OpenInterestStatisticsContractTypeParameter(value)
	for _, existing := range AllowedOpenInterestStatisticsContractTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OpenInterestStatisticsContractTypeParameter", value)
}

// NewOpenInterestStatisticsContractTypeParameterFromValue returns a pointer to a valid OpenInterestStatisticsContractTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOpenInterestStatisticsContractTypeParameterFromValue(v string) (*OpenInterestStatisticsContractTypeParameter, error) {
	ev := OpenInterestStatisticsContractTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OpenInterestStatisticsContractTypeParameter: valid values are %v", v, AllowedOpenInterestStatisticsContractTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OpenInterestStatisticsContractTypeParameter) IsValid() bool {
	for _, existing := range AllowedOpenInterestStatisticsContractTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to openInterestStatistics_contractType_parameter value
func (v OpenInterestStatisticsContractTypeParameter) Ptr() *OpenInterestStatisticsContractTypeParameter {
	return &v
}

type NullableOpenInterestStatisticsContractTypeParameter struct {
	value *OpenInterestStatisticsContractTypeParameter
	isSet bool
}

func (v NullableOpenInterestStatisticsContractTypeParameter) Get() *OpenInterestStatisticsContractTypeParameter {
	return v.value
}

func (v *NullableOpenInterestStatisticsContractTypeParameter) Set(val *OpenInterestStatisticsContractTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestStatisticsContractTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestStatisticsContractTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestStatisticsContractTypeParameter(val *OpenInterestStatisticsContractTypeParameter) *NullableOpenInterestStatisticsContractTypeParameter {
	return &NullableOpenInterestStatisticsContractTypeParameter{value: val, isSet: true}
}

func (v NullableOpenInterestStatisticsContractTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestStatisticsContractTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
