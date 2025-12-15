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

// checks if the RequestForDetailMinerListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RequestForDetailMinerListResponse{}

// RequestForDetailMinerListResponse struct for RequestForDetailMinerListResponse
type RequestForDetailMinerListResponse struct {
	Code                 *int64                                       `json:"code,omitempty"`
	Msg                  *string                                      `json:"msg,omitempty"`
	Data                 []RequestForDetailMinerListResponseDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestForDetailMinerListResponse RequestForDetailMinerListResponse

// NewRequestForDetailMinerListResponse instantiates a new RequestForDetailMinerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestForDetailMinerListResponse() *RequestForDetailMinerListResponse {
	this := RequestForDetailMinerListResponse{}
	return &this
}

// NewRequestForDetailMinerListResponseWithDefaults instantiates a new RequestForDetailMinerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestForDetailMinerListResponseWithDefaults() *RequestForDetailMinerListResponse {
	this := RequestForDetailMinerListResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *RequestForDetailMinerListResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *RequestForDetailMinerListResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RequestForDetailMinerListResponse) GetData() []RequestForDetailMinerListResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []RequestForDetailMinerListResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestForDetailMinerListResponse) GetDataOk() ([]RequestForDetailMinerListResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RequestForDetailMinerListResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RequestForDetailMinerListResponseDataInner and assigns it to the Data field.
func (o *RequestForDetailMinerListResponse) SetData(v []RequestForDetailMinerListResponseDataInner) {
	o.Data = v
}

func (o RequestForDetailMinerListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestForDetailMinerListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *RequestForDetailMinerListResponse) UnmarshalJSON(data []byte) (err error) {
	varRequestForDetailMinerListResponse := _RequestForDetailMinerListResponse{}

	err = json.Unmarshal(data, &varRequestForDetailMinerListResponse)

	if err != nil {
		return err
	}

	*o = RequestForDetailMinerListResponse(varRequestForDetailMinerListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestForDetailMinerListResponse struct {
	value *RequestForDetailMinerListResponse
	isSet bool
}

func (v NullableRequestForDetailMinerListResponse) Get() *RequestForDetailMinerListResponse {
	return v.value
}

func (v *NullableRequestForDetailMinerListResponse) Set(val *RequestForDetailMinerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestForDetailMinerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestForDetailMinerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestForDetailMinerListResponse(val *RequestForDetailMinerListResponse) *NullableRequestForDetailMinerListResponse {
	return &NullableRequestForDetailMinerListResponse{value: val, isSet: true}
}

func (v NullableRequestForDetailMinerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestForDetailMinerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
