/*
Binance Derivatives Trading Options REST API

OpenAPI Specification for the Binance Derivatives Trading Options REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OpenInterestResponseInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OpenInterestResponseInner{}

// OpenInterestResponseInner struct for OpenInterestResponseInner
type OpenInterestResponseInner struct {
	Symbol               *string `json:"symbol,omitempty"`
	SumOpenInterest      *string `json:"sumOpenInterest,omitempty"`
	SumOpenInterestUsd   *string `json:"sumOpenInterestUsd,omitempty"`
	Timestamp            *string `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OpenInterestResponseInner OpenInterestResponseInner

// NewOpenInterestResponseInner instantiates a new OpenInterestResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenInterestResponseInner() *OpenInterestResponseInner {
	this := OpenInterestResponseInner{}
	return &this
}

// NewOpenInterestResponseInnerWithDefaults instantiates a new OpenInterestResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenInterestResponseInnerWithDefaults() *OpenInterestResponseInner {
	this := OpenInterestResponseInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *OpenInterestResponseInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSumOpenInterest returns the SumOpenInterest field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSumOpenInterest() string {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		var ret string
		return ret
	}
	return *o.SumOpenInterest
}

// GetSumOpenInterestOk returns a tuple with the SumOpenInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSumOpenInterestOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterest) {
		return nil, false
	}
	return o.SumOpenInterest, true
}

// HasSumOpenInterest returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSumOpenInterest() bool {
	if o != nil && !common.IsNil(o.SumOpenInterest) {
		return true
	}

	return false
}

// SetSumOpenInterest gets a reference to the given string and assigns it to the SumOpenInterest field.
func (o *OpenInterestResponseInner) SetSumOpenInterest(v string) {
	o.SumOpenInterest = &v
}

// GetSumOpenInterestUsd returns the SumOpenInterestUsd field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetSumOpenInterestUsd() string {
	if o == nil || common.IsNil(o.SumOpenInterestUsd) {
		var ret string
		return ret
	}
	return *o.SumOpenInterestUsd
}

// GetSumOpenInterestUsdOk returns a tuple with the SumOpenInterestUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetSumOpenInterestUsdOk() (*string, bool) {
	if o == nil || common.IsNil(o.SumOpenInterestUsd) {
		return nil, false
	}
	return o.SumOpenInterestUsd, true
}

// HasSumOpenInterestUsd returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasSumOpenInterestUsd() bool {
	if o != nil && !common.IsNil(o.SumOpenInterestUsd) {
		return true
	}

	return false
}

// SetSumOpenInterestUsd gets a reference to the given string and assigns it to the SumOpenInterestUsd field.
func (o *OpenInterestResponseInner) SetSumOpenInterestUsd(v string) {
	o.SumOpenInterestUsd = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *OpenInterestResponseInner) GetTimestamp() string {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenInterestResponseInner) GetTimestampOk() (*string, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *OpenInterestResponseInner) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *OpenInterestResponseInner) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o OpenInterestResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenInterestResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.SumOpenInterest) {
		toSerialize["sumOpenInterest"] = o.SumOpenInterest
	}
	if !common.IsNil(o.SumOpenInterestUsd) {
		toSerialize["sumOpenInterestUsd"] = o.SumOpenInterestUsd
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OpenInterestResponseInner) UnmarshalJSON(data []byte) (err error) {
	varOpenInterestResponseInner := _OpenInterestResponseInner{}

	err = json.Unmarshal(data, &varOpenInterestResponseInner)

	if err != nil {
		return err
	}

	*o = OpenInterestResponseInner(varOpenInterestResponseInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "sumOpenInterest")
		delete(additionalProperties, "sumOpenInterestUsd")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOpenInterestResponseInner struct {
	value *OpenInterestResponseInner
	isSet bool
}

func (v NullableOpenInterestResponseInner) Get() *OpenInterestResponseInner {
	return v.value
}

func (v *NullableOpenInterestResponseInner) Set(val *OpenInterestResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenInterestResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenInterestResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenInterestResponseInner(val *OpenInterestResponseInner) *NullableOpenInterestResponseInner {
	return &NullableOpenInterestResponseInner{value: val, isSet: true}
}

func (v NullableOpenInterestResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenInterestResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
