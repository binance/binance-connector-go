/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the WithdrawHistoryV2ResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &WithdrawHistoryV2ResponseInner{}

// WithdrawHistoryV2ResponseInner struct for WithdrawHistoryV2ResponseInner
type WithdrawHistoryV2ResponseInner struct {
	Id                   *string `json:"id,omitempty"`
	TrId                 *int64  `json:"trId,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TransactionFee       *string `json:"transactionFee,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	WithdrawalStatus     *int64  `json:"withdrawalStatus,omitempty"`
	TravelRuleStatus     *int64  `json:"travelRuleStatus,omitempty"`
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
	Questionnaire        *string `json:"questionnaire,omitempty"`
	CompleteTime         *string `json:"completeTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WithdrawHistoryV2ResponseInner WithdrawHistoryV2ResponseInner

// NewWithdrawHistoryV2ResponseInner instantiates a new WithdrawHistoryV2ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawHistoryV2ResponseInner() *WithdrawHistoryV2ResponseInner {
	this := WithdrawHistoryV2ResponseInner{}
	return &this
}

// NewWithdrawHistoryV2ResponseInnerWithDefaults instantiates a new WithdrawHistoryV2ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawHistoryV2ResponseInnerWithDefaults() *WithdrawHistoryV2ResponseInner {
	this := WithdrawHistoryV2ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WithdrawHistoryV2ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetTrId returns the TrId field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTrId() int64 {
	if o == nil || common.IsNil(o.TrId) {
		var ret int64
		return ret
	}
	return *o.TrId
}

// GetTrIdOk returns a tuple with the TrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTrIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrId) {
		return nil, false
	}
	return o.TrId, true
}

// HasTrId returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTrId() bool {
	if o != nil && !common.IsNil(o.TrId) {
		return true
	}

	return false
}

// SetTrId gets a reference to the given int64 and assigns it to the TrId field.
func (o *WithdrawHistoryV2ResponseInner) SetTrId(v int64) {
	o.TrId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *WithdrawHistoryV2ResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTransactionFee returns the TransactionFee field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTransactionFee() string {
	if o == nil || common.IsNil(o.TransactionFee) {
		var ret string
		return ret
	}
	return *o.TransactionFee
}

// GetTransactionFeeOk returns a tuple with the TransactionFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTransactionFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionFee) {
		return nil, false
	}
	return o.TransactionFee, true
}

// HasTransactionFee returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTransactionFee() bool {
	if o != nil && !common.IsNil(o.TransactionFee) {
		return true
	}

	return false
}

// SetTransactionFee gets a reference to the given string and assigns it to the TransactionFee field.
func (o *WithdrawHistoryV2ResponseInner) SetTransactionFee(v string) {
	o.TransactionFee = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *WithdrawHistoryV2ResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetWithdrawalStatus returns the WithdrawalStatus field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetWithdrawalStatus() int64 {
	if o == nil || common.IsNil(o.WithdrawalStatus) {
		var ret int64
		return ret
	}
	return *o.WithdrawalStatus
}

// GetWithdrawalStatusOk returns a tuple with the WithdrawalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetWithdrawalStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WithdrawalStatus) {
		return nil, false
	}
	return o.WithdrawalStatus, true
}

// HasWithdrawalStatus returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasWithdrawalStatus() bool {
	if o != nil && !common.IsNil(o.WithdrawalStatus) {
		return true
	}

	return false
}

// SetWithdrawalStatus gets a reference to the given int64 and assigns it to the WithdrawalStatus field.
func (o *WithdrawHistoryV2ResponseInner) SetWithdrawalStatus(v int64) {
	o.WithdrawalStatus = &v
}

// GetTravelRuleStatus returns the TravelRuleStatus field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTravelRuleStatus() int64 {
	if o == nil || common.IsNil(o.TravelRuleStatus) {
		var ret int64
		return ret
	}
	return *o.TravelRuleStatus
}

// GetTravelRuleStatusOk returns a tuple with the TravelRuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTravelRuleStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TravelRuleStatus) {
		return nil, false
	}
	return o.TravelRuleStatus, true
}

// HasTravelRuleStatus returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTravelRuleStatus() bool {
	if o != nil && !common.IsNil(o.TravelRuleStatus) {
		return true
	}

	return false
}

// SetTravelRuleStatus gets a reference to the given int64 and assigns it to the TravelRuleStatus field.
func (o *WithdrawHistoryV2ResponseInner) SetTravelRuleStatus(v int64) {
	o.TravelRuleStatus = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *WithdrawHistoryV2ResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTxId() string {
	if o == nil || common.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *WithdrawHistoryV2ResponseInner) SetTxId(v string) {
	o.TxId = &v
}

// GetApplyTime returns the ApplyTime field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetApplyTime() string {
	if o == nil || common.IsNil(o.ApplyTime) {
		var ret string
		return ret
	}
	return *o.ApplyTime
}

// GetApplyTimeOk returns a tuple with the ApplyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetApplyTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.ApplyTime) {
		return nil, false
	}
	return o.ApplyTime, true
}

// HasApplyTime returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasApplyTime() bool {
	if o != nil && !common.IsNil(o.ApplyTime) {
		return true
	}

	return false
}

// SetApplyTime gets a reference to the given string and assigns it to the ApplyTime field.
func (o *WithdrawHistoryV2ResponseInner) SetApplyTime(v string) {
	o.ApplyTime = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *WithdrawHistoryV2ResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTransferType() int64 {
	if o == nil || common.IsNil(o.TransferType) {
		var ret int64
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTransferTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given int64 and assigns it to the TransferType field.
func (o *WithdrawHistoryV2ResponseInner) SetTransferType(v int64) {
	o.TransferType = &v
}

// GetWithdrawOrderId returns the WithdrawOrderId field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetWithdrawOrderId() string {
	if o == nil || common.IsNil(o.WithdrawOrderId) {
		var ret string
		return ret
	}
	return *o.WithdrawOrderId
}

// GetWithdrawOrderIdOk returns a tuple with the WithdrawOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetWithdrawOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WithdrawOrderId) {
		return nil, false
	}
	return o.WithdrawOrderId, true
}

// HasWithdrawOrderId returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasWithdrawOrderId() bool {
	if o != nil && !common.IsNil(o.WithdrawOrderId) {
		return true
	}

	return false
}

// SetWithdrawOrderId gets a reference to the given string and assigns it to the WithdrawOrderId field.
func (o *WithdrawHistoryV2ResponseInner) SetWithdrawOrderId(v string) {
	o.WithdrawOrderId = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetInfo() string {
	if o == nil || common.IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetInfoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasInfo() bool {
	if o != nil && !common.IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *WithdrawHistoryV2ResponseInner) SetInfo(v string) {
	o.Info = &v
}

// GetConfirmNo returns the ConfirmNo field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetConfirmNo() int64 {
	if o == nil || common.IsNil(o.ConfirmNo) {
		var ret int64
		return ret
	}
	return *o.ConfirmNo
}

// GetConfirmNoOk returns a tuple with the ConfirmNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetConfirmNoOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ConfirmNo) {
		return nil, false
	}
	return o.ConfirmNo, true
}

// HasConfirmNo returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasConfirmNo() bool {
	if o != nil && !common.IsNil(o.ConfirmNo) {
		return true
	}

	return false
}

// SetConfirmNo gets a reference to the given int64 and assigns it to the ConfirmNo field.
func (o *WithdrawHistoryV2ResponseInner) SetConfirmNo(v int64) {
	o.ConfirmNo = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetWalletType() int64 {
	if o == nil || common.IsNil(o.WalletType) {
		var ret int64
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetWalletTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasWalletType() bool {
	if o != nil && !common.IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given int64 and assigns it to the WalletType field.
func (o *WithdrawHistoryV2ResponseInner) SetWalletType(v int64) {
	o.WalletType = &v
}

// GetTxKey returns the TxKey field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetTxKey() string {
	if o == nil || common.IsNil(o.TxKey) {
		var ret string
		return ret
	}
	return *o.TxKey
}

// GetTxKeyOk returns a tuple with the TxKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetTxKeyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxKey) {
		return nil, false
	}
	return o.TxKey, true
}

// HasTxKey returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasTxKey() bool {
	if o != nil && !common.IsNil(o.TxKey) {
		return true
	}

	return false
}

// SetTxKey gets a reference to the given string and assigns it to the TxKey field.
func (o *WithdrawHistoryV2ResponseInner) SetTxKey(v string) {
	o.TxKey = &v
}

// GetQuestionnaire returns the Questionnaire field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetQuestionnaire() string {
	if o == nil || common.IsNil(o.Questionnaire) {
		var ret string
		return ret
	}
	return *o.Questionnaire
}

// GetQuestionnaireOk returns a tuple with the Questionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetQuestionnaireOk() (*string, bool) {
	if o == nil || common.IsNil(o.Questionnaire) {
		return nil, false
	}
	return o.Questionnaire, true
}

// HasQuestionnaire returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasQuestionnaire() bool {
	if o != nil && !common.IsNil(o.Questionnaire) {
		return true
	}

	return false
}

// SetQuestionnaire gets a reference to the given string and assigns it to the Questionnaire field.
func (o *WithdrawHistoryV2ResponseInner) SetQuestionnaire(v string) {
	o.Questionnaire = &v
}

// GetCompleteTime returns the CompleteTime field value if set, zero value otherwise.
func (o *WithdrawHistoryV2ResponseInner) GetCompleteTime() string {
	if o == nil || common.IsNil(o.CompleteTime) {
		var ret string
		return ret
	}
	return *o.CompleteTime
}

// GetCompleteTimeOk returns a tuple with the CompleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawHistoryV2ResponseInner) GetCompleteTimeOk() (*string, bool) {
	if o == nil || common.IsNil(o.CompleteTime) {
		return nil, false
	}
	return o.CompleteTime, true
}

// HasCompleteTime returns a boolean if a field has been set.
func (o *WithdrawHistoryV2ResponseInner) HasCompleteTime() bool {
	if o != nil && !common.IsNil(o.CompleteTime) {
		return true
	}

	return false
}

// SetCompleteTime gets a reference to the given string and assigns it to the CompleteTime field.
func (o *WithdrawHistoryV2ResponseInner) SetCompleteTime(v string) {
	o.CompleteTime = &v
}

func (o WithdrawHistoryV2ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawHistoryV2ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TrId) {
		toSerialize["trId"] = o.TrId
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
	if !common.IsNil(o.WithdrawalStatus) {
		toSerialize["withdrawalStatus"] = o.WithdrawalStatus
	}
	if !common.IsNil(o.TravelRuleStatus) {
		toSerialize["travelRuleStatus"] = o.TravelRuleStatus
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
	if !common.IsNil(o.Questionnaire) {
		toSerialize["questionnaire"] = o.Questionnaire
	}
	if !common.IsNil(o.CompleteTime) {
		toSerialize["completeTime"] = o.CompleteTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WithdrawHistoryV2ResponseInner) UnmarshalJSON(data []byte) (err error) {
	varWithdrawHistoryV2ResponseInner := _WithdrawHistoryV2ResponseInner{}

	err = json.Unmarshal(data, &varWithdrawHistoryV2ResponseInner)

	if err != nil {
		return err
	}

	*o = WithdrawHistoryV2ResponseInner(varWithdrawHistoryV2ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "trId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "transactionFee")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "withdrawalStatus")
		delete(additionalProperties, "travelRuleStatus")
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
		delete(additionalProperties, "questionnaire")
		delete(additionalProperties, "completeTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWithdrawHistoryV2ResponseInner struct {
	value *WithdrawHistoryV2ResponseInner
	isSet bool
}

func (v NullableWithdrawHistoryV2ResponseInner) Get() *WithdrawHistoryV2ResponseInner {
	return v.value
}

func (v *NullableWithdrawHistoryV2ResponseInner) Set(val *WithdrawHistoryV2ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawHistoryV2ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawHistoryV2ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawHistoryV2ResponseInner(val *WithdrawHistoryV2ResponseInner) *NullableWithdrawHistoryV2ResponseInner {
	return &NullableWithdrawHistoryV2ResponseInner{value: val, isSet: true}
}

func (v NullableWithdrawHistoryV2ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawHistoryV2ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
