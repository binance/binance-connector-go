/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetNFTWithdrawHistoryResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTWithdrawHistoryResponseListInner{}

// GetNFTWithdrawHistoryResponseListInner struct for GetNFTWithdrawHistoryResponseListInner
type GetNFTWithdrawHistoryResponseListInner struct {
	Network              *string  `json:"network,omitempty"`
	TxID                 *string  `json:"txID,omitempty"`
	ContractAdrress      *string  `json:"contractAdrress,omitempty"`
	TokenId              *string  `json:"tokenId,omitempty"`
	Timestamp            *int64   `json:"timestamp,omitempty"`
	Fee                  *float32 `json:"fee,omitempty"`
	FeeAsset             *string  `json:"feeAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTWithdrawHistoryResponseListInner GetNFTWithdrawHistoryResponseListInner

// NewGetNFTWithdrawHistoryResponseListInner instantiates a new GetNFTWithdrawHistoryResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTWithdrawHistoryResponseListInner() *GetNFTWithdrawHistoryResponseListInner {
	this := GetNFTWithdrawHistoryResponseListInner{}
	return &this
}

// NewGetNFTWithdrawHistoryResponseListInnerWithDefaults instantiates a new GetNFTWithdrawHistoryResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTWithdrawHistoryResponseListInnerWithDefaults() *GetNFTWithdrawHistoryResponseListInner {
	this := GetNFTWithdrawHistoryResponseListInner{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetNetwork(v string) {
	o.Network = &v
}

// GetTxID returns the TxID field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTxID() string {
	if o == nil || common.IsNil(o.TxID) {
		var ret string
		return ret
	}
	return *o.TxID
}

// GetTxIDOk returns a tuple with the TxID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTxIDOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxID) {
		return nil, false
	}
	return o.TxID, true
}

// HasTxID returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasTxID() bool {
	if o != nil && !common.IsNil(o.TxID) {
		return true
	}

	return false
}

// SetTxID gets a reference to the given string and assigns it to the TxID field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetTxID(v string) {
	o.TxID = &v
}

// GetContractAdrress returns the ContractAdrress field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetContractAdrress() string {
	if o == nil || common.IsNil(o.ContractAdrress) {
		var ret string
		return ret
	}
	return *o.ContractAdrress
}

// GetContractAdrressOk returns a tuple with the ContractAdrress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetContractAdrressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAdrress) {
		return nil, false
	}
	return o.ContractAdrress, true
}

// HasContractAdrress returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasContractAdrress() bool {
	if o != nil && !common.IsNil(o.ContractAdrress) {
		return true
	}

	return false
}

// SetContractAdrress gets a reference to the given string and assigns it to the ContractAdrress field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetContractAdrress(v string) {
	o.ContractAdrress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetFee() float32 {
	if o == nil || common.IsNil(o.Fee) {
		var ret float32
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetFeeOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given float32 and assigns it to the Fee field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetFee(v float32) {
	o.Fee = &v
}

// GetFeeAsset returns the FeeAsset field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponseListInner) GetFeeAsset() string {
	if o == nil || common.IsNil(o.FeeAsset) {
		var ret string
		return ret
	}
	return *o.FeeAsset
}

// GetFeeAssetOk returns a tuple with the FeeAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) GetFeeAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.FeeAsset) {
		return nil, false
	}
	return o.FeeAsset, true
}

// HasFeeAsset returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponseListInner) HasFeeAsset() bool {
	if o != nil && !common.IsNil(o.FeeAsset) {
		return true
	}

	return false
}

// SetFeeAsset gets a reference to the given string and assigns it to the FeeAsset field.
func (o *GetNFTWithdrawHistoryResponseListInner) SetFeeAsset(v string) {
	o.FeeAsset = &v
}

func (o GetNFTWithdrawHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTWithdrawHistoryResponseListInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.FeeAsset) {
		toSerialize["feeAsset"] = o.FeeAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNFTWithdrawHistoryResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetNFTWithdrawHistoryResponseListInner := _GetNFTWithdrawHistoryResponseListInner{}

	err = json.Unmarshal(data, &varGetNFTWithdrawHistoryResponseListInner)

	if err != nil {
		return err
	}

	*o = GetNFTWithdrawHistoryResponseListInner(varGetNFTWithdrawHistoryResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "network")
		delete(additionalProperties, "txID")
		delete(additionalProperties, "contractAdrress")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "feeAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTWithdrawHistoryResponseListInner struct {
	value *GetNFTWithdrawHistoryResponseListInner
	isSet bool
}

func (v NullableGetNFTWithdrawHistoryResponseListInner) Get() *GetNFTWithdrawHistoryResponseListInner {
	return v.value
}

func (v *NullableGetNFTWithdrawHistoryResponseListInner) Set(val *GetNFTWithdrawHistoryResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTWithdrawHistoryResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTWithdrawHistoryResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTWithdrawHistoryResponseListInner(val *GetNFTWithdrawHistoryResponseListInner) *NullableGetNFTWithdrawHistoryResponseListInner {
	return &NullableGetNFTWithdrawHistoryResponseListInner{value: val, isSet: true}
}

func (v NullableGetNFTWithdrawHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTWithdrawHistoryResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
