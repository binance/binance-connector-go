/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OpenOrderLossOInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenOrderLossOInner{}

// OpenOrderLossOInner struct for OpenOrderLossOInner
type OpenOrderLossOInner struct {
	// Asset
	Smalla *string `json:"a,omitempty"`
	// Amount
	Smallo               *string `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenOrderLossOInner OpenOrderLossOInner

// NewOpenOrderLossOInner instantiates a new OpenOrderLossOInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenOrderLossOInner() *OpenOrderLossOInner {
	this := OpenOrderLossOInner{}
	return &this
}

// NewOpenOrderLossOInnerWithDefaults instantiates a new OpenOrderLossOInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenOrderLossOInnerWithDefaults() *OpenOrderLossOInner {
	this := OpenOrderLossOInner{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OpenOrderLossOInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderLossOInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OpenOrderLossOInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OpenOrderLossOInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OpenOrderLossOInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenOrderLossOInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OpenOrderLossOInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *OpenOrderLossOInner) SetSmallo(v string) {
	o.Smallo = &v
}

func (o OpenOrderLossOInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenOrderLossOInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenOrderLossOInner) UnmarshalJSON(data []byte) (err error) {
	varOpenOrderLossOInner := _OpenOrderLossOInner{}

	err = json.Unmarshal(data, &varOpenOrderLossOInner)

	if err != nil {
		return err
	}

	*o = OpenOrderLossOInner(varOpenOrderLossOInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "a")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenOrderLossOInner struct {
	value *OpenOrderLossOInner
	isSet bool
}

func (v NullableOpenOrderLossOInner) Get() *OpenOrderLossOInner {
	return v.value
}

func (v *NullableOpenOrderLossOInner) Set(val *OpenOrderLossOInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenOrderLossOInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenOrderLossOInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenOrderLossOInner(val *OpenOrderLossOInner) *NullableOpenOrderLossOInner {
	return &NullableOpenOrderLossOInner{value: val, isSet: true}
}

func (v NullableOpenOrderLossOInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenOrderLossOInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
