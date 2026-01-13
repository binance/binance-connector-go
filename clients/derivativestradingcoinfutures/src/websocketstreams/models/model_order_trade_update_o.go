/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderTradeUpdateO type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderTradeUpdateO{}

// OrderTradeUpdateO struct for OrderTradeUpdateO
type OrderTradeUpdateO struct {
	Smalls               *string `json:"s,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	S                    *string `json:"S,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallap              *string `json:"ap,omitempty"`
	Smallsp              *string `json:"sp,omitempty"`
	Smallx               *string `json:"x,omitempty"`
	X                    *string `json:"X,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallz               *string `json:"z,omitempty"`
	L                    *string `json:"L,omitempty"`
	Smallma              *string `json:"ma,omitempty"`
	N                    *string `json:"N,omitempty"`
	Smalln               *string `json:"n,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallt               *int64  `json:"t,omitempty"`
	Smallrp              *string `json:"rp,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smallm               *bool   `json:"m,omitempty"`
	R                    *bool   `json:"R,omitempty"`
	Smallwt              *string `json:"wt,omitempty"`
	Smallot              *string `json:"ot,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallcp              *bool   `json:"cp,omitempty"`
	AP                   *string `json:"AP,omitempty"`
	Smallcr              *string `json:"cr,omitempty"`
	PP                   *bool   `json:"pP,omitempty"`
	V                    *string `json:"V,omitempty"`
	Smallpm              *string `json:"pm,omitempty"`
	Smaller              *string `json:"er,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderTradeUpdateO OrderTradeUpdateO

// NewOrderTradeUpdateO instantiates a new OrderTradeUpdateO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTradeUpdateO() *OrderTradeUpdateO {
	this := OrderTradeUpdateO{}
	return &this
}

// NewOrderTradeUpdateOWithDefaults instantiates a new OrderTradeUpdateO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTradeUpdateOWithDefaults() *OrderTradeUpdateO {
	this := OrderTradeUpdateO{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderTradeUpdateO) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *OrderTradeUpdateO) SetSmallc(v string) {
	o.Smallc = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OrderTradeUpdateO) SetS(v string) {
	o.S = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *OrderTradeUpdateO) SetSmallo(v string) {
	o.Smallo = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *OrderTradeUpdateO) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *OrderTradeUpdateO) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *OrderTradeUpdateO) SetSmallp(v string) {
	o.Smallp = &v
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallap() string {
	if o == nil || common.IsNil(o.Smallap) {
		var ret string
		return ret
	}
	return *o.Smallap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallap) {
		return nil, false
	}
	return o.Smallap, true
}

// HasAp returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallap() bool {
	if o != nil && !common.IsNil(o.Smallap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *OrderTradeUpdateO) SetSmallap(v string) {
	o.Smallap = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallsp() string {
	if o == nil || common.IsNil(o.Smallsp) {
		var ret string
		return ret
	}
	return *o.Smallsp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallspOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallsp) {
		return nil, false
	}
	return o.Smallsp, true
}

// HasSp returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallsp() bool {
	if o != nil && !common.IsNil(o.Smallsp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given string and assigns it to the Sp field.
func (o *OrderTradeUpdateO) SetSmallsp(v string) {
	o.Smallsp = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallx() string {
	if o == nil || common.IsNil(o.Smallx) {
		var ret string
		return ret
	}
	return *o.Smallx
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallxOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallx) {
		return nil, false
	}
	return o.Smallx, true
}

// HasX returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallx() bool {
	if o != nil && !common.IsNil(o.Smallx) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *OrderTradeUpdateO) SetSmallx(v string) {
	o.Smallx = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *OrderTradeUpdateO) SetX(v string) {
	o.X = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *OrderTradeUpdateO) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *OrderTradeUpdateO) SetSmalll(v string) {
	o.Smalll = &v
}

// GetZ returns the Z field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallz() string {
	if o == nil || common.IsNil(o.Smallz) {
		var ret string
		return ret
	}
	return *o.Smallz
}

// GetZOk returns a tuple with the Z field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallzOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallz) {
		return nil, false
	}
	return o.Smallz, true
}

// HasZ returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallz() bool {
	if o != nil && !common.IsNil(o.Smallz) {
		return true
	}

	return false
}

// SetZ gets a reference to the given string and assigns it to the Z field.
func (o *OrderTradeUpdateO) SetSmallz(v string) {
	o.Smallz = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *OrderTradeUpdateO) SetL(v string) {
	o.L = &v
}

// GetMa returns the Ma field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallma() string {
	if o == nil || common.IsNil(o.Smallma) {
		var ret string
		return ret
	}
	return *o.Smallma
}

// GetMaOk returns a tuple with the Ma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallmaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallma) {
		return nil, false
	}
	return o.Smallma, true
}

// HasMa returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallma() bool {
	if o != nil && !common.IsNil(o.Smallma) {
		return true
	}

	return false
}

// SetMa gets a reference to the given string and assigns it to the Ma field.
func (o *OrderTradeUpdateO) SetSmallma(v string) {
	o.Smallma = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetN() string {
	if o == nil || common.IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetNOk() (*string, bool) {
	if o == nil || common.IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasN() bool {
	if o != nil && !common.IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OrderTradeUpdateO) SetN(v string) {
	o.N = &v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmalln() string {
	if o == nil || common.IsNil(o.Smalln) {
		var ret string
		return ret
	}
	return *o.Smalln
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallnOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalln) {
		return nil, false
	}
	return o.Smalln, true
}

// HasN returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmalln() bool {
	if o != nil && !common.IsNil(o.Smalln) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *OrderTradeUpdateO) SetSmalln(v string) {
	o.Smalln = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderTradeUpdateO) SetT(v int64) {
	o.T = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *OrderTradeUpdateO) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetRp returns the Rp field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallrp() string {
	if o == nil || common.IsNil(o.Smallrp) {
		var ret string
		return ret
	}
	return *o.Smallrp
}

// GetRpOk returns a tuple with the Rp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallrpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallrp) {
		return nil, false
	}
	return o.Smallrp, true
}

// HasRp returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallrp() bool {
	if o != nil && !common.IsNil(o.Smallrp) {
		return true
	}

	return false
}

// SetRp gets a reference to the given string and assigns it to the Rp field.
func (o *OrderTradeUpdateO) SetSmallrp(v string) {
	o.Smallrp = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *OrderTradeUpdateO) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OrderTradeUpdateO) SetSmalla(v string) {
	o.Smalla = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *OrderTradeUpdateO) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetR() bool {
	if o == nil || common.IsNil(o.R) {
		var ret bool
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetROk() (*bool, bool) {
	if o == nil || common.IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasR() bool {
	if o != nil && !common.IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given bool and assigns it to the R field.
func (o *OrderTradeUpdateO) SetR(v bool) {
	o.R = &v
}

// GetWt returns the Wt field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallwt() string {
	if o == nil || common.IsNil(o.Smallwt) {
		var ret string
		return ret
	}
	return *o.Smallwt
}

// GetWtOk returns a tuple with the Wt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallwtOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallwt) {
		return nil, false
	}
	return o.Smallwt, true
}

// HasWt returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallwt() bool {
	if o != nil && !common.IsNil(o.Smallwt) {
		return true
	}

	return false
}

// SetWt gets a reference to the given string and assigns it to the Wt field.
func (o *OrderTradeUpdateO) SetSmallwt(v string) {
	o.Smallwt = &v
}

// GetOt returns the Ot field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallot() string {
	if o == nil || common.IsNil(o.Smallot) {
		var ret string
		return ret
	}
	return *o.Smallot
}

// GetOtOk returns a tuple with the Ot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallotOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallot) {
		return nil, false
	}
	return o.Smallot, true
}

// HasOt returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallot() bool {
	if o != nil && !common.IsNil(o.Smallot) {
		return true
	}

	return false
}

// SetOt gets a reference to the given string and assigns it to the Ot field.
func (o *OrderTradeUpdateO) SetSmallot(v string) {
	o.Smallot = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *OrderTradeUpdateO) SetSmallps(v string) {
	o.Smallps = &v
}

// GetCp returns the Cp field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallcp() bool {
	if o == nil || common.IsNil(o.Smallcp) {
		var ret bool
		return ret
	}
	return *o.Smallcp
}

// GetCpOk returns a tuple with the Cp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallcpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallcp) {
		return nil, false
	}
	return o.Smallcp, true
}

// HasCp returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallcp() bool {
	if o != nil && !common.IsNil(o.Smallcp) {
		return true
	}

	return false
}

// SetCp gets a reference to the given bool and assigns it to the Cp field.
func (o *OrderTradeUpdateO) SetSmallcp(v bool) {
	o.Smallcp = &v
}

// GetAP returns the AP field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetAP() string {
	if o == nil || common.IsNil(o.AP) {
		var ret string
		return ret
	}
	return *o.AP
}

// GetAPOk returns a tuple with the AP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetAPOk() (*string, bool) {
	if o == nil || common.IsNil(o.AP) {
		return nil, false
	}
	return o.AP, true
}

// HasAP returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasAP() bool {
	if o != nil && !common.IsNil(o.AP) {
		return true
	}

	return false
}

// SetAP gets a reference to the given string and assigns it to the AP field.
func (o *OrderTradeUpdateO) SetAP(v string) {
	o.AP = &v
}

// GetCr returns the Cr field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallcr() string {
	if o == nil || common.IsNil(o.Smallcr) {
		var ret string
		return ret
	}
	return *o.Smallcr
}

// GetCrOk returns a tuple with the Cr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallcrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcr) {
		return nil, false
	}
	return o.Smallcr, true
}

// HasCr returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallcr() bool {
	if o != nil && !common.IsNil(o.Smallcr) {
		return true
	}

	return false
}

// SetCr gets a reference to the given string and assigns it to the Cr field.
func (o *OrderTradeUpdateO) SetSmallcr(v string) {
	o.Smallcr = &v
}

// GetPP returns the PP field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetPP() bool {
	if o == nil || common.IsNil(o.PP) {
		var ret bool
		return ret
	}
	return *o.PP
}

// GetPPOk returns a tuple with the PP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetPPOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PP) {
		return nil, false
	}
	return o.PP, true
}

// HasPP returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasPP() bool {
	if o != nil && !common.IsNil(o.PP) {
		return true
	}

	return false
}

// SetPP gets a reference to the given bool and assigns it to the PP field.
func (o *OrderTradeUpdateO) SetPP(v bool) {
	o.PP = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *OrderTradeUpdateO) SetV(v string) {
	o.V = &v
}

// GetPm returns the Pm field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmallpm() string {
	if o == nil || common.IsNil(o.Smallpm) {
		var ret string
		return ret
	}
	return *o.Smallpm
}

// GetPmOk returns a tuple with the Pm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallpmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpm) {
		return nil, false
	}
	return o.Smallpm, true
}

// HasPm returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmallpm() bool {
	if o != nil && !common.IsNil(o.Smallpm) {
		return true
	}

	return false
}

// SetPm gets a reference to the given string and assigns it to the Pm field.
func (o *OrderTradeUpdateO) SetSmallpm(v string) {
	o.Smallpm = &v
}

// GetEr returns the Er field value if set, zero value otherwise.
func (o *OrderTradeUpdateO) GetSmaller() string {
	if o == nil || common.IsNil(o.Smaller) {
		var ret string
		return ret
	}
	return *o.Smaller
}

// GetErOk returns a tuple with the Er field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTradeUpdateO) GetSmallerOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smaller) {
		return nil, false
	}
	return o.Smaller, true
}

// HasEr returns a boolean if a field has been set.
func (o *OrderTradeUpdateO) HasSmaller() bool {
	if o != nil && !common.IsNil(o.Smaller) {
		return true
	}

	return false
}

// SetEr gets a reference to the given string and assigns it to the Er field.
func (o *OrderTradeUpdateO) SetSmaller(v string) {
	o.Smaller = &v
}

func (o OrderTradeUpdateO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTradeUpdateO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !common.IsNil(o.Smallap) {
		toSerialize["ap"] = o.Smallap
	}
	if !common.IsNil(o.Smallsp) {
		toSerialize["sp"] = o.Smallsp
	}
	if !common.IsNil(o.Smallx) {
		toSerialize["x"] = o.Smallx
	}
	if !common.IsNil(o.X) {
		toSerialize["X"] = o.X
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
	if !common.IsNil(o.Smallma) {
		toSerialize["ma"] = o.Smallma
	}
	if !common.IsNil(o.N) {
		toSerialize["N"] = o.N
	}
	if !common.IsNil(o.Smalln) {
		toSerialize["n"] = o.Smalln
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smallrp) {
		toSerialize["rp"] = o.Smallrp
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.R) {
		toSerialize["R"] = o.R
	}
	if !common.IsNil(o.Smallwt) {
		toSerialize["wt"] = o.Smallwt
	}
	if !common.IsNil(o.Smallot) {
		toSerialize["ot"] = o.Smallot
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallcp) {
		toSerialize["cp"] = o.Smallcp
	}
	if !common.IsNil(o.AP) {
		toSerialize["AP"] = o.AP
	}
	if !common.IsNil(o.Smallcr) {
		toSerialize["cr"] = o.Smallcr
	}
	if !common.IsNil(o.PP) {
		toSerialize["pP"] = o.PP
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.Smallpm) {
		toSerialize["pm"] = o.Smallpm
	}
	if !common.IsNil(o.Smaller) {
		toSerialize["er"] = o.Smaller
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderTradeUpdateO) UnmarshalJSON(data []byte) (err error) {
	varOrderTradeUpdateO := _OrderTradeUpdateO{}

	err = json.Unmarshal(data, &varOrderTradeUpdateO)

	if err != nil {
		return err
	}

	*o = OrderTradeUpdateO(varOrderTradeUpdateO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "S")
		delete(additionalProperties, "o")
		delete(additionalProperties, "f")
		delete(additionalProperties, "q")
		delete(additionalProperties, "p")
		delete(additionalProperties, "ap")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "x")
		delete(additionalProperties, "X")
		delete(additionalProperties, "i")
		delete(additionalProperties, "l")
		delete(additionalProperties, "z")
		delete(additionalProperties, "L")
		delete(additionalProperties, "ma")
		delete(additionalProperties, "N")
		delete(additionalProperties, "n")
		delete(additionalProperties, "T")
		delete(additionalProperties, "t")
		delete(additionalProperties, "rp")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		delete(additionalProperties, "m")
		delete(additionalProperties, "R")
		delete(additionalProperties, "wt")
		delete(additionalProperties, "ot")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "cp")
		delete(additionalProperties, "AP")
		delete(additionalProperties, "cr")
		delete(additionalProperties, "pP")
		delete(additionalProperties, "V")
		delete(additionalProperties, "pm")
		delete(additionalProperties, "er")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderTradeUpdateO struct {
	value *OrderTradeUpdateO
	isSet bool
}

func (v NullableOrderTradeUpdateO) Get() *OrderTradeUpdateO {
	return v.value
}

func (v *NullableOrderTradeUpdateO) Set(val *OrderTradeUpdateO) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTradeUpdateO) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTradeUpdateO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTradeUpdateO(val *OrderTradeUpdateO) *NullableOrderTradeUpdateO {
	return &NullableOrderTradeUpdateO{value: val, isSet: true}
}

func (v NullableOrderTradeUpdateO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTradeUpdateO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
