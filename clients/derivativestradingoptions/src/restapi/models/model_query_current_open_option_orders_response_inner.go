/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentOpenOptionOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentOpenOptionOrdersResponseInner{}

// QueryCurrentOpenOptionOrdersResponseInner struct for QueryCurrentOpenOptionOrdersResponseInner
type QueryCurrentOpenOptionOrdersResponseInner struct {
	OrderId              *int64  `json:"orderId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	ExecutedQty          *string `json:"executedQty,omitempty"`
	Fee                  *string `json:"fee,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Type                 *string `json:"type,omitempty"`
	TimeInForce          *string `json:"timeInForce,omitempty"`
	ReduceOnly           *bool   `json:"reduceOnly,omitempty"`
	PostOnly             *bool   `json:"postOnly,omitempty"`
	CreateTime           *int64  `json:"createTime,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	Status               *string `json:"status,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	PriceScale           *int64  `json:"priceScale,omitempty"`
	QuantityScale        *int64  `json:"quantityScale,omitempty"`
	OptionSide           *string `json:"optionSide,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	Mmp                  *bool   `json:"mmp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCurrentOpenOptionOrdersResponseInner QueryCurrentOpenOptionOrdersResponseInner

// NewQueryCurrentOpenOptionOrdersResponseInner instantiates a new QueryCurrentOpenOptionOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentOpenOptionOrdersResponseInner() *QueryCurrentOpenOptionOrdersResponseInner {
	this := QueryCurrentOpenOptionOrdersResponseInner{}
	return &this
}

// NewQueryCurrentOpenOptionOrdersResponseInnerWithDefaults instantiates a new QueryCurrentOpenOptionOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentOpenOptionOrdersResponseInnerWithDefaults() *QueryCurrentOpenOptionOrdersResponseInner {
	this := QueryCurrentOpenOptionOrdersResponseInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetFee(v string) {
	o.Fee = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetType(v string) {
	o.Type = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPostOnly() bool {
	if o == nil || common.IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPostOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasPostOnly() bool {
	if o != nil && !common.IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetOptionSide() string {
	if o == nil || common.IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetOptionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasOptionSide() bool {
	if o != nil && !common.IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetMmp returns the Mmp field value if set, zero value otherwise.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetMmp() bool {
	if o == nil || common.IsNil(o.Mmp) {
		var ret bool
		return ret
	}
	return *o.Mmp
}

// GetMmpOk returns a tuple with the Mmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) GetMmpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Mmp) {
		return nil, false
	}
	return o.Mmp, true
}

// HasMmp returns a boolean if a field has been set.
func (o *QueryCurrentOpenOptionOrdersResponseInner) HasMmp() bool {
	if o != nil && !common.IsNil(o.Mmp) {
		return true
	}

	return false
}

// SetMmp gets a reference to the given bool and assigns it to the Mmp field.
func (o *QueryCurrentOpenOptionOrdersResponseInner) SetMmp(v bool) {
	o.Mmp = &v
}

func (o QueryCurrentOpenOptionOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentOpenOptionOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.Fee) {
		toSerialize["fee"] = o.Fee
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.PostOnly) {
		toSerialize["postOnly"] = o.PostOnly
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
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
	if !common.IsNil(o.Mmp) {
		toSerialize["mmp"] = o.Mmp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCurrentOpenOptionOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCurrentOpenOptionOrdersResponseInner := _QueryCurrentOpenOptionOrdersResponseInner{}

	err = json.Unmarshal(data, &varQueryCurrentOpenOptionOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCurrentOpenOptionOrdersResponseInner(varQueryCurrentOpenOptionOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "fee")
		delete(additionalProperties, "side")
		delete(additionalProperties, "type")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "postOnly")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "status")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "priceScale")
		delete(additionalProperties, "quantityScale")
		delete(additionalProperties, "optionSide")
		delete(additionalProperties, "quoteAsset")
		delete(additionalProperties, "mmp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCurrentOpenOptionOrdersResponseInner struct {
	value *QueryCurrentOpenOptionOrdersResponseInner
	isSet bool
}

func (v NullableQueryCurrentOpenOptionOrdersResponseInner) Get() *QueryCurrentOpenOptionOrdersResponseInner {
	return v.value
}

func (v *NullableQueryCurrentOpenOptionOrdersResponseInner) Set(val *QueryCurrentOpenOptionOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentOpenOptionOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentOpenOptionOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCurrentOpenOptionOrdersResponseInner(val *QueryCurrentOpenOptionOrdersResponseInner) *NullableQueryCurrentOpenOptionOrdersResponseInner {
	return &NullableQueryCurrentOpenOptionOrdersResponseInner{value: val, isSet: true}
}

func (v NullableQueryCurrentOpenOptionOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentOpenOptionOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
