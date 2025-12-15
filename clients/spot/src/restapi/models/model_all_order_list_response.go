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

// checks if the AllOrderListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllOrderListResponse{}

// AllOrderListResponse struct for AllOrderListResponse
type AllOrderListResponse struct {
	Items []AllOrderListResponseInner
}

// NewAllOrderListResponse instantiates a new AllOrderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllOrderListResponse() *AllOrderListResponse {
	this := AllOrderListResponse{}
	return &this
}

// NewAllOrderListResponseWithDefaults instantiates a new AllOrderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllOrderListResponseWithDefaults() *AllOrderListResponse {
	this := AllOrderListResponse{}
	return &this
}

func (o AllOrderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllOrderListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllOrderListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllOrderListResponse struct {
	value AllOrderListResponse
	isSet bool
}

func (v NullableAllOrderListResponse) Get() AllOrderListResponse {
	return v.value
}

func (v *NullableAllOrderListResponse) Set(val AllOrderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllOrderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllOrderListResponse) Unset() {
	v.value = AllOrderListResponse{}
	v.isSet = false
}

func NewNullableAllOrderListResponse(val AllOrderListResponse) *NullableAllOrderListResponse {
	return &NullableAllOrderListResponse{value: val, isSet: true}
}

func (v NullableAllOrderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllOrderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
