/*
Spot WebSocket Market Streams

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KlineResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlineResponse{}

// KlineResponse struct for KlineResponse
type KlineResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time
	E *int64 `json:"E,omitempty"`
	// Symbol
	Smalls               *string         `json:"s,omitempty"`
	Smallk               *KlineResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlineResponse KlineResponse

// NewKlineResponse instantiates a new KlineResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlineResponse() *KlineResponse {
	this := KlineResponse{}
	return &this
}

// NewKlineResponseWithDefaults instantiates a new KlineResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlineResponseWithDefaults() *KlineResponse {
	this := KlineResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *KlineResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *KlineResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *KlineResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *KlineResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *KlineResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *KlineResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *KlineResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *KlineResponse) GetSmallk() KlineResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret KlineResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlineResponse) GetSmallkOk() (*KlineResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *KlineResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given KlineResponseK and assigns it to the K field.
func (o *KlineResponse) SetSmallk(v KlineResponseK) {
	o.Smallk = &v
}

func (o KlineResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlineResponse) ToMap() (map[string]interface{}, error) {
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

func (o *KlineResponse) UnmarshalJSON(data []byte) (err error) {
	varKlineResponse := _KlineResponse{}

	err = json.Unmarshal(data, &varKlineResponse)

	if err != nil {
		return err
	}

	*o = KlineResponse(varKlineResponse)

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

type NullableKlineResponse struct {
	value *KlineResponse
	isSet bool
}

func (v NullableKlineResponse) Get() *KlineResponse {
	return v.value
}

func (v *NullableKlineResponse) Set(val *KlineResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlineResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlineResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlineResponse(val *KlineResponse) *NullableKlineResponse {
	return &NullableKlineResponse{value: val, isSet: true}
}

func (v NullableKlineResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlineResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
