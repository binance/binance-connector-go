/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the BalancePositionUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalancePositionUpdate{}

// BalancePositionUpdate struct for BalancePositionUpdate
type BalancePositionUpdate struct {
	E                    *int64                        `json:"E,omitempty"`
	T                    *int64                        `json:"T,omitempty"`
	Smallm               *string                       `json:"m,omitempty"`
	B                    []BalancePositionUpdateBInner `json:"B,omitempty"`
	P                    []BalancePositionUpdatePInner `json:"P,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BalancePositionUpdate BalancePositionUpdate

// NewBalancePositionUpdate instantiates a new BalancePositionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalancePositionUpdate() *BalancePositionUpdate {
	this := BalancePositionUpdate{}
	return &this
}

// NewBalancePositionUpdateWithDefaults instantiates a new BalancePositionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalancePositionUpdateWithDefaults() *BalancePositionUpdate {
	this := BalancePositionUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BalancePositionUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *BalancePositionUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *BalancePositionUpdate) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *BalancePositionUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *BalancePositionUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *BalancePositionUpdate) SetT(v int64) {
	o.T = &v
}

// GetM returns the M field value if set, zero value otherwise.
func (o *BalancePositionUpdate) GetSmallm() string {
	if o == nil || common.IsNil(o.Smallm) {
		var ret string
		return ret
	}
	return *o.Smallm
}

// GetMOk returns a tuple with the M field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdate) GetSmallmOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallm) {
		return nil, false
	}
	return o.Smallm, true
}

// HasM returns a boolean if a field has been set.
func (o *BalancePositionUpdate) HasSmallm() bool {
	if o != nil && !common.IsNil(o.Smallm) {
		return true
	}

	return false
}

// SetM gets a reference to the given string and assigns it to the M field.
func (o *BalancePositionUpdate) SetSmallm(v string) {
	o.Smallm = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *BalancePositionUpdate) GetB() []BalancePositionUpdateBInner {
	if o == nil || common.IsNil(o.B) {
		var ret []BalancePositionUpdateBInner
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdate) GetBOk() ([]BalancePositionUpdateBInner, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *BalancePositionUpdate) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given []BalancePositionUpdateBInner and assigns it to the B field.
func (o *BalancePositionUpdate) SetB(v []BalancePositionUpdateBInner) {
	o.B = v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *BalancePositionUpdate) GetP() []BalancePositionUpdatePInner {
	if o == nil || common.IsNil(o.P) {
		var ret []BalancePositionUpdatePInner
		return ret
	}
	return o.P
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalancePositionUpdate) GetPOk() ([]BalancePositionUpdatePInner, bool) {
	if o == nil || common.IsNil(o.P) {
		return nil, false
	}
	return o.P, true
}

// HasP returns a boolean if a field has been set.
func (o *BalancePositionUpdate) HasP() bool {
	if o != nil && !common.IsNil(o.P) {
		return true
	}

	return false
}

// SetP gets a reference to the given []BalancePositionUpdatePInner and assigns it to the P field.
func (o *BalancePositionUpdate) SetP(v []BalancePositionUpdatePInner) {
	o.P = v
}

func (o BalancePositionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalancePositionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smallm) {
		toSerialize["m"] = o.Smallm
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.P) {
		toSerialize["P"] = o.P
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BalancePositionUpdate) UnmarshalJSON(data []byte) (err error) {
	varBalancePositionUpdate := _BalancePositionUpdate{}

	err = json.Unmarshal(data, &varBalancePositionUpdate)

	if err != nil {
		return err
	}

	*o = BalancePositionUpdate(varBalancePositionUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "m")
		delete(additionalProperties, "B")
		delete(additionalProperties, "P")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalancePositionUpdate struct {
	value *BalancePositionUpdate
	isSet bool
}

func (v NullableBalancePositionUpdate) Get() *BalancePositionUpdate {
	return v.value
}

func (v *NullableBalancePositionUpdate) Set(val *BalancePositionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBalancePositionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBalancePositionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalancePositionUpdate(val *BalancePositionUpdate) *NullableBalancePositionUpdate {
	return &NullableBalancePositionUpdate{value: val, isSet: true}
}

func (v NullableBalancePositionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalancePositionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
