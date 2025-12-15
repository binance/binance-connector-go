/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMiniTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMiniTickerResponse{}

// AllMiniTickerResponse struct for AllMiniTickerResponse
type AllMiniTickerResponse struct {
	Items []AllMiniTickerResponseInner
}

// NewAllMiniTickerResponse instantiates a new AllMiniTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMiniTickerResponse() *AllMiniTickerResponse {
	this := AllMiniTickerResponse{}
	return &this
}

// NewAllMiniTickerResponseWithDefaults instantiates a new AllMiniTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMiniTickerResponseWithDefaults() *AllMiniTickerResponse {
	this := AllMiniTickerResponse{}
	return &this
}

func (o AllMiniTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMiniTickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllMiniTickerResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllMiniTickerResponse struct {
	value AllMiniTickerResponse
	isSet bool
}

func (v NullableAllMiniTickerResponse) Get() AllMiniTickerResponse {
	return v.value
}

func (v *NullableAllMiniTickerResponse) Set(val AllMiniTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMiniTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMiniTickerResponse) Unset() {
	v.value = AllMiniTickerResponse{}
	v.isSet = false
}

func NewNullableAllMiniTickerResponse(val AllMiniTickerResponse) *NullableAllMiniTickerResponse {
	return &NullableAllMiniTickerResponse{value: val, isSet: true}
}

func (v NullableAllMiniTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMiniTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
