/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceKlineCandlestickDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceKlineCandlestickDataResponse{}

// IndexPriceKlineCandlestickDataResponse struct for IndexPriceKlineCandlestickDataResponse
type IndexPriceKlineCandlestickDataResponse struct {
	Items []IndexPriceKlineCandlestickDataResponseItem
}

// NewIndexPriceKlineCandlestickDataResponse instantiates a new IndexPriceKlineCandlestickDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceKlineCandlestickDataResponse() *IndexPriceKlineCandlestickDataResponse {
	this := IndexPriceKlineCandlestickDataResponse{}
	return &this
}

// NewIndexPriceKlineCandlestickDataResponseWithDefaults instantiates a new IndexPriceKlineCandlestickDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceKlineCandlestickDataResponseWithDefaults() *IndexPriceKlineCandlestickDataResponse {
	this := IndexPriceKlineCandlestickDataResponse{}
	return &this
}

func (o IndexPriceKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceKlineCandlestickDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *IndexPriceKlineCandlestickDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableIndexPriceKlineCandlestickDataResponse struct {
	value IndexPriceKlineCandlestickDataResponse
	isSet bool
}

func (v NullableIndexPriceKlineCandlestickDataResponse) Get() IndexPriceKlineCandlestickDataResponse {
	return v.value
}

func (v *NullableIndexPriceKlineCandlestickDataResponse) Set(val IndexPriceKlineCandlestickDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceKlineCandlestickDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceKlineCandlestickDataResponse) Unset() {
	v.value = IndexPriceKlineCandlestickDataResponse{}
	v.isSet = false
}

func NewNullableIndexPriceKlineCandlestickDataResponse(val IndexPriceKlineCandlestickDataResponse) *NullableIndexPriceKlineCandlestickDataResponse {
	return &NullableIndexPriceKlineCandlestickDataResponse{value: val, isSet: true}
}

func (v NullableIndexPriceKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceKlineCandlestickDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
