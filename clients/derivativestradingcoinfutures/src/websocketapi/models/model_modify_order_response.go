/*
Binance Derivatives Trading COIN Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading COIN Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the ModifyOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &ModifyOrderResponse{}

// ModifyOrderResponse struct for ModifyOrderResponse
type ModifyOrderResponse struct {
	Id                   *string                              `json:"id,omitempty"`
	Status               *int64                               `json:"status,omitempty"`
	Result               *ModifyOrderResponseResult           `json:"result,omitempty"`
	RateLimits           []CancelOrderResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModifyOrderResponse ModifyOrderResponse

// NewModifyOrderResponse instantiates a new ModifyOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyOrderResponse() *ModifyOrderResponse {
	this := ModifyOrderResponse{}
	return &this
}

// NewModifyOrderResponseWithDefaults instantiates a new ModifyOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyOrderResponseWithDefaults() *ModifyOrderResponse {
	this := ModifyOrderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModifyOrderResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModifyOrderResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModifyOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModifyOrderResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyOrderResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModifyOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ModifyOrderResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ModifyOrderResponse) GetResult() ModifyOrderResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret ModifyOrderResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyOrderResponse) GetResultOk() (*ModifyOrderResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ModifyOrderResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ModifyOrderResponseResult and assigns it to the Result field.
func (o *ModifyOrderResponse) SetResult(v ModifyOrderResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *ModifyOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []CancelOrderResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyOrderResponse) GetRateLimitsOk() ([]CancelOrderResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *ModifyOrderResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []CancelOrderResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *ModifyOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o ModifyOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *ModifyOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varModifyOrderResponse := _ModifyOrderResponse{}

	err = json.Unmarshal(data, &varModifyOrderResponse)

	if err != nil {
		return err
	}

	*o = ModifyOrderResponse(varModifyOrderResponse)

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

type NullableModifyOrderResponse struct {
	value *ModifyOrderResponse
	isSet bool
}

func (v NullableModifyOrderResponse) Get() *ModifyOrderResponse {
	return v.value
}

func (v *NullableModifyOrderResponse) Set(val *ModifyOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyOrderResponse(val *ModifyOrderResponse) *NullableModifyOrderResponse {
	return &NullableModifyOrderResponse{value: val, isSet: true}
}

func (v NullableModifyOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
