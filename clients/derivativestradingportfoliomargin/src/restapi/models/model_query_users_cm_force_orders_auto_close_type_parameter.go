/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"
	"fmt"
)

// QueryUsersCmForceOrdersAutoCloseTypeParameter the model 'QueryUsersCmForceOrdersAutoCloseTypeParameter'
type QueryUsersCmForceOrdersAutoCloseTypeParameter string

// List of queryUsersCmForceOrders_autoCloseType_parameter
const (
	QueryUsersCmForceOrdersAutoCloseTypeParameterLiquidation QueryUsersCmForceOrdersAutoCloseTypeParameter = "LIQUIDATION"
	QueryUsersCmForceOrdersAutoCloseTypeParameterAdl         QueryUsersCmForceOrdersAutoCloseTypeParameter = "ADL"
)

// All allowed values of QueryUsersCmForceOrdersAutoCloseTypeParameter enum
var AllowedQueryUsersCmForceOrdersAutoCloseTypeParameterEnumValues = []QueryUsersCmForceOrdersAutoCloseTypeParameter{
	"LIQUIDATION",
	"ADL",
}

func (v *QueryUsersCmForceOrdersAutoCloseTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryUsersCmForceOrdersAutoCloseTypeParameter(value)
	for _, existing := range AllowedQueryUsersCmForceOrdersAutoCloseTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryUsersCmForceOrdersAutoCloseTypeParameter", value)
}

// NewQueryUsersCmForceOrdersAutoCloseTypeParameterFromValue returns a pointer to a valid QueryUsersCmForceOrdersAutoCloseTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryUsersCmForceOrdersAutoCloseTypeParameterFromValue(v string) (*QueryUsersCmForceOrdersAutoCloseTypeParameter, error) {
	ev := QueryUsersCmForceOrdersAutoCloseTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryUsersCmForceOrdersAutoCloseTypeParameter: valid values are %v", v, AllowedQueryUsersCmForceOrdersAutoCloseTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryUsersCmForceOrdersAutoCloseTypeParameter) IsValid() bool {
	for _, existing := range AllowedQueryUsersCmForceOrdersAutoCloseTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to queryUsersCmForceOrders_autoCloseType_parameter value
func (v QueryUsersCmForceOrdersAutoCloseTypeParameter) Ptr() *QueryUsersCmForceOrdersAutoCloseTypeParameter {
	return &v
}

type NullableQueryUsersCmForceOrdersAutoCloseTypeParameter struct {
	value *QueryUsersCmForceOrdersAutoCloseTypeParameter
	isSet bool
}

func (v NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) Get() *QueryUsersCmForceOrdersAutoCloseTypeParameter {
	return v.value
}

func (v *NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) Set(val *QueryUsersCmForceOrdersAutoCloseTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUsersCmForceOrdersAutoCloseTypeParameter(val *QueryUsersCmForceOrdersAutoCloseTypeParameter) *NullableQueryUsersCmForceOrdersAutoCloseTypeParameter {
	return &NullableQueryUsersCmForceOrdersAutoCloseTypeParameter{value: val, isSet: true}
}

func (v NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUsersCmForceOrdersAutoCloseTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
