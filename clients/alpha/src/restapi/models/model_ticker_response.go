/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the TickerResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TickerResponse{}

// TickerResponse struct for TickerResponse
type TickerResponse struct {
	Code                 *string             `json:"code,omitempty"`
	Message              *string             `json:"message,omitempty"`
	MessageDetail        *string             `json:"messageDetail,omitempty"`
	Data                 *TickerResponseData `json:"data,omitempty"`
	Success              *bool               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TickerResponse TickerResponse

// NewTickerResponse instantiates a new TickerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTickerResponse() *TickerResponse {
	this := TickerResponse{}
	return &this
}

// NewTickerResponseWithDefaults instantiates a new TickerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTickerResponseWithDefaults() *TickerResponse {
	this := TickerResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *TickerResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *TickerResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *TickerResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TickerResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TickerResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TickerResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMessageDetail returns the MessageDetail field value if set, zero value otherwise.
func (o *TickerResponse) GetMessageDetail() string {
	if o == nil || common.IsNil(o.MessageDetail) {
		var ret string
		return ret
	}
	return *o.MessageDetail
}

// GetMessageDetailOk returns a tuple with the MessageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerResponse) GetMessageDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageDetail) {
		return nil, false
	}
	return o.MessageDetail, true
}

// HasMessageDetail returns a boolean if a field has been set.
func (o *TickerResponse) HasMessageDetail() bool {
	if o != nil && !common.IsNil(o.MessageDetail) {
		return true
	}

	return false
}

// SetMessageDetail gets a reference to the given string and assigns it to the MessageDetail field.
func (o *TickerResponse) SetMessageDetail(v string) {
	o.MessageDetail = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TickerResponse) GetData() TickerResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret TickerResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerResponse) GetDataOk() (*TickerResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TickerResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TickerResponseData and assigns it to the Data field.
func (o *TickerResponse) SetData(v TickerResponseData) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *TickerResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TickerResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *TickerResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *TickerResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o TickerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TickerResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TickerResponse) UnmarshalJSON(data []byte) (err error) {
	varTickerResponse := _TickerResponse{}

	err = json.Unmarshal(data, &varTickerResponse)

	if err != nil {
		return err
	}

	*o = TickerResponse(varTickerResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "messageDetail")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTickerResponse struct {
	value *TickerResponse
	isSet bool
}

func (v NullableTickerResponse) Get() *TickerResponse {
	return v.value
}

func (v *NullableTickerResponse) Set(val *TickerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTickerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTickerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTickerResponse(val *TickerResponse) *NullableTickerResponse {
	return &NullableTickerResponse{value: val, isSet: true}
}

func (v NullableTickerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTickerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
