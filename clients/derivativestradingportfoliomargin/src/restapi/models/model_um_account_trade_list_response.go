/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmAccountTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmAccountTradeListResponse{}

// UmAccountTradeListResponse struct for UmAccountTradeListResponse
type UmAccountTradeListResponse struct {
	Items []UmAccountTradeListResponseInner
}

// NewUmAccountTradeListResponse instantiates a new UmAccountTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmAccountTradeListResponse() *UmAccountTradeListResponse {
	this := UmAccountTradeListResponse{}
	return &this
}

// NewUmAccountTradeListResponseWithDefaults instantiates a new UmAccountTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmAccountTradeListResponseWithDefaults() *UmAccountTradeListResponse {
	this := UmAccountTradeListResponse{}
	return &this
}

func (o UmAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmAccountTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *UmAccountTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableUmAccountTradeListResponse struct {
	value UmAccountTradeListResponse
	isSet bool
}

func (v NullableUmAccountTradeListResponse) Get() UmAccountTradeListResponse {
	return v.value
}

func (v *NullableUmAccountTradeListResponse) Set(val UmAccountTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUmAccountTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUmAccountTradeListResponse) Unset() {
	v.value = UmAccountTradeListResponse{}
	v.isSet = false
}

func NewNullableUmAccountTradeListResponse(val UmAccountTradeListResponse) *NullableUmAccountTradeListResponse {
	return &NullableUmAccountTradeListResponse{value: val, isSet: true}
}

func (v NullableUmAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmAccountTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
