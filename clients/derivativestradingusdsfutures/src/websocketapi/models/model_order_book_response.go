/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderBookResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderBookResponse{}

// OrderBookResponse struct for OrderBookResponse
type OrderBookResponse struct {
	Id                   *string                            `json:"id,omitempty"`
	Status               *int64                             `json:"status,omitempty"`
	Result               *OrderBookResponseResult           `json:"result,omitempty"`
	RateLimits           []OrderBookResponseRateLimitsInner `json:"rateLimits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderBookResponse OrderBookResponse

// NewOrderBookResponse instantiates a new OrderBookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderBookResponse() *OrderBookResponse {
	this := OrderBookResponse{}
	return &this
}

// NewOrderBookResponseWithDefaults instantiates a new OrderBookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderBookResponseWithDefaults() *OrderBookResponse {
	this := OrderBookResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderBookResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderBookResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderBookResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderBookResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderBookResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderBookResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderBookResponse) GetResult() OrderBookResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderBookResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetResultOk() (*OrderBookResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderBookResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderBookResponseResult and assigns it to the Result field.
func (o *OrderBookResponse) SetResult(v OrderBookResponseResult) {
	o.Result = &v
}

// GetRateLimits returns the RateLimits field value if set, zero value otherwise.
func (o *OrderBookResponse) GetRateLimits() []OrderBookResponseRateLimitsInner {
	if o == nil || common.IsNil(o.RateLimits) {
		var ret []OrderBookResponseRateLimitsInner
		return ret
	}
	return o.RateLimits
}

// GetRateLimitsOk returns a tuple with the RateLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderBookResponse) GetRateLimitsOk() ([]OrderBookResponseRateLimitsInner, bool) {
	if o == nil || common.IsNil(o.RateLimits) {
		return nil, false
	}
	return o.RateLimits, true
}

// HasRateLimits returns a boolean if a field has been set.
func (o *OrderBookResponse) HasRateLimits() bool {
	if o != nil && !common.IsNil(o.RateLimits) {
		return true
	}

	return false
}

// SetRateLimits gets a reference to the given []OrderBookResponseRateLimitsInner and assigns it to the RateLimits field.
func (o *OrderBookResponse) SetRateLimits(v []OrderBookResponseRateLimitsInner) {
	o.RateLimits = v
}

func (o OrderBookResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderBookResponse) ToMap() (map[string]interface{}, error) {
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

func (o *OrderBookResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderBookResponse := _OrderBookResponse{}

	err = json.Unmarshal(data, &varOrderBookResponse)

	if err != nil {
		return err
	}

	*o = OrderBookResponse(varOrderBookResponse)

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

type NullableOrderBookResponse struct {
	value *OrderBookResponse
	isSet bool
}

func (v NullableOrderBookResponse) Get() *OrderBookResponse {
	return v.value
}

func (v *NullableOrderBookResponse) Set(val *OrderBookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderBookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderBookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderBookResponse(val *OrderBookResponse) *NullableOrderBookResponse {
	return &NullableOrderBookResponse{value: val, isSet: true}
}

func (v NullableOrderBookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderBookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
