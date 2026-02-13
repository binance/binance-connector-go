/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryAllCurrentUmOpenConditionalOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryAllCurrentUmOpenConditionalOrdersResponseInner{}

// QueryAllCurrentUmOpenConditionalOrdersResponseInner struct for QueryAllCurrentUmOpenConditionalOrdersResponseInner
type QueryAllCurrentUmOpenConditionalOrdersResponseInner struct {
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
	BookTime                *int64  `json:"bookTime,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	ActivatePrice           *string `json:"activatePrice,omitempty"`
	PriceRate               *string `json:"priceRate,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            *int64  `json:"goodTillDate,omitempty"`
	PriceMatch              *string `json:"priceMatch,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryAllCurrentUmOpenConditionalOrdersResponseInner QueryAllCurrentUmOpenConditionalOrdersResponseInner

// NewQueryAllCurrentUmOpenConditionalOrdersResponseInner instantiates a new QueryAllCurrentUmOpenConditionalOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAllCurrentUmOpenConditionalOrdersResponseInner() *QueryAllCurrentUmOpenConditionalOrdersResponseInner {
	this := QueryAllCurrentUmOpenConditionalOrdersResponseInner{}
	return &this
}

// NewQueryAllCurrentUmOpenConditionalOrdersResponseInnerWithDefaults instantiates a new QueryAllCurrentUmOpenConditionalOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAllCurrentUmOpenConditionalOrdersResponseInnerWithDefaults() *QueryAllCurrentUmOpenConditionalOrdersResponseInner {
	this := QueryAllCurrentUmOpenConditionalOrdersResponseInner{}
	return &this
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetNewClientStrategyId() string {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasNewClientStrategyId() bool {
	if o != nil && !common.IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyId() int64 {
	if o == nil || common.IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasStrategyId() bool {
	if o != nil && !common.IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyStatus() string {
	if o == nil || common.IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasStrategyStatus() bool {
	if o != nil && !common.IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyType() string {
	if o == nil || common.IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStrategyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasStrategyType() bool {
	if o != nil && !common.IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetActivatePrice() string {
	if o == nil || common.IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetActivatePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasActivatePrice() bool {
	if o != nil && !common.IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPriceRate() string {
	if o == nil || common.IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPriceRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasPriceRate() bool {
	if o != nil && !common.IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

func (o QueryAllCurrentUmOpenConditionalOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAllCurrentUmOpenConditionalOrdersResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !common.IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.GoodTillDate) {
		toSerialize["goodTillDate"] = o.GoodTillDate
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryAllCurrentUmOpenConditionalOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varQueryAllCurrentUmOpenConditionalOrdersResponseInner := _QueryAllCurrentUmOpenConditionalOrdersResponseInner{}

	err = json.Unmarshal(data, &varQueryAllCurrentUmOpenConditionalOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = QueryAllCurrentUmOpenConditionalOrdersResponseInner(varQueryAllCurrentUmOpenConditionalOrdersResponseInner)

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
		delete(additionalProperties, "bookTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "activatePrice")
		delete(additionalProperties, "priceRate")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "goodTillDate")
		delete(additionalProperties, "priceMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner struct {
	value *QueryAllCurrentUmOpenConditionalOrdersResponseInner
	isSet bool
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) Get() *QueryAllCurrentUmOpenConditionalOrdersResponseInner {
	return v.value
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) Set(val *QueryAllCurrentUmOpenConditionalOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAllCurrentUmOpenConditionalOrdersResponseInner(val *QueryAllCurrentUmOpenConditionalOrdersResponseInner) *NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner {
	return &NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner{value: val, isSet: true}
}

func (v NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAllCurrentUmOpenConditionalOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
