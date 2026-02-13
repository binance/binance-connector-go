/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the BalanceUpdate type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &BalanceUpdate{}

// BalanceUpdate struct for BalanceUpdate
type BalanceUpdate struct {
	E                    *int64  `json:"E,omitempty"`
	A                    *string `json:"a,omitempty"`
	D                    *string `json:"d,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BalanceUpdate BalanceUpdate

// NewBalanceUpdate instantiates a new BalanceUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalanceUpdate() *BalanceUpdate {
	this := BalanceUpdate{}
	return &this
}

// NewBalanceUpdateWithDefaults instantiates a new BalanceUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceUpdateWithDefaults() *BalanceUpdate {
	this := BalanceUpdate{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *BalanceUpdate) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceUpdate) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *BalanceUpdate) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *BalanceUpdate) SetE(v int64) {
	o.E = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *BalanceUpdate) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceUpdate) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *BalanceUpdate) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *BalanceUpdate) SetA(v string) {
	o.A = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *BalanceUpdate) GetD() string {
	if o == nil || common.IsNil(o.D) {
		var ret string
		return ret
	}
	return *o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceUpdate) GetDOk() (*string, bool) {
	if o == nil || common.IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *BalanceUpdate) HasD() bool {
	if o != nil && !common.IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given string and assigns it to the D field.
func (o *BalanceUpdate) SetD(v string) {
	o.D = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *BalanceUpdate) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BalanceUpdate) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *BalanceUpdate) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *BalanceUpdate) SetT(v int64) {
	o.T = &v
}

func (o BalanceUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BalanceUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !common.IsNil(o.D) {
		toSerialize["d"] = o.D
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BalanceUpdate) UnmarshalJSON(data []byte) (err error) {
	varBalanceUpdate := _BalanceUpdate{}

	err = json.Unmarshal(data, &varBalanceUpdate)

	if err != nil {
		return err
	}

	*o = BalanceUpdate(varBalanceUpdate)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "a")
		delete(additionalProperties, "d")
		delete(additionalProperties, "T")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalanceUpdate struct {
	value *BalanceUpdate
	isSet bool
}

func (v NullableBalanceUpdate) Get() *BalanceUpdate {
	return v.value
}

func (v *NullableBalanceUpdate) Set(val *BalanceUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableBalanceUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableBalanceUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalanceUpdate(val *BalanceUpdate) *NullableBalanceUpdate {
	return &NullableBalanceUpdate{value: val, isSet: true}
}

func (v NullableBalanceUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalanceUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
