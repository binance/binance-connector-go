/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesKREQUITY type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesKREQUITY{}

// TradingScheduleResponseMarketSchedulesKREQUITY struct for TradingScheduleResponseMarketSchedulesKREQUITY
type TradingScheduleResponseMarketSchedulesKREQUITY struct {
	Sessions             []TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner `json:"sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesKREQUITY TradingScheduleResponseMarketSchedulesKREQUITY

// NewTradingScheduleResponseMarketSchedulesKREQUITY instantiates a new TradingScheduleResponseMarketSchedulesKREQUITY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesKREQUITY() *TradingScheduleResponseMarketSchedulesKREQUITY {
	this := TradingScheduleResponseMarketSchedulesKREQUITY{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesKREQUITYWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesKREQUITY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesKREQUITYWithDefaults() *TradingScheduleResponseMarketSchedulesKREQUITY {
	this := TradingScheduleResponseMarketSchedulesKREQUITY{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesKREQUITY) GetSessions() []TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner {
	if o == nil || common.IsNil(o.Sessions) {
		var ret []TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesKREQUITY) GetSessionsOk() ([]TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner, bool) {
	if o == nil || common.IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesKREQUITY) HasSessions() bool {
	if o != nil && !common.IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner and assigns it to the Sessions field.
func (o *TradingScheduleResponseMarketSchedulesKREQUITY) SetSessions(v []TradingScheduleResponseMarketSchedulesKREQUITYSessionsInner) {
	o.Sessions = v
}

func (o TradingScheduleResponseMarketSchedulesKREQUITY) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesKREQUITY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesKREQUITY) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesKREQUITY := _TradingScheduleResponseMarketSchedulesKREQUITY{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesKREQUITY)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesKREQUITY(varTradingScheduleResponseMarketSchedulesKREQUITY)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesKREQUITY struct {
	value *TradingScheduleResponseMarketSchedulesKREQUITY
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesKREQUITY) Get() *TradingScheduleResponseMarketSchedulesKREQUITY {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesKREQUITY) Set(val *TradingScheduleResponseMarketSchedulesKREQUITY) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesKREQUITY) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesKREQUITY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesKREQUITY(val *TradingScheduleResponseMarketSchedulesKREQUITY) *NullableTradingScheduleResponseMarketSchedulesKREQUITY {
	return &NullableTradingScheduleResponseMarketSchedulesKREQUITY{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesKREQUITY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesKREQUITY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
