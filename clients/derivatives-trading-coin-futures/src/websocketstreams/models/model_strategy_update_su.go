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

// checks if the StrategyUpdateSu type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &StrategyUpdateSu{}

// StrategyUpdateSu struct for StrategyUpdateSu
type StrategyUpdateSu struct {
	Smallsi              *int64  `json:"si,omitempty"`
	Smallst              *string `json:"st,omitempty"`
	Smallss              *string `json:"ss,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallut              *int64  `json:"ut,omitempty"`
	Smallc               *int64  `json:"c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StrategyUpdateSu StrategyUpdateSu

// NewStrategyUpdateSu instantiates a new StrategyUpdateSu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrategyUpdateSu() *StrategyUpdateSu {
	this := StrategyUpdateSu{}
	return &this
}

// NewStrategyUpdateSuWithDefaults instantiates a new StrategyUpdateSu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrategyUpdateSuWithDefaults() *StrategyUpdateSu {
	this := StrategyUpdateSu{}
	return &this
}

// GetSi returns the Si field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmallsi() int64 {
	if o == nil || common.IsNil(o.Smallsi) {
		var ret int64
		return ret
	}
	return *o.Smallsi
}

// GetSiOk returns a tuple with the Si field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallsiOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallsi) {
		return nil, false
	}
	return o.Smallsi, true
}

// HasSi returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmallsi() bool {
	if o != nil && !common.IsNil(o.Smallsi) {
		return true
	}

	return false
}

// SetSi gets a reference to the given int64 and assigns it to the Si field.
func (o *StrategyUpdateSu) SetSmallsi(v int64) {
	o.Smallsi = &v
}

// GetSt returns the St field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmallst() string {
	if o == nil || common.IsNil(o.Smallst) {
		var ret string
		return ret
	}
	return *o.Smallst
}

// GetStOk returns a tuple with the St field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallstOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallst) {
		return nil, false
	}
	return o.Smallst, true
}

// HasSt returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmallst() bool {
	if o != nil && !common.IsNil(o.Smallst) {
		return true
	}

	return false
}

// SetSt gets a reference to the given string and assigns it to the St field.
func (o *StrategyUpdateSu) SetSmallst(v string) {
	o.Smallst = &v
}

// GetSs returns the Ss field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmallss() string {
	if o == nil || common.IsNil(o.Smallss) {
		var ret string
		return ret
	}
	return *o.Smallss
}

// GetSsOk returns a tuple with the Ss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallssOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallss) {
		return nil, false
	}
	return o.Smallss, true
}

// HasSs returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmallss() bool {
	if o != nil && !common.IsNil(o.Smallss) {
		return true
	}

	return false
}

// SetSs gets a reference to the given string and assigns it to the Ss field.
func (o *StrategyUpdateSu) SetSmallss(v string) {
	o.Smallss = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *StrategyUpdateSu) SetSmalls(v string) {
	o.Smalls = &v
}

// GetUt returns the Ut field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmallut() int64 {
	if o == nil || common.IsNil(o.Smallut) {
		var ret int64
		return ret
	}
	return *o.Smallut
}

// GetUtOk returns a tuple with the Ut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallutOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallut) {
		return nil, false
	}
	return o.Smallut, true
}

// HasUt returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmallut() bool {
	if o != nil && !common.IsNil(o.Smallut) {
		return true
	}

	return false
}

// SetUt gets a reference to the given int64 and assigns it to the Ut field.
func (o *StrategyUpdateSu) SetSmallut(v int64) {
	o.Smallut = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *StrategyUpdateSu) GetSmallc() int64 {
	if o == nil || common.IsNil(o.Smallc) {
		var ret int64
		return ret
	}
	return *o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StrategyUpdateSu) GetSmallcOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *StrategyUpdateSu) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given int64 and assigns it to the C field.
func (o *StrategyUpdateSu) SetSmallc(v int64) {
	o.Smallc = &v
}

func (o StrategyUpdateSu) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StrategyUpdateSu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smallsi) {
		toSerialize["si"] = o.Smallsi
	}
	if !common.IsNil(o.Smallst) {
		toSerialize["st"] = o.Smallst
	}
	if !common.IsNil(o.Smallss) {
		toSerialize["ss"] = o.Smallss
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallut) {
		toSerialize["ut"] = o.Smallut
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StrategyUpdateSu) UnmarshalJSON(data []byte) (err error) {
	varStrategyUpdateSu := _StrategyUpdateSu{}

	err = json.Unmarshal(data, &varStrategyUpdateSu)

	if err != nil {
		return err
	}

	*o = StrategyUpdateSu(varStrategyUpdateSu)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "si")
		delete(additionalProperties, "st")
		delete(additionalProperties, "ss")
		delete(additionalProperties, "s")
		delete(additionalProperties, "ut")
		delete(additionalProperties, "c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStrategyUpdateSu struct {
	value *StrategyUpdateSu
	isSet bool
}

func (v NullableStrategyUpdateSu) Get() *StrategyUpdateSu {
	return v.value
}

func (v *NullableStrategyUpdateSu) Set(val *StrategyUpdateSu) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategyUpdateSu) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategyUpdateSu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategyUpdateSu(val *StrategyUpdateSu) *NullableStrategyUpdateSu {
	return &NullableStrategyUpdateSu{value: val, isSet: true}
}

func (v NullableStrategyUpdateSu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategyUpdateSu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
