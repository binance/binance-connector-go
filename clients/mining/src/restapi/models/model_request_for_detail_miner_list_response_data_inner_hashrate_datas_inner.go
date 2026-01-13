/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the RequestForDetailMinerListResponseDataInnerHashrateDatasInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestForDetailMinerListResponseDataInnerHashrateDatasInner{}

// RequestForDetailMinerListResponseDataInnerHashrateDatasInner struct for RequestForDetailMinerListResponseDataInnerHashrateDatasInner
type RequestForDetailMinerListResponseDataInnerHashrateDatasInner struct {
	Time                 *int64  `json:"time,omitempty"`
	Hashrate             *string `json:"hashrate,omitempty"`
	Reject               *int64  `json:"reject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestForDetailMinerListResponseDataInnerHashrateDatasInner RequestForDetailMinerListResponseDataInnerHashrateDatasInner

// NewRequestForDetailMinerListResponseDataInnerHashrateDatasInner instantiates a new RequestForDetailMinerListResponseDataInnerHashrateDatasInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestForDetailMinerListResponseDataInnerHashrateDatasInner() *RequestForDetailMinerListResponseDataInnerHashrateDatasInner {
	this := RequestForDetailMinerListResponseDataInnerHashrateDatasInner{}
	return &this
}

// NewRequestForDetailMinerListResponseDataInnerHashrateDatasInnerWithDefaults instantiates a new RequestForDetailMinerListResponseDataInnerHashrateDatasInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestForDetailMinerListResponseDataInnerHashrateDatasInnerWithDefaults() *RequestForDetailMinerListResponseDataInnerHashrateDatasInner {
	this := RequestForDetailMinerListResponseDataInnerHashrateDatasInner{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetTime() int64 {
	if o == nil || common.IsNil(o.Time) {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) HasTime() bool {
	if o != nil && !common.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) SetTime(v int64) {
	o.Time = &v
}

// GetHashrate returns the Hashrate field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetHashrate() string {
	if o == nil || common.IsNil(o.Hashrate) {
		var ret string
		return ret
	}
	return *o.Hashrate
}

// GetHashrateOk returns a tuple with the Hashrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetHashrateOk() (*string, bool) {
	if o == nil || common.IsNil(o.Hashrate) {
		return nil, false
	}
	return o.Hashrate, true
}

// HasHashrate returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) HasHashrate() bool {
	if o != nil && !common.IsNil(o.Hashrate) {
		return true
	}

	return false
}

// SetHashrate gets a reference to the given string and assigns it to the Hashrate field.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) SetHashrate(v string) {
	o.Hashrate = &v
}

// GetReject returns the Reject field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetReject() int64 {
	if o == nil || common.IsNil(o.Reject) {
		var ret int64
		return ret
	}
	return *o.Reject
}

// GetRejectOk returns a tuple with the Reject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) GetRejectOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Reject) {
		return nil, false
	}
	return o.Reject, true
}

// HasReject returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) HasReject() bool {
	if o != nil && !common.IsNil(o.Reject) {
		return true
	}

	return false
}

// SetReject gets a reference to the given int64 and assigns it to the Reject field.
func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) SetReject(v int64) {
	o.Reject = &v
}

func (o RequestForDetailMinerListResponseDataInnerHashrateDatasInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestForDetailMinerListResponseDataInnerHashrateDatasInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !common.IsNil(o.Hashrate) {
		toSerialize["hashrate"] = o.Hashrate
	}
	if !common.IsNil(o.Reject) {
		toSerialize["reject"] = o.Reject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) UnmarshalJSON(data []byte) (err error) {
	varRequestForDetailMinerListResponseDataInnerHashrateDatasInner := _RequestForDetailMinerListResponseDataInnerHashrateDatasInner{}

	err = json.Unmarshal(data, &varRequestForDetailMinerListResponseDataInnerHashrateDatasInner)

	if err != nil {
		return err
	}

	*o = RequestForDetailMinerListResponseDataInnerHashrateDatasInner(varRequestForDetailMinerListResponseDataInnerHashrateDatasInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "time")
		delete(additionalProperties, "hashrate")
		delete(additionalProperties, "reject")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner struct {
	value *RequestForDetailMinerListResponseDataInnerHashrateDatasInner
	isSet bool
}

func (v NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) Get() *RequestForDetailMinerListResponseDataInnerHashrateDatasInner {
	return v.value
}

func (v *NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) Set(val *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner(val *RequestForDetailMinerListResponseDataInnerHashrateDatasInner) *NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner {
	return &NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner{value: val, isSet: true}
}

func (v NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestForDetailMinerListResponseDataInnerHashrateDatasInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
