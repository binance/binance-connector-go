/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AdlRiskResponse1 type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdlRiskResponse1{}

// AdlRiskResponse1 struct for AdlRiskResponse1
type AdlRiskResponse1 struct {
	Symbol               *string `json:"symbol,omitempty"`
	AdlRisk              *string `json:"adlRisk,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdlRiskResponse1 AdlRiskResponse1

// NewAdlRiskResponse1 instantiates a new AdlRiskResponse1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdlRiskResponse1() *AdlRiskResponse1 {
	this := AdlRiskResponse1{}
	return &this
}

// NewAdlRiskResponse1WithDefaults instantiates a new AdlRiskResponse1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdlRiskResponse1WithDefaults() *AdlRiskResponse1 {
	this := AdlRiskResponse1{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AdlRiskResponse1) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse1) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AdlRiskResponse1) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AdlRiskResponse1) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdlRisk returns the AdlRisk field value if set, zero value otherwise.
func (o *AdlRiskResponse1) GetAdlRisk() string {
	if o == nil || common.IsNil(o.AdlRisk) {
		var ret string
		return ret
	}
	return *o.AdlRisk
}

// GetAdlRiskOk returns a tuple with the AdlRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse1) GetAdlRiskOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdlRisk) {
		return nil, false
	}
	return o.AdlRisk, true
}

// HasAdlRisk returns a boolean if a field has been set.
func (o *AdlRiskResponse1) HasAdlRisk() bool {
	if o != nil && !common.IsNil(o.AdlRisk) {
		return true
	}

	return false
}

// SetAdlRisk gets a reference to the given string and assigns it to the AdlRisk field.
func (o *AdlRiskResponse1) SetAdlRisk(v string) {
	o.AdlRisk = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AdlRiskResponse1) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse1) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AdlRiskResponse1) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AdlRiskResponse1) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AdlRiskResponse1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdlRiskResponse1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.AdlRisk) {
		toSerialize["adlRisk"] = o.AdlRisk
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdlRiskResponse1) UnmarshalJSON(data []byte) (err error) {
	varAdlRiskResponse1 := _AdlRiskResponse1{}

	err = json.Unmarshal(data, &varAdlRiskResponse1)

	if err != nil {
		return err
	}

	*o = AdlRiskResponse1(varAdlRiskResponse1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adlRisk")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdlRiskResponse1 struct {
	value *AdlRiskResponse1
	isSet bool
}

func (v NullableAdlRiskResponse1) Get() *AdlRiskResponse1 {
	return v.value
}

func (v *NullableAdlRiskResponse1) Set(val *AdlRiskResponse1) {
	v.value = val
	v.isSet = true
}

func (v NullableAdlRiskResponse1) IsSet() bool {
	return v.isSet
}

func (v *NullableAdlRiskResponse1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdlRiskResponse1(val *AdlRiskResponse1) *NullableAdlRiskResponse1 {
	return &NullableAdlRiskResponse1{value: val, isSet: true}
}

func (v NullableAdlRiskResponse1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdlRiskResponse1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
