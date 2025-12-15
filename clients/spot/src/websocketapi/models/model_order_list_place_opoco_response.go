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

// checks if the OrderListPlaceOpocoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderListPlaceOpocoResponse{}

// OrderListPlaceOpocoResponse struct for OrderListPlaceOpocoResponse
type OrderListPlaceOpocoResponse struct {
	Id                   *string                            `json:"id,omitempty"`
	Status               *int64                             `json:"status,omitempty"`
	Result               *OrderListPlaceOpocoResponseResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderListPlaceOpocoResponse OrderListPlaceOpocoResponse

// NewOrderListPlaceOpocoResponse instantiates a new OrderListPlaceOpocoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderListPlaceOpocoResponse() *OrderListPlaceOpocoResponse {
	this := OrderListPlaceOpocoResponse{}
	return &this
}

// NewOrderListPlaceOpocoResponseWithDefaults instantiates a new OrderListPlaceOpocoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderListPlaceOpocoResponseWithDefaults() *OrderListPlaceOpocoResponse {
	this := OrderListPlaceOpocoResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderListPlaceOpocoResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *OrderListPlaceOpocoResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *OrderListPlaceOpocoResponse) GetResult() OrderListPlaceOpocoResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret OrderListPlaceOpocoResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderListPlaceOpocoResponse) GetResultOk() (*OrderListPlaceOpocoResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OrderListPlaceOpocoResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given OrderListPlaceOpocoResponseResult and assigns it to the Result field.
func (o *OrderListPlaceOpocoResponse) SetResult(v OrderListPlaceOpocoResponseResult) {
	o.Result = &v
}

func (o OrderListPlaceOpocoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderListPlaceOpocoResponse) ToMap() (map[string]interface{}, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderListPlaceOpocoResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderListPlaceOpocoResponse := _OrderListPlaceOpocoResponse{}

	err = json.Unmarshal(data, &varOrderListPlaceOpocoResponse)

	if err != nil {
		return err
	}

	*o = OrderListPlaceOpocoResponse(varOrderListPlaceOpocoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderListPlaceOpocoResponse struct {
	value *OrderListPlaceOpocoResponse
	isSet bool
}

func (v NullableOrderListPlaceOpocoResponse) Get() *OrderListPlaceOpocoResponse {
	return v.value
}

func (v *NullableOrderListPlaceOpocoResponse) Set(val *OrderListPlaceOpocoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderListPlaceOpocoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderListPlaceOpocoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderListPlaceOpocoResponse(val *OrderListPlaceOpocoResponse) *NullableOrderListPlaceOpocoResponse {
	return &NullableOrderListPlaceOpocoResponse{value: val, isSet: true}
}

func (v NullableOrderListPlaceOpocoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderListPlaceOpocoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
