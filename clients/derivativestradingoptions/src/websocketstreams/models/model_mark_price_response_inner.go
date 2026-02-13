/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarkPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceResponseInner{}

// MarkPriceResponseInner struct for MarkPriceResponseInner
type MarkPriceResponseInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smallmp              *string `json:"mp,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalle               *string `json:"e,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	P                    *string `json:"P,omitempty"`
	Smallbo              *string `json:"bo,omitempty"`
	Smallao              *string `json:"ao,omitempty"`
	Smallbq              *string `json:"bq,omitempty"`
	Smallaq              *string `json:"aq,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smallhl              *string `json:"hl,omitempty"`
	Smallll              *string `json:"ll,omitempty"`
	Smallvo              *string `json:"vo,omitempty"`
	Smallrf              *string `json:"rf,omitempty"`
	Smalld               *string `json:"d,omitempty"`
	Smallt               *string `json:"t,omitempty"`
	Smallg               *string `json:"g,omitempty"`
	Smallv               *string `json:"v,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarkPriceResponseInner MarkPriceResponseInner

// NewMarkPriceResponseInner instantiates a new MarkPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceResponseInner() *MarkPriceResponseInner {
	this := MarkPriceResponseInner{}
	return &this
}

// NewMarkPriceResponseInnerWithDefaults instantiates a new MarkPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceResponseInnerWithDefaults() *MarkPriceResponseInner {
	this := MarkPriceResponseInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MarkPriceResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *MarkPriceResponseInner) SetSmallmp(v string) {
	o.Smallmp = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarkPriceResponseInner) SetE(v int64) {
	o.E = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MarkPriceResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *MarkPriceResponseInner) SetSmalli(v string) {
	o.Smalli = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *MarkPriceResponseInner) SetP(v string) {
	o.P = &v
}

// GetBo returns the Bo field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallbo() string {
	if o == nil || common.IsNil(o.Smallbo) {
		var ret string
		return ret
	}
	return *o.Smallbo
}

// GetBoOk returns a tuple with the Bo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallboOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbo) {
		return nil, false
	}
	return o.Smallbo, true
}

// HasBo returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallbo() bool {
	if o != nil && !common.IsNil(o.Smallbo) {
		return true
	}

	return false
}

// SetBo gets a reference to the given string and assigns it to the Bo field.
func (o *MarkPriceResponseInner) SetSmallbo(v string) {
	o.Smallbo = &v
}

// GetAo returns the Ao field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallao() string {
	if o == nil || common.IsNil(o.Smallao) {
		var ret string
		return ret
	}
	return *o.Smallao
}

// GetAoOk returns a tuple with the Ao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallaoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallao) {
		return nil, false
	}
	return o.Smallao, true
}

// HasAo returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallao() bool {
	if o != nil && !common.IsNil(o.Smallao) {
		return true
	}

	return false
}

// SetAo gets a reference to the given string and assigns it to the Ao field.
func (o *MarkPriceResponseInner) SetSmallao(v string) {
	o.Smallao = &v
}

// GetBq returns the Bq field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallbq() string {
	if o == nil || common.IsNil(o.Smallbq) {
		var ret string
		return ret
	}
	return *o.Smallbq
}

// GetBqOk returns a tuple with the Bq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallbqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbq) {
		return nil, false
	}
	return o.Smallbq, true
}

// HasBq returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallbq() bool {
	if o != nil && !common.IsNil(o.Smallbq) {
		return true
	}

	return false
}

// SetBq gets a reference to the given string and assigns it to the Bq field.
func (o *MarkPriceResponseInner) SetSmallbq(v string) {
	o.Smallbq = &v
}

// GetAq returns the Aq field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallaq() string {
	if o == nil || common.IsNil(o.Smallaq) {
		var ret string
		return ret
	}
	return *o.Smallaq
}

// GetAqOk returns a tuple with the Aq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallaqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallaq) {
		return nil, false
	}
	return o.Smallaq, true
}

// HasAq returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallaq() bool {
	if o != nil && !common.IsNil(o.Smallaq) {
		return true
	}

	return false
}

// SetAq gets a reference to the given string and assigns it to the Aq field.
func (o *MarkPriceResponseInner) SetSmallaq(v string) {
	o.Smallaq = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *MarkPriceResponseInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *MarkPriceResponseInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetHl returns the Hl field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallhl() string {
	if o == nil || common.IsNil(o.Smallhl) {
		var ret string
		return ret
	}
	return *o.Smallhl
}

// GetHlOk returns a tuple with the Hl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallhlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallhl) {
		return nil, false
	}
	return o.Smallhl, true
}

