/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetBoostRewardsHistoryTypeParameter the model 'GetBoostRewardsHistoryTypeParameter'
type GetBoostRewardsHistoryTypeParameter string

// List of getBoostRewardsHistory_type_parameter
const (
	GetBoostRewardsHistoryTypeParameterClaim      GetBoostRewardsHistoryTypeParameter = "CLAIM"
	GetBoostRewardsHistoryTypeParameterDistribute GetBoostRewardsHistoryTypeParameter = "DISTRIBUTE"
)

// All allowed values of GetBoostRewardsHistoryTypeParameter enum
var AllowedGetBoostRewardsHistoryTypeParameterEnumValues = []GetBoostRewardsHistoryTypeParameter{
	"CLAIM",
	"DISTRIBUTE",
}

func (v *GetBoostRewardsHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetBoostRewardsHistoryTypeParameter(value)
	for _, existing := range AllowedGetBoostRewardsHistoryTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetBoostRewardsHistoryTypeParameter", value)
}

// NewGetBoostRewardsHistoryTypeParameterFromValue returns a pointer to a valid GetBoostRewardsHistoryTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetBoostRewardsHistoryTypeParameterFromValue(v string) (*GetBoostRewardsHistoryTypeParameter, error) {
	ev := GetBoostRewardsHistoryTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetBoostRewardsHistoryTypeParameter: valid values are %v", v, AllowedGetBoostRewardsHistoryTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetBoostRewardsHistoryTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetBoostRewardsHistoryTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getBoostRewardsHistory_type_parameter value
func (v GetBoostRewardsHistoryTypeParameter) Ptr() *GetBoostRewardsHistoryTypeParameter {
	return &v
}

type NullableGetBoostRewardsHistoryTypeParameter struct {
	value *GetBoostRewardsHistoryTypeParameter
	isSet bool
}

func (v NullableGetBoostRewardsHistoryTypeParameter) Get() *GetBoostRewardsHistoryTypeParameter {
	return v.value
}

func (v *NullableGetBoostRewardsHistoryTypeParameter) Set(val *GetBoostRewardsHistoryTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBoostRewardsHistoryTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBoostRewardsHistoryTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBoostRewardsHistoryTypeParameter(val *GetBoostRewardsHistoryTypeParameter) *NullableGetBoostRewardsHistoryTypeParameter {
	return &NullableGetBoostRewardsHistoryTypeParameter{value: val, isSet: true}
}

func (v NullableGetBoostRewardsHistoryTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBoostRewardsHistoryTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
