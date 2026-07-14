/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PartialDepthStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PartialDepthStreamResponse{}

// PartialDepthStreamResponse struct for PartialDepthStreamResponse
type PartialDepthStreamResponse struct {
	// eventType
	Smalle *string `json:"e,omitempty"`
	// eventTime
	E *int64 `json:"E,omitempty"`
	// transactionTime
	T *int64 `json:"T,omitempty"`
	// firstUpdateId
	U *int64 `json:"U,omitempty"`
	// lastUpdateId
	Smallu *int64 `json:"u,omitempty"`
	// previousUpdateId
	Smallpu *int64 `json:"pu,omitempty"`
	// symbol
	Smalls *string `json:"s,omitempty"`
	// bids to be updated
	Smallb [][]string `json:"b,omitempty"`
	// asks to be updated
	Smalla               [][]string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PartialDepthStreamResponse PartialDepthStreamResponse

// NewPartialDepthStreamResponse instantiates a new PartialDepthStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartialDepthStreamResponse() *PartialDepthStreamResponse {
	this := PartialDepthStreamResponse{}
	return &this
}

// NewPartialDepthStreamResponseWithDefaults instantiates a new PartialDepthStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartialDepthStreamResponseWithDefaults() *PartialDepthStreamResponse {
	this := PartialDepthStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *PartialDepthStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *PartialDepthStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *PartialDepthStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *PartialDepthStreamResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *PartialDepthStreamResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetPu returns the Pu field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmallpu() int64 {
	if o == nil || common.IsNil(o.Smallpu) {
		var ret int64
		return ret
	}
	return *o.Smallpu
}

// GetPuOk returns a tuple with the Pu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmallpuOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallpu) {
		return nil, false
	}
	return o.Smallpu, true
}

// HasPu returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmallpu() bool {
	if o != nil && !common.IsNil(o.Smallpu) {
		return true
	}

	return false
}

// SetPu gets a reference to the given int64 and assigns it to the Pu field.
func (o *PartialDepthStreamResponse) SetSmallpu(v int64) {
	o.Smallpu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *PartialDepthStreamResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmallb() [][]string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret [][]string
		return ret
	}
	return o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmallbOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given [][]string and assigns it to the B field.
func (o *PartialDepthStreamResponse) SetSmallb(v [][]string) {
	o.Smallb = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *PartialDepthStreamResponse) GetSmalla() [][]string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret [][]string
		return ret
	}
	return o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartialDepthStreamResponse) GetSmallaOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *PartialDepthStreamResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given [][]string and assigns it to the A field.
func (o *PartialDepthStreamResponse) SetSmalla(v [][]string) {
	o.Smalla = v
}

func (o PartialDepthStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PartialDepthStreamResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PartialDepthStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varPartialDepthStreamResponse := _PartialDepthStreamResponse{}

	err = json.Unmarshal(data, &varPartialDepthStreamResponse)

	if err != nil {
		return err
	}

	*o = PartialDepthStreamResponse(varPartialDepthStreamResponse)

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

type NullablePartialDepthStreamResponse struct {
	value *PartialDepthStreamResponse
	isSet bool
}

func (v NullablePartialDepthStreamResponse) Get() *PartialDepthStreamResponse {
	return v.value
}

func (v *NullablePartialDepthStreamResponse) Set(val *PartialDepthStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialDepthStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialDepthStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialDepthStreamResponse(val *PartialDepthStreamResponse) *NullablePartialDepthStreamResponse {
	return &NullablePartialDepthStreamResponse{value: val, isSet: true}
}

func (v NullablePartialDepthStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialDepthStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
