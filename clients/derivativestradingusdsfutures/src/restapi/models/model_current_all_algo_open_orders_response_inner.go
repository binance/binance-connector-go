/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CurrentAllAlgoOpenOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CurrentAllAlgoOpenOrdersResponseInner{}

// CurrentAllAlgoOpenOrdersResponseInner struct for CurrentAllAlgoOpenOrdersResponseInner
type CurrentAllAlgoOpenOrdersResponseInner struct {
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
	ActualOrderId           *string `json:"actualOrderId,omitempty"`
	ActualPrice             *string `json:"actualPrice,omitempty"`
	TriggerPrice            *string `json:"triggerPrice,omitempty"`
	Price                   *string `json:"price,omitempty"`
	IcebergQuantity         *string `json:"icebergQuantity,omitempty"`
	TpTriggerPrice          *string `json:"tpTriggerPrice,omitempty"`
	TpPrice                 *string `json:"tpPrice,omitempty"`
	SlTriggerPrice          *string `json:"slTriggerPrice,omitempty"`
	SlPrice                 *string `json:"slPrice,omitempty"`
	TpOrderType             *string `json:"tpOrderType,omitempty"`
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

type _CurrentAllAlgoOpenOrdersResponseInner CurrentAllAlgoOpenOrdersResponseInner

// NewCurrentAllAlgoOpenOrdersResponseInner instantiates a new CurrentAllAlgoOpenOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrentAllAlgoOpenOrdersResponseInner() *CurrentAllAlgoOpenOrdersResponseInner {
	this := CurrentAllAlgoOpenOrdersResponseInner{}
	return &this
}

// NewCurrentAllAlgoOpenOrdersResponseInnerWithDefaults instantiates a new CurrentAllAlgoOpenOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrentAllAlgoOpenOrdersResponseInnerWithDefaults() *CurrentAllAlgoOpenOrdersResponseInner {
	this := CurrentAllAlgoOpenOrdersResponseInner{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetAlgoType returns the AlgoType field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoType() string {
	if o == nil || common.IsNil(o.AlgoType) {
		var ret string
		return ret
	}
	return *o.AlgoType
}

// GetAlgoTypeOk returns a tuple with the AlgoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoType) {
		return nil, false
	}
	return o.AlgoType, true
}

// HasAlgoType returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasAlgoType() bool {
	if o != nil && !common.IsNil(o.AlgoType) {
		return true
	}

	return false
}

// SetAlgoType gets a reference to the given string and assigns it to the AlgoType field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetAlgoType(v string) {
	o.AlgoType = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTimeInForce() string {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret string
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTimeInForceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given string and assigns it to the TimeInForce field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTimeInForce(v string) {
	o.TimeInForce = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetAlgoStatus returns the AlgoStatus field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoStatus() string {
	if o == nil || common.IsNil(o.AlgoStatus) {
		var ret string
		return ret
	}
	return *o.AlgoStatus
}

// GetAlgoStatusOk returns a tuple with the AlgoStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetAlgoStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlgoStatus) {
		return nil, false
	}
	return o.AlgoStatus, true
}

// HasAlgoStatus returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasAlgoStatus() bool {
	if o != nil && !common.IsNil(o.AlgoStatus) {
		return true
	}

	return false
}

// SetAlgoStatus gets a reference to the given string and assigns it to the AlgoStatus field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetAlgoStatus(v string) {
	o.AlgoStatus = &v
}

// GetActualOrderId returns the ActualOrderId field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetActualOrderId() string {
	if o == nil || common.IsNil(o.ActualOrderId) {
		var ret string
		return ret
	}
	return *o.ActualOrderId
}

// GetActualOrderIdOk returns a tuple with the ActualOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetActualOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualOrderId) {
		return nil, false
	}
	return o.ActualOrderId, true
}

// HasActualOrderId returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasActualOrderId() bool {
	if o != nil && !common.IsNil(o.ActualOrderId) {
		return true
	}

	return false
}

// SetActualOrderId gets a reference to the given string and assigns it to the ActualOrderId field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetActualOrderId(v string) {
	o.ActualOrderId = &v
}

