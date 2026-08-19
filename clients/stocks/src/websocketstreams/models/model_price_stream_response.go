/*
Stocks Trading WebSocket Streams

WebSocket stream definitions for Binance Stocks Trading. Base URL: wss://nbstream.binance.com/equity
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the PriceStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PriceStreamResponse{}

// PriceStreamResponse struct for PriceStreamResponse
type PriceStreamResponse struct {
	// Event type, always `\"price\"`.
	E *string `json:"e,omitempty"`
	// One entry per symbol.
	Rates                []PriceStreamResponseRatesInner `json:"rates,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PriceStreamResponse PriceStreamResponse

// NewPriceStreamResponse instantiates a new PriceStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceStreamResponse() *PriceStreamResponse {
	this := PriceStreamResponse{}
	return &this
}

// NewPriceStreamResponseWithDefaults instantiates a new PriceStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceStreamResponseWithDefaults() *PriceStreamResponse {
	this := PriceStreamResponse{}
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *PriceStreamResponse) GetE() string {
	if o == nil || common.IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponse) GetEOk() (*string, bool) {
	if o == nil || common.IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *PriceStreamResponse) HasE() bool {
	if o != nil && !common.IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *PriceStreamResponse) SetE(v string) {
	o.E = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *PriceStreamResponse) GetRates() []PriceStreamResponseRatesInner {
	if o == nil || common.IsNil(o.Rates) {
		var ret []PriceStreamResponseRatesInner
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceStreamResponse) GetRatesOk() ([]PriceStreamResponseRatesInner, bool) {
	if o == nil || common.IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *PriceStreamResponse) HasRates() bool {
	if o != nil && !common.IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given []PriceStreamResponseRatesInner and assigns it to the Rates field.
func (o *PriceStreamResponse) SetRates(v []PriceStreamResponseRatesInner) {
	o.Rates = v
}

func (o PriceStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	if !common.IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PriceStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varPriceStreamResponse := _PriceStreamResponse{}

	err = json.Unmarshal(data, &varPriceStreamResponse)

	if err != nil {
		return err
	}

	*o = PriceStreamResponse(varPriceStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "rates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePriceStreamResponse struct {
	value *PriceStreamResponse
	isSet bool
}

func (v NullablePriceStreamResponse) Get() *PriceStreamResponse {
	return v.value
}

func (v *NullablePriceStreamResponse) Set(val *PriceStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceStreamResponse(val *PriceStreamResponse) *NullablePriceStreamResponse {
	return &NullablePriceStreamResponse{value: val, isSet: true}
}

func (v NullablePriceStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
