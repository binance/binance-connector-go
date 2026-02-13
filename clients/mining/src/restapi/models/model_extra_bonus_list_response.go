/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ExtraBonusListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ExtraBonusListResponse{}

// ExtraBonusListResponse struct for ExtraBonusListResponse
type ExtraBonusListResponse struct {
	Code                 *int64                      `json:"code,omitempty"`
	Msg                  *string                     `json:"msg,omitempty"`
	Data                 *ExtraBonusListResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExtraBonusListResponse ExtraBonusListResponse

// NewExtraBonusListResponse instantiates a new ExtraBonusListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraBonusListResponse() *ExtraBonusListResponse {
	this := ExtraBonusListResponse{}
	return &this
}

// NewExtraBonusListResponseWithDefaults instantiates a new ExtraBonusListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtraBonusListResponseWithDefaults() *ExtraBonusListResponse {
	this := ExtraBonusListResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ExtraBonusListResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ExtraBonusListResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *ExtraBonusListResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ExtraBonusListResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ExtraBonusListResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ExtraBonusListResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExtraBonusListResponse) GetData() ExtraBonusListResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret ExtraBonusListResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtraBonusListResponse) GetDataOk() (*ExtraBonusListResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExtraBonusListResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ExtraBonusListResponseData and assigns it to the Data field.
func (o *ExtraBonusListResponse) SetData(v ExtraBonusListResponseData) {
	o.Data = &v
}

func (o ExtraBonusListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtraBonusListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ExtraBonusListResponse) UnmarshalJSON(data []byte) (err error) {
	varExtraBonusListResponse := _ExtraBonusListResponse{}

	err = json.Unmarshal(data, &varExtraBonusListResponse)

	if err != nil {
		return err
	}

	*o = ExtraBonusListResponse(varExtraBonusListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExtraBonusListResponse struct {
	value *ExtraBonusListResponse
	isSet bool
}

func (v NullableExtraBonusListResponse) Get() *ExtraBonusListResponse {
	return v.value
}

func (v *NullableExtraBonusListResponse) Set(val *ExtraBonusListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraBonusListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraBonusListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraBonusListResponse(val *ExtraBonusListResponse) *NullableExtraBonusListResponse {
	return &NullableExtraBonusListResponse{value: val, isSet: true}
}

func (v NullableExtraBonusListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraBonusListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
