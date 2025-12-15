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

// checks if the FuturesAccountBalanceV2Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceV2Response{}

// FuturesAccountBalanceV2Response struct for FuturesAccountBalanceV2Response
type FuturesAccountBalanceV2Response struct {
	Items []FuturesAccountBalanceV2ResponseInner
}

// NewFuturesAccountBalanceV2Response instantiates a new FuturesAccountBalanceV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceV2Response() *FuturesAccountBalanceV2Response {
	this := FuturesAccountBalanceV2Response{}
	return &this
}

// NewFuturesAccountBalanceV2ResponseWithDefaults instantiates a new FuturesAccountBalanceV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceV2ResponseWithDefaults() *FuturesAccountBalanceV2Response {
	this := FuturesAccountBalanceV2Response{}
	return &this
}

func (o FuturesAccountBalanceV2Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FuturesAccountBalanceV2Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFuturesAccountBalanceV2Response struct {
	value FuturesAccountBalanceV2Response
	isSet bool
}

func (v NullableFuturesAccountBalanceV2Response) Get() FuturesAccountBalanceV2Response {
	return v.value
}

func (v *NullableFuturesAccountBalanceV2Response) Set(val FuturesAccountBalanceV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceV2Response) Unset() {
	v.value = FuturesAccountBalanceV2Response{}
	v.isSet = false
}

func NewNullableFuturesAccountBalanceV2Response(val FuturesAccountBalanceV2Response) *NullableFuturesAccountBalanceV2Response {
	return &NullableFuturesAccountBalanceV2Response{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
