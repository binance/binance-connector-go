/*
Binance Derivatives Trading Options WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading Options WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceStreamsResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceStreamsResponseInner{}

// IndexPriceStreamsResponseInner struct for IndexPriceStreamsResponseInner
type IndexPriceStreamsResponseInner struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalls               *string `json:"s,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceStreamsResponseInner IndexPriceStreamsResponseInner

// NewIndexPriceStreamsResponseInner instantiates a new IndexPriceStreamsResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceStreamsResponseInner() *IndexPriceStreamsResponseInner {
	this := IndexPriceStreamsResponseInner{}
	return &this
}

// NewIndexPriceStreamsResponseInnerWithDefaults instantiates a new IndexPriceStreamsResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceStreamsResponseInnerWithDefaults() *IndexPriceStreamsResponseInner {
	this := IndexPriceStreamsResponseInner{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponseInner) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponseInner) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponseInner) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndexPriceStreamsResponseInner) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponseInner) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponseInner) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponseInner) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndexPriceStreamsResponseInner) SetE(v int64) {
	o.E = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponseInner) GetSmalls() string {
	if o == nil || common.IsNil(o.Smalls) {
		var ret string
		return ret
	}
	return *o.Smalls
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponseInner) GetSmallsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalls) {
		return nil, false
	}
	return o.Smalls, true
}

// HasS returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponseInner) HasSmalls() bool {
	if o != nil && !common.IsNil(o.Smalls) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *IndexPriceStreamsResponseInner) SetSmalls(v string) {
	o.Smalls = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *IndexPriceStreamsResponseInner) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamsResponseInner) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *IndexPriceStreamsResponseInner) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *IndexPriceStreamsResponseInner) SetSmallp(v string) {
	o.Smallp = &v
}

func (o IndexPriceStreamsResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceStreamsResponseInner) ToMap() (map[string]interface{}, error) {
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

func (o *IndexPriceStreamsResponseInner) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceStreamsResponseInner := _IndexPriceStreamsResponseInner{}

	err = json.Unmarshal(data, &varIndexPriceStreamsResponseInner)

	if err != nil {
		return err
	}

	*o = IndexPriceStreamsResponseInner(varIndexPriceStreamsResponseInner)

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

type NullableIndexPriceStreamsResponseInner struct {
	value *IndexPriceStreamsResponseInner
	isSet bool
}

func (v NullableIndexPriceStreamsResponseInner) Get() *IndexPriceStreamsResponseInner {
	return v.value
}

func (v *NullableIndexPriceStreamsResponseInner) Set(val *IndexPriceStreamsResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceStreamsResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceStreamsResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceStreamsResponseInner(val *IndexPriceStreamsResponseInner) *NullableIndexPriceStreamsResponseInner {
	return &NullableIndexPriceStreamsResponseInner{value: val, isSet: true}
}

func (v NullableIndexPriceStreamsResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceStreamsResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
