/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the NewOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &NewOrderResponse{}

// NewOrderResponse struct for NewOrderResponse
type NewOrderResponse struct {
	Id                   *string                              `json:"id,omitempty"`
	Status               *int64                               `json:"status,omitempty"`
	Result               *NewOrderResponseResult              `json:"result,omitempty"`
	RateLimits           []ModifyOrderResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NewOrderResponse NewOrderResponse

// NewNewOrderResponse instantiates a new NewOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewOrderResponse() *NewOrderResponse {
	this := NewOrderResponse{}
	return &this
}

// NewNewOrderResponseWithDefaults instantiates a new NewOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewOrderResponseWithDefaults() *NewOrderResponse {
	this := NewOrderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NewOrderResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NewOrderResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NewOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NewOrderResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewOrderResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NewOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *NewOrderResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *NewOrderResponse) GetResult() NewOrderResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret NewOrderResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewOrderResponse) GetResultOk() (*NewOrderResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *NewOrderResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given NewOrderResponseResult and assigns it to the Result field.
func (o *NewOrderResponse) SetResult(v NewOrderResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *NewOrderResponse) GetRateLimits() []ModifyOrderResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []ModifyOrderResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewOrderResponse) GetRateLimitsOk() ([]ModifyOrderResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *NewOrderResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []ModifyOrderResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *NewOrderResponse) SetRateLimits(v []ModifyOrderResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o NewOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *NewOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varNewOrderResponse := _NewOrderResponse{}

	err = json.Unmarshal(data, &varNewOrderResponse)

	if err != nil {
		return err
	}

	*o = NewOrderResponse(varNewOrderResponse)

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

type NullableNewOrderResponse struct {
	value *NewOrderResponse
	isSet bool
}

func (v NullableNewOrderResponse) Get() *NewOrderResponse {
	return v.value
}

func (v *NullableNewOrderResponse) Set(val *NewOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNewOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNewOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewOrderResponse(val *NewOrderResponse) *NullableNewOrderResponse {
	return &NullableNewOrderResponse{value: val, isSet: true}
}

func (v NullableNewOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
