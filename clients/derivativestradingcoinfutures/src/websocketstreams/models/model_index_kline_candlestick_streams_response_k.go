/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexKlineCandlestickStreamsResponseK type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexKlineCandlestickStreamsResponseK{}

// IndexKlineCandlestickStreamsResponseK struct for IndexKlineCandlestickStreamsResponseK
type IndexKlineCandlestickStreamsResponseK struct {
	Smallt               *int64  `json:"t,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	Smallf               *int64  `json:"f,omitempty"`
	L                    *int64  `json:"L,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Smallh               *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	Smalln               *int64  `json:"n,omitempty"`
	Smallx               *bool   `json:"x,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	V                    *string `json:"V,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	B                    *string `json:"B,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexKlineCandlestickStreamsResponseK IndexKlineCandlestickStreamsResponseK

// NewIndexKlineCandlestickStreamsResponseK instantiates a new IndexKlineCandlestickStreamsResponseK object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexKlineCandlestickStreamsResponseK() *IndexKlineCandlestickStreamsResponseK {
	this := IndexKlineCandlestickStreamsResponseK{}
	return &this
}

// NewIndexKlineCandlestickStreamsResponseKWithDefaults instantiates a new IndexKlineCandlestickStreamsResponseK object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexKlineCandlestickStreamsResponseKWithDefaults() *IndexKlineCandlestickStreamsResponseK {
	this := IndexKlineCandlestickStreamsResponseK{}
	return &this
}

// GetT returns the T field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *IndexKlineCandlestickStreamsResponseK) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmalls(v string) {
	o.Smalls = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmalli(v string) {
	o.Smalli = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallf() int64 {
	if o == nil || common.IsNil(o.Smallf) {
		var ret int64
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallfOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallf(v int64) {
	o.Smallf = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *IndexKlineCandlestickStreamsResponseK) SetL(v int64) {
	o.L = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallo(v string) {
	o.Smallo = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallc(v string) {
	o.Smallc = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallv(v string) {
	o.Smallv = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmalln() int64 {
	if o == nil || common.IsNil(o.Smalln) {
		var ret int64
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallnOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmalln(v int64) {
	o.Smalln = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallx() bool {
	if o == nil || common.IsNil(o.Smallx) {
		var ret bool
		return ret
	}
	return *o.Smallx
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallxOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallx) {
		return nil, false
	}
	return o.Smallx, true
}

// HasX returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallx() bool {
	if o != nil && !common.IsNil(o.Smallx) {
		return true
	}

	return false
}

// SetX gets a reference to the given bool and assigns it to the X field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallx(v bool) {
	o.Smallx = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *IndexKlineCandlestickStreamsResponseK) SetSmallq(v string) {
	o.Smallq = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *IndexKlineCandlestickStreamsResponseK) SetV(v string) {
	o.V = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *IndexKlineCandlestickStreamsResponseK) SetQ(v string) {
	o.Q = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponseK) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponseK) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponseK) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *IndexKlineCandlestickStreamsResponseK) SetB(v string) {
	o.B = &v
}

func (o IndexKlineCandlestickStreamsResponseK) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexKlineCandlestickStreamsResponseK) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
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
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}
	if !common.IsNil(o.Smallx) {
		toSerialize["x"] = o.Smallx
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

func (o *IndexKlineCandlestickStreamsResponseK) UnmarshalJSON(data []byte) (err error) {
	varIndexKlineCandlestickStreamsResponseK := _IndexKlineCandlestickStreamsResponseK{}

	err = json.Unmarshal(data, &varIndexKlineCandlestickStreamsResponseK)

	if err != nil {
		return err
	}

	*o = IndexKlineCandlestickStreamsResponseK(varIndexKlineCandlestickStreamsResponseK)

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

type NullableIndexKlineCandlestickStreamsResponseK struct {
	value *IndexKlineCandlestickStreamsResponseK
	isSet bool
}

func (v NullableIndexKlineCandlestickStreamsResponseK) Get() *IndexKlineCandlestickStreamsResponseK {
	return v.value
}

func (v *NullableIndexKlineCandlestickStreamsResponseK) Set(val *IndexKlineCandlestickStreamsResponseK) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexKlineCandlestickStreamsResponseK) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexKlineCandlestickStreamsResponseK) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexKlineCandlestickStreamsResponseK(val *IndexKlineCandlestickStreamsResponseK) *NullableIndexKlineCandlestickStreamsResponseK {
	return &NullableIndexKlineCandlestickStreamsResponseK{value: val, isSet: true}
}

func (v NullableIndexKlineCandlestickStreamsResponseK) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexKlineCandlestickStreamsResponseK) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
