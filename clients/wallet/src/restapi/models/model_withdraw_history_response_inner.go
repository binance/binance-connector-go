/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawHistoryResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawHistoryResponseInner{}

// WithdrawHistoryResponseInner struct for WithdrawHistoryResponseInner
type WithdrawHistoryResponseInner struct {
	Id                   *string `json:"id,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TransactionFee       *string `json:"transactionFee,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Status               *int64  `json:"status,omitempty"`
	Address              *string `json:"address,omitempty"`
	TxId                 *string `json:"txId,omitempty"`
	ApplyTime            *string `json:"applyTime,omitempty"`
	Network              *string `json:"network,omitempty"`
	TransferType         *int64  `json:"transferType,omitempty"`
	WithdrawOrderId      *string `json:"withdrawOrderId,omitempty"`
	Info                 *string `json:"info,omitempty"`
	ConfirmNo            *int64  `json:"confirmNo,omitempty"`
	WalletType           *int64  `json:"walletType,omitempty"`
	TxKey                *string `json:"txKey,omitempty"`
	CompleteTime         *string `json:"completeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WithdrawHistoryResponseInner WithdrawHistoryResponseInner

// NewWithdrawHistoryResponseInner instantiates a new WithdrawHistoryResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawHistoryResponseInner() *WithdrawHistoryResponseInner {
	this := WithdrawHistoryResponseInner{}
	return &this
}

// NewWithdrawHistoryResponseInnerWithDefaults instantiates a new WithdrawHistoryResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawHistoryResponseInnerWithDefaults() *WithdrawHistoryResponseInner {
	this := WithdrawHistoryResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WithdrawHistoryResponseInner) SetId(v string) {
	o.Id = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *WithdrawHistoryResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetTransactionFee() string {
	if o == nil || common.IsNil(o.TransactionFee) {
		var ret string
		return ret
	}
	return *o.TransactionFee
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetTransactionFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionFee) {
		return nil, false
	}
	return o.TransactionFee, true
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasTransactionFee() bool {
	if o != nil && !common.IsNil(o.TransactionFee) {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given string and assigns it to the TransactionFee field.
func (o *WithdrawHistoryResponseInner) SetTransactionFee(v string) {
	o.TransactionFee = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *WithdrawHistoryResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *WithdrawHistoryResponseInner) SetStatus(v int64) {
	o.Status = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WithdrawHistoryResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetTxId() string {
	if o == nil || common.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetTxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *WithdrawHistoryResponseInner) SetTxId(v string) {
	o.TxId = &v
}

// GetApplyTime returns the ApplyTime field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetApplyTime() string {
	if o == nil || common.IsNil(o.ApplyTime) {
		var ret string
		return ret
	}
	return *o.ApplyTime
}

// GetApplyTimeOk returns a tuple with the ApplyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetApplyTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApplyTime) {
		return nil, false
	}
	return o.ApplyTime, true
}

// HasApplyTime returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasApplyTime() bool {
	if o != nil && !common.IsNil(o.ApplyTime) {
		return true
	}

	return false
}

// SetApplyTime gets a reference to the given string and assigns it to the ApplyTime field.
func (o *WithdrawHistoryResponseInner) SetApplyTime(v string) {
	o.ApplyTime = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *WithdrawHistoryResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetTransferType() int64 {
	if o == nil || common.IsNil(o.TransferType) {
		var ret int64
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetTransferTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given int64 and assigns it to the TransferType field.
func (o *WithdrawHistoryResponseInner) SetTransferType(v int64) {
	o.TransferType = &v
}

// GetWithdrawOrderId returns the WithdrawOrderId field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetWithdrawOrderId() string {
	if o == nil || common.IsNil(o.WithdrawOrderId) {
		var ret string
		return ret
	}
	return *o.WithdrawOrderId
}

// GetWithdrawOrderIdOk returns a tuple with the WithdrawOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetWithdrawOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawOrderId) {
		return nil, false
	}
	return o.WithdrawOrderId, true
}

// HasWithdrawOrderId returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasWithdrawOrderId() bool {
	if o != nil && !common.IsNil(o.WithdrawOrderId) {
		return true
	}

	return false
}

// SetWithdrawOrderId gets a reference to the given string and assigns it to the WithdrawOrderId field.
func (o *WithdrawHistoryResponseInner) SetWithdrawOrderId(v string) {
	o.WithdrawOrderId = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetInfo() string {
	if o == nil || common.IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasInfo() bool {
	if o != nil && !common.IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *WithdrawHistoryResponseInner) SetInfo(v string) {
	o.Info = &v
}

// GetConfirmNo returns the ConfirmNo field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetConfirmNo() int64 {
	if o == nil || common.IsNil(o.ConfirmNo) {
		var ret int64
		return ret
	}
	return *o.ConfirmNo
}

// GetConfirmNoOk returns a tuple with the ConfirmNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetConfirmNoOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ConfirmNo) {
		return nil, false
	}
	return o.ConfirmNo, true
}

// HasConfirmNo returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasConfirmNo() bool {
	if o != nil && !common.IsNil(o.ConfirmNo) {
		return true
	}

	return false
}

// SetConfirmNo gets a reference to the given int64 and assigns it to the ConfirmNo field.
func (o *WithdrawHistoryResponseInner) SetConfirmNo(v int64) {
	o.ConfirmNo = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetWalletType() int64 {
	if o == nil || common.IsNil(o.WalletType) {
		var ret int64
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetWalletTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasWalletType() bool {
	if o != nil && !common.IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given int64 and assigns it to the WalletType field.
func (o *WithdrawHistoryResponseInner) SetWalletType(v int64) {
	o.WalletType = &v
}

// GetTxKey returns the TxKey field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetTxKey() string {
	if o == nil || common.IsNil(o.TxKey) {
		var ret string
		return ret
	}
	return *o.TxKey
}

// GetTxKeyOk returns a tuple with the TxKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetTxKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxKey) {
		return nil, false
	}
	return o.TxKey, true
}

// HasTxKey returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasTxKey() bool {
	if o != nil && !common.IsNil(o.TxKey) {
		return true
	}

	return false
}

// SetTxKey gets a reference to the given string and assigns it to the TxKey field.
func (o *WithdrawHistoryResponseInner) SetTxKey(v string) {
	o.TxKey = &v
}

// GetCompleteTime returns the CompleteTime field value if set, zero value otherwise.
func (o *WithdrawHistoryResponseInner) GetCompleteTime() string {
	if o == nil || common.IsNil(o.CompleteTime) {
		var ret string
		return ret
	}
	return *o.CompleteTime
}

// GetCompleteTimeOk returns a tuple with the CompleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryResponseInner) GetCompleteTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompleteTime) {
		return nil, false
	}
	return o.CompleteTime, true
}

// HasCompleteTime returns a boolean if a field has been set.
func (o *WithdrawHistoryResponseInner) HasCompleteTime() bool {
	if o != nil && !common.IsNil(o.CompleteTime) {
		return true
	}

	return false
}

// SetCompleteTime gets a reference to the given string and assigns it to the CompleteTime field.
func (o *WithdrawHistoryResponseInner) SetCompleteTime(v string) {
	o.CompleteTime = &v
}

func (o WithdrawHistoryResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawHistoryResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.TransactionFee) {
		toSerialize["transactionFee"] = o.TransactionFee
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !common.IsNil(o.TxId) {
		toSerialize["txId"] = o.TxId
	}
	if !common.IsNil(o.ApplyTime) {
		toSerialize["applyTime"] = o.ApplyTime
	}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.TransferType) {
		toSerialize["transferType"] = o.TransferType
	}
	if !common.IsNil(o.WithdrawOrderId) {
		toSerialize["withdrawOrderId"] = o.WithdrawOrderId
	}
	if !common.IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !common.IsNil(o.ConfirmNo) {
		toSerialize["confirmNo"] = o.ConfirmNo
	}
	if !common.IsNil(o.WalletType) {
		toSerialize["walletType"] = o.WalletType
	}
	if !common.IsNil(o.TxKey) {
		toSerialize["txKey"] = o.TxKey
	}
	if !common.IsNil(o.CompleteTime) {
		toSerialize["completeTime"] = o.CompleteTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WithdrawHistoryResponseInner) UnmarshalJSON(data []byte) (err error) {
	varWithdrawHistoryResponseInner := _WithdrawHistoryResponseInner{}

	err = json.Unmarshal(data, &varWithdrawHistoryResponseInner)

	if err != nil {
		return err
	}

	*o = WithdrawHistoryResponseInner(varWithdrawHistoryResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "transactionFee")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "status")
		delete(additionalProperties, "address")
		delete(additionalProperties, "txId")
		delete(additionalProperties, "applyTime")
		delete(additionalProperties, "network")
		delete(additionalProperties, "transferType")
		delete(additionalProperties, "withdrawOrderId")
		delete(additionalProperties, "info")
		delete(additionalProperties, "confirmNo")
		delete(additionalProperties, "walletType")
		delete(additionalProperties, "txKey")
		delete(additionalProperties, "completeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWithdrawHistoryResponseInner struct {
	value *WithdrawHistoryResponseInner
	isSet bool
}

func (v NullableWithdrawHistoryResponseInner) Get() *WithdrawHistoryResponseInner {
	return v.value
}

func (v *NullableWithdrawHistoryResponseInner) Set(val *WithdrawHistoryResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawHistoryResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawHistoryResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawHistoryResponseInner(val *WithdrawHistoryResponseInner) *NullableWithdrawHistoryResponseInner {
	return &NullableWithdrawHistoryResponseInner{value: val, isSet: true}
}

func (v NullableWithdrawHistoryResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawHistoryResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
