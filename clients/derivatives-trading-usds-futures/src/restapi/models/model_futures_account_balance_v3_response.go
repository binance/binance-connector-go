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

// checks if the FuturesAccountBalanceV3Response type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesAccountBalanceV3Response{}

// FuturesAccountBalanceV3Response struct for FuturesAccountBalanceV3Response
type FuturesAccountBalanceV3Response struct {
	Items []FuturesAccountBalanceV2ResponseInner
}

// NewFuturesAccountBalanceV3Response instantiates a new FuturesAccountBalanceV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesAccountBalanceV3Response() *FuturesAccountBalanceV3Response {
	this := FuturesAccountBalanceV3Response{}
	return &this
}

// NewFuturesAccountBalanceV3ResponseWithDefaults instantiates a new FuturesAccountBalanceV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesAccountBalanceV3ResponseWithDefaults() *FuturesAccountBalanceV3Response {
	this := FuturesAccountBalanceV3Response{}
	return &this
}

func (o FuturesAccountBalanceV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesAccountBalanceV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FuturesAccountBalanceV3Response) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFuturesAccountBalanceV3Response struct {
	value FuturesAccountBalanceV3Response
	isSet bool
}

func (v NullableFuturesAccountBalanceV3Response) Get() FuturesAccountBalanceV3Response {
	return v.value
}

func (v *NullableFuturesAccountBalanceV3Response) Set(val FuturesAccountBalanceV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesAccountBalanceV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesAccountBalanceV3Response) Unset() {
	v.value = FuturesAccountBalanceV3Response{}
	v.isSet = false
}

func NewNullableFuturesAccountBalanceV3Response(val FuturesAccountBalanceV3Response) *NullableFuturesAccountBalanceV3Response {
	return &NullableFuturesAccountBalanceV3Response{value: val, isSet: true}
}

func (v NullableFuturesAccountBalanceV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesAccountBalanceV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
