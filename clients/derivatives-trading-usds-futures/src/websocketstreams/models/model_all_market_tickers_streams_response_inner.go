/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AllMarketTickersStreamsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketTickersStreamsResponseInner{}

// AllMarketTickersStreamsResponseInner struct for AllMarketTickersStreamsResponseInner
type AllMarketTickersStreamsResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	P                    *string `json:"P,omitempty"`
	Smallw               *string `json:"w,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallh               *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	O                    *int64  `json:"O,omitempty"`
	C                    *int64  `json:"C,omitempty"`
	F                    *int64  `json:"F,omitempty"`
	L                    *int64  `json:"L,omitempty"`
	Smalln               *int64  `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMarketTickersStreamsResponseInner AllMarketTickersStreamsResponseInner

// NewAllMarketTickersStreamsResponseInner instantiates a new AllMarketTickersStreamsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketTickersStreamsResponseInner() *AllMarketTickersStreamsResponseInner {
	this := AllMarketTickersStreamsResponseInner{}
	return &this
}

// NewAllMarketTickersStreamsResponseInnerWithDefaults instantiates a new AllMarketTickersStreamsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketTickersStreamsResponseInnerWithDefaults() *AllMarketTickersStreamsResponseInner {
	this := AllMarketTickersStreamsResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllMarketTickersStreamsResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllMarketTickersStreamsResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMarketTickersStreamsResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllMarketTickersStreamsResponseInner) SetP(v string) {
	o.P = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallw() string {
	if o == nil || common.IsNil(o.Smallw) {
		var ret string
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallw(v string) {
	o.Smallw = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMarketTickersStreamsResponseInner) SetQ(v string) {
	o.Q = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *AllMarketTickersStreamsResponseInner) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMarketTickersStreamsResponseInner) SetSmallq(v string) {
	o.Smallq = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *AllMarketTickersStreamsResponseInner) SetO(v int64) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetC() int64 {
	if o == nil || common.IsNil(o.C) {
		var ret int64
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetCOk() (*int64, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given int64 and assigns it to the C field.
func (o *AllMarketTickersStreamsResponseInner) SetC(v int64) {
	o.C = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AllMarketTickersStreamsResponseInner) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AllMarketTickersStreamsResponseInner) SetL(v int64) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *AllMarketTickersStreamsResponseInner) GetSmalln() int64 {
	if o == nil || common.IsNil(o.Smalln) {
		var ret int64
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketTickersStreamsResponseInner) GetSmallnOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *AllMarketTickersStreamsResponseInner) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *AllMarketTickersStreamsResponseInner) SetSmalln(v int64) {
	o.Smalln = &v
}

func (o AllMarketTickersStreamsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketTickersStreamsResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Smallw) {
		toSerialize["w"] = o.Smallw
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}
	if !common.IsNil(o.C) {
		toSerialize["C"] = o.C
	}
	if !common.IsNil(o.F) {
		toSerialize["F"] = o.F
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllMarketTickersStreamsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAllMarketTickersStreamsResponseInner := _AllMarketTickersStreamsResponseInner{}

	err = json.Unmarshal(data, &varAllMarketTickersStreamsResponseInner)

	if err != nil {
		return err
	}

	*o = AllMarketTickersStreamsResponseInner(varAllMarketTickersStreamsResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "P")
		delete(additionalProperties, "w")
		delete(additionalProperties, "c")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		delete(additionalProperties, "O")
		delete(additionalProperties, "C")
		delete(additionalProperties, "F")
		delete(additionalProperties, "L")
		delete(additionalProperties, "n")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllMarketTickersStreamsResponseInner struct {
	value *AllMarketTickersStreamsResponseInner
	isSet bool
}

func (v NullableAllMarketTickersStreamsResponseInner) Get() *AllMarketTickersStreamsResponseInner {
	return v.value
}

func (v *NullableAllMarketTickersStreamsResponseInner) Set(val *AllMarketTickersStreamsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketTickersStreamsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketTickersStreamsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMarketTickersStreamsResponseInner(val *AllMarketTickersStreamsResponseInner) *NullableAllMarketTickersStreamsResponseInner {
	return &NullableAllMarketTickersStreamsResponseInner{value: val, isSet: true}
}

func (v NullableAllMarketTickersStreamsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketTickersStreamsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
