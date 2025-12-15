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

// checks if the AccountTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountTradeListResponse{}

// AccountTradeListResponse struct for AccountTradeListResponse
type AccountTradeListResponse struct {
	Items []AccountTradeListResponseInner
}

// NewAccountTradeListResponse instantiates a new AccountTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTradeListResponse() *AccountTradeListResponse {
	this := AccountTradeListResponse{}
	return &this
}

// NewAccountTradeListResponseWithDefaults instantiates a new AccountTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTradeListResponseWithDefaults() *AccountTradeListResponse {
	this := AccountTradeListResponse{}
	return &this
}

func (o AccountTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AccountTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAccountTradeListResponse struct {
	value AccountTradeListResponse
	isSet bool
}

func (v NullableAccountTradeListResponse) Get() AccountTradeListResponse {
	return v.value
}

func (v *NullableAccountTradeListResponse) Set(val AccountTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTradeListResponse) Unset() {
	v.value = AccountTradeListResponse{}
	v.isSet = false
}

func NewNullableAccountTradeListResponse(val AccountTradeListResponse) *NullableAccountTradeListResponse {
	return &NullableAccountTradeListResponse{value: val, isSet: true}
}

func (v NullableAccountTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
