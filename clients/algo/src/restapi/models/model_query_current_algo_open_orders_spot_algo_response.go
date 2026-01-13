/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryCurrentAlgoOpenOrdersSpotAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryCurrentAlgoOpenOrdersSpotAlgoResponse{}

// QueryCurrentAlgoOpenOrdersSpotAlgoResponse struct for QueryCurrentAlgoOpenOrdersSpotAlgoResponse
type QueryCurrentAlgoOpenOrdersSpotAlgoResponse struct {
	Total                *int64                                                  `json:"total,omitempty"`
	Orders               []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryCurrentAlgoOpenOrdersSpotAlgoResponse QueryCurrentAlgoOpenOrdersSpotAlgoResponse

// NewQueryCurrentAlgoOpenOrdersSpotAlgoResponse instantiates a new QueryCurrentAlgoOpenOrdersSpotAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCurrentAlgoOpenOrdersSpotAlgoResponse() *QueryCurrentAlgoOpenOrdersSpotAlgoResponse {
	this := QueryCurrentAlgoOpenOrdersSpotAlgoResponse{}
	return &this
}

// NewQueryCurrentAlgoOpenOrdersSpotAlgoResponseWithDefaults instantiates a new QueryCurrentAlgoOpenOrdersSpotAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCurrentAlgoOpenOrdersSpotAlgoResponseWithDefaults() *QueryCurrentAlgoOpenOrdersSpotAlgoResponse {
	this := QueryCurrentAlgoOpenOrdersSpotAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetOrders() []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetOrdersOk() ([]QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner and assigns it to the Orders field.
func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) SetOrders(v []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner) {
	o.Orders = v
}

func (o QueryCurrentAlgoOpenOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCurrentAlgoOpenOrdersSpotAlgoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryCurrentAlgoOpenOrdersSpotAlgoResponse := _QueryCurrentAlgoOpenOrdersSpotAlgoResponse{}

	err = json.Unmarshal(data, &varQueryCurrentAlgoOpenOrdersSpotAlgoResponse)

	if err != nil {
		return err
	}

	*o = QueryCurrentAlgoOpenOrdersSpotAlgoResponse(varQueryCurrentAlgoOpenOrdersSpotAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse struct {
	value *QueryCurrentAlgoOpenOrdersSpotAlgoResponse
	isSet bool
}

func (v NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) Get() *QueryCurrentAlgoOpenOrdersSpotAlgoResponse {
	return v.value
}

func (v *NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) Set(val *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse(val *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) *NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse {
	return &NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse{value: val, isSet: true}
}

func (v NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCurrentAlgoOpenOrdersSpotAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
