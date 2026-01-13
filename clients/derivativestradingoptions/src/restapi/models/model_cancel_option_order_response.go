/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelOptionOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelOptionOrderResponse{}

// CancelOptionOrderResponse struct for CancelOptionOrderResponse
type CancelOptionOrderResponse struct {
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
	CreateDate           *int64  `json:"createDate,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	Status               *string `json:"status,omitempty"`
	AvgPrice             *string `json:"avgPrice,omitempty"`
	Source               *string `json:"source,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	PriceScale           *int64  `json:"priceScale,omitempty"`
	QuantityScale        *int64  `json:"quantityScale,omitempty"`
	OptionSide           *string `json:"optionSide,omitempty"`
	QuoteAsset           *string `json:"quoteAsset,omitempty"`
	Mmp                  *bool   `json:"mmp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelOptionOrderResponse CancelOptionOrderResponse

// NewCancelOptionOrderResponse instantiates a new CancelOptionOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOptionOrderResponse() *CancelOptionOrderResponse {
	this := CancelOptionOrderResponse{}
	return &this
}

// NewCancelOptionOrderResponseWithDefaults instantiates a new CancelOptionOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOptionOrderResponseWithDefaults() *CancelOptionOrderResponse {
	this := CancelOptionOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CancelOptionOrderResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CancelOptionOrderResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CancelOptionOrderResponse) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *CancelOptionOrderResponse) SetQuantity(v string) {
	o.Quantity = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *CancelOptionOrderResponse) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetFee returns the Fee field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetFee() string {
	if o == nil || common.IsNil(o.Fee) {
		var ret string
		return ret
	}
	return *o.Fee
}

// GetFeeOk returns a tuple with the Fee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetFeeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fee) {
		return nil, false
	}
	return o.Fee, true
}

// HasFee returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasFee() bool {
	if o != nil && !common.IsNil(o.Fee) {
		return true
	}

	return false
}

// SetFee gets a reference to the given string and assigns it to the Fee field.
func (o *CancelOptionOrderResponse) SetFee(v string) {
	o.Fee = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CancelOptionOrderResponse) SetSide(v string) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CancelOptionOrderResponse) SetType(v string) {
	o.Type = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CancelOptionOrderResponse) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *CancelOptionOrderResponse) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetPostOnly() bool {
	if o == nil || common.IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetPostOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasPostOnly() bool {
	if o != nil && !common.IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *CancelOptionOrderResponse) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetCreateDate() int64 {
	if o == nil || common.IsNil(o.CreateDate) {
		var ret int64
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetCreateDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasCreateDate() bool {
	if o != nil && !common.IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given int64 and assigns it to the CreateDate field.
func (o *CancelOptionOrderResponse) SetCreateDate(v int64) {
	o.CreateDate = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *CancelOptionOrderResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CancelOptionOrderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *CancelOptionOrderResponse) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetSource() string {
	if o == nil || common.IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetSourceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasSource() bool {
	if o != nil && !common.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *CancelOptionOrderResponse) SetSource(v string) {
	o.Source = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CancelOptionOrderResponse) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPriceScale returns the PriceScale field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetPriceScale() int64 {
	if o == nil || common.IsNil(o.PriceScale) {
		var ret int64
		return ret
	}
	return *o.PriceScale
}

// GetPriceScaleOk returns a tuple with the PriceScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetPriceScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.PriceScale) {
		return nil, false
	}
	return o.PriceScale, true
}

// HasPriceScale returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasPriceScale() bool {
	if o != nil && !common.IsNil(o.PriceScale) {
		return true
	}

	return false
}

// SetPriceScale gets a reference to the given int64 and assigns it to the PriceScale field.
func (o *CancelOptionOrderResponse) SetPriceScale(v int64) {
	o.PriceScale = &v
}

// GetQuantityScale returns the QuantityScale field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetQuantityScale() int64 {
	if o == nil || common.IsNil(o.QuantityScale) {
		var ret int64
		return ret
	}
	return *o.QuantityScale
}

// GetQuantityScaleOk returns a tuple with the QuantityScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetQuantityScaleOk() (*int64, bool) {
	if o == nil || common.IsNil(o.QuantityScale) {
		return nil, false
	}
	return o.QuantityScale, true
}

// HasQuantityScale returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasQuantityScale() bool {
	if o != nil && !common.IsNil(o.QuantityScale) {
		return true
	}

	return false
}

// SetQuantityScale gets a reference to the given int64 and assigns it to the QuantityScale field.
func (o *CancelOptionOrderResponse) SetQuantityScale(v int64) {
	o.QuantityScale = &v
}

// GetOptionSide returns the OptionSide field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetOptionSide() string {
	if o == nil || common.IsNil(o.OptionSide) {
		var ret string
		return ret
	}
	return *o.OptionSide
}

// GetOptionSideOk returns a tuple with the OptionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetOptionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.OptionSide) {
		return nil, false
	}
	return o.OptionSide, true
}

// HasOptionSide returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasOptionSide() bool {
	if o != nil && !common.IsNil(o.OptionSide) {
		return true
	}

	return false
}

// SetOptionSide gets a reference to the given string and assigns it to the OptionSide field.
func (o *CancelOptionOrderResponse) SetOptionSide(v string) {
	o.OptionSide = &v
}

// GetQuoteAsset returns the QuoteAsset field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetQuoteAsset() string {
	if o == nil || common.IsNil(o.QuoteAsset) {
		var ret string
		return ret
	}
	return *o.QuoteAsset
}

// GetQuoteAssetOk returns a tuple with the QuoteAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetQuoteAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteAsset) {
		return nil, false
	}
	return o.QuoteAsset, true
}

// HasQuoteAsset returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasQuoteAsset() bool {
	if o != nil && !common.IsNil(o.QuoteAsset) {
		return true
	}

	return false
}

// SetQuoteAsset gets a reference to the given string and assigns it to the QuoteAsset field.
func (o *CancelOptionOrderResponse) SetQuoteAsset(v string) {
	o.QuoteAsset = &v
}

// GetMmp returns the Mmp field value if set, zero value otherwise.
func (o *CancelOptionOrderResponse) GetMmp() bool {
	if o == nil || common.IsNil(o.Mmp) {
		var ret bool
		return ret
	}
	return *o.Mmp
}

// GetMmpOk returns a tuple with the Mmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOptionOrderResponse) GetMmpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Mmp) {
		return nil, false
	}
	return o.Mmp, true
}

// HasMmp returns a boolean if a field has been set.
func (o *CancelOptionOrderResponse) HasMmp() bool {
	if o != nil && !common.IsNil(o.Mmp) {
		return true
	}

	return false
}

// SetMmp gets a reference to the given bool and assigns it to the Mmp field.
func (o *CancelOptionOrderResponse) SetMmp(v bool) {
	o.Mmp = &v
}

func (o CancelOptionOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOptionOrderResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
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
	if !common.IsNil(o.Source) {
		toSerialize["source"] = o.Source
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

func (o *CancelOptionOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelOptionOrderResponse := _CancelOptionOrderResponse{}

	err = json.Unmarshal(data, &varCancelOptionOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelOptionOrderResponse(varCancelOptionOrderResponse)

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
		delete(additionalProperties, "createDate")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "status")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "source")
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

type NullableCancelOptionOrderResponse struct {
	value *CancelOptionOrderResponse
	isSet bool
}

func (v NullableCancelOptionOrderResponse) Get() *CancelOptionOrderResponse {
	return v.value
}

func (v *NullableCancelOptionOrderResponse) Set(val *CancelOptionOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOptionOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOptionOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOptionOrderResponse(val *CancelOptionOrderResponse) *NullableCancelOptionOrderResponse {
	return &NullableCancelOptionOrderResponse{value: val, isSet: true}
}

func (v NullableCancelOptionOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOptionOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
