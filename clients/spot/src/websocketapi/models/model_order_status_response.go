/*
Spot WebSocket API

Access market data, manage accounts, and trade on Binance Spot.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderStatusResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderStatusResponse{}

// OrderStatusResponse struct for OrderStatusResponse
type OrderStatusResponse struct {
	Id                   *string                                    `json:"id,omitempty"`
	Status               *int64                                     `json:"status,omitempty"`
	Result               *OrderStatusResponseResult                 `json:"result,omitempty"`
	RateLimits           []AccountCommissionResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderStatusResponse OrderStatusResponse

// NewOrderStatusResponse instantiates a new OrderStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusResponse() *OrderStatusResponse {
	this := OrderStatusResponse{}
	return &this
}

// NewOrderStatusResponseWithDefaults instantiates a new OrderStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusResponseWithDefaults() *OrderStatusResponse {
	this := OrderStatusResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderStatusResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderStatusResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetResult() OrderStatusResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderStatusResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetResultOk() (*OrderStatusResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderStatusResponseResult and assigns it to the Result field.
func (o *OrderStatusResponse) SetResult(v OrderStatusResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *OrderStatusResponse) GetRateLimits() []AccountCommissionResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []AccountCommissionResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusResponse) GetRateLimitsOk() ([]AccountCommissionResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *OrderStatusResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []AccountCommissionResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *OrderStatusResponse) SetRateLimits(v []AccountCommissionResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o OrderStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderStatusResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderStatusResponse := _OrderStatusResponse{}

	err = json.Unmarshal(data, &varOrderStatusResponse)

	if err != nil {
		return err
	}

	*o = OrderStatusResponse(varOrderStatusResponse)

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

type NullableOrderStatusResponse struct {
	value *OrderStatusResponse
	isSet bool
}

func (v NullableOrderStatusResponse) Get() *OrderStatusResponse {
	return v.value
}

func (v *NullableOrderStatusResponse) Set(val *OrderStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusResponse(val *OrderStatusResponse) *NullableOrderStatusResponse {
	return &NullableOrderStatusResponse{value: val, isSet: true}
}

func (v NullableOrderStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
