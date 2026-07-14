/*
Options WebSocket Market Streams

Access market data, manage accounts, and trade Binance Options.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OptionMarkPriceResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OptionMarkPriceResponseInner{}

// OptionMarkPriceResponseInner struct for OptionMarkPriceResponseInner
type OptionMarkPriceResponseInner struct {
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Mark price
	Smallmp *string `json:"mp,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Index price
	Smalli *string `json:"i,omitempty"`
	// Estimated Settle Price, only useful in the 0.5 hour before the settlement starts
	P *string `json:"P,omitempty"`
	// The best buy price
	Smallbo *string `json:"bo,omitempty"`
	// The best sell price
	Smallao *string `json:"ao,omitempty"`
	// The best buy quantity
	Smallbq *string `json:"bq,omitempty"`
	// The best sell quantity
	Smallaq *string `json:"aq,omitempty"`
	// BuyImplied volatility
	Smallb *string `json:"b,omitempty"`
	// SellImplied volatility
	Smalla *string `json:"a,omitempty"`
	// Buy Maximum price
	Smallhl *string `json:"hl,omitempty"`
	// Sell Minimum price
	Smallll *string `json:"ll,omitempty"`
	// volatility
	Smallvo *string `json:"vo,omitempty"`
	// risk free rate
	Smallrf *string `json:"rf,omitempty"`
	// delta
	Smalld *string `json:"d,omitempty"`
	// theta
	Smallt *string `json:"t,omitempty"`
	// gamma
	Smallg *string `json:"g,omitempty"`
	// vega
	Smallv               *string `json:"v,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptionMarkPriceResponseInner OptionMarkPriceResponseInner

// NewOptionMarkPriceResponseInner instantiates a new OptionMarkPriceResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionMarkPriceResponseInner() *OptionMarkPriceResponseInner {
	this := OptionMarkPriceResponseInner{}
	return &this
}

// NewOptionMarkPriceResponseInnerWithDefaults instantiates a new OptionMarkPriceResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionMarkPriceResponseInnerWithDefaults() *OptionMarkPriceResponseInner {
	this := OptionMarkPriceResponseInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *OptionMarkPriceResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *OptionMarkPriceResponseInner) SetSmallmp(v string) {
	o.Smallmp = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *OptionMarkPriceResponseInner) SetE(v int64) {
	o.E = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *OptionMarkPriceResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *OptionMarkPriceResponseInner) SetSmalli(v string) {
	o.Smalli = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *OptionMarkPriceResponseInner) SetP(v string) {
	o.P = &v
}

// GetBo returns the Bo field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallbo() string {
	if o == nil || common.IsNil(o.Smallbo) {
		var ret string
		return ret
	}
	return *o.Smallbo
}

// GetBoOk returns a tuple with the Bo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallboOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbo) {
		return nil, false
	}
	return o.Smallbo, true
}

// HasBo returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallbo() bool {
	if o != nil && !common.IsNil(o.Smallbo) {
		return true
	}

	return false
}

// SetBo gets a reference to the given string and assigns it to the Bo field.
func (o *OptionMarkPriceResponseInner) SetSmallbo(v string) {
	o.Smallbo = &v
}

// GetAo returns the Ao field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallao() string {
	if o == nil || common.IsNil(o.Smallao) {
		var ret string
		return ret
	}
	return *o.Smallao
}

// GetAoOk returns a tuple with the Ao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallaoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallao) {
		return nil, false
	}
	return o.Smallao, true
}

// HasAo returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallao() bool {
	if o != nil && !common.IsNil(o.Smallao) {
		return true
	}

	return false
}

// SetAo gets a reference to the given string and assigns it to the Ao field.
func (o *OptionMarkPriceResponseInner) SetSmallao(v string) {
	o.Smallao = &v
}

// GetBq returns the Bq field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallbq() string {
	if o == nil || common.IsNil(o.Smallbq) {
		var ret string
		return ret
	}
	return *o.Smallbq
}

// GetBqOk returns a tuple with the Bq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallbqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallbq) {
		return nil, false
	}
	return o.Smallbq, true
}

// HasBq returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallbq() bool {
	if o != nil && !common.IsNil(o.Smallbq) {
		return true
	}

	return false
}

// SetBq gets a reference to the given string and assigns it to the Bq field.
func (o *OptionMarkPriceResponseInner) SetSmallbq(v string) {
	o.Smallbq = &v
}

// GetAq returns the Aq field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallaq() string {
	if o == nil || common.IsNil(o.Smallaq) {
		var ret string
		return ret
	}
	return *o.Smallaq
}

// GetAqOk returns a tuple with the Aq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallaqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallaq) {
		return nil, false
	}
	return o.Smallaq, true
}

// HasAq returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallaq() bool {
	if o != nil && !common.IsNil(o.Smallaq) {
		return true
	}

	return false
}

// SetAq gets a reference to the given string and assigns it to the Aq field.
func (o *OptionMarkPriceResponseInner) SetSmallaq(v string) {
	o.Smallaq = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *OptionMarkPriceResponseInner) SetSmallb(v string) {
	o.Smallb = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *OptionMarkPriceResponseInner) SetSmalla(v string) {
	o.Smalla = &v
}

// GetHl returns the Hl field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallhl() string {
	if o == nil || common.IsNil(o.Smallhl) {
		var ret string
		return ret
	}
	return *o.Smallhl
}

// GetHlOk returns a tuple with the Hl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallhlOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallhl) {
		return nil, false
	}
	return o.Smallhl, true
}

// HasHl returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallhl() bool {
	if o != nil && !common.IsNil(o.Smallhl) {
		return true
	}

	return false
}

// SetHl gets a reference to the given string and assigns it to the Hl field.
func (o *OptionMarkPriceResponseInner) SetSmallhl(v string) {
	o.Smallhl = &v
}

// GetLl returns the Ll field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallll() string {
	if o == nil || common.IsNil(o.Smallll) {
		var ret string
		return ret
	}
	return *o.Smallll
}

// GetLlOk returns a tuple with the Ll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallll) {
		return nil, false
	}
	return o.Smallll, true
}

// HasLl returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallll() bool {
	if o != nil && !common.IsNil(o.Smallll) {
		return true
	}

	return false
}

// SetLl gets a reference to the given string and assigns it to the Ll field.
func (o *OptionMarkPriceResponseInner) SetSmallll(v string) {
	o.Smallll = &v
}

// GetVo returns the Vo field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallvo() string {
	if o == nil || common.IsNil(o.Smallvo) {
		var ret string
		return ret
	}
	return *o.Smallvo
}

// GetVoOk returns a tuple with the Vo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallvoOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallvo) {
		return nil, false
	}
	return o.Smallvo, true
}

// HasVo returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallvo() bool {
	if o != nil && !common.IsNil(o.Smallvo) {
		return true
	}

	return false
}

// SetVo gets a reference to the given string and assigns it to the Vo field.
func (o *OptionMarkPriceResponseInner) SetSmallvo(v string) {
	o.Smallvo = &v
}

// GetRf returns the Rf field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallrf() string {
	if o == nil || common.IsNil(o.Smallrf) {
		var ret string
		return ret
	}
	return *o.Smallrf
}

// GetRfOk returns a tuple with the Rf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallrfOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallrf) {
		return nil, false
	}
	return o.Smallrf, true
}

// HasRf returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallrf() bool {
	if o != nil && !common.IsNil(o.Smallrf) {
		return true
	}

	return false
}

// SetRf gets a reference to the given string and assigns it to the Rf field.
func (o *OptionMarkPriceResponseInner) SetSmallrf(v string) {
	o.Smallrf = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmalld() string {
	if o == nil || common.IsNil(o.Smalld) {
		var ret string
		return ret
	}
	return *o.Smalld
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmalldOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalld) {
		return nil, false
	}
	return o.Smalld, true
}

// HasD returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmalld() bool {
	if o != nil && !common.IsNil(o.Smalld) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *OptionMarkPriceResponseInner) SetSmalld(v string) {
	o.Smalld = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *OptionMarkPriceResponseInner) SetSmallt(v string) {
	o.Smallt = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallg() string {
	if o == nil || common.IsNil(o.Smallg) {
		var ret string
		return ret
	}
	return *o.Smallg
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallg) {
		return nil, false
	}
	return o.Smallg, true
}

// HasG returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallg() bool {
	if o != nil && !common.IsNil(o.Smallg) {
		return true
	}

	return false
}

// SetG gets a reference to the given string and assigns it to the G field.
func (o *OptionMarkPriceResponseInner) SetSmallg(v string) {
	o.Smallg = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *OptionMarkPriceResponseInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionMarkPriceResponseInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *OptionMarkPriceResponseInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *OptionMarkPriceResponseInner) SetSmallv(v string) {
	o.Smallv = &v
}

func (o OptionMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionMarkPriceResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *OptionMarkPriceResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOptionMarkPriceResponseInner := _OptionMarkPriceResponseInner{}

	err = json.Unmarshal(data, &varOptionMarkPriceResponseInner)

	if err != nil {
		return err
	}

	*o = OptionMarkPriceResponseInner(varOptionMarkPriceResponseInner)

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

type NullableOptionMarkPriceResponseInner struct {
	value *OptionMarkPriceResponseInner
	isSet bool
}

func (v NullableOptionMarkPriceResponseInner) Get() *OptionMarkPriceResponseInner {
	return v.value
}

func (v *NullableOptionMarkPriceResponseInner) Set(val *OptionMarkPriceResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionMarkPriceResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionMarkPriceResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionMarkPriceResponseInner(val *OptionMarkPriceResponseInner) *NullableOptionMarkPriceResponseInner {
	return &NullableOptionMarkPriceResponseInner{value: val, isSet: true}
}

func (v NullableOptionMarkPriceResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionMarkPriceResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
