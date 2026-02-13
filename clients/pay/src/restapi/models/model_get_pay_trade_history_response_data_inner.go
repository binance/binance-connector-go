/*
Binance Pay REST API

OpenAPI Specification for the Binance Pay REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetPayTradeHistoryResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetPayTradeHistoryResponseDataInner{}

// GetPayTradeHistoryResponseDataInner struct for GetPayTradeHistoryResponseDataInner
type GetPayTradeHistoryResponseDataInner struct {
	OrderType            *string                                               `json:"orderType,omitempty"`
	TransactionId        *string                                               `json:"transactionId,omitempty"`
	TransactionTime      *int64                                                `json:"transactionTime,omitempty"`
	Amount               *string                                               `json:"amount,omitempty"`
	Currency             *string                                               `json:"currency,omitempty"`
	WalletType           *int64                                                `json:"walletType,omitempty"`
	WalletTypes          []int64                                               `json:"walletTypes,omitempty"`
	FundsDetail          []GetPayTradeHistoryResponseDataInnerFundsDetailInner `json:"fundsDetail,omitempty"`
	PayerInfo            *GetPayTradeHistoryResponseDataInnerPayerInfo         `json:"payerInfo,omitempty"`
	ReceiverInfo         *GetPayTradeHistoryResponseDataInnerReceiverInfo      `json:"receiverInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetPayTradeHistoryResponseDataInner GetPayTradeHistoryResponseDataInner

// NewGetPayTradeHistoryResponseDataInner instantiates a new GetPayTradeHistoryResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPayTradeHistoryResponseDataInner() *GetPayTradeHistoryResponseDataInner {
	this := GetPayTradeHistoryResponseDataInner{}
	return &this
}

// NewGetPayTradeHistoryResponseDataInnerWithDefaults instantiates a new GetPayTradeHistoryResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPayTradeHistoryResponseDataInnerWithDefaults() *GetPayTradeHistoryResponseDataInner {
	this := GetPayTradeHistoryResponseDataInner{}
	return &this
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *GetPayTradeHistoryResponseDataInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetTransactionId() string {
	if o == nil || common.IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetTransactionIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasTransactionId() bool {
	if o != nil && !common.IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *GetPayTradeHistoryResponseDataInner) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionTime returns the TransactionTime field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetTransactionTime() int64 {
	if o == nil || common.IsNil(o.TransactionTime) {
		var ret int64
		return ret
	}
	return *o.TransactionTime
}

// GetTransactionTimeOk returns a tuple with the TransactionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetTransactionTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactionTime) {
		return nil, false
	}
	return o.TransactionTime, true
}

// HasTransactionTime returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasTransactionTime() bool {
	if o != nil && !common.IsNil(o.TransactionTime) {
		return true
	}

	return false
}

// SetTransactionTime gets a reference to the given int64 and assigns it to the TransactionTime field.
func (o *GetPayTradeHistoryResponseDataInner) SetTransactionTime(v int64) {
	o.TransactionTime = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetPayTradeHistoryResponseDataInner) SetAmount(v string) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetCurrency() string {
	if o == nil || common.IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasCurrency() bool {
	if o != nil && !common.IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *GetPayTradeHistoryResponseDataInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetWalletType returns the WalletType field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetWalletType() int64 {
	if o == nil || common.IsNil(o.WalletType) {
		var ret int64
		return ret
	}
	return *o.WalletType
}

// GetWalletTypeOk returns a tuple with the WalletType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.WalletType) {
		return nil, false
	}
	return o.WalletType, true
}

// HasWalletType returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasWalletType() bool {
	if o != nil && !common.IsNil(o.WalletType) {
		return true
	}

	return false
}

// SetWalletType gets a reference to the given int64 and assigns it to the WalletType field.
func (o *GetPayTradeHistoryResponseDataInner) SetWalletType(v int64) {
	o.WalletType = &v
}

// GetWalletTypes returns the WalletTypes field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypes() []int64 {
	if o == nil || common.IsNil(o.WalletTypes) {
		var ret []int64
		return ret
	}
	return o.WalletTypes
}

// GetWalletTypesOk returns a tuple with the WalletTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypesOk() ([]int64, bool) {
	if o == nil || common.IsNil(o.WalletTypes) {
		return nil, false
	}
	return o.WalletTypes, true
}

// HasWalletTypes returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasWalletTypes() bool {
	if o != nil && !common.IsNil(o.WalletTypes) {
		return true
	}

	return false
}

// SetWalletTypes gets a reference to the given []int64 and assigns it to the WalletTypes field.
func (o *GetPayTradeHistoryResponseDataInner) SetWalletTypes(v []int64) {
	o.WalletTypes = v
}

// GetFundsDetail returns the FundsDetail field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetFundsDetail() []GetPayTradeHistoryResponseDataInnerFundsDetailInner {
	if o == nil || common.IsNil(o.FundsDetail) {
		var ret []GetPayTradeHistoryResponseDataInnerFundsDetailInner
		return ret
	}
	return o.FundsDetail
}

// GetFundsDetailOk returns a tuple with the FundsDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetFundsDetailOk() ([]GetPayTradeHistoryResponseDataInnerFundsDetailInner, bool) {
	if o == nil || common.IsNil(o.FundsDetail) {
		return nil, false
	}
	return o.FundsDetail, true
}

// HasFundsDetail returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasFundsDetail() bool {
	if o != nil && !common.IsNil(o.FundsDetail) {
		return true
	}

	return false
}

// SetFundsDetail gets a reference to the given []GetPayTradeHistoryResponseDataInnerFundsDetailInner and assigns it to the FundsDetail field.
func (o *GetPayTradeHistoryResponseDataInner) SetFundsDetail(v []GetPayTradeHistoryResponseDataInnerFundsDetailInner) {
	o.FundsDetail = v
}

// GetPayerInfo returns the PayerInfo field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetPayerInfo() GetPayTradeHistoryResponseDataInnerPayerInfo {
	if o == nil || common.IsNil(o.PayerInfo) {
		var ret GetPayTradeHistoryResponseDataInnerPayerInfo
		return ret
	}
	return *o.PayerInfo
}

// GetPayerInfoOk returns a tuple with the PayerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetPayerInfoOk() (*GetPayTradeHistoryResponseDataInnerPayerInfo, bool) {
	if o == nil || common.IsNil(o.PayerInfo) {
		return nil, false
	}
	return o.PayerInfo, true
}

// HasPayerInfo returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasPayerInfo() bool {
	if o != nil && !common.IsNil(o.PayerInfo) {
		return true
	}

	return false
}

// SetPayerInfo gets a reference to the given GetPayTradeHistoryResponseDataInnerPayerInfo and assigns it to the PayerInfo field.
func (o *GetPayTradeHistoryResponseDataInner) SetPayerInfo(v GetPayTradeHistoryResponseDataInnerPayerInfo) {
	o.PayerInfo = &v
}

// GetReceiverInfo returns the ReceiverInfo field value if set, zero value otherwise.
func (o *GetPayTradeHistoryResponseDataInner) GetReceiverInfo() GetPayTradeHistoryResponseDataInnerReceiverInfo {
	if o == nil || common.IsNil(o.ReceiverInfo) {
		var ret GetPayTradeHistoryResponseDataInnerReceiverInfo
		return ret
	}
	return *o.ReceiverInfo
}

// GetReceiverInfoOk returns a tuple with the ReceiverInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPayTradeHistoryResponseDataInner) GetReceiverInfoOk() (*GetPayTradeHistoryResponseDataInnerReceiverInfo, bool) {
	if o == nil || common.IsNil(o.ReceiverInfo) {
		return nil, false
	}
	return o.ReceiverInfo, true
}

// HasReceiverInfo returns a boolean if a field has been set.
func (o *GetPayTradeHistoryResponseDataInner) HasReceiverInfo() bool {
	if o != nil && !common.IsNil(o.ReceiverInfo) {
		return true
	}

	return false
}

// SetReceiverInfo gets a reference to the given GetPayTradeHistoryResponseDataInnerReceiverInfo and assigns it to the ReceiverInfo field.
func (o *GetPayTradeHistoryResponseDataInner) SetReceiverInfo(v GetPayTradeHistoryResponseDataInnerReceiverInfo) {
	o.ReceiverInfo = &v
}

func (o GetPayTradeHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPayTradeHistoryResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !common.IsNil(o.TransactionTime) {
		toSerialize["transactionTime"] = o.TransactionTime
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !common.IsNil(o.WalletType) {
		toSerialize["walletType"] = o.WalletType
	}
	if !common.IsNil(o.WalletTypes) {
		toSerialize["walletTypes"] = o.WalletTypes
	}
	if !common.IsNil(o.FundsDetail) {
		toSerialize["fundsDetail"] = o.FundsDetail
	}
	if !common.IsNil(o.PayerInfo) {
		toSerialize["payerInfo"] = o.PayerInfo
	}
	if !common.IsNil(o.ReceiverInfo) {
		toSerialize["receiverInfo"] = o.ReceiverInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetPayTradeHistoryResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetPayTradeHistoryResponseDataInner := _GetPayTradeHistoryResponseDataInner{}

	err = json.Unmarshal(data, &varGetPayTradeHistoryResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetPayTradeHistoryResponseDataInner(varGetPayTradeHistoryResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "transactionId")
		delete(additionalProperties, "transactionTime")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "currency")
		delete(additionalProperties, "walletType")
		delete(additionalProperties, "walletTypes")
		delete(additionalProperties, "fundsDetail")
		delete(additionalProperties, "payerInfo")
		delete(additionalProperties, "receiverInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetPayTradeHistoryResponseDataInner struct {
	value *GetPayTradeHistoryResponseDataInner
	isSet bool
}

func (v NullableGetPayTradeHistoryResponseDataInner) Get() *GetPayTradeHistoryResponseDataInner {
	return v.value
}

func (v *NullableGetPayTradeHistoryResponseDataInner) Set(val *GetPayTradeHistoryResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPayTradeHistoryResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPayTradeHistoryResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPayTradeHistoryResponseDataInner(val *GetPayTradeHistoryResponseDataInner) *NullableGetPayTradeHistoryResponseDataInner {
	return &NullableGetPayTradeHistoryResponseDataInner{value: val, isSet: true}
}

func (v NullableGetPayTradeHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPayTradeHistoryResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
