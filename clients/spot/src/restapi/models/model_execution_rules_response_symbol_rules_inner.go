/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponseSymbolRulesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponseSymbolRulesInner{}

// ExecutionRulesResponseSymbolRulesInner struct for ExecutionRulesResponseSymbolRulesInner
type ExecutionRulesResponseSymbolRulesInner struct {
	Symbol               *string                                            `json:"symbol,omitempty"`
	Rules                []ExecutionRulesResponseSymbolRulesInnerRulesInner `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponseSymbolRulesInner ExecutionRulesResponseSymbolRulesInner

// NewExecutionRulesResponseSymbolRulesInner instantiates a new ExecutionRulesResponseSymbolRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponseSymbolRulesInner() *ExecutionRulesResponseSymbolRulesInner {
	this := ExecutionRulesResponseSymbolRulesInner{}
	return &this
}

// NewExecutionRulesResponseSymbolRulesInnerWithDefaults instantiates a new ExecutionRulesResponseSymbolRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseSymbolRulesInnerWithDefaults() *ExecutionRulesResponseSymbolRulesInner {
	this := ExecutionRulesResponseSymbolRulesInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExecutionRulesResponseSymbolRulesInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInner) GetRules() []ExecutionRulesResponseSymbolRulesInnerRulesInner {
	if o == nil || common.IsNil(o.Rules) {
		var ret []ExecutionRulesResponseSymbolRulesInnerRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInner) GetRulesOk() ([]ExecutionRulesResponseSymbolRulesInnerRulesInner, bool) {
	if o == nil || common.IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInner) HasRules() bool {
	if o != nil && !common.IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ExecutionRulesResponseSymbolRulesInnerRulesInner and assigns it to the Rules field.
func (o *ExecutionRulesResponseSymbolRulesInner) SetRules(v []ExecutionRulesResponseSymbolRulesInnerRulesInner) {
	o.Rules = v
}

func (o ExecutionRulesResponseSymbolRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponseSymbolRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionRulesResponseSymbolRulesInner) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponseSymbolRulesInner := _ExecutionRulesResponseSymbolRulesInner{}

	err = json.Unmarshal(data, &varExecutionRulesResponseSymbolRulesInner)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponseSymbolRulesInner(varExecutionRulesResponseSymbolRulesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionRulesResponseSymbolRulesInner struct {
	value *ExecutionRulesResponseSymbolRulesInner
	isSet bool
}

func (v NullableExecutionRulesResponseSymbolRulesInner) Get() *ExecutionRulesResponseSymbolRulesInner {
	return v.value
}

func (v *NullableExecutionRulesResponseSymbolRulesInner) Set(val *ExecutionRulesResponseSymbolRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponseSymbolRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponseSymbolRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponseSymbolRulesInner(val *ExecutionRulesResponseSymbolRulesInner) *NullableExecutionRulesResponseSymbolRulesInner {
	return &NullableExecutionRulesResponseSymbolRulesInner{value: val, isSet: true}
}

func (v NullableExecutionRulesResponseSymbolRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponseSymbolRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
