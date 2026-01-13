/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderAmendmentsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendmentsResponse{}

// OrderAmendmentsResponse struct for OrderAmendmentsResponse
type OrderAmendmentsResponse struct {
	Items []OrderAmendmentsResponseInner
}

// NewOrderAmendmentsResponse instantiates a new OrderAmendmentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendmentsResponse() *OrderAmendmentsResponse {
	this := OrderAmendmentsResponse{}
	return &this
}

// NewOrderAmendmentsResponseWithDefaults instantiates a new OrderAmendmentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendmentsResponseWithDefaults() *OrderAmendmentsResponse {
	this := OrderAmendmentsResponse{}
	return &this
}

func (o OrderAmendmentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendmentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OrderAmendmentsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOrderAmendmentsResponse struct {
	value OrderAmendmentsResponse
	isSet bool
}

func (v NullableOrderAmendmentsResponse) Get() OrderAmendmentsResponse {
	return v.value
}

func (v *NullableOrderAmendmentsResponse) Set(val OrderAmendmentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendmentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendmentsResponse) Unset() {
	v.value = OrderAmendmentsResponse{}
	v.isSet = false
}

func NewNullableOrderAmendmentsResponse(val OrderAmendmentsResponse) *NullableOrderAmendmentsResponse {
	return &NullableOrderAmendmentsResponse{value: val, isSet: true}
}

func (v NullableOrderAmendmentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendmentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
