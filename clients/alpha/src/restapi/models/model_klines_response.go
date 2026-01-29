/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the KlinesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KlinesResponse{}

// KlinesResponse struct for KlinesResponse
type KlinesResponse struct {
	Code                 *string                  `json:"code,omitempty"`
	Message              *string                  `json:"message,omitempty"`
	MessageDetail        *string                  `json:"messageDetail,omitempty"`
	Success              *bool                    `json:"success,omitempty"`
	Data                 []KlinesResponseDataItem `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KlinesResponse KlinesResponse

// NewKlinesResponse instantiates a new KlinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlinesResponse() *KlinesResponse {
	this := KlinesResponse{}
	return &this
}

// NewKlinesResponseWithDefaults instantiates a new KlinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlinesResponseWithDefaults() *KlinesResponse {
	this := KlinesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *KlinesResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlinesResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *KlinesResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *KlinesResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KlinesResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlinesResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KlinesResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KlinesResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMessageDetail returns the MessageDetail field value if set, zero value otherwise.
func (o *KlinesResponse) GetMessageDetail() string {
	if o == nil || common.IsNil(o.MessageDetail) {
		var ret string
		return ret
	}
	return *o.MessageDetail
}

// GetMessageDetailOk returns a tuple with the MessageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlinesResponse) GetMessageDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageDetail) {
		return nil, false
	}
	return o.MessageDetail, true
}

// HasMessageDetail returns a boolean if a field has been set.
func (o *KlinesResponse) HasMessageDetail() bool {
	if o != nil && !common.IsNil(o.MessageDetail) {
		return true
	}

	return false
}

// SetMessageDetail gets a reference to the given string and assigns it to the MessageDetail field.
func (o *KlinesResponse) SetMessageDetail(v string) {
	o.MessageDetail = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *KlinesResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlinesResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *KlinesResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *KlinesResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KlinesResponse) GetData() []KlinesResponseDataItem {
	if o == nil || common.IsNil(o.Data) {
		var ret []KlinesResponseDataItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KlinesResponse) GetDataOk() ([]KlinesResponseDataItem, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KlinesResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []KlinesResponseDataItem and assigns it to the Data field.
func (o *KlinesResponse) SetData(v []KlinesResponseDataItem) {
	o.Data = v
}

func (o KlinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KlinesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !common.IsNil(o.MessageDetail) {
		toSerialize["messageDetail"] = o.MessageDetail
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KlinesResponse) UnmarshalJSON(data []byte) (err error) {
	varKlinesResponse := _KlinesResponse{}

	err = json.Unmarshal(data, &varKlinesResponse)

	if err != nil {
		return err
	}

	*o = KlinesResponse(varKlinesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "messageDetail")
		delete(additionalProperties, "success")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKlinesResponse struct {
	value *KlinesResponse
	isSet bool
}

func (v NullableKlinesResponse) Get() *KlinesResponse {
	return v.value
}

func (v *NullableKlinesResponse) Set(val *KlinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKlinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKlinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlinesResponse(val *KlinesResponse) *NullableKlinesResponse {
	return &NullableKlinesResponse{value: val, isSet: true}
}

func (v NullableKlinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
