/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponse{}

// ExecutionRulesResponse struct for ExecutionRulesResponse
type ExecutionRulesResponse struct {
	SymbolRules          []ExecutionRulesResponseSymbolRulesInner `json:"symbolRules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponse ExecutionRulesResponse

// NewExecutionRulesResponse instantiates a new ExecutionRulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponse() *ExecutionRulesResponse {
	this := ExecutionRulesResponse{}
	return &this
}

// NewExecutionRulesResponseWithDefaults instantiates a new ExecutionRulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseWithDefaults() *ExecutionRulesResponse {
	this := ExecutionRulesResponse{}
	return &this
}

// GetSymbolRules returns the SymbolRules field value if set, zero value otherwise.
func (o *ExecutionRulesResponse) GetSymbolRules() []ExecutionRulesResponseSymbolRulesInner {
	if o == nil || common.IsNil(o.SymbolRules) {
		var ret []ExecutionRulesResponseSymbolRulesInner
		return ret
	}
	return o.SymbolRules
}

// GetSymbolRulesOk returns a tuple with the SymbolRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponse) GetSymbolRulesOk() ([]ExecutionRulesResponseSymbolRulesInner, bool) {
	if o == nil || common.IsNil(o.SymbolRules) {
		return nil, false
	}
	return o.SymbolRules, true
}

// HasSymbolRules returns a boolean if a field has been set.
func (o *ExecutionRulesResponse) HasSymbolRules() bool {
	if o != nil && !common.IsNil(o.SymbolRules) {
		return true
	}

	return false
}

// SetSymbolRules gets a reference to the given []ExecutionRulesResponseSymbolRulesInner and assigns it to the SymbolRules field.
func (o *ExecutionRulesResponse) SetSymbolRules(v []ExecutionRulesResponseSymbolRulesInner) {
	o.SymbolRules = v
}

func (o ExecutionRulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SymbolRules) {
		toSerialize["symbolRules"] = o.SymbolRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionRulesResponse) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponse := _ExecutionRulesResponse{}

	err = json.Unmarshal(data, &varExecutionRulesResponse)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponse(varExecutionRulesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbolRules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionRulesResponse struct {
	value *ExecutionRulesResponse
	isSet bool
}

func (v NullableExecutionRulesResponse) Get() *ExecutionRulesResponse {
	return v.value
}

func (v *NullableExecutionRulesResponse) Set(val *ExecutionRulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponse(val *ExecutionRulesResponse) *NullableExecutionRulesResponse {
	return &NullableExecutionRulesResponse{value: val, isSet: true}
}

func (v NullableExecutionRulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
