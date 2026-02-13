/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the CancelOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &CancelOrderResponse{}

// CancelOrderResponse struct for CancelOrderResponse
type CancelOrderResponse struct {
	Id                   *string                              `json:"id,omitempty"`
	Status               *int64                               `json:"status,omitempty"`
	Result               *CancelOrderResponseResult           `json:"result,omitempty"`
	RateLimits           []CancelOrderResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelOrderResponse CancelOrderResponse

// NewCancelOrderResponse instantiates a new CancelOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderResponse() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderResponseWithDefaults() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CancelOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *CancelOrderResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetResult() CancelOrderResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret CancelOrderResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetResultOk() (*CancelOrderResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CancelOrderResponseResult and assigns it to the Result field.
func (o *CancelOrderResponse) SetResult(v CancelOrderResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *CancelOrderResponse) GetRateLimits() []CancelOrderResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []CancelOrderResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetRateLimitsOk() ([]CancelOrderResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *CancelOrderResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []CancelOrderResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *CancelOrderResponse) SetRateLimits(v []CancelOrderResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o CancelOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CancelOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varCancelOrderResponse := _CancelOrderResponse{}

	err = json.Unmarshal(data, &varCancelOrderResponse)

	if err != nil {
		return err
	}

	*o = CancelOrderResponse(varCancelOrderResponse)

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

type NullableCancelOrderResponse struct {
	value *CancelOrderResponse
	isSet bool
}

func (v NullableCancelOrderResponse) Get() *CancelOrderResponse {
	return v.value
}

func (v *NullableCancelOrderResponse) Set(val *CancelOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderResponse(val *CancelOrderResponse) *NullableCancelOrderResponse {
	return &NullableCancelOrderResponse{value: val, isSet: true}
}

func (v NullableCancelOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
