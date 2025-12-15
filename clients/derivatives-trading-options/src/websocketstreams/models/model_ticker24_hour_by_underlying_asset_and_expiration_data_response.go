/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24HourByUnderlyingAssetAndExpirationDataResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24HourByUnderlyingAssetAndExpirationDataResponse{}

// Ticker24HourByUnderlyingAssetAndExpirationDataResponse struct for Ticker24HourByUnderlyingAssetAndExpirationDataResponse
type Ticker24HourByUnderlyingAssetAndExpirationDataResponse struct {
	Items []Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner
}

// NewTicker24HourByUnderlyingAssetAndExpirationDataResponse instantiates a new Ticker24HourByUnderlyingAssetAndExpirationDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24HourByUnderlyingAssetAndExpirationDataResponse() *Ticker24HourByUnderlyingAssetAndExpirationDataResponse {
	this := Ticker24HourByUnderlyingAssetAndExpirationDataResponse{}
	return &this
}

// NewTicker24HourByUnderlyingAssetAndExpirationDataResponseWithDefaults instantiates a new Ticker24HourByUnderlyingAssetAndExpirationDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24HourByUnderlyingAssetAndExpirationDataResponseWithDefaults() *Ticker24HourByUnderlyingAssetAndExpirationDataResponse {
	this := Ticker24HourByUnderlyingAssetAndExpirationDataResponse{}
	return &this
}

func (o Ticker24HourByUnderlyingAssetAndExpirationDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24HourByUnderlyingAssetAndExpirationDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse struct {
	value Ticker24HourByUnderlyingAssetAndExpirationDataResponse
	isSet bool
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) Get() Ticker24HourByUnderlyingAssetAndExpirationDataResponse {
	return v.value
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) Set(val Ticker24HourByUnderlyingAssetAndExpirationDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) Unset() {
	v.value = Ticker24HourByUnderlyingAssetAndExpirationDataResponse{}
	v.isSet = false
}

func NewNullableTicker24HourByUnderlyingAssetAndExpirationDataResponse(val Ticker24HourByUnderlyingAssetAndExpirationDataResponse) *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse {
	return &NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse{value: val, isSet: true}
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
