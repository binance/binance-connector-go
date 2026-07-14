/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContractKlineStreamIntervalParameter the model 'ContractKlineStreamIntervalParameter'
type ContractKlineStreamIntervalParameter string

// List of contractKlineStream_interval_parameter
const (
	ContractKlineStreamIntervalParameterInterval1s  ContractKlineStreamIntervalParameter = "1s"
	ContractKlineStreamIntervalParameterInterval1m  ContractKlineStreamIntervalParameter = "1m"
	ContractKlineStreamIntervalParameterInterval5m  ContractKlineStreamIntervalParameter = "5m"
	ContractKlineStreamIntervalParameterInterval15m ContractKlineStreamIntervalParameter = "15m"
	ContractKlineStreamIntervalParameterInterval1h  ContractKlineStreamIntervalParameter = "1h"
	ContractKlineStreamIntervalParameterInterval4h  ContractKlineStreamIntervalParameter = "4h"
	ContractKlineStreamIntervalParameterInterval1d  ContractKlineStreamIntervalParameter = "1d"
)

// All allowed values of ContractKlineStreamIntervalParameter enum
var AllowedContractKlineStreamIntervalParameterEnumValues = []ContractKlineStreamIntervalParameter{
	"1s",
	"1m",
	"5m",
	"15m",
	"1h",
	"4h",
	"1d",
}

func (v *ContractKlineStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContractKlineStreamIntervalParameter(value)
	for _, existing := range AllowedContractKlineStreamIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContractKlineStreamIntervalParameter", value)
}

// NewContractKlineStreamIntervalParameterFromValue returns a pointer to a valid ContractKlineStreamIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContractKlineStreamIntervalParameterFromValue(v string) (*ContractKlineStreamIntervalParameter, error) {
	ev := ContractKlineStreamIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContractKlineStreamIntervalParameter: valid values are %v", v, AllowedContractKlineStreamIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContractKlineStreamIntervalParameter) IsValid() bool {
	for _, existing := range AllowedContractKlineStreamIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to contractKlineStream_interval_parameter value
func (v ContractKlineStreamIntervalParameter) Ptr() *ContractKlineStreamIntervalParameter {
	return &v
}

type NullableContractKlineStreamIntervalParameter struct {
	value *ContractKlineStreamIntervalParameter
	isSet bool
}

func (v NullableContractKlineStreamIntervalParameter) Get() *ContractKlineStreamIntervalParameter {
	return v.value
}

func (v *NullableContractKlineStreamIntervalParameter) Set(val *ContractKlineStreamIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContractKlineStreamIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContractKlineStreamIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractKlineStreamIntervalParameter(val *ContractKlineStreamIntervalParameter) *NullableContractKlineStreamIntervalParameter {
	return &NullableContractKlineStreamIntervalParameter{value: val, isSet: true}
}

func (v NullableContractKlineStreamIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractKlineStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
