/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CurrentAllOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentAllOpenOrdersResponse{}

// CurrentAllOpenOrdersResponse struct for CurrentAllOpenOrdersResponse
type CurrentAllOpenOrdersResponse struct {
	Items []AllOrdersResponseInner
}

// NewCurrentAllOpenOrdersResponse instantiates a new CurrentAllOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentAllOpenOrdersResponse() *CurrentAllOpenOrdersResponse {
	this := CurrentAllOpenOrdersResponse{}
	return &this
}

// NewCurrentAllOpenOrdersResponseWithDefaults instantiates a new CurrentAllOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentAllOpenOrdersResponseWithDefaults() *CurrentAllOpenOrdersResponse {
	this := CurrentAllOpenOrdersResponse{}
	return &this
}

func (o CurrentAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentAllOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CurrentAllOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCurrentAllOpenOrdersResponse struct {
	value CurrentAllOpenOrdersResponse
	isSet bool
}

func (v NullableCurrentAllOpenOrdersResponse) Get() CurrentAllOpenOrdersResponse {
	return v.value
}

func (v *NullableCurrentAllOpenOrdersResponse) Set(val CurrentAllOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentAllOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentAllOpenOrdersResponse) Unset() {
	v.value = CurrentAllOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableCurrentAllOpenOrdersResponse(val CurrentAllOpenOrdersResponse) *NullableCurrentAllOpenOrdersResponse {
	return &NullableCurrentAllOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableCurrentAllOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentAllOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
