/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetNFTTransactionHistoryResponseListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTTransactionHistoryResponseListInner{}

// GetNFTTransactionHistoryResponseListInner struct for GetNFTTransactionHistoryResponseListInner
type GetNFTTransactionHistoryResponseListInner struct {
	OrderNo              *string                                                `json:"orderNo,omitempty"`
	Tokens               []GetNFTTransactionHistoryResponseListInnerTokensInner `json:"tokens,omitempty"`
	TradeTime            *int64                                                 `json:"tradeTime,omitempty"`
	TradeAmount          *string                                                `json:"tradeAmount,omitempty"`
	TradeCurrency        *string                                                `json:"tradeCurrency,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTTransactionHistoryResponseListInner GetNFTTransactionHistoryResponseListInner

// NewGetNFTTransactionHistoryResponseListInner instantiates a new GetNFTTransactionHistoryResponseListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTTransactionHistoryResponseListInner() *GetNFTTransactionHistoryResponseListInner {
	this := GetNFTTransactionHistoryResponseListInner{}
	return &this
}

// NewGetNFTTransactionHistoryResponseListInnerWithDefaults instantiates a new GetNFTTransactionHistoryResponseListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTTransactionHistoryResponseListInnerWithDefaults() *GetNFTTransactionHistoryResponseListInner {
	this := GetNFTTransactionHistoryResponseListInner{}
	return &this
}

// GetOrderNo returns the OrderNo field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInner) GetOrderNo() string {
	if o == nil || common.IsNil(o.OrderNo) {
		var ret string
		return ret
	}
	return *o.OrderNo
}

// GetOrderNoOk returns a tuple with the OrderNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInner) GetOrderNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderNo) {
		return nil, false
	}
	return o.OrderNo, true
}

// HasOrderNo returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInner) HasOrderNo() bool {
	if o != nil && !common.IsNil(o.OrderNo) {
		return true
	}

	return false
}

// SetOrderNo gets a reference to the given string and assigns it to the OrderNo field.
func (o *GetNFTTransactionHistoryResponseListInner) SetOrderNo(v string) {
	o.OrderNo = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInner) GetTokens() []GetNFTTransactionHistoryResponseListInnerTokensInner {
	if o == nil || common.IsNil(o.Tokens) {
		var ret []GetNFTTransactionHistoryResponseListInnerTokensInner
		return ret
	}
	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInner) GetTokensOk() ([]GetNFTTransactionHistoryResponseListInnerTokensInner, bool) {
	if o == nil || common.IsNil(o.Tokens) {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInner) HasTokens() bool {
	if o != nil && !common.IsNil(o.Tokens) {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []GetNFTTransactionHistoryResponseListInnerTokensInner and assigns it to the Tokens field.
func (o *GetNFTTransactionHistoryResponseListInner) SetTokens(v []GetNFTTransactionHistoryResponseListInnerTokensInner) {
	o.Tokens = v
}

// GetTradeTime returns the TradeTime field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeTime() int64 {
	if o == nil || common.IsNil(o.TradeTime) {
		var ret int64
		return ret
	}
	return *o.TradeTime
}

// GetTradeTimeOk returns a tuple with the TradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeTime) {
		return nil, false
	}
	return o.TradeTime, true
}

// HasTradeTime returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInner) HasTradeTime() bool {
	if o != nil && !common.IsNil(o.TradeTime) {
		return true
	}

	return false
}

// SetTradeTime gets a reference to the given int64 and assigns it to the TradeTime field.
func (o *GetNFTTransactionHistoryResponseListInner) SetTradeTime(v int64) {
	o.TradeTime = &v
}

// GetTradeAmount returns the TradeAmount field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeAmount() string {
	if o == nil || common.IsNil(o.TradeAmount) {
		var ret string
		return ret
	}
	return *o.TradeAmount
}

// GetTradeAmountOk returns a tuple with the TradeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeAmount) {
		return nil, false
	}
	return o.TradeAmount, true
}

// HasTradeAmount returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInner) HasTradeAmount() bool {
	if o != nil && !common.IsNil(o.TradeAmount) {
		return true
	}

	return false
}

// SetTradeAmount gets a reference to the given string and assigns it to the TradeAmount field.
func (o *GetNFTTransactionHistoryResponseListInner) SetTradeAmount(v string) {
	o.TradeAmount = &v
}

// GetTradeCurrency returns the TradeCurrency field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeCurrency() string {
	if o == nil || common.IsNil(o.TradeCurrency) {
		var ret string
		return ret
	}
	return *o.TradeCurrency
}

// GetTradeCurrencyOk returns a tuple with the TradeCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponseListInner) GetTradeCurrencyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeCurrency) {
		return nil, false
	}
	return o.TradeCurrency, true
}

// HasTradeCurrency returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponseListInner) HasTradeCurrency() bool {
	if o != nil && !common.IsNil(o.TradeCurrency) {
		return true
	}

	return false
}

// SetTradeCurrency gets a reference to the given string and assigns it to the TradeCurrency field.
func (o *GetNFTTransactionHistoryResponseListInner) SetTradeCurrency(v string) {
	o.TradeCurrency = &v
}

func (o GetNFTTransactionHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTTransactionHistoryResponseListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderNo) {
		toSerialize["orderNo"] = o.OrderNo
	}
	if !common.IsNil(o.Tokens) {
		toSerialize["tokens"] = o.Tokens
	}
	if !common.IsNil(o.TradeTime) {
		toSerialize["tradeTime"] = o.TradeTime
	}
	if !common.IsNil(o.TradeAmount) {
		toSerialize["tradeAmount"] = o.TradeAmount
	}
	if !common.IsNil(o.TradeCurrency) {
		toSerialize["tradeCurrency"] = o.TradeCurrency
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetNFTTransactionHistoryResponseListInner) UnmarshalJSON(data []byte) (err error) {
	varGetNFTTransactionHistoryResponseListInner := _GetNFTTransactionHistoryResponseListInner{}

	err = json.Unmarshal(data, &varGetNFTTransactionHistoryResponseListInner)

	if err != nil {
		return err
	}

	*o = GetNFTTransactionHistoryResponseListInner(varGetNFTTransactionHistoryResponseListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderNo")
		delete(additionalProperties, "tokens")
		delete(additionalProperties, "tradeTime")
		delete(additionalProperties, "tradeAmount")
		delete(additionalProperties, "tradeCurrency")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTTransactionHistoryResponseListInner struct {
	value *GetNFTTransactionHistoryResponseListInner
	isSet bool
}

func (v NullableGetNFTTransactionHistoryResponseListInner) Get() *GetNFTTransactionHistoryResponseListInner {
	return v.value
}

func (v *NullableGetNFTTransactionHistoryResponseListInner) Set(val *GetNFTTransactionHistoryResponseListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTTransactionHistoryResponseListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTTransactionHistoryResponseListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTTransactionHistoryResponseListInner(val *GetNFTTransactionHistoryResponseListInner) *NullableGetNFTTransactionHistoryResponseListInner {
	return &NullableGetNFTTransactionHistoryResponseListInner{value: val, isSet: true}
}

func (v NullableGetNFTTransactionHistoryResponseListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTTransactionHistoryResponseListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
