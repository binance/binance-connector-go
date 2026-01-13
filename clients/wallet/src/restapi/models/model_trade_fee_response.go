/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TradeFeeResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradeFeeResponse{}

// TradeFeeResponse struct for TradeFeeResponse
type TradeFeeResponse struct {
	Items []TradeFeeResponseInner
}

// NewTradeFeeResponse instantiates a new TradeFeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeFeeResponse() *TradeFeeResponse {
	this := TradeFeeResponse{}
	return &this
}

// NewTradeFeeResponseWithDefaults instantiates a new TradeFeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeFeeResponseWithDefaults() *TradeFeeResponse {
	this := TradeFeeResponse{}
	return &this
}

func (o TradeFeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeFeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TradeFeeResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTradeFeeResponse struct {
	value TradeFeeResponse
	isSet bool
}

func (v NullableTradeFeeResponse) Get() TradeFeeResponse {
	return v.value
}

func (v *NullableTradeFeeResponse) Set(val TradeFeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeFeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeFeeResponse) Unset() {
	v.value = TradeFeeResponse{}
	v.isSet = false
}

func NewNullableTradeFeeResponse(val TradeFeeResponse) *NullableTradeFeeResponse {
	return &NullableTradeFeeResponse{value: val, isSet: true}
}

func (v NullableTradeFeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeFeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
