/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the VaspListResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &VaspListResponseInner{}

// VaspListResponseInner struct for VaspListResponseInner
type VaspListResponseInner struct {
	VaspCode             *string `json:"vaspCode,omitempty"`
	VaspName             *string `json:"vaspName,omitempty"`
	Identifier           *string `json:"identifier,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VaspListResponseInner VaspListResponseInner

// NewVaspListResponseInner instantiates a new VaspListResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaspListResponseInner() *VaspListResponseInner {
	this := VaspListResponseInner{}
	return &this
}

// NewVaspListResponseInnerWithDefaults instantiates a new VaspListResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaspListResponseInnerWithDefaults() *VaspListResponseInner {
	this := VaspListResponseInner{}
	return &this
}

// GetVaspCode returns the VaspCode field value if set, zero value otherwise.
func (o *VaspListResponseInner) GetVaspCode() string {
	if o == nil || common.IsNil(o.VaspCode) {
		var ret string
		return ret
	}
	return *o.VaspCode
}

// GetVaspCodeOk returns a tuple with the VaspCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaspListResponseInner) GetVaspCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.VaspCode) {
		return nil, false
	}
	return o.VaspCode, true
}

// HasVaspCode returns a boolean if a field has been set.
func (o *VaspListResponseInner) HasVaspCode() bool {
	if o != nil && !common.IsNil(o.VaspCode) {
		return true
	}

	return false
}

// SetVaspCode gets a reference to the given string and assigns it to the VaspCode field.
func (o *VaspListResponseInner) SetVaspCode(v string) {
	o.VaspCode = &v
}

// GetVaspName returns the VaspName field value if set, zero value otherwise.
func (o *VaspListResponseInner) GetVaspName() string {
	if o == nil || common.IsNil(o.VaspName) {
		var ret string
		return ret
	}
	return *o.VaspName
}

// GetVaspNameOk returns a tuple with the VaspName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaspListResponseInner) GetVaspNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.VaspName) {
		return nil, false
	}
	return o.VaspName, true
}

// HasVaspName returns a boolean if a field has been set.
func (o *VaspListResponseInner) HasVaspName() bool {
	if o != nil && !common.IsNil(o.VaspName) {
		return true
	}

	return false
}

// SetVaspName gets a reference to the given string and assigns it to the VaspName field.
func (o *VaspListResponseInner) SetVaspName(v string) {
	o.VaspName = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *VaspListResponseInner) GetIdentifier() string {
	if o == nil || common.IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaspListResponseInner) GetIdentifierOk() (*string, bool) {
	if o == nil || common.IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *VaspListResponseInner) HasIdentifier() bool {
	if o != nil && !common.IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *VaspListResponseInner) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o VaspListResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaspListResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.VaspCode) {
		toSerialize["vaspCode"] = o.VaspCode
	}
	if !common.IsNil(o.VaspName) {
		toSerialize["vaspName"] = o.VaspName
	}
	if !common.IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VaspListResponseInner) UnmarshalJSON(data []byte) (err error) {
	varVaspListResponseInner := _VaspListResponseInner{}

	err = json.Unmarshal(data, &varVaspListResponseInner)

	if err != nil {
		return err
	}

	*o = VaspListResponseInner(varVaspListResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "vaspCode")
		delete(additionalProperties, "vaspName")
		delete(additionalProperties, "identifier")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVaspListResponseInner struct {
	value *VaspListResponseInner
	isSet bool
}

func (v NullableVaspListResponseInner) Get() *VaspListResponseInner {
	return v.value
}

func (v *NullableVaspListResponseInner) Set(val *VaspListResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVaspListResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVaspListResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaspListResponseInner(val *VaspListResponseInner) *NullableVaspListResponseInner {
	return &NullableVaspListResponseInner{value: val, isSet: true}
}

func (v NullableVaspListResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaspListResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
