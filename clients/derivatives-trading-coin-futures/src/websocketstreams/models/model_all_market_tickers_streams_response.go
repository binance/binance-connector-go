/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMarketTickersStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketTickersStreamsResponse{}

// AllMarketTickersStreamsResponse struct for AllMarketTickersStreamsResponse
type AllMarketTickersStreamsResponse struct {
	Items []AllMarketTickersStreamsResponseInner
}

// NewAllMarketTickersStreamsResponse instantiates a new AllMarketTickersStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketTickersStreamsResponse() *AllMarketTickersStreamsResponse {
	this := AllMarketTickersStreamsResponse{}
	return &this
}

// NewAllMarketTickersStreamsResponseWithDefaults instantiates a new AllMarketTickersStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketTickersStreamsResponseWithDefaults() *AllMarketTickersStreamsResponse {
	this := AllMarketTickersStreamsResponse{}
	return &this
}

func (o AllMarketTickersStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketTickersStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllMarketTickersStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllMarketTickersStreamsResponse struct {
	value AllMarketTickersStreamsResponse
	isSet bool
}

func (v NullableAllMarketTickersStreamsResponse) Get() AllMarketTickersStreamsResponse {
	return v.value
}

func (v *NullableAllMarketTickersStreamsResponse) Set(val AllMarketTickersStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketTickersStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketTickersStreamsResponse) Unset() {
	v.value = AllMarketTickersStreamsResponse{}
	v.isSet = false
}

func NewNullableAllMarketTickersStreamsResponse(val AllMarketTickersStreamsResponse) *NullableAllMarketTickersStreamsResponse {
	return &NullableAllMarketTickersStreamsResponse{value: val, isSet: true}
}

func (v NullableAllMarketTickersStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketTickersStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
