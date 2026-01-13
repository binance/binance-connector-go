/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the ConditionalOrderTradeUpdateSo type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ConditionalOrderTradeUpdateSo{}

// ConditionalOrderTradeUpdateSo struct for ConditionalOrderTradeUpdateSo
type ConditionalOrderTradeUpdateSo struct {
	Smalls               *string `json:"s,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	Smallsi              *int64  `json:"si,omitempty"`
	S                    *string `json:"S,omitempty"`
	Smallst              *string `json:"st,omitempty"`
	Smallf               *string `json:"f,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallsp              *string `json:"sp,omitempty"`
	Smallos              *string `json:"os,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallut              *int64  `json:"ut,omitempty"`
	R                    *bool   `json:"R,omitempty"`
	Smallwt              *string `json:"wt,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallcp              *bool   `json:"cp,omitempty"`
	AP                   *string `json:"AP,omitempty"`
	Smallcr              *string `json:"cr,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	V                    *string `json:"V,omitempty"`
	Gtd                  *int64  `json:"gtd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConditionalOrderTradeUpdateSo ConditionalOrderTradeUpdateSo

// NewConditionalOrderTradeUpdateSo instantiates a new ConditionalOrderTradeUpdateSo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalOrderTradeUpdateSo() *ConditionalOrderTradeUpdateSo {
	this := ConditionalOrderTradeUpdateSo{}
	return &this
}

// NewConditionalOrderTradeUpdateSoWithDefaults instantiates a new ConditionalOrderTradeUpdateSo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalOrderTradeUpdateSoWithDefaults() *ConditionalOrderTradeUpdateSo {
	this := ConditionalOrderTradeUpdateSo{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ConditionalOrderTradeUpdateSo) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallc(v string) {
	o.Smallc = &v
}

// GetSi returns the Si field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallsi() int64 {
	if o == nil || common.IsNil(o.Smallsi) {
		var ret int64
		return ret
	}
	return *o.Smallsi
}

// GetSiOk returns a tuple with the Si field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallsiOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallsi) {
		return nil, false
	}
	return o.Smallsi, true
}

// HasSi returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallsi() bool {
	if o != nil && !common.IsNil(o.Smallsi) {
		return true
	}

	return false
}

// SetSi gets a reference to the given int64 and assigns it to the Si field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallsi(v int64) {
	o.Smallsi = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *ConditionalOrderTradeUpdateSo) SetS(v string) {
	o.S = &v
}

// GetSt returns the St field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallst() string {
	if o == nil || common.IsNil(o.Smallst) {
		var ret string
		return ret
	}
	return *o.Smallst
}

// GetStOk returns a tuple with the St field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallstOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallst) {
		return nil, false
	}
	return o.Smallst, true
}

// HasSt returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallst() bool {
	if o != nil && !common.IsNil(o.Smallst) {
		return true
	}

	return false
}

// SetSt gets a reference to the given string and assigns it to the St field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallst(v string) {
	o.Smallst = &v
}

// GetF returns the F field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallf() string {
	if o == nil || common.IsNil(o.Smallf) {
		var ret string
		return ret
	}
	return *o.Smallf
}

// GetFOk returns a tuple with the F field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallf) {
		return nil, false
	}
	return o.Smallf, true
}

// HasF returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallf() bool {
	if o != nil && !common.IsNil(o.Smallf) {
		return true
	}

	return false
}

// SetF gets a reference to the given string and assigns it to the F field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallf(v string) {
	o.Smallf = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallp(v string) {
	o.Smallp = &v
}

// GetSp returns the Sp field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallsp() string {
	if o == nil || common.IsNil(o.Smallsp) {
		var ret string
		return ret
	}
	return *o.Smallsp
}

// GetSpOk returns a tuple with the Sp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallspOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallsp) {
		return nil, false
	}
	return o.Smallsp, true
}

// HasSp returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallsp() bool {
	if o != nil && !common.IsNil(o.Smallsp) {
		return true
	}

	return false
}

