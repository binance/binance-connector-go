/*
Binance C2C REST API

OpenAPI Specification for the Binance C2C REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetC2CTradeHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetC2CTradeHistoryResponse{}

// GetC2CTradeHistoryResponse struct for GetC2CTradeHistoryResponse
type GetC2CTradeHistoryResponse struct {
	Code                 *string                               `json:"code,omitempty"`
	Message              *string                               `json:"message,omitempty"`
	Data                 []GetC2CTradeHistoryResponseDataInner `json:"data,omitempty"`
	Total                *int64                                `json:"total,omitempty"`
	Success              *bool                                 `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetC2CTradeHistoryResponse GetC2CTradeHistoryResponse

// NewGetC2CTradeHistoryResponse instantiates a new GetC2CTradeHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetC2CTradeHistoryResponse() *GetC2CTradeHistoryResponse {
	this := GetC2CTradeHistoryResponse{}
	return &this
}

// NewGetC2CTradeHistoryResponseWithDefaults instantiates a new GetC2CTradeHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetC2CTradeHistoryResponseWithDefaults() *GetC2CTradeHistoryResponse {
	this := GetC2CTradeHistoryResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetC2CTradeHistoryResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetC2CTradeHistoryResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetData() []GetC2CTradeHistoryResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []GetC2CTradeHistoryResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetDataOk() ([]GetC2CTradeHistoryResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetC2CTradeHistoryResponseDataInner and assigns it to the Data field.
func (o *GetC2CTradeHistoryResponse) SetData(v []GetC2CTradeHistoryResponseDataInner) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetC2CTradeHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetC2CTradeHistoryResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetC2CTradeHistoryResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetC2CTradeHistoryResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetC2CTradeHistoryResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o GetC2CTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetC2CTradeHistoryResponse) ToMap() (map[string]interface{}, error) {
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
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetC2CTradeHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetC2CTradeHistoryResponse := _GetC2CTradeHistoryResponse{}

	err = json.Unmarshal(data, &varGetC2CTradeHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetC2CTradeHistoryResponse(varGetC2CTradeHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		delete(additionalProperties, "total")
		delete(additionalProperties, "success")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetC2CTradeHistoryResponse struct {
	value *GetC2CTradeHistoryResponse
	isSet bool
}

func (v NullableGetC2CTradeHistoryResponse) Get() *GetC2CTradeHistoryResponse {
	return v.value
}

func (v *NullableGetC2CTradeHistoryResponse) Set(val *GetC2CTradeHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetC2CTradeHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetC2CTradeHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetC2CTradeHistoryResponse(val *GetC2CTradeHistoryResponse) *NullableGetC2CTradeHistoryResponse {
	return &NullableGetC2CTradeHistoryResponse{value: val, isSet: true}
}

func (v NullableGetC2CTradeHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetC2CTradeHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
