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

// checks if the PlaceMultipleOrdersResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PlaceMultipleOrdersResponseInner{}

// PlaceMultipleOrdersResponseInner struct for PlaceMultipleOrdersResponseInner
type PlaceMultipleOrdersResponseInner struct {
	OrderId              *int64  `json:"orderId,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	Price                *string `json:"price,omitempty"`
	Quantity             *string `json:"quantity,omitempty"`
	Side                 *string `json:"side,omitempty"`
	Type                 *string `json:"type,omitempty"`
	ReduceOnly           *bool   `json:"reduceOnly,omitempty"`
	PostOnly             *bool   `json:"postOnly,omitempty"`
	ClientOrderId        *string `json:"clientOrderId,omitempty"`
	Mmp                  *bool   `json:"mmp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlaceMultipleOrdersResponseInner PlaceMultipleOrdersResponseInner

// NewPlaceMultipleOrdersResponseInner instantiates a new PlaceMultipleOrdersResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceMultipleOrdersResponseInner() *PlaceMultipleOrdersResponseInner {
	this := PlaceMultipleOrdersResponseInner{}
	return &this
}

// NewPlaceMultipleOrdersResponseInnerWithDefaults instantiates a new PlaceMultipleOrdersResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceMultipleOrdersResponseInnerWithDefaults() *PlaceMultipleOrdersResponseInner {
	this := PlaceMultipleOrdersResponseInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetOrderId() int64 {
	if o == nil || common.IsNil(o.OrderId) {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetOrderIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasOrderId() bool {
	if o != nil && !common.IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *PlaceMultipleOrdersResponseInner) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PlaceMultipleOrdersResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *PlaceMultipleOrdersResponseInner) SetPrice(v string) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetQuantity() string {
	if o == nil || common.IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetQuantityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *PlaceMultipleOrdersResponseInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *PlaceMultipleOrdersResponseInner) SetSide(v string) {
	o.Side = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PlaceMultipleOrdersResponseInner) SetType(v string) {
	o.Type = &v
}

// GetReduceOnly returns the ReduceOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetReduceOnly() bool {
	if o == nil || common.IsNil(o.ReduceOnly) {
		var ret bool
		return ret
	}
	return *o.ReduceOnly
}

// GetReduceOnlyOk returns a tuple with the ReduceOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetReduceOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ReduceOnly) {
		return nil, false
	}
	return o.ReduceOnly, true
}

// HasReduceOnly returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasReduceOnly() bool {
	if o != nil && !common.IsNil(o.ReduceOnly) {
		return true
	}

	return false
}

// SetReduceOnly gets a reference to the given bool and assigns it to the ReduceOnly field.
func (o *PlaceMultipleOrdersResponseInner) SetReduceOnly(v bool) {
	o.ReduceOnly = &v
}

// GetPostOnly returns the PostOnly field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetPostOnly() bool {
	if o == nil || common.IsNil(o.PostOnly) {
		var ret bool
		return ret
	}
	return *o.PostOnly
}

// GetPostOnlyOk returns a tuple with the PostOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetPostOnlyOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PostOnly) {
		return nil, false
	}
	return o.PostOnly, true
}

// HasPostOnly returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasPostOnly() bool {
	if o != nil && !common.IsNil(o.PostOnly) {
		return true
	}

	return false
}

// SetPostOnly gets a reference to the given bool and assigns it to the PostOnly field.
func (o *PlaceMultipleOrdersResponseInner) SetPostOnly(v bool) {
	o.PostOnly = &v
}

// GetClientOrderId returns the ClientOrderId field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetClientOrderId() string {
	if o == nil || common.IsNil(o.ClientOrderId) {
		var ret string
		return ret
	}
	return *o.ClientOrderId
}

// GetClientOrderIdOk returns a tuple with the ClientOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetClientOrderIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientOrderId) {
		return nil, false
	}
	return o.ClientOrderId, true
}

// HasClientOrderId returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasClientOrderId() bool {
	if o != nil && !common.IsNil(o.ClientOrderId) {
		return true
	}

	return false
}

// SetClientOrderId gets a reference to the given string and assigns it to the ClientOrderId field.
func (o *PlaceMultipleOrdersResponseInner) SetClientOrderId(v string) {
	o.ClientOrderId = &v
}

// GetMmp returns the Mmp field value if set, zero value otherwise.
func (o *PlaceMultipleOrdersResponseInner) GetMmp() bool {
	if o == nil || common.IsNil(o.Mmp) {
		var ret bool
		return ret
	}
	return *o.Mmp
}

// GetMmpOk returns a tuple with the Mmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaceMultipleOrdersResponseInner) GetMmpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Mmp) {
		return nil, false
	}
	return o.Mmp, true
}

// HasMmp returns a boolean if a field has been set.
func (o *PlaceMultipleOrdersResponseInner) HasMmp() bool {
	if o != nil && !common.IsNil(o.Mmp) {
		return true
	}

	return false
}

// SetMmp gets a reference to the given bool and assigns it to the Mmp field.
func (o *PlaceMultipleOrdersResponseInner) SetMmp(v bool) {
	o.Mmp = &v
}

func (o PlaceMultipleOrdersResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceMultipleOrdersResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !common.IsNil(o.ReduceOnly) {
		toSerialize["reduceOnly"] = o.ReduceOnly
	}
	if !common.IsNil(o.PostOnly) {
		toSerialize["postOnly"] = o.PostOnly
	}
	if !common.IsNil(o.ClientOrderId) {
		toSerialize["clientOrderId"] = o.ClientOrderId
	}
	if !common.IsNil(o.Mmp) {
		toSerialize["mmp"] = o.Mmp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlaceMultipleOrdersResponseInner) UnmarshalJSON(data []byte) (err error) {
	varPlaceMultipleOrdersResponseInner := _PlaceMultipleOrdersResponseInner{}

	err = json.Unmarshal(data, &varPlaceMultipleOrdersResponseInner)

	if err != nil {
		return err
	}

	*o = PlaceMultipleOrdersResponseInner(varPlaceMultipleOrdersResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orderId")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "price")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "side")
		delete(additionalProperties, "type")
		delete(additionalProperties, "reduceOnly")
		delete(additionalProperties, "postOnly")
		delete(additionalProperties, "clientOrderId")
		delete(additionalProperties, "mmp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlaceMultipleOrdersResponseInner struct {
	value *PlaceMultipleOrdersResponseInner
	isSet bool
}

func (v NullablePlaceMultipleOrdersResponseInner) Get() *PlaceMultipleOrdersResponseInner {
	return v.value
}

func (v *NullablePlaceMultipleOrdersResponseInner) Set(val *PlaceMultipleOrdersResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceMultipleOrdersResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceMultipleOrdersResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceMultipleOrdersResponseInner(val *PlaceMultipleOrdersResponseInner) *NullablePlaceMultipleOrdersResponseInner {
	return &NullablePlaceMultipleOrdersResponseInner{value: val, isSet: true}
}

func (v NullablePlaceMultipleOrdersResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceMultipleOrdersResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