// GetActualPrice returns the ActualPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetActualPrice() string {
	if o == nil || common.IsNil(o.ActualPrice) {
		var ret string
		return ret
	}
	return *o.ActualPrice
}

// GetActualPriceOk returns a tuple with the ActualPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetActualPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.ActualPrice) {
		return nil, false
	}
	return o.ActualPrice, true
}

// HasActualPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasActualPrice() bool {
	if o != nil && !common.IsNil(o.ActualPrice) {
		return true
	}

	return false
}

// SetActualPrice gets a reference to the given string and assigns it to the ActualPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetActualPrice(v string) {
	o.ActualPrice = &v
}

// GetTriggerPrice returns the TriggerPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTriggerPrice() string {
	if o == nil || common.IsNil(o.TriggerPrice) {
		var ret string
		return ret
	}
	return *o.TriggerPrice
}

// GetTriggerPriceOk returns a tuple with the TriggerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTriggerPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TriggerPrice) {
		return nil, false
	}
	return o.TriggerPrice, true
}

// HasTriggerPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTriggerPrice() bool {
	if o != nil && !common.IsNil(o.TriggerPrice) {
		return true
	}

	return false
}

// SetTriggerPrice gets a reference to the given string and assigns it to the TriggerPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTriggerPrice(v string) {
	o.TriggerPrice = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetIcebergQuantity returns the IcebergQuantity field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetIcebergQuantity() string {
	if o == nil || common.IsNil(o.IcebergQuantity) {
		var ret string
		return ret
	}
	return *o.IcebergQuantity
}

// GetIcebergQuantityOk returns a tuple with the IcebergQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetIcebergQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.IcebergQuantity) {
		return nil, false
	}
	return o.IcebergQuantity, true
}

// HasIcebergQuantity returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasIcebergQuantity() bool {
	if o != nil && !common.IsNil(o.IcebergQuantity) {
		return true
	}

	return false
}

// SetIcebergQuantity gets a reference to the given string and assigns it to the IcebergQuantity field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetIcebergQuantity(v string) {
	o.IcebergQuantity = &v
}

// GetTpTriggerPrice returns the TpTriggerPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpTriggerPrice() string {
	if o == nil || common.IsNil(o.TpTriggerPrice) {
		var ret string
		return ret
	}
	return *o.TpTriggerPrice
}

// GetTpTriggerPriceOk returns a tuple with the TpTriggerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpTriggerPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TpTriggerPrice) {
		return nil, false
	}
	return o.TpTriggerPrice, true
}

// HasTpTriggerPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTpTriggerPrice() bool {
	if o != nil && !common.IsNil(o.TpTriggerPrice) {
		return true
	}

	return false
}

// SetTpTriggerPrice gets a reference to the given string and assigns it to the TpTriggerPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTpTriggerPrice(v string) {
	o.TpTriggerPrice = &v
}

// GetTpPrice returns the TpPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpPrice() string {
	if o == nil || common.IsNil(o.TpPrice) {
		var ret string
		return ret
	}
	return *o.TpPrice
}

// GetTpPriceOk returns a tuple with the TpPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TpPrice) {
		return nil, false
	}
	return o.TpPrice, true
}

// HasTpPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTpPrice() bool {
	if o != nil && !common.IsNil(o.TpPrice) {
		return true
	}

	return false
}

// SetTpPrice gets a reference to the given string and assigns it to the TpPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTpPrice(v string) {
	o.TpPrice = &v
}

// GetSlTriggerPrice returns the SlTriggerPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSlTriggerPrice() string {
	if o == nil || common.IsNil(o.SlTriggerPrice) {
		var ret string
		return ret
	}
	return *o.SlTriggerPrice
}

// GetSlTriggerPriceOk returns a tuple with the SlTriggerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSlTriggerPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SlTriggerPrice) {
		return nil, false
	}
	return o.SlTriggerPrice, true
}

// HasSlTriggerPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasSlTriggerPrice() bool {
	if o != nil && !common.IsNil(o.SlTriggerPrice) {
		return true
	}

	return false
}

