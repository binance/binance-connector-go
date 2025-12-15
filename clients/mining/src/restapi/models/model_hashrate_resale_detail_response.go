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

// checks if the HashrateResaleDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &HashrateResaleDetailResponse{}

// HashrateResaleDetailResponse struct for HashrateResaleDetailResponse
type HashrateResaleDetailResponse struct {
	Code                 *int64                            `json:"code,omitempty"`
	Msg                  *string                           `json:"msg,omitempty"`
	Data                 *HashrateResaleDetailResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HashrateResaleDetailResponse HashrateResaleDetailResponse

// NewHashrateResaleDetailResponse instantiates a new HashrateResaleDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashrateResaleDetailResponse() *HashrateResaleDetailResponse {
	this := HashrateResaleDetailResponse{}
	return &this
}

// NewHashrateResaleDetailResponseWithDefaults instantiates a new HashrateResaleDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashrateResaleDetailResponseWithDefaults() *HashrateResaleDetailResponse {
	this := HashrateResaleDetailResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *HashrateResaleDetailResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *HashrateResaleDetailResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HashrateResaleDetailResponse) GetData() HashrateResaleDetailResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret HashrateResaleDetailResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashrateResaleDetailResponse) GetDataOk() (*HashrateResaleDetailResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HashrateResaleDetailResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given HashrateResaleDetailResponseData and assigns it to the Data field.
func (o *HashrateResaleDetailResponse) SetData(v HashrateResaleDetailResponseData) {
	o.Data = &v
}

func (o HashrateResaleDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashrateResaleDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HashrateResaleDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varHashrateResaleDetailResponse := _HashrateResaleDetailResponse{}

	err = json.Unmarshal(data, &varHashrateResaleDetailResponse)

	if err != nil {
		return err
	}

	*o = HashrateResaleDetailResponse(varHashrateResaleDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHashrateResaleDetailResponse struct {
	value *HashrateResaleDetailResponse
	isSet bool
}

func (v NullableHashrateResaleDetailResponse) Get() *HashrateResaleDetailResponse {
	return v.value
}

func (v *NullableHashrateResaleDetailResponse) Set(val *HashrateResaleDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHashrateResaleDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHashrateResaleDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashrateResaleDetailResponse(val *HashrateResaleDetailResponse) *NullableHashrateResaleDetailResponse {
	return &NullableHashrateResaleDetailResponse{value: val, isSet: true}
}

func (v NullableHashrateResaleDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashrateResaleDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
