/*
Binance Derivatives Trading COIN Futures REST API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PlaceMultipleOrdersBatchOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersBatchOrdersParameterInner{}

// PlaceMultipleOrdersBatchOrdersParameterInner struct for PlaceMultipleOrdersBatchOrdersParameterInner
type PlaceMultipleOrdersBatchOrdersParameterInner struct {
	Symbol                  *string                                                              `json:"symbol,omitempty"`
	Side                    *ModifyMultipleOrdersBatchOrdersParameterInnerSide                   `json:"side,omitempty"`
	PositionSide            *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide            `json:"positionSide,omitempty"`
	Type                    *PlaceMultipleOrdersBatchOrdersParameterInnerType                    `json:"type,omitempty"`
	TimeInForce             *PlaceMultipleOrdersBatchOrdersParameterInnerTimeInForce             `json:"timeInForce,omitempty"`
	Quantity                *string                                                              `json:"quantity,omitempty"`
	ReduceOnly              *string                                                              `json:"reduceOnly,omitempty"`
	Price                   *string                                                              `json:"price,omitempty"`
	NewClientOrderId        *string                                                              `json:"newClientOrderId,omitempty"`
	StopPrice               *string                                                              `json:"stopPrice,omitempty"`
	ActivationPrice         *string                                                              `json:"activationPrice,omitempty"`
	CallbackRate            *string                                                              `json:"callbackRate,omitempty"`
	WorkingType             *PlaceMultipleOrdersBatchOrdersParameterInnerWorkingType             `json:"workingType,omitempty"`
	PriceProtect            *string                                                              `json:"priceProtect,omitempty"`
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
func NewPlaceMultipleOrdersBatchOrdersParameterInner() *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	return &this
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() ModifyMultipleOrdersBatchOrdersParameterInnerSide {
	if o == nil || common.IsNil(o.Side) {
		var ret ModifyMultipleOrdersBatchOrdersParameterInnerSide
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*ModifyMultipleOrdersBatchOrdersParameterInnerSide, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given ModifyMultipleOrdersBatchOrdersParameterInnerSide and assigns it to the Side field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v ModifyMultipleOrdersBatchOrdersParameterInnerSide) {
	o.Side = &v
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

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetType() PlaceMultipleOrdersBatchOrdersParameterInnerType {
	if o == nil || common.IsNil(o.Type) {
		var ret PlaceMultipleOrdersBatchOrdersParameterInnerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersBatchOrdersParameterInnerType, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PlaceMultipleOrdersBatchOrdersParameterInnerType and assigns it to the Type field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetType(v PlaceMultipleOrdersBatchOrdersParameterInnerType) {
	o.Type = &v
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

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnly() string {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret string
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnlyOk() (*string, bool) {
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

// SetReduceOnly gets a reference to the given string and assigns it to the ReduceOnly field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetReduceOnly(v string) {
	o.ReduceOnly = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*string, bool) {
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

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPrice(v string) {
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
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPrice() string {
	if o == nil || common.IsNil(o.StopPrice) {
		var ret string
		return ret
	}
	return *o.StopPrice
}

// GetStopPriceOk returns a tuple with the StopPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetStopPriceOk() (*string, bool) {
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

// SetStopPrice gets a reference to the given string and assigns it to the StopPrice field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetStopPrice(v string) {
	o.StopPrice = &v
}

// GetActivationPrice returns the ActivationPrice field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPrice() string {
	if o == nil || common.IsNil(o.ActivationPrice) {
		var ret string
		return ret
	}
	return *o.ActivationPrice
}

// GetActivationPriceOk returns a tuple with the ActivationPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetActivationPriceOk() (*string, bool) {
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

// SetActivationPrice gets a reference to the given string and assigns it to the ActivationPrice field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetActivationPrice(v string) {
	o.ActivationPrice = &v
}

// GetCallbackRate returns the CallbackRate field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRate() string {
	if o == nil || common.IsNil(o.CallbackRate) {
		var ret string
		return ret
	}
	return *o.CallbackRate
}

// GetCallbackRateOk returns a tuple with the CallbackRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetCallbackRateOk() (*string, bool) {
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

// SetCallbackRate gets a reference to the given string and assigns it to the CallbackRate field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetCallbackRate(v string) {
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
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtect() string {
	if o == nil || common.IsNil(o.PriceProtect) {
		var ret string
		return ret
	}
	return *o.PriceProtect
}

// GetPriceProtectOk returns a tuple with the PriceProtect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceProtectOk() (*string, bool) {
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

// SetPriceProtect gets a reference to the given string and assigns it to the PriceProtect field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceProtect(v string) {
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
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
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
