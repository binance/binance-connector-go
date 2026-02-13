/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedules type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedules{}

// TradingScheduleResponseMarketSchedules struct for TradingScheduleResponseMarketSchedules
type TradingScheduleResponseMarketSchedules struct {
	EQUITY               *TradingScheduleResponseMarketSchedulesEQUITY    `json:"EQUITY,omitempty"`
	COMMODITY            *TradingScheduleResponseMarketSchedulesCOMMODITY `json:"COMMODITY,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedules TradingScheduleResponseMarketSchedules

// NewTradingScheduleResponseMarketSchedules instantiates a new TradingScheduleResponseMarketSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedules() *TradingScheduleResponseMarketSchedules {
	this := TradingScheduleResponseMarketSchedules{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesWithDefaults instantiates a new TradingScheduleResponseMarketSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesWithDefaults() *TradingScheduleResponseMarketSchedules {
	this := TradingScheduleResponseMarketSchedules{}
	return &this
}

// GetEQUITY returns the EQUITY field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedules) GetEQUITY() TradingScheduleResponseMarketSchedulesEQUITY {
	if o == nil || common.IsNil(o.EQUITY) {
		var ret TradingScheduleResponseMarketSchedulesEQUITY
		return ret
	}
	return *o.EQUITY
}

// GetEQUITYOk returns a tuple with the EQUITY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedules) GetEQUITYOk() (*TradingScheduleResponseMarketSchedulesEQUITY, bool) {
	if o == nil || common.IsNil(o.EQUITY) {
		return nil, false
	}
	return o.EQUITY, true
}

// HasEQUITY returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedules) HasEQUITY() bool {
	if o != nil && !common.IsNil(o.EQUITY) {
		return true
	}

	return false
}

// SetEQUITY gets a reference to the given TradingScheduleResponseMarketSchedulesEQUITY and assigns it to the EQUITY field.
func (o *TradingScheduleResponseMarketSchedules) SetEQUITY(v TradingScheduleResponseMarketSchedulesEQUITY) {
	o.EQUITY = &v
}

// GetCOMMODITY returns the COMMODITY field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedules) GetCOMMODITY() TradingScheduleResponseMarketSchedulesCOMMODITY {
	if o == nil || common.IsNil(o.COMMODITY) {
		var ret TradingScheduleResponseMarketSchedulesCOMMODITY
		return ret
	}
	return *o.COMMODITY
}

// GetCOMMODITYOk returns a tuple with the COMMODITY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedules) GetCOMMODITYOk() (*TradingScheduleResponseMarketSchedulesCOMMODITY, bool) {
	if o == nil || common.IsNil(o.COMMODITY) {
		return nil, false
	}
	return o.COMMODITY, true
}

// HasCOMMODITY returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedules) HasCOMMODITY() bool {
	if o != nil && !common.IsNil(o.COMMODITY) {
		return true
	}

	return false
}

// SetCOMMODITY gets a reference to the given TradingScheduleResponseMarketSchedulesCOMMODITY and assigns it to the COMMODITY field.
func (o *TradingScheduleResponseMarketSchedules) SetCOMMODITY(v TradingScheduleResponseMarketSchedulesCOMMODITY) {
	o.COMMODITY = &v
}

func (o TradingScheduleResponseMarketSchedules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.EQUITY) {
		toSerialize["EQUITY"] = o.EQUITY
	}
	if !common.IsNil(o.COMMODITY) {
		toSerialize["COMMODITY"] = o.COMMODITY
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedules) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedules := _TradingScheduleResponseMarketSchedules{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedules)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedules(varTradingScheduleResponseMarketSchedules)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "EQUITY")
		delete(additionalProperties, "COMMODITY")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedules struct {
	value *TradingScheduleResponseMarketSchedules
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedules) Get() *TradingScheduleResponseMarketSchedules {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedules) Set(val *TradingScheduleResponseMarketSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedules(val *TradingScheduleResponseMarketSchedules) *NullableTradingScheduleResponseMarketSchedules {
	return &NullableTradingScheduleResponseMarketSchedules{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
