/*
Futures (COIN-M) REST API

Access market data, manage accounts, and trade COIN-M perpetual and delivery futures.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PlaceMultipleOrdersBatchOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersBatchOrdersParameterInner{}

// PlaceMultipleOrdersBatchOrdersParameterInner struct for PlaceMultipleOrdersBatchOrdersParameterInner
type PlaceMultipleOrdersBatchOrdersParameterInner struct {
	// Symbol
	Symbol       string                                                    `json:"symbol"`
	Side         PlaceMultipleOrdersBatchOrdersParameterInnerSide          `json:"side"`
	PositionSide *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide `json:"positionSide,omitempty"`
	Type         PlaceMultipleOrdersBatchOrdersParameterInnerType          `json:"type"`
	TimeInForce  *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce  `json:"timeInForce,omitempty"`
	// quantity measured by contract number
	Quantity   float32                                                 `json:"quantity"`
	ReduceOnly *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly `json:"reduceOnly,omitempty"`
	// Order price
	Price *float32 `json:"price,omitempty"`
	// A unique id among open orders. Automatically generated if not sent. Can only be string following the rule: `^[\\.A-Z\\:/a-z0-9_-]{1,36}$`
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	// Used with `STOP/STOP_MARKET` or `TAKE_PROFIT/TAKE_PROFIT_MARKET` orders.
	StopPrice *float32 `json:"stopPrice,omitempty"`
	// Used with `TRAILING_STOP_MARKET` orders, default as the latest price(supporting different `workingType`)
	ActivationPrice *float32 `json:"activationPrice,omitempty"`
	// Used with `TRAILING_STOP_MARKET` orders, min 0.1, max 4 where 1 for 1%
	CallbackRate            *float32                                                             `json:"callbackRate,omitempty"`
	WorkingType             *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType             `json:"workingType,omitempty"`
	PriceProtect            *PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect            `json:"priceProtect,omitempty"`
	NewOrderRespType        *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType        `json:"newOrderRespType,omitempty"`
	PriceMatch              *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch              `json:"priceMatch,omitempty"`
	SelfTradePreventionMode *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _PlaceMultipleOrdersBatchOrdersParameterInner PlaceMultipleOrdersBatchOrdersParameterInner

// NewPlaceMultipleOrdersBatchOrdersParameterInner instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersBatchOrdersParameterInner(symbol string, side PlaceMultipleOrdersBatchOrdersParameterInnerSide, type_ PlaceMultipleOrdersBatchOrdersParameterInnerType, quantity float32) *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	this.Symbol = symbol
	this.Side = side
	var positionSide = PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth
	this.PositionSide = &positionSide
	this.Type = type_
	var timeInForce = PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceGtc
	this.TimeInForce = &timeInForce
	this.Quantity = quantity
	var reduceOnly = PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyTrue
	this.ReduceOnly = &reduceOnly
	var workingType = PlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeMarkPrice
	this.WorkingType = &workingType
	var priceProtect = PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectTrue
	this.PriceProtect = &priceProtect
	var newOrderRespType = PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var priceMatch = PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent
	this.PriceMatch = &priceMatch
	var selfTradePreventionMode = PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeExpireTaker
	this.SelfTradePreventionMode = &selfTradePreventionMode
	return &this
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	var positionSide = PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth
	this.PositionSide = &positionSide
	var timeInForce = PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForceGtc
	this.TimeInForce = &timeInForce
	var reduceOnly = PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyTrue
	this.ReduceOnly = &reduceOnly
	var workingType = PlaceMultipleOrdersBatchOrdersParameterInnerWorkingTypeMarkPrice
	this.WorkingType = &workingType
	var priceProtect = PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtectTrue
	this.PriceProtect = &priceProtect
	var newOrderRespType = PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var priceMatch = PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent
	this.PriceMatch = &priceMatch
	var selfTradePreventionMode = PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeExpireTaker
	this.SelfTradePreventionMode = &selfTradePreventionMode
	return &this
}

// GetSymbol returns the Symbol field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetSide returns the Side field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() PlaceMultipleOrdersBatchOrdersParameterInnerSide {
	if o == nil {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerSide
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v PlaceMultipleOrdersBatchOrdersParameterInnerSide) {
	o.Side = v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSide() PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSideOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide and assigns it to the PositionSide field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPositionSide(v PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide) {
	o.PositionSide = &v
}

// GetType returns the Type field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetType() PlaceMultipleOrdersBatchOrdersParameterInnerType {
	if o == nil {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetType(v PlaceMultipleOrdersBatchOrdersParameterInnerType) {
	o.Type = v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForce() PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForceOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce and assigns it to the TimeInForce field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetTimeInForce(v PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce) {
	o.TimeInForce = &v
}

// GetQuantity returns the Quantity field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v float32) {
	o.Quantity = v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnly() PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnlyOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly and assigns it to the ReduceOnly field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetReduceOnly(v PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly) {
	o.ReduceOnly = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPrice() float32 {
	if o == nil || common.IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPrice(v float32) {
	o.Price = &v
}

// GetNewClientOrderId returns the NewClientOrderId field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderId() string {
	if o == nil || common.IsNil(o.NewClientOrderId) {
		var ret string
		return ret
	}
	return *o.NewClientOrderId
}

// GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.NewClientOrderId) {
		return nil, false
	}
	return o.NewClientOrderId, true
}

// HasNewClientOrderId returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewClientOrderId() bool {
	if o != nil && !common.IsNil(o.NewClientOrderId) {
		return true
	}

	return false
}

// SetNewClientOrderId gets a reference to the given string and assigns it to the NewClientOrderId field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewClientOrderId(v string) {
	o.NewClientOrderId = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPrice() float32 {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret float32
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given float32 and assigns it to the StopPrice field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v float32) {
	o.StopPrice = &v
}

// GetActivationPrice returns the ActivationPrice field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPrice() float32 {
	if o == nil || common.IsNil(o.ActivationPrice) {
		var ret float32
		return ret
	}
	return *o.ActivationPrice
}

// GetActivationPriceOk returns a tuple with the ActivationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPriceOk() (*float32, bool) {
	if o == nil || common.IsNil(o.ActivationPrice) {
		return nil, false
	}
	return o.ActivationPrice, true
}

// HasActivationPrice returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasActivationPrice() bool {
	if o != nil && !common.IsNil(o.ActivationPrice) {
		return true
	}

	return false
}

// SetActivationPrice gets a reference to the given float32 and assigns it to the ActivationPrice field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetActivationPrice(v float32) {
	o.ActivationPrice = &v
}

// GetCallbackRate returns the CallbackRate field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRate() float32 {
	if o == nil || common.IsNil(o.CallbackRate) {
		var ret float32
		return ret
	}
	return *o.CallbackRate
}

// GetCallbackRateOk returns a tuple with the CallbackRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRateOk() (*float32, bool) {
	if o == nil || common.IsNil(o.CallbackRate) {
		return nil, false
	}
	return o.CallbackRate, true
}

// HasCallbackRate returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasCallbackRate() bool {
	if o != nil && !common.IsNil(o.CallbackRate) {
		return true
	}

	return false
}

// SetCallbackRate gets a reference to the given float32 and assigns it to the CallbackRate field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetCallbackRate(v float32) {
	o.CallbackRate = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetWorkingType() PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetWorkingTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType and assigns it to the WorkingType field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetWorkingType(v PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType) {
	o.WorkingType = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtect() PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtectOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect and assigns it to the PriceProtect field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceProtect(v PlaceMultipleOrdersBatchOrdersParameterInnerPriceProtect) {
	o.PriceProtect = &v
}

// GetNewOrderRespType returns the NewOrderRespType field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespType() PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType {
	if o == nil || common.IsNil(o.NewOrderRespType) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType
		return ret
	}
	return *o.NewOrderRespType
}

// GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType, bool) {
	if o == nil || common.IsNil(o.NewOrderRespType) {
		return nil, false
	}
	return o.NewOrderRespType, true
}

// HasNewOrderRespType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewOrderRespType() bool {
	if o != nil && !common.IsNil(o.NewOrderRespType) {
		return true
	}

	return false
}

// SetNewOrderRespType gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType and assigns it to the NewOrderRespType field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewOrderRespType(v PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType) {
	o.NewOrderRespType = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatch() PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatchOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch and assigns it to the PriceMatch field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceMatch(v PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch) {
	o.PriceMatch = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionMode() PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionModeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode and assigns it to the SelfTradePreventionMode field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSelfTradePreventionMode(v PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode) {
	o.SelfTradePreventionMode = &v
}

func (o PlaceMultipleOrdersBatchOrdersParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceMultipleOrdersBatchOrdersParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["side"] = o.Side
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	toSerialize["type"] = o.Type
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	toSerialize["quantity"] = o.Quantity
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.NewClientOrderId) {
		toSerialize["newClientOrderId"] = o.NewClientOrderId
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.ActivationPrice) {
		toSerialize["activationPrice"] = o.ActivationPrice
	}
	if !common.IsNil(o.CallbackRate) {
		toSerialize["callbackRate"] = o.CallbackRate
	}
	if !common.IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	if !common.IsNil(o.PriceProtect) {
		toSerialize["priceProtect"] = o.PriceProtect
	}
	if !common.IsNil(o.NewOrderRespType) {
		toSerialize["newOrderRespType"] = o.NewOrderRespType
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceMultipleOrdersBatchOrdersParameterInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"side",
		"type",
		"quantity",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPlaceMultipleOrdersBatchOrdersParameterInner := _PlaceMultipleOrdersBatchOrdersParameterInner{}

	err = json.Unmarshal(data, &varPlaceMultipleOrdersBatchOrdersParameterInner)

	if err != nil {
		return err
	}

	*o = PlaceMultipleOrdersBatchOrdersParameterInner(varPlaceMultipleOrdersBatchOrdersParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "type")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "price")
		delete(additionalProperties, "newClientOrderId")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "activationPrice")
		delete(additionalProperties, "callbackRate")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "newOrderRespType")
		delete(additionalProperties, "priceMatch")
		delete(additionalProperties, "selfTradePreventionMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceMultipleOrdersBatchOrdersParameterInner struct {
	value *PlaceMultipleOrdersBatchOrdersParameterInner
	isSet bool
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInner) Get() *PlaceMultipleOrdersBatchOrdersParameterInner {
	return v.value
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInner) Set(val *PlaceMultipleOrdersBatchOrdersParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersBatchOrdersParameterInner(val *PlaceMultipleOrdersBatchOrdersParameterInner) *NullablePlaceMultipleOrdersBatchOrdersParameterInner {
	return &NullablePlaceMultipleOrdersBatchOrdersParameterInner{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersBatchOrdersParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersBatchOrdersParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
