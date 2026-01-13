/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CurrentAllAlgoOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentAllAlgoOpenOrdersResponse{}

// CurrentAllAlgoOpenOrdersResponse struct for CurrentAllAlgoOpenOrdersResponse
type CurrentAllAlgoOpenOrdersResponse struct {
	Items []CurrentAllAlgoOpenOrdersResponseInner
}

// NewCurrentAllAlgoOpenOrdersResponse instantiates a new CurrentAllAlgoOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentAllAlgoOpenOrdersResponse() *CurrentAllAlgoOpenOrdersResponse {
	this := CurrentAllAlgoOpenOrdersResponse{}
	return &this
}

// NewCurrentAllAlgoOpenOrdersResponseWithDefaults instantiates a new CurrentAllAlgoOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentAllAlgoOpenOrdersResponseWithDefaults() *CurrentAllAlgoOpenOrdersResponse {
	this := CurrentAllAlgoOpenOrdersResponse{}
	return &this
}

func (o CurrentAllAlgoOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentAllAlgoOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *CurrentAllAlgoOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableCurrentAllAlgoOpenOrdersResponse struct {
	value CurrentAllAlgoOpenOrdersResponse
	isSet bool
}

func (v NullableCurrentAllAlgoOpenOrdersResponse) Get() CurrentAllAlgoOpenOrdersResponse {
	return v.value
}

func (v *NullableCurrentAllAlgoOpenOrdersResponse) Set(val CurrentAllAlgoOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentAllAlgoOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentAllAlgoOpenOrdersResponse) Unset() {
	v.value = CurrentAllAlgoOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableCurrentAllAlgoOpenOrdersResponse(val CurrentAllAlgoOpenOrdersResponse) *NullableCurrentAllAlgoOpenOrdersResponse {
	return &NullableCurrentAllAlgoOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableCurrentAllAlgoOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentAllAlgoOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
