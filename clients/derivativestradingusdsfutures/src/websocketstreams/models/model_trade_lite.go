/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradeLite type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradeLite{}

// TradeLite struct for TradeLite
type TradeLite struct {
	E                    *int64  `json:"E,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallq               *string `json:"q,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smallm               *bool   `json:"m,omitempty"`
	Smallc               *string `json:"c,omitempty"`
	S                    *string `json:"S,omitempty"`
	L                    *string `json:"L,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	Smallt               *int64  `json:"t,omitempty"`
	Smalli               *int64  `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradeLite TradeLite

// NewTradeLite instantiates a new TradeLite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeLite() *TradeLite {
	this := TradeLite{}
	return &this
}

// NewTradeLiteWithDefaults instantiates a new TradeLite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeLiteWithDefaults() *TradeLite {
	this := TradeLite{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradeLite) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TradeLite) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *TradeLite) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradeLite) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TradeLite) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradeLite) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TradeLite) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *TradeLite) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TradeLite) SetSmalls(v string) {
	o.Smalls = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TradeLite) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *TradeLite) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TradeLite) SetSmallq(v string) {
	o.Smallq = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *TradeLite) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *TradeLite) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *TradeLite) SetSmallp(v string) {
	o.Smallp = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *TradeLite) GetSmallm() bool {
	if o == nil || common.IsNil(o.Smallm) {
		var ret bool
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmallmOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *TradeLite) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given bool and assigns it to the M field.
func (o *TradeLite) SetSmallm(v bool) {
	o.Smallm = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *TradeLite) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *TradeLite) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *TradeLite) SetSmallc(v string) {
	o.Smallc = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TradeLite) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *TradeLite) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TradeLite) SetS(v string) {
	o.S = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *TradeLite) GetL() string {
	if o == nil || common.IsNil(o.L) {
		var ret string
		return ret
	}
	return *o.L
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetLOk() (*string, bool) {
	if o == nil || common.IsNil(o.L) {
		return nil, false
	}
	return o.L, true
}

// HasL returns a boolean if a field has been set.
func (o *TradeLite) HasL() bool {
	if o != nil && !common.IsNil(o.L) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *TradeLite) SetL(v string) {
	o.L = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *TradeLite) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *TradeLite) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *TradeLite) SetSmalll(v string) {
	o.Smalll = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradeLite) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *TradeLite) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradeLite) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *TradeLite) GetSmalli() int64 {
	if o == nil || common.IsNil(o.Smalli) {
		var ret int64
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeLite) GetSmalliOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *TradeLite) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given int64 and assigns it to the I field.
func (o *TradeLite) SetSmalli(v int64) {
	o.Smalli = &v
}

func (o TradeLite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradeLite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}
	if !common.IsNil(o.L) {
		toSerialize["L"] = o.L
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradeLite) UnmarshalJSON(data []byte) (err error) {
	varTradeLite := _TradeLite{}

	err = json.Unmarshal(data, &varTradeLite)

	if err != nil {
		return err
	}

	*o = TradeLite(varTradeLite)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "q")
		delete(additionalProperties, "p")
		delete(additionalProperties, "m")
		delete(additionalProperties, "c")
		delete(additionalProperties, "S")
		delete(additionalProperties, "L")
		delete(additionalProperties, "l")
		delete(additionalProperties, "t")
		delete(additionalProperties, "i")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradeLite struct {
	value *TradeLite
	isSet bool
}

func (v NullableTradeLite) Get() *TradeLite {
	return v.value
}

func (v *NullableTradeLite) Set(val *TradeLite) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeLite) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeLite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeLite(val *TradeLite) *NullableTradeLite {
	return &NullableTradeLite{value: val, isSet: true}
}

func (v NullableTradeLite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeLite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
