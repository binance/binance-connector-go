/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AlgoUpdateO type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AlgoUpdateO{}

// AlgoUpdateO struct for AlgoUpdateO
type AlgoUpdateO struct {
	Caid                 *string `json:"caid,omitempty"`
	Aid                  *int64  `json:"aid,omitempty"`
	Smallat              *string `json:"at,omitempty"`
	Smallo               *string `json:"o,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	S                    *string `json:"S,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	X                    *string `json:"X,omitempty"`
	Smallai              *string `json:"ai,omitempty"`
	Smallap              *string `json:"ap,omitempty"`
	Smallaq              *string `json:"aq,omitempty"`
	Act                  *string `json:"act,omitempty"`
	Smalltp              *string `json:"tp,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	V                    *string `json:"V,omitempty"`
	Smallwt              *string `json:"wt,omitempty"`
	Smallpm              *string `json:"pm,omitempty"`
	Smallcp              *bool   `json:"cp,omitempty"`
	PP                   *bool   `json:"pP,omitempty"`
	R                    *bool   `json:"R,omitempty"`
	Smalltt              *int64  `json:"tt,omitempty"`
	Gtd                  *int64  `json:"gtd,omitempty"`
	Smallrm              *string `json:"rm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AlgoUpdateO AlgoUpdateO

// NewAlgoUpdateO instantiates a new AlgoUpdateO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlgoUpdateO() *AlgoUpdateO {
	this := AlgoUpdateO{}
	return &this
}

// NewAlgoUpdateOWithDefaults instantiates a new AlgoUpdateO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlgoUpdateOWithDefaults() *AlgoUpdateO {
	this := AlgoUpdateO{}
	return &this
}

// GetCaid returns the Caid field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetCaid() string {
	if o == nil || common.IsNil(o.Caid) {
		var ret string
		return ret
	}
	return *o.Caid
}

// GetCaidOk returns a tuple with the Caid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetCaidOk() (*string, bool) {
	if o == nil || common.IsNil(o.Caid) {
		return nil, false
	}
	return o.Caid, true
}

// HasCaid returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasCaid() bool {
	if o != nil && !common.IsNil(o.Caid) {
		return true
	}

	return false
}

// SetCaid gets a reference to the given string and assigns it to the Caid field.
func (o *AlgoUpdateO) SetCaid(v string) {
	o.Caid = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetAid() int64 {
	if o == nil || common.IsNil(o.Aid) {
		var ret int64
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetAidOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasAid() bool {
	if o != nil && !common.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given int64 and assigns it to the Aid field.
func (o *AlgoUpdateO) SetAid(v int64) {
	o.Aid = &v
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallat() string {
	if o == nil || common.IsNil(o.Smallat) {
		var ret string
		return ret
	}
	return *o.Smallat
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallatOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallat) {
		return nil, false
	}
	return o.Smallat, true
}

// HasAt returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallat() bool {
	if o != nil && !common.IsNil(o.Smallat) {
		return true
	}

	return false
}

// SetAt gets a reference to the given string and assigns it to the At field.
func (o *AlgoUpdateO) SetSmallat(v string) {
	o.Smallat = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AlgoUpdateO) SetSmallo(v string) {
	o.Smallo = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AlgoUpdateO) SetSmalls(v string) {
	o.Smalls = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AlgoUpdateO) SetS(v string) {
	o.S = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *AlgoUpdateO) SetSmallps(v string) {
	o.Smallps = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *AlgoUpdateO) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AlgoUpdateO) SetSmallq(v string) {
	o.Smallq = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetX() string {
	if o == nil || common.IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetXOk() (*string, bool) {
	if o == nil || common.IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasX() bool {
	if o != nil && !common.IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *AlgoUpdateO) SetX(v string) {
	o.X = &v
}

// GetAi returns the Ai field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallai() string {
	if o == nil || common.IsNil(o.Smallai) {
		var ret string
		return ret
	}
	return *o.Smallai
}

// GetAiOk returns a tuple with the Ai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallaiOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallai) {
		return nil, false
	}
	return o.Smallai, true
}

// HasAi returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallai() bool {
	if o != nil && !common.IsNil(o.Smallai) {
		return true
	}

	return false
}

// SetAi gets a reference to the given string and assigns it to the Ai field.
func (o *AlgoUpdateO) SetSmallai(v string) {
	o.Smallai = &v
}

// GetAp returns the Ap field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallap() string {
	if o == nil || common.IsNil(o.Smallap) {
		var ret string
		return ret
	}
	return *o.Smallap
}

