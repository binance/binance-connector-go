/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the SymbolOrderBookTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolOrderBookTickerResponse{}

// SymbolOrderBookTickerResponse struct for SymbolOrderBookTickerResponse
type SymbolOrderBookTickerResponse struct {
	Items []SymbolOrderBookTickerResponseInner
}

// NewSymbolOrderBookTickerResponse instantiates a new SymbolOrderBookTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolOrderBookTickerResponse() *SymbolOrderBookTickerResponse {
	this := SymbolOrderBookTickerResponse{}
	return &this
}

// NewSymbolOrderBookTickerResponseWithDefaults instantiates a new SymbolOrderBookTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolOrderBookTickerResponseWithDefaults() *SymbolOrderBookTickerResponse {
	this := SymbolOrderBookTickerResponse{}
	return &this
}

func (o SymbolOrderBookTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolOrderBookTickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SymbolOrderBookTickerResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSymbolOrderBookTickerResponse struct {
	value SymbolOrderBookTickerResponse
	isSet bool
}

func (v NullableSymbolOrderBookTickerResponse) Get() SymbolOrderBookTickerResponse {
	return v.value
}

func (v *NullableSymbolOrderBookTickerResponse) Set(val SymbolOrderBookTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolOrderBookTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolOrderBookTickerResponse) Unset() {
	v.value = SymbolOrderBookTickerResponse{}
	v.isSet = false
}

func NewNullableSymbolOrderBookTickerResponse(val SymbolOrderBookTickerResponse) *NullableSymbolOrderBookTickerResponse {
	return &NullableSymbolOrderBookTickerResponse{value: val, isSet: true}
}

func (v NullableSymbolOrderBookTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolOrderBookTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
