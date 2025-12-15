/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PartialBookDepthStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartialBookDepthStreamsResponse{}

// PartialBookDepthStreamsResponse struct for PartialBookDepthStreamsResponse
type PartialBookDepthStreamsResponse struct {
	Smalle               *string                                `json:"e,omitempty"`
	E                    *int64                                 `json:"E,omitempty"`
	T                    *int64                                 `json:"T,omitempty"`
	Smalls               *string                                `json:"s,omitempty"`
	U                    *int64                                 `json:"U,omitempty"`
	Smallu               *int64                                 `json:"u,omitempty"`
	Smallpu              *int64                                 `json:"pu,omitempty"`
	Smallb               []PartialBookDepthStreamsResponseBItem `json:"b,omitempty"`
	Smalla               []PartialBookDepthStreamsResponseAItem `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartialBookDepthStreamsResponse PartialBookDepthStreamsResponse

// NewPartialBookDepthStreamsResponse instantiates a new PartialBookDepthStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialBookDepthStreamsResponse() *PartialBookDepthStreamsResponse {
	this := PartialBookDepthStreamsResponse{}
	return &this
}

// NewPartialBookDepthStreamsResponseWithDefaults instantiates a new PartialBookDepthStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialBookDepthStreamsResponseWithDefaults() *PartialBookDepthStreamsResponse {
	this := PartialBookDepthStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *PartialBookDepthStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *PartialBookDepthStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *PartialBookDepthStreamsResponse) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *PartialBookDepthStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *PartialBookDepthStreamsResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *PartialBookDepthStreamsResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetPu returns the Pu field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmallpu() int64 {
	if o == nil || common.IsNil(o.Smallpu) {
		var ret int64
		return ret
	}
	return *o.Smallpu
}

// GetPuOk returns a tuple with the Pu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmallpuOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallpu) {
		return nil, false
	}
	return o.Smallpu, true
}

// HasPu returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmallpu() bool {
	if o != nil && !common.IsNil(o.Smallpu) {
		return true
	}

	return false
}

// SetPu gets a reference to the given int64 and assigns it to the Pu field.
func (o *PartialBookDepthStreamsResponse) SetSmallpu(v int64) {
	o.Smallpu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmallb() []PartialBookDepthStreamsResponseBItem {
	if o == nil || common.IsNil(o.Smallb) {
		var ret []PartialBookDepthStreamsResponseBItem
		return ret
	}
	return o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmallbOk() ([]PartialBookDepthStreamsResponseBItem, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given []PartialBookDepthStreamsResponseBItem and assigns it to the B field.
func (o *PartialBookDepthStreamsResponse) SetSmallb(v []PartialBookDepthStreamsResponseBItem) {
	o.Smallb = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *PartialBookDepthStreamsResponse) GetSmalla() []PartialBookDepthStreamsResponseAItem {
	if o == nil || common.IsNil(o.Smalla) {
		var ret []PartialBookDepthStreamsResponseAItem
		return ret
	}
	return o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialBookDepthStreamsResponse) GetSmallaOk() ([]PartialBookDepthStreamsResponseAItem, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *PartialBookDepthStreamsResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given []PartialBookDepthStreamsResponseAItem and assigns it to the A field.
func (o *PartialBookDepthStreamsResponse) SetSmalla(v []PartialBookDepthStreamsResponseAItem) {
	o.Smalla = v
}

func (o PartialBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialBookDepthStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smallpu) {
		toSerialize["pu"] = o.Smallpu
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PartialBookDepthStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varPartialBookDepthStreamsResponse := _PartialBookDepthStreamsResponse{}

	err = json.Unmarshal(data, &varPartialBookDepthStreamsResponse)

	if err != nil {
		return err
	}

	*o = PartialBookDepthStreamsResponse(varPartialBookDepthStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "U")
		delete(additionalProperties, "u")
		delete(additionalProperties, "pu")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePartialBookDepthStreamsResponse struct {
	value *PartialBookDepthStreamsResponse
	isSet bool
}

func (v NullablePartialBookDepthStreamsResponse) Get() *PartialBookDepthStreamsResponse {
	return v.value
}

func (v *NullablePartialBookDepthStreamsResponse) Set(val *PartialBookDepthStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialBookDepthStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialBookDepthStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialBookDepthStreamsResponse(val *PartialBookDepthStreamsResponse) *NullablePartialBookDepthStreamsResponse {
	return &NullablePartialBookDepthStreamsResponse{value: val, isSet: true}
}

func (v NullablePartialBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialBookDepthStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
