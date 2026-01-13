/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderListPlaceOtoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOtoResponse{}

// OrderListPlaceOtoResponse struct for OrderListPlaceOtoResponse
type OrderListPlaceOtoResponse struct {
	Id                   *string                          `json:"id,omitempty"`
	Status               *int64                           `json:"status,omitempty"`
	Result               *OrderListPlaceOtoResponseResult `json:"result,omitempty"`
	RateLimits           []RateLimits                     `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOtoResponse OrderListPlaceOtoResponse

// NewOrderListPlaceOtoResponse instantiates a new OrderListPlaceOtoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOtoResponse() *OrderListPlaceOtoResponse {
	this := OrderListPlaceOtoResponse{}
	return &this
}

// NewOrderListPlaceOtoResponseWithDefaults instantiates a new OrderListPlaceOtoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOtoResponseWithDefaults() *OrderListPlaceOtoResponse {
	this := OrderListPlaceOtoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderListPlaceOtoResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderListPlaceOtoResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponse) GetResult() OrderListPlaceOtoResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderListPlaceOtoResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponse) GetResultOk() (*OrderListPlaceOtoResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderListPlaceOtoResponseResult and assigns it to the Result field.
func (o *OrderListPlaceOtoResponse) SetResult(v OrderListPlaceOtoResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *OrderListPlaceOtoResponse) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtoResponse) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *OrderListPlaceOtoResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *OrderListPlaceOtoResponse) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

func (o OrderListPlaceOtoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOtoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOtoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOtoResponse := _OrderListPlaceOtoResponse{}

	err = json.Unmarshal(data, &varOrderListPlaceOtoResponse)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOtoResponse(varOrderListPlaceOtoResponse)

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

type NullableOrderListPlaceOtoResponse struct {
	value *OrderListPlaceOtoResponse
	isSet bool
}

func (v NullableOrderListPlaceOtoResponse) Get() *OrderListPlaceOtoResponse {
	return v.value
}

func (v *NullableOrderListPlaceOtoResponse) Set(val *OrderListPlaceOtoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOtoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOtoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOtoResponse(val *OrderListPlaceOtoResponse) *NullableOrderListPlaceOtoResponse {
	return &NullableOrderListPlaceOtoResponse{value: val, isSet: true}
}

func (v NullableOrderListPlaceOtoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOtoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
