/*
Staking REST API

Subscribe to staking products, track positions, and query rewards via the Binance Staking API.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// RedeemEthAssetParameter the model 'RedeemEthAssetParameter'
type RedeemEthAssetParameter string

// List of redeemEth_asset_parameter
const (
	RedeemEthAssetParameterWbeth RedeemEthAssetParameter = "WBETH"
	RedeemEthAssetParameterBeth  RedeemEthAssetParameter = "BETH"
)

// All allowed values of RedeemEthAssetParameter enum
var AllowedRedeemEthAssetParameterEnumValues = []RedeemEthAssetParameter{
	"WBETH",
	"BETH",
}

func (v *RedeemEthAssetParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedeemEthAssetParameter(value)
	for _, existing := range AllowedRedeemEthAssetParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedeemEthAssetParameter", value)
}

// NewRedeemEthAssetParameterFromValue returns a pointer to a valid RedeemEthAssetParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedeemEthAssetParameterFromValue(v string) (*RedeemEthAssetParameter, error) {
	ev := RedeemEthAssetParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedeemEthAssetParameter: valid values are %v", v, AllowedRedeemEthAssetParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedeemEthAssetParameter) IsValid() bool {
	for _, existing := range AllowedRedeemEthAssetParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to redeemEth_asset_parameter value
func (v RedeemEthAssetParameter) Ptr() *RedeemEthAssetParameter {
	return &v
}

type NullableRedeemEthAssetParameter struct {
	value *RedeemEthAssetParameter
	isSet bool
}

func (v NullableRedeemEthAssetParameter) Get() *RedeemEthAssetParameter {
	return v.value
}

func (v *NullableRedeemEthAssetParameter) Set(val *RedeemEthAssetParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemEthAssetParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemEthAssetParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemEthAssetParameter(val *RedeemEthAssetParameter) *NullableRedeemEthAssetParameter {
	return &NullableRedeemEthAssetParameter{value: val, isSet: true}
}

func (v NullableRedeemEthAssetParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemEthAssetParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
