/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesAccountBalanceResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceResponse{}

// FuturesAccountBalanceResponse struct for FuturesAccountBalanceResponse
type FuturesAccountBalanceResponse struct {
	Items []FuturesAccountBalanceResponseInner
}

// NewFuturesAccountBalanceResponse instantiates a new FuturesAccountBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceResponse() *FuturesAccountBalanceResponse {
	this := FuturesAccountBalanceResponse{}
	return &this
}

// NewFuturesAccountBalanceResponseWithDefaults instantiates a new FuturesAccountBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceResponseWithDefaults() *FuturesAccountBalanceResponse {
	this := FuturesAccountBalanceResponse{}
	return &this
}

func (o FuturesAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FuturesAccountBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFuturesAccountBalanceResponse struct {
	value FuturesAccountBalanceResponse
	isSet bool
}

func (v NullableFuturesAccountBalanceResponse) Get() FuturesAccountBalanceResponse {
	return v.value
}

func (v *NullableFuturesAccountBalanceResponse) Set(val FuturesAccountBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceResponse) Unset() {
	v.value = FuturesAccountBalanceResponse{}
	v.isSet = false
}

func NewNullableFuturesAccountBalanceResponse(val FuturesAccountBalanceResponse) *NullableFuturesAccountBalanceResponse {
	return &NullableFuturesAccountBalanceResponse{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
