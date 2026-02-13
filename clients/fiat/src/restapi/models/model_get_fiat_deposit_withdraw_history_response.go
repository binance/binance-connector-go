/*
Binance Fiat REST API

OpenAPI Specification for the Binance Fiat REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetFiatDepositWithdrawHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFiatDepositWithdrawHistoryResponse{}

// GetFiatDepositWithdrawHistoryResponse struct for GetFiatDepositWithdrawHistoryResponse
type GetFiatDepositWithdrawHistoryResponse struct {
	Code                 *string                                          `json:"code,omitempty"`
	Message              *string                                          `json:"message,omitempty"`
	Data                 []GetFiatDepositWithdrawHistoryResponseDataInner `json:"data,omitempty"`
	Total                *int64                                           `json:"total,omitempty"`
	Success              *bool                                            `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFiatDepositWithdrawHistoryResponse GetFiatDepositWithdrawHistoryResponse

// NewGetFiatDepositWithdrawHistoryResponse instantiates a new GetFiatDepositWithdrawHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiatDepositWithdrawHistoryResponse() *GetFiatDepositWithdrawHistoryResponse {
	this := GetFiatDepositWithdrawHistoryResponse{}
	return &this
}

// NewGetFiatDepositWithdrawHistoryResponseWithDefaults instantiates a new GetFiatDepositWithdrawHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiatDepositWithdrawHistoryResponseWithDefaults() *GetFiatDepositWithdrawHistoryResponse {
	this := GetFiatDepositWithdrawHistoryResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetFiatDepositWithdrawHistoryResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetFiatDepositWithdrawHistoryResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponse) GetData() []GetFiatDepositWithdrawHistoryResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []GetFiatDepositWithdrawHistoryResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) GetDataOk() ([]GetFiatDepositWithdrawHistoryResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetFiatDepositWithdrawHistoryResponseDataInner and assigns it to the Data field.
func (o *GetFiatDepositWithdrawHistoryResponse) SetData(v []GetFiatDepositWithdrawHistoryResponseDataInner) {
	o.Data = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetFiatDepositWithdrawHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetFiatDepositWithdrawHistoryResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetFiatDepositWithdrawHistoryResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetFiatDepositWithdrawHistoryResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o GetFiatDepositWithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiatDepositWithdrawHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFiatDepositWithdrawHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFiatDepositWithdrawHistoryResponse := _GetFiatDepositWithdrawHistoryResponse{}

	err = json.Unmarshal(data, &varGetFiatDepositWithdrawHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetFiatDepositWithdrawHistoryResponse(varGetFiatDepositWithdrawHistoryResponse)

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

type NullableGetFiatDepositWithdrawHistoryResponse struct {
	value *GetFiatDepositWithdrawHistoryResponse
	isSet bool
}

func (v NullableGetFiatDepositWithdrawHistoryResponse) Get() *GetFiatDepositWithdrawHistoryResponse {
	return v.value
}

func (v *NullableGetFiatDepositWithdrawHistoryResponse) Set(val *GetFiatDepositWithdrawHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiatDepositWithdrawHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiatDepositWithdrawHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiatDepositWithdrawHistoryResponse(val *GetFiatDepositWithdrawHistoryResponse) *NullableGetFiatDepositWithdrawHistoryResponse {
	return &NullableGetFiatDepositWithdrawHistoryResponse{value: val, isSet: true}
}

func (v NullableGetFiatDepositWithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiatDepositWithdrawHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
