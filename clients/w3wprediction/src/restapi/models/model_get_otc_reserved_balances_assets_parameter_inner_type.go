/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetOtcReservedBalancesAssetsParameterInnerType the model 'GetOtcReservedBalancesAssetsParameterInnerType'
type GetOtcReservedBalancesAssetsParameterInnerType string

// List of getOtcReservedBalances_assets_parameter_inner_type
const (
	GetOtcReservedBalancesAssetsParameterInnerTypeUsdt  GetOtcReservedBalancesAssetsParameterInnerType = "USDT"
	GetOtcReservedBalancesAssetsParameterInnerTypeShare GetOtcReservedBalancesAssetsParameterInnerType = "SHARE"
)

// All allowed values of GetOtcReservedBalancesAssetsParameterInnerType enum
var AllowedGetOtcReservedBalancesAssetsParameterInnerTypeEnumValues = []GetOtcReservedBalancesAssetsParameterInnerType{
	"USDT",
	"SHARE",
}

func (v *GetOtcReservedBalancesAssetsParameterInnerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetOtcReservedBalancesAssetsParameterInnerType(value)
	for _, existing := range AllowedGetOtcReservedBalancesAssetsParameterInnerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetOtcReservedBalancesAssetsParameterInnerType", value)
}

// NewGetOtcReservedBalancesAssetsParameterInnerTypeFromValue returns a pointer to a valid GetOtcReservedBalancesAssetsParameterInnerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetOtcReservedBalancesAssetsParameterInnerTypeFromValue(v string) (*GetOtcReservedBalancesAssetsParameterInnerType, error) {
	ev := GetOtcReservedBalancesAssetsParameterInnerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetOtcReservedBalancesAssetsParameterInnerType: valid values are %v", v, AllowedGetOtcReservedBalancesAssetsParameterInnerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetOtcReservedBalancesAssetsParameterInnerType) IsValid() bool {
	for _, existing := range AllowedGetOtcReservedBalancesAssetsParameterInnerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getOtcReservedBalances_assets_parameter_inner_type value
func (v GetOtcReservedBalancesAssetsParameterInnerType) Ptr() *GetOtcReservedBalancesAssetsParameterInnerType {
	return &v
}

type NullableGetOtcReservedBalancesAssetsParameterInnerType struct {
	value *GetOtcReservedBalancesAssetsParameterInnerType
	isSet bool
}

func (v NullableGetOtcReservedBalancesAssetsParameterInnerType) Get() *GetOtcReservedBalancesAssetsParameterInnerType {
	return v.value
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInnerType) Set(val *GetOtcReservedBalancesAssetsParameterInnerType) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOtcReservedBalancesAssetsParameterInnerType) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOtcReservedBalancesAssetsParameterInnerType(val *GetOtcReservedBalancesAssetsParameterInnerType) *NullableGetOtcReservedBalancesAssetsParameterInnerType {
	return &NullableGetOtcReservedBalancesAssetsParameterInnerType{value: val, isSet: true}
}

func (v NullableGetOtcReservedBalancesAssetsParameterInnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOtcReservedBalancesAssetsParameterInnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
