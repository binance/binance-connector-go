/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineStreamIntervalParameter the model 'KlineStreamIntervalParameter'
type KlineStreamIntervalParameter string

// List of klineStream_interval_parameter
const (
	KlineStreamIntervalParameterInterval1m  KlineStreamIntervalParameter = "1m"
	KlineStreamIntervalParameterInterval3m  KlineStreamIntervalParameter = "3m"
	KlineStreamIntervalParameterInterval5m  KlineStreamIntervalParameter = "5m"
	KlineStreamIntervalParameterInterval15m KlineStreamIntervalParameter = "15m"
	KlineStreamIntervalParameterInterval30m KlineStreamIntervalParameter = "30m"
	KlineStreamIntervalParameterInterval1h  KlineStreamIntervalParameter = "1h"
	KlineStreamIntervalParameterInterval2h  KlineStreamIntervalParameter = "2h"
	KlineStreamIntervalParameterInterval4h  KlineStreamIntervalParameter = "4h"
	KlineStreamIntervalParameterInterval6h  KlineStreamIntervalParameter = "6h"
	KlineStreamIntervalParameterInterval8h  KlineStreamIntervalParameter = "8h"
	KlineStreamIntervalParameterInterval12h KlineStreamIntervalParameter = "12h"
	KlineStreamIntervalParameterInterval1d  KlineStreamIntervalParameter = "1d"
	KlineStreamIntervalParameterInterval3d  KlineStreamIntervalParameter = "3d"
	KlineStreamIntervalParameterInterval1w  KlineStreamIntervalParameter = "1w"
	KlineStreamIntervalParameterInterval1M  KlineStreamIntervalParameter = "1M"
)

// All allowed values of KlineStreamIntervalParameter enum
var AllowedKlineStreamIntervalParameterEnumValues = []KlineStreamIntervalParameter{
	"1m",
	"3m",
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"4h",
	"6h",
	"8h",
	"12h",
	"1d",
	"3d",
	"1w",
	"1M",
}

func (v *KlineStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KlineStreamIntervalParameter(value)
	for _, existing := range AllowedKlineStreamIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KlineStreamIntervalParameter", value)
}

// NewKlineStreamIntervalParameterFromValue returns a pointer to a valid KlineStreamIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKlineStreamIntervalParameterFromValue(v string) (*KlineStreamIntervalParameter, error) {
	ev := KlineStreamIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KlineStreamIntervalParameter: valid values are %v", v, AllowedKlineStreamIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KlineStreamIntervalParameter) IsValid() bool {
	for _, existing := range AllowedKlineStreamIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to klineStream_interval_parameter value
func (v KlineStreamIntervalParameter) Ptr() *KlineStreamIntervalParameter {
	return &v
}

type NullableKlineStreamIntervalParameter struct {
	value *KlineStreamIntervalParameter
	isSet bool
}

func (v NullableKlineStreamIntervalParameter) Get() *KlineStreamIntervalParameter {
	return v.value
}

func (v *NullableKlineStreamIntervalParameter) Set(val *KlineStreamIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineStreamIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineStreamIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineStreamIntervalParameter(val *KlineStreamIntervalParameter) *NullableKlineStreamIntervalParameter {
	return &NullableKlineStreamIntervalParameter{value: val, isSet: true}
}

func (v NullableKlineStreamIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineStreamIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
