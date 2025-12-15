/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner{}

// Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner struct for Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner
type Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallh               *string `json:"h,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	V                    *string `json:"V,omitempty"`
	A                    *string `json:"A,omitempty"`
	P                    *string `json:"P,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	F                    *string `json:"F,omitempty"`
	L                    *string `json:"L,omitempty"`
	Smalln               *int64  `json:"n,omitempty"`
	Smallbo              *string `json:"bo,omitempty"`
	Smallao              *string `json:"ao,omitempty"`
	Smallbq              *string `json:"bq,omitempty"`
	Smallaq              *string `json:"aq,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	Smallt               *string `json:"t,omitempty"`
	Smallg               *string `json:"g,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	Smallvo              *string `json:"vo,omitempty"`
	Smallmp              *string `json:"mp,omitempty"`
	Smallhl              *string `json:"hl,omitempty"`
	Smallll              *string `json:"ll,omitempty"`
	Eep                  *string `json:"eep,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner

// NewTicker24HourByUnderlyingAssetAndExpirationDataResponseInner instantiates a new Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTicker24HourByUnderlyingAssetAndExpirationDataResponseInner() *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner {
	this := Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner{}
	return &this
}

// NewTicker24HourByUnderlyingAssetAndExpirationDataResponseInnerWithDefaults instantiates a new Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTicker24HourByUnderlyingAssetAndExpirationDataResponseInnerWithDefaults() *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner {
	this := Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalll(v string) {
	o.Smalll = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetV(v string) {
	o.V = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetA(v string) {
	o.A = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetP(v string) {
	o.P = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetQ(v string) {
	o.Q = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetF() string {
	if o == nil || common.IsNil(o.F) {
		var ret string
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetFOk() (*string, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetF(v string) {
	o.F = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetL(v string) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalln() int64 {
	if o == nil || common.IsNil(o.Smalln) {
		var ret int64
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallnOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given int64 and assigns it to the N field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalln(v int64) {
	o.Smalln = &v
}

// GetBo returns the Bo field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallbo() string {
	if o == nil || common.IsNil(o.Smallbo) {
		var ret string
		return ret
	}
	return *o.Smallbo
}

// GetBoOk returns a tuple with the Bo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallboOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbo) {
		return nil, false
	}
	return o.Smallbo, true
}

// HasBo returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallbo() bool {
	if o != nil && !common.IsNil(o.Smallbo) {
		return true
	}

	return false
}

// SetBo gets a reference to the given string and assigns it to the Bo field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallbo(v string) {
	o.Smallbo = &v
}

// GetAo returns the Ao field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallao() string {
	if o == nil || common.IsNil(o.Smallao) {
		var ret string
		return ret
	}
	return *o.Smallao
}

// GetAoOk returns a tuple with the Ao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallaoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallao) {
		return nil, false
	}
	return o.Smallao, true
}

// HasAo returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallao() bool {
	if o != nil && !common.IsNil(o.Smallao) {
		return true
	}

	return false
}

// SetAo gets a reference to the given string and assigns it to the Ao field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallao(v string) {
	o.Smallao = &v
}

// GetBq returns the Bq field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallbq() string {
	if o == nil || common.IsNil(o.Smallbq) {
		var ret string
		return ret
	}
	return *o.Smallbq
}

// GetBqOk returns a tuple with the Bq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallbqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbq) {
		return nil, false
	}
	return o.Smallbq, true
}

// HasBq returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallbq() bool {
	if o != nil && !common.IsNil(o.Smallbq) {
		return true
	}

	return false
}

// SetBq gets a reference to the given string and assigns it to the Bq field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallbq(v string) {
	o.Smallbq = &v
}

// GetAq returns the Aq field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallaq() string {
	if o == nil || common.IsNil(o.Smallaq) {
		var ret string
		return ret
	}
	return *o.Smallaq
}

// GetAqOk returns a tuple with the Aq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallaqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallaq) {
		return nil, false
	}
	return o.Smallaq, true
}

// HasAq returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallaq() bool {
	if o != nil && !common.IsNil(o.Smallaq) {
		return true
	}

	return false
}

// SetAq gets a reference to the given string and assigns it to the Aq field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallaq(v string) {
	o.Smallaq = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmalld(v string) {
	o.Smalld = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallt(v string) {
	o.Smallt = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallg() string {
	if o == nil || common.IsNil(o.Smallg) {
		var ret string
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallg(v string) {
	o.Smallg = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallv(v string) {
	o.Smallv = &v
}

// GetVo returns the Vo field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallvo() string {
	if o == nil || common.IsNil(o.Smallvo) {
		var ret string
		return ret
	}
	return *o.Smallvo
}

// GetVoOk returns a tuple with the Vo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallvoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallvo) {
		return nil, false
	}
	return o.Smallvo, true
}

// HasVo returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallvo() bool {
	if o != nil && !common.IsNil(o.Smallvo) {
		return true
	}

	return false
}

// SetVo gets a reference to the given string and assigns it to the Vo field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallvo(v string) {
	o.Smallvo = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallmp(v string) {
	o.Smallmp = &v
}

// GetHl returns the Hl field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallhl() string {
	if o == nil || common.IsNil(o.Smallhl) {
		var ret string
		return ret
	}
	return *o.Smallhl
}

// GetHlOk returns a tuple with the Hl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallhlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallhl) {
		return nil, false
	}
	return o.Smallhl, true
}

// HasHl returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallhl() bool {
	if o != nil && !common.IsNil(o.Smallhl) {
		return true
	}

	return false
}

// SetHl gets a reference to the given string and assigns it to the Hl field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallhl(v string) {
	o.Smallhl = &v
}

// GetLl returns the Ll field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallll() string {
	if o == nil || common.IsNil(o.Smallll) {
		var ret string
		return ret
	}
	return *o.Smallll
}

// GetLlOk returns a tuple with the Ll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetSmallllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallll) {
		return nil, false
	}
	return o.Smallll, true
}

// HasLl returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasSmallll() bool {
	if o != nil && !common.IsNil(o.Smallll) {
		return true
	}

	return false
}

// SetLl gets a reference to the given string and assigns it to the Ll field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetSmallll(v string) {
	o.Smallll = &v
}

// GetEep returns the Eep field value if set, zero value otherwise.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetEep() string {
	if o == nil || common.IsNil(o.Eep) {
		var ret string
		return ret
	}
	return *o.Eep
}

// GetEepOk returns a tuple with the Eep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) GetEepOk() (*string, bool) {
	if o == nil || common.IsNil(o.Eep) {
		return nil, false
	}
	return o.Eep, true
}

// HasEep returns a boolean if a field has been set.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) HasEep() bool {
	if o != nil && !common.IsNil(o.Eep) {
		return true
	}

	return false
}

// SetEep gets a reference to the given string and assigns it to the Eep field.
func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) SetEep(v string) {
	o.Eep = &v
}

func (o Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
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
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.A) {
		toSerialize["A"] = o.A
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
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
	if !common.IsNil(o.Smallbo) {
		toSerialize["bo"] = o.Smallbo
	}
	if !common.IsNil(o.Smallao) {
		toSerialize["ao"] = o.Smallao
	}
	if !common.IsNil(o.Smallbq) {
		toSerialize["bq"] = o.Smallbq
	}
	if !common.IsNil(o.Smallaq) {
		toSerialize["aq"] = o.Smallaq
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallg) {
		toSerialize["g"] = o.Smallg
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallvo) {
		toSerialize["vo"] = o.Smallvo
	}
	if !common.IsNil(o.Smallmp) {
		toSerialize["mp"] = o.Smallmp
	}
	if !common.IsNil(o.Smallhl) {
		toSerialize["hl"] = o.Smallhl
	}
	if !common.IsNil(o.Smallll) {
		toSerialize["ll"] = o.Smallll
	}
	if !common.IsNil(o.Eep) {
		toSerialize["eep"] = o.Eep
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) UnmarshalJSON(data []byte) (err error) {
	varTicker24HourByUnderlyingAssetAndExpirationDataResponseInner := _Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner{}

	err = json.Unmarshal(data, &varTicker24HourByUnderlyingAssetAndExpirationDataResponseInner)

	if err != nil {
		return err
	}

	*o = Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner(varTicker24HourByUnderlyingAssetAndExpirationDataResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "c")
		delete(additionalProperties, "V")
		delete(additionalProperties, "A")
		delete(additionalProperties, "P")
		delete(additionalProperties, "p")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "F")
		delete(additionalProperties, "L")
		delete(additionalProperties, "n")
		delete(additionalProperties, "bo")
		delete(additionalProperties, "ao")
		delete(additionalProperties, "bq")
		delete(additionalProperties, "aq")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		delete(additionalProperties, "d")
		delete(additionalProperties, "t")
		delete(additionalProperties, "g")
		delete(additionalProperties, "v")
		delete(additionalProperties, "vo")
		delete(additionalProperties, "mp")
		delete(additionalProperties, "hl")
		delete(additionalProperties, "ll")
		delete(additionalProperties, "eep")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner struct {
	value *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner
	isSet bool
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) Get() *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner {
	return v.value
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) Set(val *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner(val *Ticker24HourByUnderlyingAssetAndExpirationDataResponseInner) *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner {
	return &NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner{value: val, isSet: true}
}

func (v NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTicker24HourByUnderlyingAssetAndExpirationDataResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
