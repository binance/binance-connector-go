/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CurrentOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentOpenOrdersResponse{}

// CurrentOpenOrdersResponse Array of open order objects. Fields are identical to the order rows in Equity Order History.
type CurrentOpenOrdersResponse struct {
	Items []CurrentOpenOrdersResponseInner
}

// NewCurrentOpenOrdersResponse instantiates a new CurrentOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentOpenOrdersResponse() *CurrentOpenOrdersResponse {
	this := CurrentOpenOrdersResponse{}
	return &this
}

// NewCurrentOpenOrdersResponseWithDefaults instantiates a new CurrentOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentOpenOrdersResponseWithDefaults() *CurrentOpenOrdersResponse {
	this := CurrentOpenOrdersResponse{}
	return &this
}

func (o CurrentOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CurrentOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCurrentOpenOrdersResponse struct {
	value CurrentOpenOrdersResponse
	isSet bool
}

func (v NullableCurrentOpenOrdersResponse) Get() CurrentOpenOrdersResponse {
	return v.value
}

func (v *NullableCurrentOpenOrdersResponse) Set(val CurrentOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentOpenOrdersResponse) Unset() {
	v.value = CurrentOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableCurrentOpenOrdersResponse(val CurrentOpenOrdersResponse) *NullableCurrentOpenOrdersResponse {
	return &NullableCurrentOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableCurrentOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
