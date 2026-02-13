/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSubAccountDepositHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSubAccountDepositHistoryResponseInner{}

// GetSubAccountDepositHistoryResponseInner struct for GetSubAccountDepositHistoryResponseInner
type GetSubAccountDepositHistoryResponseInner struct {
	Id                   *string `json:"id,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Network              *string `json:"network,omitempty"`
	Status               *int64  `json:"status,omitempty"`
	Address              *string `json:"address,omitempty"`
	AddressTag           *string `json:"addressTag,omitempty"`
	TxId                 *string `json:"txId,omitempty"`
	InsertTime           *int64  `json:"insertTime,omitempty"`
	TransferType         *int64  `json:"transferType,omitempty"`
	ConfirmTimes         *string `json:"confirmTimes,omitempty"`
	UnlockConfirm        *int64  `json:"unlockConfirm,omitempty"`
	WalletType           *int64  `json:"walletType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSubAccountDepositHistoryResponseInner GetSubAccountDepositHistoryResponseInner

// NewGetSubAccountDepositHistoryResponseInner instantiates a new GetSubAccountDepositHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubAccountDepositHistoryResponseInner() *GetSubAccountDepositHistoryResponseInner {
	this := GetSubAccountDepositHistoryResponseInner{}
	return &this
}

// NewGetSubAccountDepositHistoryResponseInnerWithDefaults instantiates a new GetSubAccountDepositHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubAccountDepositHistoryResponseInnerWithDefaults() *GetSubAccountDepositHistoryResponseInner {
	this := GetSubAccountDepositHistoryResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetSubAccountDepositHistoryResponseInner) SetId(v string) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetSubAccountDepositHistoryResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *GetSubAccountDepositHistoryResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *GetSubAccountDepositHistoryResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *GetSubAccountDepositHistoryResponseInner) SetStatus(v int64) {
	o.Status = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetSubAccountDepositHistoryResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetAddressTag returns the AddressTag field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetAddressTag() string {
	if o == nil || common.IsNil(o.AddressTag) {
		var ret string
		return ret
	}
	return *o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetAddressTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.AddressTag) {
		return nil, false
	}
	return o.AddressTag, true
}

// HasAddressTag returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasAddressTag() bool {
	if o != nil && !common.IsNil(o.AddressTag) {
		return true
	}

	return false
}

// SetAddressTag gets a reference to the given string and assigns it to the AddressTag field.
func (o *GetSubAccountDepositHistoryResponseInner) SetAddressTag(v string) {
	o.AddressTag = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetTxId() string {
	if o == nil || common.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetTxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *GetSubAccountDepositHistoryResponseInner) SetTxId(v string) {
	o.TxId = &v
}

// GetInsertTime returns the InsertTime field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetInsertTime() int64 {
	if o == nil || common.IsNil(o.InsertTime) {
		var ret int64
		return ret
	}
	return *o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetInsertTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InsertTime) {
		return nil, false
	}
	return o.InsertTime, true
}

// HasInsertTime returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasInsertTime() bool {
	if o != nil && !common.IsNil(o.InsertTime) {
		return true
	}

	return false
}

// SetInsertTime gets a reference to the given int64 and assigns it to the InsertTime field.
func (o *GetSubAccountDepositHistoryResponseInner) SetInsertTime(v int64) {
	o.InsertTime = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetTransferType() int64 {
	if o == nil || common.IsNil(o.TransferType) {
		var ret int64
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetTransferTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given int64 and assigns it to the TransferType field.
func (o *GetSubAccountDepositHistoryResponseInner) SetTransferType(v int64) {
	o.TransferType = &v
}

// GetConfirmTimes returns the ConfirmTimes field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetConfirmTimes() string {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		var ret string
		return ret
	}
	return *o.ConfirmTimes
}

// GetConfirmTimesOk returns a tuple with the ConfirmTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetConfirmTimesOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		return nil, false
	}
	return o.ConfirmTimes, true
}

