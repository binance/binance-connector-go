/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FundingWalletResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FundingWalletResponse{}

// FundingWalletResponse struct for FundingWalletResponse
type FundingWalletResponse struct {
	Items []FundingWalletResponseInner
}

// NewFundingWalletResponse instantiates a new FundingWalletResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundingWalletResponse() *FundingWalletResponse {
	this := FundingWalletResponse{}
	return &this
}

// NewFundingWalletResponseWithDefaults instantiates a new FundingWalletResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundingWalletResponseWithDefaults() *FundingWalletResponse {
	this := FundingWalletResponse{}
	return &this
}

func (o FundingWalletResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FundingWalletResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *FundingWalletResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableFundingWalletResponse struct {
	value FundingWalletResponse
	isSet bool
}

func (v NullableFundingWalletResponse) Get() FundingWalletResponse {
	return v.value
}

func (v *NullableFundingWalletResponse) Set(val FundingWalletResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFundingWalletResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFundingWalletResponse) Unset() {
	v.value = FundingWalletResponse{}
	v.isSet = false
}

func NewNullableFundingWalletResponse(val FundingWalletResponse) *NullableFundingWalletResponse {
	return &NullableFundingWalletResponse{value: val, isSet: true}
}

func (v NullableFundingWalletResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundingWalletResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
