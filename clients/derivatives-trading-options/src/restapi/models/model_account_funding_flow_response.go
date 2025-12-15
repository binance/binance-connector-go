/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountFundingFlowResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountFundingFlowResponse{}

// AccountFundingFlowResponse struct for AccountFundingFlowResponse
type AccountFundingFlowResponse struct {
	Items []AccountFundingFlowResponseInner
}

// NewAccountFundingFlowResponse instantiates a new AccountFundingFlowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFundingFlowResponse() *AccountFundingFlowResponse {
	this := AccountFundingFlowResponse{}
	return &this
}

// NewAccountFundingFlowResponseWithDefaults instantiates a new AccountFundingFlowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFundingFlowResponseWithDefaults() *AccountFundingFlowResponse {
	this := AccountFundingFlowResponse{}
	return &this
}

func (o AccountFundingFlowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountFundingFlowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return map[string]interface{}{
		"items": toSerialize,
	}, nil
}

func (o *AccountFundingFlowResponse) UnmarshalJSON(data []byte) (err error) {
	return json.Unmarshal(data, &o.Items)
}

type NullableAccountFundingFlowResponse struct {
	value AccountFundingFlowResponse
	isSet bool
}

func (v NullableAccountFundingFlowResponse) Get() AccountFundingFlowResponse {
	return v.value
}

func (v *NullableAccountFundingFlowResponse) Set(val AccountFundingFlowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFundingFlowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFundingFlowResponse) Unset() {
	v.value = AccountFundingFlowResponse{}
	v.isSet = false
}

func NewNullableAccountFundingFlowResponse(val AccountFundingFlowResponse) *NullableAccountFundingFlowResponse {
	return &NullableAccountFundingFlowResponse{value: val, isSet: true}
}

func (v NullableAccountFundingFlowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFundingFlowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
