/*
Binance Derivatives Trading Portfolio Margin WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the Liabilitychange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &Liabilitychange{}

// Liabilitychange struct for Liabilitychange
type Liabilitychange struct {
	E                    *int64  `json:"E,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smallt               *string `json:"t,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	Smalll               *string `json:"l,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Liabilitychange Liabilitychange

// NewLiabilitychange instantiates a new Liabilitychange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiabilitychange() *Liabilitychange {
	this := Liabilitychange{}
	return &this
}

// NewLiabilitychangeWithDefaults instantiates a new Liabilitychange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiabilitychangeWithDefaults() *Liabilitychange {
	this := Liabilitychange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *Liabilitychange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *Liabilitychange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *Liabilitychange) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *Liabilitychange) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *Liabilitychange) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *Liabilitychange) SetSmalla(v string) {
	o.Smalla = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Liabilitychange) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *Liabilitychange) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *Liabilitychange) SetSmallt(v string) {
	o.Smallt = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *Liabilitychange) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *Liabilitychange) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *Liabilitychange) SetT(v int64) {
	o.T = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *Liabilitychange) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *Liabilitychange) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *Liabilitychange) SetSmallp(v string) {
	o.Smallp = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *Liabilitychange) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *Liabilitychange) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *Liabilitychange) SetSmalli(v string) {
	o.Smalli = &v
}

// GetL returns the L field value if set, zero value otherwise.
func (o *Liabilitychange) GetSmalll() string {
	if o == nil || common.IsNil(o.Smalll) {
		var ret string
		return ret
	}
	return *o.Smalll
}

// GetLOk returns a tuple with the L field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Liabilitychange) GetSmalllOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalll) {
		return nil, false
	}
	return o.Smalll, true
}

// HasL returns a boolean if a field has been set.
func (o *Liabilitychange) HasSmalll() bool {
	if o != nil && !common.IsNil(o.Smalll) {
		return true
	}

	return false
}

// SetL gets a reference to the given string and assigns it to the L field.
func (o *Liabilitychange) SetSmalll(v string) {
	o.Smalll = &v
}

func (o Liabilitychange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Liabilitychange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smalll) {
		toSerialize["l"] = o.Smalll
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Liabilitychange) UnmarshalJSON(data []byte) (err error) {
	varLiabilitychange := _Liabilitychange{}

	err = json.Unmarshal(data, &varLiabilitychange)

	if err != nil {
		return err
	}

	*o = Liabilitychange(varLiabilitychange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "t")
		delete(additionalProperties, "T")
		delete(additionalProperties, "p")
		delete(additionalProperties, "i")
		delete(additionalProperties, "l")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiabilitychange struct {
	value *Liabilitychange
	isSet bool
}

func (v NullableLiabilitychange) Get() *Liabilitychange {
	return v.value
}

func (v *NullableLiabilitychange) Set(val *Liabilitychange) {
	v.value = val
	v.isSet = true
}

func (v NullableLiabilitychange) IsSet() bool {
	return v.isSet
}

func (v *NullableLiabilitychange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiabilitychange(val *Liabilitychange) *NullableLiabilitychange {
	return &NullableLiabilitychange{value: val, isSet: true}
}

func (v NullableLiabilitychange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiabilitychange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
