/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionReport type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionReport{}

// ExecutionReport struct for ExecutionReport
type ExecutionReport struct {
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
	Smallv               *int64  `json:"v,omitempty"`
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
	Smalld               *int64  `json:"d,omitempty"`
	D                    *int64  `json:"D,omitempty"`
	Smallj               *int64  `json:"j,omitempty"`
	J                    *int64  `json:"J,omitempty"`
	A                    *string `json:"A,omitempty"`
	B                    *string `json:"B,omitempty"`
	Smallu               *int64  `json:"u,omitempty"`
	U                    *int64  `json:"U,omitempty"`
	Cs                   *string `json:"Cs,omitempty"`
	Smallpl              *string `json:"pl,omitempty"`
	PL                   *string `json:"pL,omitempty"`
	PY                   *string `json:"pY,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *int64  `json:"a,omitempty"`
	Smallk               *string `json:"k,omitempty"`
	US                   *bool   `json:"uS,omitempty"`
	GP                   *string `json:"gP,omitempty"`
	GOT                  *string `json:"gOT,omitempty"`
	GOV                  *int64  `json:"gOV,omitempty"`
	Smallgp              *string `json:"gp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionReport ExecutionReport

// NewExecutionReport instantiates a new ExecutionReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionReport() *ExecutionReport {
	this := ExecutionReport{}
	return &this
}

// NewExecutionReportWithDefaults instantiates a new ExecutionReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionReportWithDefaults() *ExecutionReport {
	this := ExecutionReport{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *ExecutionReport) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *ExecutionReport) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *ExecutionReport) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ExecutionReport) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ExecutionReport) SetSmallc(v string) {
	o.Smallc = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ExecutionReport) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *ExecutionReport) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ExecutionReport) SetS(v string) {
	o.S = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *ExecutionReport) SetSmallo(v string) {
	o.Smallo = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *ExecutionReport) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ExecutionReport) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *ExecutionReport) SetSmallp(v string) {
	o.Smallp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *ExecutionReport) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *ExecutionReport) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *ExecutionReport) SetP(v string) {
	o.P = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *ExecutionReport) GetF() string {
	if o == nil || common.IsNil(o.F) {
		var ret string
		return ret
	}
	return *o.F
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetFOk() (*string, bool) {
	if o == nil || common.IsNil(o.F) {
		return nil, false
	}
	return o.F, true
}

// HasF returns a boolean if a field has been set.
func (o *ExecutionReport) HasF() bool {
	if o != nil && !common.IsNil(o.F) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *ExecutionReport) SetF(v string) {
	o.F = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallg() int64 {
	if o == nil || common.IsNil(o.Smallg) {
		var ret int64
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallgOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given int64 and assigns it to the G field.
func (o *ExecutionReport) SetSmallg(v int64) {
	o.Smallg = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ExecutionReport) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ExecutionReport) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ExecutionReport) SetC(v string) {
	o.C = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallx() string {
	if o == nil || common.IsNil(o.Smallx) {
		var ret string
		return ret
	}
	return *o.Smallx
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallxOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallx) {
		return nil, false
	}
	return o.Smallx, true
}

// HasX returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallx() bool {
	if o != nil && !common.IsNil(o.Smallx) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *ExecutionReport) SetSmallx(v string) {
	o.Smallx = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *ExecutionReport) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *ExecutionReport) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *ExecutionReport) SetX(v string) {
	o.X = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallr() string {
	if o == nil || common.IsNil(o.Smallr) {
		var ret string
		return ret
	}
	return *o.Smallr
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallr) {
		return nil, false
	}
	return o.Smallr, true
}

// HasR returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallr() bool {
	if o != nil && !common.IsNil(o.Smallr) {
		return true
	}

	return false
}

// SetR gets a reference to the given string and assigns it to the R field.
func (o *ExecutionReport) SetSmallr(v string) {
	o.Smallr = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ExecutionReport) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *ExecutionReport) SetSmalll(v string) {
	o.Smalll = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallz() string {
	if o == nil || common.IsNil(o.Smallz) {
		var ret string
		return ret
	}
	return *o.Smallz
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallzOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallz) {
		return nil, false
	}
	return o.Smallz, true
}

