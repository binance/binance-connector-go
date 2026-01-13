/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the FuturesTradingQuantitativeRulesIndicatorsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTradingQuantitativeRulesIndicatorsResponse{}

// FuturesTradingQuantitativeRulesIndicatorsResponse struct for FuturesTradingQuantitativeRulesIndicatorsResponse
type FuturesTradingQuantitativeRulesIndicatorsResponse struct {
	Indicators           *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators `json:"indicators,omitempty"`
	UpdateTime           *int64                                                       `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTradingQuantitativeRulesIndicatorsResponse FuturesTradingQuantitativeRulesIndicatorsResponse

// NewFuturesTradingQuantitativeRulesIndicatorsResponse instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTradingQuantitativeRulesIndicatorsResponse() *FuturesTradingQuantitativeRulesIndicatorsResponse {
	this := FuturesTradingQuantitativeRulesIndicatorsResponse{}
	return &this
}

// NewFuturesTradingQuantitativeRulesIndicatorsResponseWithDefaults instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTradingQuantitativeRulesIndicatorsResponseWithDefaults() *FuturesTradingQuantitativeRulesIndicatorsResponse {
	this := FuturesTradingQuantitativeRulesIndicatorsResponse{}
	return &this
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) GetIndicators() FuturesTradingQuantitativeRulesIndicatorsResponseIndicators {
	if o == nil || common.IsNil(o.Indicators) {
		var ret FuturesTradingQuantitativeRulesIndicatorsResponseIndicators
		return ret
	}
	return *o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) GetIndicatorsOk() (*FuturesTradingQuantitativeRulesIndicatorsResponseIndicators, bool) {
	if o == nil || common.IsNil(o.Indicators) {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) HasIndicators() bool {
	if o != nil && !common.IsNil(o.Indicators) {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given FuturesTradingQuantitativeRulesIndicatorsResponseIndicators and assigns it to the Indicators field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) SetIndicators(v FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) {
	o.Indicators = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Indicators) {
		toSerialize["indicators"] = o.Indicators
	}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesTradingQuantitativeRulesIndicatorsResponse) UnmarshalJSON(data []byte) (err error) {
	varFuturesTradingQuantitativeRulesIndicatorsResponse := _FuturesTradingQuantitativeRulesIndicatorsResponse{}

	err = json.Unmarshal(data, &varFuturesTradingQuantitativeRulesIndicatorsResponse)

	if err != nil {
		return err
	}

	*o = FuturesTradingQuantitativeRulesIndicatorsResponse(varFuturesTradingQuantitativeRulesIndicatorsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "indicators")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesTradingQuantitativeRulesIndicatorsResponse struct {
	value *FuturesTradingQuantitativeRulesIndicatorsResponse
	isSet bool
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponse) Get() *FuturesTradingQuantitativeRulesIndicatorsResponse {
	return v.value
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponse) Set(val *FuturesTradingQuantitativeRulesIndicatorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTradingQuantitativeRulesIndicatorsResponse(val *FuturesTradingQuantitativeRulesIndicatorsResponse) *NullableFuturesTradingQuantitativeRulesIndicatorsResponse {
	return &NullableFuturesTradingQuantitativeRulesIndicatorsResponse{value: val, isSet: true}
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
