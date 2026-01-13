/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginAccountTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginAccountTradeListResponse{}

// MarginAccountTradeListResponse struct for MarginAccountTradeListResponse
type MarginAccountTradeListResponse struct {
	Items []MarginAccountTradeListResponseInner
}

// NewMarginAccountTradeListResponse instantiates a new MarginAccountTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginAccountTradeListResponse() *MarginAccountTradeListResponse {
	this := MarginAccountTradeListResponse{}
	return &this
}

// NewMarginAccountTradeListResponseWithDefaults instantiates a new MarginAccountTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginAccountTradeListResponseWithDefaults() *MarginAccountTradeListResponse {
	this := MarginAccountTradeListResponse{}
	return &this
}

func (o MarginAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginAccountTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *MarginAccountTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableMarginAccountTradeListResponse struct {
	value MarginAccountTradeListResponse
	isSet bool
}

func (v NullableMarginAccountTradeListResponse) Get() MarginAccountTradeListResponse {
	return v.value
}

func (v *NullableMarginAccountTradeListResponse) Set(val MarginAccountTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginAccountTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginAccountTradeListResponse) Unset() {
	v.value = MarginAccountTradeListResponse{}
	v.isSet = false
}

func NewNullableMarginAccountTradeListResponse(val MarginAccountTradeListResponse) *NullableMarginAccountTradeListResponse {
	return &NullableMarginAccountTradeListResponse{value: val, isSet: true}
}

func (v NullableMarginAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginAccountTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
