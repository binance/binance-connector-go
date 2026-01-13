/*
Binance Derivatives Trading Portfolio Margin REST API

OpenAPI Specification for the Binance Derivatives Trading Portfolio Margin REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner{}

// QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner struct for QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner
type QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner struct {
	StartTime            *int64                                                                    `json:"startTime,omitempty"`
	EndTime              *int64                                                                    `json:"endTime,omitempty"`
	Details              []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner

// NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner{}
	return &this
}

// NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerWithDefaults instantiates a new QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerWithDefaults() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner {
	this := QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetStartTime() int64 {
	if o == nil || common.IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) HasStartTime() bool {
	if o != nil && !common.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetEndTime() int64 {
	if o == nil || common.IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetEndTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) HasEndTime() bool {
	if o != nil && !common.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetDetails() []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner {
	if o == nil || common.IsNil(o.Details) {
		var ret []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) GetDetailsOk() ([]QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner, bool) {
	if o == nil || common.IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) HasDetails() bool {
	if o != nil && !common.IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner and assigns it to the Details field.
func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) SetDetails(v []QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInnerDetailsInner) {
	o.Details = v
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !common.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !common.IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner := _QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner{}

	err = json.Unmarshal(data, &varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner)

	if err != nil {
		return err
	}

	*o = QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner(varQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "startTime")
		delete(additionalProperties, "endTime")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner struct {
	value *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner
	isSet bool
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) Get() *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner {
	return v.value
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) Set(val *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner(val *QueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner {
	return &NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner{value: val, isSet: true}
}

func (v NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUserNegativeBalanceAutoExchangeRecordResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
