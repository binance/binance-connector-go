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

// checks if the GetFuturesLeadTradingSymbolWhitelistResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFuturesLeadTradingSymbolWhitelistResponse{}

// GetFuturesLeadTradingSymbolWhitelistResponse struct for GetFuturesLeadTradingSymbolWhitelistResponse
type GetFuturesLeadTradingSymbolWhitelistResponse struct {
	Code                 *string                                                 `json:"code,omitempty"`
	Message              *string                                                 `json:"message,omitempty"`
	Data                 []GetFuturesLeadTradingSymbolWhitelistResponseDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetFuturesLeadTradingSymbolWhitelistResponse GetFuturesLeadTradingSymbolWhitelistResponse

// NewGetFuturesLeadTradingSymbolWhitelistResponse instantiates a new GetFuturesLeadTradingSymbolWhitelistResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFuturesLeadTradingSymbolWhitelistResponse() *GetFuturesLeadTradingSymbolWhitelistResponse {
	this := GetFuturesLeadTradingSymbolWhitelistResponse{}
	return &this
}

// NewGetFuturesLeadTradingSymbolWhitelistResponseWithDefaults instantiates a new GetFuturesLeadTradingSymbolWhitelistResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFuturesLeadTradingSymbolWhitelistResponseWithDefaults() *GetFuturesLeadTradingSymbolWhitelistResponse {
	this := GetFuturesLeadTradingSymbolWhitelistResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetData() []GetFuturesLeadTradingSymbolWhitelistResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []GetFuturesLeadTradingSymbolWhitelistResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) GetDataOk() ([]GetFuturesLeadTradingSymbolWhitelistResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetFuturesLeadTradingSymbolWhitelistResponseDataInner and assigns it to the Data field.
func (o *GetFuturesLeadTradingSymbolWhitelistResponse) SetData(v []GetFuturesLeadTradingSymbolWhitelistResponseDataInner) {
	o.Data = v
}

func (o GetFuturesLeadTradingSymbolWhitelistResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFuturesLeadTradingSymbolWhitelistResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetFuturesLeadTradingSymbolWhitelistResponse) UnmarshalJSON(data []byte) (err error) {
	varGetFuturesLeadTradingSymbolWhitelistResponse := _GetFuturesLeadTradingSymbolWhitelistResponse{}

	err = json.Unmarshal(data, &varGetFuturesLeadTradingSymbolWhitelistResponse)

	if err != nil {
		return err
	}

	*o = GetFuturesLeadTradingSymbolWhitelistResponse(varGetFuturesLeadTradingSymbolWhitelistResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFuturesLeadTradingSymbolWhitelistResponse struct {
	value *GetFuturesLeadTradingSymbolWhitelistResponse
	isSet bool
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponse) Get() *GetFuturesLeadTradingSymbolWhitelistResponse {
	return v.value
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponse) Set(val *GetFuturesLeadTradingSymbolWhitelistResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFuturesLeadTradingSymbolWhitelistResponse(val *GetFuturesLeadTradingSymbolWhitelistResponse) *NullableGetFuturesLeadTradingSymbolWhitelistResponse {
	return &NullableGetFuturesLeadTradingSymbolWhitelistResponse{value: val, isSet: true}
}

func (v NullableGetFuturesLeadTradingSymbolWhitelistResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFuturesLeadTradingSymbolWhitelistResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
