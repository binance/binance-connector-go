/*
Simple Earn REST API

Earn rewards by subscribing to flexible or locked Simple Earn products.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetFlexibleRewardsHistoryTypeParameter the model 'GetFlexibleRewardsHistoryTypeParameter'
type GetFlexibleRewardsHistoryTypeParameter string

// List of getFlexibleRewardsHistory_type_parameter
const (
	GetFlexibleRewardsHistoryTypeParameterBonus    GetFlexibleRewardsHistoryTypeParameter = "BONUS"
	GetFlexibleRewardsHistoryTypeParameterRealtime GetFlexibleRewardsHistoryTypeParameter = "REALTIME"
	GetFlexibleRewardsHistoryTypeParameterRewards  GetFlexibleRewardsHistoryTypeParameter = "REWARDS"
	GetFlexibleRewardsHistoryTypeParameterAll      GetFlexibleRewardsHistoryTypeParameter = "ALL"
)

// All allowed values of GetFlexibleRewardsHistoryTypeParameter enum
var AllowedGetFlexibleRewardsHistoryTypeParameterEnumValues = []GetFlexibleRewardsHistoryTypeParameter{
	"BONUS",
	"REALTIME",
	"REWARDS",
	"ALL",
}

func (v *GetFlexibleRewardsHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetFlexibleRewardsHistoryTypeParameter(value)
	for _, existing := range AllowedGetFlexibleRewardsHistoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetFlexibleRewardsHistoryTypeParameter", value)
}

// NewGetFlexibleRewardsHistoryTypeParameterFromValue returns a pointer to a valid GetFlexibleRewardsHistoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetFlexibleRewardsHistoryTypeParameterFromValue(v string) (*GetFlexibleRewardsHistoryTypeParameter, error) {
	ev := GetFlexibleRewardsHistoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetFlexibleRewardsHistoryTypeParameter: valid values are %v", v, AllowedGetFlexibleRewardsHistoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetFlexibleRewardsHistoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetFlexibleRewardsHistoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getFlexibleRewardsHistory_type_parameter value
func (v GetFlexibleRewardsHistoryTypeParameter) Ptr() *GetFlexibleRewardsHistoryTypeParameter {
	return &v
}

type NullableGetFlexibleRewardsHistoryTypeParameter struct {
	value *GetFlexibleRewardsHistoryTypeParameter
	isSet bool
}

func (v NullableGetFlexibleRewardsHistoryTypeParameter) Get() *GetFlexibleRewardsHistoryTypeParameter {
	return v.value
}

func (v *NullableGetFlexibleRewardsHistoryTypeParameter) Set(val *GetFlexibleRewardsHistoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleRewardsHistoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleRewardsHistoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleRewardsHistoryTypeParameter(val *GetFlexibleRewardsHistoryTypeParameter) *NullableGetFlexibleRewardsHistoryTypeParameter {
	return &NullableGetFlexibleRewardsHistoryTypeParameter{value: val, isSet: true}
}

func (v NullableGetFlexibleRewardsHistoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleRewardsHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
