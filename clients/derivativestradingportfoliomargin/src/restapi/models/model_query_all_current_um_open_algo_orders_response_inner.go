/*
Portfolio Margin REST API

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryAllCurrentUmOpenAlgoOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentUmOpenAlgoOrdersResponseInner{}

// QueryAllCurrentUmOpenAlgoOrdersResponseInner struct for QueryAllCurrentUmOpenAlgoOrdersResponseInner
type QueryAllCurrentUmOpenAlgoOrdersResponseInner struct {
	AlgoId                  *int64  `json:"algoId,omitempty"`
	ClientAlgoId            *string `json:"clientAlgoId,omitempty"`
	AlgoType                *string `json:"algoType,omitempty"`
	OrderType               *string `json:"orderType,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	Side                    *string `json:"side,omitempty"`
	PositionSide            *string `json:"positionSide,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Quantity                *string `json:"quantity,omitempty"`
	AlgoStatus              *string `json:"algoStatus,omitempty"`
	TriggerPrice            *string `json:"triggerPrice,omitempty"`
	Price                   *string `json:"price,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	WorkingType             *string `json:"workingType,omitempty"`
	PriceMatch              *string `json:"priceMatch,omitempty"`
	ClosePosition           *bool   `json:"closePosition,omitempty"`
	PriceProtect            *bool   `json:"priceProtect,omitempty"`
	ReduceOnly              *bool   `json:"reduceOnly,omitempty"`
	CreateTime              *int64  `json:"createTime,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	TriggerTime             *int64  `json:"triggerTime,omitempty"`
	GoodTillDate            *int64  `json:"goodTillDate,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryAllCurrentUmOpenAlgoOrdersResponseInner QueryAllCurrentUmOpenAlgoOrdersResponseInner

// NewQueryAllCurrentUmOpenAlgoOrdersResponseInner instantiates a new QueryAllCurrentUmOpenAlgoOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentUmOpenAlgoOrdersResponseInner() *QueryAllCurrentUmOpenAlgoOrdersResponseInner {
	this := QueryAllCurrentUmOpenAlgoOrdersResponseInner{}
	return &this
}

// NewQueryAllCurrentUmOpenAlgoOrdersResponseInnerWithDefaults instantiates a new QueryAllCurrentUmOpenAlgoOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentUmOpenAlgoOrdersResponseInnerWithDefaults() *QueryAllCurrentUmOpenAlgoOrdersResponseInner {
	this := QueryAllCurrentUmOpenAlgoOrdersResponseInner{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetAlgoType returns the AlgoType field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoType() string {
	if o == nil || common.IsNil(o.AlgoType) {
		var ret string
		return ret
	}
	return *o.AlgoType
}

// GetAlgoTypeOk returns a tuple with the AlgoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoType) {
		return nil, false
	}
	return o.AlgoType, true
}

// HasAlgoType returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasAlgoType() bool {
	if o != nil && !common.IsNil(o.AlgoType) {
		return true
	}

	return false
}

// SetAlgoType gets a reference to the given string and assigns it to the AlgoType field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetAlgoType(v string) {
	o.AlgoType = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetAlgoStatus returns the AlgoStatus field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoStatus() string {
	if o == nil || common.IsNil(o.AlgoStatus) {
		var ret string
		return ret
	}
	return *o.AlgoStatus
}

// GetAlgoStatusOk returns a tuple with the AlgoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetAlgoStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoStatus) {
		return nil, false
	}
	return o.AlgoStatus, true
}

// HasAlgoStatus returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasAlgoStatus() bool {
	if o != nil && !common.IsNil(o.AlgoStatus) {
		return true
	}

	return false
}

// SetAlgoStatus gets a reference to the given string and assigns it to the AlgoStatus field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetAlgoStatus(v string) {
	o.AlgoStatus = &v
}

// GetTriggerPrice returns the TriggerPrice field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTriggerPrice() string {
	if o == nil || common.IsNil(o.TriggerPrice) {
		var ret string
		return ret
	}
	return *o.TriggerPrice
}

// GetTriggerPriceOk returns a tuple with the TriggerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTriggerPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TriggerPrice) {
		return nil, false
	}
	return o.TriggerPrice, true
}

// HasTriggerPrice returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasTriggerPrice() bool {
	if o != nil && !common.IsNil(o.TriggerPrice) {
		return true
	}

	return false
}

// SetTriggerPrice gets a reference to the given string and assigns it to the TriggerPrice field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetTriggerPrice(v string) {
	o.TriggerPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetClosePosition() bool {
	if o == nil || common.IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetClosePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasClosePosition() bool {
	if o != nil && !common.IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetTriggerTime returns the TriggerTime field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTriggerTime() int64 {
	if o == nil || common.IsNil(o.TriggerTime) {
		var ret int64
		return ret
	}
	return *o.TriggerTime
}

// GetTriggerTimeOk returns a tuple with the TriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetTriggerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TriggerTime) {
		return nil, false
	}
	return o.TriggerTime, true
}

// HasTriggerTime returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasTriggerTime() bool {
	if o != nil && !common.IsNil(o.TriggerTime) {
		return true
	}

	return false
}

// SetTriggerTime gets a reference to the given int64 and assigns it to the TriggerTime field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetTriggerTime(v int64) {
	o.TriggerTime = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

func (o QueryAllCurrentUmOpenAlgoOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentUmOpenAlgoOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
	}
	if !common.IsNil(o.AlgoType) {
		toSerialize["algoType"] = o.AlgoType
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.AlgoStatus) {
		toSerialize["algoStatus"] = o.AlgoStatus
	}
	if !common.IsNil(o.TriggerPrice) {
		toSerialize["triggerPrice"] = o.TriggerPrice
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}
	if !common.IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
	}
	if !common.IsNil(o.PriceProtect) {
		toSerialize["priceProtect"] = o.PriceProtect
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.TriggerTime) {
		toSerialize["triggerTime"] = o.TriggerTime
	}
	if !common.IsNil(o.GoodTillDate) {
		toSerialize["goodTillDate"] = o.GoodTillDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryAllCurrentUmOpenAlgoOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryAllCurrentUmOpenAlgoOrdersResponseInner := _QueryAllCurrentUmOpenAlgoOrdersResponseInner{}

	err = json.Unmarshal(data, &varQueryAllCurrentUmOpenAlgoOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = QueryAllCurrentUmOpenAlgoOrdersResponseInner(varQueryAllCurrentUmOpenAlgoOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "clientAlgoId")
		delete(additionalProperties, "algoType")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "algoStatus")
		delete(additionalProperties, "triggerPrice")
		delete(additionalProperties, "price")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceMatch")
		delete(additionalProperties, "closePosition")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "createTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "triggerTime")
		delete(additionalProperties, "goodTillDate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner struct {
	value *QueryAllCurrentUmOpenAlgoOrdersResponseInner
	isSet bool
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) Get() *QueryAllCurrentUmOpenAlgoOrdersResponseInner {
	return v.value
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) Set(val *QueryAllCurrentUmOpenAlgoOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAllCurrentUmOpenAlgoOrdersResponseInner(val *QueryAllCurrentUmOpenAlgoOrdersResponseInner) *NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner {
	return &NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner{value: val, isSet: true}
}

func (v NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentUmOpenAlgoOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
