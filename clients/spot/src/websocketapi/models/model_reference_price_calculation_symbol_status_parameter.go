/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ReferencePriceCalculationSymbolStatusParameter the model 'ReferencePriceCalculationSymbolStatusParameter'
type ReferencePriceCalculationSymbolStatusParameter string

// List of referencePriceCalculation_symbolStatus_parameter
const (
	ReferencePriceCalculationSymbolStatusParameterTrading ReferencePriceCalculationSymbolStatusParameter = "TRADING"
	ReferencePriceCalculationSymbolStatusParameterHalt    ReferencePriceCalculationSymbolStatusParameter = "HALT"
	ReferencePriceCalculationSymbolStatusParameterBreak   ReferencePriceCalculationSymbolStatusParameter = "BREAK"
)

// All allowed values of ReferencePriceCalculationSymbolStatusParameter enum
var AllowedReferencePriceCalculationSymbolStatusParameterEnumValues = []ReferencePriceCalculationSymbolStatusParameter{
	"TRADING",
	"HALT",
	"BREAK",
}

func (v *ReferencePriceCalculationSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReferencePriceCalculationSymbolStatusParameter(value)
	for _, existing := range AllowedReferencePriceCalculationSymbolStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReferencePriceCalculationSymbolStatusParameter", value)
}

// NewReferencePriceCalculationSymbolStatusParameterFromValue returns a pointer to a valid ReferencePriceCalculationSymbolStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReferencePriceCalculationSymbolStatusParameterFromValue(v string) (*ReferencePriceCalculationSymbolStatusParameter, error) {
	ev := ReferencePriceCalculationSymbolStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReferencePriceCalculationSymbolStatusParameter: valid values are %v", v, AllowedReferencePriceCalculationSymbolStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReferencePriceCalculationSymbolStatusParameter) IsValid() bool {
	for _, existing := range AllowedReferencePriceCalculationSymbolStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to referencePriceCalculation_symbolStatus_parameter value
func (v ReferencePriceCalculationSymbolStatusParameter) Ptr() *ReferencePriceCalculationSymbolStatusParameter {
	return &v
}

type NullableReferencePriceCalculationSymbolStatusParameter struct {
	value *ReferencePriceCalculationSymbolStatusParameter
	isSet bool
}

func (v NullableReferencePriceCalculationSymbolStatusParameter) Get() *ReferencePriceCalculationSymbolStatusParameter {
	return v.value
}

func (v *NullableReferencePriceCalculationSymbolStatusParameter) Set(val *ReferencePriceCalculationSymbolStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableReferencePriceCalculationSymbolStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableReferencePriceCalculationSymbolStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferencePriceCalculationSymbolStatusParameter(val *ReferencePriceCalculationSymbolStatusParameter) *NullableReferencePriceCalculationSymbolStatusParameter {
	return &NullableReferencePriceCalculationSymbolStatusParameter{value: val, isSet: true}
}

func (v NullableReferencePriceCalculationSymbolStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferencePriceCalculationSymbolStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
