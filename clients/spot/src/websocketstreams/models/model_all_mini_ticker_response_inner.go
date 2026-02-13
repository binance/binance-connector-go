/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllMiniTickerResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMiniTickerResponseInner{}

// AllMiniTickerResponseInner struct for AllMiniTickerResponseInner
type AllMiniTickerResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	S                    *string `json:"s,omitempty"`
	C                    *string `json:"c,omitempty"`
	O                    *string `json:"o,omitempty"`
	H                    *string `json:"h,omitempty"`
	L                    *string `json:"l,omitempty"`
	V                    *string `json:"v,omitempty"`
	Q                    *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMiniTickerResponseInner AllMiniTickerResponseInner

// NewAllMiniTickerResponseInner instantiates a new AllMiniTickerResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMiniTickerResponseInner() *AllMiniTickerResponseInner {
	this := AllMiniTickerResponseInner{}
	return &this
}

// NewAllMiniTickerResponseInnerWithDefaults instantiates a new AllMiniTickerResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMiniTickerResponseInnerWithDefaults() *AllMiniTickerResponseInner {
	this := AllMiniTickerResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllMiniTickerResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllMiniTickerResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMiniTickerResponseInner) SetS(v string) {
	o.S = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *AllMiniTickerResponseInner) SetC(v string) {
	o.C = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetO() string {
	if o == nil || common.IsNil(o.O) {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetOOk() (*string, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AllMiniTickerResponseInner) SetO(v string) {
	o.O = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetH() string {
	if o == nil || common.IsNil(o.H) {
		var ret string
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetHOk() (*string, bool) {
	if o == nil || common.IsNil(o.H) {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasH() bool {
	if o != nil && !common.IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *AllMiniTickerResponseInner) SetH(v string) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *AllMiniTickerResponseInner) SetL(v string) {
	o.L = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AllMiniTickerResponseInner) SetV(v string) {
	o.V = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMiniTickerResponseInner) SetQ(v string) {
	o.Q = &v
}

func (o AllMiniTickerResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMiniTickerResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	if !common.IsNil(o.O) {
		toSerialize["o"] = o.O
	}
	if !common.IsNil(o.H) {
		toSerialize["h"] = o.H
	}
	if !common.IsNil(o.L) {
		toSerialize["l"] = o.L
	}
	if !common.IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !common.IsNil(o.Q) {
		toSerialize["q"] = o.Q
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllMiniTickerResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAllMiniTickerResponseInner := _AllMiniTickerResponseInner{}

	err = json.Unmarshal(data, &varAllMiniTickerResponseInner)

	if err != nil {
		return err
	}

	*o = AllMiniTickerResponseInner(varAllMiniTickerResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllMiniTickerResponseInner struct {
	value *AllMiniTickerResponseInner
	isSet bool
}

func (v NullableAllMiniTickerResponseInner) Get() *AllMiniTickerResponseInner {
	return v.value
}

func (v *NullableAllMiniTickerResponseInner) Set(val *AllMiniTickerResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMiniTickerResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMiniTickerResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMiniTickerResponseInner(val *AllMiniTickerResponseInner) *NullableAllMiniTickerResponseInner {
	return &NullableAllMiniTickerResponseInner{value: val, isSet: true}
}

func (v NullableAllMiniTickerResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMiniTickerResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
