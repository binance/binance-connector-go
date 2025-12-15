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

// checks if the LiquidationOrderStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &LiquidationOrderStreamsResponse{}

// LiquidationOrderStreamsResponse struct for LiquidationOrderStreamsResponse
type LiquidationOrderStreamsResponse struct {
	Smalle               *string                                    `json:"e,omitempty"`
	E                    *int64                                     `json:"E,omitempty"`
	Smallo               *AllMarketLiquidationOrderStreamsResponseO `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LiquidationOrderStreamsResponse LiquidationOrderStreamsResponse

// NewLiquidationOrderStreamsResponse instantiates a new LiquidationOrderStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiquidationOrderStreamsResponse() *LiquidationOrderStreamsResponse {
	this := LiquidationOrderStreamsResponse{}
	return &this
}

// NewLiquidationOrderStreamsResponseWithDefaults instantiates a new LiquidationOrderStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiquidationOrderStreamsResponseWithDefaults() *LiquidationOrderStreamsResponse {
	this := LiquidationOrderStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *LiquidationOrderStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationOrderStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *LiquidationOrderStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *LiquidationOrderStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *LiquidationOrderStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationOrderStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *LiquidationOrderStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *LiquidationOrderStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *LiquidationOrderStreamsResponse) GetSmallo() AllMarketLiquidationOrderStreamsResponseO {
	if o == nil || common.IsNil(o.Smallo) {
		var ret AllMarketLiquidationOrderStreamsResponseO
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiquidationOrderStreamsResponse) GetSmalloOk() (*AllMarketLiquidationOrderStreamsResponseO, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *LiquidationOrderStreamsResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given AllMarketLiquidationOrderStreamsResponseO and assigns it to the O field.
func (o *LiquidationOrderStreamsResponse) SetSmallo(v AllMarketLiquidationOrderStreamsResponseO) {
	o.Smallo = &v
}

func (o LiquidationOrderStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiquidationOrderStreamsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Smalle) {
		toSerialize["e"] = o.Smalle
	}
	if !common.IsNil(o.E) {
		toSerialize["E"] = o.E
	}
	if !common.IsNil(o.Smallo) {
		toSerialize["o"] = o.Smallo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LiquidationOrderStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varLiquidationOrderStreamsResponse := _LiquidationOrderStreamsResponse{}

	err = json.Unmarshal(data, &varLiquidationOrderStreamsResponse)

	if err != nil {
		return err
	}

	*o = LiquidationOrderStreamsResponse(varLiquidationOrderStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLiquidationOrderStreamsResponse struct {
	value *LiquidationOrderStreamsResponse
	isSet bool
}

func (v NullableLiquidationOrderStreamsResponse) Get() *LiquidationOrderStreamsResponse {
	return v.value
}

func (v *NullableLiquidationOrderStreamsResponse) Set(val *LiquidationOrderStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiquidationOrderStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiquidationOrderStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiquidationOrderStreamsResponse(val *LiquidationOrderStreamsResponse) *NullableLiquidationOrderStreamsResponse {
	return &NullableLiquidationOrderStreamsResponse{value: val, isSet: true}
}

func (v NullableLiquidationOrderStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiquidationOrderStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
