/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMarketMiniTickersStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketMiniTickersStreamResponse{}

// AllMarketMiniTickersStreamResponse struct for AllMarketMiniTickersStreamResponse
type AllMarketMiniTickersStreamResponse struct {
	Items []AllMarketMiniTickersStreamResponseInner
}

// NewAllMarketMiniTickersStreamResponse instantiates a new AllMarketMiniTickersStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketMiniTickersStreamResponse() *AllMarketMiniTickersStreamResponse {
	this := AllMarketMiniTickersStreamResponse{}
	return &this
}

// NewAllMarketMiniTickersStreamResponseWithDefaults instantiates a new AllMarketMiniTickersStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketMiniTickersStreamResponseWithDefaults() *AllMarketMiniTickersStreamResponse {
	this := AllMarketMiniTickersStreamResponse{}
	return &this
}

func (o AllMarketMiniTickersStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketMiniTickersStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AllMarketMiniTickersStreamResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAllMarketMiniTickersStreamResponse struct {
	value AllMarketMiniTickersStreamResponse
	isSet bool
}

func (v NullableAllMarketMiniTickersStreamResponse) Get() AllMarketMiniTickersStreamResponse {
	return v.value
}

func (v *NullableAllMarketMiniTickersStreamResponse) Set(val AllMarketMiniTickersStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketMiniTickersStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketMiniTickersStreamResponse) Unset() {
	v.value = AllMarketMiniTickersStreamResponse{}
	v.isSet = false
}

func NewNullableAllMarketMiniTickersStreamResponse(val AllMarketMiniTickersStreamResponse) *NullableAllMarketMiniTickersStreamResponse {
	return &NullableAllMarketMiniTickersStreamResponse{value: val, isSet: true}
}

func (v NullableAllMarketMiniTickersStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketMiniTickersStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