// HasConfirmTimes returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasConfirmTimes() bool {
	if o != nil && !common.IsNil(o.ConfirmTimes) {
		return true
	}

	return false
}

// SetConfirmTimes gets a reference to the given string and assigns it to the ConfirmTimes field.
func (o *GetSubAccountDepositHistoryResponseInner) SetConfirmTimes(v string) {
	o.ConfirmTimes = &v
}

// GetUnlockConfirm returns the UnlockConfirm field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetUnlockConfirm() int64 {
	if o == nil || common.IsNil(o.UnlockConfirm) {
		var ret int64
		return ret
	}
	return *o.UnlockConfirm
}

// GetUnlockConfirmOk returns a tuple with the UnlockConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetUnlockConfirmOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UnlockConfirm) {
		return nil, false
	}
	return o.UnlockConfirm, true
}

// HasUnlockConfirm returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasUnlockConfirm() bool {
	if o != nil && !common.IsNil(o.UnlockConfirm) {
		return true
	}

	return false
}

// SetUnlockConfirm gets a reference to the given int64 and assigns it to the UnlockConfirm field.
func (o *GetSubAccountDepositHistoryResponseInner) SetUnlockConfirm(v int64) {
	o.UnlockConfirm = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *GetSubAccountDepositHistoryResponseInner) GetWalletType() int64 {
	if o == nil || common.IsNil(o.WalletType) {
		var ret int64
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubAccountDepositHistoryResponseInner) GetWalletTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *GetSubAccountDepositHistoryResponseInner) HasWalletType() bool {
	if o != nil && !common.IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given int64 and assigns it to the WalletType field.
func (o *GetSubAccountDepositHistoryResponseInner) SetWalletType(v int64) {
	o.WalletType = &v
}

func (o GetSubAccountDepositHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubAccountDepositHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.AddressTag) {
		toSerialize["addressTag"] = o.AddressTag
	}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !common.IsNil(o.InsertTime) {
		toSerialize["insertTime"] = o.InsertTime
	}
	if !common.IsNil(o.TransferType) {
		toSerialize["transferType"] = o.TransferType
	}
	if !common.IsNil(o.ConfirmTimes) {
		toSerialize["confirmTimes"] = o.ConfirmTimes
	}
	if !common.IsNil(o.UnlockConfirm) {
		toSerialize["unlockConfirm"] = o.UnlockConfirm
	}
	if !common.IsNil(o.WalletType) {
		toSerialize["walletType"] = o.WalletType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSubAccountDepositHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varGetSubAccountDepositHistoryResponseInner := _GetSubAccountDepositHistoryResponseInner{}

	err = json.Unmarshal(data, &varGetSubAccountDepositHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = GetSubAccountDepositHistoryResponseInner(varGetSubAccountDepositHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "network")
		delete(additionalProperties, "status")
		delete(additionalProperties, "address")
		delete(additionalProperties, "addressTag")
		delete(additionalProperties, "txId")
		delete(additionalProperties, "insertTime")
		delete(additionalProperties, "transferType")
		delete(additionalProperties, "confirmTimes")
		delete(additionalProperties, "unlockConfirm")
		delete(additionalProperties, "walletType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSubAccountDepositHistoryResponseInner struct {
	value *GetSubAccountDepositHistoryResponseInner
	isSet bool
}

func (v NullableGetSubAccountDepositHistoryResponseInner) Get() *GetSubAccountDepositHistoryResponseInner {
	return v.value
}

func (v *NullableGetSubAccountDepositHistoryResponseInner) Set(val *GetSubAccountDepositHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubAccountDepositHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubAccountDepositHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubAccountDepositHistoryResponseInner(val *GetSubAccountDepositHistoryResponseInner) *NullableGetSubAccountDepositHistoryResponseInner {
	return &NullableGetSubAccountDepositHistoryResponseInner{value: val, isSet: true}
}

func (v NullableGetSubAccountDepositHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubAccountDepositHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
