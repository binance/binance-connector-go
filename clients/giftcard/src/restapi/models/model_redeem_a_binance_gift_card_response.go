/*
Binance Gift Card REST API

OpenAPI Specification for the Binance Gift Card REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the RedeemABinanceGiftCardResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &RedeemABinanceGiftCardResponse{}

// RedeemABinanceGiftCardResponse struct for RedeemABinanceGiftCardResponse
type RedeemABinanceGiftCardResponse struct {
	Code                 *string                             `json:"code,omitempty"`
	Message              *string                             `json:"message,omitempty"`
	Data                 *RedeemABinanceGiftCardResponseData `json:"data,omitempty"`
	Success              *bool                               `json:"success,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RedeemABinanceGiftCardResponse RedeemABinanceGiftCardResponse

// NewRedeemABinanceGiftCardResponse instantiates a new RedeemABinanceGiftCardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedeemABinanceGiftCardResponse() *RedeemABinanceGiftCardResponse {
	this := RedeemABinanceGiftCardResponse{}
	return &this
}

// NewRedeemABinanceGiftCardResponseWithDefaults instantiates a new RedeemABinanceGiftCardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedeemABinanceGiftCardResponseWithDefaults() *RedeemABinanceGiftCardResponse {
	this := RedeemABinanceGiftCardResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *RedeemABinanceGiftCardResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RedeemABinanceGiftCardResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponse) GetData() RedeemABinanceGiftCardResponseData {
	if o == nil || common.IsNil(o.Data) {
		var ret RedeemABinanceGiftCardResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponse) GetDataOk() (*RedeemABinanceGiftCardResponseData, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RedeemABinanceGiftCardResponseData and assigns it to the Data field.
func (o *RedeemABinanceGiftCardResponse) SetData(v RedeemABinanceGiftCardResponseData) {
	o.Data = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *RedeemABinanceGiftCardResponse) GetSuccess() bool {
	if o == nil || common.IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedeemABinanceGiftCardResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *RedeemABinanceGiftCardResponse) HasSuccess() bool {
	if o != nil && !common.IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *RedeemABinanceGiftCardResponse) SetSuccess(v bool) {
	o.Success = &v
}

func (o RedeemABinanceGiftCardResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedeemABinanceGiftCardResponse) ToMap() (map[string]interface{}, error) {
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

func (o *RedeemABinanceGiftCardResponse) UnmarshalJSON(data []byte) (err error) {
	varRedeemABinanceGiftCardResponse := _RedeemABinanceGiftCardResponse{}

	err = json.Unmarshal(data, &varRedeemABinanceGiftCardResponse)

	if err != nil {
		return err
	}

	*o = RedeemABinanceGiftCardResponse(varRedeemABinanceGiftCardResponse)

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

type NullableRedeemABinanceGiftCardResponse struct {
	value *RedeemABinanceGiftCardResponse
	isSet bool
}

func (v NullableRedeemABinanceGiftCardResponse) Get() *RedeemABinanceGiftCardResponse {
	return v.value
}

func (v *NullableRedeemABinanceGiftCardResponse) Set(val *RedeemABinanceGiftCardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedeemABinanceGiftCardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedeemABinanceGiftCardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedeemABinanceGiftCardResponse(val *RedeemABinanceGiftCardResponse) *NullableRedeemABinanceGiftCardResponse {
	return &NullableRedeemABinanceGiftCardResponse{value: val, isSet: true}
}

func (v NullableRedeemABinanceGiftCardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedeemABinanceGiftCardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
