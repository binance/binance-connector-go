/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountSnapshotResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountSnapshotResponse{}

// QueryManagedSubAccountSnapshotResponse struct for QueryManagedSubAccountSnapshotResponse
type QueryManagedSubAccountSnapshotResponse struct {
	Code                 *int64                                                   `json:"code,omitempty"`
	Msg                  *string                                                  `json:"msg,omitempty"`
	SnapshotVos          []QueryManagedSubAccountSnapshotResponseSnapshotVosInner `json:"snapshotVos,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryManagedSubAccountSnapshotResponse QueryManagedSubAccountSnapshotResponse

// NewQueryManagedSubAccountSnapshotResponse instantiates a new QueryManagedSubAccountSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountSnapshotResponse() *QueryManagedSubAccountSnapshotResponse {
	this := QueryManagedSubAccountSnapshotResponse{}
	return &this
}

// NewQueryManagedSubAccountSnapshotResponseWithDefaults instantiates a new QueryManagedSubAccountSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountSnapshotResponseWithDefaults() *QueryManagedSubAccountSnapshotResponse {
	this := QueryManagedSubAccountSnapshotResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *QueryManagedSubAccountSnapshotResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *QueryManagedSubAccountSnapshotResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetSnapshotVos returns the SnapshotVos field value if set, zero value otherwise.
func (o *QueryManagedSubAccountSnapshotResponse) GetSnapshotVos() []QueryManagedSubAccountSnapshotResponseSnapshotVosInner {
	if o == nil || common.IsNil(o.SnapshotVos) {
		var ret []QueryManagedSubAccountSnapshotResponseSnapshotVosInner
		return ret
	}
	return o.SnapshotVos
}

// GetSnapshotVosOk returns a tuple with the SnapshotVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountSnapshotResponse) GetSnapshotVosOk() ([]QueryManagedSubAccountSnapshotResponseSnapshotVosInner, bool) {
	if o == nil || common.IsNil(o.SnapshotVos) {
		return nil, false
	}
	return o.SnapshotVos, true
}

// HasSnapshotVos returns a boolean if a field has been set.
func (o *QueryManagedSubAccountSnapshotResponse) HasSnapshotVos() bool {
	if o != nil && !common.IsNil(o.SnapshotVos) {
		return true
	}

	return false
}

// SetSnapshotVos gets a reference to the given []QueryManagedSubAccountSnapshotResponseSnapshotVosInner and assigns it to the SnapshotVos field.
func (o *QueryManagedSubAccountSnapshotResponse) SetSnapshotVos(v []QueryManagedSubAccountSnapshotResponseSnapshotVosInner) {
	o.SnapshotVos = v
}

func (o QueryManagedSubAccountSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountSnapshotResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountSnapshotResponse := _QueryManagedSubAccountSnapshotResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountSnapshotResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountSnapshotResponse(varQueryManagedSubAccountSnapshotResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "snapshotVos")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountSnapshotResponse struct {
	value *QueryManagedSubAccountSnapshotResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountSnapshotResponse) Get() *QueryManagedSubAccountSnapshotResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountSnapshotResponse) Set(val *QueryManagedSubAccountSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountSnapshotResponse(val *QueryManagedSubAccountSnapshotResponse) *NullableQueryManagedSubAccountSnapshotResponse {
	return &NullableQueryManagedSubAccountSnapshotResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
