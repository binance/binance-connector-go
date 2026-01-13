/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DeleteOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DeleteOpenOrdersResponse{}

// DeleteOpenOrdersResponse struct for DeleteOpenOrdersResponse
type DeleteOpenOrdersResponse struct {
	Items []DeleteOpenOrdersResponseInner
}

// NewDeleteOpenOrdersResponse instantiates a new DeleteOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteOpenOrdersResponse() *DeleteOpenOrdersResponse {
	this := DeleteOpenOrdersResponse{}
	return &this
}

// NewDeleteOpenOrdersResponseWithDefaults instantiates a new DeleteOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteOpenOrdersResponseWithDefaults() *DeleteOpenOrdersResponse {
	this := DeleteOpenOrdersResponse{}
	return &this
}

func (o DeleteOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *DeleteOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableDeleteOpenOrdersResponse struct {
	value DeleteOpenOrdersResponse
	isSet bool
}

func (v NullableDeleteOpenOrdersResponse) Get() DeleteOpenOrdersResponse {
	return v.value
}

func (v *NullableDeleteOpenOrdersResponse) Set(val DeleteOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOpenOrdersResponse) Unset() {
	v.value = DeleteOpenOrdersResponse{}
	v.isSet = false
}

func NewNullableDeleteOpenOrdersResponse(val DeleteOpenOrdersResponse) *NullableDeleteOpenOrdersResponse {
	return &NullableDeleteOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableDeleteOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
