/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PlaceMultipleOrdersOrdersParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersOrdersParameterInner{}

// PlaceMultipleOrdersOrdersParameterInner struct for PlaceMultipleOrdersOrdersParameterInner
type PlaceMultipleOrdersOrdersParameterInner struct {
	Symbol               *string                                                  `json:"symbol,omitempty"`
	Side                 *PlaceMultipleOrdersOrdersParameterInnerSide             `json:"side,omitempty"`
	Type                 *PlaceMultipleOrdersOrdersParameterInnerType             `json:"type,omitempty"`
	Quantity             *string                                                  `json:"quantity,omitempty"`
	Price                *string                                                  `json:"price,omitempty"`
	TimeInForce          *PlaceMultipleOrdersOrdersParameterInnerTimeInForce      `json:"timeInForce,omitempty"`
	ReduceOnly           *string                                                  `json:"reduceOnly,omitempty"`
	PostOnly             *string                                                  `json:"postOnly,omitempty"`
	NewOrderRespType     *PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType `json:"newOrderRespType,omitempty"`
	ClientOrderId        *string                                                  `json:"clientOrderId,omitempty"`
	IsMmp                *string                                                  `json:"isMmp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceMultipleOrdersOrdersParameterInner PlaceMultipleOrdersOrdersParameterInner

// NewPlaceMultipleOrdersOrdersParameterInner instantiates a new PlaceMultipleOrdersOrdersParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersOrdersParameterInner() *PlaceMultipleOrdersOrdersParameterInner {
	this := PlaceMultipleOrdersOrdersParameterInner{}
	return &this
}

// NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersOrdersParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersOrdersParameterInner {
	this := PlaceMultipleOrdersOrdersParameterInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSide() PlaceMultipleOrdersOrdersParameterInnerSide {
	if o == nil || common.IsNil(o.Side) {
		var ret PlaceMultipleOrdersOrdersParameterInnerSide
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetSideOk() (*PlaceMultipleOrdersOrdersParameterInnerSide, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given PlaceMultipleOrdersOrdersParameterInnerSide and assigns it to the Side field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetSide(v PlaceMultipleOrdersOrdersParameterInnerSide) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetType() PlaceMultipleOrdersOrdersParameterInnerType {
	if o == nil || common.IsNil(o.Type) {
		var ret PlaceMultipleOrdersOrdersParameterInnerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersOrdersParameterInnerType, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PlaceMultipleOrdersOrdersParameterInnerType and assigns it to the Type field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetType(v PlaceMultipleOrdersOrdersParameterInnerType) {
	o.Type = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPriceOk() (*string, bool) {
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

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetPrice(v string) {
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
func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnly() string {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret string
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnlyOk() (*string, bool) {
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

// SetReduceOnly gets a reference to the given string and assigns it to the ReduceOnly field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetReduceOnly(v string) {
	o.ReduceOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnly() string {
	if o == nil || common.IsNil(o.PostOnly) {
		var ret string
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnlyOk() (*string, bool) {
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

// SetPostOnly gets a reference to the given string and assigns it to the PostOnly field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetPostOnly(v string) {
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
func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmp() string {
	if o == nil || common.IsNil(o.IsMmp) {
		var ret string
		return ret
	}
	return *o.IsMmp
}

// GetIsMmpOk returns a tuple with the IsMmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmpOk() (*string, bool) {
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

// SetIsMmp gets a reference to the given string and assigns it to the IsMmp field.
func (o *PlaceMultipleOrdersOrdersParameterInner) SetIsMmp(v string) {
	o.IsMmp = &v
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
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceMultipleOrdersOrdersParameterInner) UnmarshalJSON(data []byte) (err error) {
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
