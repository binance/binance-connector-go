/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DepositHistoryTravelRuleResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryTravelRuleResponseInner{}

// DepositHistoryTravelRuleResponseInner struct for DepositHistoryTravelRuleResponseInner
type DepositHistoryTravelRuleResponseInner struct {
	TrId                 *int64  `json:"trId,omitempty"`
	TranId               *int64  `json:"tranId,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	Coin                 *string `json:"coin,omitempty"`
	Network              *string `json:"network,omitempty"`
	DepositStatus        *int64  `json:"depositStatus,omitempty"`
	TravelRuleStatus     *int64  `json:"travelRuleStatus,omitempty"`
	Address              *string `json:"address,omitempty"`
	AddressTag           *string `json:"addressTag,omitempty"`
	TxId                 *string `json:"txId,omitempty"`
	InsertTime           *int64  `json:"insertTime,omitempty"`
	TransferType         *int64  `json:"transferType,omitempty"`
	ConfirmTimes         *string `json:"confirmTimes,omitempty"`
	UnlockConfirm        *int64  `json:"unlockConfirm,omitempty"`
	WalletType           *int64  `json:"walletType,omitempty"`
	RequireQuestionnaire *bool   `json:"requireQuestionnaire,omitempty"`
	Questionnaire        *string `json:"questionnaire,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositHistoryTravelRuleResponseInner DepositHistoryTravelRuleResponseInner

// NewDepositHistoryTravelRuleResponseInner instantiates a new DepositHistoryTravelRuleResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryTravelRuleResponseInner() *DepositHistoryTravelRuleResponseInner {
	this := DepositHistoryTravelRuleResponseInner{}
	return &this
}

// NewDepositHistoryTravelRuleResponseInnerWithDefaults instantiates a new DepositHistoryTravelRuleResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryTravelRuleResponseInnerWithDefaults() *DepositHistoryTravelRuleResponseInner {
	this := DepositHistoryTravelRuleResponseInner{}
	return &this
}

// GetTrId returns the TrId field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetTrId() int64 {
	if o == nil || common.IsNil(o.TrId) {
		var ret int64
		return ret
	}
	return *o.TrId
}

// GetTrIdOk returns a tuple with the TrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetTrIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TrId) {
		return nil, false
	}
	return o.TrId, true
}

// HasTrId returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasTrId() bool {
	if o != nil && !common.IsNil(o.TrId) {
		return true
	}

	return false
}

// SetTrId gets a reference to the given int64 and assigns it to the TrId field.
func (o *DepositHistoryTravelRuleResponseInner) SetTrId(v int64) {
	o.TrId = &v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetTranId() int64 {
	if o == nil || common.IsNil(o.TranId) {
		var ret int64
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetTranIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasTranId() bool {
	if o != nil && !common.IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given int64 and assigns it to the TranId field.
func (o *DepositHistoryTravelRuleResponseInner) SetTranId(v int64) {
	o.TranId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *DepositHistoryTravelRuleResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *DepositHistoryTravelRuleResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *DepositHistoryTravelRuleResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetDepositStatus returns the DepositStatus field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetDepositStatus() int64 {
	if o == nil || common.IsNil(o.DepositStatus) {
		var ret int64
		return ret
	}
	return *o.DepositStatus
}

// GetDepositStatusOk returns a tuple with the DepositStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetDepositStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DepositStatus) {
		return nil, false
	}
	return o.DepositStatus, true
}

// HasDepositStatus returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasDepositStatus() bool {
	if o != nil && !common.IsNil(o.DepositStatus) {
		return true
	}

	return false
}

// SetDepositStatus gets a reference to the given int64 and assigns it to the DepositStatus field.
func (o *DepositHistoryTravelRuleResponseInner) SetDepositStatus(v int64) {
	o.DepositStatus = &v
}

// GetTravelRuleStatus returns the TravelRuleStatus field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetTravelRuleStatus() int64 {
	if o == nil || common.IsNil(o.TravelRuleStatus) {
		var ret int64
		return ret
	}
	return *o.TravelRuleStatus
}

// GetTravelRuleStatusOk returns a tuple with the TravelRuleStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetTravelRuleStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TravelRuleStatus) {
		return nil, false
	}
	return o.TravelRuleStatus, true
}

