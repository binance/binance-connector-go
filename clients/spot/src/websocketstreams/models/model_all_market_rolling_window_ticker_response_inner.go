/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllMarketRollingWindowTickerResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketRollingWindowTickerResponseInner{}

// AllMarketRollingWindowTickerResponseInner struct for AllMarketRollingWindowTickerResponseInner
type AllMarketRollingWindowTickerResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	S                    *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	P                    *string `json:"P,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	H                    *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	W                    *string `json:"w,omitempty"`
	V                    *string `json:"v,omitempty"`
	Q                    *string `json:"q,omitempty"`
	O                    *int64  `json:"O,omitempty"`
	C                    *int64  `json:"C,omitempty"`
	F                    *int64  `json:"F,omitempty"`
	L                    *int64  `json:"L,omitempty"`
	N                    *int64  `json:"n,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMarketRollingWindowTickerResponseInner AllMarketRollingWindowTickerResponseInner

// NewAllMarketRollingWindowTickerResponseInner instantiates a new AllMarketRollingWindowTickerResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketRollingWindowTickerResponseInner() *AllMarketRollingWindowTickerResponseInner {
	this := AllMarketRollingWindowTickerResponseInner{}
	return &this
}

// NewAllMarketRollingWindowTickerResponseInnerWithDefaults instantiates a new AllMarketRollingWindowTickerResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketRollingWindowTickerResponseInnerWithDefaults() *AllMarketRollingWindowTickerResponseInner {
	this := AllMarketRollingWindowTickerResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllMarketRollingWindowTickerResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllMarketRollingWindowTickerResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMarketRollingWindowTickerResponseInner) SetS(v string) {
	o.S = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllMarketRollingWindowTickerResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllMarketRollingWindowTickerResponseInner) SetP(v string) {
	o.P = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AllMarketRollingWindowTickerResponseInner) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetH() string {
	if o == nil || common.IsNil(o.H) {
		var ret string
		return ret
	}
	return *o.H
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetHOk() (*string, bool) {
	if o == nil || common.IsNil(o.H) {
		return nil, false
	}
	return o.H, true
}

// HasH returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasH() bool {
	if o != nil && !common.IsNil(o.H) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *AllMarketRollingWindowTickerResponseInner) SetH(v string) {
	o.H = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *AllMarketRollingWindowTickerResponseInner) SetSmalll(v string) {
	o.Smalll = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *AllMarketRollingWindowTickerResponseInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetW() string {
	if o == nil || common.IsNil(o.W) {
		var ret string
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetWOk() (*string, bool) {
	if o == nil || common.IsNil(o.W) {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasW() bool {
	if o != nil && !common.IsNil(o.W) {
		return true
	}

	return false
}

// SetW gets a reference to the given string and assigns it to the W field.
func (o *AllMarketRollingWindowTickerResponseInner) SetW(v string) {
	o.W = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AllMarketRollingWindowTickerResponseInner) SetV(v string) {
	o.V = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMarketRollingWindowTickerResponseInner) SetQ(v string) {
	o.Q = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *AllMarketRollingWindowTickerResponseInner) SetO(v int64) {
	o.O = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetC() int64 {
	if o == nil || common.IsNil(o.C) {
		var ret int64
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetCOk() (*int64, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given int64 and assigns it to the C field.
func (o *AllMarketRollingWindowTickerResponseInner) SetC(v int64) {
	o.C = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetF() int64 {
	if o == nil || common.IsNil(o.F) {
		var ret int64
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetFOk() (*int64, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given int64 and assigns it to the F field.
func (o *AllMarketRollingWindowTickerResponseInner) SetF(v int64) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetL() int64 {
	if o == nil || common.IsNil(o.L) {
		var ret int64
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetLOk() (*int64, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given int64 and assigns it to the L field.
func (o *AllMarketRollingWindowTickerResponseInner) SetL(v int64) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *AllMarketRollingWindowTickerResponseInner) GetN() int64 {
	if o == nil || common.IsNil(o.N) {
		var ret int64
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketRollingWindowTickerResponseInner) GetNOk() (*int64, bool) {
	if o == nil || common.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *AllMarketRollingWindowTickerResponseInner) HasN() bool {
	if o != nil && !common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *AllMarketRollingWindowTickerResponseInner) SetN(v int64) {
	o.N = &v
}

func (o AllMarketRollingWindowTickerResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketRollingWindowTickerResponseInner) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.H) {
		toSerialize["h"] = o.H
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.W) {
		toSerialize["w"] = o.W
	}
	if !common.IsNil(o.V) {
		toSerialize["v"] = o.V
	}
	if !common.IsNil(o.Q) {
		toSerialize["q"] = o.Q
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
	if !common.IsNil(o.N) {
		toSerialize["n"] = o.N
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllMarketRollingWindowTickerResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAllMarketRollingWindowTickerResponseInner := _AllMarketRollingWindowTickerResponseInner{}

	err = json.Unmarshal(data, &varAllMarketRollingWindowTickerResponseInner)

	if err != nil {
		return err
	}

	*o = AllMarketRollingWindowTickerResponseInner(varAllMarketRollingWindowTickerResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "P")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "c")
		delete(additionalProperties, "w")
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

type NullableAllMarketRollingWindowTickerResponseInner struct {
	value *AllMarketRollingWindowTickerResponseInner
	isSet bool
}

func (v NullableAllMarketRollingWindowTickerResponseInner) Get() *AllMarketRollingWindowTickerResponseInner {
	return v.value
}

func (v *NullableAllMarketRollingWindowTickerResponseInner) Set(val *AllMarketRollingWindowTickerResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketRollingWindowTickerResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketRollingWindowTickerResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMarketRollingWindowTickerResponseInner(val *AllMarketRollingWindowTickerResponseInner) *NullableAllMarketRollingWindowTickerResponseInner {
	return &NullableAllMarketRollingWindowTickerResponseInner{value: val, isSet: true}
}

func (v NullableAllMarketRollingWindowTickerResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketRollingWindowTickerResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
