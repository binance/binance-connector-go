/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderListPlaceOcoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOcoResponse{}

// OrderListPlaceOcoResponse struct for OrderListPlaceOcoResponse
type OrderListPlaceOcoResponse struct {
	Id                   *string                          `json:"id,omitempty"`
	Status               *int64                           `json:"status,omitempty"`
	Result               *OrderListPlaceOcoResponseResult `json:"result,omitempty"`
	RateLimits           []RateLimits                     `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOcoResponse OrderListPlaceOcoResponse

// NewOrderListPlaceOcoResponse instantiates a new OrderListPlaceOcoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOcoResponse() *OrderListPlaceOcoResponse {
	this := OrderListPlaceOcoResponse{}
	return &this
}

// NewOrderListPlaceOcoResponseWithDefaults instantiates a new OrderListPlaceOcoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOcoResponseWithDefaults() *OrderListPlaceOcoResponse {
	this := OrderListPlaceOcoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderListPlaceOcoResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderListPlaceOcoResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponse) GetResult() OrderListPlaceOcoResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderListPlaceOcoResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponse) GetResultOk() (*OrderListPlaceOcoResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderListPlaceOcoResponseResult and assigns it to the Result field.
func (o *OrderListPlaceOcoResponse) SetResult(v OrderListPlaceOcoResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *OrderListPlaceOcoResponse) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOcoResponse) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *OrderListPlaceOcoResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *OrderListPlaceOcoResponse) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

func (o OrderListPlaceOcoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOcoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOcoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOcoResponse := _OrderListPlaceOcoResponse{}

	err = json.Unmarshal(data, &varOrderListPlaceOcoResponse)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOcoResponse(varOrderListPlaceOcoResponse)

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

type NullableOrderListPlaceOcoResponse struct {
	value *OrderListPlaceOcoResponse
	isSet bool
}

func (v NullableOrderListPlaceOcoResponse) Get() *OrderListPlaceOcoResponse {
	return v.value
}

func (v *NullableOrderListPlaceOcoResponse) Set(val *OrderListPlaceOcoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOcoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOcoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOcoResponse(val *OrderListPlaceOcoResponse) *NullableOrderListPlaceOcoResponse {
	return &NullableOrderListPlaceOcoResponse{value: val, isSet: true}
}

func (v NullableOrderListPlaceOcoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOcoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
