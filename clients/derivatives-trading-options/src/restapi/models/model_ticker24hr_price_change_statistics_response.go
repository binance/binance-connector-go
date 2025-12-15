/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24hrPriceChangeStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24hrPriceChangeStatisticsResponse{}

// Ticker24hrPriceChangeStatisticsResponse struct for Ticker24hrPriceChangeStatisticsResponse
type Ticker24hrPriceChangeStatisticsResponse struct {
	Items []Ticker24hrPriceChangeStatisticsResponseInner
}

// NewTicker24hrPriceChangeStatisticsResponse instantiates a new Ticker24hrPriceChangeStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24hrPriceChangeStatisticsResponse() *Ticker24hrPriceChangeStatisticsResponse {
	this := Ticker24hrPriceChangeStatisticsResponse{}
	return &this
}

// NewTicker24hrPriceChangeStatisticsResponseWithDefaults instantiates a new Ticker24hrPriceChangeStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24hrPriceChangeStatisticsResponseWithDefaults() *Ticker24hrPriceChangeStatisticsResponse {
	this := Ticker24hrPriceChangeStatisticsResponse{}
	return &this
}

func (o Ticker24hrPriceChangeStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24hrPriceChangeStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *Ticker24hrPriceChangeStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTicker24hrPriceChangeStatisticsResponse struct {
	value Ticker24hrPriceChangeStatisticsResponse
	isSet bool
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) Get() Ticker24hrPriceChangeStatisticsResponse {
	return v.value
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) Set(val Ticker24hrPriceChangeStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) Unset() {
	v.value = Ticker24hrPriceChangeStatisticsResponse{}
	v.isSet = false
}

func NewNullableTicker24hrPriceChangeStatisticsResponse(val Ticker24hrPriceChangeStatisticsResponse) *NullableTicker24hrPriceChangeStatisticsResponse {
	return &NullableTicker24hrPriceChangeStatisticsResponse{value: val, isSet: true}
}

func (v NullableTicker24hrPriceChangeStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24hrPriceChangeStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
