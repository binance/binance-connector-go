/*
Binance Spot REST API

OpenAPI Specifications for the Binance Spot REST API  API documents:   - [Github rest-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/rest-api.md)   - [General API information for rest-api on website](https://developers.binance.com/docs/binance-spot-api-docs/rest-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the OrderAmendKeepPriorityResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponse{}

// OrderAmendKeepPriorityResponse struct for OrderAmendKeepPriorityResponse
type OrderAmendKeepPriorityResponse struct {
	TransactTime         *int64                                      `json:"transactTime,omitempty"`
	ExecutionId          *int64                                      `json:"executionId,omitempty"`
	AmendedOrder         *OrderAmendKeepPriorityResponseAmendedOrder `json:"amendedOrder,omitempty"`
	ListStatus           *OrderAmendKeepPriorityResponseListStatus   `json:"listStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderAmendKeepPriorityResponse OrderAmendKeepPriorityResponse

// NewOrderAmendKeepPriorityResponse instantiates a new OrderAmendKeepPriorityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponse() *OrderAmendKeepPriorityResponse {
	this := OrderAmendKeepPriorityResponse{}
	return &this
}

// NewOrderAmendKeepPriorityResponseWithDefaults instantiates a new OrderAmendKeepPriorityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseWithDefaults() *OrderAmendKeepPriorityResponse {
	this := OrderAmendKeepPriorityResponse{}
	return &this
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponse) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponse) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponse) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *OrderAmendKeepPriorityResponse) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponse) GetExecutionId() int64 {
	if o == nil || common.IsNil(o.ExecutionId) {
		var ret int64
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponse) GetExecutionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponse) HasExecutionId() bool {
	if o != nil && !common.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given int64 and assigns it to the ExecutionId field.
func (o *OrderAmendKeepPriorityResponse) SetExecutionId(v int64) {
	o.ExecutionId = &v
}

// GetAmendedOrder returns the AmendedOrder field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponse) GetAmendedOrder() OrderAmendKeepPriorityResponseAmendedOrder {
	if o == nil || common.IsNil(o.AmendedOrder) {
		var ret OrderAmendKeepPriorityResponseAmendedOrder
		return ret
	}
	return *o.AmendedOrder
}

// GetAmendedOrderOk returns a tuple with the AmendedOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponse) GetAmendedOrderOk() (*OrderAmendKeepPriorityResponseAmendedOrder, bool) {
	if o == nil || common.IsNil(o.AmendedOrder) {
		return nil, false
	}
	return o.AmendedOrder, true
}

// HasAmendedOrder returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponse) HasAmendedOrder() bool {
	if o != nil && !common.IsNil(o.AmendedOrder) {
		return true
	}

	return false
}

// SetAmendedOrder gets a reference to the given OrderAmendKeepPriorityResponseAmendedOrder and assigns it to the AmendedOrder field.
func (o *OrderAmendKeepPriorityResponse) SetAmendedOrder(v OrderAmendKeepPriorityResponseAmendedOrder) {
	o.AmendedOrder = &v
}

// GetListStatus returns the ListStatus field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponse) GetListStatus() OrderAmendKeepPriorityResponseListStatus {
	if o == nil || common.IsNil(o.ListStatus) {
		var ret OrderAmendKeepPriorityResponseListStatus
		return ret
	}
	return *o.ListStatus
}

// GetListStatusOk returns a tuple with the ListStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponse) GetListStatusOk() (*OrderAmendKeepPriorityResponseListStatus, bool) {
	if o == nil || common.IsNil(o.ListStatus) {
		return nil, false
	}
	return o.ListStatus, true
}

// HasListStatus returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponse) HasListStatus() bool {
	if o != nil && !common.IsNil(o.ListStatus) {
		return true
	}

	return false
}

// SetListStatus gets a reference to the given OrderAmendKeepPriorityResponseListStatus and assigns it to the ListStatus field.
func (o *OrderAmendKeepPriorityResponse) SetListStatus(v OrderAmendKeepPriorityResponseListStatus) {
	o.ListStatus = &v
}

func (o OrderAmendKeepPriorityResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TransactTime) {
		toSerialize["transactTime"] = o.TransactTime
	}
	if !common.IsNil(o.ExecutionId) {
		toSerialize["executionId"] = o.ExecutionId
	}
	if !common.IsNil(o.AmendedOrder) {
		toSerialize["amendedOrder"] = o.AmendedOrder
	}
	if !common.IsNil(o.ListStatus) {
		toSerialize["listStatus"] = o.ListStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderAmendKeepPriorityResponse) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponse := _OrderAmendKeepPriorityResponse{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponse)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponse(varOrderAmendKeepPriorityResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "transactTime")
		delete(additionalProperties, "executionId")
		delete(additionalProperties, "amendedOrder")
		delete(additionalProperties, "listStatus")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderAmendKeepPriorityResponse struct {
	value *OrderAmendKeepPriorityResponse
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponse) Get() *OrderAmendKeepPriorityResponse {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponse) Set(val *OrderAmendKeepPriorityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponse(val *OrderAmendKeepPriorityResponse) *NullableOrderAmendKeepPriorityResponse {
	return &NullableOrderAmendKeepPriorityResponse{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
