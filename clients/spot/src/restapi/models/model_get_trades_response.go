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

// checks if the GetTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetTradesResponse{}

// GetTradesResponse struct for GetTradesResponse
type GetTradesResponse struct {
	Items []HistoricalTradesResponseInner
}

// NewGetTradesResponse instantiates a new GetTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTradesResponse() *GetTradesResponse {
	this := GetTradesResponse{}
	return &this
}

// NewGetTradesResponseWithDefaults instantiates a new GetTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTradesResponseWithDefaults() *GetTradesResponse {
	this := GetTradesResponse{}
	return &this
}

func (o GetTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *GetTradesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableGetTradesResponse struct {
	value GetTradesResponse
	isSet bool
}

func (v NullableGetTradesResponse) Get() GetTradesResponse {
	return v.value
}

func (v *NullableGetTradesResponse) Set(val GetTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTradesResponse) Unset() {
	v.value = GetTradesResponse{}
	v.isSet = false
}

func NewNullableGetTradesResponse(val GetTradesResponse) *NullableGetTradesResponse {
	return &NullableGetTradesResponse{value: val, isSet: true}
}

func (v NullableGetTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
