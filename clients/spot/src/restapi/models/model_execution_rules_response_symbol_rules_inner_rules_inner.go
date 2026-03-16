/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponseSymbolRulesInnerRulesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponseSymbolRulesInnerRulesInner{}

// ExecutionRulesResponseSymbolRulesInnerRulesInner struct for ExecutionRulesResponseSymbolRulesInnerRulesInner
type ExecutionRulesResponseSymbolRulesInnerRulesInner struct {
	RuleType             *string `json:"ruleType,omitempty"`
	BidLimitMultUp       *string `json:"bidLimitMultUp,omitempty"`
	BidLimitMultDown     *string `json:"bidLimitMultDown,omitempty"`
	AskLimitMultUp       *string `json:"askLimitMultUp,omitempty"`
	AskLimitMultDown     *string `json:"askLimitMultDown,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponseSymbolRulesInnerRulesInner ExecutionRulesResponseSymbolRulesInnerRulesInner

// NewExecutionRulesResponseSymbolRulesInnerRulesInner instantiates a new ExecutionRulesResponseSymbolRulesInnerRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponseSymbolRulesInnerRulesInner() *ExecutionRulesResponseSymbolRulesInnerRulesInner {
	this := ExecutionRulesResponseSymbolRulesInnerRulesInner{}
	return &this
}

// NewExecutionRulesResponseSymbolRulesInnerRulesInnerWithDefaults instantiates a new ExecutionRulesResponseSymbolRulesInnerRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseSymbolRulesInnerRulesInnerWithDefaults() *ExecutionRulesResponseSymbolRulesInnerRulesInner {
	this := ExecutionRulesResponseSymbolRulesInnerRulesInner{}
	return &this
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetRuleType() string {
	if o == nil || common.IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetRuleTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) HasRuleType() bool {
	if o != nil && !common.IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) SetRuleType(v string) {
	o.RuleType = &v
}

// GetBidLimitMultUp returns the BidLimitMultUp field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetBidLimitMultUp() string {
	if o == nil || common.IsNil(o.BidLimitMultUp) {
		var ret string
		return ret
	}
	return *o.BidLimitMultUp
}

// GetBidLimitMultUpOk returns a tuple with the BidLimitMultUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetBidLimitMultUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidLimitMultUp) {
		return nil, false
	}
	return o.BidLimitMultUp, true
}

// HasBidLimitMultUp returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) HasBidLimitMultUp() bool {
	if o != nil && !common.IsNil(o.BidLimitMultUp) {
		return true
	}

	return false
}

// SetBidLimitMultUp gets a reference to the given string and assigns it to the BidLimitMultUp field.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) SetBidLimitMultUp(v string) {
	o.BidLimitMultUp = &v
}

// GetBidLimitMultDown returns the BidLimitMultDown field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetBidLimitMultDown() string {
	if o == nil || common.IsNil(o.BidLimitMultDown) {
		var ret string
		return ret
	}
	return *o.BidLimitMultDown
}

// GetBidLimitMultDownOk returns a tuple with the BidLimitMultDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetBidLimitMultDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidLimitMultDown) {
		return nil, false
	}
	return o.BidLimitMultDown, true
}

// HasBidLimitMultDown returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) HasBidLimitMultDown() bool {
	if o != nil && !common.IsNil(o.BidLimitMultDown) {
		return true
	}

	return false
}

// SetBidLimitMultDown gets a reference to the given string and assigns it to the BidLimitMultDown field.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) SetBidLimitMultDown(v string) {
	o.BidLimitMultDown = &v
}

// GetAskLimitMultUp returns the AskLimitMultUp field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetAskLimitMultUp() string {
	if o == nil || common.IsNil(o.AskLimitMultUp) {
		var ret string
		return ret
	}
	return *o.AskLimitMultUp
}

// GetAskLimitMultUpOk returns a tuple with the AskLimitMultUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetAskLimitMultUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskLimitMultUp) {
		return nil, false
	}
	return o.AskLimitMultUp, true
}

// HasAskLimitMultUp returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) HasAskLimitMultUp() bool {
	if o != nil && !common.IsNil(o.AskLimitMultUp) {
		return true
	}

	return false
}

// SetAskLimitMultUp gets a reference to the given string and assigns it to the AskLimitMultUp field.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) SetAskLimitMultUp(v string) {
	o.AskLimitMultUp = &v
}

// GetAskLimitMultDown returns the AskLimitMultDown field value if set, zero value otherwise.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetAskLimitMultDown() string {
	if o == nil || common.IsNil(o.AskLimitMultDown) {
		var ret string
		return ret
	}
	return *o.AskLimitMultDown
}

// GetAskLimitMultDownOk returns a tuple with the AskLimitMultDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) GetAskLimitMultDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskLimitMultDown) {
		return nil, false
	}
	return o.AskLimitMultDown, true
}

// HasAskLimitMultDown returns a boolean if a field has been set.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) HasAskLimitMultDown() bool {
	if o != nil && !common.IsNil(o.AskLimitMultDown) {
		return true
	}

	return false
}

// SetAskLimitMultDown gets a reference to the given string and assigns it to the AskLimitMultDown field.
func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) SetAskLimitMultDown(v string) {
	o.AskLimitMultDown = &v
}

func (o ExecutionRulesResponseSymbolRulesInnerRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponseSymbolRulesInnerRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.RuleType) {
		toSerialize["ruleType"] = o.RuleType
	}
	if !common.IsNil(o.BidLimitMultUp) {
		toSerialize["bidLimitMultUp"] = o.BidLimitMultUp
	}
	if !common.IsNil(o.BidLimitMultDown) {
		toSerialize["bidLimitMultDown"] = o.BidLimitMultDown
	}
	if !common.IsNil(o.AskLimitMultUp) {
		toSerialize["askLimitMultUp"] = o.AskLimitMultUp
	}
	if !common.IsNil(o.AskLimitMultDown) {
		toSerialize["askLimitMultDown"] = o.AskLimitMultDown
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExecutionRulesResponseSymbolRulesInnerRulesInner) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponseSymbolRulesInnerRulesInner := _ExecutionRulesResponseSymbolRulesInnerRulesInner{}

	err = json.Unmarshal(data, &varExecutionRulesResponseSymbolRulesInnerRulesInner)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponseSymbolRulesInnerRulesInner(varExecutionRulesResponseSymbolRulesInnerRulesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ruleType")
		delete(additionalProperties, "bidLimitMultUp")
		delete(additionalProperties, "bidLimitMultDown")
		delete(additionalProperties, "askLimitMultUp")
		delete(additionalProperties, "askLimitMultDown")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExecutionRulesResponseSymbolRulesInnerRulesInner struct {
	value *ExecutionRulesResponseSymbolRulesInnerRulesInner
	isSet bool
}

func (v NullableExecutionRulesResponseSymbolRulesInnerRulesInner) Get() *ExecutionRulesResponseSymbolRulesInnerRulesInner {
	return v.value
}

func (v *NullableExecutionRulesResponseSymbolRulesInnerRulesInner) Set(val *ExecutionRulesResponseSymbolRulesInnerRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponseSymbolRulesInnerRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponseSymbolRulesInnerRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponseSymbolRulesInnerRulesInner(val *ExecutionRulesResponseSymbolRulesInnerRulesInner) *NullableExecutionRulesResponseSymbolRulesInnerRulesInner {
	return &NullableExecutionRulesResponseSymbolRulesInnerRulesInner{value: val, isSet: true}
}

func (v NullableExecutionRulesResponseSymbolRulesInnerRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponseSymbolRulesInnerRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
