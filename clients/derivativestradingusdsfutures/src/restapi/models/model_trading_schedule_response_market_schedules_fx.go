/*
Futures (USDⓈ-M) REST API

Access market data, manage accounts, and trade USDⓈ-M perpetual futures.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesFX type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesFX{}

// TradingScheduleResponseMarketSchedulesFX struct for TradingScheduleResponseMarketSchedulesFX
type TradingScheduleResponseMarketSchedulesFX struct {
	Sessions             []TradingScheduleResponseMarketSchedulesFXSessionsInner `json:"sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesFX TradingScheduleResponseMarketSchedulesFX

// NewTradingScheduleResponseMarketSchedulesFX instantiates a new TradingScheduleResponseMarketSchedulesFX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesFX() *TradingScheduleResponseMarketSchedulesFX {
	this := TradingScheduleResponseMarketSchedulesFX{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesFXWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesFX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesFXWithDefaults() *TradingScheduleResponseMarketSchedulesFX {
	this := TradingScheduleResponseMarketSchedulesFX{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesFX) GetSessions() []TradingScheduleResponseMarketSchedulesFXSessionsInner {
	if o == nil || common.IsNil(o.Sessions) {
		var ret []TradingScheduleResponseMarketSchedulesFXSessionsInner
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesFX) GetSessionsOk() ([]TradingScheduleResponseMarketSchedulesFXSessionsInner, bool) {
	if o == nil || common.IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesFX) HasSessions() bool {
	if o != nil && !common.IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []TradingScheduleResponseMarketSchedulesFXSessionsInner and assigns it to the Sessions field.
func (o *TradingScheduleResponseMarketSchedulesFX) SetSessions(v []TradingScheduleResponseMarketSchedulesFXSessionsInner) {
	o.Sessions = v
}

func (o TradingScheduleResponseMarketSchedulesFX) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesFX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesFX) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesFX := _TradingScheduleResponseMarketSchedulesFX{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesFX)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesFX(varTradingScheduleResponseMarketSchedulesFX)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesFX struct {
	value *TradingScheduleResponseMarketSchedulesFX
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesFX) Get() *TradingScheduleResponseMarketSchedulesFX {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesFX) Set(val *TradingScheduleResponseMarketSchedulesFX) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesFX) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesFX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesFX(val *TradingScheduleResponseMarketSchedulesFX) *NullableTradingScheduleResponseMarketSchedulesFX {
	return &NullableTradingScheduleResponseMarketSchedulesFX{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesFX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesFX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
