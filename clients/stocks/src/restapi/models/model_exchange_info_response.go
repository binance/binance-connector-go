/*
Stocks Trading REST API

REST APIs for Binance Stocks Trading. All endpoints under `/sapi/v1/equity/_*`.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExchangeInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExchangeInfoResponse{}

// ExchangeInfoResponse struct for ExchangeInfoResponse
type ExchangeInfoResponse struct {
	// Always `UTC`.
	Timezone *string `json:"timezone,omitempty"`
	// List of tradable symbols with trading rules.
	Symbols              []ExchangeInfoResponseSymbolsInner `json:"symbols,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExchangeInfoResponse ExchangeInfoResponse

// NewExchangeInfoResponse instantiates a new ExchangeInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeInfoResponse() *ExchangeInfoResponse {
	this := ExchangeInfoResponse{}
	return &this
}

// NewExchangeInfoResponseWithDefaults instantiates a new ExchangeInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeInfoResponseWithDefaults() *ExchangeInfoResponse {
	this := ExchangeInfoResponse{}
	return &this
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ExchangeInfoResponse) GetTimezone() string {
	if o == nil || common.IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponse) GetTimezoneOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ExchangeInfoResponse) HasTimezone() bool {
	if o != nil && !common.IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ExchangeInfoResponse) SetTimezone(v string) {
	o.Timezone = &v
}

// GetSymbols returns the Symbols field value if set, zero value otherwise.
func (o *ExchangeInfoResponse) GetSymbols() []ExchangeInfoResponseSymbolsInner {
	if o == nil || common.IsNil(o.Symbols) {
		var ret []ExchangeInfoResponseSymbolsInner
		return ret
	}
	return o.Symbols
}

// GetSymbolsOk returns a tuple with the Symbols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeInfoResponse) GetSymbolsOk() ([]ExchangeInfoResponseSymbolsInner, bool) {
	if o == nil || common.IsNil(o.Symbols) {
		return nil, false
	}
	return o.Symbols, true
}

// HasSymbols returns a boolean if a field has been set.
func (o *ExchangeInfoResponse) HasSymbols() bool {
	if o != nil && !common.IsNil(o.Symbols) {
		return true
	}

	return false
}

// SetSymbols gets a reference to the given []ExchangeInfoResponseSymbolsInner and assigns it to the Symbols field.
func (o *ExchangeInfoResponse) SetSymbols(v []ExchangeInfoResponseSymbolsInner) {
	o.Symbols = v
}

func (o ExchangeInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !common.IsNil(o.Symbols) {
		toSerialize["symbols"] = o.Symbols
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExchangeInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varExchangeInfoResponse := _ExchangeInfoResponse{}

	err = json.Unmarshal(data, &varExchangeInfoResponse)

	if err != nil {
		return err
	}

	*o = ExchangeInfoResponse(varExchangeInfoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "symbols")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExchangeInfoResponse struct {
	value *ExchangeInfoResponse
	isSet bool
}

func (v NullableExchangeInfoResponse) Get() *ExchangeInfoResponse {
	return v.value
}

func (v *NullableExchangeInfoResponse) Set(val *ExchangeInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeInfoResponse(val *ExchangeInfoResponse) *NullableExchangeInfoResponse {
	return &NullableExchangeInfoResponse{value: val, isSet: true}
}

func (v NullableExchangeInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
