/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetNFTTransactionHistoryResponseListInnerTokensInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTTransactionHistoryResponseListInnerTokensInner{}

// GetNFTTransactionHistoryResponseListInnerTokensInner struct for GetNFTTransactionHistoryResponseListInnerTokensInner
type GetNFTTransactionHistoryResponseListInnerTokensInner struct {
	Network              *string `json:"network,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	ContractAddress      *string `json:"contractAddress,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTTransactionHistoryResponseListInnerTokensInner GetNFTTransactionHistoryResponseListInnerTokensInner

// NewGetNFTTransactionHistoryResponseListInnerTokensInner instantiates a new GetNFTTransactionHistoryResponseListInnerTokensInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTTransactionHistoryResponseListInnerTokensInner() *GetNFTTransactionHistoryResponseListInnerTokensInner {
	this := GetNFTTransactionHistoryResponseListInnerTokensInner{}
	return &this
}

// NewGetNFTTransactionHistoryResponseListInnerTokensInnerWithDefaults instantiates a new GetNFTTransactionHistoryResponseListInnerTokensInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTTransactionHistoryResponseListInnerTokensInnerWithDefaults() *GetNFTTransactionHistoryResponseListInnerTokensInner {
	this := GetNFTTransactionHistoryResponseListInnerTokensInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) SetNetwork(v string) {
	o.Network = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetContractAddress() string {
	if o == nil || common.IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) GetContractAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) HasContractAddress() bool {
	if o != nil && !common.IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) SetContractAddress(v string) {
	o.ContractAddress = &v
}

func (o GetNFTTransactionHistoryResponseListInnerTokensInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTTransactionHistoryResponseListInnerTokensInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNFTTransactionHistoryResponseListInnerTokensInner) UnmarshalJSON(data []byte) (err error) {
	varGetNFTTransactionHistoryResponseListInnerTokensInner := _GetNFTTransactionHistoryResponseListInnerTokensInner{}

	err = json.Unmarshal(data, &varGetNFTTransactionHistoryResponseListInnerTokensInner)

	if err != nil {
		return err
	}

	*o = GetNFTTransactionHistoryResponseListInnerTokensInner(varGetNFTTransactionHistoryResponseListInnerTokensInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "contractAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTTransactionHistoryResponseListInnerTokensInner struct {
	value *GetNFTTransactionHistoryResponseListInnerTokensInner
	isSet bool
}

func (v NullableGetNFTTransactionHistoryResponseListInnerTokensInner) Get() *GetNFTTransactionHistoryResponseListInnerTokensInner {
	return v.value
}

func (v *NullableGetNFTTransactionHistoryResponseListInnerTokensInner) Set(val *GetNFTTransactionHistoryResponseListInnerTokensInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTTransactionHistoryResponseListInnerTokensInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTTransactionHistoryResponseListInnerTokensInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTTransactionHistoryResponseListInnerTokensInner(val *GetNFTTransactionHistoryResponseListInnerTokensInner) *NullableGetNFTTransactionHistoryResponseListInnerTokensInner {
	return &NullableGetNFTTransactionHistoryResponseListInnerTokensInner{value: val, isSet: true}
}

func (v NullableGetNFTTransactionHistoryResponseListInnerTokensInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTTransactionHistoryResponseListInnerTokensInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
