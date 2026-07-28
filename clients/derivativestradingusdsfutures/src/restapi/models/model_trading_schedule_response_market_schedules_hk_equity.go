/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesHKEQUITY type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesHKEQUITY{}

// TradingScheduleResponseMarketSchedulesHKEQUITY struct for TradingScheduleResponseMarketSchedulesHKEQUITY
type TradingScheduleResponseMarketSchedulesHKEQUITY struct {
	Sessions             []TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner `json:"sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesHKEQUITY TradingScheduleResponseMarketSchedulesHKEQUITY

// NewTradingScheduleResponseMarketSchedulesHKEQUITY instantiates a new TradingScheduleResponseMarketSchedulesHKEQUITY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesHKEQUITY() *TradingScheduleResponseMarketSchedulesHKEQUITY {
	this := TradingScheduleResponseMarketSchedulesHKEQUITY{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesHKEQUITYWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesHKEQUITY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesHKEQUITYWithDefaults() *TradingScheduleResponseMarketSchedulesHKEQUITY {
	this := TradingScheduleResponseMarketSchedulesHKEQUITY{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesHKEQUITY) GetSessions() []TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner {
	if o == nil || common.IsNil(o.Sessions) {
		var ret []TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesHKEQUITY) GetSessionsOk() ([]TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner, bool) {
	if o == nil || common.IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesHKEQUITY) HasSessions() bool {
	if o != nil && !common.IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner and assigns it to the Sessions field.
func (o *TradingScheduleResponseMarketSchedulesHKEQUITY) SetSessions(v []TradingScheduleResponseMarketSchedulesHKEQUITYSessionsInner) {
	o.Sessions = v
}

func (o TradingScheduleResponseMarketSchedulesHKEQUITY) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesHKEQUITY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesHKEQUITY) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesHKEQUITY := _TradingScheduleResponseMarketSchedulesHKEQUITY{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesHKEQUITY)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesHKEQUITY(varTradingScheduleResponseMarketSchedulesHKEQUITY)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesHKEQUITY struct {
	value *TradingScheduleResponseMarketSchedulesHKEQUITY
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesHKEQUITY) Get() *TradingScheduleResponseMarketSchedulesHKEQUITY {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesHKEQUITY) Set(val *TradingScheduleResponseMarketSchedulesHKEQUITY) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesHKEQUITY) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesHKEQUITY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesHKEQUITY(val *TradingScheduleResponseMarketSchedulesHKEQUITY) *NullableTradingScheduleResponseMarketSchedulesHKEQUITY {
	return &NullableTradingScheduleResponseMarketSchedulesHKEQUITY{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesHKEQUITY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesHKEQUITY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
