/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TakerBuySellVolumeContractTypeParameter the model 'TakerBuySellVolumeContractTypeParameter'
type TakerBuySellVolumeContractTypeParameter string

// List of takerBuySellVolume_contractType_parameter
const (
	TakerBuySellVolumeContractTypeParameterAll            TakerBuySellVolumeContractTypeParameter = "ALL"
	TakerBuySellVolumeContractTypeParameterPerpetual      TakerBuySellVolumeContractTypeParameter = "PERPETUAL"
	TakerBuySellVolumeContractTypeParameterCurrentQuarter TakerBuySellVolumeContractTypeParameter = "CURRENT_QUARTER"
	TakerBuySellVolumeContractTypeParameterNextQuarter    TakerBuySellVolumeContractTypeParameter = "NEXT_QUARTER"
)

// All allowed values of TakerBuySellVolumeContractTypeParameter enum
var AllowedTakerBuySellVolumeContractTypeParameterEnumValues = []TakerBuySellVolumeContractTypeParameter{
	"ALL",
	"PERPETUAL",
	"CURRENT_QUARTER",
	"NEXT_QUARTER",
}

func (v *TakerBuySellVolumeContractTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TakerBuySellVolumeContractTypeParameter(value)
	for _, existing := range AllowedTakerBuySellVolumeContractTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TakerBuySellVolumeContractTypeParameter", value)
}

// NewTakerBuySellVolumeContractTypeParameterFromValue returns a pointer to a valid TakerBuySellVolumeContractTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTakerBuySellVolumeContractTypeParameterFromValue(v string) (*TakerBuySellVolumeContractTypeParameter, error) {
	ev := TakerBuySellVolumeContractTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TakerBuySellVolumeContractTypeParameter: valid values are %v", v, AllowedTakerBuySellVolumeContractTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TakerBuySellVolumeContractTypeParameter) IsValid() bool {
	for _, existing := range AllowedTakerBuySellVolumeContractTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to takerBuySellVolume_contractType_parameter value
func (v TakerBuySellVolumeContractTypeParameter) Ptr() *TakerBuySellVolumeContractTypeParameter {
	return &v
}

type NullableTakerBuySellVolumeContractTypeParameter struct {
	value *TakerBuySellVolumeContractTypeParameter
	isSet bool
}

func (v NullableTakerBuySellVolumeContractTypeParameter) Get() *TakerBuySellVolumeContractTypeParameter {
	return v.value
}

func (v *NullableTakerBuySellVolumeContractTypeParameter) Set(val *TakerBuySellVolumeContractTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableTakerBuySellVolumeContractTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableTakerBuySellVolumeContractTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTakerBuySellVolumeContractTypeParameter(val *TakerBuySellVolumeContractTypeParameter) *NullableTakerBuySellVolumeContractTypeParameter {
	return &NullableTakerBuySellVolumeContractTypeParameter{value: val, isSet: true}
}

func (v NullableTakerBuySellVolumeContractTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakerBuySellVolumeContractTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
