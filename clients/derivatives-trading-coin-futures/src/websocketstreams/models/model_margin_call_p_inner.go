/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarginCallPInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginCallPInner{}

// MarginCallPInner struct for MarginCallPInner
type MarginCallPInner struct {
	Smalls               *string `json:"s,omitempty"`
	Smallps              *string `json:"ps,omitempty"`
	Smallpa              *string `json:"pa,omitempty"`
	Smallmt              *string `json:"mt,omitempty"`
	Smalliw              *string `json:"iw,omitempty"`
	Smallmp              *string `json:"mp,omitempty"`
	Smallup              *string `json:"up,omitempty"`
	Smallmm              *string `json:"mm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCallPInner MarginCallPInner

// NewMarginCallPInner instantiates a new MarginCallPInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCallPInner() *MarginCallPInner {
	this := MarginCallPInner{}
	return &this
}

// NewMarginCallPInnerWithDefaults instantiates a new MarginCallPInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCallPInnerWithDefaults() *MarginCallPInner {
	this := MarginCallPInner{}
	return &this
}

// GetS returns the S field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *MarginCallPInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *MarginCallPInner) SetSmallps(v string) {
	o.Smallps = &v
}

// GetPa returns the Pa field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallpa() string {
	if o == nil || common.IsNil(o.Smallpa) {
		var ret string
		return ret
	}
	return *o.Smallpa
}

// GetPaOk returns a tuple with the Pa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallpaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallpa) {
		return nil, false
	}
	return o.Smallpa, true
}

// HasPa returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallpa() bool {
	if o != nil && !common.IsNil(o.Smallpa) {
		return true
	}

	return false
}

// SetPa gets a reference to the given string and assigns it to the Pa field.
func (o *MarginCallPInner) SetSmallpa(v string) {
	o.Smallpa = &v
}

// GetMt returns the Mt field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallmt() string {
	if o == nil || common.IsNil(o.Smallmt) {
		var ret string
		return ret
	}
	return *o.Smallmt
}

// GetMtOk returns a tuple with the Mt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallmtOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmt) {
		return nil, false
	}
	return o.Smallmt, true
}

// HasMt returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallmt() bool {
	if o != nil && !common.IsNil(o.Smallmt) {
		return true
	}

	return false
}

// SetMt gets a reference to the given string and assigns it to the Mt field.
func (o *MarginCallPInner) SetSmallmt(v string) {
	o.Smallmt = &v
}

// GetIw returns the Iw field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmalliw() string {
	if o == nil || common.IsNil(o.Smalliw) {
		var ret string
		return ret
	}
	return *o.Smalliw
}

// GetIwOk returns a tuple with the Iw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmalliwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalliw) {
		return nil, false
	}
	return o.Smalliw, true
}

// HasIw returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmalliw() bool {
	if o != nil && !common.IsNil(o.Smalliw) {
		return true
	}

	return false
}

// SetIw gets a reference to the given string and assigns it to the Iw field.
func (o *MarginCallPInner) SetSmalliw(v string) {
	o.Smalliw = &v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallmp() string {
	if o == nil || common.IsNil(o.Smallmp) {
		var ret string
		return ret
	}
	return *o.Smallmp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallmpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmp) {
		return nil, false
	}
	return o.Smallmp, true
}

// HasMp returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallmp() bool {
	if o != nil && !common.IsNil(o.Smallmp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given string and assigns it to the Mp field.
func (o *MarginCallPInner) SetSmallmp(v string) {
	o.Smallmp = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallup() string {
	if o == nil || common.IsNil(o.Smallup) {
		var ret string
		return ret
	}
	return *o.Smallup
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallupOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallup) {
		return nil, false
	}
	return o.Smallup, true
}

// HasUp returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallup() bool {
	if o != nil && !common.IsNil(o.Smallup) {
		return true
	}

	return false
}

// SetUp gets a reference to the given string and assigns it to the Up field.
func (o *MarginCallPInner) SetSmallup(v string) {
	o.Smallup = &v
}

// GetMm returns the Mm field value if set, zero value otherwise.
func (o *MarginCallPInner) GetSmallmm() string {
	if o == nil || common.IsNil(o.Smallmm) {
		var ret string
		return ret
	}
	return *o.Smallmm
}

// GetMmOk returns a tuple with the Mm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCallPInner) GetSmallmmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallmm) {
		return nil, false
	}
	return o.Smallmm, true
}

// HasMm returns a boolean if a field has been set.
func (o *MarginCallPInner) HasSmallmm() bool {
	if o != nil && !common.IsNil(o.Smallmm) {
		return true
	}

	return false
}

// SetMm gets a reference to the given string and assigns it to the Mm field.
func (o *MarginCallPInner) SetSmallmm(v string) {
	o.Smallmm = &v
}

func (o MarginCallPInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCallPInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallpa) {
		toSerialize["pa"] = o.Smallpa
	}
	if !common.IsNil(o.Smallmt) {
		toSerialize["mt"] = o.Smallmt
	}
	if !common.IsNil(o.Smalliw) {
		toSerialize["iw"] = o.Smalliw
	}
	if !common.IsNil(o.Smallmp) {
		toSerialize["mp"] = o.Smallmp
	}
	if !common.IsNil(o.Smallup) {
		toSerialize["up"] = o.Smallup
	}
	if !common.IsNil(o.Smallmm) {
		toSerialize["mm"] = o.Smallmm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginCallPInner) UnmarshalJSON(data []byte) (err error) {
	varMarginCallPInner := _MarginCallPInner{}

	err = json.Unmarshal(data, &varMarginCallPInner)

	if err != nil {
		return err
	}

	*o = MarginCallPInner(varMarginCallPInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "pa")
		delete(additionalProperties, "mt")
		delete(additionalProperties, "iw")
		delete(additionalProperties, "mp")
		delete(additionalProperties, "up")
		delete(additionalProperties, "mm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCallPInner struct {
	value *MarginCallPInner
	isSet bool
}

func (v NullableMarginCallPInner) Get() *MarginCallPInner {
	return v.value
}

func (v *NullableMarginCallPInner) Set(val *MarginCallPInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCallPInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCallPInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCallPInner(val *MarginCallPInner) *NullableMarginCallPInner {
	return &NullableMarginCallPInner{value: val, isSet: true}
}

func (v NullableMarginCallPInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCallPInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
