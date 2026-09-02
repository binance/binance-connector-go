/*
Options REST API

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"
	"fmt"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PlaceMultipleOrdersOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersOrdersParameterInner{}

// PlaceMultipleOrdersOrdersParameterInner struct for PlaceMultipleOrdersOrdersParameterInner
type PlaceMultipleOrdersOrdersParameterInner struct {
	Symbol string                                      `json:"symbol"`
	Side   PlaceMultipleOrdersOrdersParameterInnerSide `json:"side"`
	Type   PlaceMultipleOrdersOrdersParameterInnerType `json:"type"`
	// Order Quantity
	Quantity float64 `json:"quantity"`
	// Order Price
	Price            *float64                                                 `json:"price,omitempty"`
	TimeInForce      *PlaceMultipleOrdersOrdersParameterInnerTimeInForce      `json:"timeInForce,omitempty"`
	ReduceOnly       *bool                                                    `json:"reduceOnly,omitempty"`
	PostOnly         *bool                                                    `json:"postOnly,omitempty"`
	NewOrderRespType *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType `json:"newOrderRespType,omitempty"`
	// User-defined order ID cannot be repeated in pending orders
	ClientOrderId *string `json:"clientOrderId,omitempty"`
	// is market maker protection order
	IsMmp                   *bool                                                           `json:"isMmp,omitempty"`
	SelfTradePreventionMode *PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode `json:"selfTradePreventionMode,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _PlaceMultipleOrdersOrdersParameterInner PlaceMultipleOrdersOrdersParameterInner

// NewPlaceMultipleOrdersOrdersParameterInner instantiates a new PlaceMultipleOrdersOrdersParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersOrdersParameterInner(symbol string, side PlaceMultipleOrdersOrdersParameterInnerSide, type_ PlaceMultipleOrdersOrdersParameterInnerType, quantity float64) *PlaceMultipleOrdersOrdersParameterInner {
	this := PlaceMultipleOrdersOrdersParameterInner{}
	this.Symbol = symbol
	this.Side = side
	this.Type = type_
	this.Quantity = quantity
	var timeInForce = PlaceMultipleOrdersOrdersParameterInnerTimeInForceGtc
	this.TimeInForce = &timeInForce
	var reduceOnly = false
	this.ReduceOnly = &reduceOnly
	var postOnly = false
	this.PostOnly = &postOnly
	var newOrderRespType = PlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var selfTradePreventionMode = PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeExpireTaker
	this.SelfTradePreventionMode = &selfTradePreventionMode
	return &this
}

// NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersOrdersParameterInner {
	this := PlaceMultipleOrdersOrdersParameterInner{}
	var timeInForce = PlaceMultipleOrdersOrdersParameterInnerTimeInForceGtc
	this.TimeInForce = &timeInForce
	var reduceOnly = false
	this.ReduceOnly = &reduceOnly
	var postOnly = false
	this.PostOnly = &postOnly
	var newOrderRespType = PlaceMultipleOrdersOrdersParameterInnerNewOrderRespTypeAck
	this.NewOrderRespType = &newOrderRespType
	var selfTradePreventionMode = PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionModeExpireTaker
	this.SelfTradePreventionMode = &selfTradePreventionMode
	return &this
}

// GetSymbol returns the Symbol field value
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *PlaceMultipleOrdersOrdersParameterInner) SetSymbol(v string) {
	o.Symbol = v
}

// GetSide returns the Side field value
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSide() PlaceMultipleOrdersOrdersParameterInnerSide {
	if o == nil {
		var ret PlaceMultipleOrdersOrdersParameterInnerSide
		return ret
	}

	return o.Side
}

// GetSideOk returns a tuple with the Side field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSideOk() (*PlaceMultipleOrdersOrdersParameterInnerSide, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Side, true
}

// SetSide sets field value
func (o *PlaceMultipleOrdersOrdersParameterInner) SetSide(v PlaceMultipleOrdersOrdersParameterInnerSide) {
	o.Side = v
}

// GetType returns the Type field value
func (o *PlaceMultipleOrdersOrdersParameterInner) GetType() PlaceMultipleOrdersOrdersParameterInnerType {
	if o == nil {
		var ret PlaceMultipleOrdersOrdersParameterInnerType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersOrdersParameterInnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PlaceMultipleOrdersOrdersParameterInner) SetType(v PlaceMultipleOrdersOrdersParameterInnerType) {
	o.Type = v
}

// GetQuantity returns the Quantity field value
func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *PlaceMultipleOrdersOrdersParameterInner) SetQuantity(v float64) {
	o.Quantity = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPrice() float64 {
	if o == nil || common.IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPriceOk() (*float64, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetPrice(v float64) {
	o.Price = &v
}

// GetTimeInForce returns the TimeInForce field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetTimeInForce() PlaceMultipleOrdersOrdersParameterInnerTimeInForce {
	if o == nil || common.IsNil(o.TimeInForce) {
		var ret PlaceMultipleOrdersOrdersParameterInnerTimeInForce
		return ret
	}
	return *o.TimeInForce
}

// GetTimeInForceOk returns a tuple with the TimeInForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetTimeInForceOk() (*PlaceMultipleOrdersOrdersParameterInnerTimeInForce, bool) {
	if o == nil || common.IsNil(o.TimeInForce) {
		return nil, false
	}
	return o.TimeInForce, true
}

// HasTimeInForce returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasTimeInForce() bool {
	if o != nil && !common.IsNil(o.TimeInForce) {
		return true
	}

	return false
}

// SetTimeInForce gets a reference to the given PlaceMultipleOrdersOrdersParameterInnerTimeInForce and assigns it to the TimeInForce field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetTimeInForce(v PlaceMultipleOrdersOrdersParameterInnerTimeInForce) {
	o.TimeInForce = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnly() bool {
	if o == nil || common.IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasPostOnly() bool {
	if o != nil && !common.IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetNewOrderRespType returns the NewOrderRespType field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetNewOrderRespType() PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType {
	if o == nil || common.IsNil(o.NewOrderRespType) {
		var ret PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType
		return ret
	}
	return *o.NewOrderRespType
}

// GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetNewOrderRespTypeOk() (*PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType, bool) {
	if o == nil || common.IsNil(o.NewOrderRespType) {
		return nil, false
	}
	return o.NewOrderRespType, true
}

// HasNewOrderRespType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasNewOrderRespType() bool {
	if o != nil && !common.IsNil(o.NewOrderRespType) {
		return true
	}

	return false
}

// SetNewOrderRespType gets a reference to the given PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType and assigns it to the NewOrderRespType field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetNewOrderRespType(v PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType) {
	o.NewOrderRespType = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetIsMmp returns the IsMmp field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmp() bool {
	if o == nil || common.IsNil(o.IsMmp) {
		var ret bool
		return ret
	}
	return *o.IsMmp
}

// GetIsMmpOk returns a tuple with the IsMmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMmp) {
		return nil, false
	}
	return o.IsMmp, true
}

// HasIsMmp returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasIsMmp() bool {
	if o != nil && !common.IsNil(o.IsMmp) {
		return true
	}

	return false
}

// SetIsMmp gets a reference to the given bool and assigns it to the IsMmp field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetIsMmp(v bool) {
	o.IsMmp = &v
}

// GetSelfTradePreventionMode returns the SelfTradePreventionMode field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSelfTradePreventionMode() PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		var ret PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode
		return ret
	}
	return *o.SelfTradePreventionMode
}

// GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSelfTradePreventionModeOk() (*PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode, bool) {
	if o == nil || common.IsNil(o.SelfTradePreventionMode) {
		return nil, false
	}
	return o.SelfTradePreventionMode, true
}

// HasSelfTradePreventionMode returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasSelfTradePreventionMode() bool {
	if o != nil && !common.IsNil(o.SelfTradePreventionMode) {
		return true
	}

	return false
}

// SetSelfTradePreventionMode gets a reference to the given PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode and assigns it to the SelfTradePreventionMode field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetSelfTradePreventionMode(v PlaceMultipleOrdersOrdersParameterInnerSelfTradePreventionMode) {
	o.SelfTradePreventionMode = &v
}

func (o PlaceMultipleOrdersOrdersParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceMultipleOrdersOrdersParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["side"] = o.Side
	toSerialize["type"] = o.Type
	toSerialize["quantity"] = o.Quantity
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.TimeInForce) {
		toSerialize["timeInForce"] = o.TimeInForce
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.PostOnly) {
		toSerialize["postOnly"] = o.PostOnly
	}
	if !common.IsNil(o.NewOrderRespType) {
		toSerialize["newOrderRespType"] = o.NewOrderRespType
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.IsMmp) {
		toSerialize["isMmp"] = o.IsMmp
	}
	if !common.IsNil(o.SelfTradePreventionMode) {
		toSerialize["selfTradePreventionMode"] = o.SelfTradePreventionMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceMultipleOrdersOrdersParameterInner) UnmarshalJSON(data []byte) (err error) {
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

	varPlaceMultipleOrdersOrdersParameterInner := _PlaceMultipleOrdersOrdersParameterInner{}

	err = json.Unmarshal(data, &varPlaceMultipleOrdersOrdersParameterInner)

	if err != nil {
		return err
	}

	*o = PlaceMultipleOrdersOrdersParameterInner(varPlaceMultipleOrdersOrdersParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "side")
		delete(additionalProperties, "type")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "price")
		delete(additionalProperties, "timeInForce")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "postOnly")
		delete(additionalProperties, "newOrderRespType")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "isMmp")
		delete(additionalProperties, "selfTradePreventionMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceMultipleOrdersOrdersParameterInner struct {
	value *PlaceMultipleOrdersOrdersParameterInner
	isSet bool
}

func (v NullablePlaceMultipleOrdersOrdersParameterInner) Get() *PlaceMultipleOrdersOrdersParameterInner {
	return v.value
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInner) Set(val *PlaceMultipleOrdersOrdersParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersOrdersParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersOrdersParameterInner(val *PlaceMultipleOrdersOrdersParameterInner) *NullablePlaceMultipleOrdersOrdersParameterInner {
	return &NullablePlaceMultipleOrdersOrdersParameterInner{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersOrdersParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersOrdersParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
