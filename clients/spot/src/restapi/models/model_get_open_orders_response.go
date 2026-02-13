/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOpenOrdersResponse{}

// GetOpenOrdersResponse struct for GetOpenOrdersResponse
type GetOpenOrdersResponse struct {
	Items []AllOrdersResponseInner
}

// NewGetOpenOrdersResponse instantiates a new GetOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOpenOrdersResponse() *GetOpenOrdersResponse {
	this := GetOpenOrdersResponse{}
	return &this
}

// NewGetOpenOrdersResponseWithDefaults instantiates a new GetOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOpenOrdersResponseWithDefaults() *GetOpenOrdersResponse {
	this := GetOpenOrdersResponse{}
	return &this
}

func (o GetOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetOpenOrdersResponse struct {
	value GetOpenOrdersResponse
	isSet bool
}

func (v NullableGetOpenOrdersResponse) Get() GetOpenOrdersResponse {
	return v.value
}

func (v *NullableGetOpenOrdersResponse) Set(val GetOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOpenOrdersResponse) Unset() {
	v.value = GetOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableGetOpenOrdersResponse(val GetOpenOrdersResponse) *NullableGetOpenOrdersResponse {
	return &NullableGetOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableGetOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
