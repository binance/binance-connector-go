/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the IndexKlineCandlestickStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &IndexKlineCandlestickStreamsResponse{}

// IndexKlineCandlestickStreamsResponse struct for IndexKlineCandlestickStreamsResponse
type IndexKlineCandlestickStreamsResponse struct {
	Smalle               *string                                `json:"e,omitempty"`
	E                    *int64                                 `json:"E,omitempty"`
	Smallps              *string                                `json:"ps,omitempty"`
	Smallk               *IndexKlineCandlestickStreamsResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IndexKlineCandlestickStreamsResponse IndexKlineCandlestickStreamsResponse

// NewIndexKlineCandlestickStreamsResponse instantiates a new IndexKlineCandlestickStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexKlineCandlestickStreamsResponse() *IndexKlineCandlestickStreamsResponse {
	this := IndexKlineCandlestickStreamsResponse{}
	return &this
}

// NewIndexKlineCandlestickStreamsResponseWithDefaults instantiates a new IndexKlineCandlestickStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexKlineCandlestickStreamsResponseWithDefaults() *IndexKlineCandlestickStreamsResponse {
	this := IndexKlineCandlestickStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *IndexKlineCandlestickStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *IndexKlineCandlestickStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *IndexKlineCandlestickStreamsResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *IndexKlineCandlestickStreamsResponse) GetSmallk() IndexKlineCandlestickStreamsResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret IndexKlineCandlestickStreamsResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexKlineCandlestickStreamsResponse) GetSmallkOk() (*IndexKlineCandlestickStreamsResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *IndexKlineCandlestickStreamsResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given IndexKlineCandlestickStreamsResponseK and assigns it to the K field.
func (o *IndexKlineCandlestickStreamsResponse) SetSmallk(v IndexKlineCandlestickStreamsResponseK) {
	o.Smallk = &v
}

func (o IndexKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndexKlineCandlestickStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallps) {
		toSerialize["ps"] = o.Smallps
	}
	if !common.IsNil(o.Smallk) {
		toSerialize["k"] = o.Smallk
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IndexKlineCandlestickStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varIndexKlineCandlestickStreamsResponse := _IndexKlineCandlestickStreamsResponse{}

	err = json.Unmarshal(data, &varIndexKlineCandlestickStreamsResponse)

	if err != nil {
		return err
	}

	*o = IndexKlineCandlestickStreamsResponse(varIndexKlineCandlestickStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "ps")
		delete(additionalProperties, "k")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIndexKlineCandlestickStreamsResponse struct {
	value *IndexKlineCandlestickStreamsResponse
	isSet bool
}

func (v NullableIndexKlineCandlestickStreamsResponse) Get() *IndexKlineCandlestickStreamsResponse {
	return v.value
}

func (v *NullableIndexKlineCandlestickStreamsResponse) Set(val *IndexKlineCandlestickStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexKlineCandlestickStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexKlineCandlestickStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexKlineCandlestickStreamsResponse(val *IndexKlineCandlestickStreamsResponse) *NullableIndexKlineCandlestickStreamsResponse {
	return &NullableIndexKlineCandlestickStreamsResponse{value: val, isSet: true}
}

func (v NullableIndexKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexKlineCandlestickStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