// GetApOk returns a tuple with the Ap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallapOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallap) {
		return nil, false
	}
	return o.Smallap, true
}

// HasAp returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallap() bool {
	if o != nil && !common.IsNil(o.Smallap) {
		return true
	}

	return false
}

// SetAp gets a reference to the given string and assigns it to the Ap field.
func (o *AlgoUpdateO) SetSmallap(v string) {
	o.Smallap = &v
}

// GetAq returns the Aq field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallaq() string {
	if o == nil || common.IsNil(o.Smallaq) {
		var ret string
		return ret
	}
	return *o.Smallaq
}

// GetAqOk returns a tuple with the Aq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallaqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallaq) {
		return nil, false
	}
	return o.Smallaq, true
}

// HasAq returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallaq() bool {
	if o != nil && !common.IsNil(o.Smallaq) {
		return true
	}

	return false
}

// SetAq gets a reference to the given string and assigns it to the Aq field.
func (o *AlgoUpdateO) SetSmallaq(v string) {
	o.Smallaq = &v
}

// GetAct returns the Act field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetAct() string {
	if o == nil || common.IsNil(o.Act) {
		var ret string
		return ret
	}
	return *o.Act
}

// GetActOk returns a tuple with the Act field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetActOk() (*string, bool) {
	if o == nil || common.IsNil(o.Act) {
		return nil, false
	}
	return o.Act, true
}

// HasAct returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasAct() bool {
	if o != nil && !common.IsNil(o.Act) {
		return true
	}

	return false
}

// SetAct gets a reference to the given string and assigns it to the Act field.
func (o *AlgoUpdateO) SetAct(v string) {
	o.Act = &v
}

// GetTp returns the Tp field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmalltp() string {
	if o == nil || common.IsNil(o.Smalltp) {
		var ret string
		return ret
	}
	return *o.Smalltp
}

// GetTpOk returns a tuple with the Tp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmalltpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalltp) {
		return nil, false
	}
	return o.Smalltp, true
}

// HasTp returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmalltp() bool {
	if o != nil && !common.IsNil(o.Smalltp) {
		return true
	}

	return false
}

// SetTp gets a reference to the given string and assigns it to the Tp field.
func (o *AlgoUpdateO) SetSmalltp(v string) {
	o.Smalltp = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AlgoUpdateO) SetSmallp(v string) {
	o.Smallp = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AlgoUpdateO) SetV(v string) {
	o.V = &v
}

// GetWt returns the Wt field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallwt() string {
	if o == nil || common.IsNil(o.Smallwt) {
		var ret string
		return ret
	}
	return *o.Smallwt
}

// GetWtOk returns a tuple with the Wt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallwtOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallwt) {
		return nil, false
	}
	return o.Smallwt, true
}

// HasWt returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallwt() bool {
	if o != nil && !common.IsNil(o.Smallwt) {
		return true
	}

	return false
}

// SetWt gets a reference to the given string and assigns it to the Wt field.
func (o *AlgoUpdateO) SetSmallwt(v string) {
	o.Smallwt = &v
}

// GetPm returns the Pm field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallpm() string {
	if o == nil || common.IsNil(o.Smallpm) {
		var ret string
		return ret
	}
	return *o.Smallpm
}

// GetPmOk returns a tuple with the Pm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallpmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpm) {
		return nil, false
	}
	return o.Smallpm, true
}

// HasPm returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallpm() bool {
	if o != nil && !common.IsNil(o.Smallpm) {
		return true
	}

	return false
}

// SetPm gets a reference to the given string and assigns it to the Pm field.
func (o *AlgoUpdateO) SetSmallpm(v string) {
	o.Smallpm = &v
}

// GetCp returns the Cp field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallcp() bool {
	if o == nil || common.IsNil(o.Smallcp) {
		var ret bool
		return ret
	}
	return *o.Smallcp
}

// GetCpOk returns a tuple with the Cp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallcpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallcp) {
		return nil, false
	}
	return o.Smallcp, true
}

// HasCp returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallcp() bool {
	if o != nil && !common.IsNil(o.Smallcp) {
		return true
	}

	return false
}