// HasZ returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallz() bool {
	if o != nil && !common.IsNil(o.Smallz) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *ExecutionReport) SetSmallz(v string) {
	o.Smallz = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *ExecutionReport) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *ExecutionReport) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *ExecutionReport) SetL(v string) {
	o.L = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalln() string {
	if o == nil || common.IsNil(o.Smalln) {
		var ret string
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *ExecutionReport) SetSmalln(v string) {
	o.Smalln = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *ExecutionReport) GetN() string {
	if o == nil || common.IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetNOk() (*string, bool) {
	if o == nil || common.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *ExecutionReport) HasN() bool {
	if o != nil && !common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *ExecutionReport) SetN(v string) {
	o.N = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ExecutionReport) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ExecutionReport) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ExecutionReport) SetT(v int64) {
	o.T = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ExecutionReport) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallv() int64 {
	if o == nil || common.IsNil(o.Smallv) {
		var ret int64
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallvOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given int64 and assigns it to the V field.
func (o *ExecutionReport) SetSmallv(v int64) {
	o.Smallv = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ExecutionReport) GetI() int64 {
	if o == nil || common.IsNil(o.I) {
		var ret int64
		return ret
	}
	return *o.I
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetIOk() (*int64, bool) {
	if o == nil || common.IsNil(o.I) {
		return nil, false
	}
	return o.I, true
}

// HasI returns a boolean if a field has been set.
func (o *ExecutionReport) HasI() bool {
	if o != nil && !common.IsNil(o.I) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ExecutionReport) SetI(v int64) {
	o.I = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallw() bool {
	if o == nil || common.IsNil(o.Smallw) {
		var ret bool
		return ret
	}
	return *o.Smallw
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallwOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallw) {
		return nil, false
	}
	return o.Smallw, true
}

// HasW returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallw() bool {
	if o != nil && !common.IsNil(o.Smallw) {
		return true
	}

	return false
}

// SetW gets a reference to the given bool and assigns it to the W field.
func (o *ExecutionReport) SetSmallw(v bool) {
	o.Smallw = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *ExecutionReport) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *ExecutionReport) GetM() bool {
	if o == nil || common.IsNil(o.M) {
		var ret bool
		return ret
	}
	return *o.M
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetMOk() (*bool, bool) {
	if o == nil || common.IsNil(o.M) {
		return nil, false
	}
	return o.M, true
}

// HasM returns a boolean if a field has been set.
func (o *ExecutionReport) HasM() bool {
	if o != nil && !common.IsNil(o.M) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *ExecutionReport) SetM(v bool) {
	o.M = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *ExecutionReport) GetO() int64 {
	if o == nil || common.IsNil(o.O) {
		var ret int64
		return ret
	}
	return *o.O
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetOOk() (*int64, bool) {
	if o == nil || common.IsNil(o.O) {
		return nil, false
	}
	return o.O, true
}

// HasO returns a boolean if a field has been set.
func (o *ExecutionReport) HasO() bool {
	if o != nil && !common.IsNil(o.O) {
		return true
	}

	return false
}

// SetO gets a reference to the given int64 and assigns it to the O field.
func (o *ExecutionReport) SetO(v int64) {
	o.O = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *ExecutionReport) GetZ() string {
	if o == nil || common.IsNil(o.Z) {
		var ret string
		return ret
	}
	return *o.Z
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetZOk() (*string, bool) {
	if o == nil || common.IsNil(o.Z) {
		return nil, false
	}
	return o.Z, true
}

// HasZ returns a boolean if a field has been set.
func (o *ExecutionReport) HasZ() bool {
	if o != nil && !common.IsNil(o.Z) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *ExecutionReport) SetZ(v string) {
	o.Z = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *ExecutionReport) GetY() string {
	if o == nil || common.IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetYOk() (*string, bool) {
	if o == nil || common.IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *ExecutionReport) HasY() bool {
	if o != nil && !common.IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *ExecutionReport) SetY(v string) {
	o.Y = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ExecutionReport) GetQ() string {
	if o == nil || common.IsNil(o.Q) {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetQOk() (*string, bool) {
	if o == nil || common.IsNil(o.Q) {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *ExecutionReport) HasQ() bool {
	if o != nil && !common.IsNil(o.Q) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ExecutionReport) SetQ(v string) {
	o.Q = &v
}

// GetW returns the W field value if set, zero value otherwise.
func (o *ExecutionReport) GetW() int64 {
	if o == nil || common.IsNil(o.W) {
		var ret int64
		return ret
	}
	return *o.W
}

// GetWOk returns a tuple with the W field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetWOk() (*int64, bool) {
	if o == nil || common.IsNil(o.W) {
		return nil, false
	}
	return o.W, true
}

// HasW returns a boolean if a field has been set.
func (o *ExecutionReport) HasW() bool {
	if o != nil && !common.IsNil(o.W) {
		return true
	}

	return false
}

// SetW gets a reference to the given int64 and assigns it to the W field.
func (o *ExecutionReport) SetW(v int64) {
	o.W = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *ExecutionReport) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *ExecutionReport) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *ExecutionReport) SetV(v string) {
	o.V = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalld() int64 {
	if o == nil || common.IsNil(o.Smalld) {
		var ret int64
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalldOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given int64 and assigns it to the D field.
func (o *ExecutionReport) SetSmalld(v int64) {
	o.Smalld = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *ExecutionReport) GetD() int64 {
	if o == nil || common.IsNil(o.D) {
		var ret int64
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetDOk() (*int64, bool) {
	if o == nil || common.IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *ExecutionReport) HasD() bool {
	if o != nil && !common.IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given int64 and assigns it to the D field.
func (o *ExecutionReport) SetD(v int64) {
	o.D = &v
}

// GetJ returns the J field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallj() int64 {
	if o == nil || common.IsNil(o.Smallj) {
		var ret int64
		return ret
	}
	return *o.Smallj
}

// GetJOk returns a tuple with the J field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalljOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallj) {
		return nil, false
	}
	return o.Smallj, true
}

// HasJ returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallj() bool {
	if o != nil && !common.IsNil(o.Smallj) {
		return true
	}

	return false
}

// SetJ gets a reference to the given int64 and assigns it to the J field.
func (o *ExecutionReport) SetSmallj(v int64) {
	o.Smallj = &v
}

// GetJ returns the J field value if set, zero value otherwise.
func (o *ExecutionReport) GetJ() int64 {
	if o == nil || common.IsNil(o.J) {
		var ret int64
		return ret
	}
	return *o.J
}

// GetJOk returns a tuple with the J field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetJOk() (*int64, bool) {
	if o == nil || common.IsNil(o.J) {
		return nil, false
	}
	return o.J, true
}

// HasJ returns a boolean if a field has been set.
func (o *ExecutionReport) HasJ() bool {
	if o != nil && !common.IsNil(o.J) {
		return true
	}

	return false
}

// SetJ gets a reference to the given int64 and assigns it to the J field.
func (o *ExecutionReport) SetJ(v int64) {
	o.J = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *ExecutionReport) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *ExecutionReport) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *ExecutionReport) SetA(v string) {
	o.A = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *ExecutionReport) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *ExecutionReport) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *ExecutionReport) SetB(v string) {
	o.B = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *ExecutionReport) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *ExecutionReport) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *ExecutionReport) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *ExecutionReport) SetU(v int64) {
	o.U = &v
}

