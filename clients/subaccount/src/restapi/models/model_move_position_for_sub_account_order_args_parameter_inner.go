/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MovePositionForSubAccountOrderArgsParameterInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MovePositionForSubAccountOrderArgsParameterInner{}

// MovePositionForSubAccountOrderArgsParameterInner struct for MovePositionForSubAccountOrderArgsParameterInner
type MovePositionForSubAccountOrderArgsParameterInner struct {
	Symbol               *string  `json:"symbol,omitempty"`
	Quantity             *float32 `json:"quantity,omitempty"`
	PositionSide         *string  `json:"positionSide,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MovePositionForSubAccountOrderArgsParameterInner MovePositionForSubAccountOrderArgsParameterInner

// NewMovePositionForSubAccountOrderArgsParameterInner instantiates a new MovePositionForSubAccountOrderArgsParameterInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovePositionForSubAccountOrderArgsParameterInner() *MovePositionForSubAccountOrderArgsParameterInner {
	this := MovePositionForSubAccountOrderArgsParameterInner{}
	return &this
}

// NewMovePositionForSubAccountOrderArgsParameterInnerWithDefaults instantiates a new MovePositionForSubAccountOrderArgsParameterInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovePositionForSubAccountOrderArgsParameterInnerWithDefaults() *MovePositionForSubAccountOrderArgsParameterInner {
	this := MovePositionForSubAccountOrderArgsParameterInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MovePositionForSubAccountOrderArgsParameterInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetQuantity() float32 {
	if o == nil || common.IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetQuantityOk() (*float32, bool) {
	if o == nil || common.IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) HasQuantity() bool {
	if o != nil && !common.IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *MovePositionForSubAccountOrderArgsParameterInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetPositionSide returns the PositionSide field value if set, zero value otherwise.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetPositionSide() string {
	if o == nil || common.IsNil(o.PositionSide) {
		var ret string
		return ret
	}
	return *o.PositionSide
}

// GetPositionSideOk returns a tuple with the PositionSide field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) GetPositionSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.PositionSide) {
		return nil, false
	}
	return o.PositionSide, true
}

// HasPositionSide returns a boolean if a field has been set.
func (o *MovePositionForSubAccountOrderArgsParameterInner) HasPositionSide() bool {
	if o != nil && !common.IsNil(o.PositionSide) {
		return true
	}

	return false
}

// SetPositionSide gets a reference to the given string and assigns it to the PositionSide field.
func (o *MovePositionForSubAccountOrderArgsParameterInner) SetPositionSide(v string) {
	o.PositionSide = &v
}

func (o MovePositionForSubAccountOrderArgsParameterInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovePositionForSubAccountOrderArgsParameterInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !common.IsNil(o.PositionSide) {
		toSerialize["positionSide"] = o.PositionSide
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MovePositionForSubAccountOrderArgsParameterInner) UnmarshalJSON(data []byte) (err error) {
	varMovePositionForSubAccountOrderArgsParameterInner := _MovePositionForSubAccountOrderArgsParameterInner{}

	err = json.Unmarshal(data, &varMovePositionForSubAccountOrderArgsParameterInner)

	if err != nil {
		return err
	}

	*o = MovePositionForSubAccountOrderArgsParameterInner(varMovePositionForSubAccountOrderArgsParameterInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "positionSide")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMovePositionForSubAccountOrderArgsParameterInner struct {
	value *MovePositionForSubAccountOrderArgsParameterInner
	isSet bool
}

func (v NullableMovePositionForSubAccountOrderArgsParameterInner) Get() *MovePositionForSubAccountOrderArgsParameterInner {
	return v.value
}

func (v *NullableMovePositionForSubAccountOrderArgsParameterInner) Set(val *MovePositionForSubAccountOrderArgsParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePositionForSubAccountOrderArgsParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePositionForSubAccountOrderArgsParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePositionForSubAccountOrderArgsParameterInner(val *MovePositionForSubAccountOrderArgsParameterInner) *NullableMovePositionForSubAccountOrderArgsParameterInner {
	return &NullableMovePositionForSubAccountOrderArgsParameterInner{value: val, isSet: true}
}

func (v NullableMovePositionForSubAccountOrderArgsParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePositionForSubAccountOrderArgsParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
