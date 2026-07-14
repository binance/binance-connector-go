/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineStreamResponse{}

// KlineStreamResponse struct for KlineStreamResponse
type KlineStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// symbol
	Smalls               *string               `json:"s,omitempty"`
	Smallk               *KlineStreamResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineStreamResponse KlineStreamResponse

// NewKlineStreamResponse instantiates a new KlineStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineStreamResponse() *KlineStreamResponse {
	this := KlineStreamResponse{}
	return &this
}

// NewKlineStreamResponseWithDefaults instantiates a new KlineStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineStreamResponseWithDefaults() *KlineStreamResponse {
	this := KlineStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *KlineStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *KlineStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *KlineStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *KlineStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *KlineStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *KlineStreamResponse) GetSmallk() KlineStreamResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret KlineStreamResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineStreamResponse) GetSmallkOk() (*KlineStreamResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *KlineStreamResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given KlineStreamResponseK and assigns it to the K field.
func (o *KlineStreamResponse) SetSmallk(v KlineStreamResponseK) {
	o.Smallk = &v
}

func (o KlineStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineStreamResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallk) {
		toSerialize["k"] = o.Smallk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KlineStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varKlineStreamResponse := _KlineStreamResponse{}

	err = json.Unmarshal(data, &varKlineStreamResponse)

	if err != nil {
		return err
	}

	*o = KlineStreamResponse(varKlineStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "k")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKlineStreamResponse struct {
	value *KlineStreamResponse
	isSet bool
}

func (v NullableKlineStreamResponse) Get() *KlineStreamResponse {
	return v.value
}

func (v *NullableKlineStreamResponse) Set(val *KlineStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineStreamResponse(val *KlineStreamResponse) *NullableKlineStreamResponse {
	return &NullableKlineStreamResponse{value: val, isSet: true}
}

func (v NullableKlineStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
