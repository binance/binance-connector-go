/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RequestForMinerListResponseDataWorkerDatasInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestForMinerListResponseDataWorkerDatasInner{}

// RequestForMinerListResponseDataWorkerDatasInner struct for RequestForMinerListResponseDataWorkerDatasInner
type RequestForMinerListResponseDataWorkerDatasInner struct {
	WorkerId             *string  `json:"workerId,omitempty"`
	WorkerName           *string  `json:"workerName,omitempty"`
	Status               *int64   `json:"status,omitempty"`
	HashRate             *int64   `json:"hashRate,omitempty"`
	DayHashRate          *float32 `json:"dayHashRate,omitempty"`
	RejectRate           *int64   `json:"rejectRate,omitempty"`
	LastShareTime        *int64   `json:"lastShareTime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestForMinerListResponseDataWorkerDatasInner RequestForMinerListResponseDataWorkerDatasInner

// NewRequestForMinerListResponseDataWorkerDatasInner instantiates a new RequestForMinerListResponseDataWorkerDatasInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestForMinerListResponseDataWorkerDatasInner() *RequestForMinerListResponseDataWorkerDatasInner {
	this := RequestForMinerListResponseDataWorkerDatasInner{}
	return &this
}

// NewRequestForMinerListResponseDataWorkerDatasInnerWithDefaults instantiates a new RequestForMinerListResponseDataWorkerDatasInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestForMinerListResponseDataWorkerDatasInnerWithDefaults() *RequestForMinerListResponseDataWorkerDatasInner {
	this := RequestForMinerListResponseDataWorkerDatasInner{}
	return &this
}

// GetWorkerId returns the WorkerId field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetWorkerId() string {
	if o == nil || common.IsNil(o.WorkerId) {
		var ret string
		return ret
	}
	return *o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetWorkerIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkerId) {
		return nil, false
	}
	return o.WorkerId, true
}

// HasWorkerId returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasWorkerId() bool {
	if o != nil && !common.IsNil(o.WorkerId) {
		return true
	}

	return false
}

// SetWorkerId gets a reference to the given string and assigns it to the WorkerId field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetWorkerId(v string) {
	o.WorkerId = &v
}

// GetWorkerName returns the WorkerName field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetWorkerName() string {
	if o == nil || common.IsNil(o.WorkerName) {
		var ret string
		return ret
	}
	return *o.WorkerName
}

// GetWorkerNameOk returns a tuple with the WorkerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetWorkerNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.WorkerName) {
		return nil, false
	}
	return o.WorkerName, true
}

// HasWorkerName returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasWorkerName() bool {
	if o != nil && !common.IsNil(o.WorkerName) {
		return true
	}

	return false
}

// SetWorkerName gets a reference to the given string and assigns it to the WorkerName field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetWorkerName(v string) {
	o.WorkerName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetStatus(v int64) {
	o.Status = &v
}

// GetHashRate returns the HashRate field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetHashRate() int64 {
	if o == nil || common.IsNil(o.HashRate) {
		var ret int64
		return ret
	}
	return *o.HashRate
}

// GetHashRateOk returns a tuple with the HashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetHashRateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.HashRate) {
		return nil, false
	}
	return o.HashRate, true
}

// HasHashRate returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasHashRate() bool {
	if o != nil && !common.IsNil(o.HashRate) {
		return true
	}

	return false
}

// SetHashRate gets a reference to the given int64 and assigns it to the HashRate field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetHashRate(v int64) {
	o.HashRate = &v
}

// GetDayHashRate returns the DayHashRate field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetDayHashRate() float32 {
	if o == nil || common.IsNil(o.DayHashRate) {
		var ret float32
		return ret
	}
	return *o.DayHashRate
}

// GetDayHashRateOk returns a tuple with the DayHashRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetDayHashRateOk() (*float32, bool) {
	if o == nil || common.IsNil(o.DayHashRate) {
		return nil, false
	}
	return o.DayHashRate, true
}

// HasDayHashRate returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasDayHashRate() bool {
	if o != nil && !common.IsNil(o.DayHashRate) {
		return true
	}

	return false
}

// SetDayHashRate gets a reference to the given float32 and assigns it to the DayHashRate field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetDayHashRate(v float32) {
	o.DayHashRate = &v
}

// GetRejectRate returns the RejectRate field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetRejectRate() int64 {
	if o == nil || common.IsNil(o.RejectRate) {
		var ret int64
		return ret
	}
	return *o.RejectRate
}

// GetRejectRateOk returns a tuple with the RejectRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetRejectRateOk() (*int64, bool) {
	if o == nil || common.IsNil(o.RejectRate) {
		return nil, false
	}
	return o.RejectRate, true
}

// HasRejectRate returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasRejectRate() bool {
	if o != nil && !common.IsNil(o.RejectRate) {
		return true
	}

	return false
}

// SetRejectRate gets a reference to the given int64 and assigns it to the RejectRate field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetRejectRate(v int64) {
	o.RejectRate = &v
}

// GetLastShareTime returns the LastShareTime field value if set, zero value otherwise.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetLastShareTime() int64 {
	if o == nil || common.IsNil(o.LastShareTime) {
		var ret int64
		return ret
	}
	return *o.LastShareTime
}

// GetLastShareTimeOk returns a tuple with the LastShareTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) GetLastShareTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.LastShareTime) {
		return nil, false
	}
	return o.LastShareTime, true
}

// HasLastShareTime returns a boolean if a field has been set.
func (o *RequestForMinerListResponseDataWorkerDatasInner) HasLastShareTime() bool {
	if o != nil && !common.IsNil(o.LastShareTime) {
		return true
	}

	return false
}

// SetLastShareTime gets a reference to the given int64 and assigns it to the LastShareTime field.
func (o *RequestForMinerListResponseDataWorkerDatasInner) SetLastShareTime(v int64) {
	o.LastShareTime = &v
}

func (o RequestForMinerListResponseDataWorkerDatasInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestForMinerListResponseDataWorkerDatasInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.WorkerId) {
		toSerialize["workerId"] = o.WorkerId
	}
	if !common.IsNil(o.WorkerName) {
		toSerialize["workerName"] = o.WorkerName
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.HashRate) {
		toSerialize["hashRate"] = o.HashRate
	}
	if !common.IsNil(o.DayHashRate) {
		toSerialize["dayHashRate"] = o.DayHashRate
	}
	if !common.IsNil(o.RejectRate) {
		toSerialize["rejectRate"] = o.RejectRate
	}
	if !common.IsNil(o.LastShareTime) {
		toSerialize["lastShareTime"] = o.LastShareTime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestForMinerListResponseDataWorkerDatasInner) UnmarshalJSON(data []byte) (err error) {
	varRequestForMinerListResponseDataWorkerDatasInner := _RequestForMinerListResponseDataWorkerDatasInner{}

	err = json.Unmarshal(data, &varRequestForMinerListResponseDataWorkerDatasInner)

	if err != nil {
		return err
	}

	*o = RequestForMinerListResponseDataWorkerDatasInner(varRequestForMinerListResponseDataWorkerDatasInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "workerId")
		delete(additionalProperties, "workerName")
		delete(additionalProperties, "status")
		delete(additionalProperties, "hashRate")
		delete(additionalProperties, "dayHashRate")
		delete(additionalProperties, "rejectRate")
		delete(additionalProperties, "lastShareTime")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestForMinerListResponseDataWorkerDatasInner struct {
	value *RequestForMinerListResponseDataWorkerDatasInner
	isSet bool
}

func (v NullableRequestForMinerListResponseDataWorkerDatasInner) Get() *RequestForMinerListResponseDataWorkerDatasInner {
	return v.value
}

func (v *NullableRequestForMinerListResponseDataWorkerDatasInner) Set(val *RequestForMinerListResponseDataWorkerDatasInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestForMinerListResponseDataWorkerDatasInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestForMinerListResponseDataWorkerDatasInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestForMinerListResponseDataWorkerDatasInner(val *RequestForMinerListResponseDataWorkerDatasInner) *NullableRequestForMinerListResponseDataWorkerDatasInner {
	return &NullableRequestForMinerListResponseDataWorkerDatasInner{value: val, isSet: true}
}

func (v NullableRequestForMinerListResponseDataWorkerDatasInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestForMinerListResponseDataWorkerDatasInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