// SetSp gets a reference to the given string and assigns it to the Sp field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallsp(v string) {
	o.Smallsp = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallos() string {
	if o == nil || common.IsNil(o.Smallos) {
		var ret string
		return ret
	}
	return *o.Smallos
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallosOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallos) {
		return nil, false
	}
	return o.Smallos, true
}

// HasOs returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallos() bool {
	if o != nil && !common.IsNil(o.Smallos) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallos(v string) {
	o.Smallos = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *ConditionalOrderTradeUpdateSo) SetT(v int64) {
	o.T = &v
}

// GetUt returns the Ut field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallut() int64 {
	if o == nil || common.IsNil(o.Smallut) {
		var ret int64
		return ret
	}
	return *o.Smallut
}

// GetUtOk returns a tuple with the Ut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallutOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallut) {
		return nil, false
	}
	return o.Smallut, true
}

// HasUt returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallut() bool {
	if o != nil && !common.IsNil(o.Smallut) {
		return true
	}

	return false
}

// SetUt gets a reference to the given int64 and assigns it to the Ut field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallut(v int64) {
	o.Smallut = &v
}

// GetR returns the R field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetR() bool {
	if o == nil || common.IsNil(o.R) {
		var ret bool
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetROk() (*bool, bool) {
	if o == nil || common.IsNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasR() bool {
	if o != nil && !common.IsNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given bool and assigns it to the R field.
func (o *ConditionalOrderTradeUpdateSo) SetR(v bool) {
	o.R = &v
}

// GetWt returns the Wt field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallwt() string {
	if o == nil || common.IsNil(o.Smallwt) {
		var ret string
		return ret
	}
	return *o.Smallwt
}

// GetWtOk returns a tuple with the Wt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallwtOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallwt) {
		return nil, false
	}
	return o.Smallwt, true
}

// HasWt returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallwt() bool {
	if o != nil && !common.IsNil(o.Smallwt) {
		return true
	}

	return false
}

// SetWt gets a reference to the given string and assigns it to the Wt field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallwt(v string) {
	o.Smallwt = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallps(v string) {
	o.Smallps = &v
}

// GetCp returns the Cp field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallcp() bool {
	if o == nil || common.IsNil(o.Smallcp) {
		var ret bool
		return ret
	}
	return *o.Smallcp
}

// GetCpOk returns a tuple with the Cp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallcpOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallcp) {
		return nil, false
	}
	return o.Smallcp, true
}

// HasCp returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallcp() bool {
	if o != nil && !common.IsNil(o.Smallcp) {
		return true
	}

	return false
}

// SetCp gets a reference to the given bool and assigns it to the Cp field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallcp(v bool) {
	o.Smallcp = &v
}

// GetAP returns the AP field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetAP() string {
	if o == nil || common.IsNil(o.AP) {
		var ret string
		return ret
	}
	return *o.AP
}

// GetAPOk returns a tuple with the AP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetAPOk() (*string, bool) {
	if o == nil || common.IsNil(o.AP) {
		return nil, false
	}
	return o.AP, true
}

// HasAP returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasAP() bool {
	if o != nil && !common.IsNil(o.AP) {
		return true
	}

	return false
}

// SetAP gets a reference to the given string and assigns it to the AP field.
func (o *ConditionalOrderTradeUpdateSo) SetAP(v string) {
	o.AP = &v
}

// GetCr returns the Cr field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmallcr() string {
	if o == nil || common.IsNil(o.Smallcr) {
		var ret string
		return ret
	}
	return *o.Smallcr
}

// GetCrOk returns a tuple with the Cr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmallcrOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcr) {
		return nil, false
	}
	return o.Smallcr, true
}

// HasCr returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmallcr() bool {
	if o != nil && !common.IsNil(o.Smallcr) {
		return true
	}

	return false
}

