/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExecutionRulesResponseResultSymbolRulesInnerRulesInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExecutionRulesResponseResultSymbolRulesInnerRulesInner{}

// ExecutionRulesResponseResultSymbolRulesInnerRulesInner struct for ExecutionRulesResponseResultSymbolRulesInnerRulesInner
type ExecutionRulesResponseResultSymbolRulesInnerRulesInner struct {
	RuleType             *string `json:"ruleType,omitempty"`
	BidLimitMultUp       *string `json:"bidLimitMultUp,omitempty"`
	BidLimitMultDown     *string `json:"bidLimitMultDown,omitempty"`
	AskLimitMultUp       *string `json:"askLimitMultUp,omitempty"`
	AskLimitMultDown     *string `json:"askLimitMultDown,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExecutionRulesResponseResultSymbolRulesInnerRulesInner ExecutionRulesResponseResultSymbolRulesInnerRulesInner

// NewExecutionRulesResponseResultSymbolRulesInnerRulesInner instantiates a new ExecutionRulesResponseResultSymbolRulesInnerRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionRulesResponseResultSymbolRulesInnerRulesInner() *ExecutionRulesResponseResultSymbolRulesInnerRulesInner {
	this := ExecutionRulesResponseResultSymbolRulesInnerRulesInner{}
	return &this
}

// NewExecutionRulesResponseResultSymbolRulesInnerRulesInnerWithDefaults instantiates a new ExecutionRulesResponseResultSymbolRulesInnerRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionRulesResponseResultSymbolRulesInnerRulesInnerWithDefaults() *ExecutionRulesResponseResultSymbolRulesInnerRulesInner {
	this := ExecutionRulesResponseResultSymbolRulesInnerRulesInner{}
	return &this
}

// GetRuleType returns the RuleType field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetRuleType() string {
	if o == nil || common.IsNil(o.RuleType) {
		var ret string
		return ret
	}
	return *o.RuleType
}

// GetRuleTypeOk returns a tuple with the RuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetRuleTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.RuleType) {
		return nil, false
	}
	return o.RuleType, true
}

// HasRuleType returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) HasRuleType() bool {
	if o != nil && !common.IsNil(o.RuleType) {
		return true
	}

	return false
}

// SetRuleType gets a reference to the given string and assigns it to the RuleType field.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) SetRuleType(v string) {
	o.RuleType = &v
}

// GetBidLimitMultUp returns the BidLimitMultUp field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetBidLimitMultUp() string {
	if o == nil || common.IsNil(o.BidLimitMultUp) {
		var ret string
		return ret
	}
	return *o.BidLimitMultUp
}

// GetBidLimitMultUpOk returns a tuple with the BidLimitMultUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetBidLimitMultUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidLimitMultUp) {
		return nil, false
	}
	return o.BidLimitMultUp, true
}

// HasBidLimitMultUp returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) HasBidLimitMultUp() bool {
	if o != nil && !common.IsNil(o.BidLimitMultUp) {
		return true
	}

	return false
}

// SetBidLimitMultUp gets a reference to the given string and assigns it to the BidLimitMultUp field.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) SetBidLimitMultUp(v string) {
	o.BidLimitMultUp = &v
}

// GetBidLimitMultDown returns the BidLimitMultDown field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetBidLimitMultDown() string {
	if o == nil || common.IsNil(o.BidLimitMultDown) {
		var ret string
		return ret
	}
	return *o.BidLimitMultDown
}

// GetBidLimitMultDownOk returns a tuple with the BidLimitMultDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetBidLimitMultDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.BidLimitMultDown) {
		return nil, false
	}
	return o.BidLimitMultDown, true
}

// HasBidLimitMultDown returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) HasBidLimitMultDown() bool {
	if o != nil && !common.IsNil(o.BidLimitMultDown) {
		return true
	}

	return false
}

// SetBidLimitMultDown gets a reference to the given string and assigns it to the BidLimitMultDown field.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) SetBidLimitMultDown(v string) {
	o.BidLimitMultDown = &v
}

// GetAskLimitMultUp returns the AskLimitMultUp field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetAskLimitMultUp() string {
	if o == nil || common.IsNil(o.AskLimitMultUp) {
		var ret string
		return ret
	}
	return *o.AskLimitMultUp
}

// GetAskLimitMultUpOk returns a tuple with the AskLimitMultUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetAskLimitMultUpOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskLimitMultUp) {
		return nil, false
	}
	return o.AskLimitMultUp, true
}

// HasAskLimitMultUp returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) HasAskLimitMultUp() bool {
	if o != nil && !common.IsNil(o.AskLimitMultUp) {
		return true
	}

	return false
}

// SetAskLimitMultUp gets a reference to the given string and assigns it to the AskLimitMultUp field.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) SetAskLimitMultUp(v string) {
	o.AskLimitMultUp = &v
}

// GetAskLimitMultDown returns the AskLimitMultDown field value if set, zero value otherwise.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetAskLimitMultDown() string {
	if o == nil || common.IsNil(o.AskLimitMultDown) {
		var ret string
		return ret
	}
	return *o.AskLimitMultDown
}

// GetAskLimitMultDownOk returns a tuple with the AskLimitMultDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) GetAskLimitMultDownOk() (*string, bool) {
	if o == nil || common.IsNil(o.AskLimitMultDown) {
		return nil, false
	}
	return o.AskLimitMultDown, true
}

// HasAskLimitMultDown returns a boolean if a field has been set.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) HasAskLimitMultDown() bool {
	if o != nil && !common.IsNil(o.AskLimitMultDown) {
		return true
	}

	return false
}

// SetAskLimitMultDown gets a reference to the given string and assigns it to the AskLimitMultDown field.
func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) SetAskLimitMultDown(v string) {
	o.AskLimitMultDown = &v
}

func (o ExecutionRulesResponseResultSymbolRulesInnerRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionRulesResponseResultSymbolRulesInnerRulesInner) ToMap() (map[string]interface{}, error) {
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

func (o *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) UnmarshalJSON(data []byte) (err error) {
	varExecutionRulesResponseResultSymbolRulesInnerRulesInner := _ExecutionRulesResponseResultSymbolRulesInnerRulesInner{}

	err = json.Unmarshal(data, &varExecutionRulesResponseResultSymbolRulesInnerRulesInner)

	if err != nil {
		return err
	}

	*o = ExecutionRulesResponseResultSymbolRulesInnerRulesInner(varExecutionRulesResponseResultSymbolRulesInnerRulesInner)

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

type NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner struct {
	value *ExecutionRulesResponseResultSymbolRulesInnerRulesInner
	isSet bool
}

func (v NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) Get() *ExecutionRulesResponseResultSymbolRulesInnerRulesInner {
	return v.value
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) Set(val *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionRulesResponseResultSymbolRulesInnerRulesInner(val *ExecutionRulesResponseResultSymbolRulesInnerRulesInner) *NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner {
	return &NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner{value: val, isSet: true}
}

func (v NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionRulesResponseResultSymbolRulesInnerRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
