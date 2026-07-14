/*
C2C REST API

Query fiat transaction history via the C2C REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetC2CTradeHistoryTradeTypeParameter the model 'GetC2CTradeHistoryTradeTypeParameter'
type GetC2CTradeHistoryTradeTypeParameter string

// List of getC2CTradeHistory_tradeType_parameter
const (
	GetC2CTradeHistoryTradeTypeParameterBuy  GetC2CTradeHistoryTradeTypeParameter = "BUY"
	GetC2CTradeHistoryTradeTypeParameterSell GetC2CTradeHistoryTradeTypeParameter = "SELL"
)

// All allowed values of GetC2CTradeHistoryTradeTypeParameter enum
var AllowedGetC2CTradeHistoryTradeTypeParameterEnumValues = []GetC2CTradeHistoryTradeTypeParameter{
	"BUY",
	"SELL",
}

func (v *GetC2CTradeHistoryTradeTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetC2CTradeHistoryTradeTypeParameter(value)
	for _, existing := range AllowedGetC2CTradeHistoryTradeTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetC2CTradeHistoryTradeTypeParameter", value)
}

// NewGetC2CTradeHistoryTradeTypeParameterFromValue returns a pointer to a valid GetC2CTradeHistoryTradeTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetC2CTradeHistoryTradeTypeParameterFromValue(v string) (*GetC2CTradeHistoryTradeTypeParameter, error) {
	ev := GetC2CTradeHistoryTradeTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetC2CTradeHistoryTradeTypeParameter: valid values are %v", v, AllowedGetC2CTradeHistoryTradeTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetC2CTradeHistoryTradeTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetC2CTradeHistoryTradeTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getC2CTradeHistory_tradeType_parameter value
func (v GetC2CTradeHistoryTradeTypeParameter) Ptr() *GetC2CTradeHistoryTradeTypeParameter {
	return &v
}

type NullableGetC2CTradeHistoryTradeTypeParameter struct {
	value *GetC2CTradeHistoryTradeTypeParameter
	isSet bool
}

func (v NullableGetC2CTradeHistoryTradeTypeParameter) Get() *GetC2CTradeHistoryTradeTypeParameter {
	return v.value
}

func (v *NullableGetC2CTradeHistoryTradeTypeParameter) Set(val *GetC2CTradeHistoryTradeTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetC2CTradeHistoryTradeTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetC2CTradeHistoryTradeTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetC2CTradeHistoryTradeTypeParameter(val *GetC2CTradeHistoryTradeTypeParameter) *NullableGetC2CTradeHistoryTradeTypeParameter {
	return &NullableGetC2CTradeHistoryTradeTypeParameter{value: val, isSet: true}
}

func (v NullableGetC2CTradeHistoryTradeTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetC2CTradeHistoryTradeTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
