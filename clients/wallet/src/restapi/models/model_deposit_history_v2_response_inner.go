/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DepositHistoryV2ResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DepositHistoryV2ResponseInner{}

// DepositHistoryV2ResponseInner struct for DepositHistoryV2ResponseInner
type DepositHistoryV2ResponseInner struct {
	DepositId            *string                                     `json:"depositId,omitempty"`
	Amount               *string                                     `json:"amount,omitempty"`
	Network              *string                                     `json:"network,omitempty"`
	Coin                 *string                                     `json:"coin,omitempty"`
	DepositStatus        *int64                                      `json:"depositStatus,omitempty"`
	TravelRuleReqStatus  *int64                                      `json:"travelRuleReqStatus,omitempty"`
	Address              *string                                     `json:"address,omitempty"`
	AddressTag           *string                                     `json:"addressTag,omitempty"`
	TxId                 *string                                     `json:"txId,omitempty"`
	TransferType         *int64                                      `json:"transferType,omitempty"`
	ConfirmTimes         *string                                     `json:"confirmTimes,omitempty"`
	RequireQuestionnaire *bool                                       `json:"requireQuestionnaire,omitempty"`
	Questionnaire        *DepositHistoryV2ResponseInnerQuestionnaire `json:"questionnaire,omitempty"`
	InsertTime           *int64                                      `json:"insertTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DepositHistoryV2ResponseInner DepositHistoryV2ResponseInner

// NewDepositHistoryV2ResponseInner instantiates a new DepositHistoryV2ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositHistoryV2ResponseInner() *DepositHistoryV2ResponseInner {
	this := DepositHistoryV2ResponseInner{}
	return &this
}

// NewDepositHistoryV2ResponseInnerWithDefaults instantiates a new DepositHistoryV2ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositHistoryV2ResponseInnerWithDefaults() *DepositHistoryV2ResponseInner {
	this := DepositHistoryV2ResponseInner{}
	return &this
}

// GetDepositId returns the DepositId field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetDepositId() string {
	if o == nil || common.IsNil(o.DepositId) {
		var ret string
		return ret
	}
	return *o.DepositId
}

// GetDepositIdOk returns a tuple with the DepositId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetDepositIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.DepositId) {
		return nil, false
	}
	return o.DepositId, true
}

// HasDepositId returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasDepositId() bool {
	if o != nil && !common.IsNil(o.DepositId) {
		return true
	}

	return false
}

// SetDepositId gets a reference to the given string and assigns it to the DepositId field.
func (o *DepositHistoryV2ResponseInner) SetDepositId(v string) {
	o.DepositId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *DepositHistoryV2ResponseInner) SetAmount(v string) {
	o.Amount = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetNetwork() string {
	if o == nil || common.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetNetworkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasNetwork() bool {
	if o != nil && !common.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *DepositHistoryV2ResponseInner) SetNetwork(v string) {
	o.Network = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetCoin() string {
	if o == nil || common.IsNil(o.Coin) {
		var ret string
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetCoinOk() (*string, bool) {
	if o == nil || common.IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasCoin() bool {
	if o != nil && !common.IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given string and assigns it to the Coin field.
func (o *DepositHistoryV2ResponseInner) SetCoin(v string) {
	o.Coin = &v
}

// GetDepositStatus returns the DepositStatus field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetDepositStatus() int64 {
	if o == nil || common.IsNil(o.DepositStatus) {
		var ret int64
		return ret
	}
	return *o.DepositStatus
}

// GetDepositStatusOk returns a tuple with the DepositStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetDepositStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.DepositStatus) {
		return nil, false
	}
	return o.DepositStatus, true
}

// HasDepositStatus returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasDepositStatus() bool {
	if o != nil && !common.IsNil(o.DepositStatus) {
		return true
	}

	return false
}

// SetDepositStatus gets a reference to the given int64 and assigns it to the DepositStatus field.
func (o *DepositHistoryV2ResponseInner) SetDepositStatus(v int64) {
	o.DepositStatus = &v
}

// GetTravelRuleReqStatus returns the TravelRuleReqStatus field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetTravelRuleReqStatus() int64 {
	if o == nil || common.IsNil(o.TravelRuleReqStatus) {
		var ret int64
		return ret
	}
	return *o.TravelRuleReqStatus
}

// GetTravelRuleReqStatusOk returns a tuple with the TravelRuleReqStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetTravelRuleReqStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TravelRuleReqStatus) {
		return nil, false
	}
	return o.TravelRuleReqStatus, true
}

// HasTravelRuleReqStatus returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasTravelRuleReqStatus() bool {
	if o != nil && !common.IsNil(o.TravelRuleReqStatus) {
		return true
	}

	return false
}

// SetTravelRuleReqStatus gets a reference to the given int64 and assigns it to the TravelRuleReqStatus field.
func (o *DepositHistoryV2ResponseInner) SetTravelRuleReqStatus(v int64) {
	o.TravelRuleReqStatus = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetAddress() string {
	if o == nil || common.IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasAddress() bool {
	if o != nil && !common.IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DepositHistoryV2ResponseInner) SetAddress(v string) {
	o.Address = &v
}

// GetAddressTag returns the AddressTag field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetAddressTag() string {
	if o == nil || common.IsNil(o.AddressTag) {
		var ret string
		return ret
	}
	return *o.AddressTag
}

// GetAddressTagOk returns a tuple with the AddressTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetAddressTagOk() (*string, bool) {
	if o == nil || common.IsNil(o.AddressTag) {
		return nil, false
	}
	return o.AddressTag, true
}

// HasAddressTag returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasAddressTag() bool {
	if o != nil && !common.IsNil(o.AddressTag) {
		return true
	}

	return false
}

// SetAddressTag gets a reference to the given string and assigns it to the AddressTag field.
func (o *DepositHistoryV2ResponseInner) SetAddressTag(v string) {
	o.AddressTag = &v
}

// GetTxId returns the TxId field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetTxId() string {
	if o == nil || common.IsNil(o.TxId) {
		var ret string
		return ret
	}
	return *o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetTxIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TxId) {
		return nil, false
	}
	return o.TxId, true
}

// HasTxId returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasTxId() bool {
	if o != nil && !common.IsNil(o.TxId) {
		return true
	}

	return false
}

// SetTxId gets a reference to the given string and assigns it to the TxId field.
func (o *DepositHistoryV2ResponseInner) SetTxId(v string) {
	o.TxId = &v
}

// GetTransferType returns the TransferType field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetTransferType() int64 {
	if o == nil || common.IsNil(o.TransferType) {
		var ret int64
		return ret
	}
	return *o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetTransferTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransferType) {
		return nil, false
	}
	return o.TransferType, true
}

// HasTransferType returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasTransferType() bool {
	if o != nil && !common.IsNil(o.TransferType) {
		return true
	}

	return false
}

// SetTransferType gets a reference to the given int64 and assigns it to the TransferType field.
func (o *DepositHistoryV2ResponseInner) SetTransferType(v int64) {
	o.TransferType = &v
}

// GetConfirmTimes returns the ConfirmTimes field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetConfirmTimes() string {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		var ret string
		return ret
	}
	return *o.ConfirmTimes
}

// GetConfirmTimesOk returns a tuple with the ConfirmTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetConfirmTimesOk() (*string, bool) {
	if o == nil || common.IsNil(o.ConfirmTimes) {
		return nil, false
	}
	return o.ConfirmTimes, true
}

// HasConfirmTimes returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasConfirmTimes() bool {
	if o != nil && !common.IsNil(o.ConfirmTimes) {
		return true
	}

	return false
}

// SetConfirmTimes gets a reference to the given string and assigns it to the ConfirmTimes field.
func (o *DepositHistoryV2ResponseInner) SetConfirmTimes(v string) {
	o.ConfirmTimes = &v
}

// GetRequireQuestionnaire returns the RequireQuestionnaire field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetRequireQuestionnaire() bool {
	if o == nil || common.IsNil(o.RequireQuestionnaire) {
		var ret bool
		return ret
	}
	return *o.RequireQuestionnaire
}

// GetRequireQuestionnaireOk returns a tuple with the RequireQuestionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetRequireQuestionnaireOk() (*bool, bool) {
	if o == nil || common.IsNil(o.RequireQuestionnaire) {
		return nil, false
	}
	return o.RequireQuestionnaire, true
}

// HasRequireQuestionnaire returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasRequireQuestionnaire() bool {
	if o != nil && !common.IsNil(o.RequireQuestionnaire) {
		return true
	}

	return false
}

// SetRequireQuestionnaire gets a reference to the given bool and assigns it to the RequireQuestionnaire field.
func (o *DepositHistoryV2ResponseInner) SetRequireQuestionnaire(v bool) {
	o.RequireQuestionnaire = &v
}

// GetQuestionnaire returns the Questionnaire field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetQuestionnaire() DepositHistoryV2ResponseInnerQuestionnaire {
	if o == nil || common.IsNil(o.Questionnaire) {
		var ret DepositHistoryV2ResponseInnerQuestionnaire
		return ret
	}
	return *o.Questionnaire
}

// GetQuestionnaireOk returns a tuple with the Questionnaire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetQuestionnaireOk() (*DepositHistoryV2ResponseInnerQuestionnaire, bool) {
	if o == nil || common.IsNil(o.Questionnaire) {
		return nil, false
	}
	return o.Questionnaire, true
}

// HasQuestionnaire returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasQuestionnaire() bool {
	if o != nil && !common.IsNil(o.Questionnaire) {
		return true
	}

	return false
}

// SetQuestionnaire gets a reference to the given DepositHistoryV2ResponseInnerQuestionnaire and assigns it to the Questionnaire field.
func (o *DepositHistoryV2ResponseInner) SetQuestionnaire(v DepositHistoryV2ResponseInnerQuestionnaire) {
	o.Questionnaire = &v
}

// GetInsertTime returns the InsertTime field value if set, zero value otherwise.
func (o *DepositHistoryV2ResponseInner) GetInsertTime() int64 {
	if o == nil || common.IsNil(o.InsertTime) {
		var ret int64
		return ret
	}
	return *o.InsertTime
}

// GetInsertTimeOk returns a tuple with the InsertTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositHistoryV2ResponseInner) GetInsertTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.InsertTime) {
		return nil, false
	}
	return o.InsertTime, true
}

// HasInsertTime returns a boolean if a field has been set.
func (o *DepositHistoryV2ResponseInner) HasInsertTime() bool {
	if o != nil && !common.IsNil(o.InsertTime) {
		return true
	}

	return false
}

// SetInsertTime gets a reference to the given int64 and assigns it to the InsertTime field.
func (o *DepositHistoryV2ResponseInner) SetInsertTime(v int64) {
	o.InsertTime = &v
}

func (o DepositHistoryV2ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositHistoryV2ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.DepositId) {
		toSerialize["depositId"] = o.DepositId
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !common.IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !common.IsNil(o.DepositStatus) {
		toSerialize["depositStatus"] = o.DepositStatus
	}
	if !common.IsNil(o.TravelRuleReqStatus) {
		toSerialize["travelRuleReqStatus"] = o.TravelRuleReqStatus
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
	if !common.IsNil(o.TransferType) {
		toSerialize["transferType"] = o.TransferType
	}
	if !common.IsNil(o.ConfirmTimes) {
		toSerialize["confirmTimes"] = o.ConfirmTimes
	}
	if !common.IsNil(o.RequireQuestionnaire) {
		toSerialize["requireQuestionnaire"] = o.RequireQuestionnaire
	}
	if !common.IsNil(o.Questionnaire) {
		toSerialize["questionnaire"] = o.Questionnaire
	}
	if !common.IsNil(o.InsertTime) {
		toSerialize["insertTime"] = o.InsertTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DepositHistoryV2ResponseInner) UnmarshalJSON(data []byte) (err error) {
	varDepositHistoryV2ResponseInner := _DepositHistoryV2ResponseInner{}

	err = json.Unmarshal(data, &varDepositHistoryV2ResponseInner)

	if err != nil {
		return err
	}

	*o = DepositHistoryV2ResponseInner(varDepositHistoryV2ResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "depositId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "network")
		delete(additionalProperties, "coin")
		delete(additionalProperties, "depositStatus")
		delete(additionalProperties, "travelRuleReqStatus")
		delete(additionalProperties, "address")
		delete(additionalProperties, "addressTag")
		delete(additionalProperties, "txId")
		delete(additionalProperties, "transferType")
		delete(additionalProperties, "confirmTimes")
		delete(additionalProperties, "requireQuestionnaire")
		delete(additionalProperties, "questionnaire")
		delete(additionalProperties, "insertTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositHistoryV2ResponseInner struct {
	value *DepositHistoryV2ResponseInner
	isSet bool
}

func (v NullableDepositHistoryV2ResponseInner) Get() *DepositHistoryV2ResponseInner {
	return v.value
}

func (v *NullableDepositHistoryV2ResponseInner) Set(val *DepositHistoryV2ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositHistoryV2ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositHistoryV2ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositHistoryV2ResponseInner(val *DepositHistoryV2ResponseInner) *NullableDepositHistoryV2ResponseInner {
	return &NullableDepositHistoryV2ResponseInner{value: val, isSet: true}
}

func (v NullableDepositHistoryV2ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositHistoryV2ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
