/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentMarginOpenOrderResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentMarginOpenOrderResponseInner{}

// QueryCurrentMarginOpenOrderResponseInner struct for QueryCurrentMarginOpenOrderResponseInner
type QueryCurrentMarginOpenOrderResponseInner struct {
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	CummulativeQuoteQty     *string `json:"cummulativeQuoteQty,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	IcebergQty              *string `json:"icebergQty,omitempty"`
	IsWorking               *bool   `json:"isWorking,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	Price                   *string `json:"price,omitempty"`
	Side                    *string `json:"side,omitempty"`
	Status                  *string `json:"status,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	Time                    *int64  `json:"time,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	AccountId               *int64  `json:"accountId,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	PreventedMatchId        *string `json:"preventedMatchId,omitempty"`
	PreventedQuantity       *string `json:"preventedQuantity,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryCurrentMarginOpenOrderResponseInner QueryCurrentMarginOpenOrderResponseInner

// NewQueryCurrentMarginOpenOrderResponseInner instantiates a new QueryCurrentMarginOpenOrderResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentMarginOpenOrderResponseInner() *QueryCurrentMarginOpenOrderResponseInner {
	this := QueryCurrentMarginOpenOrderResponseInner{}
	return &this
}

// NewQueryCurrentMarginOpenOrderResponseInnerWithDefaults instantiates a new QueryCurrentMarginOpenOrderResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentMarginOpenOrderResponseInnerWithDefaults() *QueryCurrentMarginOpenOrderResponseInner {
	this := QueryCurrentMarginOpenOrderResponseInner{}
	return &this
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCummulativeQuoteQty returns the CummulativeQuoteQty field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetCummulativeQuoteQty() string {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		var ret string
		return ret
	}
	return *o.CummulativeQuoteQty
}

// GetCummulativeQuoteQtyOk returns a tuple with the CummulativeQuoteQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetCummulativeQuoteQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CummulativeQuoteQty) {
		return nil, false
	}
	return o.CummulativeQuoteQty, true
}

// HasCummulativeQuoteQty returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasCummulativeQuoteQty() bool {
	if o != nil && !common.IsNil(o.CummulativeQuoteQty) {
		return true
	}

	return false
}

// SetCummulativeQuoteQty gets a reference to the given string and assigns it to the CummulativeQuoteQty field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetCummulativeQuoteQty(v string) {
	o.CummulativeQuoteQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetIcebergQty returns the IcebergQty field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetIcebergQty() string {
	if o == nil || common.IsNil(o.IcebergQty) {
		var ret string
		return ret
	}
	return *o.IcebergQty
}

// GetIcebergQtyOk returns a tuple with the IcebergQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetIcebergQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.IcebergQty) {
		return nil, false
	}
	return o.IcebergQty, true
}

// HasIcebergQty returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasIcebergQty() bool {
	if o != nil && !common.IsNil(o.IcebergQty) {
		return true
	}

	return false
}

// SetIcebergQty gets a reference to the given string and assigns it to the IcebergQty field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetIcebergQty(v string) {
	o.IcebergQty = &v
}

// GetIsWorking returns the IsWorking field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetIsWorking() bool {
	if o == nil || common.IsNil(o.IsWorking) {
		var ret bool
		return ret
	}
	return *o.IsWorking
}

// GetIsWorkingOk returns a tuple with the IsWorking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetIsWorkingOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsWorking) {
		return nil, false
	}
	return o.IsWorking, true
}

// HasIsWorking returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasIsWorking() bool {
	if o != nil && !common.IsNil(o.IsWorking) {
		return true
	}

	return false
}

// SetIsWorking gets a reference to the given bool and assigns it to the IsWorking field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetIsWorking(v bool) {
	o.IsWorking = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetType(v string) {
	o.Type = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetAccountId() int64 {
	if o == nil || common.IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetAccountIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasAccountId() bool {
	if o != nil && !common.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetPreventedMatchId returns the PreventedMatchId field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPreventedMatchId() string {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		var ret string
		return ret
	}
	return *o.PreventedMatchId
}

// GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPreventedMatchIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreventedMatchId) {
		return nil, false
	}
	return o.PreventedMatchId, true
}

// HasPreventedMatchId returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasPreventedMatchId() bool {
	if o != nil && !common.IsNil(o.PreventedMatchId) {
		return true
	}

	return false
}

// SetPreventedMatchId gets a reference to the given string and assigns it to the PreventedMatchId field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetPreventedMatchId(v string) {
	o.PreventedMatchId = &v
}

// GetPreventedQuantity returns the PreventedQuantity field value if set, zero value otherwise.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPreventedQuantity() string {
	if o == nil || common.IsNil(o.PreventedQuantity) {
		var ret string
		return ret
	}
	return *o.PreventedQuantity
}

// GetPreventedQuantityOk returns a tuple with the PreventedQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) GetPreventedQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.PreventedQuantity) {
		return nil, false
	}
	return o.PreventedQuantity, true
}

// HasPreventedQuantity returns a boolean if a field has been set.
func (o *QueryCurrentMarginOpenOrderResponseInner) HasPreventedQuantity() bool {
	if o != nil && !common.IsNil(o.PreventedQuantity) {
		return true
	}

	return false
}

// SetPreventedQuantity gets a reference to the given string and assigns it to the PreventedQuantity field.
func (o *QueryCurrentMarginOpenOrderResponseInner) SetPreventedQuantity(v string) {
	o.PreventedQuantity = &v
}

func (o QueryCurrentMarginOpenOrderResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentMarginOpenOrderResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.CummulativeQuoteQty) {
		toSerialize["cummulativeQuoteQty"] = o.CummulativeQuoteQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.IcebergQty) {
		toSerialize["icebergQty"] = o.IcebergQty
	}
	if !common.IsNil(o.IsWorking) {
		toSerialize["isWorking"] = o.IsWorking
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.PreventedMatchId) {
		toSerialize["preventedMatchId"] = o.PreventedMatchId
	}
	if !common.IsNil(o.PreventedQuantity) {
		toSerialize["preventedQuantity"] = o.PreventedQuantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryCurrentMarginOpenOrderResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryCurrentMarginOpenOrderResponseInner := _QueryCurrentMarginOpenOrderResponseInner{}

	err = json.Unmarshal(data, &varQueryCurrentMarginOpenOrderResponseInner)

	if err != nil {
		return err
	}

	*o = QueryCurrentMarginOpenOrderResponseInner(varQueryCurrentMarginOpenOrderResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "cummulativeQuoteQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "icebergQty")
		delete(additionalProperties, "isWorking")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "price")
		delete(additionalProperties, "side")
		delete(additionalProperties, "status")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "preventedMatchId")
		delete(additionalProperties, "preventedQuantity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCurrentMarginOpenOrderResponseInner struct {
	value *QueryCurrentMarginOpenOrderResponseInner
	isSet bool
}

func (v NullableQueryCurrentMarginOpenOrderResponseInner) Get() *QueryCurrentMarginOpenOrderResponseInner {
	return v.value
}

func (v *NullableQueryCurrentMarginOpenOrderResponseInner) Set(val *QueryCurrentMarginOpenOrderResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentMarginOpenOrderResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentMarginOpenOrderResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCurrentMarginOpenOrderResponseInner(val *QueryCurrentMarginOpenOrderResponseInner) *NullableQueryCurrentMarginOpenOrderResponseInner {
	return &NullableQueryCurrentMarginOpenOrderResponseInner{value: val, isSet: true}
}

func (v NullableQueryCurrentMarginOpenOrderResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentMarginOpenOrderResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
