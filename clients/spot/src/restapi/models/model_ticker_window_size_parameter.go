/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerWindowSizeParameter the model 'TickerWindowSizeParameter'
type TickerWindowSizeParameter string

// List of ticker_windowSize_parameter
const (
	TickerWindowSizeParameterWindowSize1m  TickerWindowSizeParameter = "1m"
	TickerWindowSizeParameterWindowSize2m  TickerWindowSizeParameter = "2m"
	TickerWindowSizeParameterWindowSize3m  TickerWindowSizeParameter = "3m"
	TickerWindowSizeParameterWindowSize4m  TickerWindowSizeParameter = "4m"
	TickerWindowSizeParameterWindowSize5m  TickerWindowSizeParameter = "5m"
	TickerWindowSizeParameterWindowSize6m  TickerWindowSizeParameter = "6m"
	TickerWindowSizeParameterWindowSize7m  TickerWindowSizeParameter = "7m"
	TickerWindowSizeParameterWindowSize8m  TickerWindowSizeParameter = "8m"
	TickerWindowSizeParameterWindowSize9m  TickerWindowSizeParameter = "9m"
	TickerWindowSizeParameterWindowSize10m TickerWindowSizeParameter = "10m"
	TickerWindowSizeParameterWindowSize11m TickerWindowSizeParameter = "11m"
	TickerWindowSizeParameterWindowSize12m TickerWindowSizeParameter = "12m"
	TickerWindowSizeParameterWindowSize13m TickerWindowSizeParameter = "13m"
	TickerWindowSizeParameterWindowSize14m TickerWindowSizeParameter = "14m"
	TickerWindowSizeParameterWindowSize15m TickerWindowSizeParameter = "15m"
	TickerWindowSizeParameterWindowSize16m TickerWindowSizeParameter = "16m"
	TickerWindowSizeParameterWindowSize17m TickerWindowSizeParameter = "17m"
	TickerWindowSizeParameterWindowSize18m TickerWindowSizeParameter = "18m"
	TickerWindowSizeParameterWindowSize19m TickerWindowSizeParameter = "19m"
	TickerWindowSizeParameterWindowSize20m TickerWindowSizeParameter = "20m"
	TickerWindowSizeParameterWindowSize21m TickerWindowSizeParameter = "21m"
	TickerWindowSizeParameterWindowSize22m TickerWindowSizeParameter = "22m"
	TickerWindowSizeParameterWindowSize23m TickerWindowSizeParameter = "23m"
	TickerWindowSizeParameterWindowSize24m TickerWindowSizeParameter = "24m"
	TickerWindowSizeParameterWindowSize25m TickerWindowSizeParameter = "25m"
	TickerWindowSizeParameterWindowSize26m TickerWindowSizeParameter = "26m"
	TickerWindowSizeParameterWindowSize27m TickerWindowSizeParameter = "27m"
	TickerWindowSizeParameterWindowSize28m TickerWindowSizeParameter = "28m"
	TickerWindowSizeParameterWindowSize29m TickerWindowSizeParameter = "29m"
	TickerWindowSizeParameterWindowSize30m TickerWindowSizeParameter = "30m"
	TickerWindowSizeParameterWindowSize31m TickerWindowSizeParameter = "31m"
	TickerWindowSizeParameterWindowSize32m TickerWindowSizeParameter = "32m"
	TickerWindowSizeParameterWindowSize33m TickerWindowSizeParameter = "33m"
	TickerWindowSizeParameterWindowSize34m TickerWindowSizeParameter = "34m"
	TickerWindowSizeParameterWindowSize35m TickerWindowSizeParameter = "35m"
	TickerWindowSizeParameterWindowSize36m TickerWindowSizeParameter = "36m"
	TickerWindowSizeParameterWindowSize37m TickerWindowSizeParameter = "37m"
	TickerWindowSizeParameterWindowSize38m TickerWindowSizeParameter = "38m"
	TickerWindowSizeParameterWindowSize39m TickerWindowSizeParameter = "39m"
	TickerWindowSizeParameterWindowSize40m TickerWindowSizeParameter = "40m"
	TickerWindowSizeParameterWindowSize41m TickerWindowSizeParameter = "41m"
	TickerWindowSizeParameterWindowSize42m TickerWindowSizeParameter = "42m"
	TickerWindowSizeParameterWindowSize43m TickerWindowSizeParameter = "43m"
	TickerWindowSizeParameterWindowSize44m TickerWindowSizeParameter = "44m"
	TickerWindowSizeParameterWindowSize45m TickerWindowSizeParameter = "45m"
	TickerWindowSizeParameterWindowSize46m TickerWindowSizeParameter = "46m"
	TickerWindowSizeParameterWindowSize47m TickerWindowSizeParameter = "47m"
	TickerWindowSizeParameterWindowSize48m TickerWindowSizeParameter = "48m"
	TickerWindowSizeParameterWindowSize49m TickerWindowSizeParameter = "49m"
	TickerWindowSizeParameterWindowSize50m TickerWindowSizeParameter = "50m"
	TickerWindowSizeParameterWindowSize51m TickerWindowSizeParameter = "51m"
	TickerWindowSizeParameterWindowSize52m TickerWindowSizeParameter = "52m"
	TickerWindowSizeParameterWindowSize53m TickerWindowSizeParameter = "53m"
	TickerWindowSizeParameterWindowSize54m TickerWindowSizeParameter = "54m"
	TickerWindowSizeParameterWindowSize55m TickerWindowSizeParameter = "55m"
	TickerWindowSizeParameterWindowSize56m TickerWindowSizeParameter = "56m"
	TickerWindowSizeParameterWindowSize57m TickerWindowSizeParameter = "57m"
	TickerWindowSizeParameterWindowSize58m TickerWindowSizeParameter = "58m"
	TickerWindowSizeParameterWindowSize59m TickerWindowSizeParameter = "59m"
	TickerWindowSizeParameterWindowSize1h  TickerWindowSizeParameter = "1h"
	TickerWindowSizeParameterWindowSize2h  TickerWindowSizeParameter = "2h"
	TickerWindowSizeParameterWindowSize3h  TickerWindowSizeParameter = "3h"
	TickerWindowSizeParameterWindowSize4h  TickerWindowSizeParameter = "4h"
	TickerWindowSizeParameterWindowSize5h  TickerWindowSizeParameter = "5h"
	TickerWindowSizeParameterWindowSize6h  TickerWindowSizeParameter = "6h"
	TickerWindowSizeParameterWindowSize7h  TickerWindowSizeParameter = "7h"
	TickerWindowSizeParameterWindowSize8h  TickerWindowSizeParameter = "8h"
	TickerWindowSizeParameterWindowSize9h  TickerWindowSizeParameter = "9h"
	TickerWindowSizeParameterWindowSize10h TickerWindowSizeParameter = "10h"
	TickerWindowSizeParameterWindowSize11h TickerWindowSizeParameter = "11h"
	TickerWindowSizeParameterWindowSize12h TickerWindowSizeParameter = "12h"
	TickerWindowSizeParameterWindowSize13h TickerWindowSizeParameter = "13h"
	TickerWindowSizeParameterWindowSize14h TickerWindowSizeParameter = "14h"
	TickerWindowSizeParameterWindowSize15h TickerWindowSizeParameter = "15h"
	TickerWindowSizeParameterWindowSize16h TickerWindowSizeParameter = "16h"
	TickerWindowSizeParameterWindowSize17h TickerWindowSizeParameter = "17h"
	TickerWindowSizeParameterWindowSize18h TickerWindowSizeParameter = "18h"
	TickerWindowSizeParameterWindowSize19h TickerWindowSizeParameter = "19h"
	TickerWindowSizeParameterWindowSize20h TickerWindowSizeParameter = "20h"
	TickerWindowSizeParameterWindowSize21h TickerWindowSizeParameter = "21h"
	TickerWindowSizeParameterWindowSize22h TickerWindowSizeParameter = "22h"
	TickerWindowSizeParameterWindowSize23h TickerWindowSizeParameter = "23h"
	TickerWindowSizeParameterWindowSize1d  TickerWindowSizeParameter = "1d"
	TickerWindowSizeParameterWindowSize2d  TickerWindowSizeParameter = "2d"
	TickerWindowSizeParameterWindowSize3d  TickerWindowSizeParameter = "3d"
	TickerWindowSizeParameterWindowSize4d  TickerWindowSizeParameter = "4d"
	TickerWindowSizeParameterWindowSize5d  TickerWindowSizeParameter = "5d"
	TickerWindowSizeParameterWindowSize6d  TickerWindowSizeParameter = "6d"
)

