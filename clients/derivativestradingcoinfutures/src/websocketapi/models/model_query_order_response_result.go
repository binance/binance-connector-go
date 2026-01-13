/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryOrderResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderResponseResult{}

// QueryOrderResponseResult struct for QueryOrderResponseResult
type QueryOrderResponseResult struct {
	OrderId                 *int64  `json:"orderId,omitempty"`
	Symbol                  *string `json:"symbol,omitempty"`
	Pair                    *string `json:"pair,omitempty"`
	Status                  *string `json:"status,omitempty"`
	ClientOrderId           *string `json:"clientOrderId,omitempty"`
	Price                   *string `json:"price,omitempty"`
	AvgPrice                *string `json:"avgPrice,omitempty"`
	OrigQty                 *string `json:"origQty,omitempty"`
	ExecutedQty             *string `json:"executedQty,omitempty"`
	CumBase                 *string `json:"cumBase,omitempty"`
	TimeInForce             *string `json:"timeInForce,omitempty"`
	Type                    *string `json:"type,omitempty"`
	ReduceOnly              *bool   `json:"reduceOnly,omitempty"`
	ClosePosition           *bool   `json:"closePosition,omitempty"`
	Side                    *string `json:"side,omitempty"`
	PositionSide            *string `json:"positionSide,omitempty"`
	StopPrice               *string `json:"stopPrice,omitempty"`
	WorkingType             *string `json:"workingType,omitempty"`
	PriceProtect            *bool   `json:"priceProtect,omitempty"`
	OrigType                *string `json:"origType,omitempty"`
	SelfTradePreventionMode *string `json:"selfTradePreventionMode,omitempty"`
	Time                    *int64  `json:"time,omitempty"`
	UpdateTime              *int64  `json:"updateTime,omitempty"`
	PriceMatch              *string `json:"priceMatch,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _QueryOrderResponseResult QueryOrderResponseResult

// NewQueryOrderResponseResult instantiates a new QueryOrderResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderResponseResult() *QueryOrderResponseResult {
	this := QueryOrderResponseResult{}
	return &this
}

// NewQueryOrderResponseResultWithDefaults instantiates a new QueryOrderResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderResponseResultWithDefaults() *QueryOrderResponseResult {
	this := QueryOrderResponseResult{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *QueryOrderResponseResult) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *QueryOrderResponseResult) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPair returns the Pair field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetPair() string {
	if o == nil || common.IsNil(o.Pair) {
		var ret string
		return ret
	}
	return *o.Pair
}

// GetPairOk returns a tuple with the Pair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetPairOk() (*string, bool) {
	if o == nil || common.IsNil(o.Pair) {
		return nil, false
	}
	return o.Pair, true
}

// HasPair returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasPair() bool {
	if o != nil && !common.IsNil(o.Pair) {
		return true
	}

	return false
}

// SetPair gets a reference to the given string and assigns it to the Pair field.
func (o *QueryOrderResponseResult) SetPair(v string) {
	o.Pair = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *QueryOrderResponseResult) SetStatus(v string) {
	o.Status = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *QueryOrderResponseResult) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *QueryOrderResponseResult) SetPrice(v string) {
	o.Price = &v
}

// GetAvgPrice returns the AvgPrice field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetAvgPrice() string {
	if o == nil || common.IsNil(o.AvgPrice) {
		var ret string
		return ret
	}
	return *o.AvgPrice
}

// GetAvgPriceOk returns a tuple with the AvgPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetAvgPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.AvgPrice) {
		return nil, false
	}
	return o.AvgPrice, true
}

// HasAvgPrice returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasAvgPrice() bool {
	if o != nil && !common.IsNil(o.AvgPrice) {
		return true
	}

	return false
}

// SetAvgPrice gets a reference to the given string and assigns it to the AvgPrice field.
func (o *QueryOrderResponseResult) SetAvgPrice(v string) {
	o.AvgPrice = &v
}

// GetOrigQty returns the OrigQty field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetOrigQty() string {
	if o == nil || common.IsNil(o.OrigQty) {
		var ret string
		return ret
	}
	return *o.OrigQty
}

// GetOrigQtyOk returns a tuple with the OrigQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetOrigQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigQty) {
		return nil, false
	}
	return o.OrigQty, true
}

// HasOrigQty returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasOrigQty() bool {
	if o != nil && !common.IsNil(o.OrigQty) {
		return true
	}

	return false
}

// SetOrigQty gets a reference to the given string and assigns it to the OrigQty field.
func (o *QueryOrderResponseResult) SetOrigQty(v string) {
	o.OrigQty = &v
}

// GetExecutedQty returns the ExecutedQty field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetExecutedQty() string {
	if o == nil || common.IsNil(o.ExecutedQty) {
		var ret string
		return ret
	}
	return *o.ExecutedQty
}

// GetExecutedQtyOk returns a tuple with the ExecutedQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetExecutedQtyOk() (*string, bool) {
	if o == nil || common.IsNil(o.ExecutedQty) {
		return nil, false
	}
	return o.ExecutedQty, true
}

// HasExecutedQty returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasExecutedQty() bool {
	if o != nil && !common.IsNil(o.ExecutedQty) {
		return true
	}

	return false
}

// SetExecutedQty gets a reference to the given string and assigns it to the ExecutedQty field.
func (o *QueryOrderResponseResult) SetExecutedQty(v string) {
	o.ExecutedQty = &v
}

// GetCumBase returns the CumBase field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetCumBase() string {
	if o == nil || common.IsNil(o.CumBase) {
		var ret string
		return ret
	}
	return *o.CumBase
}

// GetCumBaseOk returns a tuple with the CumBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetCumBaseOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumBase) {
		return nil, false
	}
	return o.CumBase, true
}

// HasCumBase returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasCumBase() bool {
	if o != nil && !common.IsNil(o.CumBase) {
		return true
	}

	return false
}

// SetCumBase gets a reference to the given string and assigns it to the CumBase field.
func (o *QueryOrderResponseResult) SetCumBase(v string) {
	o.CumBase = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *QueryOrderResponseResult) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *QueryOrderResponseResult) SetType(v string) {
	o.Type = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *QueryOrderResponseResult) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetClosePosition() bool {
	if o == nil || common.IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetClosePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasClosePosition() bool {
	if o != nil && !common.IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *QueryOrderResponseResult) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *QueryOrderResponseResult) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *QueryOrderResponseResult) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetStopPrice returns the StopPrice field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetStopPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.StopPrice) {
		return nil, false
	}
	return o.StopPrice, true
}

// HasStopPrice returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasStopPrice() bool {
	if o != nil && !common.IsNil(o.StopPrice) {
		return true
	}

	return false
}

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *QueryOrderResponseResult) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *QueryOrderResponseResult) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *QueryOrderResponseResult) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetOrigType returns the OrigType field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetOrigType() string {
	if o == nil || common.IsNil(o.OrigType) {
		var ret string
		return ret
	}
	return *o.OrigType
}

// GetOrigTypeOk returns a tuple with the OrigType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetOrigTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrigType) {
		return nil, false
	}
	return o.OrigType, true
}

// HasOrigType returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasOrigType() bool {
	if o != nil && !common.IsNil(o.OrigType) {
		return true
	}

	return false
}

// SetOrigType gets a reference to the given string and assigns it to the OrigType field.
func (o *QueryOrderResponseResult) SetOrigType(v string) {
	o.OrigType = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *QueryOrderResponseResult) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *QueryOrderResponseResult) SetTime(v int64) {
	o.Time = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *QueryOrderResponseResult) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *QueryOrderResponseResult) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponseResult) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *QueryOrderResponseResult) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *QueryOrderResponseResult) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

func (o QueryOrderResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Pair) {
		toSerialize["pair"] = o.Pair
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.AvgPrice) {
		toSerialize["avgPrice"] = o.AvgPrice
	}
	if !common.IsNil(o.OrigQty) {
		toSerialize["origQty"] = o.OrigQty
	}
	if !common.IsNil(o.ExecutedQty) {
		toSerialize["executedQty"] = o.ExecutedQty
	}
	if !common.IsNil(o.CumBase) {
		toSerialize["cumBase"] = o.CumBase
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.ClosePosition) {
		toSerialize["closePosition"] = o.ClosePosition
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
	if !common.IsNil(o.WorkingType) {
		toSerialize["workingType"] = o.WorkingType
	}
	if !common.IsNil(o.PriceProtect) {
		toSerialize["priceProtect"] = o.PriceProtect
	}
	if !common.IsNil(o.OrigType) {
		toSerialize["origType"] = o.OrigType
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryOrderResponseResult) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderResponseResult := _QueryOrderResponseResult{}

	err = json.Unmarshal(data, &varQueryOrderResponseResult)

	if err != nil {
		return err
	}

	*o = QueryOrderResponseResult(varQueryOrderResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "pair")
		delete(additionalProperties, "status")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "price")
		delete(additionalProperties, "avgPrice")
		delete(additionalProperties, "origQty")
		delete(additionalProperties, "executedQty")
		delete(additionalProperties, "cumBase")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "type")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "closePosition")
		delete(additionalProperties, "side")
		delete(additionalProperties, "positionSide")
		delete(additionalProperties, "stopPrice")
		delete(additionalProperties, "workingType")
		delete(additionalProperties, "priceProtect")
		delete(additionalProperties, "origType")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "time")
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "priceMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderResponseResult struct {
	value *QueryOrderResponseResult
	isSet bool
}

func (v NullableQueryOrderResponseResult) Get() *QueryOrderResponseResult {
	return v.value
}

func (v *NullableQueryOrderResponseResult) Set(val *QueryOrderResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderResponseResult(val *QueryOrderResponseResult) *NullableQueryOrderResponseResult {
	return &NullableQueryOrderResponseResult{value: val, isSet: true}
}

func (v NullableQueryOrderResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
