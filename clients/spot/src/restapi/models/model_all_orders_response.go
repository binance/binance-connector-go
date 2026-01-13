/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllOrdersResponse{}

// AllOrdersResponse struct for AllOrdersResponse
type AllOrdersResponse struct {
	Items []AllOrdersResponseInner
}

// NewAllOrdersResponse instantiates a new AllOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllOrdersResponse() *AllOrdersResponse {
	this := AllOrdersResponse{}
	return &this
}

// NewAllOrdersResponseWithDefaults instantiates a new AllOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllOrdersResponseWithDefaults() *AllOrdersResponse {
	this := AllOrdersResponse{}
	return &this
}

func (o AllOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllOrdersResponse struct {
	value AllOrdersResponse
	isSet bool
}

func (v NullableAllOrdersResponse) Get() AllOrdersResponse {
	return v.value
}

func (v *NullableAllOrdersResponse) Set(val AllOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllOrdersResponse) Unset() {
	v.value = AllOrdersResponse{}
	v.isSet = false
}

func NewNullableAllOrdersResponse(val AllOrdersResponse) *NullableAllOrdersResponse {
	return &NullableAllOrdersResponse{value: val, isSet: true}
}

func (v NullableAllOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
