/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// IndexPriceKlineCandlestickDataIntervalParameter the model 'IndexPriceKlineCandlestickDataIntervalParameter'
type IndexPriceKlineCandlestickDataIntervalParameter string

// List of indexPriceKlineCandlestickData_interval_parameter
const (
	IndexPriceKlineCandlestickDataIntervalParameterInterval1m  IndexPriceKlineCandlestickDataIntervalParameter = "1m"
	IndexPriceKlineCandlestickDataIntervalParameterInterval3m  IndexPriceKlineCandlestickDataIntervalParameter = "3m"
	IndexPriceKlineCandlestickDataIntervalParameterInterval5m  IndexPriceKlineCandlestickDataIntervalParameter = "5m"
	IndexPriceKlineCandlestickDataIntervalParameterInterval15m IndexPriceKlineCandlestickDataIntervalParameter = "15m"
	IndexPriceKlineCandlestickDataIntervalParameterInterval30m IndexPriceKlineCandlestickDataIntervalParameter = "30m"
	IndexPriceKlineCandlestickDataIntervalParameterInterval1h  IndexPriceKlineCandlestickDataIntervalParameter = "1h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval2h  IndexPriceKlineCandlestickDataIntervalParameter = "2h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval4h  IndexPriceKlineCandlestickDataIntervalParameter = "4h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval6h  IndexPriceKlineCandlestickDataIntervalParameter = "6h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval8h  IndexPriceKlineCandlestickDataIntervalParameter = "8h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval12h IndexPriceKlineCandlestickDataIntervalParameter = "12h"
	IndexPriceKlineCandlestickDataIntervalParameterInterval1d  IndexPriceKlineCandlestickDataIntervalParameter = "1d"
	IndexPriceKlineCandlestickDataIntervalParameterInterval3d  IndexPriceKlineCandlestickDataIntervalParameter = "3d"
	IndexPriceKlineCandlestickDataIntervalParameterInterval1w  IndexPriceKlineCandlestickDataIntervalParameter = "1w"
	IndexPriceKlineCandlestickDataIntervalParameterInterval1M  IndexPriceKlineCandlestickDataIntervalParameter = "1M"
)

// All allowed values of IndexPriceKlineCandlestickDataIntervalParameter enum
var AllowedIndexPriceKlineCandlestickDataIntervalParameterEnumValues = []IndexPriceKlineCandlestickDataIntervalParameter{
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

func (v *IndexPriceKlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IndexPriceKlineCandlestickDataIntervalParameter(value)
	for _, existing := range AllowedIndexPriceKlineCandlestickDataIntervalParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IndexPriceKlineCandlestickDataIntervalParameter", value)
}

// NewIndexPriceKlineCandlestickDataIntervalParameterFromValue returns a pointer to a valid IndexPriceKlineCandlestickDataIntervalParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndexPriceKlineCandlestickDataIntervalParameterFromValue(v string) (*IndexPriceKlineCandlestickDataIntervalParameter, error) {
	ev := IndexPriceKlineCandlestickDataIntervalParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IndexPriceKlineCandlestickDataIntervalParameter: valid values are %v", v, AllowedIndexPriceKlineCandlestickDataIntervalParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IndexPriceKlineCandlestickDataIntervalParameter) IsValid() bool {
	for _, existing := range AllowedIndexPriceKlineCandlestickDataIntervalParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to indexPriceKlineCandlestickData_interval_parameter value
func (v IndexPriceKlineCandlestickDataIntervalParameter) Ptr() *IndexPriceKlineCandlestickDataIntervalParameter {
	return &v
}

type NullableIndexPriceKlineCandlestickDataIntervalParameter struct {
	value *IndexPriceKlineCandlestickDataIntervalParameter
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataIntervalParameter) Get() *IndexPriceKlineCandlestickDataIntervalParameter {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataIntervalParameter) Set(val *IndexPriceKlineCandlestickDataIntervalParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataIntervalParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataIntervalParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataIntervalParameter(val *IndexPriceKlineCandlestickDataIntervalParameter) *NullableIndexPriceKlineCandlestickDataIntervalParameter {
	return &NullableIndexPriceKlineCandlestickDataIntervalParameter{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataIntervalParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataIntervalParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
