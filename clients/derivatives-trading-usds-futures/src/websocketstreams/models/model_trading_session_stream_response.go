/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TradingSessionStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingSessionStreamResponse{}

// TradingSessionStreamResponse struct for TradingSessionStreamResponse
type TradingSessionStreamResponse struct {
	Smalle               *string `json:"e,omitempty"`
	E                    *int64  `json:"E,omitempty"`
	Smallt               *int64  `json:"t,omitempty"`
	T                    *int64  `json:"T,omitempty"`
	S                    *string `json:"S,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingSessionStreamResponse TradingSessionStreamResponse

// NewTradingSessionStreamResponse instantiates a new TradingSessionStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingSessionStreamResponse() *TradingSessionStreamResponse {
	this := TradingSessionStreamResponse{}
	return &this
}

// NewTradingSessionStreamResponseWithDefaults instantiates a new TradingSessionStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingSessionStreamResponseWithDefaults() *TradingSessionStreamResponse {
	this := TradingSessionStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradingSessionStreamResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingSessionStreamResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *TradingSessionStreamResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *TradingSessionStreamResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *TradingSessionStreamResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingSessionStreamResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *TradingSessionStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *TradingSessionStreamResponse) SetE(v int64) {
	o.E = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradingSessionStreamResponse) GetSmallt() int64 {
	if o == nil || common.IsNil(o.Smallt) {
		var ret int64
		return ret
	}
	return *o.Smallt
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingSessionStreamResponse) GetSmalltOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Smallt) {
		return nil, false
	}
	return o.Smallt, true
}

// HasT returns a boolean if a field has been set.
func (o *TradingSessionStreamResponse) HasSmallt() bool {
	if o != nil && !common.IsNil(o.Smallt) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradingSessionStreamResponse) SetSmallt(v int64) {
	o.Smallt = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *TradingSessionStreamResponse) GetT() int64 {
	if o == nil || common.IsNil(o.T) {
		var ret int64
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingSessionStreamResponse) GetTOk() (*int64, bool) {
	if o == nil || common.IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *TradingSessionStreamResponse) HasT() bool {
	if o != nil && !common.IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given int64 and assigns it to the T field.
func (o *TradingSessionStreamResponse) SetT(v int64) {
	o.T = &v
}

// GetS returns the S field value if set, zero value otherwise.
func (o *TradingSessionStreamResponse) GetS() string {
	if o == nil || common.IsNil(o.S) {
		var ret string
		return ret
	}
	return *o.S
}

// GetSOk returns a tuple with the S field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingSessionStreamResponse) GetSOk() (*string, bool) {
	if o == nil || common.IsNil(o.S) {
		return nil, false
	}
	return o.S, true
}

// HasS returns a boolean if a field has been set.
func (o *TradingSessionStreamResponse) HasS() bool {
	if o != nil && !common.IsNil(o.S) {
		return true
	}

	return false
}

// SetS gets a reference to the given string and assigns it to the S field.
func (o *TradingSessionStreamResponse) SetS(v string) {
	o.S = &v
}

func (o TradingSessionStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingSessionStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallt) {
		toSerialize["t"] = o.Smallt
	}
	if !common.IsNil(o.T) {
		toSerialize["T"] = o.T
	}
	if !common.IsNil(o.S) {
		toSerialize["S"] = o.S
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingSessionStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varTradingSessionStreamResponse := _TradingSessionStreamResponse{}

	err = json.Unmarshal(data, &varTradingSessionStreamResponse)

	if err != nil {
		return err
	}

	*o = TradingSessionStreamResponse(varTradingSessionStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "t")
		delete(additionalProperties, "T")
		delete(additionalProperties, "S")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingSessionStreamResponse struct {
	value *TradingSessionStreamResponse
	isSet bool
}

func (v NullableTradingSessionStreamResponse) Get() *TradingSessionStreamResponse {
	return v.value
}

func (v *NullableTradingSessionStreamResponse) Set(val *TradingSessionStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingSessionStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingSessionStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingSessionStreamResponse(val *TradingSessionStreamResponse) *NullableTradingSessionStreamResponse {
	return &NullableTradingSessionStreamResponse{value: val, isSet: true}
}

func (v NullableTradingSessionStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingSessionStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
