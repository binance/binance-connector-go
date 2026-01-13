/*
Binance Mining REST API

OpenAPI Specification for the Binance Mining REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the AcquiringCoinnameResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &AcquiringCoinnameResponse{}

// AcquiringCoinnameResponse struct for AcquiringCoinnameResponse
type AcquiringCoinnameResponse struct {
	Code                 *int64                               `json:"code,omitempty"`
	Msg                  *string                              `json:"msg,omitempty"`
	Data                 []AcquiringCoinnameResponseDataInner `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AcquiringCoinnameResponse AcquiringCoinnameResponse

// NewAcquiringCoinnameResponse instantiates a new AcquiringCoinnameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcquiringCoinnameResponse() *AcquiringCoinnameResponse {
	this := AcquiringCoinnameResponse{}
	return &this
}

// NewAcquiringCoinnameResponseWithDefaults instantiates a new AcquiringCoinnameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcquiringCoinnameResponseWithDefaults() *AcquiringCoinnameResponse {
	this := AcquiringCoinnameResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponse) GetCode() int64 {
	if o == nil || common.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponse) GetCodeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *AcquiringCoinnameResponse) SetCode(v int64) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *AcquiringCoinnameResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AcquiringCoinnameResponse) GetData() []AcquiringCoinnameResponseDataInner {
	if o == nil || common.IsNil(o.Data) {
		var ret []AcquiringCoinnameResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcquiringCoinnameResponse) GetDataOk() ([]AcquiringCoinnameResponseDataInner, bool) {
	if o == nil || common.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AcquiringCoinnameResponse) HasData() bool {
	if o != nil && !common.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AcquiringCoinnameResponseDataInner and assigns it to the Data field.
func (o *AcquiringCoinnameResponse) SetData(v []AcquiringCoinnameResponseDataInner) {
	o.Data = v
}

func (o AcquiringCoinnameResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcquiringCoinnameResponse) ToMap() (map[string]interface{}, error) {
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

func (o *AcquiringCoinnameResponse) UnmarshalJSON(data []byte) (err error) {
	varAcquiringCoinnameResponse := _AcquiringCoinnameResponse{}

	err = json.Unmarshal(data, &varAcquiringCoinnameResponse)

	if err != nil {
		return err
	}

	*o = AcquiringCoinnameResponse(varAcquiringCoinnameResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAcquiringCoinnameResponse struct {
	value *AcquiringCoinnameResponse
	isSet bool
}

func (v NullableAcquiringCoinnameResponse) Get() *AcquiringCoinnameResponse {
	return v.value
}

func (v *NullableAcquiringCoinnameResponse) Set(val *AcquiringCoinnameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAcquiringCoinnameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAcquiringCoinnameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcquiringCoinnameResponse(val *AcquiringCoinnameResponse) *NullableAcquiringCoinnameResponse {
	return &NullableAcquiringCoinnameResponse{value: val, isSet: true}
}

func (v NullableAcquiringCoinnameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcquiringCoinnameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
