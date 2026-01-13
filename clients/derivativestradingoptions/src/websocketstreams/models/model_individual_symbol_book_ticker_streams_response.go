/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndividualSymbolBookTickerStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndividualSymbolBookTickerStreamsResponse{}

// IndividualSymbolBookTickerStreamsResponse struct for IndividualSymbolBookTickerStreamsResponse
type IndividualSymbolBookTickerStreamsResponse struct {
	Smalle               *string `json:"e,omitempty"`
	Smallu               *int64  `json:"u,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallb               *string `json:"b,omitempty"`
	B                    *string `json:"B,omitempty"`
	Smalla               *string `json:"a,omitempty"`
	A                    *string `json:"A,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndividualSymbolBookTickerStreamsResponse IndividualSymbolBookTickerStreamsResponse

// NewIndividualSymbolBookTickerStreamsResponse instantiates a new IndividualSymbolBookTickerStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndividualSymbolBookTickerStreamsResponse() *IndividualSymbolBookTickerStreamsResponse {
	this := IndividualSymbolBookTickerStreamsResponse{}
	return &this
}

// NewIndividualSymbolBookTickerStreamsResponseWithDefaults instantiates a new IndividualSymbolBookTickerStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndividualSymbolBookTickerStreamsResponseWithDefaults() *IndividualSymbolBookTickerStreamsResponse {
	this := IndividualSymbolBookTickerStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmallu() int64 {
	if o == nil || common.IsNil(o.Smallu) {
		var ret int64
		return ret
	}
	return *o.Smallu
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmalluOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallu) {
		return nil, false
	}
	return o.Smallu, true
}

// HasU returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasSmallu() bool {
	if o != nil && !common.IsNil(o.Smallu) {
		return true
	}

	return false
}

// SetU gets a reference to the given int64 and assigns it to the U field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetSmallu(v int64) {
	o.Smallu = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmallb() string {
	if o == nil || common.IsNil(o.Smallb) {
		var ret string
		return ret
	}
	return *o.Smallb
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmallbOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallb) {
		return nil, false
	}
	return o.Smallb, true
}

// HasB returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasSmallb() bool {
	if o != nil && !common.IsNil(o.Smallb) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetSmallb(v string) {
	o.Smallb = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetB() string {
	if o == nil || common.IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetBOk() (*string, bool) {
	if o == nil || common.IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasB() bool {
	if o != nil && !common.IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetB(v string) {
	o.B = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmalla() string {
	if o == nil || common.IsNil(o.Smalla) {
		var ret string
		return ret
	}
	return *o.Smalla
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetSmallaOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalla) {
		return nil, false
	}
	return o.Smalla, true
}

// HasA returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasSmalla() bool {
	if o != nil && !common.IsNil(o.Smalla) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetSmalla(v string) {
	o.Smalla = &v
}

// GetA returns the A field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetA() string {
	if o == nil || common.IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetAOk() (*string, bool) {
	if o == nil || common.IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasA() bool {
	if o != nil && !common.IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetA(v string) {
	o.A = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetT(v int64) {
	o.T = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndividualSymbolBookTickerStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndividualSymbolBookTickerStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndividualSymbolBookTickerStreamsResponse) SetE(v int64) {
	o.E = &v
}

func (o IndividualSymbolBookTickerStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndividualSymbolBookTickerStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.Smallu) {
		toSerialize["u"] = o.Smallu
	}
	if !common.IsNil(o.Smalls) {
		toSerialize["s"] = o.Smalls
	}
	if !common.IsNil(o.Smallb) {
		toSerialize["b"] = o.Smallb
	}
	if !common.IsNil(o.B) {
		toSerialize["B"] = o.B
	}
	if !common.IsNil(o.Smalla) {
		toSerialize["a"] = o.Smalla
	}
	if !common.IsNil(o.A) {
		toSerialize["A"] = o.A
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndividualSymbolBookTickerStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varIndividualSymbolBookTickerStreamsResponse := _IndividualSymbolBookTickerStreamsResponse{}

	err = json.Unmarshal(data, &varIndividualSymbolBookTickerStreamsResponse)

	if err != nil {
		return err
	}

	*o = IndividualSymbolBookTickerStreamsResponse(varIndividualSymbolBookTickerStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "u")
		delete(additionalProperties, "s")
		delete(additionalProperties, "b")
		delete(additionalProperties, "B")
		delete(additionalProperties, "a")
		delete(additionalProperties, "A")
		delete(additionalProperties, "T")
		delete(additionalProperties, "E")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndividualSymbolBookTickerStreamsResponse struct {
	value *IndividualSymbolBookTickerStreamsResponse
	isSet bool
}

func (v NullableIndividualSymbolBookTickerStreamsResponse) Get() *IndividualSymbolBookTickerStreamsResponse {
	return v.value
}

func (v *NullableIndividualSymbolBookTickerStreamsResponse) Set(val *IndividualSymbolBookTickerStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndividualSymbolBookTickerStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndividualSymbolBookTickerStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndividualSymbolBookTickerStreamsResponse(val *IndividualSymbolBookTickerStreamsResponse) *NullableIndividualSymbolBookTickerStreamsResponse {
	return &NullableIndividualSymbolBookTickerStreamsResponse{value: val, isSet: true}
}

func (v NullableIndividualSymbolBookTickerStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndividualSymbolBookTickerStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
