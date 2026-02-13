/*
Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin Pro WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the Risklevelchange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Risklevelchange{}

// Risklevelchange struct for Risklevelchange
type Risklevelchange struct {
	E                    *int64  `json:"E,omitempty"`
	Smallu               *string `json:"u,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smalleq              *string `json:"eq,omitempty"`
	Smallae              *string `json:"ae,omitempty"`
	Smallm               *string `json:"m,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Risklevelchange Risklevelchange

// NewRisklevelchange instantiates a new Risklevelchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRisklevelchange() *Risklevelchange {
	this := Risklevelchange{}
	return &this
}

// NewRisklevelchangeWithDefaults instantiates a new Risklevelchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRisklevelchangeWithDefaults() *Risklevelchange {
	this := Risklevelchange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Risklevelchange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Risklevelchange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Risklevelchange) SetE(v int64) {
	o.E = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *Risklevelchange) GetSmallu() string {
	if o == nil || common.IsNil(o.Smallu) {
		var ret string
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetSmalluOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *Risklevelchange) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *Risklevelchange) SetSmallu(v string) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *Risklevelchange) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *Risklevelchange) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *Risklevelchange) SetSmalls(v string) {
	o.Smalls = &v
}

// GetEq returns the Eq field value if set, zero value otherwise.
func (o *Risklevelchange) GetSmalleq() string {
	if o == nil || common.IsNil(o.Smalleq) {
		var ret string
		return ret
	}
	return *o.Smalleq
}

// GetEqOk returns a tuple with the Eq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetSmalleqOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalleq) {
		return nil, false
	}
	return o.Smalleq, true
}

// HasEq returns a boolean if a field has been set.
func (o *Risklevelchange) HasSmalleq() bool {
	if o != nil && !common.IsNil(o.Smalleq) {
		return true
	}

	return false
}

// SetEq gets a reference to the given string and assigns it to the Eq field.
func (o *Risklevelchange) SetSmalleq(v string) {
	o.Smalleq = &v
}

// GetAe returns the Ae field value if set, zero value otherwise.
func (o *Risklevelchange) GetSmallae() string {
	if o == nil || common.IsNil(o.Smallae) {
		var ret string
		return ret
	}
	return *o.Smallae
}

// GetAeOk returns a tuple with the Ae field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetSmallaeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallae) {
		return nil, false
	}
	return o.Smallae, true
}

// HasAe returns a boolean if a field has been set.
func (o *Risklevelchange) HasSmallae() bool {
	if o != nil && !common.IsNil(o.Smallae) {
		return true
	}

	return false
}

// SetAe gets a reference to the given string and assigns it to the Ae field.
func (o *Risklevelchange) SetSmallae(v string) {
	o.Smallae = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *Risklevelchange) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Risklevelchange) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *Risklevelchange) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *Risklevelchange) SetSmallm(v string) {
	o.Smallm = &v
}

func (o Risklevelchange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Risklevelchange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smalleq) {
		toSerialize["eq"] = o.Smalleq
	}
	if !common.IsNil(o.Smallae) {
		toSerialize["ae"] = o.Smallae
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Risklevelchange) UnmarshalJSON(data []byte) (err error) {
	varRisklevelchange := _Risklevelchange{}

	err = json.Unmarshal(data, &varRisklevelchange)

	if err != nil {
		return err
	}

	*o = Risklevelchange(varRisklevelchange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "u")
		delete(additionalProperties, "s")
		delete(additionalProperties, "eq")
		delete(additionalProperties, "ae")
		delete(additionalProperties, "m")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRisklevelchange struct {
	value *Risklevelchange
	isSet bool
}

func (v NullableRisklevelchange) Get() *Risklevelchange {
	return v.value
}

func (v *NullableRisklevelchange) Set(val *Risklevelchange) {
	v.value = val
	v.isSet = true
}

func (v NullableRisklevelchange) IsSet() bool {
	return v.isSet
}

func (v *NullableRisklevelchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRisklevelchange(val *Risklevelchange) *NullableRisklevelchange {
	return &NullableRisklevelchange{value: val, isSet: true}
}

func (v NullableRisklevelchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRisklevelchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
