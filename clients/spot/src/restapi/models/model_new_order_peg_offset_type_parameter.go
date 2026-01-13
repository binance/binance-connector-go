/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderPegOffsetTypeParameter the model 'NewOrderPegOffsetTypeParameter'
type NewOrderPegOffsetTypeParameter string

// List of newOrder_pegOffsetType_parameter
const (
	NewOrderPegOffsetTypeParameterPriceLevel       NewOrderPegOffsetTypeParameter = "PRICE_LEVEL"
	NewOrderPegOffsetTypeParameterNonRepresentable NewOrderPegOffsetTypeParameter = "NON_REPRESENTABLE"
)

// All allowed values of NewOrderPegOffsetTypeParameter enum
var AllowedNewOrderPegOffsetTypeParameterEnumValues = []NewOrderPegOffsetTypeParameter{
	"PRICE_LEVEL",
	"NON_REPRESENTABLE",
}

func (v *NewOrderPegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderPegOffsetTypeParameter(value)
	for _, existing := range AllowedNewOrderPegOffsetTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderPegOffsetTypeParameter", value)
}

// NewNewOrderPegOffsetTypeParameterFromValue returns a pointer to a valid NewOrderPegOffsetTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderPegOffsetTypeParameterFromValue(v string) (*NewOrderPegOffsetTypeParameter, error) {
	ev := NewOrderPegOffsetTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderPegOffsetTypeParameter: valid values are %v", v, AllowedNewOrderPegOffsetTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderPegOffsetTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderPegOffsetTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_pegOffsetType_parameter value
func (v NewOrderPegOffsetTypeParameter) Ptr() *NewOrderPegOffsetTypeParameter {
	return &v
}

type NullableNewOrderPegOffsetTypeParameter struct {
	value *NewOrderPegOffsetTypeParameter
	isSet bool
}

func (v NullableNewOrderPegOffsetTypeParameter) Get() *NewOrderPegOffsetTypeParameter {
	return v.value
}

func (v *NullableNewOrderPegOffsetTypeParameter) Set(val *NewOrderPegOffsetTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderPegOffsetTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderPegOffsetTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderPegOffsetTypeParameter(val *NewOrderPegOffsetTypeParameter) *NullableNewOrderPegOffsetTypeParameter {
	return &NullableNewOrderPegOffsetTypeParameter{value: val, isSet: true}
}

func (v NullableNewOrderPegOffsetTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderPegOffsetTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
