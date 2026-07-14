/*
Futures (COIN-M) WebSocket Market Streams

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickStreamsContractTypeParameter the model 'ContinuousContractKlineCandlestickStreamsContractTypeParameter'
type ContinuousContractKlineCandlestickStreamsContractTypeParameter string

// List of continuousContractKlineCandlestickStreams_contractType_parameter
const (
	ContinuousContractKlineCandlestickStreamsContractTypeParameterPerpetual      ContinuousContractKlineCandlestickStreamsContractTypeParameter = "perpetual"
	ContinuousContractKlineCandlestickStreamsContractTypeParameterCurrentQuarter ContinuousContractKlineCandlestickStreamsContractTypeParameter = "current_quarter"
	ContinuousContractKlineCandlestickStreamsContractTypeParameterNextQuarter    ContinuousContractKlineCandlestickStreamsContractTypeParameter = "next_quarter"
)

// All allowed values of ContinuousContractKlineCandlestickStreamsContractTypeParameter enum
var AllowedContinuousContractKlineCandlestickStreamsContractTypeParameterEnumValues = []ContinuousContractKlineCandlestickStreamsContractTypeParameter{
	"perpetual",
	"current_quarter",
	"next_quarter",
}

func (v *ContinuousContractKlineCandlestickStreamsContractTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContinuousContractKlineCandlestickStreamsContractTypeParameter(value)
	for _, existing := range AllowedContinuousContractKlineCandlestickStreamsContractTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContinuousContractKlineCandlestickStreamsContractTypeParameter", value)
}

// NewContinuousContractKlineCandlestickStreamsContractTypeParameterFromValue returns a pointer to a valid ContinuousContractKlineCandlestickStreamsContractTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContinuousContractKlineCandlestickStreamsContractTypeParameterFromValue(v string) (*ContinuousContractKlineCandlestickStreamsContractTypeParameter, error) {
	ev := ContinuousContractKlineCandlestickStreamsContractTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContinuousContractKlineCandlestickStreamsContractTypeParameter: valid values are %v", v, AllowedContinuousContractKlineCandlestickStreamsContractTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContinuousContractKlineCandlestickStreamsContractTypeParameter) IsValid() bool {
	for _, existing := range AllowedContinuousContractKlineCandlestickStreamsContractTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to continuousContractKlineCandlestickStreams_contractType_parameter value
func (v ContinuousContractKlineCandlestickStreamsContractTypeParameter) Ptr() *ContinuousContractKlineCandlestickStreamsContractTypeParameter {
	return &v
}

type NullableContinuousContractKlineCandlestickStreamsContractTypeParameter struct {
	value *ContinuousContractKlineCandlestickStreamsContractTypeParameter
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) Get() *ContinuousContractKlineCandlestickStreamsContractTypeParameter {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) Set(val *ContinuousContractKlineCandlestickStreamsContractTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickStreamsContractTypeParameter(val *ContinuousContractKlineCandlestickStreamsContractTypeParameter) *NullableContinuousContractKlineCandlestickStreamsContractTypeParameter {
	return &NullableContinuousContractKlineCandlestickStreamsContractTypeParameter{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickStreamsContractTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
