/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"
	"fmt"
)

// OrderListPlaceOpoWorkingTypeParameter the model 'OrderListPlaceOpoWorkingTypeParameter'
type OrderListPlaceOpoWorkingTypeParameter string

// List of orderListPlaceOpo_workingType_parameter
const (
	OrderListPlaceOpoWorkingTypeParameterLimit      OrderListPlaceOpoWorkingTypeParameter = "LIMIT"
	OrderListPlaceOpoWorkingTypeParameterLimitMaker OrderListPlaceOpoWorkingTypeParameter = "LIMIT_MAKER"
)

// All allowed values of OrderListPlaceOpoWorkingTypeParameter enum
var AllowedOrderListPlaceOpoWorkingTypeParameterEnumValues = []OrderListPlaceOpoWorkingTypeParameter{
	"LIMIT",
	"LIMIT_MAKER",
}

func (v *OrderListPlaceOpoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderListPlaceOpoWorkingTypeParameter(value)
	for _, existing := range AllowedOrderListPlaceOpoWorkingTypeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderListPlaceOpoWorkingTypeParameter", value)
}

// NewOrderListPlaceOpoWorkingTypeParameterFromValue returns a pointer to a valid OrderListPlaceOpoWorkingTypeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderListPlaceOpoWorkingTypeParameterFromValue(v string) (*OrderListPlaceOpoWorkingTypeParameter, error) {
	ev := OrderListPlaceOpoWorkingTypeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderListPlaceOpoWorkingTypeParameter: valid values are %v", v, AllowedOrderListPlaceOpoWorkingTypeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderListPlaceOpoWorkingTypeParameter) IsValid() bool {
	for _, existing := range AllowedOrderListPlaceOpoWorkingTypeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to orderListPlaceOpo_workingType_parameter value
func (v OrderListPlaceOpoWorkingTypeParameter) Ptr() *OrderListPlaceOpoWorkingTypeParameter {
	return &v
}

type NullableOrderListPlaceOpoWorkingTypeParameter struct {
	value *OrderListPlaceOpoWorkingTypeParameter
	isSet bool
}

func (v NullableOrderListPlaceOpoWorkingTypeParameter) Get() *OrderListPlaceOpoWorkingTypeParameter {
	return v.value
}

func (v *NullableOrderListPlaceOpoWorkingTypeParameter) Set(val *OrderListPlaceOpoWorkingTypeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOpoWorkingTypeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOpoWorkingTypeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOpoWorkingTypeParameter(val *OrderListPlaceOpoWorkingTypeParameter) *NullableOrderListPlaceOpoWorkingTypeParameter {
	return &NullableOrderListPlaceOpoWorkingTypeParameter{value: val, isSet: true}
}

func (v NullableOrderListPlaceOpoWorkingTypeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOpoWorkingTypeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
