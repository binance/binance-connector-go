/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TokenizedAssetsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenizedAssetsResponse{}

// TokenizedAssetsResponse struct for TokenizedAssetsResponse
type TokenizedAssetsResponse struct {
	Items []TokenizedAssetsResponseInner
}

// NewTokenizedAssetsResponse instantiates a new TokenizedAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenizedAssetsResponse() *TokenizedAssetsResponse {
	this := TokenizedAssetsResponse{}
	return &this
}

// NewTokenizedAssetsResponseWithDefaults instantiates a new TokenizedAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenizedAssetsResponseWithDefaults() *TokenizedAssetsResponse {
	this := TokenizedAssetsResponse{}
	return &this
}

func (o TokenizedAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenizedAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *TokenizedAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableTokenizedAssetsResponse struct {
	value TokenizedAssetsResponse
	isSet bool
}

func (v NullableTokenizedAssetsResponse) Get() TokenizedAssetsResponse {
	return v.value
}

func (v *NullableTokenizedAssetsResponse) Set(val TokenizedAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenizedAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenizedAssetsResponse) Unset() {
	v.value = TokenizedAssetsResponse{}
	v.isSet = false
}

func NewNullableTokenizedAssetsResponse(val TokenizedAssetsResponse) *NullableTokenizedAssetsResponse {
	return &NullableTokenizedAssetsResponse{value: val, isSet: true}
}

func (v NullableTokenizedAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenizedAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
