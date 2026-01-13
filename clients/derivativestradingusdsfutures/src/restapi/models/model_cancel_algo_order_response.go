/*
Binance Derivatives Trading USDS Futures REST API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the CancelAlgoOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelAlgoOrderResponse{}

// CancelAlgoOrderResponse struct for CancelAlgoOrderResponse
type CancelAlgoOrderResponse struct {
	AlgoId               *int64  `json:"algoId,omitempty"`
	ClientAlgoId         *string `json:"clientAlgoId,omitempty"`
	Code                 *string `json:"code,omitempty"`
	Msg                  *string `json:"msg,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelAlgoOrderResponse CancelAlgoOrderResponse

// NewCancelAlgoOrderResponse instantiates a new CancelAlgoOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelAlgoOrderResponse() *CancelAlgoOrderResponse {
	this := CancelAlgoOrderResponse{}
	return &this
}

// NewCancelAlgoOrderResponseWithDefaults instantiates a new CancelAlgoOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelAlgoOrderResponseWithDefaults() *CancelAlgoOrderResponse {
	this := CancelAlgoOrderResponse{}
	return &this
}

// GetAlgoId returns the AlgoId field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponse) GetAlgoId() int64 {
	if o == nil || common.IsNil(o.AlgoId) {
		var ret int64
		return ret
	}
	return *o.AlgoId
}

// GetAlgoIdOk returns a tuple with the AlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponse) GetAlgoIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.AlgoId) {
		return nil, false
	}
	return o.AlgoId, true
}

// HasAlgoId returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponse) HasAlgoId() bool {
	if o != nil && !common.IsNil(o.AlgoId) {
		return true
	}

	return false
}

// SetAlgoId gets a reference to the given int64 and assigns it to the AlgoId field.
func (o *CancelAlgoOrderResponse) SetAlgoId(v int64) {
	o.AlgoId = &v
}

// GetClientAlgoId returns the ClientAlgoId field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponse) GetClientAlgoId() string {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		var ret string
		return ret
	}
	return *o.ClientAlgoId
}

// GetClientAlgoIdOk returns a tuple with the ClientAlgoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ClientAlgoId) {
		return nil, false
	}
	return o.ClientAlgoId, true
}

// HasClientAlgoId returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponse) HasClientAlgoId() bool {
	if o != nil && !common.IsNil(o.ClientAlgoId) {
		return true
	}

	return false
}

// SetClientAlgoId gets a reference to the given string and assigns it to the ClientAlgoId field.
func (o *CancelAlgoOrderResponse) SetClientAlgoId(v string) {
	o.ClientAlgoId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponse) GetCode() string {
	if o == nil || common.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponse) GetCodeOk() (*string, bool) {
	if o == nil || common.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponse) HasCode() bool {
	if o != nil && !common.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CancelAlgoOrderResponse) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CancelAlgoOrderResponse) GetMsg() string {
	if o == nil || common.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelAlgoOrderResponse) GetMsgOk() (*string, bool) {
	if o == nil || common.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CancelAlgoOrderResponse) HasMsg() bool {
	if o != nil && !common.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CancelAlgoOrderResponse) SetMsg(v string) {
	o.Msg = &v
}

func (o CancelAlgoOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelAlgoOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.AlgoId) {
		toSerialize["algoId"] = o.AlgoId
	}
	if !common.IsNil(o.ClientAlgoId) {
		toSerialize["clientAlgoId"] = o.ClientAlgoId
	}
	if !common.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !common.IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CancelAlgoOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelAlgoOrderResponse := _CancelAlgoOrderResponse{}

	err = json.Unmarshal(data, &varCancelAlgoOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelAlgoOrderResponse(varCancelAlgoOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "algoId")
		delete(additionalProperties, "clientAlgoId")
		delete(additionalProperties, "code")
		delete(additionalProperties, "msg")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelAlgoOrderResponse struct {
	value *CancelAlgoOrderResponse
	isSet bool
}

func (v NullableCancelAlgoOrderResponse) Get() *CancelAlgoOrderResponse {
	return v.value
}

func (v *NullableCancelAlgoOrderResponse) Set(val *CancelAlgoOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelAlgoOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelAlgoOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelAlgoOrderResponse(val *CancelAlgoOrderResponse) *NullableCancelAlgoOrderResponse {
	return &NullableCancelAlgoOrderResponse{value: val, isSet: true}
}

func (v NullableCancelAlgoOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelAlgoOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
