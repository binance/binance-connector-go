/*
Binance Derivatives Trading USDS Futures WebSocket Market Streams

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket Market Streams
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AllMarketLiquidationOrderStreamsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AllMarketLiquidationOrderStreamsResponse{}

// AllMarketLiquidationOrderStreamsResponse struct for AllMarketLiquidationOrderStreamsResponse
type AllMarketLiquidationOrderStreamsResponse struct {
	Smalle               *string                                    `json:"e,omitempty"`
	E                    *int64                                     `json:"E,omitempty"`
	Smallo               *AllMarketLiquidationOrderStreamsResponseO `json:"o,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AllMarketLiquidationOrderStreamsResponse AllMarketLiquidationOrderStreamsResponse

// NewAllMarketLiquidationOrderStreamsResponse instantiates a new AllMarketLiquidationOrderStreamsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllMarketLiquidationOrderStreamsResponse() *AllMarketLiquidationOrderStreamsResponse {
	this := AllMarketLiquidationOrderStreamsResponse{}
	return &this
}

// NewAllMarketLiquidationOrderStreamsResponseWithDefaults instantiates a new AllMarketLiquidationOrderStreamsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllMarketLiquidationOrderStreamsResponseWithDefaults() *AllMarketLiquidationOrderStreamsResponse {
	this := AllMarketLiquidationOrderStreamsResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponse) GetSmalle() string {
	if o == nil || common.IsNil(o.Smalle) {
		var ret string
		return ret
	}
	return *o.Smalle
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) GetSmalleOk() (*string, bool) {
	if o == nil || common.IsNil(o.Smalle) {
		return nil, false
	}
	return o.Smalle, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) HasSmalle() bool {
	if o != nil && !common.IsNil(o.Smalle) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AllMarketLiquidationOrderStreamsResponse) SetSmalle(v string) {
	o.Smalle = &v
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponse) GetE() int64 {
	if o == nil || common.IsNil(o.E) {
		var ret int64
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) GetEOk() (*int64, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given int64 and assigns it to the E field.
func (o *AllMarketLiquidationOrderStreamsResponse) SetE(v int64) {
	o.E = &v
}

// GetO returns the O field value if set, zero value otherwise.
func (o *AllMarketLiquidationOrderStreamsResponse) GetSmallo() AllMarketLiquidationOrderStreamsResponseO {
	if o == nil || common.IsNil(o.Smallo) {
		var ret AllMarketLiquidationOrderStreamsResponseO
		return ret
	}
	return *o.Smallo
}

// GetOOk returns a tuple with the O field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) GetSmalloOk() (*AllMarketLiquidationOrderStreamsResponseO, bool) {
	if o == nil || common.IsNil(o.Smallo) {
		return nil, false
	}
	return o.Smallo, true
}

// HasO returns a boolean if a field has been set.
func (o *AllMarketLiquidationOrderStreamsResponse) HasSmallo() bool {
	if o != nil && !common.IsNil(o.Smallo) {
		return true
	}

	return false
}

// SetO gets a reference to the given AllMarketLiquidationOrderStreamsResponseO and assigns it to the O field.
func (o *AllMarketLiquidationOrderStreamsResponse) SetSmallo(v AllMarketLiquidationOrderStreamsResponseO) {
	o.Smallo = &v
}

func (o AllMarketLiquidationOrderStreamsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllMarketLiquidationOrderStreamsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AllMarketLiquidationOrderStreamsResponse) UnmarshalJSON(data []byte) (err error) {
	varAllMarketLiquidationOrderStreamsResponse := _AllMarketLiquidationOrderStreamsResponse{}

	err = json.Unmarshal(data, &varAllMarketLiquidationOrderStreamsResponse)

	if err != nil {
		return err
	}

	*o = AllMarketLiquidationOrderStreamsResponse(varAllMarketLiquidationOrderStreamsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "E")
		delete(additionalProperties, "o")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAllMarketLiquidationOrderStreamsResponse struct {
	value *AllMarketLiquidationOrderStreamsResponse
	isSet bool
}

func (v NullableAllMarketLiquidationOrderStreamsResponse) Get() *AllMarketLiquidationOrderStreamsResponse {
	return v.value
}

func (v *NullableAllMarketLiquidationOrderStreamsResponse) Set(val *AllMarketLiquidationOrderStreamsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAllMarketLiquidationOrderStreamsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAllMarketLiquidationOrderStreamsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllMarketLiquidationOrderStreamsResponse(val *AllMarketLiquidationOrderStreamsResponse) *NullableAllMarketLiquidationOrderStreamsResponse {
	return &NullableAllMarketLiquidationOrderStreamsResponse{value: val, isSet: true}
}

func (v NullableAllMarketLiquidationOrderStreamsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllMarketLiquidationOrderStreamsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
