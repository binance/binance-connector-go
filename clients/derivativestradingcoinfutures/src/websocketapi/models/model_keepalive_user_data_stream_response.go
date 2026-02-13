/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the KeepaliveUserDataStreamResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &KeepaliveUserDataStreamResponse{}

// KeepaliveUserDataStreamResponse struct for KeepaliveUserDataStreamResponse
type KeepaliveUserDataStreamResponse struct {
	Id                   *string                                      `json:"id,omitempty"`
	Status               *int64                                       `json:"status,omitempty"`
	Result               *KeepaliveUserDataStreamResponseResult       `json:"result,omitempty"`
	RateLimits           []CloseUserDataStreamResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KeepaliveUserDataStreamResponse KeepaliveUserDataStreamResponse

// NewKeepaliveUserDataStreamResponse instantiates a new KeepaliveUserDataStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeepaliveUserDataStreamResponse() *KeepaliveUserDataStreamResponse {
	this := KeepaliveUserDataStreamResponse{}
	return &this
}

// NewKeepaliveUserDataStreamResponseWithDefaults instantiates a new KeepaliveUserDataStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeepaliveUserDataStreamResponseWithDefaults() *KeepaliveUserDataStreamResponse {
	this := KeepaliveUserDataStreamResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeepaliveUserDataStreamResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *KeepaliveUserDataStreamResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponse) GetResult() KeepaliveUserDataStreamResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret KeepaliveUserDataStreamResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponse) GetResultOk() (*KeepaliveUserDataStreamResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given KeepaliveUserDataStreamResponseResult and assigns it to the Result field.
func (o *KeepaliveUserDataStreamResponse) SetResult(v KeepaliveUserDataStreamResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *KeepaliveUserDataStreamResponse) GetRateLimits() []CloseUserDataStreamResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []CloseUserDataStreamResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeepaliveUserDataStreamResponse) GetRateLimitsOk() ([]CloseUserDataStreamResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *KeepaliveUserDataStreamResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []CloseUserDataStreamResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *KeepaliveUserDataStreamResponse) SetRateLimits(v []CloseUserDataStreamResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o KeepaliveUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeepaliveUserDataStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !common.IsNil(o.RateLimits) {
		toSerialize["rateLimits"] = o.RateLimits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *KeepaliveUserDataStreamResponse) UnmarshalJSON(data []byte) (err error) {
	varKeepaliveUserDataStreamResponse := _KeepaliveUserDataStreamResponse{}

	err = json.Unmarshal(data, &varKeepaliveUserDataStreamResponse)

	if err != nil {
		return err
	}

	*o = KeepaliveUserDataStreamResponse(varKeepaliveUserDataStreamResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		delete(additionalProperties, "rateLimits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKeepaliveUserDataStreamResponse struct {
	value *KeepaliveUserDataStreamResponse
	isSet bool
}

func (v NullableKeepaliveUserDataStreamResponse) Get() *KeepaliveUserDataStreamResponse {
	return v.value
}

func (v *NullableKeepaliveUserDataStreamResponse) Set(val *KeepaliveUserDataStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepaliveUserDataStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepaliveUserDataStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepaliveUserDataStreamResponse(val *KeepaliveUserDataStreamResponse) *NullableKeepaliveUserDataStreamResponse {
	return &NullableKeepaliveUserDataStreamResponse{value: val, isSet: true}
}

func (v NullableKeepaliveUserDataStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepaliveUserDataStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
