/*
Portfolio Margin WebSocket Market Streams

Access account information, manage margin positions, and trade with Binance Portfolio Margin.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AlgoOrderUpdateAo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AlgoOrderUpdateAo{}

// AlgoOrderUpdateAo Algo order info
type AlgoOrderUpdateAo struct {
	// Client Algo Id
	Caid *string `json:"caid,omitempty"`
	// Algo Id
	Aid *int64 `json:"aid,omitempty"`
	// Algo Type
	Smallat *string `json:"at,omitempty"`
	// Order Type
	Smallo *string `json:"o,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Side
	S *string `json:"S,omitempty"`
	// Position Side
	Smallps *string `json:"ps,omitempty"`
	// Time in Force
	Smallf *string `json:"f,omitempty"`
	// Quantity
	Smallq *string `json:"q,omitempty"`
	// Algo Status: NEW, CANCELED, TRIGGERING, TRIGGERED, FINISHED, REJECTED, EXPIRED
	X *string `json:"X,omitempty"`
	// Actual order ID in matching engine
	Smallai *string `json:"ai,omitempty"`
	// Avg fill price in matching engine
	Smallap *string `json:"ap,omitempty"`
	// Executed quantity in matching engine
	Smallaq *string `json:"aq,omitempty"`
	// Actual order type in matching engine
	Act *string `json:"act,omitempty"`
	// Trigger Price
	Smalltp *string `json:"tp,omitempty"`
	// Order Price
	Smallp *string `json:"p,omitempty"`
	// Self Trade Prevention Mode
	V *string `json:"V,omitempty"`
	// Working Type
	Smallwt *string `json:"wt,omitempty"`
	// Price Match
	Smallpm *string `json:"pm,omitempty"`
	// If Close-All
	Smallcp *bool `json:"cp,omitempty"`
	// If price protection is on
	SmallpP *bool `json:"pP,omitempty"`
	// Is reduce only
	R *bool `json:"R,omitempty"`
	// Trigger Time
	Smalltt *int64 `json:"tt,omitempty"`
	// Good Till Date
	Gtd *int64 `json:"gtd,omitempty"`
	// Algo order failed reason
	Smallrm              *string `json:"rm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlgoOrderUpdateAo AlgoOrderUpdateAo

// NewAlgoOrderUpdateAo instantiates a new AlgoOrderUpdateAo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgoOrderUpdateAo() *AlgoOrderUpdateAo {
	this := AlgoOrderUpdateAo{}
	return &this
}

// NewAlgoOrderUpdateAoWithDefaults instantiates a new AlgoOrderUpdateAo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgoOrderUpdateAoWithDefaults() *AlgoOrderUpdateAo {
	this := AlgoOrderUpdateAo{}
	return &this
}

// GetCaid returns the Caid field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetCaid() string {
	if o == nil || common.IsNil(o.Caid) {
		var ret string
		return ret
	}
	return *o.Caid
}

// GetCaidOk returns a tuple with the Caid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetCaidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Caid) {
		return nil, false
	}
	return o.Caid, true
}

// HasCaid returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasCaid() bool {
	if o != nil && !common.IsNil(o.Caid) {
		return true
	}

	return false
}

// SetCaid gets a reference to the given string and assigns it to the Caid field.
func (o *AlgoOrderUpdateAo) SetCaid(v string) {
	o.Caid = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetAid() int64 {
	if o == nil || common.IsNil(o.Aid) {
		var ret int64
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetAidOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasAid() bool {
	if o != nil && !common.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given int64 and assigns it to the Aid field.
func (o *AlgoOrderUpdateAo) SetAid(v int64) {
	o.Aid = &v
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallat() string {
	if o == nil || common.IsNil(o.Smallat) {
		var ret string
		return ret
	}
	return *o.Smallat
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallatOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallat) {
		return nil, false
	}
	return o.Smallat, true
}

// HasAt returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallat() bool {
	if o != nil && !common.IsNil(o.Smallat) {
		return true
	}

	return false
}

// SetAt gets a reference to the given string and assigns it to the At field.
func (o *AlgoOrderUpdateAo) SetSmallat(v string) {
	o.Smallat = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AlgoOrderUpdateAo) SetSmallo(v string) {
	o.Smallo = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AlgoOrderUpdateAo) SetSmalls(v string) {
	o.Smalls = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AlgoOrderUpdateAo) SetS(v string) {
	o.S = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *AlgoOrderUpdateAo) SetSmallps(v string) {
	o.Smallps = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *AlgoOrderUpdateAo) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AlgoOrderUpdateAo) SetSmallq(v string) {
	o.Smallq = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *AlgoOrderUpdateAo) SetX(v string) {
	o.X = &v
}

// GetAi returns the Ai field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallai() string {
	if o == nil || common.IsNil(o.Smallai) {
		var ret string
		return ret
	}
	return *o.Smallai
}

// GetAiOk returns a tuple with the Ai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallaiOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallai) {
		return nil, false
	}
	return o.Smallai, true
}

// HasAi returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallai() bool {
	if o != nil && !common.IsNil(o.Smallai) {
		return true
	}

	return false
}

// SetAi gets a reference to the given string and assigns it to the Ai field.
func (o *AlgoOrderUpdateAo) SetSmallai(v string) {
	o.Smallai = &v
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallap() string {
	if o == nil || common.IsNil(o.Smallap) {
		var ret string
		return ret
	}
	return *o.Smallap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallap) {
		return nil, false
	}
	return o.Smallap, true
}

// HasAp returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallap() bool {
	if o != nil && !common.IsNil(o.Smallap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *AlgoOrderUpdateAo) SetSmallap(v string) {
	o.Smallap = &v
}

// GetAq returns the Aq field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallaq() string {
	if o == nil || common.IsNil(o.Smallaq) {
		var ret string
		return ret
	}
	return *o.Smallaq
}

// GetAqOk returns a tuple with the Aq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallaqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallaq) {
		return nil, false
	}
	return o.Smallaq, true
}

// HasAq returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallaq() bool {
	if o != nil && !common.IsNil(o.Smallaq) {
		return true
	}

	return false
}

// SetAq gets a reference to the given string and assigns it to the Aq field.
func (o *AlgoOrderUpdateAo) SetSmallaq(v string) {
	o.Smallaq = &v
}

// GetAct returns the Act field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetAct() string {
	if o == nil || common.IsNil(o.Act) {
		var ret string
		return ret
	}
	return *o.Act
}

// GetActOk returns a tuple with the Act field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetActOk() (*string, bool) {
	if o == nil || common.IsNil(o.Act) {
		return nil, false
	}
	return o.Act, true
}

// HasAct returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasAct() bool {
	if o != nil && !common.IsNil(o.Act) {
		return true
	}

	return false
}

// SetAct gets a reference to the given string and assigns it to the Act field.
func (o *AlgoOrderUpdateAo) SetAct(v string) {
	o.Act = &v
}

// GetTp returns the Tp field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmalltp() string {
	if o == nil || common.IsNil(o.Smalltp) {
		var ret string
		return ret
	}
	return *o.Smalltp
}

// GetTpOk returns a tuple with the Tp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmalltpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalltp) {
		return nil, false
	}
	return o.Smalltp, true
}

// HasTp returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmalltp() bool {
	if o != nil && !common.IsNil(o.Smalltp) {
		return true
	}

	return false
}

// SetTp gets a reference to the given string and assigns it to the Tp field.
func (o *AlgoOrderUpdateAo) SetSmalltp(v string) {
	o.Smalltp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AlgoOrderUpdateAo) SetSmallp(v string) {
	o.Smallp = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AlgoOrderUpdateAo) SetV(v string) {
	o.V = &v
}

// GetWt returns the Wt field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallwt() string {
	if o == nil || common.IsNil(o.Smallwt) {
		var ret string
		return ret
	}
	return *o.Smallwt
}

// GetWtOk returns a tuple with the Wt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallwtOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallwt) {
		return nil, false
	}
	return o.Smallwt, true
}

// HasWt returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallwt() bool {
	if o != nil && !common.IsNil(o.Smallwt) {
		return true
	}

	return false
}

// SetWt gets a reference to the given string and assigns it to the Wt field.
func (o *AlgoOrderUpdateAo) SetSmallwt(v string) {
	o.Smallwt = &v
}

// GetPm returns the Pm field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallpm() string {
	if o == nil || common.IsNil(o.Smallpm) {
		var ret string
		return ret
	}
	return *o.Smallpm
}

// GetPmOk returns a tuple with the Pm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallpmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpm) {
		return nil, false
	}
	return o.Smallpm, true
}

// HasPm returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallpm() bool {
	if o != nil && !common.IsNil(o.Smallpm) {
		return true
	}

	return false
}

// SetPm gets a reference to the given string and assigns it to the Pm field.
func (o *AlgoOrderUpdateAo) SetSmallpm(v string) {
	o.Smallpm = &v
}

// GetCp returns the Cp field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallcp() bool {
	if o == nil || common.IsNil(o.Smallcp) {
		var ret bool
		return ret
	}
	return *o.Smallcp
}

// GetCpOk returns a tuple with the Cp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallcpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallcp) {
		return nil, false
	}
	return o.Smallcp, true
}

// HasCp returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallcp() bool {
	if o != nil && !common.IsNil(o.Smallcp) {
		return true
	}

	return false
}

// SetCp gets a reference to the given bool and assigns it to the Cp field.
func (o *AlgoOrderUpdateAo) SetSmallcp(v bool) {
	o.Smallcp = &v
}

// GetPP returns the PP field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallpP() bool {
	if o == nil || common.IsNil(o.SmallpP) {
		var ret bool
		return ret
	}
	return *o.SmallpP
}

// GetPPOk returns a tuple with the PP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallpPOk() (*bool, bool) {
	if o == nil || common.IsNil(o.SmallpP) {
		return nil, false
	}
	return o.SmallpP, true
}

// HasPP returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallpP() bool {
	if o != nil && !common.IsNil(o.SmallpP) {
		return true
	}

	return false
}

// SetPP gets a reference to the given bool and assigns it to the PP field.
func (o *AlgoOrderUpdateAo) SetSmallpP(v bool) {
	o.SmallpP = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetR() bool {
	if o == nil || common.IsNil(o.R) {
		var ret bool
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetROk() (*bool, bool) {
	if o == nil || common.IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasR() bool {
	if o != nil && !common.IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given bool and assigns it to the R field.
func (o *AlgoOrderUpdateAo) SetR(v bool) {
	o.R = &v
}

// GetTt returns the Tt field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmalltt() int64 {
	if o == nil || common.IsNil(o.Smalltt) {
		var ret int64
		return ret
	}
	return *o.Smalltt
}

// GetTtOk returns a tuple with the Tt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallttOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalltt) {
		return nil, false
	}
	return o.Smalltt, true
}

// HasTt returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmalltt() bool {
	if o != nil && !common.IsNil(o.Smalltt) {
		return true
	}

	return false
}

// SetTt gets a reference to the given int64 and assigns it to the Tt field.
func (o *AlgoOrderUpdateAo) SetSmalltt(v int64) {
	o.Smalltt = &v
}

// GetGtd returns the Gtd field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetGtd() int64 {
	if o == nil || common.IsNil(o.Gtd) {
		var ret int64
		return ret
	}
	return *o.Gtd
}

// GetGtdOk returns a tuple with the Gtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetGtdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Gtd) {
		return nil, false
	}
	return o.Gtd, true
}

// HasGtd returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasGtd() bool {
	if o != nil && !common.IsNil(o.Gtd) {
		return true
	}

	return false
}

// SetGtd gets a reference to the given int64 and assigns it to the Gtd field.
func (o *AlgoOrderUpdateAo) SetGtd(v int64) {
	o.Gtd = &v
}

// GetRm returns the Rm field value if set, zero value otherwise.
func (o *AlgoOrderUpdateAo) GetSmallrm() string {
	if o == nil || common.IsNil(o.Smallrm) {
		var ret string
		return ret
	}
	return *o.Smallrm
}

// GetRmOk returns a tuple with the Rm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoOrderUpdateAo) GetSmallrmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallrm) {
		return nil, false
	}
	return o.Smallrm, true
}

// HasRm returns a boolean if a field has been set.
func (o *AlgoOrderUpdateAo) HasSmallrm() bool {
	if o != nil && !common.IsNil(o.Smallrm) {
		return true
	}

	return false
}

// SetRm gets a reference to the given string and assigns it to the Rm field.
func (o *AlgoOrderUpdateAo) SetSmallrm(v string) {
	o.Smallrm = &v
}

func (o AlgoOrderUpdateAo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgoOrderUpdateAo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Caid) {
		toSerialize["caid"] = o.Caid
	}
	if !common.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !common.IsNil(o.Smallat) {
		toSerialize["at"] = o.Smallat
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallf) {
		toSerialize["f"] = o.Smallf
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.X) {
		toSerialize["X"] = o.X
	}
	if !common.IsNil(o.Smallai) {
		toSerialize["ai"] = o.Smallai
	}
	if !common.IsNil(o.Smallap) {
		toSerialize["ap"] = o.Smallap
	}
	if !common.IsNil(o.Smallaq) {
		toSerialize["aq"] = o.Smallaq
	}
	if !common.IsNil(o.Act) {
		toSerialize["act"] = o.Act
	}
	if !common.IsNil(o.Smalltp) {
		toSerialize["tp"] = o.Smalltp
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.Smallwt) {
		toSerialize["wt"] = o.Smallwt
	}
	if !common.IsNil(o.Smallpm) {
		toSerialize["pm"] = o.Smallpm
	}
	if !common.IsNil(o.Smallcp) {
		toSerialize["cp"] = o.Smallcp
	}
	if !common.IsNil(o.SmallpP) {
		toSerialize["pP"] = o.SmallpP
	}
	if !common.IsNil(o.R) {
		toSerialize["R"] = o.R
	}
	if !common.IsNil(o.Smalltt) {
		toSerialize["tt"] = o.Smalltt
	}
	if !common.IsNil(o.Gtd) {
		toSerialize["gtd"] = o.Gtd
	}
	if !common.IsNil(o.Smallrm) {
		toSerialize["rm"] = o.Smallrm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AlgoOrderUpdateAo) UnmarshalJSON(data []byte) (err error) {
	varAlgoOrderUpdateAo := _AlgoOrderUpdateAo{}

	err = json.Unmarshal(data, &varAlgoOrderUpdateAo)

	if err != nil {
		return err
	}

	*o = AlgoOrderUpdateAo(varAlgoOrderUpdateAo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caid")
		delete(additionalProperties, "aid")
		delete(additionalProperties, "at")
		delete(additionalProperties, "o")
		delete(additionalProperties, "s")
		delete(additionalProperties, "S")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "f")
		delete(additionalProperties, "q")
		delete(additionalProperties, "X")
		delete(additionalProperties, "ai")
		delete(additionalProperties, "ap")
		delete(additionalProperties, "aq")
		delete(additionalProperties, "act")
		delete(additionalProperties, "tp")
		delete(additionalProperties, "p")
		delete(additionalProperties, "V")
		delete(additionalProperties, "wt")
		delete(additionalProperties, "pm")
		delete(additionalProperties, "cp")
		delete(additionalProperties, "pP")
		delete(additionalProperties, "R")
		delete(additionalProperties, "tt")
		delete(additionalProperties, "gtd")
		delete(additionalProperties, "rm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAlgoOrderUpdateAo struct {
	value *AlgoOrderUpdateAo
	isSet bool
}

func (v NullableAlgoOrderUpdateAo) Get() *AlgoOrderUpdateAo {
	return v.value
}

func (v *NullableAlgoOrderUpdateAo) Set(val *AlgoOrderUpdateAo) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgoOrderUpdateAo) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgoOrderUpdateAo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgoOrderUpdateAo(val *AlgoOrderUpdateAo) *NullableAlgoOrderUpdateAo {
	return &NullableAlgoOrderUpdateAo{value: val, isSet: true}
}

func (v NullableAlgoOrderUpdateAo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgoOrderUpdateAo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
