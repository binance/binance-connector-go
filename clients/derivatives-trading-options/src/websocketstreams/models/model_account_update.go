/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AccountUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AccountUpdate{}

// AccountUpdate struct for AccountUpdate
type AccountUpdate struct {
	E                    *int64                `json:"E,omitempty"`
	B                    []AccountUpdateBInner `json:"B,omitempty"`
	G                    []AccountUpdateGInner `json:"G,omitempty"`
	P                    []AccountUpdatePInner `json:"P,omitempty"`
	Uid                  *int64                `json:"uid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUpdate AccountUpdate

// NewAccountUpdate instantiates a new AccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUpdate() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// NewAccountUpdateWithDefaults instantiates a new AccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUpdateWithDefaults() *AccountUpdate {
	this := AccountUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AccountUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AccountUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AccountUpdate) SetE(v int64) {
	o.E = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *AccountUpdate) GetB() []AccountUpdateBInner {
	if o == nil || common.IsNil(o.B) {
		var ret []AccountUpdateBInner
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetBOk() ([]AccountUpdateBInner, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *AccountUpdate) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given []AccountUpdateBInner and assigns it to the B field.
func (o *AccountUpdate) SetB(v []AccountUpdateBInner) {
	o.B = v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *AccountUpdate) GetG() []AccountUpdateGInner {
	if o == nil || common.IsNil(o.G) {
		var ret []AccountUpdateGInner
		return ret
	}
	return o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetGOk() ([]AccountUpdateGInner, bool) {
	if o == nil || common.IsNil(o.G) {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *AccountUpdate) HasG() bool {
	if o != nil && !common.IsNil(o.G) {
		return true
	}

	return false
}

// SetG gets a reference to the given []AccountUpdateGInner and assigns it to the G field.
func (o *AccountUpdate) SetG(v []AccountUpdateGInner) {
	o.G = v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *AccountUpdate) GetP() []AccountUpdatePInner {
	if o == nil || common.IsNil(o.P) {
		var ret []AccountUpdatePInner
		return ret
	}
	return o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetPOk() ([]AccountUpdatePInner, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *AccountUpdate) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given []AccountUpdatePInner and assigns it to the P field.
func (o *AccountUpdate) SetP(v []AccountUpdatePInner) {
	o.P = v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *AccountUpdate) GetUid() int64 {
	if o == nil || common.IsNil(o.Uid) {
		var ret int64
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUpdate) GetUidOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *AccountUpdate) HasUid() bool {
	if o != nil && !common.IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int64 and assigns it to the Uid field.
func (o *AccountUpdate) SetUid(v int64) {
	o.Uid = &v
}

func (o AccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.G) {
		toSerialize["G"] = o.G
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}
	if !common.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccountUpdate) UnmarshalJSON(data []byte) (err error) {
	varAccountUpdate := _AccountUpdate{}

	err = json.Unmarshal(data, &varAccountUpdate)

	if err != nil {
		return err
	}

	*o = AccountUpdate(varAccountUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "B")
		delete(additionalProperties, "G")
		delete(additionalProperties, "P")
		delete(additionalProperties, "uid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUpdate struct {
	value *AccountUpdate
	isSet bool
}

func (v NullableAccountUpdate) Get() *AccountUpdate {
	return v.value
}

func (v *NullableAccountUpdate) Set(val *AccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUpdate(val *AccountUpdate) *NullableAccountUpdate {
	return &NullableAccountUpdate{value: val, isSet: true}
}

func (v NullableAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