// HasHl returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallhl() bool {
	if o != nil && !common.IsNil(o.Smallhl) {
		return true
	}

	return false
}

// SetHl gets a reference to the given string and assigns it to the Hl field.
func (o *MarkPriceResponseInner) SetSmallhl(v string) {
	o.Smallhl = &v
}

// GetLl returns the Ll field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallll() string {
	if o == nil || common.IsNil(o.Smallll) {
		var ret string
		return ret
	}
	return *o.Smallll
}

// GetLlOk returns a tuple with the Ll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallll) {
		return nil, false
	}
	return o.Smallll, true
}

// HasLl returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallll() bool {
	if o != nil && !common.IsNil(o.Smallll) {
		return true
	}

	return false
}

// SetLl gets a reference to the given string and assigns it to the Ll field.
func (o *MarkPriceResponseInner) SetSmallll(v string) {
	o.Smallll = &v
}

// GetVo returns the Vo field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallvo() string {
	if o == nil || common.IsNil(o.Smallvo) {
		var ret string
		return ret
	}
	return *o.Smallvo
}

// GetVoOk returns a tuple with the Vo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallvoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallvo) {
		return nil, false
	}
	return o.Smallvo, true
}

// HasVo returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallvo() bool {
	if o != nil && !common.IsNil(o.Smallvo) {
		return true
	}

	return false
}

// SetVo gets a reference to the given string and assigns it to the Vo field.
func (o *MarkPriceResponseInner) SetSmallvo(v string) {
	o.Smallvo = &v
}

// GetRf returns the Rf field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallrf() string {
	if o == nil || common.IsNil(o.Smallrf) {
		var ret string
		return ret
	}
	return *o.Smallrf
}

// GetRfOk returns a tuple with the Rf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallrfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallrf) {
		return nil, false
	}
	return o.Smallrf, true
}

// HasRf returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallrf() bool {
	if o != nil && !common.IsNil(o.Smallrf) {
		return true
	}

	return false
}

// SetRf gets a reference to the given string and assigns it to the Rf field.
func (o *MarkPriceResponseInner) SetSmallrf(v string) {
	o.Smallrf = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *MarkPriceResponseInner) SetSmalld(v string) {
	o.Smalld = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *MarkPriceResponseInner) SetSmallt(v string) {
	o.Smallt = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallg() string {
	if o == nil || common.IsNil(o.Smallg) {
		var ret string
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *MarkPriceResponseInner) SetSmallg(v string) {
	o.Smallg = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *MarkPriceResponseInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceResponseInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *MarkPriceResponseInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *MarkPriceResponseInner) SetSmallv(v string) {
	o.Smallv = &v
}

func (o MarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallmp) {
		toSerialize["mp"] = o.Smallmp
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
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
	if !common.IsNil(o.Smallhl) {
		toSerialize["hl"] = o.Smallhl
	}
	if !common.IsNil(o.Smallll) {
		toSerialize["ll"] = o.Smallll
	}
	if !common.IsNil(o.Smallvo) {
		toSerialize["vo"] = o.Smallvo
	}
	if !common.IsNil(o.Smallrf) {
		toSerialize["rf"] = o.Smallrf
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarkPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varMarkPriceResponseInner := _MarkPriceResponseInner{}

	err = json.Unmarshal(data, &varMarkPriceResponseInner)

	if err != nil {
		return err
	}

	*o = MarkPriceResponseInner(varMarkPriceResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "mp")
		delete(additionalProperties, "E")
		delete(additionalProperties, "e")
		delete(additionalProperties, "i")
		delete(additionalProperties, "P")
		delete(additionalProperties, "bo")
		delete(additionalProperties, "ao")
		delete(additionalProperties, "bq")
		delete(additionalProperties, "aq")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		delete(additionalProperties, "hl")
		delete(additionalProperties, "ll")
		delete(additionalProperties, "vo")
		delete(additionalProperties, "rf")
		delete(additionalProperties, "d")
		delete(additionalProperties, "t")
		delete(additionalProperties, "g")
		delete(additionalProperties, "v")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarkPriceResponseInner struct {
	value *MarkPriceResponseInner
	isSet bool
}

func (v NullableMarkPriceResponseInner) Get() *MarkPriceResponseInner {
	return v.value
}

func (v *NullableMarkPriceResponseInner) Set(val *MarkPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceResponseInner(val *MarkPriceResponseInner) *NullableMarkPriceResponseInner {
	return &NullableMarkPriceResponseInner{value: val, isSet: true}
}

func (v NullableMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
