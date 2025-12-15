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

// checks if the PlaceMultipleOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersResponse{}

// PlaceMultipleOrdersResponse struct for PlaceMultipleOrdersResponse
type PlaceMultipleOrdersResponse struct {
	Items []PlaceMultipleOrdersResponseInner
}

// NewPlaceMultipleOrdersResponse instantiates a new PlaceMultipleOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersResponse() *PlaceMultipleOrdersResponse {
	this := PlaceMultipleOrdersResponse{}
	return &this
}

// NewPlaceMultipleOrdersResponseWithDefaults instantiates a new PlaceMultipleOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersResponseWithDefaults() *PlaceMultipleOrdersResponse {
	this := PlaceMultipleOrdersResponse{}
	return &this
}

func (o PlaceMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceMultipleOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *PlaceMultipleOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullablePlaceMultipleOrdersResponse struct {
	value PlaceMultipleOrdersResponse
	isSet bool
}

func (v NullablePlaceMultipleOrdersResponse) Get() PlaceMultipleOrdersResponse {
	return v.value
}

func (v *NullablePlaceMultipleOrdersResponse) Set(val PlaceMultipleOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersResponse) Unset() {
	v.value = PlaceMultipleOrdersResponse{}
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersResponse(val PlaceMultipleOrdersResponse) *NullablePlaceMultipleOrdersResponse {
	return &NullablePlaceMultipleOrdersResponse{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
