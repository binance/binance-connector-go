/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AdlRiskResponse2Inner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AdlRiskResponse2Inner{}

// AdlRiskResponse2Inner struct for AdlRiskResponse2Inner
type AdlRiskResponse2Inner struct {
	Symbol               *string `json:"symbol,omitempty"`
	AdlRisk              *string `json:"adlRisk,omitempty"`
	UpdateTime           *int64  `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdlRiskResponse2Inner AdlRiskResponse2Inner

// NewAdlRiskResponse2Inner instantiates a new AdlRiskResponse2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdlRiskResponse2Inner() *AdlRiskResponse2Inner {
	this := AdlRiskResponse2Inner{}
	return &this
}

// NewAdlRiskResponse2InnerWithDefaults instantiates a new AdlRiskResponse2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdlRiskResponse2InnerWithDefaults() *AdlRiskResponse2Inner {
	this := AdlRiskResponse2Inner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AdlRiskResponse2Inner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse2Inner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AdlRiskResponse2Inner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AdlRiskResponse2Inner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetAdlRisk returns the AdlRisk field value if set, zero value otherwise.
func (o *AdlRiskResponse2Inner) GetAdlRisk() string {
	if o == nil || common.IsNil(o.AdlRisk) {
		var ret string
		return ret
	}
	return *o.AdlRisk
}

// GetAdlRiskOk returns a tuple with the AdlRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse2Inner) GetAdlRiskOk() (*string, bool) {
	if o == nil || common.IsNil(o.AdlRisk) {
		return nil, false
	}
	return o.AdlRisk, true
}

// HasAdlRisk returns a boolean if a field has been set.
func (o *AdlRiskResponse2Inner) HasAdlRisk() bool {
	if o != nil && !common.IsNil(o.AdlRisk) {
		return true
	}

	return false
}

// SetAdlRisk gets a reference to the given string and assigns it to the AdlRisk field.
func (o *AdlRiskResponse2Inner) SetAdlRisk(v string) {
	o.AdlRisk = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *AdlRiskResponse2Inner) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdlRiskResponse2Inner) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *AdlRiskResponse2Inner) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *AdlRiskResponse2Inner) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o AdlRiskResponse2Inner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdlRiskResponse2Inner) ToMap() (map[string]interface{}, error) {
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

func (o *AdlRiskResponse2Inner) UnmarshalJSON(data []byte) (err error) {
	varAdlRiskResponse2Inner := _AdlRiskResponse2Inner{}

	err = json.Unmarshal(data, &varAdlRiskResponse2Inner)

	if err != nil {
		return err
	}

	*o = AdlRiskResponse2Inner(varAdlRiskResponse2Inner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "adlRisk")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdlRiskResponse2Inner struct {
	value *AdlRiskResponse2Inner
	isSet bool
}

func (v NullableAdlRiskResponse2Inner) Get() *AdlRiskResponse2Inner {
	return v.value
}

func (v *NullableAdlRiskResponse2Inner) Set(val *AdlRiskResponse2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullableAdlRiskResponse2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullableAdlRiskResponse2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdlRiskResponse2Inner(val *AdlRiskResponse2Inner) *NullableAdlRiskResponse2Inner {
	return &NullableAdlRiskResponse2Inner{value: val, isSet: true}
}

func (v NullableAdlRiskResponse2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdlRiskResponse2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
