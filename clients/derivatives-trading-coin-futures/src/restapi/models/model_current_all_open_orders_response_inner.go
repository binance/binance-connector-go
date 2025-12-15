/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CurrentAllOpenOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentAllOpenOrdersResponseInner{}

// CurrentAllOpenOrdersResponseInner struct for CurrentAllOpenOrdersResponseInner
type CurrentAllOpenOrdersResponseInner struct {
	AvgPrice                *string `json:"avgPrice,omitempty"`
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	CumBase                 *string `json:"cumBase,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	OrderId                 *int64  `json:"orderId,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	OrigType                *string `json:"origType,omitempty"`
	Price                   *string `json:"price,omitempty"`
	ReduceOnly              *bool   `json:"reduceOnly,omitempty"`
	Side                    *string `json:"side,omitempty"`
	PositionSide            *string `json:"positionSide,omitempty"`
	Status                  *string `json:"status,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	ClosePosition           *bool   `json:"closePosition,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	Time                    *int64  `json:"time,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	ActivatePrice           *string `json:"activatePrice,omitempty"`
	PriceRate               *string `json:"priceRate,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	WorkingType             *string `json:"workingType,omitempty"`
	PriceProtect            *bool   `json:"priceProtect,omitempty"`
	PriceMatch              *string `json:"priceMatch,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _CurrentAllOpenOrdersResponseInner CurrentAllOpenOrdersResponseInner

// NewCurrentAllOpenOrdersResponseInner instantiates a new CurrentAllOpenOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentAllOpenOrdersResponseInner() *CurrentAllOpenOrdersResponseInner {
	this := CurrentAllOpenOrdersResponseInner{}
	return &this
}

// NewCurrentAllOpenOrdersResponseInnerWithDefaults instantiates a new CurrentAllOpenOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentAllOpenOrdersResponseInnerWithDefaults() *CurrentAllOpenOrdersResponseInner {
	this := CurrentAllOpenOrdersResponseInner{}
	return &this
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *CurrentAllOpenOrdersResponseInner) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *CurrentAllOpenOrdersResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetCumBase returns the CumBase field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetCumBase() string {
	if o == nil || common.IsNil(o.CumBase) {
		var ret string
		return ret
	}
	return *o.CumBase
}

// GetCumBaseOk returns a tuple with the CumBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetCumBaseOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumBase) {
		return nil, false
	}
	return o.CumBase, true
}

// HasCumBase returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasCumBase() bool {
	if o != nil && !common.IsNil(o.CumBase) {
		return true
	}

	return false
}

// SetCumBase gets a reference to the given string and assigns it to the CumBase field.
func (o *CurrentAllOpenOrdersResponseInner) SetCumBase(v string) {
	o.CumBase = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *CurrentAllOpenOrdersResponseInner) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *CurrentAllOpenOrdersResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *CurrentAllOpenOrdersResponseInner) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetOrigType() string {
	if o == nil || common.IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetOrigTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasOrigType() bool {
	if o != nil && !common.IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *CurrentAllOpenOrdersResponseInner) SetOrigType(v string) {
	o.OrigType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CurrentAllOpenOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *CurrentAllOpenOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CurrentAllOpenOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *CurrentAllOpenOrdersResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CurrentAllOpenOrdersResponseInner) SetStatus(v string) {
	o.Status = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *CurrentAllOpenOrdersResponseInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetClosePosition() bool {
	if o == nil || common.IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetClosePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasClosePosition() bool {
	if o != nil && !common.IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *CurrentAllOpenOrdersResponseInner) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CurrentAllOpenOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *CurrentAllOpenOrdersResponseInner) SetTime(v int64) {
	o.Time = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CurrentAllOpenOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CurrentAllOpenOrdersResponseInner) SetType(v string) {
	o.Type = &v
}

// GetActivatePrice returns the ActivatePrice field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetActivatePrice() string {
	if o == nil || common.IsNil(o.ActivatePrice) {
		var ret string
		return ret
	}
	return *o.ActivatePrice
}

// GetActivatePriceOk returns a tuple with the ActivatePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetActivatePriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActivatePrice) {
		return nil, false
	}
	return o.ActivatePrice, true
}

// HasActivatePrice returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasActivatePrice() bool {
	if o != nil && !common.IsNil(o.ActivatePrice) {
		return true
	}

	return false
}

// SetActivatePrice gets a reference to the given string and assigns it to the ActivatePrice field.
func (o *CurrentAllOpenOrdersResponseInner) SetActivatePrice(v string) {
	o.ActivatePrice = &v
}

// GetPriceRate returns the PriceRate field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceRate() string {
	if o == nil || common.IsNil(o.PriceRate) {
		var ret string
		return ret
	}
	return *o.PriceRate
}

// GetPriceRateOk returns a tuple with the PriceRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceRate) {
		return nil, false
	}
	return o.PriceRate, true
}

// HasPriceRate returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasPriceRate() bool {
	if o != nil && !common.IsNil(o.PriceRate) {
		return true
	}

	return false
}

// SetPriceRate gets a reference to the given string and assigns it to the PriceRate field.
func (o *CurrentAllOpenOrdersResponseInner) SetPriceRate(v string) {
	o.PriceRate = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *CurrentAllOpenOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *CurrentAllOpenOrdersResponseInner) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *CurrentAllOpenOrdersResponseInner) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *CurrentAllOpenOrdersResponseInner) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *CurrentAllOpenOrdersResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllOpenOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *CurrentAllOpenOrdersResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *CurrentAllOpenOrdersResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

func (o CurrentAllOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentAllOpenOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.CumBase) {
		toSerialize["cumBase"] = o.CumBase
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.OrigType) {
		toSerialize["origType"] = o.OrigType
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
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.StopPrice) {
		toSerialize["stopPrice"] = o.StopPrice
	}
	if !common.IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
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
	if !common.IsNil(o.ActivatePrice) {
		toSerialize["activatePrice"] = o.ActivatePrice
	}
	if !common.IsNil(o.PriceRate) {
		toSerialize["priceRate"] = o.PriceRate
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

func (o *CurrentAllOpenOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCurrentAllOpenOrdersResponseInner := _CurrentAllOpenOrdersResponseInner{}

	err = json.Unmarshal(data, &varCurrentAllOpenOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = CurrentAllOpenOrdersResponseInner(varCurrentAllOpenOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "cumBase")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "origType")
		delete(additionalProperties, "price")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "status")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "closePosition")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "time")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "activatePrice")
		delete(additionalProperties, "priceRate")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "priceMatch")
		delete(additionalProperties, "selfTradePreventionMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCurrentAllOpenOrdersResponseInner struct {
	value *CurrentAllOpenOrdersResponseInner
	isSet bool
}

func (v NullableCurrentAllOpenOrdersResponseInner) Get() *CurrentAllOpenOrdersResponseInner {
	return v.value
}

func (v *NullableCurrentAllOpenOrdersResponseInner) Set(val *CurrentAllOpenOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentAllOpenOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentAllOpenOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentAllOpenOrdersResponseInner(val *CurrentAllOpenOrdersResponseInner) *NullableCurrentAllOpenOrdersResponseInner {
	return &NullableCurrentAllOpenOrdersResponseInner{value: val, isSet: true}
}

func (v NullableCurrentAllOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentAllOpenOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