// SetSlTriggerPrice gets a reference to the given string and assigns it to the SlTriggerPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetSlTriggerPrice(v string) {
	o.SlTriggerPrice = &v
}

// GetSlPrice returns the SlPrice field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSlPrice() string {
	if o == nil || common.IsNil(o.SlPrice) {
		var ret string
		return ret
	}
	return *o.SlPrice
}

// GetSlPriceOk returns a tuple with the SlPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSlPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.SlPrice) {
		return nil, false
	}
	return o.SlPrice, true
}

// HasSlPrice returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasSlPrice() bool {
	if o != nil && !common.IsNil(o.SlPrice) {
		return true
	}

	return false
}

// SetSlPrice gets a reference to the given string and assigns it to the SlPrice field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetSlPrice(v string) {
	o.SlPrice = &v
}

// GetTpOrderType returns the TpOrderType field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpOrderType() string {
	if o == nil || common.IsNil(o.TpOrderType) {
		var ret string
		return ret
	}
	return *o.TpOrderType
}

// GetTpOrderTypeOk returns a tuple with the TpOrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTpOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.TpOrderType) {
		return nil, false
	}
	return o.TpOrderType, true
}

// HasTpOrderType returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTpOrderType() bool {
	if o != nil && !common.IsNil(o.TpOrderType) {
		return true
	}

	return false
}

// SetTpOrderType gets a reference to the given string and assigns it to the TpOrderType field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTpOrderType(v string) {
	o.TpOrderType = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSelfTradePreventionMode() string {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret string
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given string and assigns it to the SelfTradePreventionMode field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetSelfTradePreventionMode(v string) {
	o.SelfTradePreventionMode = &v
}

// GetWorkingType returns the WorkingType field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetWorkingType() string {
	if o == nil || common.IsNil(o.WorkingType) {
		var ret string
		return ret
	}
	return *o.WorkingType
}

// GetWorkingTypeOk returns a tuple with the WorkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetWorkingTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkingType) {
		return nil, false
	}
	return o.WorkingType, true
}

// HasWorkingType returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasWorkingType() bool {
	if o != nil && !common.IsNil(o.WorkingType) {
		return true
	}

	return false
}

// SetWorkingType gets a reference to the given string and assigns it to the WorkingType field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetWorkingType(v string) {
	o.WorkingType = &v
}

// GetPriceMatch returns the PriceMatch field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPriceMatch() string {
	if o == nil || common.IsNil(o.PriceMatch) {
		var ret string
		return ret
	}
	return *o.PriceMatch
}

// GetPriceMatchOk returns a tuple with the PriceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPriceMatchOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceMatch) {
		return nil, false
	}
	return o.PriceMatch, true
}

// HasPriceMatch returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasPriceMatch() bool {
	if o != nil && !common.IsNil(o.PriceMatch) {
		return true
	}

	return false
}

// SetPriceMatch gets a reference to the given string and assigns it to the PriceMatch field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetPriceMatch(v string) {
	o.PriceMatch = &v
}

// GetClosePosition returns the ClosePosition field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetClosePosition() bool {
	if o == nil || common.IsNil(o.ClosePosition) {
		var ret bool
		return ret
	}
	return *o.ClosePosition
}

// GetClosePositionOk returns a tuple with the ClosePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetClosePositionOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ClosePosition) {
		return nil, false
	}
	return o.ClosePosition, true
}

// HasClosePosition returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasClosePosition() bool {
	if o != nil && !common.IsNil(o.ClosePosition) {
		return true
	}

	return false
}

// SetClosePosition gets a reference to the given bool and assigns it to the ClosePosition field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetClosePosition(v bool) {
	o.ClosePosition = &v
}

// GetPriceProtect returns the PriceProtect field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPriceProtect() bool {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret bool
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetPriceProtectOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PriceProtect) {
		return nil, false
	}
	return o.PriceProtect, true
}

// HasPriceProtect returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasPriceProtect() bool {
	if o != nil && !common.IsNil(o.PriceProtect) {
		return true
	}

	return false
}

