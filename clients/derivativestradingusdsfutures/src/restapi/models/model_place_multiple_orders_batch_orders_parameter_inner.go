/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PlaceMultipleOrdersBatchOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersBatchOrdersParameterInner{}

// PlaceMultipleOrdersBatchOrdersParameterInner struct for PlaceMultipleOrdersBatchOrdersParameterInner
type PlaceMultipleOrdersBatchOrdersParameterInner struct {
	Symbol                  *string                                                              `json:"symbol,omitempty"`
	Side                    *NewAlgoOrderSideParameter                                           `json:"side,omitempty"`
	PositionSide            *PlaceMultipleOrdersBatchOrdersParameterInnerPositionSide            `json:"positionSide,omitempty"`
	Type                    *PlaceMultipleOrdersBatchOrdersParameterInnerType                    `json:"type,omitempty"`
	TimeInForce             *NewAlgoOrderTimeInForceParameter                                    `json:"timeInForce,omitempty"`
	Quantity                *float32                                                             `json:"quantity,omitempty"`
	ReduceOnly              *PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnly              `json:"reduceOnly,omitempty"`
	Price                   *float32                                                             `json:"price,omitempty"`
	NewClientOrderId        *string                                                              `json:"newClientOrderId,omitempty"`
	NewOrderRespType        *PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespType        `json:"newOrderRespType,omitempty"`
	PriceMatch              *PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatch              `json:"priceMatch,omitempty"`
	SelfTradePreventionMode *PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	// Auto-cancel time for `GTD` orders.
	GoodTillDate         *int64 `json:"goodTillDate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceMultipleOrdersBatchOrdersParameterInner PlaceMultipleOrdersBatchOrdersParameterInner

// NewPlaceMultipleOrdersBatchOrdersParameterInner instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersBatchOrdersParameterInner() *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	var side = NewAlgoOrderSideParameterBuy
	this.Side = &side
	var positionSide = PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth
	this.PositionSide = &positionSide
	var type_ = PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit
	this.Type = &type_
	var timeInForce = NewAlgoOrderTimeInForceParameterGtc
	this.TimeInForce = &timeInForce
	var reduceOnly = PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyTrue
	this.ReduceOnly = &reduceOnly
	var newOrderRespType = PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var priceMatch = PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent
	this.PriceMatch = &priceMatch
	var selfTradePreventionMode = PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeNone
	this.SelfTradePreventionMode = &selfTradePreventionMode
	return &this
}

// NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersBatchOrdersParameterInner {
	this := PlaceMultipleOrdersBatchOrdersParameterInner{}
	var side = NewAlgoOrderSideParameterBuy
	this.Side = &side
	var positionSide = PlaceMultipleOrdersBatchOrdersParameterInnerPositionSideBoth
	this.PositionSide = &positionSide
	var type_ = PlaceMultipleOrdersBatchOrdersParameterInnerTypeLimit
	this.Type = &type_
	var timeInForce = NewAlgoOrderTimeInForceParameterGtc
	this.TimeInForce = &timeInForce
	var reduceOnly = PlaceMultipleOrdersBatchOrdersParameterInnerReduceOnlyTrue
	this.ReduceOnly = &reduceOnly
	var newOrderRespType = PlaceMultipleOrdersBatchOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var priceMatch = PlaceMultipleOrdersBatchOrdersParameterInnerPriceMatchOpponent
	this.PriceMatch = &priceMatch
	var selfTradePreventionMode = PlaceMultipleOrdersBatchOrdersParameterInnerSelfTradePreventionModeNone
	this.SelfTradePreventionMode = &selfTradePreventionMode
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
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() NewAlgoOrderSideParameter {
	if o == nil || common.IsNil(o.Side) {
		var ret NewAlgoOrderSideParameter
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*NewAlgoOrderSideParameter, bool) {
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

// SetSide gets a reference to the given NewAlgoOrderSideParameter and assigns it to the Side field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v NewAlgoOrderSideParameter) {
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
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForce() NewAlgoOrderTimeInForceParameter {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret NewAlgoOrderTimeInForceParameter
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForceOk() (*NewAlgoOrderTimeInForceParameter, bool) {
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

// SetTimeInForce gets a reference to the given NewAlgoOrderTimeInForceParameter and assigns it to the TimeInForce field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetTimeInForce(v NewAlgoOrderTimeInForceParameter) {
	o.TimeInForce = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() float32 {
	if o == nil || common.IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*float32, bool) {
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

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v float32) {
	o.Quantity = &v
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

// GetGoodTillDate returns the GoodTillDate field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetGoodTillDate() int64 {
	if o == nil || common.IsNil(o.GoodTillDate) {
		var ret int64
		return ret
	}
	return *o.GoodTillDate
}

// GetGoodTillDateOk returns a tuple with the GoodTillDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetGoodTillDateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GoodTillDate) {
		return nil, false
	}
	return o.GoodTillDate, true
}

// HasGoodTillDate returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasGoodTillDate() bool {
	if o != nil && !common.IsNil(o.GoodTillDate) {
		return true
	}

	return false
}

// SetGoodTillDate gets a reference to the given int64 and assigns it to the GoodTillDate field.
func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetGoodTillDate(v int64) {
	o.GoodTillDate = &v
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
	if !common.IsNil(o.NewOrderRespType) {
		toSerialize["newOrderRespType"] = o.NewOrderRespType
	}
	if !common.IsNil(o.PriceMatch) {
		toSerialize["priceMatch"] = o.PriceMatch
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
		delete(additionalProperties, "newOrderRespType")
		delete(additionalProperties, "priceMatch")
		delete(additionalProperties, "selfTradePreventionMode")
		delete(additionalProperties, "goodTillDate")
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
