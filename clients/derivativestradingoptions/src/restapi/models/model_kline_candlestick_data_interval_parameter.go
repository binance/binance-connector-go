/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// KlineCandlestickDataIntervalParameter the model 'KlineCandlestickDataIntervalParameter'
type KlineCandlestickDataIntervalParameter string

// List of klineCandlestickData_interval_parameter
const (
	KlineCandlestickDataIntervalParameterInterval1m  KlineCandlestickDataIntervalParameter = "1m"
	KlineCandlestickDataIntervalParameterInterval3m  KlineCandlestickDataIntervalParameter = "3m"
	KlineCandlestickDataIntervalParameterInterval5m  KlineCandlestickDataIntervalParameter = "5m"
	KlineCandlestickDataIntervalParameterInterval15m KlineCandlestickDataIntervalParameter = "15m"
	KlineCandlestickDataIntervalParameterInterval30m KlineCandlestickDataIntervalParameter = "30m"
	KlineCandlestickDataIntervalParameterInterval1h  KlineCandlestickDataIntervalParameter = "1h"
	KlineCandlestickDataIntervalParameterInterval2h  KlineCandlestickDataIntervalParameter = "2h"
	KlineCandlestickDataIntervalParameterInterval4h  KlineCandlestickDataIntervalParameter = "4h"
	KlineCandlestickDataIntervalParameterInterval6h  KlineCandlestickDataIntervalParameter = "6h"
	KlineCandlestickDataIntervalParameterInterval8h  KlineCandlestickDataIntervalParameter = "8h"
	KlineCandlestickDataIntervalParameterInterval12h KlineCandlestickDataIntervalParameter = "12h"
	KlineCandlestickDataIntervalParameterInterval1d  KlineCandlestickDataIntervalParameter = "1d"
	KlineCandlestickDataIntervalParameterInterval3d  KlineCandlestickDataIntervalParameter = "3d"
	KlineCandlestickDataIntervalParameterInterval1w  KlineCandlestickDataIntervalParameter = "1w"
	KlineCandlestickDataIntervalParameterInterval1M  KlineCandlestickDataIntervalParameter = "1M"
)

// All allowed values of KlineCandlestickDataIntervalParameter enum
var AllowedKlineCandlestickDataIntervalParameterEnumValues = []KlineCandlestickDataIntervalParameter{
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

func (v *KlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KlineCandlestickDataIntervalParameter(value)
	for _, existing := range AllowedKlineCandlestickDataIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KlineCandlestickDataIntervalParameter", value)
}

// NewKlineCandlestickDataIntervalParameterFromValue returns a pointer to a valid KlineCandlestickDataIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKlineCandlestickDataIntervalParameterFromValue(v string) (*KlineCandlestickDataIntervalParameter, error) {
	ev := KlineCandlestickDataIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KlineCandlestickDataIntervalParameter: valid values are %v", v, AllowedKlineCandlestickDataIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KlineCandlestickDataIntervalParameter) IsValid() bool {
	for _, existing := range AllowedKlineCandlestickDataIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to klineCandlestickData_interval_parameter value
func (v KlineCandlestickDataIntervalParameter) Ptr() *KlineCandlestickDataIntervalParameter {
	return &v
}

type NullableKlineCandlestickDataIntervalParameter struct {
	value *KlineCandlestickDataIntervalParameter
	isSet bool
}

func (v NullableKlineCandlestickDataIntervalParameter) Get() *KlineCandlestickDataIntervalParameter {
	return v.value
}

func (v *NullableKlineCandlestickDataIntervalParameter) Set(val *KlineCandlestickDataIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickDataIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickDataIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickDataIntervalParameter(val *KlineCandlestickDataIntervalParameter) *NullableKlineCandlestickDataIntervalParameter {
	return &NullableKlineCandlestickDataIntervalParameter{value: val, isSet: true}
}

func (v NullableKlineCandlestickDataIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