// GetCs returns the Cs field value if set, zero value otherwise.
func (o *ExecutionReport) GetCs() string {
	if o == nil || common.IsNil(o.Cs) {
		var ret string
		return ret
	}
	return *o.Cs
}

// GetCsOk returns a tuple with the Cs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetCsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Cs) {
		return nil, false
	}
	return o.Cs, true
}

// HasCs returns a boolean if a field has been set.
func (o *ExecutionReport) HasCs() bool {
	if o != nil && !common.IsNil(o.Cs) {
		return true
	}

	return false
}

// SetCs gets a reference to the given string and assigns it to the Cs field.
func (o *ExecutionReport) SetCs(v string) {
	o.Cs = &v
}

// GetPl returns the Pl field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallpl() string {
	if o == nil || common.IsNil(o.Smallpl) {
		var ret string
		return ret
	}
	return *o.Smallpl
}

// GetPlOk returns a tuple with the Pl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallplOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpl) {
		return nil, false
	}
	return o.Smallpl, true
}

// HasPl returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallpl() bool {
	if o != nil && !common.IsNil(o.Smallpl) {
		return true
	}

	return false
}

// SetPl gets a reference to the given string and assigns it to the Pl field.
func (o *ExecutionReport) SetSmallpl(v string) {
	o.Smallpl = &v
}

