/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentAlgoOpenOrdersFutureAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentAlgoOpenOrdersFutureAlgoResponse{}

// QueryCurrentAlgoOpenOrdersFutureAlgoResponse struct for QueryCurrentAlgoOpenOrdersFutureAlgoResponse
type QueryCurrentAlgoOpenOrdersFutureAlgoResponse struct {
	Total                *int64                                                    `json:"total,omitempty"`
	Orders               []QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCurrentAlgoOpenOrdersFutureAlgoResponse QueryCurrentAlgoOpenOrdersFutureAlgoResponse

// NewQueryCurrentAlgoOpenOrdersFutureAlgoResponse instantiates a new QueryCurrentAlgoOpenOrdersFutureAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentAlgoOpenOrdersFutureAlgoResponse() *QueryCurrentAlgoOpenOrdersFutureAlgoResponse {
	this := QueryCurrentAlgoOpenOrdersFutureAlgoResponse{}
	return &this
}

// NewQueryCurrentAlgoOpenOrdersFutureAlgoResponseWithDefaults instantiates a new QueryCurrentAlgoOpenOrdersFutureAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentAlgoOpenOrdersFutureAlgoResponseWithDefaults() *QueryCurrentAlgoOpenOrdersFutureAlgoResponse {
	this := QueryCurrentAlgoOpenOrdersFutureAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) GetOrders() []QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) GetOrdersOk() ([]QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner and assigns it to the Orders field.
func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) SetOrders(v []QueryCurrentAlgoOpenOrdersFutureAlgoResponseOrdersInner) {
	o.Orders = v
}

func (o QueryCurrentAlgoOpenOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentAlgoOpenOrdersFutureAlgoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryCurrentAlgoOpenOrdersFutureAlgoResponse := _QueryCurrentAlgoOpenOrdersFutureAlgoResponse{}

	err = json.Unmarshal(data, &varQueryCurrentAlgoOpenOrdersFutureAlgoResponse)

	if err != nil {
		return err
	}

	*o = QueryCurrentAlgoOpenOrdersFutureAlgoResponse(varQueryCurrentAlgoOpenOrdersFutureAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse struct {
	value *QueryCurrentAlgoOpenOrdersFutureAlgoResponse
	isSet bool
}

func (v NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) Get() *QueryCurrentAlgoOpenOrdersFutureAlgoResponse {
	return v.value
}

func (v *NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) Set(val *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse(val *QueryCurrentAlgoOpenOrdersFutureAlgoResponse) *NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse {
	return &NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentAlgoOpenOrdersFutureAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
