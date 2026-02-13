/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner{}

// TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner struct for TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner
type TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner struct {
	StartTime            *int64  `json:"startTime,omitempty"`
	EndTime              *int64  `json:"endTime,omitempty"`
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner

// NewTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner instantiates a new TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner() *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner {
	this := TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner{}
	return &this
}

// NewTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInnerWithDefaults instantiates a new TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInnerWithDefaults() *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner {
	this := TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetStartTime() int64 {
	if o == nil || common.IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetType() string {
	if o == nil || common.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) GetTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) HasType() bool {
	if o != nil && !common.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) SetType(v string) {
	o.Type = &v
}

func (o TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) ToMap() (map[string]interface{}, error) {
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

func (o *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner := _TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner{}

	err = json.Unmarshal(data, &varTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner(varTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner struct {
	value *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner
	isSet bool
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) Get() *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner {
	return v.value
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) Set(val *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner(val *TradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) *NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner {
	return &NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner{value: val, isSet: true}
}

func (v NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponseMarketSchedulesCOMMODITYSessionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
