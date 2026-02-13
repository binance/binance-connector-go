/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CmAccountTradeListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CmAccountTradeListResponseInner{}

// CmAccountTradeListResponseInner struct for CmAccountTradeListResponseInner
type CmAccountTradeListResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Id                   *int64  `json:"id,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	Pair                 *string `json:"pair,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	RealizedPnl          *string `json:"realizedPnl,omitempty"`
	MarginAsset          *string `json:"marginAsset,omitempty"`
	BaseQty              *string `json:"baseQty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	Buyer                *bool   `json:"buyer,omitempty"`
	Maker                *bool   `json:"maker,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CmAccountTradeListResponseInner CmAccountTradeListResponseInner

// NewCmAccountTradeListResponseInner instantiates a new CmAccountTradeListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmAccountTradeListResponseInner() *CmAccountTradeListResponseInner {
	this := CmAccountTradeListResponseInner{}
	return &this
}

// NewCmAccountTradeListResponseInnerWithDefaults instantiates a new CmAccountTradeListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmAccountTradeListResponseInnerWithDefaults() *CmAccountTradeListResponseInner {
	this := CmAccountTradeListResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CmAccountTradeListResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CmAccountTradeListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CmAccountTradeListResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *CmAccountTradeListResponseInner) SetPair(v string) {
	o.Pair = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CmAccountTradeListResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CmAccountTradeListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *CmAccountTradeListResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetRealizedPnl returns the RealizedPnl field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetRealizedPnl() string {
	if o == nil || common.IsNil(o.RealizedPnl) {
		var ret string
		return ret
	}
	return *o.RealizedPnl
}

// GetRealizedPnlOk returns a tuple with the RealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedPnl) {
		return nil, false
	}
	return o.RealizedPnl, true
}

// HasRealizedPnl returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasRealizedPnl() bool {
	if o != nil && !common.IsNil(o.RealizedPnl) {
		return true
	}

	return false
}

// SetRealizedPnl gets a reference to the given string and assigns it to the RealizedPnl field.
func (o *CmAccountTradeListResponseInner) SetRealizedPnl(v string) {
	o.RealizedPnl = &v
}

// GetMarginAsset returns the MarginAsset field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetMarginAsset() string {
	if o == nil || common.IsNil(o.MarginAsset) {
		var ret string
		return ret
	}
	return *o.MarginAsset
}

// GetMarginAssetOk returns a tuple with the MarginAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetMarginAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarginAsset) {
		return nil, false
	}
	return o.MarginAsset, true
}

// HasMarginAsset returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasMarginAsset() bool {
	if o != nil && !common.IsNil(o.MarginAsset) {
		return true
	}

	return false
}

// SetMarginAsset gets a reference to the given string and assigns it to the MarginAsset field.
func (o *CmAccountTradeListResponseInner) SetMarginAsset(v string) {
	o.MarginAsset = &v
}

// GetBaseQty returns the BaseQty field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetBaseQty() string {
	if o == nil || common.IsNil(o.BaseQty) {
		var ret string
		return ret
	}
	return *o.BaseQty
}

// GetBaseQtyOk returns a tuple with the BaseQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetBaseQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.BaseQty) {
		return nil, false
	}
	return o.BaseQty, true
}

// HasBaseQty returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasBaseQty() bool {
	if o != nil && !common.IsNil(o.BaseQty) {
		return true
	}

	return false
}

// SetBaseQty gets a reference to the given string and assigns it to the BaseQty field.
func (o *CmAccountTradeListResponseInner) SetBaseQty(v string) {
	o.BaseQty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *CmAccountTradeListResponseInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *CmAccountTradeListResponseInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *CmAccountTradeListResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *CmAccountTradeListResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetBuyer() bool {
	if o == nil || common.IsNil(o.Buyer) {
		var ret bool
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given bool and assigns it to the Buyer field.
func (o *CmAccountTradeListResponseInner) SetBuyer(v bool) {
	o.Buyer = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *CmAccountTradeListResponseInner) GetMaker() bool {
	if o == nil || common.IsNil(o.Maker) {
		var ret bool
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmAccountTradeListResponseInner) GetMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *CmAccountTradeListResponseInner) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given bool and assigns it to the Maker field.
func (o *CmAccountTradeListResponseInner) SetMaker(v bool) {
	o.Maker = &v
}

func (o CmAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CmAccountTradeListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Qty) {
		toSerialize["qty"] = o.Qty
	}
	if !common.IsNil(o.RealizedPnl) {
		toSerialize["realizedPnl"] = o.RealizedPnl
	}
	if !common.IsNil(o.MarginAsset) {
		toSerialize["marginAsset"] = o.MarginAsset
	}
	if !common.IsNil(o.BaseQty) {
		toSerialize["baseQty"] = o.BaseQty
	}
	if !common.IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !common.IsNil(o.CommissionAsset) {
		toSerialize["commissionAsset"] = o.CommissionAsset
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !common.IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CmAccountTradeListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCmAccountTradeListResponseInner := _CmAccountTradeListResponseInner{}

	err = json.Unmarshal(data, &varCmAccountTradeListResponseInner)

	if err != nil {
		return err
	}

	*o = CmAccountTradeListResponseInner(varCmAccountTradeListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "side")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "realizedPnl")
		delete(additionalProperties, "marginAsset")
		delete(additionalProperties, "baseQty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "buyer")
		delete(additionalProperties, "maker")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCmAccountTradeListResponseInner struct {
	value *CmAccountTradeListResponseInner
	isSet bool
}

func (v NullableCmAccountTradeListResponseInner) Get() *CmAccountTradeListResponseInner {
	return v.value
}

func (v *NullableCmAccountTradeListResponseInner) Set(val *CmAccountTradeListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCmAccountTradeListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCmAccountTradeListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmAccountTradeListResponseInner(val *CmAccountTradeListResponseInner) *NullableCmAccountTradeListResponseInner {
	return &NullableCmAccountTradeListResponseInner{value: val, isSet: true}
}

func (v NullableCmAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmAccountTradeListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
