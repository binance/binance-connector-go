/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CompositeIndexSymbolInformationStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CompositeIndexSymbolInformationStreamsResponse{}

// CompositeIndexSymbolInformationStreamsResponse struct for CompositeIndexSymbolInformationStreamsResponse
type CompositeIndexSymbolInformationStreamsResponse struct {
	Smalle               *string                                                `json:"e,omitempty"`
	E                    *int64                                                 `json:"E,omitempty"`
	Smalls               *string                                                `json:"s,omitempty"`
	Smallp               *string                                                `json:"p,omitempty"`
	C                    *string                                                `json:"C,omitempty"`
	Smallc               []CompositeIndexSymbolInformationStreamsResponseCInner `json:"c,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CompositeIndexSymbolInformationStreamsResponse CompositeIndexSymbolInformationStreamsResponse

// NewCompositeIndexSymbolInformationStreamsResponse instantiates a new CompositeIndexSymbolInformationStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompositeIndexSymbolInformationStreamsResponse() *CompositeIndexSymbolInformationStreamsResponse {
	this := CompositeIndexSymbolInformationStreamsResponse{}
	return &this
}

// NewCompositeIndexSymbolInformationStreamsResponseWithDefaults instantiates a new CompositeIndexSymbolInformationStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositeIndexSymbolInformationStreamsResponseWithDefaults() *CompositeIndexSymbolInformationStreamsResponse {
	this := CompositeIndexSymbolInformationStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetSmallp(v string) {
	o.Smallp = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetC() string {
	if o == nil || common.IsNil(o.C) {
		var ret string
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetCOk() (*string, bool) {
	if o == nil || common.IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasC() bool {
	if o != nil && !common.IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given string and assigns it to the C field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetC(v string) {
	o.C = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmallc() []CompositeIndexSymbolInformationStreamsResponseCInner {
	if o == nil || common.IsNil(o.Smallc) {
		var ret []CompositeIndexSymbolInformationStreamsResponseCInner
		return ret
	}
	return o.Smallc
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) GetSmallcOk() ([]CompositeIndexSymbolInformationStreamsResponseCInner, bool) {
	if o == nil || common.IsNil(o.Smallc) {
		return nil, false
	}
	return o.Smallc, true
}

// HasC returns a boolean if a field has been set.
func (o *CompositeIndexSymbolInformationStreamsResponse) HasSmallc() bool {
	if o != nil && !common.IsNil(o.Smallc) {
		return true
	}

	return false
}

// SetC gets a reference to the given []CompositeIndexSymbolInformationStreamsResponseCInner and assigns it to the C field.
func (o *CompositeIndexSymbolInformationStreamsResponse) SetSmallc(v []CompositeIndexSymbolInformationStreamsResponseCInner) {
	o.Smallc = v
}

func (o CompositeIndexSymbolInformationStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompositeIndexSymbolInformationStreamsResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}
	if !common.IsNil(o.C) {
		toSerialize["C"] = o.C
	}
	if !common.IsNil(o.Smallc) {
		toSerialize["c"] = o.Smallc
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CompositeIndexSymbolInformationStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varCompositeIndexSymbolInformationStreamsResponse := _CompositeIndexSymbolInformationStreamsResponse{}

	err = json.Unmarshal(data, &varCompositeIndexSymbolInformationStreamsResponse)

	if err != nil {
		return err
	}

	*o = CompositeIndexSymbolInformationStreamsResponse(varCompositeIndexSymbolInformationStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		delete(additionalProperties, "C")
		delete(additionalProperties, "c")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCompositeIndexSymbolInformationStreamsResponse struct {
	value *CompositeIndexSymbolInformationStreamsResponse
	isSet bool
}

func (v NullableCompositeIndexSymbolInformationStreamsResponse) Get() *CompositeIndexSymbolInformationStreamsResponse {
	return v.value
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponse) Set(val *CompositeIndexSymbolInformationStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCompositeIndexSymbolInformationStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompositeIndexSymbolInformationStreamsResponse(val *CompositeIndexSymbolInformationStreamsResponse) *NullableCompositeIndexSymbolInformationStreamsResponse {
	return &NullableCompositeIndexSymbolInformationStreamsResponse{value: val, isSet: true}
}

func (v NullableCompositeIndexSymbolInformationStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompositeIndexSymbolInformationStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
