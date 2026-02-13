/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RateLimitOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RateLimitOrderResponse{}

// RateLimitOrderResponse struct for RateLimitOrderResponse
type RateLimitOrderResponse struct {
	Items []RateLimitOrderResponseInner
}

// NewRateLimitOrderResponse instantiates a new RateLimitOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitOrderResponse() *RateLimitOrderResponse {
	this := RateLimitOrderResponse{}
	return &this
}

// NewRateLimitOrderResponseWithDefaults instantiates a new RateLimitOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitOrderResponseWithDefaults() *RateLimitOrderResponse {
	this := RateLimitOrderResponse{}
	return &this
}

func (o RateLimitOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *RateLimitOrderResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableRateLimitOrderResponse struct {
	value RateLimitOrderResponse
	isSet bool
}

func (v NullableRateLimitOrderResponse) Get() RateLimitOrderResponse {
	return v.value
}

func (v *NullableRateLimitOrderResponse) Set(val RateLimitOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitOrderResponse) Unset() {
	v.value = RateLimitOrderResponse{}
	v.isSet = false
}

func NewNullableRateLimitOrderResponse(val RateLimitOrderResponse) *NullableRateLimitOrderResponse {
	return &NullableRateLimitOrderResponse{value: val, isSet: true}
}

func (v NullableRateLimitOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
