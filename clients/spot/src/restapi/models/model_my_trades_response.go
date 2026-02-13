/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MyTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyTradesResponse{}

// MyTradesResponse struct for MyTradesResponse
type MyTradesResponse struct {
	Items []MyTradesResponseInner
}

// NewMyTradesResponse instantiates a new MyTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyTradesResponse() *MyTradesResponse {
	this := MyTradesResponse{}
	return &this
}

// NewMyTradesResponseWithDefaults instantiates a new MyTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyTradesResponseWithDefaults() *MyTradesResponse {
	this := MyTradesResponse{}
	return &this
}

func (o MyTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyTradesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MyTradesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMyTradesResponse struct {
	value MyTradesResponse
	isSet bool
}

func (v NullableMyTradesResponse) Get() MyTradesResponse {
	return v.value
}

func (v *NullableMyTradesResponse) Set(val MyTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMyTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMyTradesResponse) Unset() {
	v.value = MyTradesResponse{}
	v.isSet = false
}

func NewNullableMyTradesResponse(val MyTradesResponse) *NullableMyTradesResponse {
	return &NullableMyTradesResponse{value: val, isSet: true}
}

func (v NullableMyTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
