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

// checks if the CancelHashrateResaleConfigurationResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelHashrateResaleConfigurationResponse{}

// CancelHashrateResaleConfigurationResponse struct for CancelHashrateResaleConfigurationResponse
type CancelHashrateResaleConfigurationResponse struct {
	Code                 *int64  `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	Data                 *bool   `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelHashrateResaleConfigurationResponse CancelHashrateResaleConfigurationResponse

// NewCancelHashrateResaleConfigurationResponse instantiates a new CancelHashrateResaleConfigurationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelHashrateResaleConfigurationResponse() *CancelHashrateResaleConfigurationResponse {
	this := CancelHashrateResaleConfigurationResponse{}
	return &this
}

// NewCancelHashrateResaleConfigurationResponseWithDefaults instantiates a new CancelHashrateResaleConfigurationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelHashrateResaleConfigurationResponseWithDefaults() *CancelHashrateResaleConfigurationResponse {
	this := CancelHashrateResaleConfigurationResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelHashrateResaleConfigurationResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelHashrateResaleConfigurationResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelHashrateResaleConfigurationResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *CancelHashrateResaleConfigurationResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelHashrateResaleConfigurationResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelHashrateResaleConfigurationResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelHashrateResaleConfigurationResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelHashrateResaleConfigurationResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CancelHashrateResaleConfigurationResponse) GetData() bool {
	if o == nil || common.IsNil(o.Data) {
		var ret bool
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelHashrateResaleConfigurationResponse) GetDataOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CancelHashrateResaleConfigurationResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given bool and assigns it to the Data field.
func (o *CancelHashrateResaleConfigurationResponse) SetData(v bool) {
	o.Data = &v
}

func (o CancelHashrateResaleConfigurationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelHashrateResaleConfigurationResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelHashrateResaleConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelHashrateResaleConfigurationResponse := _CancelHashrateResaleConfigurationResponse{}

	err = json.Unmarshal(data, &varCancelHashrateResaleConfigurationResponse)

	if err != nil {
		return err
	}

	*o = CancelHashrateResaleConfigurationResponse(varCancelHashrateResaleConfigurationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelHashrateResaleConfigurationResponse struct {
	value *CancelHashrateResaleConfigurationResponse
	isSet bool
}

func (v NullableCancelHashrateResaleConfigurationResponse) Get() *CancelHashrateResaleConfigurationResponse {
	return v.value
}

func (v *NullableCancelHashrateResaleConfigurationResponse) Set(val *CancelHashrateResaleConfigurationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelHashrateResaleConfigurationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelHashrateResaleConfigurationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelHashrateResaleConfigurationResponse(val *CancelHashrateResaleConfigurationResponse) *NullableCancelHashrateResaleConfigurationResponse {
	return &NullableCancelHashrateResaleConfigurationResponse{value: val, isSet: true}
}

func (v NullableCancelHashrateResaleConfigurationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelHashrateResaleConfigurationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
