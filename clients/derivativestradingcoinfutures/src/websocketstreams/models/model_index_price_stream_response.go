/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexPriceStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexPriceStreamResponse{}

// IndexPriceStreamResponse struct for IndexPriceStreamResponse
type IndexPriceStreamResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smalli               *string `json:"i,omitempty"`
	Smallp               *string `json:"p,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexPriceStreamResponse IndexPriceStreamResponse

// NewIndexPriceStreamResponse instantiates a new IndexPriceStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexPriceStreamResponse() *IndexPriceStreamResponse {
	this := IndexPriceStreamResponse{}
	return &this
}

// NewIndexPriceStreamResponseWithDefaults instantiates a new IndexPriceStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexPriceStreamResponseWithDefaults() *IndexPriceStreamResponse {
	this := IndexPriceStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndexPriceStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexPriceStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexPriceStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndexPriceStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetI returns the I field value if set, zero value otherwise.
func (o *IndexPriceStreamResponse) GetSmalli() string {
	if o == nil || common.IsNil(o.Smalli) {
		var ret string
		return ret
	}
	return *o.Smalli
}

// GetIOk returns a tuple with the I field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamResponse) GetSmalliOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalli) {
		return nil, false
	}
	return o.Smalli, true
}

// HasI returns a boolean if a field has been set.
func (o *IndexPriceStreamResponse) HasSmalli() bool {
	if o != nil && !common.IsNil(o.Smalli) {
		return true
	}

	return false
}

// SetI gets a reference to the given string and assigns it to the I field.
func (o *IndexPriceStreamResponse) SetSmalli(v string) {
	o.Smalli = &v
}

// GetP returns the P field value if set, zero value otherwise.
func (o *IndexPriceStreamResponse) GetSmallp() string {
	if o == nil || common.IsNil(o.Smallp) {
		var ret string
		return ret
	}
	return *o.Smallp
}

// GetPOk returns a tuple with the P field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexPriceStreamResponse) GetSmallpOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallp) {
		return nil, false
	}
	return o.Smallp, true
}

// HasP returns a boolean if a field has been set.
func (o *IndexPriceStreamResponse) HasSmallp() bool {
	if o != nil && !common.IsNil(o.Smallp) {
		return true
	}

	return false
}

// SetP gets a reference to the given string and assigns it to the P field.
func (o *IndexPriceStreamResponse) SetSmallp(v string) {
	o.Smallp = &v
}

func (o IndexPriceStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexPriceStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smalli) {
		toSerialize["i"] = o.Smalli
	}
	if !common.IsNil(o.Smallp) {
		toSerialize["p"] = o.Smallp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexPriceStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varIndexPriceStreamResponse := _IndexPriceStreamResponse{}

	err = json.Unmarshal(data, &varIndexPriceStreamResponse)

	if err != nil {
		return err
	}

	*o = IndexPriceStreamResponse(varIndexPriceStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "i")
		delete(additionalProperties, "p")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexPriceStreamResponse struct {
	value *IndexPriceStreamResponse
	isSet bool
}

func (v NullableIndexPriceStreamResponse) Get() *IndexPriceStreamResponse {
	return v.value
}

func (v *NullableIndexPriceStreamResponse) Set(val *IndexPriceStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexPriceStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexPriceStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexPriceStreamResponse(val *IndexPriceStreamResponse) *NullableIndexPriceStreamResponse {
	return &NullableIndexPriceStreamResponse{value: val, isSet: true}
}

func (v NullableIndexPriceStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexPriceStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
