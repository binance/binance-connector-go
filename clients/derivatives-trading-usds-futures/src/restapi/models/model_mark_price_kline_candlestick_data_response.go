/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceKlineCandlestickDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceKlineCandlestickDataResponse{}

// MarkPriceKlineCandlestickDataResponse struct for MarkPriceKlineCandlestickDataResponse
type MarkPriceKlineCandlestickDataResponse struct {
	Items []MarkPriceKlineCandlestickDataResponseItem
}

// NewMarkPriceKlineCandlestickDataResponse instantiates a new MarkPriceKlineCandlestickDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceKlineCandlestickDataResponse() *MarkPriceKlineCandlestickDataResponse {
	this := MarkPriceKlineCandlestickDataResponse{}
	return &this
}

// NewMarkPriceKlineCandlestickDataResponseWithDefaults instantiates a new MarkPriceKlineCandlestickDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceKlineCandlestickDataResponseWithDefaults() *MarkPriceKlineCandlestickDataResponse {
	this := MarkPriceKlineCandlestickDataResponse{}
	return &this
}

func (o MarkPriceKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceKlineCandlestickDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceKlineCandlestickDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceKlineCandlestickDataResponse struct {
	value MarkPriceKlineCandlestickDataResponse
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickDataResponse) Get() MarkPriceKlineCandlestickDataResponse {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickDataResponse) Set(val MarkPriceKlineCandlestickDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickDataResponse) Unset() {
	v.value = MarkPriceKlineCandlestickDataResponse{}
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickDataResponse(val MarkPriceKlineCandlestickDataResponse) *NullableMarkPriceKlineCandlestickDataResponse {
	return &NullableMarkPriceKlineCandlestickDataResponse{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
