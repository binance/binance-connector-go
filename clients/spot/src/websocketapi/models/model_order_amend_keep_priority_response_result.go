/*
Binance Spot WebSocket API

OpenAPI Specifications for the Binance Spot WebSocket API  API documents:   - [Github web-socket-api documentation file](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-api.md)   - [General API information for web-socket-api on website](https://developers.binance.com/docs/binance-spot-api-docs/web-socket-api/general-api-information)
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the OrderAmendKeepPriorityResponseResult type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &OrderAmendKeepPriorityResponseResult{}

// OrderAmendKeepPriorityResponseResult struct for OrderAmendKeepPriorityResponseResult
type OrderAmendKeepPriorityResponseResult struct {
	TransactTime         *int64                                            `json:"transactTime,omitempty"`
	ExecutionId          *int64                                            `json:"executionId,omitempty"`
	AmendedOrder         *OrderAmendKeepPriorityResponseResultAmendedOrder `json:"amendedOrder,omitempty"`
	ListStatus           *OrderAmendKeepPriorityResponseResultListStatus   `json:"listStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderAmendKeepPriorityResponseResult OrderAmendKeepPriorityResponseResult

// NewOrderAmendKeepPriorityResponseResult instantiates a new OrderAmendKeepPriorityResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmendKeepPriorityResponseResult() *OrderAmendKeepPriorityResponseResult {
	this := OrderAmendKeepPriorityResponseResult{}
	return &this
}

// NewOrderAmendKeepPriorityResponseResultWithDefaults instantiates a new OrderAmendKeepPriorityResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmendKeepPriorityResponseResultWithDefaults() *OrderAmendKeepPriorityResponseResult {
	this := OrderAmendKeepPriorityResponseResult{}
	return &this
}

// GetTransactTime returns the TransactTime field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResult) GetTransactTime() int64 {
	if o == nil || common.IsNil(o.TransactTime) {
		var ret int64
		return ret
	}
	return *o.TransactTime
}

// GetTransactTimeOk returns a tuple with the TransactTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResult) GetTransactTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TransactTime) {
		return nil, false
	}
	return o.TransactTime, true
}

// HasTransactTime returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResult) HasTransactTime() bool {
	if o != nil && !common.IsNil(o.TransactTime) {
		return true
	}

	return false
}

// SetTransactTime gets a reference to the given int64 and assigns it to the TransactTime field.
func (o *OrderAmendKeepPriorityResponseResult) SetTransactTime(v int64) {
	o.TransactTime = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResult) GetExecutionId() int64 {
	if o == nil || common.IsNil(o.ExecutionId) {
		var ret int64
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResult) GetExecutionIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResult) HasExecutionId() bool {
	if o != nil && !common.IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given int64 and assigns it to the ExecutionId field.
func (o *OrderAmendKeepPriorityResponseResult) SetExecutionId(v int64) {
	o.ExecutionId = &v
}

// GetAmendedOrder returns the AmendedOrder field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResult) GetAmendedOrder() OrderAmendKeepPriorityResponseResultAmendedOrder {
	if o == nil || common.IsNil(o.AmendedOrder) {
		var ret OrderAmendKeepPriorityResponseResultAmendedOrder
		return ret
	}
	return *o.AmendedOrder
}

// GetAmendedOrderOk returns a tuple with the AmendedOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResult) GetAmendedOrderOk() (*OrderAmendKeepPriorityResponseResultAmendedOrder, bool) {
	if o == nil || common.IsNil(o.AmendedOrder) {
		return nil, false
	}
	return o.AmendedOrder, true
}

// HasAmendedOrder returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResult) HasAmendedOrder() bool {
	if o != nil && !common.IsNil(o.AmendedOrder) {
		return true
	}

	return false
}

// SetAmendedOrder gets a reference to the given OrderAmendKeepPriorityResponseResultAmendedOrder and assigns it to the AmendedOrder field.
func (o *OrderAmendKeepPriorityResponseResult) SetAmendedOrder(v OrderAmendKeepPriorityResponseResultAmendedOrder) {
	o.AmendedOrder = &v
}

// GetListStatus returns the ListStatus field value if set, zero value otherwise.
func (o *OrderAmendKeepPriorityResponseResult) GetListStatus() OrderAmendKeepPriorityResponseResultListStatus {
	if o == nil || common.IsNil(o.ListStatus) {
		var ret OrderAmendKeepPriorityResponseResultListStatus
		return ret
	}
	return *o.ListStatus
}

// GetListStatusOk returns a tuple with the ListStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAmendKeepPriorityResponseResult) GetListStatusOk() (*OrderAmendKeepPriorityResponseResultListStatus, bool) {
	if o == nil || common.IsNil(o.ListStatus) {
		return nil, false
	}
	return o.ListStatus, true
}

// HasListStatus returns a boolean if a field has been set.
func (o *OrderAmendKeepPriorityResponseResult) HasListStatus() bool {
	if o != nil && !common.IsNil(o.ListStatus) {
		return true
	}

	return false
}

// SetListStatus gets a reference to the given OrderAmendKeepPriorityResponseResultListStatus and assigns it to the ListStatus field.
func (o *OrderAmendKeepPriorityResponseResult) SetListStatus(v OrderAmendKeepPriorityResponseResultListStatus) {
	o.ListStatus = &v
}

func (o OrderAmendKeepPriorityResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderAmendKeepPriorityResponseResult) ToMap() (map[string]interface{}, error) {
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

func (o *OrderAmendKeepPriorityResponseResult) UnmarshalJSON(data []byte) (err error) {
	varOrderAmendKeepPriorityResponseResult := _OrderAmendKeepPriorityResponseResult{}

	err = json.Unmarshal(data, &varOrderAmendKeepPriorityResponseResult)

	if err != nil {
		return err
	}

	*o = OrderAmendKeepPriorityResponseResult(varOrderAmendKeepPriorityResponseResult)

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

type NullableOrderAmendKeepPriorityResponseResult struct {
	value *OrderAmendKeepPriorityResponseResult
	isSet bool
}

func (v NullableOrderAmendKeepPriorityResponseResult) Get() *OrderAmendKeepPriorityResponseResult {
	return v.value
}

func (v *NullableOrderAmendKeepPriorityResponseResult) Set(val *OrderAmendKeepPriorityResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmendKeepPriorityResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmendKeepPriorityResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmendKeepPriorityResponseResult(val *OrderAmendKeepPriorityResponseResult) *NullableOrderAmendKeepPriorityResponseResult {
	return &NullableOrderAmendKeepPriorityResponseResult{value: val, isSet: true}
}

func (v NullableOrderAmendKeepPriorityResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmendKeepPriorityResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
