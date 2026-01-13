/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators{}

// PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators struct for PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators
type PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators struct {
	BTCUSDT              []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner `json:"BTCUSDT,omitempty"`
	ACCOUNT              []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner `json:"ACCOUNT,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators{}
	return &this
}

// NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsWithDefaults instantiates a new PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsWithDefaults() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators {
	this := PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators{}
	return &this
}

// GetBTCUSDT returns the BTCUSDT field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) GetBTCUSDT() []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	if o == nil || common.IsNil(o.BTCUSDT) {
		var ret []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
		return ret
	}
	return o.BTCUSDT
}

// GetBTCUSDTOk returns a tuple with the BTCUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) GetBTCUSDTOk() ([]PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner, bool) {
	if o == nil || common.IsNil(o.BTCUSDT) {
		return nil, false
	}
	return o.BTCUSDT, true
}

// HasBTCUSDT returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) HasBTCUSDT() bool {
	if o != nil && !common.IsNil(o.BTCUSDT) {
		return true
	}

	return false
}

// SetBTCUSDT gets a reference to the given []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner and assigns it to the BTCUSDT field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) SetBTCUSDT(v []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) {
	o.BTCUSDT = v
}

// GetACCOUNT returns the ACCOUNT field value if set, zero value otherwise.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) GetACCOUNT() []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	if o == nil || common.IsNil(o.ACCOUNT) {
		var ret []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
		return ret
	}
	return o.ACCOUNT
}

// GetACCOUNTOk returns a tuple with the ACCOUNT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) GetACCOUNTOk() ([]PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner, bool) {
	if o == nil || common.IsNil(o.ACCOUNT) {
		return nil, false
	}
	return o.ACCOUNT, true
}

// HasACCOUNT returns a boolean if a field has been set.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) HasACCOUNT() bool {
	if o != nil && !common.IsNil(o.ACCOUNT) {
		return true
	}

	return false
}

// SetACCOUNT gets a reference to the given []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner and assigns it to the ACCOUNT field.
func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) SetACCOUNT(v []PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) {
	o.ACCOUNT = v
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BTCUSDT) {
		toSerialize["BTCUSDT"] = o.BTCUSDT
	}
	if !common.IsNil(o.ACCOUNT) {
		toSerialize["ACCOUNT"] = o.ACCOUNT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) UnmarshalJSON(data []byte) (err error) {
	varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators := _PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators{}

	err = json.Unmarshal(data, &varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators)

	if err != nil {
		return err
	}

	*o = PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators(varPortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "BTCUSDT")
		delete(additionalProperties, "ACCOUNT")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators struct {
	value *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators
	isSet bool
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) Get() *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators {
	return v.value
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) Set(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators(val *PortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators {
	return &NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators{value: val, isSet: true}
}

func (v NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioMarginUmTradingQuantitativeRulesIndicatorsResponseIndicators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
