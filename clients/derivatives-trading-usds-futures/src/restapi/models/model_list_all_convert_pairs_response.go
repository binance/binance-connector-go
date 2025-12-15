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

// checks if the ListAllConvertPairsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ListAllConvertPairsResponse{}

// ListAllConvertPairsResponse struct for ListAllConvertPairsResponse
type ListAllConvertPairsResponse struct {
	Items []ListAllConvertPairsResponseInner
}

// NewListAllConvertPairsResponse instantiates a new ListAllConvertPairsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAllConvertPairsResponse() *ListAllConvertPairsResponse {
	this := ListAllConvertPairsResponse{}
	return &this
}

// NewListAllConvertPairsResponseWithDefaults instantiates a new ListAllConvertPairsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAllConvertPairsResponseWithDefaults() *ListAllConvertPairsResponse {
	this := ListAllConvertPairsResponse{}
	return &this
}

func (o ListAllConvertPairsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListAllConvertPairsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *ListAllConvertPairsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableListAllConvertPairsResponse struct {
	value ListAllConvertPairsResponse
	isSet bool
}

func (v NullableListAllConvertPairsResponse) Get() ListAllConvertPairsResponse {
	return v.value
}

func (v *NullableListAllConvertPairsResponse) Set(val ListAllConvertPairsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListAllConvertPairsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListAllConvertPairsResponse) Unset() {
	v.value = ListAllConvertPairsResponse{}
	v.isSet = false
}

func NewNullableListAllConvertPairsResponse(val ListAllConvertPairsResponse) *NullableListAllConvertPairsResponse {
	return &NullableListAllConvertPairsResponse{value: val, isSet: true}
}

func (v NullableListAllConvertPairsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAllConvertPairsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
