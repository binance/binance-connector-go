/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TradingScheduleResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TradingScheduleResponse{}

// TradingScheduleResponse struct for TradingScheduleResponse
type TradingScheduleResponse struct {
	UpdateTime           *int64                                  `json:"updateTime,omitempty"`
	MarketSchedules      *TradingScheduleResponseMarketSchedules `json:"marketSchedules,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TradingScheduleResponse TradingScheduleResponse

// NewTradingScheduleResponse instantiates a new TradingScheduleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradingScheduleResponse() *TradingScheduleResponse {
	this := TradingScheduleResponse{}
	return &this
}

// NewTradingScheduleResponseWithDefaults instantiates a new TradingScheduleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradingScheduleResponseWithDefaults() *TradingScheduleResponse {
	this := TradingScheduleResponse{}
	return &this
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *TradingScheduleResponse) GetUpdateTime() int64 {
	if o == nil || common.IsNil(o.UpdateTime) {
		var ret int64
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponse) GetUpdateTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *TradingScheduleResponse) HasUpdateTime() bool {
	if o != nil && !common.IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int64 and assigns it to the UpdateTime field.
func (o *TradingScheduleResponse) SetUpdateTime(v int64) {
	o.UpdateTime = &v
}

// GetMarketSchedules returns the MarketSchedules field value if set, zero value otherwise.
func (o *TradingScheduleResponse) GetMarketSchedules() TradingScheduleResponseMarketSchedules {
	if o == nil || common.IsNil(o.MarketSchedules) {
		var ret TradingScheduleResponseMarketSchedules
		return ret
	}
	return *o.MarketSchedules
}

// GetMarketSchedulesOk returns a tuple with the MarketSchedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradingScheduleResponse) GetMarketSchedulesOk() (*TradingScheduleResponseMarketSchedules, bool) {
	if o == nil || common.IsNil(o.MarketSchedules) {
		return nil, false
	}
	return o.MarketSchedules, true
}

// HasMarketSchedules returns a boolean if a field has been set.
func (o *TradingScheduleResponse) HasMarketSchedules() bool {
	if o != nil && !common.IsNil(o.MarketSchedules) {
		return true
	}

	return false
}

// SetMarketSchedules gets a reference to the given TradingScheduleResponseMarketSchedules and assigns it to the MarketSchedules field.
func (o *TradingScheduleResponse) SetMarketSchedules(v TradingScheduleResponseMarketSchedules) {
	o.MarketSchedules = &v
}

func (o TradingScheduleResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TradingScheduleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	if !common.IsNil(o.MarketSchedules) {
		toSerialize["marketSchedules"] = o.MarketSchedules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TradingScheduleResponse) UnmarshalJSON(data []byte) (err error) {
	varTradingScheduleResponse := _TradingScheduleResponse{}

	err = json.Unmarshal(data, &varTradingScheduleResponse)

	if err != nil {
		return err
	}

	*o = TradingScheduleResponse(varTradingScheduleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "updateTime")
		delete(additionalProperties, "marketSchedules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTradingScheduleResponse struct {
	value *TradingScheduleResponse
	isSet bool
}

func (v NullableTradingScheduleResponse) Get() *TradingScheduleResponse {
	return v.value
}

func (v *NullableTradingScheduleResponse) Set(val *TradingScheduleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTradingScheduleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTradingScheduleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradingScheduleResponse(val *TradingScheduleResponse) *NullableTradingScheduleResponse {
	return &NullableTradingScheduleResponse{value: val, isSet: true}
}

func (v NullableTradingScheduleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradingScheduleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
