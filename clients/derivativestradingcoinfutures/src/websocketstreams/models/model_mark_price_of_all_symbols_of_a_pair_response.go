/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceOfAllSymbolsOfAPairResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceOfAllSymbolsOfAPairResponse{}

// MarkPriceOfAllSymbolsOfAPairResponse struct for MarkPriceOfAllSymbolsOfAPairResponse
type MarkPriceOfAllSymbolsOfAPairResponse struct {
	Items []MarkPriceOfAllSymbolsOfAPairResponseInner
}

// NewMarkPriceOfAllSymbolsOfAPairResponse instantiates a new MarkPriceOfAllSymbolsOfAPairResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceOfAllSymbolsOfAPairResponse() *MarkPriceOfAllSymbolsOfAPairResponse {
	this := MarkPriceOfAllSymbolsOfAPairResponse{}
	return &this
}

// NewMarkPriceOfAllSymbolsOfAPairResponseWithDefaults instantiates a new MarkPriceOfAllSymbolsOfAPairResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceOfAllSymbolsOfAPairResponseWithDefaults() *MarkPriceOfAllSymbolsOfAPairResponse {
	this := MarkPriceOfAllSymbolsOfAPairResponse{}
	return &this
}

func (o MarkPriceOfAllSymbolsOfAPairResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceOfAllSymbolsOfAPairResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarkPriceOfAllSymbolsOfAPairResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarkPriceOfAllSymbolsOfAPairResponse struct {
	value MarkPriceOfAllSymbolsOfAPairResponse
	isSet bool
}

func (v NullableMarkPriceOfAllSymbolsOfAPairResponse) Get() MarkPriceOfAllSymbolsOfAPairResponse {
	return v.value
}

func (v *NullableMarkPriceOfAllSymbolsOfAPairResponse) Set(val MarkPriceOfAllSymbolsOfAPairResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceOfAllSymbolsOfAPairResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceOfAllSymbolsOfAPairResponse) Unset() {
	v.value = MarkPriceOfAllSymbolsOfAPairResponse{}
	v.isSet = false
}

func NewNullableMarkPriceOfAllSymbolsOfAPairResponse(val MarkPriceOfAllSymbolsOfAPairResponse) *NullableMarkPriceOfAllSymbolsOfAPairResponse {
	return &NullableMarkPriceOfAllSymbolsOfAPairResponse{value: val, isSet: true}
}

func (v NullableMarkPriceOfAllSymbolsOfAPairResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceOfAllSymbolsOfAPairResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
