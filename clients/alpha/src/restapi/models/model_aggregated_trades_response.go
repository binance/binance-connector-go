/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API

*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AggregatedTradesResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AggregatedTradesResponse{}

// AggregatedTradesResponse struct for AggregatedTradesResponse
type AggregatedTradesResponse struct {
	Code                 *string                             `json:"code,omitempty"`
	Message              *string                             `json:"message,omitempty"`
	MessageDetail        *string                             `json:"messageDetail,omitempty"`
	Data                 []AggregatedTradesResponseDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AggregatedTradesResponse AggregatedTradesResponse

// NewAggregatedTradesResponse instantiates a new AggregatedTradesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregatedTradesResponse() *AggregatedTradesResponse {
	this := AggregatedTradesResponse{}
	return &this
}

// NewAggregatedTradesResponseWithDefaults instantiates a new AggregatedTradesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatedTradesResponseWithDefaults() *AggregatedTradesResponse {
	this := AggregatedTradesResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AggregatedTradesResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AggregatedTradesResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AggregatedTradesResponse) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AggregatedTradesResponse) GetMessage() string {
	if o == nil || common.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponse) GetMessageOk() (*string, bool) {
	if o == nil || common.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AggregatedTradesResponse) HasMessage() bool {
	if o != nil && !common.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AggregatedTradesResponse) SetMessage(v string) {
	o.Message = &v
}

// GetMessageDetail returns the MessageDetail field value if set, zero value otherwise.
func (o *AggregatedTradesResponse) GetMessageDetail() string {
	if o == nil || common.IsNil(o.MessageDetail) {
		var ret string
		return ret
	}
	return *o.MessageDetail
}

// GetMessageDetailOk returns a tuple with the MessageDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponse) GetMessageDetailOk() (*string, bool) {
	if o == nil || common.IsNil(o.MessageDetail) {
		return nil, false
	}
	return o.MessageDetail, true
}

// HasMessageDetail returns a boolean if a field has been set.
func (o *AggregatedTradesResponse) HasMessageDetail() bool {
	if o != nil && !common.IsNil(o.MessageDetail) {
		return true
	}

	return false
}

// SetMessageDetail gets a reference to the given string and assigns it to the MessageDetail field.
func (o *AggregatedTradesResponse) SetMessageDetail(v string) {
	o.MessageDetail = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AggregatedTradesResponse) GetData() []AggregatedTradesResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []AggregatedTradesResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregatedTradesResponse) GetDataOk() ([]AggregatedTradesResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AggregatedTradesResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AggregatedTradesResponseDataInner and assigns it to the Data field.
func (o *AggregatedTradesResponse) SetData(v []AggregatedTradesResponseDataInner) {
	o.Data = v
}

func (o AggregatedTradesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregatedTradesResponse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AggregatedTradesResponse) UnmarshalJSON(data []byte) (err error) {
	varAggregatedTradesResponse := _AggregatedTradesResponse{}

	err = json.Unmarshal(data, &varAggregatedTradesResponse)

	if err != nil {
		return err
	}

	*o = AggregatedTradesResponse(varAggregatedTradesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "messageDetail")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAggregatedTradesResponse struct {
	value *AggregatedTradesResponse
	isSet bool
}

func (v NullableAggregatedTradesResponse) Get() *AggregatedTradesResponse {
	return v.value
}

func (v *NullableAggregatedTradesResponse) Set(val *AggregatedTradesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregatedTradesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregatedTradesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregatedTradesResponse(val *AggregatedTradesResponse) *NullableAggregatedTradesResponse {
	return &NullableAggregatedTradesResponse{value: val, isSet: true}
}

func (v NullableAggregatedTradesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregatedTradesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