// All allowed values of TickerWindowSizeParameter enum
var AllowedTickerWindowSizeParameterEnumValues = []TickerWindowSizeParameter{
	"1m",
	"2m",
	"3m",
	"4m",
	"5m",
	"6m",
	"7m",
	"8m",
	"9m",
	"10m",
	"11m",
	"12m",
	"13m",
	"14m",
	"15m",
	"16m",
	"17m",
	"18m",
	"19m",
	"20m",
	"21m",
	"22m",
	"23m",
	"24m",
	"25m",
	"26m",
	"27m",
	"28m",
	"29m",
	"30m",
	"31m",
	"32m",
	"33m",
	"34m",
	"35m",
	"36m",
	"37m",
	"38m",
	"39m",
	"40m",
	"41m",
	"42m",
	"43m",
	"44m",
	"45m",
	"46m",
	"47m",
	"48m",
	"49m",
	"50m",
	"51m",
	"52m",
	"53m",
	"54m",
	"55m",
	"56m",
	"57m",
	"58m",
	"59m",
	"1h",
	"2h",
	"3h",
	"4h",
	"5h",
	"6h",
	"7h",
	"8h",
	"9h",
	"10h",
	"11h",
	"12h",
	"13h",
	"14h",
	"15h",
	"16h",
	"17h",
	"18h",
	"19h",
	"20h",
	"21h",
	"22h",
	"23h",
	"1d",
	"2d",
	"3d",
	"4d",
	"5d",
	"6d",
}

func (v *TickerWindowSizeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TickerWindowSizeParameter(value)
	for _, existing := range AllowedTickerWindowSizeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TickerWindowSizeParameter", value)
}

// NewTickerWindowSizeParameterFromValue returns a pointer to a valid TickerWindowSizeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTickerWindowSizeParameterFromValue(v string) (*TickerWindowSizeParameter, error) {
	ev := TickerWindowSizeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TickerWindowSizeParameter: valid values are %v", v, AllowedTickerWindowSizeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TickerWindowSizeParameter) IsValid() bool {
	for _, existing := range AllowedTickerWindowSizeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ticker_windowSize_parameter value
func (v TickerWindowSizeParameter) Ptr() *TickerWindowSizeParameter {
	return &v
}

type NullableTickerWindowSizeParameter struct {
	value *TickerWindowSizeParameter
	isSet bool
}

func (v NullableTickerWindowSizeParameter) Get() *TickerWindowSizeParameter {
	return v.value
}

func (v *NullableTickerWindowSizeParameter) Set(val *TickerWindowSizeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerWindowSizeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerWindowSizeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerWindowSizeParameter(val *TickerWindowSizeParameter) *NullableTickerWindowSizeParameter {
	return &NullableTickerWindowSizeParameter{value: val, isSet: true}
}

func (v NullableTickerWindowSizeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerWindowSizeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
