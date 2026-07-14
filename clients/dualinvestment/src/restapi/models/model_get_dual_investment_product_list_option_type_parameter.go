/*
Dual Investment REST API

Query products, request quotes, and subscribe to Advanced Earn Dual Investment strategies.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// GetDualInvestmentProductListOptionTypeParameter the model 'GetDualInvestmentProductListOptionTypeParameter'
type GetDualInvestmentProductListOptionTypeParameter string

// List of getDualInvestmentProductList_optionType_parameter
const (
	GetDualInvestmentProductListOptionTypeParameterCall GetDualInvestmentProductListOptionTypeParameter = "CALL"
	GetDualInvestmentProductListOptionTypeParameterPut  GetDualInvestmentProductListOptionTypeParameter = "PUT"
)

// All allowed values of GetDualInvestmentProductListOptionTypeParameter enum
var AllowedGetDualInvestmentProductListOptionTypeParameterEnumValues = []GetDualInvestmentProductListOptionTypeParameter{
	"CALL",
	"PUT",
}

func (v *GetDualInvestmentProductListOptionTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetDualInvestmentProductListOptionTypeParameter(value)
	for _, existing := range AllowedGetDualInvestmentProductListOptionTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetDualInvestmentProductListOptionTypeParameter", value)
}

// NewGetDualInvestmentProductListOptionTypeParameterFromValue returns a pointer to a valid GetDualInvestmentProductListOptionTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetDualInvestmentProductListOptionTypeParameterFromValue(v string) (*GetDualInvestmentProductListOptionTypeParameter, error) {
	ev := GetDualInvestmentProductListOptionTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetDualInvestmentProductListOptionTypeParameter: valid values are %v", v, AllowedGetDualInvestmentProductListOptionTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetDualInvestmentProductListOptionTypeParameter) IsValid() bool {
	for _, existing := range AllowedGetDualInvestmentProductListOptionTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getDualInvestmentProductList_optionType_parameter value
func (v GetDualInvestmentProductListOptionTypeParameter) Ptr() *GetDualInvestmentProductListOptionTypeParameter {
	return &v
}

type NullableGetDualInvestmentProductListOptionTypeParameter struct {
	value *GetDualInvestmentProductListOptionTypeParameter
	isSet bool
}

func (v NullableGetDualInvestmentProductListOptionTypeParameter) Get() *GetDualInvestmentProductListOptionTypeParameter {
	return v.value
}

func (v *NullableGetDualInvestmentProductListOptionTypeParameter) Set(val *GetDualInvestmentProductListOptionTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentProductListOptionTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentProductListOptionTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentProductListOptionTypeParameter(val *GetDualInvestmentProductListOptionTypeParameter) *NullableGetDualInvestmentProductListOptionTypeParameter {
	return &NullableGetDualInvestmentProductListOptionTypeParameter{value: val, isSet: true}
}

func (v NullableGetDualInvestmentProductListOptionTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentProductListOptionTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
