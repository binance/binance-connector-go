/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllTokens24hTickerStreamResponseDInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllTokens24hTickerStreamResponseDInner{}

// AllTokens24hTickerStreamResponseDInner struct for AllTokens24hTickerStreamResponseDInner
type AllTokens24hTickerStreamResponseDInner struct {
	// Contract address@chain ID
	Ca *string `json:"ca,omitempty"`
	// Number of trades in the last 24 hours
	Cnt24 *int64 `json:"cnt24,omitempty"`
	// Fully diluted valuation
	Fdv *string `json:"fdv,omitempty"`
	// Number of holders
	Hc *string `json:"hc,omitempty"`
	// Liquidity
	Liq *string `json:"liq,omitempty"`
	// Market cap
	Mc *string `json:"mc,omitempty"`
	// Current price
	P *string `json:"p,omitempty"`
	// 24-hour price change percent
	Pc24 *string `json:"pc24,omitempty"`
	// Token short identifier
	S *string `json:"s,omitempty"`
	// Event timestamp in milliseconds
	T *int64 `json:"t,omitempty"`
	// 24-hour volume
	Vol24                *string `json:"vol24,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllTokens24hTickerStreamResponseDInner AllTokens24hTickerStreamResponseDInner

// NewAllTokens24hTickerStreamResponseDInner instantiates a new AllTokens24hTickerStreamResponseDInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllTokens24hTickerStreamResponseDInner() *AllTokens24hTickerStreamResponseDInner {
	this := AllTokens24hTickerStreamResponseDInner{}
	return &this
}

// NewAllTokens24hTickerStreamResponseDInnerWithDefaults instantiates a new AllTokens24hTickerStreamResponseDInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllTokens24hTickerStreamResponseDInnerWithDefaults() *AllTokens24hTickerStreamResponseDInner {
	this := AllTokens24hTickerStreamResponseDInner{}
	return &this
}

// GetCa returns the Ca field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetCa() string {
	if o == nil || common.IsNil(o.Ca) {
		var ret string
		return ret
	}
	return *o.Ca
}

// GetCaOk returns a tuple with the Ca field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetCaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Ca) {
		return nil, false
	}
	return o.Ca, true
}

// HasCa returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasCa() bool {
	if o != nil && !common.IsNil(o.Ca) {
		return true
	}

	return false
}

// SetCa gets a reference to the given string and assigns it to the Ca field.
func (o *AllTokens24hTickerStreamResponseDInner) SetCa(v string) {
	o.Ca = &v
}

// GetCnt24 returns the Cnt24 field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetCnt24() int64 {
	if o == nil || common.IsNil(o.Cnt24) {
		var ret int64
		return ret
	}
	return *o.Cnt24
}

// GetCnt24Ok returns a tuple with the Cnt24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetCnt24Ok() (*int64, bool) {
	if o == nil || common.IsNil(o.Cnt24) {
		return nil, false
	}
	return o.Cnt24, true
}

// HasCnt24 returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasCnt24() bool {
	if o != nil && !common.IsNil(o.Cnt24) {
		return true
	}

	return false
}

// SetCnt24 gets a reference to the given int64 and assigns it to the Cnt24 field.
func (o *AllTokens24hTickerStreamResponseDInner) SetCnt24(v int64) {
	o.Cnt24 = &v
}

// GetFdv returns the Fdv field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetFdv() string {
	if o == nil || common.IsNil(o.Fdv) {
		var ret string
		return ret
	}
	return *o.Fdv
}

// GetFdvOk returns a tuple with the Fdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetFdvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fdv) {
		return nil, false
	}
	return o.Fdv, true
}

// HasFdv returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasFdv() bool {
	if o != nil && !common.IsNil(o.Fdv) {
		return true
	}

	return false
}

// SetFdv gets a reference to the given string and assigns it to the Fdv field.
func (o *AllTokens24hTickerStreamResponseDInner) SetFdv(v string) {
	o.Fdv = &v
}

// GetHc returns the Hc field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetHc() string {
	if o == nil || common.IsNil(o.Hc) {
		var ret string
		return ret
	}
	return *o.Hc
}

// GetHcOk returns a tuple with the Hc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetHcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Hc) {
		return nil, false
	}
	return o.Hc, true
}

// HasHc returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasHc() bool {
	if o != nil && !common.IsNil(o.Hc) {
		return true
	}

	return false
}

// SetHc gets a reference to the given string and assigns it to the Hc field.
func (o *AllTokens24hTickerStreamResponseDInner) SetHc(v string) {
	o.Hc = &v
}

// GetLiq returns the Liq field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetLiq() string {
	if o == nil || common.IsNil(o.Liq) {
		var ret string
		return ret
	}
	return *o.Liq
}

// GetLiqOk returns a tuple with the Liq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetLiqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liq) {
		return nil, false
	}
	return o.Liq, true
}

// HasLiq returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasLiq() bool {
	if o != nil && !common.IsNil(o.Liq) {
		return true
	}

	return false
}

// SetLiq gets a reference to the given string and assigns it to the Liq field.
func (o *AllTokens24hTickerStreamResponseDInner) SetLiq(v string) {
	o.Liq = &v
}

// GetMc returns the Mc field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetMc() string {
	if o == nil || common.IsNil(o.Mc) {
		var ret string
		return ret
	}
	return *o.Mc
}

// GetMcOk returns a tuple with the Mc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetMcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Mc) {
		return nil, false
	}
	return o.Mc, true
}

// HasMc returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasMc() bool {
	if o != nil && !common.IsNil(o.Mc) {
		return true
	}

	return false
}

// SetMc gets a reference to the given string and assigns it to the Mc field.
func (o *AllTokens24hTickerStreamResponseDInner) SetMc(v string) {
	o.Mc = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetP() string {
	if o == nil || common.IsNil(o.P) {
		var ret string
		return ret
	}
	return *o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetPOk() (*string, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *AllTokens24hTickerStreamResponseDInner) SetP(v string) {
	o.P = &v
}

// GetPc24 returns the Pc24 field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetPc24() string {
	if o == nil || common.IsNil(o.Pc24) {
		var ret string
		return ret
	}
	return *o.Pc24
}

// GetPc24Ok returns a tuple with the Pc24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetPc24Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Pc24) {
		return nil, false
	}
	return o.Pc24, true
}

// HasPc24 returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasPc24() bool {
	if o != nil && !common.IsNil(o.Pc24) {
		return true
	}

	return false
}

// SetPc24 gets a reference to the given string and assigns it to the Pc24 field.
func (o *AllTokens24hTickerStreamResponseDInner) SetPc24(v string) {
	o.Pc24 = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllTokens24hTickerStreamResponseDInner) SetS(v string) {
	o.S = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *AllTokens24hTickerStreamResponseDInner) SetT(v int64) {
	o.T = &v
}

// GetVol24 returns the Vol24 field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponseDInner) GetVol24() string {
	if o == nil || common.IsNil(o.Vol24) {
		var ret string
		return ret
	}
	return *o.Vol24
}

// GetVol24Ok returns a tuple with the Vol24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponseDInner) GetVol24Ok() (*string, bool) {
	if o == nil || common.IsNil(o.Vol24) {
		return nil, false
	}
	return o.Vol24, true
}

// HasVol24 returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponseDInner) HasVol24() bool {
	if o != nil && !common.IsNil(o.Vol24) {
		return true
	}

	return false
}

// SetVol24 gets a reference to the given string and assigns it to the Vol24 field.
func (o *AllTokens24hTickerStreamResponseDInner) SetVol24(v string) {
	o.Vol24 = &v
}

func (o AllTokens24hTickerStreamResponseDInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllTokens24hTickerStreamResponseDInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Ca) {
		toSerialize["ca"] = o.Ca
	}
	if !common.IsNil(o.Cnt24) {
		toSerialize["cnt24"] = o.Cnt24
	}
	if !common.IsNil(o.Fdv) {
		toSerialize["fdv"] = o.Fdv
	}
	if !common.IsNil(o.Hc) {
		toSerialize["hc"] = o.Hc
	}
	if !common.IsNil(o.Liq) {
		toSerialize["liq"] = o.Liq
	}
	if !common.IsNil(o.Mc) {
		toSerialize["mc"] = o.Mc
	}
	if !common.IsNil(o.P) {
		toSerialize["p"] = o.P
	}
	if !common.IsNil(o.Pc24) {
		toSerialize["pc24"] = o.Pc24
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.T) {
		toSerialize["t"] = o.T
	}
	if !common.IsNil(o.Vol24) {
		toSerialize["vol24"] = o.Vol24
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllTokens24hTickerStreamResponseDInner) UnmarshalJSON(data []byte) (err error) {
	varAllTokens24hTickerStreamResponseDInner := _AllTokens24hTickerStreamResponseDInner{}

	err = json.Unmarshal(data, &varAllTokens24hTickerStreamResponseDInner)

	if err != nil {
		return err
	}

	*o = AllTokens24hTickerStreamResponseDInner(varAllTokens24hTickerStreamResponseDInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ca")
		delete(additionalProperties, "cnt24")
		delete(additionalProperties, "fdv")
		delete(additionalProperties, "hc")
		delete(additionalProperties, "liq")
		delete(additionalProperties, "mc")
		delete(additionalProperties, "p")
		delete(additionalProperties, "pc24")
		delete(additionalProperties, "s")
		delete(additionalProperties, "t")
		delete(additionalProperties, "vol24")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllTokens24hTickerStreamResponseDInner struct {
	value *AllTokens24hTickerStreamResponseDInner
	isSet bool
}

func (v NullableAllTokens24hTickerStreamResponseDInner) Get() *AllTokens24hTickerStreamResponseDInner {
	return v.value
}

func (v *NullableAllTokens24hTickerStreamResponseDInner) Set(val *AllTokens24hTickerStreamResponseDInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllTokens24hTickerStreamResponseDInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllTokens24hTickerStreamResponseDInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllTokens24hTickerStreamResponseDInner(val *AllTokens24hTickerStreamResponseDInner) *NullableAllTokens24hTickerStreamResponseDInner {
	return &NullableAllTokens24hTickerStreamResponseDInner{value: val, isSet: true}
}

func (v NullableAllTokens24hTickerStreamResponseDInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllTokens24hTickerStreamResponseDInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
