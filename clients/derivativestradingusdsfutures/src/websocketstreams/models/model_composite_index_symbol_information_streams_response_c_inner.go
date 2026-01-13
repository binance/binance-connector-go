/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CompositeIndexSymbolInformationStreamsResponseCInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompositeIndexSymbolInformationStreamsResponseCInner{}

// CompositeIndexSymbolInformationStreamsResponseCInner struct for CompositeIndexSymbolInformationStreamsResponseCInner
type CompositeIndexSymbolInformationStreamsResponseCInner struct {
	Smallb               *string `json:"b,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallw               *string `json:"w,omitempty"`
	W                    *string `json:"W,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompositeIndexSymbolInformationStreamsResponseCInner CompositeIndexSymbolInformationStreamsResponseCInner

// NewCompositeIndexSymbolInformationStreamsResponseCInner instantiates a new CompositeIndexSymbolInformationStreamsResponseCInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeIndexSymbolInformationStreamsResponseCInner() *CompositeIndexSymbolInformationStreamsResponseCInner {
	this := CompositeIndexSymbolInformationStreamsResponseCInner{}
	return &this
}

// NewCompositeIndexSymbolInformationStreamsResponseCInnerWithDefaults instantiates a new CompositeIndexSymbolInformationStreamsResponseCInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeIndexSymbolInformationStreamsResponseCInnerWithDefaults() *CompositeIndexSymbolInformationStreamsResponseCInner {
	this := CompositeIndexSymbolInformationStreamsResponseCInner{}
	return &this
}

// GetB returns the B field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallw() string {
	if o == nil || common.IsNil(o.Smallw) {
		var ret string
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmallwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) SetSmallw(v string) {
	o.Smallw = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetW() string {
	if o == nil || common.IsNil(o.W) {
		var ret string
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetWOk() (*string, bool) {
	if o == nil || common.IsNil(o.W) {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) HasW() bool {
	if o != nil && !common.IsNil(o.W) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) SetW(v string) {
	o.W = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *CompositeIndexSymbolInformationStreamsResponseCInner) SetSmalli(v string) {
	o.Smalli = &v
}

func (o CompositeIndexSymbolInformationStreamsResponseCInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeIndexSymbolInformationStreamsResponseCInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallw) {
		toSerialize["w"] = o.Smallw
	}
	if !common.IsNil(o.W) {
		toSerialize["W"] = o.W
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompositeIndexSymbolInformationStreamsResponseCInner) UnmarshalJSON(data []byte) (err error) {
	varCompositeIndexSymbolInformationStreamsResponseCInner := _CompositeIndexSymbolInformationStreamsResponseCInner{}

	err = json.Unmarshal(data, &varCompositeIndexSymbolInformationStreamsResponseCInner)

	if err != nil {
		return err
	}

	*o = CompositeIndexSymbolInformationStreamsResponseCInner(varCompositeIndexSymbolInformationStreamsResponseCInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "b")
		delete(additionalProperties, "q")
		delete(additionalProperties, "w")
		delete(additionalProperties, "W")
		delete(additionalProperties, "i")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompositeIndexSymbolInformationStreamsResponseCInner struct {
	value *CompositeIndexSymbolInformationStreamsResponseCInner
	isSet bool
}

func (v NullableCompositeIndexSymbolInformationStreamsResponseCInner) Get() *CompositeIndexSymbolInformationStreamsResponseCInner {
	return v.value
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponseCInner) Set(val *CompositeIndexSymbolInformationStreamsResponseCInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeIndexSymbolInformationStreamsResponseCInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponseCInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeIndexSymbolInformationStreamsResponseCInner(val *CompositeIndexSymbolInformationStreamsResponseCInner) *NullableCompositeIndexSymbolInformationStreamsResponseCInner {
	return &NullableCompositeIndexSymbolInformationStreamsResponseCInner{value: val, isSet: true}
}

func (v NullableCompositeIndexSymbolInformationStreamsResponseCInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponseCInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
