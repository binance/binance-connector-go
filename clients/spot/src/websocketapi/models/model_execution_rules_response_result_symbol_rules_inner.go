/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponseResultSymbolRulesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponseResultSymbolRulesInner{}

// ExecutionRulesResponseResultSymbolRulesInner struct for ExecutionRulesResponseResultSymbolRulesInner
type ExecutionRulesResponseResultSymbolRulesInner struct {
	Symbol               *string                                                  `json:"symbol,omitempty"`
	Rules                []ExecutionRulesResponseResultSymbolRulesInnerRulesInner `json:"rules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponseResultSymbolRulesInner ExecutionRulesResponseResultSymbolRulesInner

// NewExecutionRulesResponseResultSymbolRulesInner instantiates a new ExecutionRulesResponseResultSymbolRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponseResultSymbolRulesInner() *ExecutionRulesResponseResultSymbolRulesInner {
	this := ExecutionRulesResponseResultSymbolRulesInner{}
	return &this
}

// NewExecutionRulesResponseResultSymbolRulesInnerWithDefaults instantiates a new ExecutionRulesResponseResultSymbolRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseResultSymbolRulesInnerWithDefaults() *ExecutionRulesResponseResultSymbolRulesInner {
	this := ExecutionRulesResponseResultSymbolRulesInner{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ExecutionRulesResponseResultSymbolRulesInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInner) GetRules() []ExecutionRulesResponseResultSymbolRulesInnerRulesInner {
	if o == nil || common.IsNil(o.Rules) {
		var ret []ExecutionRulesResponseResultSymbolRulesInnerRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInner) GetRulesOk() ([]ExecutionRulesResponseResultSymbolRulesInnerRulesInner, bool) {
	if o == nil || common.IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInner) HasRules() bool {
	if o != nil && !common.IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ExecutionRulesResponseResultSymbolRulesInnerRulesInner and assigns it to the Rules field.
func (o *ExecutionRulesResponseResultSymbolRulesInner) SetRules(v []ExecutionRulesResponseResultSymbolRulesInnerRulesInner) {
	o.Rules = v
}

func (o ExecutionRulesResponseResultSymbolRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponseResultSymbolRulesInner) ToMap() (map[string]interface{}, error) {
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

func (o *ExecutionRulesResponseResultSymbolRulesInner) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponseResultSymbolRulesInner := _ExecutionRulesResponseResultSymbolRulesInner{}

	err = json.Unmarshal(data, &varExecutionRulesResponseResultSymbolRulesInner)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponseResultSymbolRulesInner(varExecutionRulesResponseResultSymbolRulesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionRulesResponseResultSymbolRulesInner struct {
	value *ExecutionRulesResponseResultSymbolRulesInner
	isSet bool
}

func (v NullableExecutionRulesResponseResultSymbolRulesInner) Get() *ExecutionRulesResponseResultSymbolRulesInner {
	return v.value
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInner) Set(val *ExecutionRulesResponseResultSymbolRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponseResultSymbolRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponseResultSymbolRulesInner(val *ExecutionRulesResponseResultSymbolRulesInner) *NullableExecutionRulesResponseResultSymbolRulesInner {
	return &NullableExecutionRulesResponseResultSymbolRulesInner{value: val, isSet: true}
}

func (v NullableExecutionRulesResponseResultSymbolRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
