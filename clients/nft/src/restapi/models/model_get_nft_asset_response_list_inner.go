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

// checks if the GetNFTAssetResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTAssetResponseListInner{}

// GetNFTAssetResponseListInner struct for GetNFTAssetResponseListInner
type GetNFTAssetResponseListInner struct {
	Network              *string `json:"network,omitempty"`
	ContractAddress      *string `json:"contractAddress,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTAssetResponseListInner GetNFTAssetResponseListInner

// NewGetNFTAssetResponseListInner instantiates a new GetNFTAssetResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTAssetResponseListInner() *GetNFTAssetResponseListInner {
	this := GetNFTAssetResponseListInner{}
	return &this
}

// NewGetNFTAssetResponseListInnerWithDefaults instantiates a new GetNFTAssetResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTAssetResponseListInnerWithDefaults() *GetNFTAssetResponseListInner {
	this := GetNFTAssetResponseListInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetNFTAssetResponseListInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTAssetResponseListInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetNFTAssetResponseListInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetNFTAssetResponseListInner) SetNetwork(v string) {
	o.Network = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *GetNFTAssetResponseListInner) GetContractAddress() string {
	if o == nil || common.IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTAssetResponseListInner) GetContractAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *GetNFTAssetResponseListInner) HasContractAddress() bool {
	if o != nil && !common.IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *GetNFTAssetResponseListInner) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetNFTAssetResponseListInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTAssetResponseListInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetNFTAssetResponseListInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetNFTAssetResponseListInner) SetTokenId(v string) {
	o.TokenId = &v
}

func (o GetNFTAssetResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTAssetResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNFTAssetResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetNFTAssetResponseListInner := _GetNFTAssetResponseListInner{}

	err = json.Unmarshal(data, &varGetNFTAssetResponseListInner)

	if err != nil {
		return err
	}

	*o = GetNFTAssetResponseListInner(varGetNFTAssetResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "contractAddress")
		delete(additionalProperties, "tokenId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTAssetResponseListInner struct {
	value *GetNFTAssetResponseListInner
	isSet bool
}

func (v NullableGetNFTAssetResponseListInner) Get() *GetNFTAssetResponseListInner {
	return v.value
}

func (v *NullableGetNFTAssetResponseListInner) Set(val *GetNFTAssetResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTAssetResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTAssetResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTAssetResponseListInner(val *GetNFTAssetResponseListInner) *NullableGetNFTAssetResponseListInner {
	return &NullableGetNFTAssetResponseListInner{value: val, isSet: true}
}

func (v NullableGetNFTAssetResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTAssetResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