// SetCp gets a reference to the given bool and assigns it to the Cp field.
func (o *AlgoUpdateO) SetSmallcp(v bool) {
	o.Smallcp = &v
}

// GetPP returns the PP field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetPP() bool {
	if o == nil || common.IsNil(o.PP) {
		var ret bool
		return ret
	}
	return *o.PP
}

// GetPPOk returns a tuple with the PP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetPPOk() (*bool, bool) {
	if o == nil || common.IsNil(o.PP) {
		return nil, false
	}
	return o.PP, true
}

// HasPP returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasPP() bool {
	if o != nil && !common.IsNil(o.PP) {
		return true
	}

	return false
}

// SetPP gets a reference to the given bool and assigns it to the PP field.
func (o *AlgoUpdateO) SetPP(v bool) {
	o.PP = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetR() bool {
	if o == nil || common.IsNil(o.R) {
		var ret bool
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetROk() (*bool, bool) {
	if o == nil || common.IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasR() bool {
	if o != nil && !common.IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given bool and assigns it to the R field.
func (o *AlgoUpdateO) SetR(v bool) {
	o.R = &v
}

// GetTt returns the Tt field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmalltt() int64 {
	if o == nil || common.IsNil(o.Smalltt) {
		var ret int64
		return ret
	}
	return *o.Smalltt
}

// GetTtOk returns a tuple with the Tt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallttOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalltt) {
		return nil, false
	}
	return o.Smalltt, true
}

// HasTt returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmalltt() bool {
	if o != nil && !common.IsNil(o.Smalltt) {
		return true
	}

	return false
}

// SetTt gets a reference to the given int64 and assigns it to the Tt field.
func (o *AlgoUpdateO) SetSmalltt(v int64) {
	o.Smalltt = &v
}

// GetGtd returns the Gtd field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetGtd() int64 {
	if o == nil || common.IsNil(o.Gtd) {
		var ret int64
		return ret
	}
	return *o.Gtd
}

// GetGtdOk returns a tuple with the Gtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetGtdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Gtd) {
		return nil, false
	}
	return o.Gtd, true
}

// HasGtd returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasGtd() bool {
	if o != nil && !common.IsNil(o.Gtd) {
		return true
	}

	return false
}

// SetGtd gets a reference to the given int64 and assigns it to the Gtd field.
func (o *AlgoUpdateO) SetGtd(v int64) {
	o.Gtd = &v
}

// GetRm returns the Rm field value if set, zero value otherwise.
func (o *AlgoUpdateO) GetSmallrm() string {
	if o == nil || common.IsNil(o.Smallrm) {
		var ret string
		return ret
	}
	return *o.Smallrm
}

// GetRmOk returns a tuple with the Rm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlgoUpdateO) GetSmallrmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallrm) {
		return nil, false
	}
	return o.Smallrm, true
}

// HasRm returns a boolean if a field has been set.
func (o *AlgoUpdateO) HasSmallrm() bool {
	if o != nil && !common.IsNil(o.Smallrm) {
		return true
	}

	return false
}

// SetRm gets a reference to the given string and assigns it to the Rm field.
func (o *AlgoUpdateO) SetSmallrm(v string) {
	o.Smallrm = &v
}

func (o AlgoUpdateO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlgoUpdateO) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.PP) {
		toSerialize["pP"] = o.PP
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

func (o *AlgoUpdateO) UnmarshalJSON(data []byte) (err error) {
	varAlgoUpdateO := _AlgoUpdateO{}

	err = json.Unmarshal(data, &varAlgoUpdateO)

	if err != nil {
		return err
	}

	*o = AlgoUpdateO(varAlgoUpdateO)

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

type NullableAlgoUpdateO struct {
	value *AlgoUpdateO
	isSet bool
}

func (v NullableAlgoUpdateO) Get() *AlgoUpdateO {
	return v.value
}

func (v *NullableAlgoUpdateO) Set(val *AlgoUpdateO) {
	v.value = val
	v.isSet = true
}

func (v NullableAlgoUpdateO) IsSet() bool {
	return v.isSet
}

func (v *NullableAlgoUpdateO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlgoUpdateO(val *AlgoUpdateO) *NullableAlgoUpdateO {
	return &NullableAlgoUpdateO{value: val, isSet: true}
}

func (v NullableAlgoUpdateO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlgoUpdateO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
