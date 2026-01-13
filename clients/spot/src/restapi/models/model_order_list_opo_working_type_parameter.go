/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListOpoWorkingTypeParameter the model 'OrderListOpoWorkingTypeParameter'
type OrderListOpoWorkingTypeParameter string

// List of orderListOpo_workingType_parameter
const (
	OrderListOpoWorkingTypeParameterLimit      OrderListOpoWorkingTypeParameter = "LIMIT"
	OrderListOpoWorkingTypeParameterLimitMaker OrderListOpoWorkingTypeParameter = "LIMIT_MAKER"
)

// All allowed values of OrderListOpoWorkingTypeParameter enum
var AllowedOrderListOpoWorkingTypeParameterEnumValues = []OrderListOpoWorkingTypeParameter{
	"LIMIT",
	"LIMIT_MAKER",
}

func (v *OrderListOpoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListOpoWorkingTypeParameter(value)
	for _, existing := range AllowedOrderListOpoWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListOpoWorkingTypeParameter", value)
}

// NewOrderListOpoWorkingTypeParameterFromValue returns a pointer to a valid OrderListOpoWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListOpoWorkingTypeParameterFromValue(v string) (*OrderListOpoWorkingTypeParameter, error) {
	ev := OrderListOpoWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListOpoWorkingTypeParameter: valid values are %v", v, AllowedOrderListOpoWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListOpoWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListOpoWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListOpo_workingType_parameter value
func (v OrderListOpoWorkingTypeParameter) Ptr() *OrderListOpoWorkingTypeParameter {
	return &v
}

type NullableOrderListOpoWorkingTypeParameter struct {
	value *OrderListOpoWorkingTypeParameter
	isSet bool
}

func (v NullableOrderListOpoWorkingTypeParameter) Get() *OrderListOpoWorkingTypeParameter {
	return v.value
}

func (v *NullableOrderListOpoWorkingTypeParameter) Set(val *OrderListOpoWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListOpoWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListOpoWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListOpoWorkingTypeParameter(val *OrderListOpoWorkingTypeParameter) *NullableOrderListOpoWorkingTypeParameter {
	return &NullableOrderListOpoWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListOpoWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListOpoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
