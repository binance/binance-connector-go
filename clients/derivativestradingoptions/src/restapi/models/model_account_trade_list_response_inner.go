/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountTradeListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountTradeListResponseInner{}

// AccountTradeListResponseInner struct for AccountTradeListResponseInner
type AccountTradeListResponseInner struct {
	Id                   *int64  `json:"id,omitempty"`
	TradeId              *int64  `json:"tradeId,omitempty"`
	OrderId              *int64  `json:"orderId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	RealizedProfit       *string `json:"realizedProfit,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Type                 *string `json:"type,omitempty"`
	Liquidity            *string `json:"liquidity,omitempty"`
	Time                 *int64  `json:"time,omitempty"`
	PriceScale           *int64  `json:"priceScale,omitempty"`
	QuantityScale        *int64  `json:"quantityScale,omitempty"`
	OptionSide           *string `json:"optionSide,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountTradeListResponseInner AccountTradeListResponseInner

// NewAccountTradeListResponseInner instantiates a new AccountTradeListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTradeListResponseInner() *AccountTradeListResponseInner {
	this := AccountTradeListResponseInner{}
	return &this
}

// NewAccountTradeListResponseInnerWithDefaults instantiates a new AccountTradeListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTradeListResponseInnerWithDefaults() *AccountTradeListResponseInner {
	this := AccountTradeListResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetId() int64 {
	if o == nil || common.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AccountTradeListResponseInner) SetId(v int64) {
	o.Id = &v
}

// GetTradeId returns the TradeId field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetTradeId() int64 {
	if o == nil || common.IsNil(o.TradeId) {
		var ret int64
		return ret
	}
	return *o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetTradeIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeId) {
		return nil, false
	}
	return o.TradeId, true
}

// HasTradeId returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasTradeId() bool {
	if o != nil && !common.IsNil(o.TradeId) {
		return true
	}

	return false
}

// SetTradeId gets a reference to the given int64 and assigns it to the TradeId field.
func (o *AccountTradeListResponseInner) SetTradeId(v int64) {
	o.TradeId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *AccountTradeListResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountTradeListResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *AccountTradeListResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *AccountTradeListResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *AccountTradeListResponseInner) SetFee(v string) {
	o.Fee = &v
}

// GetRealizedProfit returns the RealizedProfit field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetRealizedProfit() string {
	if o == nil || common.IsNil(o.RealizedProfit) {
		var ret string
		return ret
	}
	return *o.RealizedProfit
}

// GetRealizedProfitOk returns a tuple with the RealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetRealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.RealizedProfit) {
		return nil, false
	}
	return o.RealizedProfit, true
}

// HasRealizedProfit returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasRealizedProfit() bool {
	if o != nil && !common.IsNil(o.RealizedProfit) {
		return true
	}

	return false
}

// SetRealizedProfit gets a reference to the given string and assigns it to the RealizedProfit field.
func (o *AccountTradeListResponseInner) SetRealizedProfit(v string) {
	o.RealizedProfit = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *AccountTradeListResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccountTradeListResponseInner) SetType(v string) {
	o.Type = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *AccountTradeListResponseInner) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *AccountTradeListResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *AccountTradeListResponseInner) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *AccountTradeListResponseInner) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetOptionSide() string {
	if o == nil || common.IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetOptionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasOptionSide() bool {
	if o != nil && !common.IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *AccountTradeListResponseInner) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *AccountTradeListResponseInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountTradeListResponseInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *AccountTradeListResponseInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *AccountTradeListResponseInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

func (o AccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountTradeListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.TradeId) {
		toSerialize["tradeId"] = o.TradeId
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.RealizedProfit) {
		toSerialize["realizedProfit"] = o.RealizedProfit
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.PriceScale) {
		toSerialize["priceScale"] = o.PriceScale
	}
	if !common.IsNil(o.QuantityScale) {
		toSerialize["quantityScale"] = o.QuantityScale
	}
	if !common.IsNil(o.OptionSide) {
		toSerialize["optionSide"] = o.OptionSide
	}
	if !common.IsNil(o.QuoteAsset) {
		toSerialize["quoteAsset"] = o.QuoteAsset
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountTradeListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAccountTradeListResponseInner := _AccountTradeListResponseInner{}

	err = json.Unmarshal(data, &varAccountTradeListResponseInner)

	if err != nil {
		return err
	}

	*o = AccountTradeListResponseInner(varAccountTradeListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "tradeId")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "realizedProfit")
		delete(additionalProperties, "side")
		delete(additionalProperties, "type")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "time")
		delete(additionalProperties, "priceScale")
		delete(additionalProperties, "quantityScale")
		delete(additionalProperties, "optionSide")
		delete(additionalProperties, "quoteAsset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountTradeListResponseInner struct {
	value *AccountTradeListResponseInner
	isSet bool
}

func (v NullableAccountTradeListResponseInner) Get() *AccountTradeListResponseInner {
	return v.value
}

func (v *NullableAccountTradeListResponseInner) Set(val *AccountTradeListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTradeListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTradeListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTradeListResponseInner(val *AccountTradeListResponseInner) *NullableAccountTradeListResponseInner {
	return &NullableAccountTradeListResponseInner{value: val, isSet: true}
}

func (v NullableAccountTradeListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTradeListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
