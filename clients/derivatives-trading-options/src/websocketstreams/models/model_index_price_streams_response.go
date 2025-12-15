/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the IndexPriceStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceStreamsResponse{}

// IndexPriceStreamsResponse struct for IndexPriceStreamsResponse
type IndexPriceStreamsResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceStreamsResponse IndexPriceStreamsResponse

// NewIndexPriceStreamsResponse instantiates a new IndexPriceStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceStreamsResponse() *IndexPriceStreamsResponse {
	this := IndexPriceStreamsResponse{}
	return &this
}

// NewIndexPriceStreamsResponseWithDefaults instantiates a new IndexPriceStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceStreamsResponseWithDefaults() *IndexPriceStreamsResponse {
	this := IndexPriceStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndexPriceStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndexPriceStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponse) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponse) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponse) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *IndexPriceStreamsResponse) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *IndexPriceStreamsResponse) SetSmallp(v string) {
	o.Smallp = &v
}

func (o IndexPriceStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceStreamsResponse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexPriceStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceStreamsResponse := _IndexPriceStreamsResponse{}

	err = json.Unmarshal(data, &varIndexPriceStreamsResponse)

	if err != nil {
		return err
	}

	*o = IndexPriceStreamsResponse(varIndexPriceStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "s")
		delete(additionalProperties, "p")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexPriceStreamsResponse struct {
	value *IndexPriceStreamsResponse
	isSet bool
}

func (v NullableIndexPriceStreamsResponse) Get() *IndexPriceStreamsResponse {
	return v.value
}

func (v *NullableIndexPriceStreamsResponse) Set(val *IndexPriceStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceStreamsResponse(val *IndexPriceStreamsResponse) *NullableIndexPriceStreamsResponse {
	return &NullableIndexPriceStreamsResponse{value: val, isSet: true}
}

func (v NullableIndexPriceStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
