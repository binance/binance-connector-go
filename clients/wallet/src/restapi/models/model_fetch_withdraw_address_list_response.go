/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchWithdrawAddressListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchWithdrawAddressListResponse{}

// FetchWithdrawAddressListResponse struct for FetchWithdrawAddressListResponse
type FetchWithdrawAddressListResponse struct {
	Items []FetchWithdrawAddressListResponseInner
}

// NewFetchWithdrawAddressListResponse instantiates a new FetchWithdrawAddressListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchWithdrawAddressListResponse() *FetchWithdrawAddressListResponse {
	this := FetchWithdrawAddressListResponse{}
	return &this
}

// NewFetchWithdrawAddressListResponseWithDefaults instantiates a new FetchWithdrawAddressListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchWithdrawAddressListResponseWithDefaults() *FetchWithdrawAddressListResponse {
	this := FetchWithdrawAddressListResponse{}
	return &this
}

func (o FetchWithdrawAddressListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchWithdrawAddressListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FetchWithdrawAddressListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFetchWithdrawAddressListResponse struct {
	value FetchWithdrawAddressListResponse
	isSet bool
}

func (v NullableFetchWithdrawAddressListResponse) Get() FetchWithdrawAddressListResponse {
	return v.value
}

func (v *NullableFetchWithdrawAddressListResponse) Set(val FetchWithdrawAddressListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchWithdrawAddressListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchWithdrawAddressListResponse) Unset() {
	v.value = FetchWithdrawAddressListResponse{}
	v.isSet = false
}

func NewNullableFetchWithdrawAddressListResponse(val FetchWithdrawAddressListResponse) *NullableFetchWithdrawAddressListResponse {
	return &NullableFetchWithdrawAddressListResponse{value: val, isSet: true}
}

func (v NullableFetchWithdrawAddressListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchWithdrawAddressListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
