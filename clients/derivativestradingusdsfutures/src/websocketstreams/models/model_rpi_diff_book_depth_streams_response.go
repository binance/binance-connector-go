/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RpiDiffBookDepthStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RpiDiffBookDepthStreamsResponse{}

// RpiDiffBookDepthStreamsResponse struct for RpiDiffBookDepthStreamsResponse
type RpiDiffBookDepthStreamsResponse struct {
	Smalle               *string                                `json:"e,omitempty"`
	E                    *int64                                 `json:"E,omitempty"`
	T                    *int64                                 `json:"T,omitempty"`
	Smalls               *string                                `json:"s,omitempty"`
	U                    *int64                                 `json:"U,omitempty"`
	Smallu               *int64                                 `json:"u,omitempty"`
	Smallpu              *int64                                 `json:"pu,omitempty"`
	Smallb               []RpiDiffBookDepthStreamsResponseBItem `json:"b,omitempty"`
	Smalla               []RpiDiffBookDepthStreamsResponseAItem `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RpiDiffBookDepthStreamsResponse RpiDiffBookDepthStreamsResponse

// NewRpiDiffBookDepthStreamsResponse instantiates a new RpiDiffBookDepthStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpiDiffBookDepthStreamsResponse() *RpiDiffBookDepthStreamsResponse {
	this := RpiDiffBookDepthStreamsResponse{}
	return &this
}

// NewRpiDiffBookDepthStreamsResponseWithDefaults instantiates a new RpiDiffBookDepthStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpiDiffBookDepthStreamsResponseWithDefaults() *RpiDiffBookDepthStreamsResponse {
	this := RpiDiffBookDepthStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *RpiDiffBookDepthStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *RpiDiffBookDepthStreamsResponse) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *RpiDiffBookDepthStreamsResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetPu returns the Pu field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallpu() int64 {
	if o == nil || common.IsNil(o.Smallpu) {
		var ret int64
		return ret
	}
	return *o.Smallpu
}

// GetPuOk returns a tuple with the Pu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallpuOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallpu) {
		return nil, false
	}
	return o.Smallpu, true
}

// HasPu returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmallpu() bool {
	if o != nil && !common.IsNil(o.Smallpu) {
		return true
	}

	return false
}

// SetPu gets a reference to the given int64 and assigns it to the Pu field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmallpu(v int64) {
	o.Smallpu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallb() []RpiDiffBookDepthStreamsResponseBItem {
	if o == nil || common.IsNil(o.Smallb) {
		var ret []RpiDiffBookDepthStreamsResponseBItem
		return ret
	}
	return o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallbOk() ([]RpiDiffBookDepthStreamsResponseBItem, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given []RpiDiffBookDepthStreamsResponseBItem and assigns it to the B field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmallb(v []RpiDiffBookDepthStreamsResponseBItem) {
	o.Smallb = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *RpiDiffBookDepthStreamsResponse) GetSmalla() []RpiDiffBookDepthStreamsResponseAItem {
	if o == nil || common.IsNil(o.Smalla) {
		var ret []RpiDiffBookDepthStreamsResponseAItem
		return ret
	}
	return o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpiDiffBookDepthStreamsResponse) GetSmallaOk() ([]RpiDiffBookDepthStreamsResponseAItem, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *RpiDiffBookDepthStreamsResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given []RpiDiffBookDepthStreamsResponseAItem and assigns it to the A field.
func (o *RpiDiffBookDepthStreamsResponse) SetSmalla(v []RpiDiffBookDepthStreamsResponseAItem) {
	o.Smalla = v
}

func (o RpiDiffBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpiDiffBookDepthStreamsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *RpiDiffBookDepthStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varRpiDiffBookDepthStreamsResponse := _RpiDiffBookDepthStreamsResponse{}

	err = json.Unmarshal(data, &varRpiDiffBookDepthStreamsResponse)

	if err != nil {
		return err
	}

	*o = RpiDiffBookDepthStreamsResponse(varRpiDiffBookDepthStreamsResponse)

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

type NullableRpiDiffBookDepthStreamsResponse struct {
	value *RpiDiffBookDepthStreamsResponse
	isSet bool
}

func (v NullableRpiDiffBookDepthStreamsResponse) Get() *RpiDiffBookDepthStreamsResponse {
	return v.value
}

func (v *NullableRpiDiffBookDepthStreamsResponse) Set(val *RpiDiffBookDepthStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpiDiffBookDepthStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpiDiffBookDepthStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpiDiffBookDepthStreamsResponse(val *RpiDiffBookDepthStreamsResponse) *NullableRpiDiffBookDepthStreamsResponse {
	return &NullableRpiDiffBookDepthStreamsResponse{value: val, isSet: true}
}

func (v NullableRpiDiffBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpiDiffBookDepthStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
