/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TradingScheduleResponseMarketSchedulesCOMMODITY type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesCOMMODITY{}

// TradingScheduleResponseMarketSchedulesCOMMODITY struct for TradingScheduleResponseMarketSchedulesCOMMODITY
type TradingScheduleResponseMarketSchedulesCOMMODITY struct {
	Sessions             []TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner `json:"sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesCOMMODITY TradingScheduleResponseMarketSchedulesCOMMODITY

// NewTradingScheduleResponseMarketSchedulesCOMMODITY instantiates a new TradingScheduleResponseMarketSchedulesCOMMODITY object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesCOMMODITY() *TradingScheduleResponseMarketSchedulesCOMMODITY {
	this := TradingScheduleResponseMarketSchedulesCOMMODITY{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesCOMMODITYWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesCOMMODITY object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesCOMMODITYWithDefaults() *TradingScheduleResponseMarketSchedulesCOMMODITY {
	this := TradingScheduleResponseMarketSchedulesCOMMODITY{}
	return &this
}

// GetSessions returns the Sessions field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITY) GetSessions() []TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner {
	if o == nil || common.IsNil(o.Sessions) {
		var ret []TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITY) GetSessionsOk() ([]TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner, bool) {
	if o == nil || common.IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITY) HasSessions() bool {
	if o != nil && !common.IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner and assigns it to the Sessions field.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITY) SetSessions(v []TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) {
	o.Sessions = v
}

func (o TradingScheduleResponseMarketSchedulesCOMMODITY) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesCOMMODITY) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Sessions) {
		toSerialize["sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesCOMMODITY) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesCOMMODITY := _TradingScheduleResponseMarketSchedulesCOMMODITY{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesCOMMODITY)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesCOMMODITY(varTradingScheduleResponseMarketSchedulesCOMMODITY)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "sessions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesCOMMODITY struct {
	value *TradingScheduleResponseMarketSchedulesCOMMODITY
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITY) Get() *TradingScheduleResponseMarketSchedulesCOMMODITY {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITY) Set(val *TradingScheduleResponseMarketSchedulesCOMMODITY) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITY) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITY) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesCOMMODITY(val *TradingScheduleResponseMarketSchedulesCOMMODITY) *NullableTradingScheduleResponseMarketSchedulesCOMMODITY {
	return &NullableTradingScheduleResponseMarketSchedulesCOMMODITY{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITY) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITY) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
