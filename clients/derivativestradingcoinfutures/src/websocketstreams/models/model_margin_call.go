/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the MarginCall type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarginCall{}

// MarginCall struct for MarginCall
type MarginCall struct {
	E                    *int64             `json:"E,omitempty"`
	Smalli               *string            `json:"i,omitempty"`
	Smallcw              *string            `json:"cw,omitempty"`
	Smallp               []MarginCallPInner `json:"p,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarginCall MarginCall

// NewMarginCall instantiates a new MarginCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginCall() *MarginCall {
	this := MarginCall{}
	return &this
}

// NewMarginCallWithDefaults instantiates a new MarginCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginCallWithDefaults() *MarginCall {
	this := MarginCall{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarginCall) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCall) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarginCall) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarginCall) SetE(v int64) {
	o.E = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *MarginCall) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCall) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *MarginCall) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *MarginCall) SetSmalli(v string) {
	o.Smalli = &v
}

// GetCw returns the Cw field value if set, zero value otherwise.
func (o *MarginCall) GetSmallcw() string {
	if o == nil || common.IsNil(o.Smallcw) {
		var ret string
		return ret
	}
	return *o.Smallcw
}

// GetCwOk returns a tuple with the Cw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCall) GetSmallcwOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallcw) {
		return nil, false
	}
	return o.Smallcw, true
}

// HasCw returns a boolean if a field has been set.
func (o *MarginCall) HasSmallcw() bool {
	if o != nil && !common.IsNil(o.Smallcw) {
		return true
	}

	return false
}

// SetCw gets a reference to the given string and assigns it to the Cw field.
func (o *MarginCall) SetSmallcw(v string) {
	o.Smallcw = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *MarginCall) GetSmallp() []MarginCallPInner {
	if o == nil || common.IsNil(o.Smallp) {
		var ret []MarginCallPInner
		return ret
	}
	return o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginCall) GetSmallpOk() ([]MarginCallPInner, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *MarginCall) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given []MarginCallPInner and assigns it to the P field.
func (o *MarginCall) SetSmallp(v []MarginCallPInner) {
	o.Smallp = v
}

func (o MarginCall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallcw) {
		toSerialize["cw"] = o.Smallcw
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MarginCall) UnmarshalJSON(data []byte) (err error) {
	varMarginCall := _MarginCall{}

	err = json.Unmarshal(data, &varMarginCall)

	if err != nil {
		return err
	}

	*o = MarginCall(varMarginCall)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "E")
		delete(additionalProperties, "i")
		delete(additionalProperties, "cw")
		delete(additionalProperties, "p")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMarginCall struct {
	value *MarginCall
	isSet bool
}

func (v NullableMarginCall) Get() *MarginCall {
	return v.value
}

func (v *NullableMarginCall) Set(val *MarginCall) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginCall) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginCall(val *MarginCall) *NullableMarginCall {
	return &NullableMarginCall{value: val, isSet: true}
}

func (v NullableMarginCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
