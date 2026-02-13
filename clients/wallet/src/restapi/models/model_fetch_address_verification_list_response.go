/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FetchAddressVerificationListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FetchAddressVerificationListResponse{}

// FetchAddressVerificationListResponse struct for FetchAddressVerificationListResponse
type FetchAddressVerificationListResponse struct {
	Items []FetchAddressVerificationListResponseInner
}

// NewFetchAddressVerificationListResponse instantiates a new FetchAddressVerificationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchAddressVerificationListResponse() *FetchAddressVerificationListResponse {
	this := FetchAddressVerificationListResponse{}
	return &this
}

// NewFetchAddressVerificationListResponseWithDefaults instantiates a new FetchAddressVerificationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchAddressVerificationListResponseWithDefaults() *FetchAddressVerificationListResponse {
	this := FetchAddressVerificationListResponse{}
	return &this
}

func (o FetchAddressVerificationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchAddressVerificationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FetchAddressVerificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFetchAddressVerificationListResponse struct {
	value FetchAddressVerificationListResponse
	isSet bool
}

func (v NullableFetchAddressVerificationListResponse) Get() FetchAddressVerificationListResponse {
	return v.value
}

func (v *NullableFetchAddressVerificationListResponse) Set(val FetchAddressVerificationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchAddressVerificationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchAddressVerificationListResponse) Unset() {
	v.value = FetchAddressVerificationListResponse{}
	v.isSet = false
}

func NewNullableFetchAddressVerificationListResponse(val FetchAddressVerificationListResponse) *NullableFetchAddressVerificationListResponse {
	return &NullableFetchAddressVerificationListResponse{value: val, isSet: true}
}

func (v NullableFetchAddressVerificationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchAddressVerificationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
