/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// NewOrderPegPriceTypeParameter the model 'NewOrderPegPriceTypeParameter'
type NewOrderPegPriceTypeParameter string

// List of newOrder_pegPriceType_parameter
const (
	NewOrderPegPriceTypeParameterPrimaryPeg       NewOrderPegPriceTypeParameter = "PRIMARY_PEG"
	NewOrderPegPriceTypeParameterMarketPeg        NewOrderPegPriceTypeParameter = "MARKET_PEG"
	NewOrderPegPriceTypeParameterNonRepresentable NewOrderPegPriceTypeParameter = "NON_REPRESENTABLE"
)

// All allowed values of NewOrderPegPriceTypeParameter enum
var AllowedNewOrderPegPriceTypeParameterEnumValues = []NewOrderPegPriceTypeParameter{
	"PRIMARY_PEG",
	"MARKET_PEG",
	"NON_REPRESENTABLE",
}

func (v *NewOrderPegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NewOrderPegPriceTypeParameter(value)
	for _, existing := range AllowedNewOrderPegPriceTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NewOrderPegPriceTypeParameter", value)
}

// NewNewOrderPegPriceTypeParameterFromValue returns a pointer to a valid NewOrderPegPriceTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNewOrderPegPriceTypeParameterFromValue(v string) (*NewOrderPegPriceTypeParameter, error) {
	ev := NewOrderPegPriceTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NewOrderPegPriceTypeParameter: valid values are %v", v, AllowedNewOrderPegPriceTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NewOrderPegPriceTypeParameter) IsValid() bool {
	for _, existing := range AllowedNewOrderPegPriceTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to newOrder_pegPriceType_parameter value
func (v NewOrderPegPriceTypeParameter) Ptr() *NewOrderPegPriceTypeParameter {
	return &v
}

type NullableNewOrderPegPriceTypeParameter struct {
	value *NewOrderPegPriceTypeParameter
	isSet bool
}

func (v NullableNewOrderPegPriceTypeParameter) Get() *NewOrderPegPriceTypeParameter {
	return v.value
}

func (v *NullableNewOrderPegPriceTypeParameter) Set(val *NewOrderPegPriceTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderPegPriceTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderPegPriceTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderPegPriceTypeParameter(val *NewOrderPegPriceTypeParameter) *NullableNewOrderPegPriceTypeParameter {
	return &NullableNewOrderPegPriceTypeParameter{value: val, isSet: true}
}

func (v NullableNewOrderPegPriceTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderPegPriceTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
