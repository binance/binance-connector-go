/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Executionreport type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Executionreport{}

// Executionreport struct for Executionreport
type Executionreport struct {
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	S                    *string `json:"S,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	P                    *string `json:"P,omitempty"`
	F                    *string `json:"F,omitempty"`
	Smallg               *int64  `json:"g,omitempty"`
	C                    *string `json:"C,omitempty"`
	Smallx               *string `json:"x,omitempty"`
	X                    *string `json:"X,omitempty"`
	Smallr               *string `json:"r,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallz               *string `json:"z,omitempty"`
	L                    *string `json:"L,omitempty"`
	Smalln               *string `json:"n,omitempty"`
	N                    *string `json:"N,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallt               *int64  `json:"t,omitempty"`
	I                    *int64  `json:"I,omitempty"`
	Smallw               *bool   `json:"w,omitempty"`
	Smallm               *bool   `json:"m,omitempty"`
	M                    *bool   `json:"M,omitempty"`
	O                    *int64  `json:"O,omitempty"`
	Z                    *string `json:"Z,omitempty"`
	Y                    *string `json:"Y,omitempty"`
	Q                    *string `json:"Q,omitempty"`
	W                    *int64  `json:"W,omitempty"`
	V                    *string `json:"V,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	D                    *string `json:"D,omitempty"`
	Smallj               *string `json:"j,omitempty"`
	J                    *string `json:"J,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	A                    *string `json:"A,omitempty"`
	B                    *string `json:"B,omitempty"`
	Smallu               *string `json:"u,omitempty"`
	U                    *string `json:"U,omitempty"`
	Cs                   *string `json:"Cs,omitempty"`
	Smallpl              *string `json:"pl,omitempty"`
	PL                   *string `json:"pL,omitempty"`
	PY                   *string `json:"pY,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smallk               *string `json:"k,omitempty"`
	US                   *bool   `json:"uS,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Executionreport Executionreport

// NewExecutionreport instantiates a new Executionreport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionreport() *Executionreport {
	this := Executionreport{}
	return &this
}

// NewExecutionreportWithDefaults instantiates a new Executionreport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionreportWithDefaults() *Executionreport {
	this := Executionreport{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Executionreport) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Executionreport) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Executionreport) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Executionreport) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *Executionreport) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Executionreport) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *Executionreport) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *Executionreport) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *Executionreport) SetSmallc(v string) {
	o.Smallc = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Executionreport) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *Executionreport) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Executionreport) SetS(v string) {
	o.S = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Executionreport) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *Executionreport) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *Executionreport) SetSmallo(v string) {
	o.Smallo = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *Executionreport) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *Executionreport) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *Executionreport) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Executionreport) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *Executionreport) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Executionreport) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Executionreport) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *Executionreport) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Executionreport) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Executionreport) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *Executionreport) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Executionreport) SetP(v string) {
	o.P = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *Executionreport) GetF() string {
	if o == nil || common.IsNil(o.F) {
		var ret string
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetFOk() (*string, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *Executionreport) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *Executionreport) SetF(v string) {
	o.F = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *Executionreport) GetSmallg() int64 {
	if o == nil || common.IsNil(o.Smallg) {
		var ret int64
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallgOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *Executionreport) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given int64 and assigns it to the G field.
func (o *Executionreport) SetSmallg(v int64) {
	o.Smallg = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *Executionreport) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *Executionreport) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *Executionreport) SetC(v string) {
	o.C = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *Executionreport) GetSmallx() string {
	if o == nil || common.IsNil(o.Smallx) {
		var ret string
		return ret
	}
	return *o.Smallx
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallxOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallx) {
		return nil, false
	}
	return o.Smallx, true
}

// HasX returns a boolean if a field has been set.
func (o *Executionreport) HasSmallx() bool {
	if o != nil && !common.IsNil(o.Smallx) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *Executionreport) SetSmallx(v string) {
	o.Smallx = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *Executionreport) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *Executionreport) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *Executionreport) SetX(v string) {
	o.X = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *Executionreport) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *Executionreport) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *Executionreport) SetSmallr(v string) {
	o.Smallr = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *Executionreport) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *Executionreport) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *Executionreport) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Executionreport) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *Executionreport) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Executionreport) SetSmalll(v string) {
	o.Smalll = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *Executionreport) GetSmallz() string {
	if o == nil || common.IsNil(o.Smallz) {
		var ret string
		return ret
	}
	return *o.Smallz
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallzOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallz) {
		return nil, false
	}
	return o.Smallz, true
}

// HasZ returns a boolean if a field has been set.
func (o *Executionreport) HasSmallz() bool {
	if o != nil && !common.IsNil(o.Smallz) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *Executionreport) SetSmallz(v string) {
	o.Smallz = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Executionreport) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *Executionreport) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Executionreport) SetL(v string) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *Executionreport) GetSmalln() string {
	if o == nil || common.IsNil(o.Smalln) {
		var ret string
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *Executionreport) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *Executionreport) SetSmalln(v string) {
	o.Smalln = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *Executionreport) GetN() string {
	if o == nil || common.IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetNOk() (*string, bool) {
	if o == nil || common.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *Executionreport) HasN() bool {
	if o != nil && !common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *Executionreport) SetN(v string) {
	o.N = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Executionreport) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *Executionreport) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *Executionreport) SetT(v int64) {
	o.T = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Executionreport) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *Executionreport) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *Executionreport) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *Executionreport) GetI() int64 {
	if o == nil || common.IsNil(o.I) {
		var ret int64
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetIOk() (*int64, bool) {
	if o == nil || common.IsNil(o.I) {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *Executionreport) HasI() bool {
	if o != nil && !common.IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *Executionreport) SetI(v int64) {
	o.I = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *Executionreport) GetSmallw() bool {
	if o == nil || common.IsNil(o.Smallw) {
		var ret bool
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallwOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *Executionreport) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given bool and assigns it to the W field.
func (o *Executionreport) SetSmallw(v bool) {
	o.Smallw = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *Executionreport) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *Executionreport) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *Executionreport) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *Executionreport) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *Executionreport) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *Executionreport) SetM(v bool) {
	o.M = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *Executionreport) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *Executionreport) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *Executionreport) SetO(v int64) {
	o.O = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *Executionreport) GetZ() string {
	if o == nil || common.IsNil(o.Z) {
		var ret string
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetZOk() (*string, bool) {
	if o == nil || common.IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *Executionreport) HasZ() bool {
	if o != nil && !common.IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *Executionreport) SetZ(v string) {
	o.Z = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *Executionreport) GetY() string {
	if o == nil || common.IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetYOk() (*string, bool) {
	if o == nil || common.IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *Executionreport) HasY() bool {
	if o != nil && !common.IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *Executionreport) SetY(v string) {
	o.Y = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *Executionreport) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *Executionreport) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *Executionreport) SetQ(v string) {
	o.Q = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *Executionreport) GetW() int64 {
	if o == nil || common.IsNil(o.W) {
		var ret int64
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetWOk() (*int64, bool) {
	if o == nil || common.IsNil(o.W) {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *Executionreport) HasW() bool {
	if o != nil && !common.IsNil(o.W) {
		return true
	}

	return false
}

// SetW gets a reference to the given int64 and assigns it to the W field.
func (o *Executionreport) SetW(v int64) {
	o.W = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Executionreport) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *Executionreport) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Executionreport) SetV(v string) {
	o.V = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *Executionreport) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *Executionreport) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *Executionreport) SetSmalld(v string) {
	o.Smalld = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *Executionreport) GetD() string {
	if o == nil || common.IsNil(o.D) {
		var ret string
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetDOk() (*string, bool) {
	if o == nil || common.IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *Executionreport) HasD() bool {
	if o != nil && !common.IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *Executionreport) SetD(v string) {
	o.D = &v
}

// GetJ returns the J field value if set, zero value otherwise.
func (o *Executionreport) GetSmallj() string {
	if o == nil || common.IsNil(o.Smallj) {
		var ret string
		return ret
	}
	return *o.Smallj
}

// GetJOk returns a tuple with the J field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalljOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallj) {
		return nil, false
	}
	return o.Smallj, true
}

// HasJ returns a boolean if a field has been set.
func (o *Executionreport) HasSmallj() bool {
	if o != nil && !common.IsNil(o.Smallj) {
		return true
	}

	return false
}

// SetJ gets a reference to the given string and assigns it to the J field.
func (o *Executionreport) SetSmallj(v string) {
	o.Smallj = &v
}

// GetJ returns the J field value if set, zero value otherwise.
func (o *Executionreport) GetJ() string {
	if o == nil || common.IsNil(o.J) {
		var ret string
		return ret
	}
	return *o.J
}

// GetJOk returns a tuple with the J field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetJOk() (*string, bool) {
	if o == nil || common.IsNil(o.J) {
		return nil, false
	}
	return o.J, true
}

// HasJ returns a boolean if a field has been set.
func (o *Executionreport) HasJ() bool {
	if o != nil && !common.IsNil(o.J) {
		return true
	}

	return false
}

// SetJ gets a reference to the given string and assigns it to the J field.
func (o *Executionreport) SetJ(v string) {
	o.J = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *Executionreport) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *Executionreport) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *Executionreport) SetSmallv(v string) {
	o.Smallv = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Executionreport) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *Executionreport) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Executionreport) SetA(v string) {
	o.A = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *Executionreport) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *Executionreport) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *Executionreport) SetB(v string) {
	o.B = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *Executionreport) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *Executionreport) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *Executionreport) SetSmallu(v string) {
	o.Smallu = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *Executionreport) GetU() string {
	if o == nil || common.IsNil(o.U) {
		var ret string
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetUOk() (*string, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *Executionreport) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *Executionreport) SetU(v string) {
	o.U = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *Executionreport) GetCs() string {
	if o == nil || common.IsNil(o.Cs) {
		var ret string
		return ret
	}
	return *o.Cs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetCsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cs) {
		return nil, false
	}
	return o.Cs, true
}

// HasCs returns a boolean if a field has been set.
func (o *Executionreport) HasCs() bool {
	if o != nil && !common.IsNil(o.Cs) {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *Executionreport) SetCs(v string) {
	o.Cs = &v
}

// GetPl returns the Pl field value if set, zero value otherwise.
func (o *Executionreport) GetSmallpl() string {
	if o == nil || common.IsNil(o.Smallpl) {
		var ret string
		return ret
	}
	return *o.Smallpl
}

// GetPlOk returns a tuple with the Pl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallplOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpl) {
		return nil, false
	}
	return o.Smallpl, true
}

// HasPl returns a boolean if a field has been set.
func (o *Executionreport) HasSmallpl() bool {
	if o != nil && !common.IsNil(o.Smallpl) {
		return true
	}

	return false
}

// SetPl gets a reference to the given string and assigns it to the Pl field.
func (o *Executionreport) SetSmallpl(v string) {
	o.Smallpl = &v
}

// GetPL returns the PL field value if set, zero value otherwise.
func (o *Executionreport) GetPL() string {
	if o == nil || common.IsNil(o.PL) {
		var ret string
		return ret
	}
	return *o.PL
}

// GetPLOk returns a tuple with the PL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetPLOk() (*string, bool) {
	if o == nil || common.IsNil(o.PL) {
		return nil, false
	}
	return o.PL, true
}

// HasPL returns a boolean if a field has been set.
func (o *Executionreport) HasPL() bool {
	if o != nil && !common.IsNil(o.PL) {
		return true
	}

	return false
}

// SetPL gets a reference to the given string and assigns it to the PL field.
func (o *Executionreport) SetPL(v string) {
	o.PL = &v
}

// GetPY returns the PY field value if set, zero value otherwise.
func (o *Executionreport) GetPY() string {
	if o == nil || common.IsNil(o.PY) {
		var ret string
		return ret
	}
	return *o.PY
}

// GetPYOk returns a tuple with the PY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetPYOk() (*string, bool) {
	if o == nil || common.IsNil(o.PY) {
		return nil, false
	}
	return o.PY, true
}

// HasPY returns a boolean if a field has been set.
func (o *Executionreport) HasPY() bool {
	if o != nil && !common.IsNil(o.PY) {
		return true
	}

	return false
}

// SetPY gets a reference to the given string and assigns it to the PY field.
func (o *Executionreport) SetPY(v string) {
	o.PY = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *Executionreport) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *Executionreport) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *Executionreport) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Executionreport) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *Executionreport) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Executionreport) SetSmalla(v string) {
	o.Smalla = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *Executionreport) GetSmallk() string {
	if o == nil || common.IsNil(o.Smallk) {
		var ret string
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetSmallkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *Executionreport) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given string and assigns it to the K field.
func (o *Executionreport) SetSmallk(v string) {
	o.Smallk = &v
}

// GetUS returns the US field value if set, zero value otherwise.
func (o *Executionreport) GetUS() bool {
	if o == nil || common.IsNil(o.US) {
		var ret bool
		return ret
	}
	return *o.US
}

// GetUSOk returns a tuple with the US field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Executionreport) GetUSOk() (*bool, bool) {
	if o == nil || common.IsNil(o.US) {
		return nil, false
	}
	return o.US, true
}

// HasUS returns a boolean if a field has been set.
func (o *Executionreport) HasUS() bool {
	if o != nil && !common.IsNil(o.US) {
		return true
	}

	return false
}

// SetUS gets a reference to the given bool and assigns it to the US field.
func (o *Executionreport) SetUS(v bool) {
	o.US = &v
}

func (o Executionreport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Executionreport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.F) {
		toSerialize["F"] = o.F
	}
	if !common.IsNil(o.Smallg) {
		toSerialize["g"] = o.Smallg
	}
	if !common.IsNil(o.C) {
		toSerialize["C"] = o.C
	}
	if !common.IsNil(o.Smallx) {
		toSerialize["x"] = o.Smallx
	}
	if !common.IsNil(o.X) {
		toSerialize["X"] = o.X
	}
	if !common.IsNil(o.Smallr) {
		toSerialize["r"] = o.Smallr
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallz) {
		toSerialize["z"] = o.Smallz
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}
	if !common.IsNil(o.N) {
		toSerialize["N"] = o.N
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.I) {
		toSerialize["I"] = o.I
	}
	if !common.IsNil(o.Smallw) {
		toSerialize["w"] = o.Smallw
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.M) {
		toSerialize["M"] = o.M
	}
	if !common.IsNil(o.O) {
		toSerialize["O"] = o.O
	}
	if !common.IsNil(o.Z) {
		toSerialize["Z"] = o.Z
	}
	if !common.IsNil(o.Y) {
		toSerialize["Y"] = o.Y
	}
	if !common.IsNil(o.Q) {
		toSerialize["Q"] = o.Q
	}
	if !common.IsNil(o.W) {
		toSerialize["W"] = o.W
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.Smalld) {
		toSerialize["d"] = o.Smalld
	}
	if !common.IsNil(o.D) {
		toSerialize["D"] = o.D
	}
	if !common.IsNil(o.Smallj) {
		toSerialize["j"] = o.Smallj
	}
	if !common.IsNil(o.J) {
		toSerialize["J"] = o.J
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.A) {
		toSerialize["A"] = o.A
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}
	if !common.IsNil(o.Cs) {
		toSerialize["Cs"] = o.Cs
	}
	if !common.IsNil(o.Smallpl) {
		toSerialize["pl"] = o.Smallpl
	}
	if !common.IsNil(o.PL) {
		toSerialize["pL"] = o.PL
	}
	if !common.IsNil(o.PY) {
		toSerialize["pY"] = o.PY
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallk) {
		toSerialize["k"] = o.Smallk
	}
	if !common.IsNil(o.US) {
		toSerialize["uS"] = o.US
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Executionreport) UnmarshalJSON(data []byte) (err error) {
	varExecutionreport := _Executionreport{}

	err = json.Unmarshal(data, &varExecutionreport)

	if err != nil {
		return err
	}

	*o = Executionreport(varExecutionreport)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "S")
		delete(additionalProperties, "o")
		delete(additionalProperties, "f")
		delete(additionalProperties, "q")
		delete(additionalProperties, "p")
		delete(additionalProperties, "P")
		delete(additionalProperties, "F")
		delete(additionalProperties, "g")
		delete(additionalProperties, "C")
		delete(additionalProperties, "x")
		delete(additionalProperties, "X")
		delete(additionalProperties, "r")
		delete(additionalProperties, "i")
		delete(additionalProperties, "l")
		delete(additionalProperties, "z")
		delete(additionalProperties, "L")
		delete(additionalProperties, "n")
		delete(additionalProperties, "N")
		delete(additionalProperties, "T")
		delete(additionalProperties, "t")
		delete(additionalProperties, "I")
		delete(additionalProperties, "w")
		delete(additionalProperties, "m")
		delete(additionalProperties, "M")
		delete(additionalProperties, "O")
		delete(additionalProperties, "Z")
		delete(additionalProperties, "Y")
		delete(additionalProperties, "Q")
		delete(additionalProperties, "W")
		delete(additionalProperties, "V")
		delete(additionalProperties, "d")
		delete(additionalProperties, "D")
		delete(additionalProperties, "j")
		delete(additionalProperties, "J")
		delete(additionalProperties, "v")
		delete(additionalProperties, "A")
		delete(additionalProperties, "B")
		delete(additionalProperties, "u")
		delete(additionalProperties, "U")
		delete(additionalProperties, "Cs")
		delete(additionalProperties, "pl")
		delete(additionalProperties, "pL")
		delete(additionalProperties, "pY")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		delete(additionalProperties, "k")
		delete(additionalProperties, "uS")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionreport struct {
	value *Executionreport
	isSet bool
}

func (v NullableExecutionreport) Get() *Executionreport {
	return v.value
}

func (v *NullableExecutionreport) Set(val *Executionreport) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionreport) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionreport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionreport(val *Executionreport) *NullableExecutionreport {
	return &NullableExecutionreport{value: val, isSet: true}
}

func (v NullableExecutionreport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionreport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
