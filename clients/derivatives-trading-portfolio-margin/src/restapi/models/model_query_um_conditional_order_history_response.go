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

// checks if the QueryUmConditionalOrderHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUmConditionalOrderHistoryResponse{}

// QueryUmConditionalOrderHistoryResponse struct for QueryUmConditionalOrderHistoryResponse
type QueryUmConditionalOrderHistoryResponse struct {
	NewClientStrategyId     *string `json:"newClientStrategyId,omitempty"`
	StrategyId              *int64  `json:"strategyId,omitempty"`
	StrategyStatus          *string `json:"strategyStatus,omitempty"`
	StrategyType            *string `json:"strategyType,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	Price                   *string `json:"price,omitempty"`
	ReduceOnly              *bool   `json:"reduceOnly,omitempty"`
	Side                    *string `json:"side,omitempty"`
	PositionSide            *string `json:"positionSide,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	Status                  *string `json:"status,omitempty"`
	BookTime                *int64  `json:"bookTime,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	TriggerTime             *int64  `json:"triggerTime,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	ActivatePrice           *string `json:"activatePrice,omitempty"`
	PriceRate               *string `json:"priceRate,omitempty"`
	WorkingType             *string `json:"workingType,omitempty"`
	PriceProtect            *bool   `json:"priceProtect,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            *int64  `json:"goodTillDate,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryUmConditionalOrderHistoryResponse QueryUmConditionalOrderHistoryResponse

// NewQueryUmConditionalOrderHistoryResponse instantiates a new QueryUmConditionalOrderHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUmConditionalOrderHistoryResponse() *QueryUmConditionalOrderHistoryResponse {
	this := QueryUmConditionalOrderHistoryResponse{}
	return &this
}

// NewQueryUmConditionalOrderHistoryResponseWithDefaults instantiates a new QueryUmConditionalOrderHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUmConditionalOrderHistoryResponseWithDefaults() *QueryUmConditionalOrderHistoryResponse {
	this := QueryUmConditionalOrderHistoryResponse{}
	return &this
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetNewClientStrategyId() string {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasNewClientStrategyId() bool {
	if o != nil && !common.IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *QueryUmConditionalOrderHistoryResponse) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyId() int64 {
	if o == nil || common.IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasStrategyId() bool {
	if o != nil && !common.IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *QueryUmConditionalOrderHistoryResponse) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyStatus() string {
	if o == nil || common.IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasStrategyStatus() bool {
	if o != nil && !common.IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *QueryUmConditionalOrderHistoryResponse) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyType() string {
	if o == nil || common.IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetStrategyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasStrategyType() bool {
	if o != nil && !common.IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *QueryUmConditionalOrderHistoryResponse) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QueryUmConditionalOrderHistoryResponse) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryUmConditionalOrderHistoryResponse) SetPrice(v string) {
	o.Price = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *QueryUmConditionalOrderHistoryResponse) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryUmConditionalOrderHistoryResponse) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryUmConditionalOrderHistoryResponse) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *QueryUmConditionalOrderHistoryResponse) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryUmConditionalOrderHistoryResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryUmConditionalOrderHistoryResponse) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryUmConditionalOrderHistoryResponse) SetStatus(v string) {
	o.Status = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *QueryUmConditionalOrderHistoryResponse) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryUmConditionalOrderHistoryResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetTriggerTime returns the TriggerTime field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetTriggerTime() int64 {
	if o == nil || common.IsNil(o.TriggerTime) {
		var ret int64
		return ret
	}
	return *o.TriggerTime
}

// GetTriggerTimeOk returns a tuple with the TriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetTriggerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TriggerTime) {
		return nil, false
	}
	return o.TriggerTime, true
}

// HasTriggerTime returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasTriggerTime() bool {
	if o != nil && !common.IsNil(o.TriggerTime) {
		return true
	}

	return false
}

// SetTriggerTime gets a reference to the given int64 and assigns it to the TriggerTime field.
func (o *QueryUmConditionalOrderHistoryResponse) SetTriggerTime(v int64) {
	o.TriggerTime = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryUmConditionalOrderHistoryResponse) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryUmConditionalOrderHistoryResponse) SetType(v string) {
	o.Type = &v
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetActivatePrice() string {
	if o == nil || common.IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetActivatePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasActivatePrice() bool {
	if o != nil && !common.IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *QueryUmConditionalOrderHistoryResponse) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetPriceRate() string {
	if o == nil || common.IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetPriceRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasPriceRate() bool {
	if o != nil && !common.IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *QueryUmConditionalOrderHistoryResponse) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *QueryUmConditionalOrderHistoryResponse) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *QueryUmConditionalOrderHistoryResponse) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryUmConditionalOrderHistoryResponse) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *QueryUmConditionalOrderHistoryResponse) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUmConditionalOrderHistoryResponse) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *QueryUmConditionalOrderHistoryResponse) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *QueryUmConditionalOrderHistoryResponse) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

func (o QueryUmConditionalOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUmConditionalOrderHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.NewClientStrategyId) {
		toSerialize["newClientStrategyId"] = o.NewClientStrategyId
	}
	if !common.IsNil(o.StrategyId) {
		toSerialize["strategyId"] = o.StrategyId
	}
	if !common.IsNil(o.StrategyStatus) {
		toSerialize["strategyStatus"] = o.StrategyStatus
	}
	if !common.IsNil(o.StrategyType) {
		toSerialize["strategyType"] = o.StrategyType
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.TriggerTime) {
		toSerialize["triggerTime"] = o.TriggerTime
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !common.IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
	}
	if !common.IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	if !common.IsNil(o.PriceProtect) {
		toSerialize["priceProtect"] = o.PriceProtect
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.GoodTillDate) {
		toSerialize["goodTillDate"] = o.GoodTillDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUmConditionalOrderHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUmConditionalOrderHistoryResponse := _QueryUmConditionalOrderHistoryResponse{}

	err = json.Unmarshal(data, &varQueryUmConditionalOrderHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryUmConditionalOrderHistoryResponse(varQueryUmConditionalOrderHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "newClientStrategyId")
		delete(additionalProperties, "strategyId")
		delete(additionalProperties, "strategyStatus")
		delete(additionalProperties, "strategyType")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "price")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "status")
		delete(additionalProperties, "bookTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "triggerTime")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "activatePrice")
		delete(additionalProperties, "priceRate")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "goodTillDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUmConditionalOrderHistoryResponse struct {
	value *QueryUmConditionalOrderHistoryResponse
	isSet bool
}

func (v NullableQueryUmConditionalOrderHistoryResponse) Get() *QueryUmConditionalOrderHistoryResponse {
	return v.value
}

func (v *NullableQueryUmConditionalOrderHistoryResponse) Set(val *QueryUmConditionalOrderHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUmConditionalOrderHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUmConditionalOrderHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUmConditionalOrderHistoryResponse(val *QueryUmConditionalOrderHistoryResponse) *NullableQueryUmConditionalOrderHistoryResponse {
	return &NullableQueryUmConditionalOrderHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUmConditionalOrderHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUmConditionalOrderHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
