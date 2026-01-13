/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the DiffBookDepthStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DiffBookDepthStreamsResponse{}

// DiffBookDepthStreamsResponse struct for DiffBookDepthStreamsResponse
type DiffBookDepthStreamsResponse struct {
	Smalle               *string                             `json:"e,omitempty"`
	E                    *int64                              `json:"E,omitempty"`
	T                    *int64                              `json:"T,omitempty"`
	Smalls               *string                             `json:"s,omitempty"`
	Smallps              *string                             `json:"ps,omitempty"`
	U                    *int64                              `json:"U,omitempty"`
	Smallu               *int64                              `json:"u,omitempty"`
	Smallpu              *int64                              `json:"pu,omitempty"`
	Smallb               []DiffBookDepthStreamsResponseBItem `json:"b,omitempty"`
	Smalla               []DiffBookDepthStreamsResponseAItem `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiffBookDepthStreamsResponse DiffBookDepthStreamsResponse

// NewDiffBookDepthStreamsResponse instantiates a new DiffBookDepthStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffBookDepthStreamsResponse() *DiffBookDepthStreamsResponse {
	this := DiffBookDepthStreamsResponse{}
	return &this
}

// NewDiffBookDepthStreamsResponseWithDefaults instantiates a new DiffBookDepthStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffBookDepthStreamsResponseWithDefaults() *DiffBookDepthStreamsResponse {
	this := DiffBookDepthStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *DiffBookDepthStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *DiffBookDepthStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *DiffBookDepthStreamsResponse) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *DiffBookDepthStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *DiffBookDepthStreamsResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *DiffBookDepthStreamsResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *DiffBookDepthStreamsResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetPu returns the Pu field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmallpu() int64 {
	if o == nil || common.IsNil(o.Smallpu) {
		var ret int64
		return ret
	}
	return *o.Smallpu
}

// GetPuOk returns a tuple with the Pu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmallpuOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallpu) {
		return nil, false
	}
	return o.Smallpu, true
}

// HasPu returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmallpu() bool {
	if o != nil && !common.IsNil(o.Smallpu) {
		return true
	}

	return false
}

// SetPu gets a reference to the given int64 and assigns it to the Pu field.
func (o *DiffBookDepthStreamsResponse) SetSmallpu(v int64) {
	o.Smallpu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmallb() []DiffBookDepthStreamsResponseBItem {
	if o == nil || common.IsNil(o.Smallb) {
		var ret []DiffBookDepthStreamsResponseBItem
		return ret
	}
	return o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmallbOk() ([]DiffBookDepthStreamsResponseBItem, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given []DiffBookDepthStreamsResponseBItem and assigns it to the B field.
func (o *DiffBookDepthStreamsResponse) SetSmallb(v []DiffBookDepthStreamsResponseBItem) {
	o.Smallb = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *DiffBookDepthStreamsResponse) GetSmalla() []DiffBookDepthStreamsResponseAItem {
	if o == nil || common.IsNil(o.Smalla) {
		var ret []DiffBookDepthStreamsResponseAItem
		return ret
	}
	return o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthStreamsResponse) GetSmallaOk() ([]DiffBookDepthStreamsResponseAItem, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *DiffBookDepthStreamsResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given []DiffBookDepthStreamsResponseAItem and assigns it to the A field.
func (o *DiffBookDepthStreamsResponse) SetSmalla(v []DiffBookDepthStreamsResponseAItem) {
	o.Smalla = v
}

func (o DiffBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiffBookDepthStreamsResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
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

func (o *DiffBookDepthStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varDiffBookDepthStreamsResponse := _DiffBookDepthStreamsResponse{}

	err = json.Unmarshal(data, &varDiffBookDepthStreamsResponse)

	if err != nil {
		return err
	}

	*o = DiffBookDepthStreamsResponse(varDiffBookDepthStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "T")
		delete(additionalProperties, "s")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "U")
		delete(additionalProperties, "u")
		delete(additionalProperties, "pu")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiffBookDepthStreamsResponse struct {
	value *DiffBookDepthStreamsResponse
	isSet bool
}

func (v NullableDiffBookDepthStreamsResponse) Get() *DiffBookDepthStreamsResponse {
	return v.value
}

func (v *NullableDiffBookDepthStreamsResponse) Set(val *DiffBookDepthStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiffBookDepthStreamsResponse(val *DiffBookDepthStreamsResponse) *NullableDiffBookDepthStreamsResponse {
	return &NullableDiffBookDepthStreamsResponse{value: val, isSet: true}
}

func (v NullableDiffBookDepthStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
