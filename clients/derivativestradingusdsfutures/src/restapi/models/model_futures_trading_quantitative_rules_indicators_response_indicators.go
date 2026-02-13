/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the FuturesTradingQuantitativeRulesIndicatorsResponseIndicators type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &FuturesTradingQuantitativeRulesIndicatorsResponseIndicators{}

// FuturesTradingQuantitativeRulesIndicatorsResponseIndicators struct for FuturesTradingQuantitativeRulesIndicatorsResponseIndicators
type FuturesTradingQuantitativeRulesIndicatorsResponseIndicators struct {
	BTCUSDT              []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner `json:"BTCUSDT,omitempty"`
	ETHUSDT              []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner `json:"ETHUSDT,omitempty"`
	ACCOUNT              []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner `json:"ACCOUNT,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FuturesTradingQuantitativeRulesIndicatorsResponseIndicators FuturesTradingQuantitativeRulesIndicatorsResponseIndicators

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicators instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicators() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicators{}
	return &this
}

// NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsWithDefaults instantiates a new FuturesTradingQuantitativeRulesIndicatorsResponseIndicators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsWithDefaults() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators {
	this := FuturesTradingQuantitativeRulesIndicatorsResponseIndicators{}
	return &this
}

// GetBTCUSDT returns the BTCUSDT field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetBTCUSDT() []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	if o == nil || common.IsNil(o.BTCUSDT) {
		var ret []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
		return ret
	}
	return o.BTCUSDT
}

// GetBTCUSDTOk returns a tuple with the BTCUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetBTCUSDTOk() ([]FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner, bool) {
	if o == nil || common.IsNil(o.BTCUSDT) {
		return nil, false
	}
	return o.BTCUSDT, true
}

// HasBTCUSDT returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) HasBTCUSDT() bool {
	if o != nil && !common.IsNil(o.BTCUSDT) {
		return true
	}

	return false
}

// SetBTCUSDT gets a reference to the given []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner and assigns it to the BTCUSDT field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) SetBTCUSDT(v []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) {
	o.BTCUSDT = v
}

// GetETHUSDT returns the ETHUSDT field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetETHUSDT() []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner {
	if o == nil || common.IsNil(o.ETHUSDT) {
		var ret []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner
		return ret
	}
	return o.ETHUSDT
}

// GetETHUSDTOk returns a tuple with the ETHUSDT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetETHUSDTOk() ([]FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner, bool) {
	if o == nil || common.IsNil(o.ETHUSDT) {
		return nil, false
	}
	return o.ETHUSDT, true
}

// HasETHUSDT returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) HasETHUSDT() bool {
	if o != nil && !common.IsNil(o.ETHUSDT) {
		return true
	}

	return false
}

// SetETHUSDT gets a reference to the given []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner and assigns it to the ETHUSDT field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) SetETHUSDT(v []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsBTCUSDTInner) {
	o.ETHUSDT = v
}

// GetACCOUNT returns the ACCOUNT field value if set, zero value otherwise.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetACCOUNT() []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner {
	if o == nil || common.IsNil(o.ACCOUNT) {
		var ret []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner
		return ret
	}
	return o.ACCOUNT
}

// GetACCOUNTOk returns a tuple with the ACCOUNT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) GetACCOUNTOk() ([]FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner, bool) {
	if o == nil || common.IsNil(o.ACCOUNT) {
		return nil, false
	}
	return o.ACCOUNT, true
}

// HasACCOUNT returns a boolean if a field has been set.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) HasACCOUNT() bool {
	if o != nil && !common.IsNil(o.ACCOUNT) {
		return true
	}

	return false
}

// SetACCOUNT gets a reference to the given []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner and assigns it to the ACCOUNT field.
func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) SetACCOUNT(v []FuturesTradingQuantitativeRulesIndicatorsResponseIndicatorsACCOUNTInner) {
	o.ACCOUNT = v
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.BTCUSDT) {
		toSerialize["BTCUSDT"] = o.BTCUSDT
	}
	if !common.IsNil(o.ETHUSDT) {
		toSerialize["ETHUSDT"] = o.ETHUSDT
	}
	if !common.IsNil(o.ACCOUNT) {
		toSerialize["ACCOUNT"] = o.ACCOUNT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) UnmarshalJSON(data []byte) (err error) {
	varFuturesTradingQuantitativeRulesIndicatorsResponseIndicators := _FuturesTradingQuantitativeRulesIndicatorsResponseIndicators{}

	err = json.Unmarshal(data, &varFuturesTradingQuantitativeRulesIndicatorsResponseIndicators)

	if err != nil {
		return err
	}

	*o = FuturesTradingQuantitativeRulesIndicatorsResponseIndicators(varFuturesTradingQuantitativeRulesIndicatorsResponseIndicators)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "BTCUSDT")
		delete(additionalProperties, "ETHUSDT")
		delete(additionalProperties, "ACCOUNT")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators struct {
	value *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators
	isSet bool
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) Get() *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators {
	return v.value
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) Set(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) {
	v.value = val
	v.isSet = true
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) IsSet() bool {
	return v.isSet
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators(val *FuturesTradingQuantitativeRulesIndicatorsResponseIndicators) *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators {
	return &NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators{value: val, isSet: true}
}

func (v NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuturesTradingQuantitativeRulesIndicatorsResponseIndicators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
