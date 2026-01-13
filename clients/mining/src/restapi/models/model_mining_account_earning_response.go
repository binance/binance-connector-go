/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the MiningAccountEarningResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &MiningAccountEarningResponse{}

// MiningAccountEarningResponse struct for MiningAccountEarningResponse
type MiningAccountEarningResponse struct {
	Code                 *int64                            `json:"code,omitempty"`
	Msg                  *string                           `json:"msg,omitempty"`
	Data                 *MiningAccountEarningResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MiningAccountEarningResponse MiningAccountEarningResponse

// NewMiningAccountEarningResponse instantiates a new MiningAccountEarningResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiningAccountEarningResponse() *MiningAccountEarningResponse {
	this := MiningAccountEarningResponse{}
	return &this
}

// NewMiningAccountEarningResponseWithDefaults instantiates a new MiningAccountEarningResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiningAccountEarningResponseWithDefaults() *MiningAccountEarningResponse {
	this := MiningAccountEarningResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MiningAccountEarningResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MiningAccountEarningResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *MiningAccountEarningResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *MiningAccountEarningResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *MiningAccountEarningResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *MiningAccountEarningResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MiningAccountEarningResponse) GetData() MiningAccountEarningResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret MiningAccountEarningResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiningAccountEarningResponse) GetDataOk() (*MiningAccountEarningResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MiningAccountEarningResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MiningAccountEarningResponseData and assigns it to the Data field.
func (o *MiningAccountEarningResponse) SetData(v MiningAccountEarningResponseData) {
	o.Data = &v
}

func (o MiningAccountEarningResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiningAccountEarningResponse) ToMap() (map[string]interface{}, error) {
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

func (o *MiningAccountEarningResponse) UnmarshalJSON(data []byte) (err error) {
	varMiningAccountEarningResponse := _MiningAccountEarningResponse{}

	err = json.Unmarshal(data, &varMiningAccountEarningResponse)

	if err != nil {
		return err
	}

	*o = MiningAccountEarningResponse(varMiningAccountEarningResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMiningAccountEarningResponse struct {
	value *MiningAccountEarningResponse
	isSet bool
}

func (v NullableMiningAccountEarningResponse) Get() *MiningAccountEarningResponse {
	return v.value
}

func (v *NullableMiningAccountEarningResponse) Set(val *MiningAccountEarningResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMiningAccountEarningResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMiningAccountEarningResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiningAccountEarningResponse(val *MiningAccountEarningResponse) *NullableMiningAccountEarningResponse {
	return &NullableMiningAccountEarningResponse{value: val, isSet: true}
}

func (v NullableMiningAccountEarningResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiningAccountEarningResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
