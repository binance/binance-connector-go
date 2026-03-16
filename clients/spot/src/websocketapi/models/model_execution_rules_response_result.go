/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponseResult{}

// ExecutionRulesResponseResult struct for ExecutionRulesResponseResult
type ExecutionRulesResponseResult struct {
	SymbolRules          []ExecutionRulesResponseResultSymbolRulesInner `json:"symbolRules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponseResult ExecutionRulesResponseResult

// NewExecutionRulesResponseResult instantiates a new ExecutionRulesResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponseResult() *ExecutionRulesResponseResult {
	this := ExecutionRulesResponseResult{}
	return &this
}

// NewExecutionRulesResponseResultWithDefaults instantiates a new ExecutionRulesResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseResultWithDefaults() *ExecutionRulesResponseResult {
	this := ExecutionRulesResponseResult{}
	return &this
}

// GetSymbolRules returns the SymbolRules field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResult) GetSymbolRules() []ExecutionRulesResponseResultSymbolRulesInner {
	if o == nil || common.IsNil(o.SymbolRules) {
		var ret []ExecutionRulesResponseResultSymbolRulesInner
		return ret
	}
	return o.SymbolRules
}

// GetSymbolRulesOk returns a tuple with the SymbolRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResult) GetSymbolRulesOk() ([]ExecutionRulesResponseResultSymbolRulesInner, bool) {
	if o == nil || common.IsNil(o.SymbolRules) {
		return nil, false
	}
	return o.SymbolRules, true
}

// HasSymbolRules returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResult) HasSymbolRules() bool {
	if o != nil && !common.IsNil(o.SymbolRules) {
		return true
	}

	return false
}

// SetSymbolRules gets a reference to the given []ExecutionRulesResponseResultSymbolRulesInner and assigns it to the SymbolRules field.
func (o *ExecutionRulesResponseResult) SetSymbolRules(v []ExecutionRulesResponseResultSymbolRulesInner) {
	o.SymbolRules = v
}

func (o ExecutionRulesResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SymbolRules) {
		toSerialize["symbolRules"] = o.SymbolRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponseResult := _ExecutionRulesResponseResult{}

	err = json.Unmarshal(data, &varExecutionRulesResponseResult)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponseResult(varExecutionRulesResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbolRules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionRulesResponseResult struct {
	value *ExecutionRulesResponseResult
	isSet bool
}

func (v NullableExecutionRulesResponseResult) Get() *ExecutionRulesResponseResult {
	return v.value
}

func (v *NullableExecutionRulesResponseResult) Set(val *ExecutionRulesResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponseResult(val *ExecutionRulesResponseResult) *NullableExecutionRulesResponseResult {
	return &NullableExecutionRulesResponseResult{value: val, isSet: true}
}

func (v NullableExecutionRulesResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
