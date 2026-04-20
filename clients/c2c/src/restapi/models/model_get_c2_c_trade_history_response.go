/*
Binance C2C REST API

OpenAPI Specification for the Binance C2C REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetC2CTradeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetC2CTradeHistoryResponse{}

// GetC2CTradeHistoryResponse struct for GetC2CTradeHistoryResponse
type GetC2CTradeHistoryResponse struct {
	OrderNumber          *string `json:"orderNumber,omitempty"`
	AdvNo                *string `json:"advNo,omitempty"`
	TradeType            *string `json:"tradeType,omitempty"`
	Asset                *string `json:"asset,omitempty"`
	Fiat                 *string `json:"fiat,omitempty"`
	FiatSymbol           *string `json:"fiatSymbol,omitempty"`
	Amount               *string `json:"amount,omitempty"`
	TotalPrice           *string `json:"totalPrice,omitempty"`
	UnitPrice            *string `json:"unitPrice,omitempty"`
	OrderStatus          *string `json:"orderStatus,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	TakerCommissionRate  *string `json:"takerCommissionRate,omitempty"`
	TakerCommission      *string `json:"takerCommission,omitempty"`
	TakerAmount          *string `json:"takerAmount,omitempty"`
	CounterPartNickName  *string `json:"counterPartNickName,omitempty"`
	PayMethodName        *string `json:"payMethodName,omitempty"`
	AdditionalKycVerify  *int64  `json:"additionalKycVerify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetC2CTradeHistoryResponse GetC2CTradeHistoryResponse

// NewGetC2CTradeHistoryResponse instantiates a new GetC2CTradeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetC2CTradeHistoryResponse() *GetC2CTradeHistoryResponse {
	this := GetC2CTradeHistoryResponse{}
	return &this
}

// NewGetC2CTradeHistoryResponseWithDefaults instantiates a new GetC2CTradeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetC2CTradeHistoryResponseWithDefaults() *GetC2CTradeHistoryResponse {
	this := GetC2CTradeHistoryResponse{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetOrderNumber() string {
	if o == nil || common.IsNil(o.OrderNumber) {
		var ret string
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetOrderNumberOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasOrderNumber() bool {
	if o != nil && !common.IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given string and assigns it to the OrderNumber field.
func (o *GetC2CTradeHistoryResponse) SetOrderNumber(v string) {
	o.OrderNumber = &v
}

// GetAdvNo returns the AdvNo field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetAdvNo() string {
	if o == nil || common.IsNil(o.AdvNo) {
		var ret string
		return ret
	}
	return *o.AdvNo
}

// GetAdvNoOk returns a tuple with the AdvNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetAdvNoOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdvNo) {
		return nil, false
	}
	return o.AdvNo, true
}

// HasAdvNo returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasAdvNo() bool {
	if o != nil && !common.IsNil(o.AdvNo) {
		return true
	}

	return false
}

// SetAdvNo gets a reference to the given string and assigns it to the AdvNo field.
func (o *GetC2CTradeHistoryResponse) SetAdvNo(v string) {
	o.AdvNo = &v
}

// GetTradeType returns the TradeType field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTradeType() string {
	if o == nil || common.IsNil(o.TradeType) {
		var ret string
		return ret
	}
	return *o.TradeType
}

// GetTradeTypeOk returns a tuple with the TradeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTradeTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TradeType) {
		return nil, false
	}
	return o.TradeType, true
}

// HasTradeType returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTradeType() bool {
	if o != nil && !common.IsNil(o.TradeType) {
		return true
	}

	return false
}

// SetTradeType gets a reference to the given string and assigns it to the TradeType field.
func (o *GetC2CTradeHistoryResponse) SetTradeType(v string) {
	o.TradeType = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetC2CTradeHistoryResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetFiat returns the Fiat field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetFiat() string {
	if o == nil || common.IsNil(o.Fiat) {
		var ret string
		return ret
	}
	return *o.Fiat
}

// GetFiatOk returns a tuple with the Fiat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetFiatOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fiat) {
		return nil, false
	}
	return o.Fiat, true
}

// HasFiat returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasFiat() bool {
	if o != nil && !common.IsNil(o.Fiat) {
		return true
	}

	return false
}

// SetFiat gets a reference to the given string and assigns it to the Fiat field.
func (o *GetC2CTradeHistoryResponse) SetFiat(v string) {
	o.Fiat = &v
}

// GetFiatSymbol returns the FiatSymbol field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetFiatSymbol() string {
	if o == nil || common.IsNil(o.FiatSymbol) {
		var ret string
		return ret
	}
	return *o.FiatSymbol
}

// GetFiatSymbolOk returns a tuple with the FiatSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetFiatSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.FiatSymbol) {
		return nil, false
	}
	return o.FiatSymbol, true
}

// HasFiatSymbol returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasFiatSymbol() bool {
	if o != nil && !common.IsNil(o.FiatSymbol) {
		return true
	}

	return false
}

// SetFiatSymbol gets a reference to the given string and assigns it to the FiatSymbol field.
func (o *GetC2CTradeHistoryResponse) SetFiatSymbol(v string) {
	o.FiatSymbol = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetAmount() string {
	if o == nil || common.IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasAmount() bool {
	if o != nil && !common.IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *GetC2CTradeHistoryResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetTotalPrice returns the TotalPrice field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTotalPrice() string {
	if o == nil || common.IsNil(o.TotalPrice) {
		var ret string
		return ret
	}
	return *o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTotalPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPrice) {
		return nil, false
	}
	return o.TotalPrice, true
}

// HasTotalPrice returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTotalPrice() bool {
	if o != nil && !common.IsNil(o.TotalPrice) {
		return true
	}

	return false
}

// SetTotalPrice gets a reference to the given string and assigns it to the TotalPrice field.
func (o *GetC2CTradeHistoryResponse) SetTotalPrice(v string) {
	o.TotalPrice = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetUnitPrice() string {
	if o == nil || common.IsNil(o.UnitPrice) {
		var ret string
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetUnitPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasUnitPrice() bool {
	if o != nil && !common.IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given string and assigns it to the UnitPrice field.
func (o *GetC2CTradeHistoryResponse) SetUnitPrice(v string) {
	o.UnitPrice = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetOrderStatus() string {
	if o == nil || common.IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasOrderStatus() bool {
	if o != nil && !common.IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *GetC2CTradeHistoryResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *GetC2CTradeHistoryResponse) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *GetC2CTradeHistoryResponse) SetCommission(v string) {
	o.Commission = &v
}

// GetTakerCommissionRate returns the TakerCommissionRate field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTakerCommissionRate() string {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		var ret string
		return ret
	}
	return *o.TakerCommissionRate
}

// GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTakerCommissionRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommissionRate) {
		return nil, false
	}
	return o.TakerCommissionRate, true
}

// HasTakerCommissionRate returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTakerCommissionRate() bool {
	if o != nil && !common.IsNil(o.TakerCommissionRate) {
		return true
	}

	return false
}

// SetTakerCommissionRate gets a reference to the given string and assigns it to the TakerCommissionRate field.
func (o *GetC2CTradeHistoryResponse) SetTakerCommissionRate(v string) {
	o.TakerCommissionRate = &v
}

// GetTakerCommission returns the TakerCommission field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTakerCommission() string {
	if o == nil || common.IsNil(o.TakerCommission) {
		var ret string
		return ret
	}
	return *o.TakerCommission
}

// GetTakerCommissionOk returns a tuple with the TakerCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTakerCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerCommission) {
		return nil, false
	}
	return o.TakerCommission, true
}

// HasTakerCommission returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTakerCommission() bool {
	if o != nil && !common.IsNil(o.TakerCommission) {
		return true
	}

	return false
}

// SetTakerCommission gets a reference to the given string and assigns it to the TakerCommission field.
func (o *GetC2CTradeHistoryResponse) SetTakerCommission(v string) {
	o.TakerCommission = &v
}

// GetTakerAmount returns the TakerAmount field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTakerAmount() string {
	if o == nil || common.IsNil(o.TakerAmount) {
		var ret string
		return ret
	}
	return *o.TakerAmount
}

// GetTakerAmountOk returns a tuple with the TakerAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTakerAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TakerAmount) {
		return nil, false
	}
	return o.TakerAmount, true
}

// HasTakerAmount returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTakerAmount() bool {
	if o != nil && !common.IsNil(o.TakerAmount) {
		return true
	}

	return false
}

// SetTakerAmount gets a reference to the given string and assigns it to the TakerAmount field.
func (o *GetC2CTradeHistoryResponse) SetTakerAmount(v string) {
	o.TakerAmount = &v
}

// GetCounterPartNickName returns the CounterPartNickName field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetCounterPartNickName() string {
	if o == nil || common.IsNil(o.CounterPartNickName) {
		var ret string
		return ret
	}
	return *o.CounterPartNickName
}

// GetCounterPartNickNameOk returns a tuple with the CounterPartNickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetCounterPartNickNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CounterPartNickName) {
		return nil, false
	}
	return o.CounterPartNickName, true
}

// HasCounterPartNickName returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasCounterPartNickName() bool {
	if o != nil && !common.IsNil(o.CounterPartNickName) {
		return true
	}

	return false
}

// SetCounterPartNickName gets a reference to the given string and assigns it to the CounterPartNickName field.
func (o *GetC2CTradeHistoryResponse) SetCounterPartNickName(v string) {
	o.CounterPartNickName = &v
}

// GetPayMethodName returns the PayMethodName field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetPayMethodName() string {
	if o == nil || common.IsNil(o.PayMethodName) {
		var ret string
		return ret
	}
	return *o.PayMethodName
}

// GetPayMethodNameOk returns a tuple with the PayMethodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetPayMethodNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.PayMethodName) {
		return nil, false
	}
	return o.PayMethodName, true
}

// HasPayMethodName returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasPayMethodName() bool {
	if o != nil && !common.IsNil(o.PayMethodName) {
		return true
	}

	return false
}

// SetPayMethodName gets a reference to the given string and assigns it to the PayMethodName field.
func (o *GetC2CTradeHistoryResponse) SetPayMethodName(v string) {
	o.PayMethodName = &v
}

// GetAdditionalKycVerify returns the AdditionalKycVerify field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetAdditionalKycVerify() int64 {
	if o == nil || common.IsNil(o.AdditionalKycVerify) {
		var ret int64
		return ret
	}
	return *o.AdditionalKycVerify
}

// GetAdditionalKycVerifyOk returns a tuple with the AdditionalKycVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetAdditionalKycVerifyOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AdditionalKycVerify) {
		return nil, false
	}
	return o.AdditionalKycVerify, true
}

// HasAdditionalKycVerify returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasAdditionalKycVerify() bool {
	if o != nil && !common.IsNil(o.AdditionalKycVerify) {
		return true
	}

	return false
}

// SetAdditionalKycVerify gets a reference to the given int64 and assigns it to the AdditionalKycVerify field.
func (o *GetC2CTradeHistoryResponse) SetAdditionalKycVerify(v int64) {
	o.AdditionalKycVerify = &v
}

func (o GetC2CTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetC2CTradeHistoryResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.TakerCommissionRate) {
		toSerialize["takerCommissionRate"] = o.TakerCommissionRate
	}
	if !common.IsNil(o.TakerCommission) {
		toSerialize["takerCommission"] = o.TakerCommission
	}
	if !common.IsNil(o.TakerAmount) {
		toSerialize["takerAmount"] = o.TakerAmount
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetC2CTradeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetC2CTradeHistoryResponse := _GetC2CTradeHistoryResponse{}

	err = json.Unmarshal(data, &varGetC2CTradeHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetC2CTradeHistoryResponse(varGetC2CTradeHistoryResponse)

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
		delete(additionalProperties, "takerCommissionRate")
		delete(additionalProperties, "takerCommission")
		delete(additionalProperties, "takerAmount")
		delete(additionalProperties, "counterPartNickName")
		delete(additionalProperties, "payMethodName")
		delete(additionalProperties, "additionalKycVerify")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetC2CTradeHistoryResponse struct {
	value *GetC2CTradeHistoryResponse
	isSet bool
}

func (v NullableGetC2CTradeHistoryResponse) Get() *GetC2CTradeHistoryResponse {
	return v.value
}

func (v *NullableGetC2CTradeHistoryResponse) Set(val *GetC2CTradeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetC2CTradeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetC2CTradeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetC2CTradeHistoryResponse(val *GetC2CTradeHistoryResponse) *NullableGetC2CTradeHistoryResponse {
	return &NullableGetC2CTradeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetC2CTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetC2CTradeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
