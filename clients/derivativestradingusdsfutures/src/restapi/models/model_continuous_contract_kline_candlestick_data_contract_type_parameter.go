/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// ContinuousContractKlineCandlestickDataContractTypeParameter the model 'ContinuousContractKlineCandlestickDataContractTypeParameter'
type ContinuousContractKlineCandlestickDataContractTypeParameter string

// List of continuousContractKlineCandlestickData_contractType_parameter
const (
	ContinuousContractKlineCandlestickDataContractTypeParameterPerpetual        ContinuousContractKlineCandlestickDataContractTypeParameter = "PERPETUAL"
	ContinuousContractKlineCandlestickDataContractTypeParameterCurrentQuarter   ContinuousContractKlineCandlestickDataContractTypeParameter = "CURRENT_QUARTER"
	ContinuousContractKlineCandlestickDataContractTypeParameterNextQuarter      ContinuousContractKlineCandlestickDataContractTypeParameter = "NEXT_QUARTER"
	ContinuousContractKlineCandlestickDataContractTypeParameterTradifiPerpetual ContinuousContractKlineCandlestickDataContractTypeParameter = "TRADIFI_PERPETUAL"
)

// All allowed values of ContinuousContractKlineCandlestickDataContractTypeParameter enum
var AllowedContinuousContractKlineCandlestickDataContractTypeParameterEnumValues = []ContinuousContractKlineCandlestickDataContractTypeParameter{
	"PERPETUAL",
	"CURRENT_QUARTER",
	"NEXT_QUARTER",
	"TRADIFI_PERPETUAL",
}

func (v *ContinuousContractKlineCandlestickDataContractTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContinuousContractKlineCandlestickDataContractTypeParameter(value)
	for _, existing := range AllowedContinuousContractKlineCandlestickDataContractTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContinuousContractKlineCandlestickDataContractTypeParameter", value)
}

// NewContinuousContractKlineCandlestickDataContractTypeParameterFromValue returns a pointer to a valid ContinuousContractKlineCandlestickDataContractTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContinuousContractKlineCandlestickDataContractTypeParameterFromValue(v string) (*ContinuousContractKlineCandlestickDataContractTypeParameter, error) {
	ev := ContinuousContractKlineCandlestickDataContractTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContinuousContractKlineCandlestickDataContractTypeParameter: valid values are %v", v, AllowedContinuousContractKlineCandlestickDataContractTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContinuousContractKlineCandlestickDataContractTypeParameter) IsValid() bool {
	for _, existing := range AllowedContinuousContractKlineCandlestickDataContractTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to continuousContractKlineCandlestickData_contractType_parameter value
func (v ContinuousContractKlineCandlestickDataContractTypeParameter) Ptr() *ContinuousContractKlineCandlestickDataContractTypeParameter {
	return &v
}

type NullableContinuousContractKlineCandlestickDataContractTypeParameter struct {
	value *ContinuousContractKlineCandlestickDataContractTypeParameter
	isSet bool
}

func (v NullableContinuousContractKlineCandlestickDataContractTypeParameter) Get() *ContinuousContractKlineCandlestickDataContractTypeParameter {
	return v.value
}

func (v *NullableContinuousContractKlineCandlestickDataContractTypeParameter) Set(val *ContinuousContractKlineCandlestickDataContractTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableContinuousContractKlineCandlestickDataContractTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableContinuousContractKlineCandlestickDataContractTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinuousContractKlineCandlestickDataContractTypeParameter(val *ContinuousContractKlineCandlestickDataContractTypeParameter) *NullableContinuousContractKlineCandlestickDataContractTypeParameter {
	return &NullableContinuousContractKlineCandlestickDataContractTypeParameter{value: val, isSet: true}
}

func (v NullableContinuousContractKlineCandlestickDataContractTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinuousContractKlineCandlestickDataContractTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
