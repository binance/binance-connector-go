/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FullDepthStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FullDepthStreamResponse{}

// FullDepthStreamResponse struct for FullDepthStreamResponse
type FullDepthStreamResponse struct {
	// Event type
	Smalle *string `json:"e,omitempty"`
	// Event time (ms)
	E *int64 `json:"E,omitempty"`
	// Matching time (ms)
	T *int64 `json:"T,omitempty"`
	// First updateId in this event
	U *int64 `json:"U,omitempty"`
	// Last updateId in this event
	Smallu *int64 `json:"u,omitempty"`
	// Previous updateId from the last push
	Smallpu *int64 `json:"pu,omitempty"`
	// Symbol
	Smalls *string `json:"s,omitempty"`
	// bids to be updated
	Smallb [][]string `json:"b,omitempty"`
	// asks to be updated
	Smalla               [][]string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FullDepthStreamResponse FullDepthStreamResponse

// NewFullDepthStreamResponse instantiates a new FullDepthStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFullDepthStreamResponse() *FullDepthStreamResponse {
	this := FullDepthStreamResponse{}
	return &this
}

// NewFullDepthStreamResponseWithDefaults instantiates a new FullDepthStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFullDepthStreamResponseWithDefaults() *FullDepthStreamResponse {
	this := FullDepthStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *FullDepthStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *FullDepthStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *FullDepthStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *FullDepthStreamResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *FullDepthStreamResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetPu returns the Pu field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmallpu() int64 {
	if o == nil || common.IsNil(o.Smallpu) {
		var ret int64
		return ret
	}
	return *o.Smallpu
}

// GetPuOk returns a tuple with the Pu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmallpuOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallpu) {
		return nil, false
	}
	return o.Smallpu, true
}

// HasPu returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmallpu() bool {
	if o != nil && !common.IsNil(o.Smallpu) {
		return true
	}

	return false
}

// SetPu gets a reference to the given int64 and assigns it to the Pu field.
func (o *FullDepthStreamResponse) SetSmallpu(v int64) {
	o.Smallpu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *FullDepthStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmallb() [][]string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret [][]string
		return ret
	}
	return o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmallbOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given [][]string and assigns it to the B field.
func (o *FullDepthStreamResponse) SetSmallb(v [][]string) {
	o.Smallb = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *FullDepthStreamResponse) GetSmalla() [][]string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret [][]string
		return ret
	}
	return o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FullDepthStreamResponse) GetSmallaOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *FullDepthStreamResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given [][]string and assigns it to the A field.
func (o *FullDepthStreamResponse) SetSmalla(v [][]string) {
	o.Smalla = v
}

func (o FullDepthStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FullDepthStreamResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smallpu) {
		toSerialize["pu"] = o.Smallpu
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
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

func (o *FullDepthStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varFullDepthStreamResponse := _FullDepthStreamResponse{}

	err = json.Unmarshal(data, &varFullDepthStreamResponse)

	if err != nil {
		return err
	}

	*o = FullDepthStreamResponse(varFullDepthStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "U")
		delete(additionalProperties, "u")
		delete(additionalProperties, "pu")
		delete(additionalProperties, "s")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFullDepthStreamResponse struct {
	value *FullDepthStreamResponse
	isSet bool
}

func (v NullableFullDepthStreamResponse) Get() *FullDepthStreamResponse {
	return v.value
}

func (v *NullableFullDepthStreamResponse) Set(val *FullDepthStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFullDepthStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFullDepthStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFullDepthStreamResponse(val *FullDepthStreamResponse) *NullableFullDepthStreamResponse {
	return &NullableFullDepthStreamResponse{value: val, isSet: true}
}

func (v NullableFullDepthStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFullDepthStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