// HasTravelRuleStatus returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasTravelRuleStatus() bool {
	if o != nil && !common.IsNil(o.TravelRuleStatus) {
		return true
	}

	return false
}

// SetTravelRuleStatus gets a reference to the given int64 and assigns it to the TravelRuleStatus field.
func (o *DepositHistoryTravelRuleResponseInner) SetTravelRuleStatus(v int64) {
	o.TravelRuleStatus = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DepositHistoryTravelRuleResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetAddressTag returns the AddressTag field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetAddressTag() string {
	if o == nil || common.IsNil(o.AddressTag) {
		var ret string
		return ret
	}
	return *o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetAddressTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.AddressTag) {
		return nil, false
	}
	return o.AddressTag, true
}

// HasAddressTag returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasAddressTag() bool {
	if o != nil && !common.IsNil(o.AddressTag) {
		return true
	}

	return false
}

// SetAddressTag gets a reference to the given string and assigns it to the AddressTag field.
func (o *DepositHistoryTravelRuleResponseInner) SetAddressTag(v string) {
	o.AddressTag = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetTxId() string {
	if o == nil || common.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetTxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *DepositHistoryTravelRuleResponseInner) SetTxId(v string) {
	o.TxId = &v
}

// GetInsertTime returns the InsertTime field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetInsertTime() int64 {
	if o == nil || common.IsNil(o.InsertTime) {
		var ret int64
		return ret
	}
	return *o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetInsertTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InsertTime) {
		return nil, false
	}
	return o.InsertTime, true
}

// HasInsertTime returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasInsertTime() bool {
	if o != nil && !common.IsNil(o.InsertTime) {
		return true
	}

	return false
}

// SetInsertTime gets a reference to the given int64 and assigns it to the InsertTime field.
func (o *DepositHistoryTravelRuleResponseInner) SetInsertTime(v int64) {
	o.InsertTime = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetTransferType() int64 {
	if o == nil || common.IsNil(o.TransferType) {
		var ret int64
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetTransferTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given int64 and assigns it to the TransferType field.
func (o *DepositHistoryTravelRuleResponseInner) SetTransferType(v int64) {
	o.TransferType = &v
}

// GetConfirmTimes returns the ConfirmTimes field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetConfirmTimes() string {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		var ret string
		return ret
	}
	return *o.ConfirmTimes
}

// GetConfirmTimesOk returns a tuple with the ConfirmTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetConfirmTimesOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		return nil, false
	}
	return o.ConfirmTimes, true
}

// HasConfirmTimes returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasConfirmTimes() bool {
	if o != nil && !common.IsNil(o.ConfirmTimes) {
		return true
	}

	return false
}

// SetConfirmTimes gets a reference to the given string and assigns it to the ConfirmTimes field.
func (o *DepositHistoryTravelRuleResponseInner) SetConfirmTimes(v string) {
	o.ConfirmTimes = &v
}

// GetUnlockConfirm returns the UnlockConfirm field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetUnlockConfirm() int64 {
	if o == nil || common.IsNil(o.UnlockConfirm) {
		var ret int64
		return ret
	}
	return *o.UnlockConfirm
}

// GetUnlockConfirmOk returns a tuple with the UnlockConfirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetUnlockConfirmOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UnlockConfirm) {
		return nil, false
	}
	return o.UnlockConfirm, true
}

// HasUnlockConfirm returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasUnlockConfirm() bool {
	if o != nil && !common.IsNil(o.UnlockConfirm) {
		return true
	}

	return false
}

// SetUnlockConfirm gets a reference to the given int64 and assigns it to the UnlockConfirm field.
func (o *DepositHistoryTravelRuleResponseInner) SetUnlockConfirm(v int64) {
	o.UnlockConfirm = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetWalletType() int64 {
	if o == nil || common.IsNil(o.WalletType) {
		var ret int64
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetWalletTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasWalletType() bool {
	if o != nil && !common.IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given int64 and assigns it to the WalletType field.
func (o *DepositHistoryTravelRuleResponseInner) SetWalletType(v int64) {
	o.WalletType = &v
}

// GetRequireQuestionnaire returns the RequireQuestionnaire field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetRequireQuestionnaire() bool {
	if o == nil || common.IsNil(o.RequireQuestionnaire) {
		var ret bool
		return ret
	}
	return *o.RequireQuestionnaire
}

// GetRequireQuestionnaireOk returns a tuple with the RequireQuestionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetRequireQuestionnaireOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequireQuestionnaire) {
		return nil, false
	}
	return o.RequireQuestionnaire, true
}

// HasRequireQuestionnaire returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasRequireQuestionnaire() bool {
	if o != nil && !common.IsNil(o.RequireQuestionnaire) {
		return true
	}

	return false
}

