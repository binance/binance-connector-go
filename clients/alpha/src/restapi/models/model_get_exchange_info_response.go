/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetExchangeInfoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetExchangeInfoResponse{}

// GetExchangeInfoResponse struct for GetExchangeInfoResponse
type GetExchangeInfoResponse struct {
	Code                 *string                      `json:"code,omitempty"`
	Message              *string                      `json:"message,omitempty"`
	MessageDetail        *string                      `json:"messageDetail,omitempty"`
	Success              *bool                        `json:"success,omitempty"`
	Data                 *GetExchangeInfoResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetExchangeInfoResponse GetExchangeInfoResponse

// NewGetExchangeInfoResponse instantiates a new GetExchangeInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExchangeInfoResponse() *GetExchangeInfoResponse {
	this := GetExchangeInfoResponse{}
	return &this
}

// NewGetExchangeInfoResponseWithDefaults instantiates a new GetExchangeInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExchangeInfoResponseWithDefaults() *GetExchangeInfoResponse {
	this := GetExchangeInfoResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetExchangeInfoResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetExchangeInfoResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetExchangeInfoResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetExchangeInfoResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetExchangeInfoResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetExchangeInfoResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMessageDetail returns the MessageDetail field value if set, zero value otherwise.
func (o *GetExchangeInfoResponse) GetMessageDetail() string {
	if o == nil || common.IsNil(o.MessageDetail) {
		var ret string
		return ret
	}
	return *o.MessageDetail
}

// GetMessageDetailOk returns a tuple with the MessageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponse) GetMessageDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageDetail) {
		return nil, false
	}
	return o.MessageDetail, true
}

// HasMessageDetail returns a boolean if a field has been set.
func (o *GetExchangeInfoResponse) HasMessageDetail() bool {
	if o != nil && !common.IsNil(o.MessageDetail) {
		return true
	}

	return false
}

// SetMessageDetail gets a reference to the given string and assigns it to the MessageDetail field.
func (o *GetExchangeInfoResponse) SetMessageDetail(v string) {
	o.MessageDetail = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetExchangeInfoResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetExchangeInfoResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetExchangeInfoResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExchangeInfoResponse) GetData() GetExchangeInfoResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret GetExchangeInfoResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExchangeInfoResponse) GetDataOk() (*GetExchangeInfoResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExchangeInfoResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetExchangeInfoResponseData and assigns it to the Data field.
func (o *GetExchangeInfoResponse) SetData(v GetExchangeInfoResponseData) {
	o.Data = &v
}

func (o GetExchangeInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExchangeInfoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetExchangeInfoResponse) UnmarshalJSON(data []byte) (err error) {
	varGetExchangeInfoResponse := _GetExchangeInfoResponse{}

	err = json.Unmarshal(data, &varGetExchangeInfoResponse)

	if err != nil {
		return err
	}

	*o = GetExchangeInfoResponse(varGetExchangeInfoResponse)

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

type NullableGetExchangeInfoResponse struct {
	value *GetExchangeInfoResponse
	isSet bool
}

func (v NullableGetExchangeInfoResponse) Get() *GetExchangeInfoResponse {
	return v.value
}

func (v *NullableGetExchangeInfoResponse) Set(val *GetExchangeInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExchangeInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExchangeInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExchangeInfoResponse(val *GetExchangeInfoResponse) *NullableGetExchangeInfoResponse {
	return &NullableGetExchangeInfoResponse{value: val, isSet: true}
}

func (v NullableGetExchangeInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExchangeInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
