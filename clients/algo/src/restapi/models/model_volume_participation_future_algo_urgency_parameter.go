/*
Algo Trading REST API

Programmatic access to Binance’s execution algorithms for creating and managing Spot and Futures algo orders.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// VolumeParticipationFutureAlgoUrgencyParameter the model 'VolumeParticipationFutureAlgoUrgencyParameter'
type VolumeParticipationFutureAlgoUrgencyParameter string

// List of volumeParticipationFutureAlgo_urgency_parameter
const (
	VolumeParticipationFutureAlgoUrgencyParameterLow    VolumeParticipationFutureAlgoUrgencyParameter = "LOW"
	VolumeParticipationFutureAlgoUrgencyParameterMedium VolumeParticipationFutureAlgoUrgencyParameter = "MEDIUM"
	VolumeParticipationFutureAlgoUrgencyParameterHigh   VolumeParticipationFutureAlgoUrgencyParameter = "HIGH"
)

// All allowed values of VolumeParticipationFutureAlgoUrgencyParameter enum
var AllowedVolumeParticipationFutureAlgoUrgencyParameterEnumValues = []VolumeParticipationFutureAlgoUrgencyParameter{
	"LOW",
	"MEDIUM",
	"HIGH",
}

func (v *VolumeParticipationFutureAlgoUrgencyParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VolumeParticipationFutureAlgoUrgencyParameter(value)
	for _, existing := range AllowedVolumeParticipationFutureAlgoUrgencyParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VolumeParticipationFutureAlgoUrgencyParameter", value)
}

// NewVolumeParticipationFutureAlgoUrgencyParameterFromValue returns a pointer to a valid VolumeParticipationFutureAlgoUrgencyParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVolumeParticipationFutureAlgoUrgencyParameterFromValue(v string) (*VolumeParticipationFutureAlgoUrgencyParameter, error) {
	ev := VolumeParticipationFutureAlgoUrgencyParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VolumeParticipationFutureAlgoUrgencyParameter: valid values are %v", v, AllowedVolumeParticipationFutureAlgoUrgencyParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VolumeParticipationFutureAlgoUrgencyParameter) IsValid() bool {
	for _, existing := range AllowedVolumeParticipationFutureAlgoUrgencyParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to volumeParticipationFutureAlgo_urgency_parameter value
func (v VolumeParticipationFutureAlgoUrgencyParameter) Ptr() *VolumeParticipationFutureAlgoUrgencyParameter {
	return &v
}

type NullableVolumeParticipationFutureAlgoUrgencyParameter struct {
	value *VolumeParticipationFutureAlgoUrgencyParameter
	isSet bool
}

func (v NullableVolumeParticipationFutureAlgoUrgencyParameter) Get() *VolumeParticipationFutureAlgoUrgencyParameter {
	return v.value
}

func (v *NullableVolumeParticipationFutureAlgoUrgencyParameter) Set(val *VolumeParticipationFutureAlgoUrgencyParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeParticipationFutureAlgoUrgencyParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeParticipationFutureAlgoUrgencyParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeParticipationFutureAlgoUrgencyParameter(val *VolumeParticipationFutureAlgoUrgencyParameter) *NullableVolumeParticipationFutureAlgoUrgencyParameter {
	return &NullableVolumeParticipationFutureAlgoUrgencyParameter{value: val, isSet: true}
}

func (v NullableVolumeParticipationFutureAlgoUrgencyParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeParticipationFutureAlgoUrgencyParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
