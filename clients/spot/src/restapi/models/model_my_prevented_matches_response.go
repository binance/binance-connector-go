/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MyPreventedMatchesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MyPreventedMatchesResponse{}

// MyPreventedMatchesResponse struct for MyPreventedMatchesResponse
type MyPreventedMatchesResponse struct {
	Items []MyPreventedMatchesResponseInner
}

// NewMyPreventedMatchesResponse instantiates a new MyPreventedMatchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyPreventedMatchesResponse() *MyPreventedMatchesResponse {
	this := MyPreventedMatchesResponse{}
	return &this
}

// NewMyPreventedMatchesResponseWithDefaults instantiates a new MyPreventedMatchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyPreventedMatchesResponseWithDefaults() *MyPreventedMatchesResponse {
	this := MyPreventedMatchesResponse{}
	return &this
}

func (o MyPreventedMatchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MyPreventedMatchesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MyPreventedMatchesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMyPreventedMatchesResponse struct {
	value MyPreventedMatchesResponse
	isSet bool
}

func (v NullableMyPreventedMatchesResponse) Get() MyPreventedMatchesResponse {
	return v.value
}

func (v *NullableMyPreventedMatchesResponse) Set(val MyPreventedMatchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMyPreventedMatchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMyPreventedMatchesResponse) Unset() {
	v.value = MyPreventedMatchesResponse{}
	v.isSet = false
}

func NewNullableMyPreventedMatchesResponse(val MyPreventedMatchesResponse) *NullableMyPreventedMatchesResponse {
	return &NullableMyPreventedMatchesResponse{value: val, isSet: true}
}

func (v NullableMyPreventedMatchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyPreventedMatchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
