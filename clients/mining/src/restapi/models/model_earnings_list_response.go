/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the EarningsListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &EarningsListResponse{}

// EarningsListResponse struct for EarningsListResponse
type EarningsListResponse struct {
	Code                 *int64                    `json:"code,omitempty"`
	Msg                  *string                   `json:"msg,omitempty"`
	Data                 *EarningsListResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EarningsListResponse EarningsListResponse

// NewEarningsListResponse instantiates a new EarningsListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsListResponse() *EarningsListResponse {
	this := EarningsListResponse{}
	return &this
}

// NewEarningsListResponseWithDefaults instantiates a new EarningsListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsListResponseWithDefaults() *EarningsListResponse {
	this := EarningsListResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EarningsListResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EarningsListResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *EarningsListResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *EarningsListResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *EarningsListResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *EarningsListResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EarningsListResponse) GetData() EarningsListResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret EarningsListResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsListResponse) GetDataOk() (*EarningsListResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EarningsListResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given EarningsListResponseData and assigns it to the Data field.
func (o *EarningsListResponse) SetData(v EarningsListResponseData) {
	o.Data = &v
}

func (o EarningsListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarningsListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *EarningsListResponse) UnmarshalJSON(data []byte) (err error) {
	varEarningsListResponse := _EarningsListResponse{}

	err = json.Unmarshal(data, &varEarningsListResponse)

	if err != nil {
		return err
	}

	*o = EarningsListResponse(varEarningsListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEarningsListResponse struct {
	value *EarningsListResponse
	isSet bool
}

func (v NullableEarningsListResponse) Get() *EarningsListResponse {
	return v.value
}

func (v *NullableEarningsListResponse) Set(val *EarningsListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsListResponse(val *EarningsListResponse) *NullableEarningsListResponse {
	return &NullableEarningsListResponse{value: val, isSet: true}
}

func (v NullableEarningsListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