// SetCr gets a reference to the given string and assigns it to the Cr field.
func (o *ConditionalOrderTradeUpdateSo) SetSmallcr(v string) {
	o.Smallcr = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *ConditionalOrderTradeUpdateSo) SetSmalli(v int64) {
	o.Smalli = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetV() string {
	if o == nil || common.IsNil(o.V) {
		var ret string
		return ret
	}
	return *o.V
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetVOk() (*string, bool) {
	if o == nil || common.IsNil(o.V) {
		return nil, false
	}
	return o.V, true
}

// HasV returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasV() bool {
	if o != nil && !common.IsNil(o.V) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *ConditionalOrderTradeUpdateSo) SetV(v string) {
	o.V = &v
}

// GetGtd returns the Gtd field value if set, zero value otherwise.
func (o *ConditionalOrderTradeUpdateSo) GetGtd() int64 {
	if o == nil || common.IsNil(o.Gtd) {
		var ret int64
		return ret
	}
	return *o.Gtd
}

// GetGtdOk returns a tuple with the Gtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalOrderTradeUpdateSo) GetGtdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Gtd) {
		return nil, false
	}
	return o.Gtd, true
}

// HasGtd returns a boolean if a field has been set.
func (o *ConditionalOrderTradeUpdateSo) HasGtd() bool {
	if o != nil && !common.IsNil(o.Gtd) {
		return true
	}

	return false
}

// SetGtd gets a reference to the given int64 and assigns it to the Gtd field.
func (o *ConditionalOrderTradeUpdateSo) SetGtd(v int64) {
	o.Gtd = &v
}

func (o ConditionalOrderTradeUpdateSo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalOrderTradeUpdateSo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallsi) {
		toSerialize["si"] = o.Smallsi
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.Smallst) {
		toSerialize["st"] = o.Smallst
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
	if !common.IsNil(o.Smallsp) {
		toSerialize["sp"] = o.Smallsp
	}
	if !common.IsNil(o.Smallos) {
		toSerialize["os"] = o.Smallos
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallut) {
		toSerialize["ut"] = o.Smallut
	}
	if !common.IsNil(o.R) {
		toSerialize["R"] = o.R
	}
	if !common.IsNil(o.Smallwt) {
		toSerialize["wt"] = o.Smallwt
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
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.V) {
		toSerialize["V"] = o.V
	}
	if !common.IsNil(o.Gtd) {
		toSerialize["gtd"] = o.Gtd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConditionalOrderTradeUpdateSo) UnmarshalJSON(data []byte) (err error) {
	varConditionalOrderTradeUpdateSo := _ConditionalOrderTradeUpdateSo{}

	err = json.Unmarshal(data, &varConditionalOrderTradeUpdateSo)

	if err != nil {
		return err
	}

	*o = ConditionalOrderTradeUpdateSo(varConditionalOrderTradeUpdateSo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "si")
		delete(additionalProperties, "S")
		delete(additionalProperties, "st")
		delete(additionalProperties, "f")
		delete(additionalProperties, "q")
		delete(additionalProperties, "p")
		delete(additionalProperties, "sp")
		delete(additionalProperties, "os")
		delete(additionalProperties, "T")
		delete(additionalProperties, "ut")
		delete(additionalProperties, "R")
		delete(additionalProperties, "wt")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "cp")
		delete(additionalProperties, "AP")
		delete(additionalProperties, "cr")
		delete(additionalProperties, "i")
		delete(additionalProperties, "V")
		delete(additionalProperties, "gtd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConditionalOrderTradeUpdateSo struct {
	value *ConditionalOrderTradeUpdateSo
	isSet bool
}

func (v NullableConditionalOrderTradeUpdateSo) Get() *ConditionalOrderTradeUpdateSo {
	return v.value
}

func (v *NullableConditionalOrderTradeUpdateSo) Set(val *ConditionalOrderTradeUpdateSo) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalOrderTradeUpdateSo) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalOrderTradeUpdateSo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalOrderTradeUpdateSo(val *ConditionalOrderTradeUpdateSo) *NullableConditionalOrderTradeUpdateSo {
	return &NullableConditionalOrderTradeUpdateSo{value: val, isSet: true}
}

func (v NullableConditionalOrderTradeUpdateSo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalOrderTradeUpdateSo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
