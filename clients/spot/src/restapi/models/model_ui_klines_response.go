/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UiKlinesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UiKlinesResponse{}

// UiKlinesResponse struct for UiKlinesResponse
type UiKlinesResponse struct {
	Items []UiKlinesItem
}

// NewUiKlinesResponse instantiates a new UiKlinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiKlinesResponse() *UiKlinesResponse {
	this := UiKlinesResponse{}
	return &this
}

// NewUiKlinesResponseWithDefaults instantiates a new UiKlinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiKlinesResponseWithDefaults() *UiKlinesResponse {
	this := UiKlinesResponse{}
	return &this
}

func (o UiKlinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiKlinesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UiKlinesResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUiKlinesResponse struct {
	value UiKlinesResponse
	isSet bool
}

func (v NullableUiKlinesResponse) Get() UiKlinesResponse {
	return v.value
}

func (v *NullableUiKlinesResponse) Set(val UiKlinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUiKlinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUiKlinesResponse) Unset() {
	v.value = UiKlinesResponse{}
	v.isSet = false
}

func NewNullableUiKlinesResponse(val UiKlinesResponse) *NullableUiKlinesResponse {
	return &NullableUiKlinesResponse{value: val, isSet: true}
}

func (v NullableUiKlinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiKlinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
