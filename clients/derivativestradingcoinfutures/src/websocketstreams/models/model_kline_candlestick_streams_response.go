/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineCandlestickStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineCandlestickStreamsResponse{}

// KlineCandlestickStreamsResponse struct for KlineCandlestickStreamsResponse
type KlineCandlestickStreamsResponse struct {
	Smalle               *string                           `json:"e,omitempty"`
	E                    *int64                            `json:"E,omitempty"`
	Smalls               *string                           `json:"s,omitempty"`
	Smallk               *KlineCandlestickStreamsResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineCandlestickStreamsResponse KlineCandlestickStreamsResponse

// NewKlineCandlestickStreamsResponse instantiates a new KlineCandlestickStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineCandlestickStreamsResponse() *KlineCandlestickStreamsResponse {
	this := KlineCandlestickStreamsResponse{}
	return &this
}

// NewKlineCandlestickStreamsResponseWithDefaults instantiates a new KlineCandlestickStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineCandlestickStreamsResponseWithDefaults() *KlineCandlestickStreamsResponse {
	this := KlineCandlestickStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineCandlestickStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineCandlestickStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *KlineCandlestickStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineCandlestickStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineCandlestickStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *KlineCandlestickStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *KlineCandlestickStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *KlineCandlestickStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *KlineCandlestickStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *KlineCandlestickStreamsResponse) GetSmallk() KlineCandlestickStreamsResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret KlineCandlestickStreamsResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineCandlestickStreamsResponse) GetSmallkOk() (*KlineCandlestickStreamsResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *KlineCandlestickStreamsResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given KlineCandlestickStreamsResponseK and assigns it to the K field.
func (o *KlineCandlestickStreamsResponse) SetSmallk(v KlineCandlestickStreamsResponseK) {
	o.Smallk = &v
}

func (o KlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineCandlestickStreamsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *KlineCandlestickStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varKlineCandlestickStreamsResponse := _KlineCandlestickStreamsResponse{}

	err = json.Unmarshal(data, &varKlineCandlestickStreamsResponse)

	if err != nil {
		return err
	}

	*o = KlineCandlestickStreamsResponse(varKlineCandlestickStreamsResponse)

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

type NullableKlineCandlestickStreamsResponse struct {
	value *KlineCandlestickStreamsResponse
	isSet bool
}

func (v NullableKlineCandlestickStreamsResponse) Get() *KlineCandlestickStreamsResponse {
	return v.value
}

func (v *NullableKlineCandlestickStreamsResponse) Set(val *KlineCandlestickStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineCandlestickStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineCandlestickStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineCandlestickStreamsResponse(val *KlineCandlestickStreamsResponse) *NullableKlineCandlestickStreamsResponse {
	return &NullableKlineCandlestickStreamsResponse{value: val, isSet: true}
}

func (v NullableKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineCandlestickStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
