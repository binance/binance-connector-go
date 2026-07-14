/*
C2C REST API

Query fiat transaction history via the C2C REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetC2CTradeHistoryResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetC2CTradeHistoryResponseDataInner{}

// GetC2CTradeHistoryResponseDataInner struct for GetC2CTradeHistoryResponseDataInner
type GetC2CTradeHistoryResponseDataInner struct {
	OrderNumber *string `json:"orderNumber,omitempty"`
	AdvNo       *string `json:"advNo,omitempty"`
	TradeType   *string `json:"tradeType,omitempty"`
	Asset       *string `json:"asset,omitempty"`
	Fiat        *string `json:"fiat,omitempty"`
	FiatSymbol  *string `json:"fiatSymbol,omitempty"`
	// Quantity (in Crypto)
	Amount *string `json:"amount,omitempty"`
	// Total order amount in fiat
	TotalPrice *string `json:"totalPrice,omitempty"`
	// Unit Price (in Fiat)
	UnitPrice   *string `json:"unitPrice,omitempty"`
	OrderStatus *string `json:"orderStatus,omitempty"`
	// Order creation timestamp in milliseconds
	CreateTime *int64 `json:"createTime,omitempty"`
	// Transaction Fee (in Crypto)
	Commission *string `json:"commission,omitempty"`
	// Counterparty nickname
	CounterPartNickName *string `json:"counterPartNickName,omitempty"`
	// Identifier of the payment method
	PayMethodName *string `json:"payMethodName,omitempty"`
	// KYC verification status. 0: not required, 1: not verified, 2: verified
	AdditionalKycVerify *int64 `json:"additionalKycVerify,omitempty"`
	// Taker commission rate
	TakerCommissionRate *string `json:"takerCommissionRate,omitempty"`
	// Taker commission amount
	TakerCommission *string `json:"takerCommission,omitempty"`
	// Taker trade amount
	TakerAmount          *string `json:"takerAmount,omitempty"`
	AdvertisementRole    *string `json:"advertisementRole,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetC2CTradeHistoryResponseDataInner GetC2CTradeHistoryResponseDataInner

// NewGetC2CTradeHistoryResponseDataInner instantiates a new GetC2CTradeHistoryResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetC2CTradeHistoryResponseDataInner() *GetC2CTradeHistoryResponseDataInner {
	this := GetC2CTradeHistoryResponseDataInner{}
	return &this
}

// NewGetC2CTradeHistoryResponseDataInnerWithDefaults instantiates a new GetC2CTradeHistoryResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetC2CTradeHistoryResponseDataInnerWithDefaults() *GetC2CTradeHistoryResponseDataInner {
	this := GetC2CTradeHistoryResponseDataInner{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetOrderNumber() string {
	if o == nil || common.IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetOrderNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasOrderNumber() bool {
	if o != nil && !common.IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *GetC2CTradeHistoryResponseDataInner) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetAdvNo returns the AdvNo field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdvNo() string {
	if o == nil || common.IsNil(o.AdvNo) {
		var ret string
		return ret
	}
	return *o.AdvNo
}

// GetAdvNoOk returns a tuple with the AdvNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdvNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdvNo) {
		return nil, false
	}
	return o.AdvNo, true
}

// HasAdvNo returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasAdvNo() bool {
	if o != nil && !common.IsNil(o.AdvNo) {
		return true
	}

	return false
}

// SetAdvNo gets a reference to the given string and assigns it to the AdvNo field.
func (o *GetC2CTradeHistoryResponseDataInner) SetAdvNo(v string) {
	o.AdvNo = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetTradeType() string {
	if o == nil || common.IsNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetTradeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeType) {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasTradeType() bool {
	if o != nil && !common.IsNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *GetC2CTradeHistoryResponseDataInner) SetTradeType(v string) {
	o.TradeType = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetC2CTradeHistoryResponseDataInner) SetAsset(v string) {
	o.Asset = &v
}

// GetFiat returns the Fiat field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetFiat() string {
	if o == nil || common.IsNil(o.Fiat) {
		var ret string
		return ret
	}
	return *o.Fiat
}

// GetFiatOk returns a tuple with the Fiat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetFiatOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fiat) {
		return nil, false
	}
	return o.Fiat, true
}

// HasFiat returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasFiat() bool {
	if o != nil && !common.IsNil(o.Fiat) {
		return true
	}

	return false
}

// SetFiat gets a reference to the given string and assigns it to the Fiat field.
func (o *GetC2CTradeHistoryResponseDataInner) SetFiat(v string) {
	o.Fiat = &v
}

// GetFiatSymbol returns the FiatSymbol field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetFiatSymbol() string {
	if o == nil || common.IsNil(o.FiatSymbol) {
		var ret string
		return ret
	}
	return *o.FiatSymbol
}

// GetFiatSymbolOk returns a tuple with the FiatSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetFiatSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.FiatSymbol) {
		return nil, false
	}
	return o.FiatSymbol, true
}

// HasFiatSymbol returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasFiatSymbol() bool {
	if o != nil && !common.IsNil(o.FiatSymbol) {
		return true
	}

	return false
}

// SetFiatSymbol gets a reference to the given string and assigns it to the FiatSymbol field.
func (o *GetC2CTradeHistoryResponseDataInner) SetFiatSymbol(v string) {
	o.FiatSymbol = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetC2CTradeHistoryResponseDataInner) SetAmount(v string) {
	o.Amount = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetTotalPrice() string {
	if o == nil || common.IsNil(o.TotalPrice) {
		var ret string
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetTotalPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPrice) {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasTotalPrice() bool {
	if o != nil && !common.IsNil(o.TotalPrice) {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given string and assigns it to the TotalPrice field.
func (o *GetC2CTradeHistoryResponseDataInner) SetTotalPrice(v string) {
	o.TotalPrice = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetUnitPrice() string {
	if o == nil || common.IsNil(o.UnitPrice) {
		var ret string
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetUnitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasUnitPrice() bool {
	if o != nil && !common.IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given string and assigns it to the UnitPrice field.
func (o *GetC2CTradeHistoryResponseDataInner) SetUnitPrice(v string) {
	o.UnitPrice = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *GetC2CTradeHistoryResponseDataInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetC2CTradeHistoryResponseDataInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *GetC2CTradeHistoryResponseDataInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCounterPartNickName returns the CounterPartNickName field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetCounterPartNickName() string {
	if o == nil || common.IsNil(o.CounterPartNickName) {
		var ret string
		return ret
	}
	return *o.CounterPartNickName
}

// GetCounterPartNickNameOk returns a tuple with the CounterPartNickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetCounterPartNickNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CounterPartNickName) {
		return nil, false
	}
	return o.CounterPartNickName, true
}

// HasCounterPartNickName returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasCounterPartNickName() bool {
	if o != nil && !common.IsNil(o.CounterPartNickName) {
		return true
	}

	return false
}

// SetCounterPartNickName gets a reference to the given string and assigns it to the CounterPartNickName field.
func (o *GetC2CTradeHistoryResponseDataInner) SetCounterPartNickName(v string) {
	o.CounterPartNickName = &v
}

// GetPayMethodName returns the PayMethodName field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetPayMethodName() string {
	if o == nil || common.IsNil(o.PayMethodName) {
		var ret string
		return ret
	}
	return *o.PayMethodName
}

// GetPayMethodNameOk returns a tuple with the PayMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetPayMethodNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayMethodName) {
		return nil, false
	}
	return o.PayMethodName, true
}

// HasPayMethodName returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasPayMethodName() bool {
	if o != nil && !common.IsNil(o.PayMethodName) {
		return true
	}

	return false
}

// SetPayMethodName gets a reference to the given string and assigns it to the PayMethodName field.
func (o *GetC2CTradeHistoryResponseDataInner) SetPayMethodName(v string) {
	o.PayMethodName = &v
}

// GetAdditionalKycVerify returns the AdditionalKycVerify field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdditionalKycVerify() int64 {
	if o == nil || common.IsNil(o.AdditionalKycVerify) {
		var ret int64
		return ret
	}
	return *o.AdditionalKycVerify
}

// GetAdditionalKycVerifyOk returns a tuple with the AdditionalKycVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdditionalKycVerifyOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AdditionalKycVerify) {
		return nil, false
	}
	return o.AdditionalKycVerify, true
}

// HasAdditionalKycVerify returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasAdditionalKycVerify() bool {
	if o != nil && !common.IsNil(o.AdditionalKycVerify) {
		return true
	}

	return false
}

// SetAdditionalKycVerify gets a reference to the given int64 and assigns it to the AdditionalKycVerify field.
func (o *GetC2CTradeHistoryResponseDataInner) SetAdditionalKycVerify(v int64) {
	o.AdditionalKycVerify = &v
}

// GetTakerCommissionRate returns the TakerCommissionRate field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerCommissionRate() string {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.TakerCommissionRate
}

// GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		return nil, false
	}
	return o.TakerCommissionRate, true
}

// HasTakerCommissionRate returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasTakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.TakerCommissionRate) {
		return true
	}

	return false
}

// SetTakerCommissionRate gets a reference to the given string and assigns it to the TakerCommissionRate field.
func (o *GetC2CTradeHistoryResponseDataInner) SetTakerCommissionRate(v string) {
	o.TakerCommissionRate = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerCommission() string {
	if o == nil || common.IsNil(o.TakerCommission) {
		var ret string
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasTakerCommission() bool {
	if o != nil && !common.IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given string and assigns it to the TakerCommission field.
func (o *GetC2CTradeHistoryResponseDataInner) SetTakerCommission(v string) {
	o.TakerCommission = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerAmount() string {
	if o == nil || common.IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetTakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasTakerAmount() bool {
	if o != nil && !common.IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *GetC2CTradeHistoryResponseDataInner) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetAdvertisementRole returns the AdvertisementRole field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdvertisementRole() string {
	if o == nil || common.IsNil(o.AdvertisementRole) {
		var ret string
		return ret
	}
	return *o.AdvertisementRole
}

// GetAdvertisementRoleOk returns a tuple with the AdvertisementRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponseDataInner) GetAdvertisementRoleOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdvertisementRole) {
		return nil, false
	}
	return o.AdvertisementRole, true
}

// HasAdvertisementRole returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponseDataInner) HasAdvertisementRole() bool {
	if o != nil && !common.IsNil(o.AdvertisementRole) {
		return true
	}

	return false
}

// SetAdvertisementRole gets a reference to the given string and assigns it to the AdvertisementRole field.
func (o *GetC2CTradeHistoryResponseDataInner) SetAdvertisementRole(v string) {
	o.AdvertisementRole = &v
}

func (o GetC2CTradeHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetC2CTradeHistoryResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !common.IsNil(o.AdvNo) {
		toSerialize["advNo"] = o.AdvNo
	}
	if !common.IsNil(o.TradeType) {
		toSerialize["tradeType"] = o.TradeType
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.Fiat) {
		toSerialize["fiat"] = o.Fiat
	}
	if !common.IsNil(o.FiatSymbol) {
		toSerialize["fiatSymbol"] = o.FiatSymbol
	}
	if !common.IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !common.IsNil(o.TotalPrice) {
		toSerialize["totalPrice"] = o.TotalPrice
	}
	if !common.IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !common.IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CounterPartNickName) {
		toSerialize["counterPartNickName"] = o.CounterPartNickName
	}
	if !common.IsNil(o.PayMethodName) {
		toSerialize["payMethodName"] = o.PayMethodName
	}
	if !common.IsNil(o.AdditionalKycVerify) {
		toSerialize["additionalKycVerify"] = o.AdditionalKycVerify
	}
	if !common.IsNil(o.TakerCommissionRate) {
		toSerialize["takerCommissionRate"] = o.TakerCommissionRate
	}
	if !common.IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	if !common.IsNil(o.TakerAmount) {
		toSerialize["takerAmount"] = o.TakerAmount
	}
	if !common.IsNil(o.AdvertisementRole) {
		toSerialize["advertisementRole"] = o.AdvertisementRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetC2CTradeHistoryResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varGetC2CTradeHistoryResponseDataInner := _GetC2CTradeHistoryResponseDataInner{}

	err = json.Unmarshal(data, &varGetC2CTradeHistoryResponseDataInner)

	if err != nil {
		return err
	}

	*o = GetC2CTradeHistoryResponseDataInner(varGetC2CTradeHistoryResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderNumber")
		delete(additionalProperties, "advNo")
		delete(additionalProperties, "tradeType")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "fiat")
		delete(additionalProperties, "fiatSymbol")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "totalPrice")
		delete(additionalProperties, "unitPrice")
		delete(additionalProperties, "orderStatus")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "counterPartNickName")
		delete(additionalProperties, "payMethodName")
		delete(additionalProperties, "additionalKycVerify")
		delete(additionalProperties, "takerCommissionRate")
		delete(additionalProperties, "takerCommission")
		delete(additionalProperties, "takerAmount")
		delete(additionalProperties, "advertisementRole")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetC2CTradeHistoryResponseDataInner struct {
	value *GetC2CTradeHistoryResponseDataInner
	isSet bool
}

func (v NullableGetC2CTradeHistoryResponseDataInner) Get() *GetC2CTradeHistoryResponseDataInner {
	return v.value
}

func (v *NullableGetC2CTradeHistoryResponseDataInner) Set(val *GetC2CTradeHistoryResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetC2CTradeHistoryResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetC2CTradeHistoryResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetC2CTradeHistoryResponseDataInner(val *GetC2CTradeHistoryResponseDataInner) *NullableGetC2CTradeHistoryResponseDataInner {
	return &NullableGetC2CTradeHistoryResponseDataInner{value: val, isSet: true}
}

func (v NullableGetC2CTradeHistoryResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetC2CTradeHistoryResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