// GetPL returns the PL field value if set, zero value otherwise.
func (o *ExecutionReport) GetPL() string {
	if o == nil || common.IsNil(o.PL) {
		var ret string
		return ret
	}
	return *o.PL
}

// GetPLOk returns a tuple with the PL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetPLOk() (*string, bool) {
	if o == nil || common.IsNil(o.PL) {
		return nil, false
	}
	return o.PL, true
}

// HasPL returns a boolean if a field has been set.
func (o *ExecutionReport) HasPL() bool {
	if o != nil && !common.IsNil(o.PL) {
		return true
	}

	return false
}

// SetPL gets a reference to the given string and assigns it to the PL field.
func (o *ExecutionReport) SetPL(v string) {
	o.PL = &v
}

// GetPY returns the PY field value if set, zero value otherwise.
func (o *ExecutionReport) GetPY() string {
	if o == nil || common.IsNil(o.PY) {
		var ret string
		return ret
	}
	return *o.PY
}

// GetPYOk returns a tuple with the PY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetPYOk() (*string, bool) {
	if o == nil || common.IsNil(o.PY) {
		return nil, false
	}
	return o.PY, true
}

// HasPY returns a boolean if a field has been set.
func (o *ExecutionReport) HasPY() bool {
	if o != nil && !common.IsNil(o.PY) {
		return true
	}

	return false
}