// SetPriceProtect gets a reference to the given bool and assigns it to the PriceProtect field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetPriceProtect(v bool) {
	o.PriceProtect = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetCreateTime() int64 {
	if o == nil || common.IsNil(o.CreateTime) {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetCreateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasCreateTime() bool {
	if o != nil && !common.IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetTriggerTime returns the TriggerTime field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTriggerTime() int64 {
	if o == nil || common.IsNil(o.TriggerTime) {
		var ret int64
		return ret
	}
	return *o.TriggerTime
}

// GetTriggerTimeOk returns a tuple with the TriggerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetTriggerTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TriggerTime) {
		return nil, false
	}
	return o.TriggerTime, true
}

// HasTriggerTime returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasTriggerTime() bool {
	if o != nil && !common.IsNil(o.TriggerTime) {
		return true
	}

	return false
}

// SetTriggerTime gets a reference to the given int64 and assigns it to the TriggerTime field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetTriggerTime(v int64) {
	o.TriggerTime = &v
}

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *CurrentAllAlgoOpenOrdersResponseInner) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *CurrentAllAlgoOpenOrdersResponseInner) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
}

func (o CurrentAllAlgoOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrentAllAlgoOpenOrdersResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.ActualOrderId) {
		toSerialize["actualOrderId"] = o.ActualOrderId
	}
	if !common.IsNil(o.ActualPrice) {
		toSerialize["actualPrice"] = o.ActualPrice
	}
	if !common.IsNil(o.TriggerPrice) {
		toSerialize["triggerPrice"] = o.TriggerPrice
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.IcebergQuantity) {
		toSerialize["icebergQuantity"] = o.IcebergQuantity
	}
	if !common.IsNil(o.TpTriggerPrice) {
		toSerialize["tpTriggerPrice"] = o.TpTriggerPrice
	}
	if !common.IsNil(o.TpPrice) {
		toSerialize["tpPrice"] = o.TpPrice
	}
	if !common.IsNil(o.SlTriggerPrice) {
		toSerialize["slTriggerPrice"] = o.SlTriggerPrice
	}
	if !common.IsNil(o.SlPrice) {
		toSerialize["slPrice"] = o.SlPrice
	}
	if !common.IsNil(o.TpOrderType) {
		toSerialize["tpOrderType"] = o.TpOrderType
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

func (o *CurrentAllAlgoOpenOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varCurrentAllAlgoOpenOrdersResponseInner := _CurrentAllAlgoOpenOrdersResponseInner{}

	err = json.Unmarshal(data, &varCurrentAllAlgoOpenOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = CurrentAllAlgoOpenOrdersResponseInner(varCurrentAllAlgoOpenOrdersResponseInner)

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
		delete(additionalProperties, "actualOrderId")
		delete(additionalProperties, "actualPrice")
		delete(additionalProperties, "triggerPrice")
		delete(additionalProperties, "price")
		delete(additionalProperties, "icebergQuantity")
		delete(additionalProperties, "tpTriggerPrice")
		delete(additionalProperties, "tpPrice")
		delete(additionalProperties, "slTriggerPrice")
		delete(additionalProperties, "slPrice")
		delete(additionalProperties, "tpOrderType")
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

type NullableCurrentAllAlgoOpenOrdersResponseInner struct {
	value *CurrentAllAlgoOpenOrdersResponseInner
	isSet bool
}

func (v NullableCurrentAllAlgoOpenOrdersResponseInner) Get() *CurrentAllAlgoOpenOrdersResponseInner {
	return v.value
}

func (v *NullableCurrentAllAlgoOpenOrdersResponseInner) Set(val *CurrentAllAlgoOpenOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrentAllAlgoOpenOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrentAllAlgoOpenOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrentAllAlgoOpenOrdersResponseInner(val *CurrentAllAlgoOpenOrdersResponseInner) *NullableCurrentAllAlgoOpenOrdersResponseInner {
	return &NullableCurrentAllAlgoOpenOrdersResponseInner{value: val, isSet: true}
}

func (v NullableCurrentAllAlgoOpenOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrentAllAlgoOpenOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
