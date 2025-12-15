/*
Binance Derivatives Trading COIN Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MarkPriceKlineCandlestickStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MarkPriceKlineCandlestickStreamsResponse{}

// MarkPriceKlineCandlestickStreamsResponse struct for MarkPriceKlineCandlestickStreamsResponse
type MarkPriceKlineCandlestickStreamsResponse struct {
	Smalle               *string                                    `json:"e,omitempty"`
	E                    *int64                                     `json:"E,omitempty"`
	Smallps              *string                                    `json:"ps,omitempty"`
	Smallk               *MarkPriceKlineCandlestickStreamsResponseK `json:"k,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MarkPriceKlineCandlestickStreamsResponse MarkPriceKlineCandlestickStreamsResponse

// NewMarkPriceKlineCandlestickStreamsResponse instantiates a new MarkPriceKlineCandlestickStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkPriceKlineCandlestickStreamsResponse() *MarkPriceKlineCandlestickStreamsResponse {
	this := MarkPriceKlineCandlestickStreamsResponse{}
	return &this
}

// NewMarkPriceKlineCandlestickStreamsResponseWithDefaults instantiates a new MarkPriceKlineCandlestickStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkPriceKlineCandlestickStreamsResponseWithDefaults() *MarkPriceKlineCandlestickStreamsResponse {
	this := MarkPriceKlineCandlestickStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *MarkPriceKlineCandlestickStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *MarkPriceKlineCandlestickStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetPs returns the Ps field value if set, zero value otherwise.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmallps() string {
	if o == nil || common.IsNil(o.Smallps) {
		var ret string
		return ret
	}
	return *o.Smallps
}

// GetPsOk returns a tuple with the Ps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmallpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smallps) {
		return nil, false
	}
	return o.Smallps, true
}

// HasPs returns a boolean if a field has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) HasSmallps() bool {
	if o != nil && !common.IsNil(o.Smallps) {
		return true
	}

	return false
}

// SetPs gets a reference to the given string and assigns it to the Ps field.
func (o *MarkPriceKlineCandlestickStreamsResponse) SetSmallps(v string) {
	o.Smallps = &v
}

// GetK returns the K field value if set, zero value otherwise.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmallk() MarkPriceKlineCandlestickStreamsResponseK {
	if o == nil || common.IsNil(o.Smallk) {
		var ret MarkPriceKlineCandlestickStreamsResponseK
		return ret
	}
	return *o.Smallk
}

// GetKOk returns a tuple with the K field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) GetSmallkOk() (*MarkPriceKlineCandlestickStreamsResponseK, bool) {
	if o == nil || common.IsNil(o.Smallk) {
		return nil, false
	}
	return o.Smallk, true
}

// HasK returns a boolean if a field has been set.
func (o *MarkPriceKlineCandlestickStreamsResponse) HasSmallk() bool {
	if o != nil && !common.IsNil(o.Smallk) {
		return true
	}

	return false
}

// SetK gets a reference to the given MarkPriceKlineCandlestickStreamsResponseK and assigns it to the K field.
func (o *MarkPriceKlineCandlestickStreamsResponse) SetSmallk(v MarkPriceKlineCandlestickStreamsResponseK) {
	o.Smallk = &v
}

func (o MarkPriceKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkPriceKlineCandlestickStreamsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *MarkPriceKlineCandlestickStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varMarkPriceKlineCandlestickStreamsResponse := _MarkPriceKlineCandlestickStreamsResponse{}

	err = json.Unmarshal(data, &varMarkPriceKlineCandlestickStreamsResponse)

	if err != nil {
		return err
	}

	*o = MarkPriceKlineCandlestickStreamsResponse(varMarkPriceKlineCandlestickStreamsResponse)

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

type NullableMarkPriceKlineCandlestickStreamsResponse struct {
	value *MarkPriceKlineCandlestickStreamsResponse
	isSet bool
}

func (v NullableMarkPriceKlineCandlestickStreamsResponse) Get() *MarkPriceKlineCandlestickStreamsResponse {
	return v.value
}

func (v *NullableMarkPriceKlineCandlestickStreamsResponse) Set(val *MarkPriceKlineCandlestickStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkPriceKlineCandlestickStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkPriceKlineCandlestickStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkPriceKlineCandlestickStreamsResponse(val *MarkPriceKlineCandlestickStreamsResponse) *NullableMarkPriceKlineCandlestickStreamsResponse {
	return &NullableMarkPriceKlineCandlestickStreamsResponse{value: val, isSet: true}
}

func (v NullableMarkPriceKlineCandlestickStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkPriceKlineCandlestickStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