// SetPY gets a reference to the given string and assigns it to the PY field.
func (o *ExecutionReport) SetPY(v string) {
	o.PY = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *ExecutionReport) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmalla() int64 {
	if o == nil || common.IsNil(o.Smalla) {
		var ret int64
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallaOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given int64 and assigns it to the A field.
func (o *ExecutionReport) SetSmalla(v int64) {
	o.Smalla = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallk() string {
	if o == nil || common.IsNil(o.Smallk) {
		var ret string
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallkOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given string and assigns it to the K field.
func (o *ExecutionReport) SetSmallk(v string) {
	o.Smallk = &v
}

// GetUS returns the US field value if set, zero value otherwise.
func (o *ExecutionReport) GetUS() bool {
	if o == nil || common.IsNil(o.US) {
		var ret bool
		return ret
	}
	return *o.US
}

// GetUSOk returns a tuple with the US field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetUSOk() (*bool, bool) {
	if o == nil || common.IsNil(o.US) {
		return nil, false
	}
	return o.US, true
}

// HasUS returns a boolean if a field has been set.
func (o *ExecutionReport) HasUS() bool {
	if o != nil && !common.IsNil(o.US) {
		return true
	}

	return false
}

// SetUS gets a reference to the given bool and assigns it to the US field.
func (o *ExecutionReport) SetUS(v bool) {
	o.US = &v
}

// GetGP returns the GP field value if set, zero value otherwise.
func (o *ExecutionReport) GetGP() string {
	if o == nil || common.IsNil(o.GP) {
		var ret string
		return ret
	}
	return *o.GP
}

// GetGPOk returns a tuple with the GP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetGPOk() (*string, bool) {
	if o == nil || common.IsNil(o.GP) {
		return nil, false
	}
	return o.GP, true
}

// HasGP returns a boolean if a field has been set.
func (o *ExecutionReport) HasGP() bool {
	if o != nil && !common.IsNil(o.GP) {
		return true
	}

	return false
}

// SetGP gets a reference to the given string and assigns it to the GP field.
func (o *ExecutionReport) SetGP(v string) {
	o.GP = &v
}

// GetGOT returns the GOT field value if set, zero value otherwise.
func (o *ExecutionReport) GetGOT() string {
	if o == nil || common.IsNil(o.GOT) {
		var ret string
		return ret
	}
	return *o.GOT
}

// GetGOTOk returns a tuple with the GOT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetGOTOk() (*string, bool) {
	if o == nil || common.IsNil(o.GOT) {
		return nil, false
	}
	return o.GOT, true
}

// HasGOT returns a boolean if a field has been set.
func (o *ExecutionReport) HasGOT() bool {
	if o != nil && !common.IsNil(o.GOT) {
		return true
	}

	return false
}

// SetGOT gets a reference to the given string and assigns it to the GOT field.
func (o *ExecutionReport) SetGOT(v string) {
	o.GOT = &v
}

// GetGOV returns the GOV field value if set, zero value otherwise.
func (o *ExecutionReport) GetGOV() int64 {
	if o == nil || common.IsNil(o.GOV) {
		var ret int64
		return ret
	}
	return *o.GOV
}

// GetGOVOk returns a tuple with the GOV field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetGOVOk() (*int64, bool) {
	if o == nil || common.IsNil(o.GOV) {
		return nil, false
	}
	return o.GOV, true
}

// HasGOV returns a boolean if a field has been set.
func (o *ExecutionReport) HasGOV() bool {
	if o != nil && !common.IsNil(o.GOV) {
		return true
	}

	return false
}

// SetGOV gets a reference to the given int64 and assigns it to the GOV field.
func (o *ExecutionReport) SetGOV(v int64) {
	o.GOV = &v
}

// GetGp returns the Gp field value if set, zero value otherwise.
func (o *ExecutionReport) GetSmallgp() string {
	if o == nil || common.IsNil(o.Smallgp) {
		var ret string
		return ret
	}
	return *o.Smallgp
}

// GetGpOk returns a tuple with the Gp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionReport) GetSmallgpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallgp) {
		return nil, false
	}
	return o.Smallgp, true
}

// HasGp returns a boolean if a field has been set.
func (o *ExecutionReport) HasSmallgp() bool {
	if o != nil && !common.IsNil(o.Smallgp) {
		return true
	}

	return false
}

// SetGp gets a reference to the given string and assigns it to the Gp field.
func (o *ExecutionReport) SetSmallgp(v string) {
	o.Smallgp = &v
}

func (o ExecutionReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionReport) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
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
	if !common.IsNil(o.GP) {
		toSerialize["gP"] = o.GP
	}
	if !common.IsNil(o.GOT) {
		toSerialize["gOT"] = o.GOT
	}
	if !common.IsNil(o.GOV) {
		toSerialize["gOV"] = o.GOV
	}
	if !common.IsNil(o.Smallgp) {
		toSerialize["gp"] = o.Smallgp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionReport) UnmarshalJSON(data []byte) (err error) {
	varExecutionReport := _ExecutionReport{}

	err = json.Unmarshal(data, &varExecutionReport)

	if err != nil {
		return err
	}

	*o = ExecutionReport(varExecutionReport)

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
		delete(additionalProperties, "v")
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
		delete(additionalProperties, "gP")
		delete(additionalProperties, "gOT")
		delete(additionalProperties, "gOV")
		delete(additionalProperties, "gp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionReport struct {
	value *ExecutionReport
	isSet bool
}

func (v NullableExecutionReport) Get() *ExecutionReport {
	return v.value
}

func (v *NullableExecutionReport) Set(val *ExecutionReport) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionReport) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionReport(val *ExecutionReport) *NullableExecutionReport {
	return &NullableExecutionReport{value: val, isSet: true}
}

func (v NullableExecutionReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
