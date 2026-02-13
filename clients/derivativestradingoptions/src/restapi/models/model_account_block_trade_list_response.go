/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AccountBlockTradeListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountBlockTradeListResponse{}

// AccountBlockTradeListResponse struct for AccountBlockTradeListResponse
type AccountBlockTradeListResponse struct {
	Items []AccountBlockTradeListResponseInner
}

// NewAccountBlockTradeListResponse instantiates a new AccountBlockTradeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBlockTradeListResponse() *AccountBlockTradeListResponse {
	this := AccountBlockTradeListResponse{}
	return &this
}

// NewAccountBlockTradeListResponseWithDefaults instantiates a new AccountBlockTradeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBlockTradeListResponseWithDefaults() *AccountBlockTradeListResponse {
	this := AccountBlockTradeListResponse{}
	return &this
}

func (o AccountBlockTradeListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBlockTradeListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AccountBlockTradeListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAccountBlockTradeListResponse struct {
	value AccountBlockTradeListResponse
	isSet bool
}

func (v NullableAccountBlockTradeListResponse) Get() AccountBlockTradeListResponse {
	return v.value
}

func (v *NullableAccountBlockTradeListResponse) Set(val AccountBlockTradeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBlockTradeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBlockTradeListResponse) Unset() {
	v.value = AccountBlockTradeListResponse{}
	v.isSet = false
}

func NewNullableAccountBlockTradeListResponse(val AccountBlockTradeListResponse) *NullableAccountBlockTradeListResponse {
	return &NullableAccountBlockTradeListResponse{value: val, isSet: true}
}

func (v NullableAccountBlockTradeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBlockTradeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
