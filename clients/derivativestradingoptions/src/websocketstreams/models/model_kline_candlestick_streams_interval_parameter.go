/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineCandlestickStreamsIntervalParameter the model 'KlineCandlestickStreamsIntervalParameter'
type KlineCandlestickStreamsIntervalParameter string

// List of klineCandlestickStreams_interval_parameter
const (
	KlineCandlestickStreamsIntervalParameterInterval1m  KlineCandlestickStreamsIntervalParameter = "1m"
	KlineCandlestickStreamsIntervalParameterInterval3m  KlineCandlestickStreamsIntervalParameter = "3m"
	KlineCandlestickStreamsIntervalParameterInterval5m  KlineCandlestickStreamsIntervalParameter = "5m"
	KlineCandlestickStreamsIntervalParameterInterval15m KlineCandlestickStreamsIntervalParameter = "15m"
	KlineCandlestickStreamsIntervalParameterInterval30m KlineCandlestickStreamsIntervalParameter = "30m"
	KlineCandlestickStreamsIntervalParameterInterval1h  KlineCandlestickStreamsIntervalParameter = "1h"
	KlineCandlestickStreamsIntervalParameterInterval2h  KlineCandlestickStreamsIntervalParameter = "2h"
	KlineCandlestickStreamsIntervalParameterInterval4h  KlineCandlestickStreamsIntervalParameter = "4h"
	KlineCandlestickStreamsIntervalParameterInterval6h  KlineCandlestickStreamsIntervalParameter = "6h"
	KlineCandlestickStreamsIntervalParameterInterval12h KlineCandlestickStreamsIntervalParameter = "12h"
	KlineCandlestickStreamsIntervalParameterInterval1d  KlineCandlestickStreamsIntervalParameter = "1d"
	KlineCandlestickStreamsIntervalParameterInterval3d  KlineCandlestickStreamsIntervalParameter = "3d"
	KlineCandlestickStreamsIntervalParameterInterval1w  KlineCandlestickStreamsIntervalParameter = "1w"
)

// All allowed values of KlineCandlestickStreamsIntervalParameter enum
var AllowedKlineCandlestickStreamsIntervalParameterEnumValues = []KlineCandlestickStreamsIntervalParameter{
	"1m",
	"3m",
	"5m",
	"15m",
	"30m",
	"1h",
	"2h",
	"4h",
	"6h",
	"12h",
	"1d",
	"3d",
	"1w",
}

func (v *KlineCandlestickStreamsIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KlineCandlestickStreamsIntervalParameter(value)
	for _, existing := range AllowedKlineCandlestickStreamsIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KlineCandlestickStreamsIntervalParameter", value)
}

// NewKlineCandlestickStreamsIntervalParameterFromValue returns a pointer to a valid KlineCandlestickStreamsIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKlineCandlestickStreamsIntervalParameterFromValue(v string) (*KlineCandlestickStreamsIntervalParameter, error) {
	ev := KlineCandlestickStreamsIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KlineCandlestickStreamsIntervalParameter: valid values are %v", v, AllowedKlineCandlestickStreamsIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KlineCandlestickStreamsIntervalParameter) IsValid() bool {
	for _, existing := range AllowedKlineCandlestickStreamsIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to klineCandlestickStreams_interval_parameter value
func (v KlineCandlestickStreamsIntervalParameter) Ptr() *KlineCandlestickStreamsIntervalParameter {
	return &v
}

type NullableKlineCandlestickStreamsIntervalParameter struct {
	value *KlineCandlestickStreamsIntervalParameter
	isSet bool
}

func (v NullableKlineCandlestickStreamsIntervalParameter) Get() *KlineCandlestickStreamsIntervalParameter {
	return v.value
}

func (v *NullableKlineCandlestickStreamsIntervalParameter) Set(val *KlineCandlestickStreamsIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickStreamsIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickStreamsIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickStreamsIntervalParameter(val *KlineCandlestickStreamsIntervalParameter) *NullableKlineCandlestickStreamsIntervalParameter {
	return &NullableKlineCandlestickStreamsIntervalParameter{value: val, isSet: true}
}

func (v NullableKlineCandlestickStreamsIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickStreamsIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
