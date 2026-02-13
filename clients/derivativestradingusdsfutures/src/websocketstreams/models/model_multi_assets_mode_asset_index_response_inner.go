/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MultiAssetsModeAssetIndexResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MultiAssetsModeAssetIndexResponseInner{}

// MultiAssetsModeAssetIndexResponseInner struct for MultiAssetsModeAssetIndexResponseInner
type MultiAssetsModeAssetIndexResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	B                    *string `json:"B,omitempty"`
	A                    *string `json:"A,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallg               *string `json:"g,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	G                    *string `json:"G,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MultiAssetsModeAssetIndexResponseInner MultiAssetsModeAssetIndexResponseInner

// NewMultiAssetsModeAssetIndexResponseInner instantiates a new MultiAssetsModeAssetIndexResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAssetsModeAssetIndexResponseInner() *MultiAssetsModeAssetIndexResponseInner {
	this := MultiAssetsModeAssetIndexResponseInner{}
	return &this
}

// NewMultiAssetsModeAssetIndexResponseInnerWithDefaults instantiates a new MultiAssetsModeAssetIndexResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAssetsModeAssetIndexResponseInnerWithDefaults() *MultiAssetsModeAssetIndexResponseInner {
	this := MultiAssetsModeAssetIndexResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmalli(v string) {
	o.Smalli = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetB(v string) {
	o.B = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetA(v string) {
	o.A = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallg() string {
	if o == nil || common.IsNil(o.Smallg) {
		var ret string
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetSmallgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetSmallg(v string) {
	o.Smallg = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetQ(v string) {
	o.Q = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *MultiAssetsModeAssetIndexResponseInner) GetG() string {
	if o == nil || common.IsNil(o.G) {
		var ret string
		return ret
	}
	return *o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) GetGOk() (*string, bool) {
	if o == nil || common.IsNil(o.G) {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *MultiAssetsModeAssetIndexResponseInner) HasG() bool {
	if o != nil && !common.IsNil(o.G) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *MultiAssetsModeAssetIndexResponseInner) SetG(v string) {
	o.G = &v
}

func (o MultiAssetsModeAssetIndexResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiAssetsModeAssetIndexResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.A) {
		toSerialize["A"] = o.A
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallg) {
		toSerialize["g"] = o.Smallg
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.G) {
		toSerialize["G"] = o.G
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MultiAssetsModeAssetIndexResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMultiAssetsModeAssetIndexResponseInner := _MultiAssetsModeAssetIndexResponseInner{}

	err = json.Unmarshal(data, &varMultiAssetsModeAssetIndexResponseInner)

	if err != nil {
		return err
	}

	*o = MultiAssetsModeAssetIndexResponseInner(varMultiAssetsModeAssetIndexResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		delete(additionalProperties, "B")
		delete(additionalProperties, "A")
		delete(additionalProperties, "q")
		delete(additionalProperties, "g")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "G")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMultiAssetsModeAssetIndexResponseInner struct {
	value *MultiAssetsModeAssetIndexResponseInner
	isSet bool
}

func (v NullableMultiAssetsModeAssetIndexResponseInner) Get() *MultiAssetsModeAssetIndexResponseInner {
	return v.value
}

func (v *NullableMultiAssetsModeAssetIndexResponseInner) Set(val *MultiAssetsModeAssetIndexResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAssetsModeAssetIndexResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAssetsModeAssetIndexResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiAssetsModeAssetIndexResponseInner(val *MultiAssetsModeAssetIndexResponseInner) *NullableMultiAssetsModeAssetIndexResponseInner {
	return &NullableMultiAssetsModeAssetIndexResponseInner{value: val, isSet: true}
}

func (v NullableMultiAssetsModeAssetIndexResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiAssetsModeAssetIndexResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
