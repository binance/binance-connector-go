/*
Binance Copy Trading REST API

OpenAPI Specification for the Binance Copy Trading REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFuturesLeadTraderStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesLeadTraderStatusResponse{}

// GetFuturesLeadTraderStatusResponse struct for GetFuturesLeadTraderStatusResponse
type GetFuturesLeadTraderStatusResponse struct {
	Code                 *string                                 `json:"code,omitempty"`
	Message              *string                                 `json:"message,omitempty"`
	Data                 *GetFuturesLeadTraderStatusResponseData `json:"data,omitempty"`
	Success              *bool                                   `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesLeadTraderStatusResponse GetFuturesLeadTraderStatusResponse

// NewGetFuturesLeadTraderStatusResponse instantiates a new GetFuturesLeadTraderStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesLeadTraderStatusResponse() *GetFuturesLeadTraderStatusResponse {
	this := GetFuturesLeadTraderStatusResponse{}
	return &this
}

// NewGetFuturesLeadTraderStatusResponseWithDefaults instantiates a new GetFuturesLeadTraderStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesLeadTraderStatusResponseWithDefaults() *GetFuturesLeadTraderStatusResponse {
	this := GetFuturesLeadTraderStatusResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetFuturesLeadTraderStatusResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetFuturesLeadTraderStatusResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponse) GetData() GetFuturesLeadTraderStatusResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret GetFuturesLeadTraderStatusResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponse) GetDataOk() (*GetFuturesLeadTraderStatusResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetFuturesLeadTraderStatusResponseData and assigns it to the Data field.
func (o *GetFuturesLeadTraderStatusResponse) SetData(v GetFuturesLeadTraderStatusResponseData) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetFuturesLeadTraderStatusResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTraderStatusResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetFuturesLeadTraderStatusResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetFuturesLeadTraderStatusResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o GetFuturesLeadTraderStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesLeadTraderStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Message) {
		toSerialize["message"] = o.Message
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

func (o *GetFuturesLeadTraderStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesLeadTraderStatusResponse := _GetFuturesLeadTraderStatusResponse{}

	err = json.Unmarshal(data, &varGetFuturesLeadTraderStatusResponse)

	if err != nil {
		return err
	}

	*o = GetFuturesLeadTraderStatusResponse(varGetFuturesLeadTraderStatusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesLeadTraderStatusResponse struct {
	value *GetFuturesLeadTraderStatusResponse
	isSet bool
}

func (v NullableGetFuturesLeadTraderStatusResponse) Get() *GetFuturesLeadTraderStatusResponse {
	return v.value
}

func (v *NullableGetFuturesLeadTraderStatusResponse) Set(val *GetFuturesLeadTraderStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesLeadTraderStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesLeadTraderStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesLeadTraderStatusResponse(val *GetFuturesLeadTraderStatusResponse) *NullableGetFuturesLeadTraderStatusResponse {
	return &NullableGetFuturesLeadTraderStatusResponse{value: val, isSet: true}
}

func (v NullableGetFuturesLeadTraderStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesLeadTraderStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
