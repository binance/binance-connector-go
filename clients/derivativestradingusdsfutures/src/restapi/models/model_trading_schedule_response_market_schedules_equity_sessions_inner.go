/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesEQUITYSessionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesEQUITYSessionsInner{}

// TradingScheduleResponseMarketSchedulesEQUITYSessionsInner struct for TradingScheduleResponseMarketSchedulesEQUITYSessionsInner
type TradingScheduleResponseMarketSchedulesEQUITYSessionsInner struct {
	StartTime            *int64  `json:"startTime,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesEQUITYSessionsInner TradingScheduleResponseMarketSchedulesEQUITYSessionsInner

// NewTradingScheduleResponseMarketSchedulesEQUITYSessionsInner instantiates a new TradingScheduleResponseMarketSchedulesEQUITYSessionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesEQUITYSessionsInner() *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner {
	this := TradingScheduleResponseMarketSchedulesEQUITYSessionsInner{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesEQUITYSessionsInnerWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesEQUITYSessionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesEQUITYSessionsInnerWithDefaults() *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner {
	this := TradingScheduleResponseMarketSchedulesEQUITYSessionsInner{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetStartTime() int64 {
	if o == nil || common.IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) SetType(v string) {
	o.Type = &v
}

func (o TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !common.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesEQUITYSessionsInner := _TradingScheduleResponseMarketSchedulesEQUITYSessionsInner{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesEQUITYSessionsInner)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesEQUITYSessionsInner(varTradingScheduleResponseMarketSchedulesEQUITYSessionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner struct {
	value *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) Get() *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) Set(val *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner(val *TradingScheduleResponseMarketSchedulesEQUITYSessionsInner) *NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner {
	return &NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesEQUITYSessionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
