/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UmAccountTradeListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UmAccountTradeListResponseInner{}

// UmAccountTradeListResponseInner struct for UmAccountTradeListResponseInner
type UmAccountTradeListResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	Id                   *int64  `json:"id,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Price                *string `json:"price,omitempty"`
	Qty                  *string `json:"qty,omitempty"`
	RealizedPnl          *string `json:"realizedPnl,omitempty"`
	QuoteQty             *string `json:"quoteQty,omitempty"`
	Commission           *string `json:"commission,omitempty"`
	CommissionAsset      *string `json:"commissionAsset,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	Buyer                *bool   `json:"buyer,omitempty"`
	Maker                *bool   `json:"maker,omitempty"`
	PositionSide         *string `json:"positionSide,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UmAccountTradeListResponseInner UmAccountTradeListResponseInner

// NewUmAccountTradeListResponseInner instantiates a new UmAccountTradeListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUmAccountTradeListResponseInner() *UmAccountTradeListResponseInner {
	this := UmAccountTradeListResponseInner{}
	return &this
}

// NewUmAccountTradeListResponseInnerWithDefaults instantiates a new UmAccountTradeListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUmAccountTradeListResponseInnerWithDefaults() *UmAccountTradeListResponseInner {
	this := UmAccountTradeListResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *UmAccountTradeListResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UmAccountTradeListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *UmAccountTradeListResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *UmAccountTradeListResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *UmAccountTradeListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQty returns the Qty field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetQty() string {
	if o == nil || common.IsNil(o.Qty) {
		var ret string
		return ret
	}
	return *o.Qty
}

// GetQtyOk returns a tuple with the Qty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.Qty) {
		return nil, false
	}
	return o.Qty, true
}

// HasQty returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasQty() bool {
	if o != nil && !common.IsNil(o.Qty) {
		return true
	}

	return false
}

// SetQty gets a reference to the given string and assigns it to the Qty field.
func (o *UmAccountTradeListResponseInner) SetQty(v string) {
	o.Qty = &v
}

// GetRealizedPnl returns the RealizedPnl field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetRealizedPnl() string {
	if o == nil || common.IsNil(o.RealizedPnl) {
		var ret string
		return ret
	}
	return *o.RealizedPnl
}

// GetRealizedPnlOk returns a tuple with the RealizedPnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetRealizedPnlOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedPnl) {
		return nil, false
	}
	return o.RealizedPnl, true
}

// HasRealizedPnl returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasRealizedPnl() bool {
	if o != nil && !common.IsNil(o.RealizedPnl) {
		return true
	}

	return false
}

// SetRealizedPnl gets a reference to the given string and assigns it to the RealizedPnl field.
func (o *UmAccountTradeListResponseInner) SetRealizedPnl(v string) {
	o.RealizedPnl = &v
}

// GetQuoteQty returns the QuoteQty field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetQuoteQty() string {
	if o == nil || common.IsNil(o.QuoteQty) {
		var ret string
		return ret
	}
	return *o.QuoteQty
}

// GetQuoteQtyOk returns a tuple with the QuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteQty) {
		return nil, false
	}
	return o.QuoteQty, true
}

// HasQuoteQty returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasQuoteQty() bool {
	if o != nil && !common.IsNil(o.QuoteQty) {
		return true
	}

	return false
}

// SetQuoteQty gets a reference to the given string and assigns it to the QuoteQty field.
func (o *UmAccountTradeListResponseInner) SetQuoteQty(v string) {
	o.QuoteQty = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetCommission() string {
	if o == nil || common.IsNil(o.Commission) {
		var ret string
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetCommissionOk() (*string, bool) {
	if o == nil || common.IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasCommission() bool {
	if o != nil && !common.IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given string and assigns it to the Commission field.
func (o *UmAccountTradeListResponseInner) SetCommission(v string) {
	o.Commission = &v
}

// GetCommissionAsset returns the CommissionAsset field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetCommissionAsset() string {
	if o == nil || common.IsNil(o.CommissionAsset) {
		var ret string
		return ret
	}
	return *o.CommissionAsset
}

// GetCommissionAssetOk returns a tuple with the CommissionAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetCommissionAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.CommissionAsset) {
		return nil, false
	}
	return o.CommissionAsset, true
}

// HasCommissionAsset returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasCommissionAsset() bool {
	if o != nil && !common.IsNil(o.CommissionAsset) {
		return true
	}

	return false
}

// SetCommissionAsset gets a reference to the given string and assigns it to the CommissionAsset field.
func (o *UmAccountTradeListResponseInner) SetCommissionAsset(v string) {
	o.CommissionAsset = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *UmAccountTradeListResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetBuyer() bool {
	if o == nil || common.IsNil(o.Buyer) {
		var ret bool
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetBuyerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Buyer) {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasBuyer() bool {
	if o != nil && !common.IsNil(o.Buyer) {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given bool and assigns it to the Buyer field.
func (o *UmAccountTradeListResponseInner) SetBuyer(v bool) {
	o.Buyer = &v
}

// GetMaker returns the Maker field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetMaker() bool {
	if o == nil || common.IsNil(o.Maker) {
		var ret bool
		return ret
	}
	return *o.Maker
}

// GetMakerOk returns a tuple with the Maker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetMakerOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Maker) {
		return nil, false
	}
	return o.Maker, true
}

// HasMaker returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasMaker() bool {
	if o != nil && !common.IsNil(o.Maker) {
		return true
	}

	return false
}

// SetMaker gets a reference to the given bool and assigns it to the Maker field.
func (o *UmAccountTradeListResponseInner) SetMaker(v bool) {
	o.Maker = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *UmAccountTradeListResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UmAccountTradeListResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *UmAccountTradeListResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *UmAccountTradeListResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

func (o UmAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UmAccountTradeListResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.QuoteQty) {
		toSerialize["quoteQty"] = o.QuoteQty
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
	if !common.IsNil(o.Buyer) {
		toSerialize["buyer"] = o.Buyer
	}
	if !common.IsNil(o.Maker) {
		toSerialize["maker"] = o.Maker
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UmAccountTradeListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varUmAccountTradeListResponseInner := _UmAccountTradeListResponseInner{}

	err = json.Unmarshal(data, &varUmAccountTradeListResponseInner)

	if err != nil {
		return err
	}

	*o = UmAccountTradeListResponseInner(varUmAccountTradeListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "id")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "side")
		delete(additionalProperties, "price")
		delete(additionalProperties, "qty")
		delete(additionalProperties, "realizedPnl")
		delete(additionalProperties, "quoteQty")
		delete(additionalProperties, "commission")
		delete(additionalProperties, "commissionAsset")
		delete(additionalProperties, "time")
		delete(additionalProperties, "buyer")
		delete(additionalProperties, "maker")
		delete(additionalProperties, "positionSide")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUmAccountTradeListResponseInner struct {
	value *UmAccountTradeListResponseInner
	isSet bool
}

func (v NullableUmAccountTradeListResponseInner) Get() *UmAccountTradeListResponseInner {
	return v.value
}

func (v *NullableUmAccountTradeListResponseInner) Set(val *UmAccountTradeListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUmAccountTradeListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUmAccountTradeListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUmAccountTradeListResponseInner(val *UmAccountTradeListResponseInner) *NullableUmAccountTradeListResponseInner {
	return &NullableUmAccountTradeListResponseInner{value: val, isSet: true}
}

func (v NullableUmAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUmAccountTradeListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
