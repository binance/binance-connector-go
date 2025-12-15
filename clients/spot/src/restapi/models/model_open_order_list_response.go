/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenOrderListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenOrderListResponse{}

// OpenOrderListResponse struct for OpenOrderListResponse
type OpenOrderListResponse struct {
	Items []OpenOrderListResponseInner
}

// NewOpenOrderListResponse instantiates a new OpenOrderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenOrderListResponse() *OpenOrderListResponse {
	this := OpenOrderListResponse{}
	return &this
}

// NewOpenOrderListResponseWithDefaults instantiates a new OpenOrderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenOrderListResponseWithDefaults() *OpenOrderListResponse {
	this := OpenOrderListResponse{}
	return &this
}

func (o OpenOrderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenOrderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *OpenOrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableOpenOrderListResponse struct {
	value OpenOrderListResponse
	isSet bool
}

func (v NullableOpenOrderListResponse) Get() OpenOrderListResponse {
	return v.value
}

func (v *NullableOpenOrderListResponse) Set(val OpenOrderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenOrderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenOrderListResponse) Unset() {
	v.value = OpenOrderListResponse{}
	v.isSet = false
}

func NewNullableOpenOrderListResponse(val OpenOrderListResponse) *NullableOpenOrderListResponse {
	return &NullableOpenOrderListResponse{value: val, isSet: true}
}

func (v NullableOpenOrderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenOrderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
