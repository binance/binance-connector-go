/*
Binance Spot WebSocket Streams

OpenAPI Specifications for the Binance Spot WebSocket Streams  API documents:   - [Github web-socket-streams documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)   - [General API information for web-socket-streams on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-streams)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DiffBookDepthResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DiffBookDepthResponse{}

// DiffBookDepthResponse struct for DiffBookDepthResponse
type DiffBookDepthResponse struct {
	Smalle               *string    `json:"e,omitempty"`
	E                    *int64     `json:"E,omitempty"`
	S                    *string    `json:"s,omitempty"`
	U                    *int64     `json:"U,omitempty"`
	Smallu               *int64     `json:"u,omitempty"`
	B                    [][]string `json:"b,omitempty"`
	A                    [][]string `json:"a,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DiffBookDepthResponse DiffBookDepthResponse

// NewDiffBookDepthResponse instantiates a new DiffBookDepthResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiffBookDepthResponse() *DiffBookDepthResponse {
	this := DiffBookDepthResponse{}
	return &this
}

// NewDiffBookDepthResponseWithDefaults instantiates a new DiffBookDepthResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiffBookDepthResponseWithDefaults() *DiffBookDepthResponse {
	this := DiffBookDepthResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *DiffBookDepthResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *DiffBookDepthResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *DiffBookDepthResponse) SetS(v string) {
	o.S = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetU() int64 {
	if o == nil || common.IsNil(o.U) {
		var ret int64
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetUOk() (*int64, bool) {
	if o == nil || common.IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasU() bool {
	if o != nil && !common.IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *DiffBookDepthResponse) SetU(v int64) {
	o.U = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *DiffBookDepthResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetB() [][]string {
	if o == nil || common.IsNil(o.B) {
		var ret [][]string
		return ret
	}
	return o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetBOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given [][]string and assigns it to the B field.
func (o *DiffBookDepthResponse) SetB(v [][]string) {
	o.B = v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *DiffBookDepthResponse) GetA() [][]string {
	if o == nil || common.IsNil(o.A) {
		var ret [][]string
		return ret
	}
	return o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiffBookDepthResponse) GetAOk() ([][]string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *DiffBookDepthResponse) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given [][]string and assigns it to the A field.
func (o *DiffBookDepthResponse) SetA(v [][]string) {
	o.A = v
}

func (o DiffBookDepthResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiffBookDepthResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.S) {
		toSerialize["s"] = o.S
	}
	if !common.IsNil(o.U) {
		toSerialize["U"] = o.U
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.B) {
		toSerialize["b"] = o.B
	}
	if !common.IsNil(o.A) {
		toSerialize["a"] = o.A
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DiffBookDepthResponse) UnmarshalJSON(data []byte) (err error) {
	varDiffBookDepthResponse := _DiffBookDepthResponse{}

	err = json.Unmarshal(data, &varDiffBookDepthResponse)

	if err != nil {
		return err
	}

	*o = DiffBookDepthResponse(varDiffBookDepthResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "U")
		delete(additionalProperties, "u")
		delete(additionalProperties, "b")
		delete(additionalProperties, "a")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDiffBookDepthResponse struct {
	value *DiffBookDepthResponse
	isSet bool
}

func (v NullableDiffBookDepthResponse) Get() *DiffBookDepthResponse {
	return v.value
}

func (v *NullableDiffBookDepthResponse) Set(val *DiffBookDepthResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDiffBookDepthResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDiffBookDepthResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiffBookDepthResponse(val *DiffBookDepthResponse) *NullableDiffBookDepthResponse {
	return &NullableDiffBookDepthResponse{value: val, isSet: true}
}

func (v NullableDiffBookDepthResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiffBookDepthResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
