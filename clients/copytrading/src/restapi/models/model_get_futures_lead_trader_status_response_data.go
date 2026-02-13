/*
Binance Copy Trading REST API

OpenAPI Specification for the Binance Copy Trading REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFuturesLeadTraderStatusResponseData type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesLeadTraderStatusResponseData{}

// GetFuturesLeadTraderStatusResponseData struct for GetFuturesLeadTraderStatusResponseData
type GetFuturesLeadTraderStatusResponseData struct {
	IsLeadTrader         *bool  `json:"isLeadTrader,omitempty"`
	Time                 *int64 `json:"time,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesLeadTraderStatusResponseData GetFuturesLeadTraderStatusResponseData

// NewGetFuturesLeadTraderStatusResponseData instantiates a new GetFuturesLeadTraderStatusResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesLeadTraderStatusResponseData() *GetFuturesLeadTraderStatusResponseData {
	this := GetFuturesLeadTraderStatusResponseData{}
	return &this
}

// NewGetFuturesLeadTraderStatusResponseDataWithDefaults instantiates a new GetFuturesLeadTraderStatusResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesLeadTraderStatusResponseDataWithDefaults() *GetFuturesLeadTraderStatusResponseData {
	this := GetFuturesLeadTraderStatusResponseData{}
	return &this
}

// GetIsLeadTrader returns the IsLeadTrader field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponseData) GetIsLeadTrader() bool {
	if o == nil || common.IsNil(o.IsLeadTrader) {
		var ret bool
		return ret
	}
	return *o.IsLeadTrader
}

// GetIsLeadTraderOk returns a tuple with the IsLeadTrader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponseData) GetIsLeadTraderOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsLeadTrader) {
		return nil, false
	}
	return o.IsLeadTrader, true
}

// HasIsLeadTrader returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponseData) HasIsLeadTrader() bool {
	if o != nil && !common.IsNil(o.IsLeadTrader) {
		return true
	}

	return false
}

// SetIsLeadTrader gets a reference to the given bool and assigns it to the IsLeadTrader field.
func (o *GetFuturesLeadTraderStatusResponseData) SetIsLeadTrader(v bool) {
	o.IsLeadTrader = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponseData) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponseData) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponseData) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *GetFuturesLeadTraderStatusResponseData) SetTime(v int64) {
	o.Time = &v
}

func (o GetFuturesLeadTraderStatusResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesLeadTraderStatusResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.IsLeadTrader) {
		toSerialize["isLeadTrader"] = o.IsLeadTrader
	}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFuturesLeadTraderStatusResponseData) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesLeadTraderStatusResponseData := _GetFuturesLeadTraderStatusResponseData{}

	err = json.Unmarshal(data, &varGetFuturesLeadTraderStatusResponseData)

	if err != nil {
		return err
	}

	*o = GetFuturesLeadTraderStatusResponseData(varGetFuturesLeadTraderStatusResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "isLeadTrader")
		delete(additionalProperties, "time")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesLeadTraderStatusResponseData struct {
	value *GetFuturesLeadTraderStatusResponseData
	isSet bool
}

func (v NullableGetFuturesLeadTraderStatusResponseData) Get() *GetFuturesLeadTraderStatusResponseData {
	return v.value
}

func (v *NullableGetFuturesLeadTraderStatusResponseData) Set(val *GetFuturesLeadTraderStatusResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesLeadTraderStatusResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesLeadTraderStatusResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesLeadTraderStatusResponseData(val *GetFuturesLeadTraderStatusResponseData) *NullableGetFuturesLeadTraderStatusResponseData {
	return &NullableGetFuturesLeadTraderStatusResponseData{value: val, isSet: true}
}

func (v NullableGetFuturesLeadTraderStatusResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesLeadTraderStatusResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
