/*
Alpha WebSocket Market Streams

Access Alpha market streams over WebSocket.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllTokens24hTickerStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllTokens24hTickerStreamResponse{}

// AllTokens24hTickerStreamResponse struct for AllTokens24hTickerStreamResponse
type AllTokens24hTickerStreamResponse struct {
	// Event type
	E *string `json:"e,omitempty"`
	// Per-token 24-hour ticker metrics
	D                    []AllTokens24hTickerStreamResponseDInner `json:"d,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllTokens24hTickerStreamResponse AllTokens24hTickerStreamResponse

// NewAllTokens24hTickerStreamResponse instantiates a new AllTokens24hTickerStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllTokens24hTickerStreamResponse() *AllTokens24hTickerStreamResponse {
	this := AllTokens24hTickerStreamResponse{}
	return &this
}

// NewAllTokens24hTickerStreamResponseWithDefaults instantiates a new AllTokens24hTickerStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllTokens24hTickerStreamResponseWithDefaults() *AllTokens24hTickerStreamResponse {
	this := AllTokens24hTickerStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllTokens24hTickerStreamResponse) SetE(v string) {
	o.E = &v
}

// GetD returns the D field value if set, zero value otherwise.
func (o *AllTokens24hTickerStreamResponse) GetD() []AllTokens24hTickerStreamResponseDInner {
	if o == nil || common.IsNil(o.D) {
		var ret []AllTokens24hTickerStreamResponseDInner
		return ret
	}
	return o.D
}

// GetDOk returns a tuple with the D field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllTokens24hTickerStreamResponse) GetDOk() ([]AllTokens24hTickerStreamResponseDInner, bool) {
	if o == nil || common.IsNil(o.D) {
		return nil, false
	}
	return o.D, true
}

// HasD returns a boolean if a field has been set.
func (o *AllTokens24hTickerStreamResponse) HasD() bool {
	if o != nil && !common.IsNil(o.D) {
		return true
	}

	return false
}

// SetD gets a reference to the given []AllTokens24hTickerStreamResponseDInner and assigns it to the D field.
func (o *AllTokens24hTickerStreamResponse) SetD(v []AllTokens24hTickerStreamResponseDInner) {
	o.D = v
}

func (o AllTokens24hTickerStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllTokens24hTickerStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.D) {
		toSerialize["d"] = o.D
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AllTokens24hTickerStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varAllTokens24hTickerStreamResponse := _AllTokens24hTickerStreamResponse{}

	err = json.Unmarshal(data, &varAllTokens24hTickerStreamResponse)

	if err != nil {
		return err
	}

	*o = AllTokens24hTickerStreamResponse(varAllTokens24hTickerStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "d")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllTokens24hTickerStreamResponse struct {
	value *AllTokens24hTickerStreamResponse
	isSet bool
}

func (v NullableAllTokens24hTickerStreamResponse) Get() *AllTokens24hTickerStreamResponse {
	return v.value
}

func (v *NullableAllTokens24hTickerStreamResponse) Set(val *AllTokens24hTickerStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllTokens24hTickerStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllTokens24hTickerStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllTokens24hTickerStreamResponse(val *AllTokens24hTickerStreamResponse) *NullableAllTokens24hTickerStreamResponse {
	return &NullableAllTokens24hTickerStreamResponse{value: val, isSet: true}
}

func (v NullableAllTokens24hTickerStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllTokens24hTickerStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
