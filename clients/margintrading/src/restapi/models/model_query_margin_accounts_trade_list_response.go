/*
Binance Margin Trading REST API

OpenAPI Specification for the Binance Margin Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryMarginAccountsTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryMarginAccountsTradeListResponse{}

// QueryMarginAccountsTradeListResponse struct for QueryMarginAccountsTradeListResponse
type QueryMarginAccountsTradeListResponse struct {
	Items []QueryMarginAccountsTradeListResponseInner
}

// NewQueryMarginAccountsTradeListResponse instantiates a new QueryMarginAccountsTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryMarginAccountsTradeListResponse() *QueryMarginAccountsTradeListResponse {
	this := QueryMarginAccountsTradeListResponse{}
	return &this
}

// NewQueryMarginAccountsTradeListResponseWithDefaults instantiates a new QueryMarginAccountsTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryMarginAccountsTradeListResponseWithDefaults() *QueryMarginAccountsTradeListResponse {
	this := QueryMarginAccountsTradeListResponse{}
	return &this
}

func (o QueryMarginAccountsTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryMarginAccountsTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *QueryMarginAccountsTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableQueryMarginAccountsTradeListResponse struct {
	value QueryMarginAccountsTradeListResponse
	isSet bool
}

func (v NullableQueryMarginAccountsTradeListResponse) Get() QueryMarginAccountsTradeListResponse {
	return v.value
}

func (v *NullableQueryMarginAccountsTradeListResponse) Set(val QueryMarginAccountsTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryMarginAccountsTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryMarginAccountsTradeListResponse) Unset() {
	v.value = QueryMarginAccountsTradeListResponse{}
	v.isSet = false
}

func NewNullableQueryMarginAccountsTradeListResponse(val QueryMarginAccountsTradeListResponse) *NullableQueryMarginAccountsTradeListResponse {
	return &NullableQueryMarginAccountsTradeListResponse{value: val, isSet: true}
}

func (v NullableQueryMarginAccountsTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryMarginAccountsTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
