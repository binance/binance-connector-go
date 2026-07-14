/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllMiniTickerResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMiniTickerResponseInner{}

// AllMiniTickerResponseInner struct for AllMiniTickerResponseInner
type AllMiniTickerResponseInner struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// Close price
	Smallc *string `json:"c,omitempty"`
	// Open price
	Smallo *string `json:"o,omitempty"`
	// High price
	Smallh *string `json:"h,omitempty"`
	// Low price
	Smalll *string `json:"l,omitempty"`
	// Total traded base asset volume
	Smallv *string `json:"v,omitempty"`
	// Total traded quote asset volume
	Smallq               *string `json:"q,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMiniTickerResponseInner AllMiniTickerResponseInner

// NewAllMiniTickerResponseInner instantiates a new AllMiniTickerResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMiniTickerResponseInner() *AllMiniTickerResponseInner {
	this := AllMiniTickerResponseInner{}
	return &this
}

// NewAllMiniTickerResponseInnerWithDefaults instantiates a new AllMiniTickerResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMiniTickerResponseInnerWithDefaults() *AllMiniTickerResponseInner {
	this := AllMiniTickerResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllMiniTickerResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllMiniTickerResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *AllMiniTickerResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmallc() string {
	if o == nil || common.IsNil(o.Smallc) {
		var ret string
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmallcOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *AllMiniTickerResponseInner) SetSmallc(v string) {
	o.Smallc = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmallo() string {
	if o == nil || common.IsNil(o.Smallo) {
		var ret string
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmalloOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given string and assigns it to the O field.
func (o *AllMiniTickerResponseInner) SetSmallo(v string) {
	o.Smallo = &v
}

// GetH returns the H field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmallh() string {
	if o == nil || common.IsNil(o.Smallh) {
		var ret string
		return ret
	}
	return *o.Smallh
}

// GetHOk returns a tuple with the H field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmallhOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallh) {
		return nil, false
	}
	return o.Smallh, true
}

// HasH returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmallh() bool {
	if o != nil && !common.IsNil(o.Smallh) {
		return true
	}

	return false
}

// SetH gets a reference to the given string and assigns it to the H field.
func (o *AllMiniTickerResponseInner) SetSmallh(v string) {
	o.Smallh = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *AllMiniTickerResponseInner) SetSmalll(v string) {
	o.Smalll = &v
}

// GetV returns the V field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmallv() string {
	if o == nil || common.IsNil(o.Smallv) {
		var ret string
		return ret
	}
	return *o.Smallv
}

// GetVOk returns a tuple with the V field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmallvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallv) {
		return nil, false
	}
	return o.Smallv, true
}

// HasV returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmallv() bool {
	if o != nil && !common.IsNil(o.Smallv) {
		return true
	}

	return false
}

// SetV gets a reference to the given string and assigns it to the V field.
func (o *AllMiniTickerResponseInner) SetSmallv(v string) {
	o.Smallv = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *AllMiniTickerResponseInner) GetSmallq() string {
	if o == nil || common.IsNil(o.Smallq) {
		var ret string
		return ret
	}
	return *o.Smallq
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMiniTickerResponseInner) GetSmallqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallq) {
		return nil, false
	}
	return o.Smallq, true
}

// HasQ returns a boolean if a field has been set.
func (o *AllMiniTickerResponseInner) HasSmallq() bool {
	if o != nil && !common.IsNil(o.Smallq) {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *AllMiniTickerResponseInner) SetSmallq(v string) {
	o.Smallq = &v
}

func (o AllMiniTickerResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMiniTickerResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}
	if !common.IsNil(o.Smallh) {
		toSerialize["h"] = o.Smallh
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}
	if !common.IsNil(o.Smallv) {
		toSerialize["v"] = o.Smallv
	}
	if !common.IsNil(o.Smallq) {
		toSerialize["q"] = o.Smallq
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllMiniTickerResponseInner) UnmarshalJSON(data []byte) (err error) {
	varAllMiniTickerResponseInner := _AllMiniTickerResponseInner{}

	err = json.Unmarshal(data, &varAllMiniTickerResponseInner)

	if err != nil {
		return err
	}

	*o = AllMiniTickerResponseInner(varAllMiniTickerResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "c")
		delete(additionalProperties, "o")
		delete(additionalProperties, "h")
		delete(additionalProperties, "l")
		delete(additionalProperties, "v")
		delete(additionalProperties, "q")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllMiniTickerResponseInner struct {
	value *AllMiniTickerResponseInner
	isSet bool
}

func (v NullableAllMiniTickerResponseInner) Get() *AllMiniTickerResponseInner {
	return v.value
}

func (v *NullableAllMiniTickerResponseInner) Set(val *AllMiniTickerResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMiniTickerResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMiniTickerResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMiniTickerResponseInner(val *AllMiniTickerResponseInner) *NullableAllMiniTickerResponseInner {
	return &NullableAllMiniTickerResponseInner{value: val, isSet: true}
}

func (v NullableAllMiniTickerResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMiniTickerResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
