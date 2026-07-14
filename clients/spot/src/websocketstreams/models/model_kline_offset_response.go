/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineOffsetResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineOffsetResponse{}

// KlineOffsetResponse struct for KlineOffsetResponse
type KlineOffsetResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls               *string         `json:"s,omitempty"`
	Smallk               *KlineResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineOffsetResponse KlineOffsetResponse

// NewKlineOffsetResponse instantiates a new KlineOffsetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineOffsetResponse() *KlineOffsetResponse {
	this := KlineOffsetResponse{}
	return &this
}

// NewKlineOffsetResponseWithDefaults instantiates a new KlineOffsetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineOffsetResponseWithDefaults() *KlineOffsetResponse {
	this := KlineOffsetResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineOffsetResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineOffsetResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineOffsetResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *KlineOffsetResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineOffsetResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineOffsetResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineOffsetResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *KlineOffsetResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *KlineOffsetResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineOffsetResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *KlineOffsetResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *KlineOffsetResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *KlineOffsetResponse) GetSmallk() KlineResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret KlineResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineOffsetResponse) GetSmallkOk() (*KlineResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *KlineOffsetResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given KlineResponseK and assigns it to the K field.
func (o *KlineOffsetResponse) SetSmallk(v KlineResponseK) {
	o.Smallk = &v
}

func (o KlineOffsetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineOffsetResponse) ToMap() (map[string]interface{}, error) {
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

func (o *KlineOffsetResponse) UnmarshalJSON(data []byte) (err error) {
	varKlineOffsetResponse := _KlineOffsetResponse{}

	err = json.Unmarshal(data, &varKlineOffsetResponse)

	if err != nil {
		return err
	}

	*o = KlineOffsetResponse(varKlineOffsetResponse)

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

type NullableKlineOffsetResponse struct {
	value *KlineOffsetResponse
	isSet bool
}

func (v NullableKlineOffsetResponse) Get() *KlineOffsetResponse {
	return v.value
}

func (v *NullableKlineOffsetResponse) Set(val *KlineOffsetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineOffsetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineOffsetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineOffsetResponse(val *KlineOffsetResponse) *NullableKlineOffsetResponse {
	return &NullableKlineOffsetResponse{value: val, isSet: true}
}

func (v NullableKlineOffsetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineOffsetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
