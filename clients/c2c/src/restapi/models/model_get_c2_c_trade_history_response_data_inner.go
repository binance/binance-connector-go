/*
Binance C2C REST API

OpenAPI Specification for the Binance C2C REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetC2CTradeHistoryResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetC2CTradeHistoryResponseDataInner{}

// GetC2CTradeHistoryResponseDataInner struct for GetC2CTradeHistoryResponseDataInner
type GetC2CTradeHistoryResponseDataInner struct {
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
	CounterPartNickName  *string `json:"counterPartNickName,omitempty"`
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
