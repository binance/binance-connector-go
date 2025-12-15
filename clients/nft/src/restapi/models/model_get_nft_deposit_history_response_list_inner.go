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

// checks if the GetNFTDepositHistoryResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTDepositHistoryResponseListInner{}

// GetNFTDepositHistoryResponseListInner struct for GetNFTDepositHistoryResponseListInner
type GetNFTDepositHistoryResponseListInner struct {
	Network              *string `json:"network,omitempty"`
	TxID                 *string `json:"txID,omitempty"`
	ContractAdrress      *string `json:"contractAdrress,omitempty"`
	TokenId              *string `json:"tokenId,omitempty"`
	Timestamp            *int64  `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTDepositHistoryResponseListInner GetNFTDepositHistoryResponseListInner

// NewGetNFTDepositHistoryResponseListInner instantiates a new GetNFTDepositHistoryResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTDepositHistoryResponseListInner() *GetNFTDepositHistoryResponseListInner {
	this := GetNFTDepositHistoryResponseListInner{}
	return &this
}

// NewGetNFTDepositHistoryResponseListInnerWithDefaults instantiates a new GetNFTDepositHistoryResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTDepositHistoryResponseListInnerWithDefaults() *GetNFTDepositHistoryResponseListInner {
	this := GetNFTDepositHistoryResponseListInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponseListInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponseListInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponseListInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetNFTDepositHistoryResponseListInner) SetNetwork(v string) {
	o.Network = &v
}

// GetTxID returns the TxID field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponseListInner) GetTxID() string {
	if o == nil || common.IsNil(o.TxID) {
		var ret string
		return ret
	}
	return *o.TxID
}

// GetTxIDOk returns a tuple with the TxID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponseListInner) GetTxIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxID) {
		return nil, false
	}
	return o.TxID, true
}

// HasTxID returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponseListInner) HasTxID() bool {
	if o != nil && !common.IsNil(o.TxID) {
		return true
	}

	return false
}

// SetTxID gets a reference to the given string and assigns it to the TxID field.
func (o *GetNFTDepositHistoryResponseListInner) SetTxID(v string) {
	o.TxID = &v
}

// GetContractAdrress returns the ContractAdrress field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponseListInner) GetContractAdrress() string {
	if o == nil || common.IsNil(o.ContractAdrress) {
		var ret string
		return ret
	}
	return *o.ContractAdrress
}

// GetContractAdrressOk returns a tuple with the ContractAdrress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponseListInner) GetContractAdrressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAdrress) {
		return nil, false
	}
	return o.ContractAdrress, true
}

// HasContractAdrress returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponseListInner) HasContractAdrress() bool {
	if o != nil && !common.IsNil(o.ContractAdrress) {
		return true
	}

	return false
}

// SetContractAdrress gets a reference to the given string and assigns it to the ContractAdrress field.
func (o *GetNFTDepositHistoryResponseListInner) SetContractAdrress(v string) {
	o.ContractAdrress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponseListInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponseListInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponseListInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetNFTDepositHistoryResponseListInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponseListInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponseListInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponseListInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetNFTDepositHistoryResponseListInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o GetNFTDepositHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTDepositHistoryResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.TxID) {
		toSerialize["txID"] = o.TxID
	}
	if !common.IsNil(o.ContractAdrress) {
		toSerialize["contractAdrress"] = o.ContractAdrress
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNFTDepositHistoryResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetNFTDepositHistoryResponseListInner := _GetNFTDepositHistoryResponseListInner{}

	err = json.Unmarshal(data, &varGetNFTDepositHistoryResponseListInner)

	if err != nil {
		return err
	}

	*o = GetNFTDepositHistoryResponseListInner(varGetNFTDepositHistoryResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "txID")
		delete(additionalProperties, "contractAdrress")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTDepositHistoryResponseListInner struct {
	value *GetNFTDepositHistoryResponseListInner
	isSet bool
}

func (v NullableGetNFTDepositHistoryResponseListInner) Get() *GetNFTDepositHistoryResponseListInner {
	return v.value
}

func (v *NullableGetNFTDepositHistoryResponseListInner) Set(val *GetNFTDepositHistoryResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTDepositHistoryResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTDepositHistoryResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTDepositHistoryResponseListInner(val *GetNFTDepositHistoryResponseListInner) *NullableGetNFTDepositHistoryResponseListInner {
	return &NullableGetNFTDepositHistoryResponseListInner{value: val, isSet: true}
}

func (v NullableGetNFTDepositHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTDepositHistoryResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
