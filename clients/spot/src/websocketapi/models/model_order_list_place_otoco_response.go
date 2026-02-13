/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderListPlaceOtocoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOtocoResponse{}

// OrderListPlaceOtocoResponse struct for OrderListPlaceOtocoResponse
type OrderListPlaceOtocoResponse struct {
	Id                   *string                            `json:"id,omitempty"`
	Status               *int64                             `json:"status,omitempty"`
	Result               *OrderListPlaceOtocoResponseResult `json:"result,omitempty"`
	RateLimits           []RateLimits                       `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOtocoResponse OrderListPlaceOtocoResponse

// NewOrderListPlaceOtocoResponse instantiates a new OrderListPlaceOtocoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOtocoResponse() *OrderListPlaceOtocoResponse {
	this := OrderListPlaceOtocoResponse{}
	return &this
}

// NewOrderListPlaceOtocoResponseWithDefaults instantiates a new OrderListPlaceOtocoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOtocoResponseWithDefaults() *OrderListPlaceOtocoResponse {
	this := OrderListPlaceOtocoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderListPlaceOtocoResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderListPlaceOtocoResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponse) GetResult() OrderListPlaceOtocoResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderListPlaceOtocoResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponse) GetResultOk() (*OrderListPlaceOtocoResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderListPlaceOtocoResponseResult and assigns it to the Result field.
func (o *OrderListPlaceOtocoResponse) SetResult(v OrderListPlaceOtocoResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *OrderListPlaceOtocoResponse) GetRateLimits() []RateLimits {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []RateLimits
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOtocoResponse) GetRateLimitsOk() ([]RateLimits, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *OrderListPlaceOtocoResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []RateLimits and assigns it to the RateLimits field.
func (o *OrderListPlaceOtocoResponse) SetRateLimits(v []RateLimits) {
	o.RateLimits = v
}

func (o OrderListPlaceOtocoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOtocoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderListPlaceOtocoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOtocoResponse := _OrderListPlaceOtocoResponse{}

	err = json.Unmarshal(data, &varOrderListPlaceOtocoResponse)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOtocoResponse(varOrderListPlaceOtocoResponse)

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

type NullableOrderListPlaceOtocoResponse struct {
	value *OrderListPlaceOtocoResponse
	isSet bool
}

func (v NullableOrderListPlaceOtocoResponse) Get() *OrderListPlaceOtocoResponse {
	return v.value
}

func (v *NullableOrderListPlaceOtocoResponse) Set(val *OrderListPlaceOtocoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOtocoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOtocoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOtocoResponse(val *OrderListPlaceOtocoResponse) *NullableOrderListPlaceOtocoResponse {
	return &NullableOrderListPlaceOtocoResponse{value: val, isSet: true}
}

func (v NullableOrderListPlaceOtocoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOtocoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
