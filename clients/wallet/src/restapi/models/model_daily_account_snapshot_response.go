/*
Binance Wallet REST API

OpenAPI Specification for the Binance Wallet REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the DailyAccountSnapshotResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &DailyAccountSnapshotResponse{}

// DailyAccountSnapshotResponse struct for DailyAccountSnapshotResponse
type DailyAccountSnapshotResponse struct {
	Code                 *int64                                         `json:"code,omitempty"`
	Msg                  *string                                        `json:"msg,omitempty"`
	SnapshotVos          []DailyAccountSnapshotResponseSnapshotVosInner `json:"snapshotVos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DailyAccountSnapshotResponse DailyAccountSnapshotResponse

// NewDailyAccountSnapshotResponse instantiates a new DailyAccountSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailyAccountSnapshotResponse() *DailyAccountSnapshotResponse {
	this := DailyAccountSnapshotResponse{}
	return &this
}

// NewDailyAccountSnapshotResponseWithDefaults instantiates a new DailyAccountSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyAccountSnapshotResponseWithDefaults() *DailyAccountSnapshotResponse {
	this := DailyAccountSnapshotResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *DailyAccountSnapshotResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *DailyAccountSnapshotResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetSnapshotVos returns the SnapshotVos field value if set, zero value otherwise.
func (o *DailyAccountSnapshotResponse) GetSnapshotVos() []DailyAccountSnapshotResponseSnapshotVosInner {
	if o == nil || common.IsNil(o.SnapshotVos) {
		var ret []DailyAccountSnapshotResponseSnapshotVosInner
		return ret
	}
	return o.SnapshotVos
}

// GetSnapshotVosOk returns a tuple with the SnapshotVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailyAccountSnapshotResponse) GetSnapshotVosOk() ([]DailyAccountSnapshotResponseSnapshotVosInner, bool) {
	if o == nil || common.IsNil(o.SnapshotVos) {
		return nil, false
	}
	return o.SnapshotVos, true
}

// HasSnapshotVos returns a boolean if a field has been set.
func (o *DailyAccountSnapshotResponse) HasSnapshotVos() bool {
	if o != nil && !common.IsNil(o.SnapshotVos) {
		return true
	}

	return false
}

// SetSnapshotVos gets a reference to the given []DailyAccountSnapshotResponseSnapshotVosInner and assigns it to the SnapshotVos field.
func (o *DailyAccountSnapshotResponse) SetSnapshotVos(v []DailyAccountSnapshotResponseSnapshotVosInner) {
	o.SnapshotVos = v
}

func (o DailyAccountSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DailyAccountSnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.SnapshotVos) {
		toSerialize["snapshotVos"] = o.SnapshotVos
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DailyAccountSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	varDailyAccountSnapshotResponse := _DailyAccountSnapshotResponse{}

	err = json.Unmarshal(data, &varDailyAccountSnapshotResponse)

	if err != nil {
		return err
	}

	*o = DailyAccountSnapshotResponse(varDailyAccountSnapshotResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "snapshotVos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDailyAccountSnapshotResponse struct {
	value *DailyAccountSnapshotResponse
	isSet bool
}

func (v NullableDailyAccountSnapshotResponse) Get() *DailyAccountSnapshotResponse {
	return v.value
}

func (v *NullableDailyAccountSnapshotResponse) Set(val *DailyAccountSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDailyAccountSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDailyAccountSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailyAccountSnapshotResponse(val *DailyAccountSnapshotResponse) *NullableDailyAccountSnapshotResponse {
	return &NullableDailyAccountSnapshotResponse{value: val, isSet: true}
}

func (v NullableDailyAccountSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailyAccountSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
