/*
Margin REST API

Access account information, borrow and repay assets, and trade with Binance Margin.
*/

package models

import (
	"encoding/json"
	"fmt"
)

// MarginAccountNewOtocoPendingAboveTypeParameter the model 'MarginAccountNewOtocoPendingAboveTypeParameter'
type MarginAccountNewOtocoPendingAboveTypeParameter string

// List of marginAccountNewOtoco_pendingAboveType_parameter
const (
	MarginAccountNewOtocoPendingAboveTypeParameterLimitMaker    MarginAccountNewOtocoPendingAboveTypeParameter = "LIMIT_MAKER"
	MarginAccountNewOtocoPendingAboveTypeParameterStopLoss      MarginAccountNewOtocoPendingAboveTypeParameter = "STOP_LOSS"
	MarginAccountNewOtocoPendingAboveTypeParameterStopLossLimit MarginAccountNewOtocoPendingAboveTypeParameter = "STOP_LOSS_LIMIT"
)

// All allowed values of MarginAccountNewOtocoPendingAboveTypeParameter enum
var AllowedMarginAccountNewOtocoPendingAboveTypeParameterEnumValues = []MarginAccountNewOtocoPendingAboveTypeParameter{
	"LIMIT_MAKER",
	"STOP_LOSS",
	"STOP_LOSS_LIMIT",
}

func (v *MarginAccountNewOtocoPendingAboveTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MarginAccountNewOtocoPendingAboveTypeParameter(value)
	for _, existing := range AllowedMarginAccountNewOtocoPendingAboveTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MarginAccountNewOtocoPendingAboveTypeParameter", value)
}

// NewMarginAccountNewOtocoPendingAboveTypeParameterFromValue returns a pointer to a valid MarginAccountNewOtocoPendingAboveTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMarginAccountNewOtocoPendingAboveTypeParameterFromValue(v string) (*MarginAccountNewOtocoPendingAboveTypeParameter, error) {
	ev := MarginAccountNewOtocoPendingAboveTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MarginAccountNewOtocoPendingAboveTypeParameter: valid values are %v", v, AllowedMarginAccountNewOtocoPendingAboveTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MarginAccountNewOtocoPendingAboveTypeParameter) IsValid() bool {
	for _, existing := range AllowedMarginAccountNewOtocoPendingAboveTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to marginAccountNewOtoco_pendingAboveType_parameter value
func (v MarginAccountNewOtocoPendingAboveTypeParameter) Ptr() *MarginAccountNewOtocoPendingAboveTypeParameter {
	return &v
}

type NullableMarginAccountNewOtocoPendingAboveTypeParameter struct {
	value *MarginAccountNewOtocoPendingAboveTypeParameter
	isSet bool
}

func (v NullableMarginAccountNewOtocoPendingAboveTypeParameter) Get() *MarginAccountNewOtocoPendingAboveTypeParameter {
	return v.value
}

func (v *NullableMarginAccountNewOtocoPendingAboveTypeParameter) Set(val *MarginAccountNewOtocoPendingAboveTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountNewOtocoPendingAboveTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountNewOtocoPendingAboveTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginAccountNewOtocoPendingAboveTypeParameter(val *MarginAccountNewOtocoPendingAboveTypeParameter) *NullableMarginAccountNewOtocoPendingAboveTypeParameter {
	return &NullableMarginAccountNewOtocoPendingAboveTypeParameter{value: val, isSet: true}
}

func (v NullableMarginAccountNewOtocoPendingAboveTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountNewOtocoPendingAboveTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
