/*
Binance Margin Trading WebSocket Market Streams

OpenAPI Specification for the Binance Margin Trading WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the UserLiabilityChange type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &UserLiabilityChange{}

// UserLiabilityChange struct for UserLiabilityChange
type UserLiabilityChange struct {
	E                    *int64  `json:"E,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	Smallt               *string `json:"t,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserLiabilityChange UserLiabilityChange

// NewUserLiabilityChange instantiates a new UserLiabilityChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLiabilityChange() *UserLiabilityChange {
	this := UserLiabilityChange{}
	return &this
}

// NewUserLiabilityChangeWithDefaults instantiates a new UserLiabilityChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLiabilityChangeWithDefaults() *UserLiabilityChange {
	this := UserLiabilityChange{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *UserLiabilityChange) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLiabilityChange) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *UserLiabilityChange) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *UserLiabilityChange) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *UserLiabilityChange) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLiabilityChange) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *UserLiabilityChange) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *UserLiabilityChange) SetSmalla(v string) {
	o.Smalla = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *UserLiabilityChange) GetSmallt() string {
	if o == nil || common.IsNil(o.Smallt) {
		var ret string
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLiabilityChange) GetSmalltOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *UserLiabilityChange) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *UserLiabilityChange) SetSmallt(v string) {
	o.Smallt = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *UserLiabilityChange) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLiabilityChange) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *UserLiabilityChange) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *UserLiabilityChange) SetSmallp(v string) {
	o.Smallp = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *UserLiabilityChange) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLiabilityChange) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *UserLiabilityChange) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *UserLiabilityChange) SetSmalli(v string) {
	o.Smalli = &v
}

func (o UserLiabilityChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserLiabilityChange) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserLiabilityChange) UnmarshalJSON(data []byte) (err error) {
	varUserLiabilityChange := _UserLiabilityChange{}

	err = json.Unmarshal(data, &varUserLiabilityChange)

	if err != nil {
		return err
	}

	*o = UserLiabilityChange(varUserLiabilityChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "t")
		delete(additionalProperties, "p")
		delete(additionalProperties, "i")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserLiabilityChange struct {
	value *UserLiabilityChange
	isSet bool
}

func (v NullableUserLiabilityChange) Get() *UserLiabilityChange {
	return v.value
}

func (v *NullableUserLiabilityChange) Set(val *UserLiabilityChange) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLiabilityChange) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLiabilityChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLiabilityChange(val *UserLiabilityChange) *NullableUserLiabilityChange {
	return &NullableUserLiabilityChange{value: val, isSet: true}
}

func (v NullableUserLiabilityChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLiabilityChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
