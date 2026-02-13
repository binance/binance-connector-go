/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MyAllocationsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyAllocationsResponse{}

// MyAllocationsResponse struct for MyAllocationsResponse
type MyAllocationsResponse struct {
	Items []MyAllocationsResponseInner
}

// NewMyAllocationsResponse instantiates a new MyAllocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyAllocationsResponse() *MyAllocationsResponse {
	this := MyAllocationsResponse{}
	return &this
}

// NewMyAllocationsResponseWithDefaults instantiates a new MyAllocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyAllocationsResponseWithDefaults() *MyAllocationsResponse {
	this := MyAllocationsResponse{}
	return &this
}

func (o MyAllocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyAllocationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MyAllocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMyAllocationsResponse struct {
	value MyAllocationsResponse
	isSet bool
}

func (v NullableMyAllocationsResponse) Get() MyAllocationsResponse {
	return v.value
}

func (v *NullableMyAllocationsResponse) Set(val MyAllocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMyAllocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMyAllocationsResponse) Unset() {
	v.value = MyAllocationsResponse{}
	v.isSet = false
}

func NewNullableMyAllocationsResponse(val MyAllocationsResponse) *NullableMyAllocationsResponse {
	return &NullableMyAllocationsResponse{value: val, isSet: true}
}

func (v NullableMyAllocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyAllocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
