/*
Spot REST API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TickerSymbolStatusParameter Filters for symbols that have this `tradingStatus`. For a single symbol, a status mismatch returns error `-1220 SYMBOL_DOES_NOT_MATCH_STATUS`. For multiple symbols, non-matching ones are simply excluded from the response.
type TickerSymbolStatusParameter string

// List of ticker_symbolStatus_parameter
const (
	TickerSymbolStatusParameterTrading TickerSymbolStatusParameter = "TRADING"
	TickerSymbolStatusParameterHalt    TickerSymbolStatusParameter = "HALT"
	TickerSymbolStatusParameterBreak   TickerSymbolStatusParameter = "BREAK"
)

// All allowed values of TickerSymbolStatusParameter enum
var AllowedTickerSymbolStatusParameterEnumValues = []TickerSymbolStatusParameter{
	"TRADING",
	"HALT",
	"BREAK",
}

func (v *TickerSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TickerSymbolStatusParameter(value)
	for _, existing := range AllowedTickerSymbolStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TickerSymbolStatusParameter", value)
}

// NewTickerSymbolStatusParameterFromValue returns a pointer to a valid TickerSymbolStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTickerSymbolStatusParameterFromValue(v string) (*TickerSymbolStatusParameter, error) {
	ev := TickerSymbolStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TickerSymbolStatusParameter: valid values are %v", v, AllowedTickerSymbolStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TickerSymbolStatusParameter) IsValid() bool {
	for _, existing := range AllowedTickerSymbolStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ticker_symbolStatus_parameter value
func (v TickerSymbolStatusParameter) Ptr() *TickerSymbolStatusParameter {
	return &v
}

type NullableTickerSymbolStatusParameter struct {
	value *TickerSymbolStatusParameter
	isSet bool
}

func (v NullableTickerSymbolStatusParameter) Get() *TickerSymbolStatusParameter {
	return v.value
}

func (v *NullableTickerSymbolStatusParameter) Set(val *TickerSymbolStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerSymbolStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerSymbolStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerSymbolStatusParameter(val *TickerSymbolStatusParameter) *NullableTickerSymbolStatusParameter {
	return &NullableTickerSymbolStatusParameter{value: val, isSet: true}
}

func (v NullableTickerSymbolStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
