/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesEQUITY type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesEQUITY{}

// TradingScheduleResponseMarketSchedulesEQUITY struct for TradingScheduleResponseMarketSchedulesEQUITY
type TradingScheduleResponseMarketSchedulesEQUITY struct {
	Sessions             []TradingScheduleResponseMarketSchedulesEQUITYSessionsInner `json:"sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesEQUITY TradingScheduleResponseMarketSchedulesEQUITY

// NewTradingScheduleResponseMarketSchedulesEQUITY instantiates a new TradingScheduleResponseMarketSchedulesEQUITY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesEQUITY() *TradingScheduleResponseMarketSchedulesEQUITY {
	this := TradingScheduleResponseMarketSchedulesEQUITY{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesEQUITYWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesEQUITY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesEQUITYWithDefaults() *TradingScheduleResponseMarketSchedulesEQUITY {
	this := TradingScheduleResponseMarketSchedulesEQUITY{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesEQUITY) GetSessions() []TradingScheduleResponseMarketSchedulesEQUITYSessionsInner {
	if o == nil || common.IsNil(o.Sessions) {
		var ret []TradingScheduleResponseMarketSchedulesEQUITYSessionsInner
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITY) GetSessionsOk() ([]TradingScheduleResponseMarketSchedulesEQUITYSessionsInner, bool) {
	if o == nil || common.IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITY) HasSessions() bool {
	if o != nil && !common.IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []TradingScheduleResponseMarketSchedulesEQUITYSessionsInner and assigns it to the Sessions field.
func (o *TradingScheduleResponseMarketSchedulesEQUITY) SetSessions(v []TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) {
	o.Sessions = v
}

func (o TradingScheduleResponseMarketSchedulesEQUITY) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesEQUITY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesEQUITY) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesEQUITY := _TradingScheduleResponseMarketSchedulesEQUITY{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesEQUITY)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesEQUITY(varTradingScheduleResponseMarketSchedulesEQUITY)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesEQUITY struct {
	value *TradingScheduleResponseMarketSchedulesEQUITY
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITY) Get() *TradingScheduleResponseMarketSchedulesEQUITY {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITY) Set(val *TradingScheduleResponseMarketSchedulesEQUITY) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITY) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesEQUITY(val *TradingScheduleResponseMarketSchedulesEQUITY) *NullableTradingScheduleResponseMarketSchedulesEQUITY {
	return &NullableTradingScheduleResponseMarketSchedulesEQUITY{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
