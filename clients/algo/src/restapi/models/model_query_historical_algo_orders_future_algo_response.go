/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryHistoricalAlgoOrdersFutureAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryHistoricalAlgoOrdersFutureAlgoResponse{}

// QueryHistoricalAlgoOrdersFutureAlgoResponse struct for QueryHistoricalAlgoOrdersFutureAlgoResponse
type QueryHistoricalAlgoOrdersFutureAlgoResponse struct {
	Total                *int64                                                   `json:"total,omitempty"`
	Orders               []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryHistoricalAlgoOrdersFutureAlgoResponse QueryHistoricalAlgoOrdersFutureAlgoResponse

// NewQueryHistoricalAlgoOrdersFutureAlgoResponse instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryHistoricalAlgoOrdersFutureAlgoResponse() *QueryHistoricalAlgoOrdersFutureAlgoResponse {
	this := QueryHistoricalAlgoOrdersFutureAlgoResponse{}
	return &this
}

// NewQueryHistoricalAlgoOrdersFutureAlgoResponseWithDefaults instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryHistoricalAlgoOrdersFutureAlgoResponseWithDefaults() *QueryHistoricalAlgoOrdersFutureAlgoResponse {
	this := QueryHistoricalAlgoOrdersFutureAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetOrders() []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetOrdersOk() ([]QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner and assigns it to the Orders field.
func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) SetOrders(v []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner) {
	o.Orders = v
}

func (o QueryHistoricalAlgoOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryHistoricalAlgoOrdersFutureAlgoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryHistoricalAlgoOrdersFutureAlgoResponse := _QueryHistoricalAlgoOrdersFutureAlgoResponse{}

	err = json.Unmarshal(data, &varQueryHistoricalAlgoOrdersFutureAlgoResponse)

	if err != nil {
		return err
	}

	*o = QueryHistoricalAlgoOrdersFutureAlgoResponse(varQueryHistoricalAlgoOrdersFutureAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryHistoricalAlgoOrdersFutureAlgoResponse struct {
	value *QueryHistoricalAlgoOrdersFutureAlgoResponse
	isSet bool
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) Get() *QueryHistoricalAlgoOrdersFutureAlgoResponse {
	return v.value
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) Set(val *QueryHistoricalAlgoOrdersFutureAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryHistoricalAlgoOrdersFutureAlgoResponse(val *QueryHistoricalAlgoOrdersFutureAlgoResponse) *NullableQueryHistoricalAlgoOrdersFutureAlgoResponse {
	return &NullableQueryHistoricalAlgoOrdersFutureAlgoResponse{value: val, isSet: true}
}

func (v NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryHistoricalAlgoOrdersFutureAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
