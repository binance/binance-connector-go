/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the NewUmConditionalOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewUmConditionalOrderResponse{}

// NewUmConditionalOrderResponse struct for NewUmConditionalOrderResponse
type NewUmConditionalOrderResponse struct {
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
	TimeInForce             *string `json:"timeInForce,omitempty"`
	ActivatePrice           *string `json:"activatePrice,omitempty"`
	PriceRate               *string `json:"priceRate,omitempty"`
	BookTime                *int64  `json:"bookTime,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	WorkingType             *string `json:"workingType,omitempty"`
	PriceProtect            *bool   `json:"priceProtect,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            *int64  `json:"goodTillDate,omitempty"`
	PriceMatch              *string `json:"priceMatch,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _NewUmConditionalOrderResponse NewUmConditionalOrderResponse

// NewNewUmConditionalOrderResponse instantiates a new NewUmConditionalOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUmConditionalOrderResponse() *NewUmConditionalOrderResponse {
	this := NewUmConditionalOrderResponse{}
	return &this
}

// NewNewUmConditionalOrderResponseWithDefaults instantiates a new NewUmConditionalOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUmConditionalOrderResponseWithDefaults() *NewUmConditionalOrderResponse {
	this := NewUmConditionalOrderResponse{}
	return &this
}

// GetNewClientStrategyId returns the NewClientStrategyId field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetNewClientStrategyId() string {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		var ret string
		return ret
	}
	return *o.NewClientStrategyId
}

// GetNewClientStrategyIdOk returns a tuple with the NewClientStrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetNewClientStrategyIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewClientStrategyId) {
		return nil, false
	}
	return o.NewClientStrategyId, true
}

// HasNewClientStrategyId returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasNewClientStrategyId() bool {
	if o != nil && !common.IsNil(o.NewClientStrategyId) {
		return true
	}

	return false
}

// SetNewClientStrategyId gets a reference to the given string and assigns it to the NewClientStrategyId field.
func (o *NewUmConditionalOrderResponse) SetNewClientStrategyId(v string) {
	o.NewClientStrategyId = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetStrategyId() int64 {
	if o == nil || common.IsNil(o.StrategyId) {
		var ret int64
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetStrategyIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StrategyId) {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasStrategyId() bool {
	if o != nil && !common.IsNil(o.StrategyId) {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given int64 and assigns it to the StrategyId field.
func (o *NewUmConditionalOrderResponse) SetStrategyId(v int64) {
	o.StrategyId = &v
}

// GetStrategyStatus returns the StrategyStatus field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetStrategyStatus() string {
	if o == nil || common.IsNil(o.StrategyStatus) {
		var ret string
		return ret
	}
	return *o.StrategyStatus
}

// GetStrategyStatusOk returns a tuple with the StrategyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetStrategyStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyStatus) {
		return nil, false
	}
	return o.StrategyStatus, true
}

// HasStrategyStatus returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasStrategyStatus() bool {
	if o != nil && !common.IsNil(o.StrategyStatus) {
		return true
	}

	return false
}

// SetStrategyStatus gets a reference to the given string and assigns it to the StrategyStatus field.
func (o *NewUmConditionalOrderResponse) SetStrategyStatus(v string) {
	o.StrategyStatus = &v
}

// GetStrategyType returns the StrategyType field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetStrategyType() string {
	if o == nil || common.IsNil(o.StrategyType) {
		var ret string
		return ret
	}
	return *o.StrategyType
}

// GetStrategyTypeOk returns a tuple with the StrategyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetStrategyTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.StrategyType) {
		return nil, false
	}
	return o.StrategyType, true
}

// HasStrategyType returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasStrategyType() bool {
	if o != nil && !common.IsNil(o.StrategyType) {
		return true
	}

	return false
}

// SetStrategyType gets a reference to the given string and assigns it to the StrategyType field.
func (o *NewUmConditionalOrderResponse) SetStrategyType(v string) {
	o.StrategyType = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *NewUmConditionalOrderResponse) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *NewUmConditionalOrderResponse) SetPrice(v string) {
	o.Price = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *NewUmConditionalOrderResponse) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *NewUmConditionalOrderResponse) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *NewUmConditionalOrderResponse) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *NewUmConditionalOrderResponse) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *NewUmConditionalOrderResponse) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *NewUmConditionalOrderResponse) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetActivatePrice() string {
	if o == nil || common.IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetActivatePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasActivatePrice() bool {
	if o != nil && !common.IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *NewUmConditionalOrderResponse) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetPriceRate() string {
	if o == nil || common.IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetPriceRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasPriceRate() bool {
	if o != nil && !common.IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *NewUmConditionalOrderResponse) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetBookTime returns the BookTime field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetBookTime() int64 {
	if o == nil || common.IsNil(o.BookTime) {
		var ret int64
		return ret
	}
	return *o.BookTime
}

// GetBookTimeOk returns a tuple with the BookTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetBookTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.BookTime) {
		return nil, false
	}
	return o.BookTime, true
}

// HasBookTime returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasBookTime() bool {
	if o != nil && !common.IsNil(o.BookTime) {
		return true
	}

	return false
}

// SetBookTime gets a reference to the given int64 and assigns it to the BookTime field.
func (o *NewUmConditionalOrderResponse) SetBookTime(v int64) {
	o.BookTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *NewUmConditionalOrderResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *NewUmConditionalOrderResponse) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *NewUmConditionalOrderResponse) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *NewUmConditionalOrderResponse) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *NewUmConditionalOrderResponse) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *NewUmConditionalOrderResponse) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUmConditionalOrderResponse) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *NewUmConditionalOrderResponse) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *NewUmConditionalOrderResponse) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

func (o NewUmConditionalOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewUmConditionalOrderResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !common.IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
	}
	if !common.IsNil(o.BookTime) {
		toSerialize["bookTime"] = o.BookTime
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
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
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NewUmConditionalOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varNewUmConditionalOrderResponse := _NewUmConditionalOrderResponse{}

	err = json.Unmarshal(data, &varNewUmConditionalOrderResponse)

	if err != nil {
		return err
	}

	*o = NewUmConditionalOrderResponse(varNewUmConditionalOrderResponse)

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
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "activatePrice")
		delete(additionalProperties, "priceRate")
		delete(additionalProperties, "bookTime")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "goodTillDate")
		delete(additionalProperties, "priceMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNewUmConditionalOrderResponse struct {
	value *NewUmConditionalOrderResponse
	isSet bool
}

func (v NullableNewUmConditionalOrderResponse) Get() *NewUmConditionalOrderResponse {
	return v.value
}

func (v *NullableNewUmConditionalOrderResponse) Set(val *NewUmConditionalOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUmConditionalOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUmConditionalOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUmConditionalOrderResponse(val *NewUmConditionalOrderResponse) *NullableNewUmConditionalOrderResponse {
	return &NullableNewUmConditionalOrderResponse{value: val, isSet: true}
}

func (v NullableNewUmConditionalOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUmConditionalOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
