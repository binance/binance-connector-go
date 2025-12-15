/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KlineResponseK type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineResponseK{}

// KlineResponseK struct for KlineResponseK
type KlineResponseK struct {
	Smallt               *int64  `json:"t,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	S                    *string `json:"s,omitempty"`
	I                    *string `json:"i,omitempty"`
	F                    *int64  `json:"f,omitempty"`
	L                    *int64  `json:"L,omitempty"`
	O                    *string `json:"o,omitempty"`
	C                    *string `json:"c,omitempty"`
	H                    *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	N                    *int64  `json:"n,omitempty"`
	X                    *bool   `json:"x,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	V                    *string `json:"V,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	B                    *string `json:"B,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineResponseK KlineResponseK

// NewKlineResponseK instantiates a new KlineResponseK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineResponseK() *KlineResponseK {
	this := KlineResponseK{}
	return &this
}

// NewKlineResponseKWithDefaults instantiates a new KlineResponseK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineResponseKWithDefaults() *KlineResponseK {
	this := KlineResponseK{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *KlineResponseK) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *KlineResponseK) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *KlineResponseK) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *KlineResponseK) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *KlineResponseK) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *KlineResponseK) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *KlineResponseK) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *KlineResponseK) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *KlineResponseK) SetS(v string) {
	o.S = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *KlineResponseK) GetI() string {
	if o == nil || common.IsNil(o.I) {
		var ret string
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetIOk() (*string, bool) {
	if o == nil || common.IsNil(o.I) {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *KlineResponseK) HasI() bool {
	if o != nil && !common.IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *KlineResponseK) SetI(v string) {
	o.I = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *KlineResponseK) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *KlineResponseK) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *KlineResponseK) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *KlineResponseK) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *KlineResponseK) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *KlineResponseK) SetL(v int64) {
	o.L = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *KlineResponseK) GetO() string {
	if o == nil || common.IsNil(o.O) {
		var ret string
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetOOk() (*string, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *KlineResponseK) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *KlineResponseK) SetO(v string) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *KlineResponseK) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *KlineResponseK) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *KlineResponseK) SetC(v string) {
	o.C = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *KlineResponseK) GetH() string {
	if o == nil || common.IsNil(o.H) {
		var ret string
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetHOk() (*string, bool) {
	if o == nil || common.IsNil(o.H) {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *KlineResponseK) HasH() bool {
	if o != nil && !common.IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *KlineResponseK) SetH(v string) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *KlineResponseK) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *KlineResponseK) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *KlineResponseK) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *KlineResponseK) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *KlineResponseK) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *KlineResponseK) SetSmallv(v string) {
	o.Smallv = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *KlineResponseK) GetN() int64 {
	if o == nil || common.IsNil(o.N) {
		var ret int64
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetNOk() (*int64, bool) {
	if o == nil || common.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *KlineResponseK) HasN() bool {
	if o != nil && !common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *KlineResponseK) SetN(v int64) {
	o.N = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *KlineResponseK) GetX() bool {
	if o == nil || common.IsNil(o.X) {
		var ret bool
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetXOk() (*bool, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *KlineResponseK) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given bool and assigns it to the X field.
func (o *KlineResponseK) SetX(v bool) {
	o.X = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *KlineResponseK) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *KlineResponseK) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *KlineResponseK) SetSmallq(v string) {
	o.Smallq = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *KlineResponseK) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *KlineResponseK) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *KlineResponseK) SetV(v string) {
	o.V = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *KlineResponseK) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *KlineResponseK) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *KlineResponseK) SetQ(v string) {
	o.Q = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *KlineResponseK) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponseK) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *KlineResponseK) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *KlineResponseK) SetB(v string) {
	o.B = &v
}

func (o KlineResponseK) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineResponseK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.I) {
		toSerialize["i"] = o.I
	}
	if !common.IsNil(o.F) {
		toSerialize["f"] = o.F
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.O) {
		toSerialize["o"] = o.O
	}
	if !common.IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	if !common.IsNil(o.H) {
		toSerialize["h"] = o.H
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !common.IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KlineResponseK) UnmarshalJSON(data []byte) (err error) {
	varKlineResponseK := _KlineResponseK{}

	err = json.Unmarshal(data, &varKlineResponseK)

	if err != nil {
		return err
	}

	*o = KlineResponseK(varKlineResponseK)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "t")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "i")
		delete(additionalProperties, "f")
		delete(additionalProperties, "L")
		delete(additionalProperties, "o")
		delete(additionalProperties, "c")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "n")
		delete(additionalProperties, "x")
		delete(additionalProperties, "q")
		delete(additionalProperties, "V")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "B")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKlineResponseK struct {
	value *KlineResponseK
	isSet bool
}

func (v NullableKlineResponseK) Get() *KlineResponseK {
	return v.value
}

func (v *NullableKlineResponseK) Set(val *KlineResponseK) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineResponseK) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineResponseK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineResponseK(val *KlineResponseK) *NullableKlineResponseK {
	return &NullableKlineResponseK{value: val, isSet: true}
}

func (v NullableKlineResponseK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineResponseK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
