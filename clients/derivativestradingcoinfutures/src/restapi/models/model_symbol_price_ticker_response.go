/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the SymbolPriceTickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &SymbolPriceTickerResponse{}

// SymbolPriceTickerResponse struct for SymbolPriceTickerResponse
type SymbolPriceTickerResponse struct {
	Items []SymbolPriceTickerResponseInner
}

// NewSymbolPriceTickerResponse instantiates a new SymbolPriceTickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSymbolPriceTickerResponse() *SymbolPriceTickerResponse {
	this := SymbolPriceTickerResponse{}
	return &this
}

// NewSymbolPriceTickerResponseWithDefaults instantiates a new SymbolPriceTickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSymbolPriceTickerResponseWithDefaults() *SymbolPriceTickerResponse {
	this := SymbolPriceTickerResponse{}
	return &this
}

func (o SymbolPriceTickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SymbolPriceTickerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *SymbolPriceTickerResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableSymbolPriceTickerResponse struct {
	value SymbolPriceTickerResponse
	isSet bool
}

func (v NullableSymbolPriceTickerResponse) Get() SymbolPriceTickerResponse {
	return v.value
}

func (v *NullableSymbolPriceTickerResponse) Set(val SymbolPriceTickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSymbolPriceTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSymbolPriceTickerResponse) Unset() {
	v.value = SymbolPriceTickerResponse{}
	v.isSet = false
}

func NewNullableSymbolPriceTickerResponse(val SymbolPriceTickerResponse) *NullableSymbolPriceTickerResponse {
	return &NullableSymbolPriceTickerResponse{value: val, isSet: true}
}

func (v NullableSymbolPriceTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSymbolPriceTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
