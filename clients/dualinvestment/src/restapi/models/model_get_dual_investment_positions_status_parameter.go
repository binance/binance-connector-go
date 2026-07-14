/*
Dual Investment REST API

Query products, request quotes, and subscribe to Advanced Earn Dual Investment strategies.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetDualInvestmentPositionsStatusParameter the model 'GetDualInvestmentPositionsStatusParameter'
type GetDualInvestmentPositionsStatusParameter string

// List of getDualInvestmentPositions_status_parameter
const (
	GetDualInvestmentPositionsStatusParameterPending         GetDualInvestmentPositionsStatusParameter = "PENDING"
	GetDualInvestmentPositionsStatusParameterPurchaseSuccess GetDualInvestmentPositionsStatusParameter = "PURCHASE_SUCCESS"
	GetDualInvestmentPositionsStatusParameterSettled         GetDualInvestmentPositionsStatusParameter = "SETTLED"
	GetDualInvestmentPositionsStatusParameterPurchaseFail    GetDualInvestmentPositionsStatusParameter = "PURCHASE_FAIL"
	GetDualInvestmentPositionsStatusParameterRefunding       GetDualInvestmentPositionsStatusParameter = "REFUNDING"
	GetDualInvestmentPositionsStatusParameterRefundSuccess   GetDualInvestmentPositionsStatusParameter = "REFUND_SUCCESS"
	GetDualInvestmentPositionsStatusParameterSettling        GetDualInvestmentPositionsStatusParameter = "SETTLING"
)

// All allowed values of GetDualInvestmentPositionsStatusParameter enum
var AllowedGetDualInvestmentPositionsStatusParameterEnumValues = []GetDualInvestmentPositionsStatusParameter{
	"PENDING",
	"PURCHASE_SUCCESS",
	"SETTLED",
	"PURCHASE_FAIL",
	"REFUNDING",
	"REFUND_SUCCESS",
	"SETTLING",
}

func (v *GetDualInvestmentPositionsStatusParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetDualInvestmentPositionsStatusParameter(value)
	for _, existing := range AllowedGetDualInvestmentPositionsStatusParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetDualInvestmentPositionsStatusParameter", value)
}

// NewGetDualInvestmentPositionsStatusParameterFromValue returns a pointer to a valid GetDualInvestmentPositionsStatusParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetDualInvestmentPositionsStatusParameterFromValue(v string) (*GetDualInvestmentPositionsStatusParameter, error) {
	ev := GetDualInvestmentPositionsStatusParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetDualInvestmentPositionsStatusParameter: valid values are %v", v, AllowedGetDualInvestmentPositionsStatusParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetDualInvestmentPositionsStatusParameter) IsValid() bool {
	for _, existing := range AllowedGetDualInvestmentPositionsStatusParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getDualInvestmentPositions_status_parameter value
func (v GetDualInvestmentPositionsStatusParameter) Ptr() *GetDualInvestmentPositionsStatusParameter {
	return &v
}

type NullableGetDualInvestmentPositionsStatusParameter struct {
	value *GetDualInvestmentPositionsStatusParameter
	isSet bool
}

func (v NullableGetDualInvestmentPositionsStatusParameter) Get() *GetDualInvestmentPositionsStatusParameter {
	return v.value
}

func (v *NullableGetDualInvestmentPositionsStatusParameter) Set(val *GetDualInvestmentPositionsStatusParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentPositionsStatusParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentPositionsStatusParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentPositionsStatusParameter(val *GetDualInvestmentPositionsStatusParameter) *NullableGetDualInvestmentPositionsStatusParameter {
	return &NullableGetDualInvestmentPositionsStatusParameter{value: val, isSet: true}
}

func (v NullableGetDualInvestmentPositionsStatusParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentPositionsStatusParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
