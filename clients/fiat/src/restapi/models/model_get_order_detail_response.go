/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetOrderDetailResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetOrderDetailResponse{}

// GetOrderDetailResponse struct for GetOrderDetailResponse
type GetOrderDetailResponse struct {
	Code                 *string                     `json:"code,omitempty"`
	Message              *string                     `json:"message,omitempty"`
	Data                 *GetOrderDetailResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetOrderDetailResponse GetOrderDetailResponse

// NewGetOrderDetailResponse instantiates a new GetOrderDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrderDetailResponse() *GetOrderDetailResponse {
	this := GetOrderDetailResponse{}
	return &this
}

// NewGetOrderDetailResponseWithDefaults instantiates a new GetOrderDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrderDetailResponseWithDefaults() *GetOrderDetailResponse {
	this := GetOrderDetailResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetOrderDetailResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetOrderDetailResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetOrderDetailResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetOrderDetailResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetOrderDetailResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetOrderDetailResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOrderDetailResponse) GetData() GetOrderDetailResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret GetOrderDetailResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrderDetailResponse) GetDataOk() (*GetOrderDetailResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOrderDetailResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetOrderDetailResponseData and assigns it to the Data field.
func (o *GetOrderDetailResponse) SetData(v GetOrderDetailResponseData) {
	o.Data = &v
}

func (o GetOrderDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrderDetailResponse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetOrderDetailResponse) UnmarshalJSON(data []byte) (err error) {
	varGetOrderDetailResponse := _GetOrderDetailResponse{}

	err = json.Unmarshal(data, &varGetOrderDetailResponse)

	if err != nil {
		return err
	}

	*o = GetOrderDetailResponse(varGetOrderDetailResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetOrderDetailResponse struct {
	value *GetOrderDetailResponse
	isSet bool
}

func (v NullableGetOrderDetailResponse) Get() *GetOrderDetailResponse {
	return v.value
}

func (v *NullableGetOrderDetailResponse) Set(val *GetOrderDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrderDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrderDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrderDetailResponse(val *GetOrderDetailResponse) *NullableGetOrderDetailResponse {
	return &NullableGetOrderDetailResponse{value: val, isSet: true}
}

func (v NullableGetOrderDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrderDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
