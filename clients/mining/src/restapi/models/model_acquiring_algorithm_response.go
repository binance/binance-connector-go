/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the AcquiringAlgorithmResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcquiringAlgorithmResponse{}

// AcquiringAlgorithmResponse struct for AcquiringAlgorithmResponse
type AcquiringAlgorithmResponse struct {
	Code                 *int64                                `json:"code,omitempty"`
	Msg                  *string                               `json:"msg,omitempty"`
	Data                 []AcquiringAlgorithmResponseDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcquiringAlgorithmResponse AcquiringAlgorithmResponse

// NewAcquiringAlgorithmResponse instantiates a new AcquiringAlgorithmResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquiringAlgorithmResponse() *AcquiringAlgorithmResponse {
	this := AcquiringAlgorithmResponse{}
	return &this
}

// NewAcquiringAlgorithmResponseWithDefaults instantiates a new AcquiringAlgorithmResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquiringAlgorithmResponseWithDefaults() *AcquiringAlgorithmResponse {
	this := AcquiringAlgorithmResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *AcquiringAlgorithmResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AcquiringAlgorithmResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AcquiringAlgorithmResponse) GetData() []AcquiringAlgorithmResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []AcquiringAlgorithmResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringAlgorithmResponse) GetDataOk() ([]AcquiringAlgorithmResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AcquiringAlgorithmResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AcquiringAlgorithmResponseDataInner and assigns it to the Data field.
func (o *AcquiringAlgorithmResponse) SetData(v []AcquiringAlgorithmResponseDataInner) {
	o.Data = v
}

func (o AcquiringAlgorithmResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcquiringAlgorithmResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !common.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AcquiringAlgorithmResponse) UnmarshalJSON(data []byte) (err error) {
	varAcquiringAlgorithmResponse := _AcquiringAlgorithmResponse{}

	err = json.Unmarshal(data, &varAcquiringAlgorithmResponse)

	if err != nil {
		return err
	}

	*o = AcquiringAlgorithmResponse(varAcquiringAlgorithmResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcquiringAlgorithmResponse struct {
	value *AcquiringAlgorithmResponse
	isSet bool
}

func (v NullableAcquiringAlgorithmResponse) Get() *AcquiringAlgorithmResponse {
	return v.value
}

func (v *NullableAcquiringAlgorithmResponse) Set(val *AcquiringAlgorithmResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquiringAlgorithmResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquiringAlgorithmResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquiringAlgorithmResponse(val *AcquiringAlgorithmResponse) *NullableAcquiringAlgorithmResponse {
	return &NullableAcquiringAlgorithmResponse{value: val, isSet: true}
}

func (v NullableAcquiringAlgorithmResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquiringAlgorithmResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
