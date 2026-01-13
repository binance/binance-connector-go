/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FetchDepositAddressListWithNetworkResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchDepositAddressListWithNetworkResponse{}

// FetchDepositAddressListWithNetworkResponse struct for FetchDepositAddressListWithNetworkResponse
type FetchDepositAddressListWithNetworkResponse struct {
	Items []FetchDepositAddressListWithNetworkResponseInner
}

// NewFetchDepositAddressListWithNetworkResponse instantiates a new FetchDepositAddressListWithNetworkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchDepositAddressListWithNetworkResponse() *FetchDepositAddressListWithNetworkResponse {
	this := FetchDepositAddressListWithNetworkResponse{}
	return &this
}

// NewFetchDepositAddressListWithNetworkResponseWithDefaults instantiates a new FetchDepositAddressListWithNetworkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchDepositAddressListWithNetworkResponseWithDefaults() *FetchDepositAddressListWithNetworkResponse {
	this := FetchDepositAddressListWithNetworkResponse{}
	return &this
}

func (o FetchDepositAddressListWithNetworkResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchDepositAddressListWithNetworkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FetchDepositAddressListWithNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFetchDepositAddressListWithNetworkResponse struct {
	value FetchDepositAddressListWithNetworkResponse
	isSet bool
}

func (v NullableFetchDepositAddressListWithNetworkResponse) Get() FetchDepositAddressListWithNetworkResponse {
	return v.value
}

func (v *NullableFetchDepositAddressListWithNetworkResponse) Set(val FetchDepositAddressListWithNetworkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchDepositAddressListWithNetworkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchDepositAddressListWithNetworkResponse) Unset() {
	v.value = FetchDepositAddressListWithNetworkResponse{}
	v.isSet = false
}

func NewNullableFetchDepositAddressListWithNetworkResponse(val FetchDepositAddressListWithNetworkResponse) *NullableFetchDepositAddressListWithNetworkResponse {
	return &NullableFetchDepositAddressListWithNetworkResponse{value: val, isSet: true}
}

func (v NullableFetchDepositAddressListWithNetworkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchDepositAddressListWithNetworkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