// SetRequireQuestionnaire gets a reference to the given bool and assigns it to the RequireQuestionnaire field.
func (o *DepositHistoryTravelRuleResponseInner) SetRequireQuestionnaire(v bool) {
	o.RequireQuestionnaire = &v
}

// GetQuestionnaire returns the Questionnaire field value if set, zero value otherwise.
func (o *DepositHistoryTravelRuleResponseInner) GetQuestionnaire() string {
	if o == nil || common.IsNil(o.Questionnaire) {
		var ret string
		return ret
	}
	return *o.Questionnaire
}

// GetQuestionnaireOk returns a tuple with the Questionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryTravelRuleResponseInner) GetQuestionnaireOk() (*string, bool) {
	if o == nil || common.IsNil(o.Questionnaire) {
		return nil, false
	}
	return o.Questionnaire, true
}

// HasQuestionnaire returns a boolean if a field has been set.
func (o *DepositHistoryTravelRuleResponseInner) HasQuestionnaire() bool {
	if o != nil && !common.IsNil(o.Questionnaire) {
		return true
	}

	return false
}

// SetQuestionnaire gets a reference to the given string and assigns it to the Questionnaire field.
func (o *DepositHistoryTravelRuleResponseInner) SetQuestionnaire(v string) {
	o.Questionnaire = &v
}

func (o DepositHistoryTravelRuleResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryTravelRuleResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TrId) {
		toSerialize["trId"] = o.TrId
	}
	if !common.IsNil(o.TranId) {
		toSerialize["tranId"] = o.TranId
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
	if !common.IsNil(o.DepositStatus) {
		toSerialize["depositStatus"] = o.DepositStatus
	}
	if !common.IsNil(o.TravelRuleStatus) {
		toSerialize["travelRuleStatus"] = o.TravelRuleStatus
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
	if !common.IsNil(o.RequireQuestionnaire) {
		toSerialize["requireQuestionnaire"] = o.RequireQuestionnaire
	}
	if !common.IsNil(o.Questionnaire) {
		toSerialize["questionnaire"] = o.Questionnaire
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositHistoryTravelRuleResponseInner) UnmarshalJSON(data []byte) (err error) {
	varDepositHistoryTravelRuleResponseInner := _DepositHistoryTravelRuleResponseInner{}

	err = json.Unmarshal(data, &varDepositHistoryTravelRuleResponseInner)

	if err != nil {
		return err
	}

	*o = DepositHistoryTravelRuleResponseInner(varDepositHistoryTravelRuleResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "trId")
		delete(additionalProperties, "tranId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "network")
		delete(additionalProperties, "depositStatus")
		delete(additionalProperties, "travelRuleStatus")
		delete(additionalProperties, "address")
		delete(additionalProperties, "addressTag")
		delete(additionalProperties, "txId")
		delete(additionalProperties, "insertTime")
		delete(additionalProperties, "transferType")
		delete(additionalProperties, "confirmTimes")
		delete(additionalProperties, "unlockConfirm")
		delete(additionalProperties, "walletType")
		delete(additionalProperties, "requireQuestionnaire")
		delete(additionalProperties, "questionnaire")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositHistoryTravelRuleResponseInner struct {
	value *DepositHistoryTravelRuleResponseInner
	isSet bool
}

func (v NullableDepositHistoryTravelRuleResponseInner) Get() *DepositHistoryTravelRuleResponseInner {
	return v.value
}

func (v *NullableDepositHistoryTravelRuleResponseInner) Set(val *DepositHistoryTravelRuleResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryTravelRuleResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryTravelRuleResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositHistoryTravelRuleResponseInner(val *DepositHistoryTravelRuleResponseInner) *NullableDepositHistoryTravelRuleResponseInner {
	return &NullableDepositHistoryTravelRuleResponseInner{value: val, isSet: true}
}

func (v NullableDepositHistoryTravelRuleResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryTravelRuleResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
