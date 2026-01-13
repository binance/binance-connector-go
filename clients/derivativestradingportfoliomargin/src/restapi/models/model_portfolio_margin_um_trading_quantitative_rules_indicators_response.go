/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse{}

// PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse struct for PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse
type PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse struct {
	Indicators           *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators `json:"indicators,omitempty"`
	UpdateTime           *int64                                                                 `json:"updateTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse{}
	return &this
}

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseWithDefaults instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseWithDefaults() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse{}
	return &this
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) GetIndicators() PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators {
	if o == nil || common.IsNil(o.Indicators) {
		var ret PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators
		return ret
	}
	return *o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) GetIndicatorsOk() (*PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators, bool) {
	if o == nil || common.IsNil(o.Indicators) {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) HasIndicators() bool {
	if o != nil && !common.IsNil(o.Indicators) {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators and assigns it to the Indicators field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) SetIndicators(v PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) {
	o.Indicators = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse := _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse{}

	err = json.Unmarshal(data, &varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse)

	if err != nil {
		return err
	}

	*o = PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse(varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "indicators")
		delete(additionalProperties, "updateTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse struct {
	value *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse
	isSet bool
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) Get() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse {
	return v.value
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) Set(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse {
	return &NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse{value: val, isSet: true}
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
