/*
Binance Algo REST API

OpenAPI Specification for the Binance Algo REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryHistoricalAlgoOrdersSpotAlgoResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryHistoricalAlgoOrdersSpotAlgoResponse{}

// QueryHistoricalAlgoOrdersSpotAlgoResponse struct for QueryHistoricalAlgoOrdersSpotAlgoResponse
type QueryHistoricalAlgoOrdersSpotAlgoResponse struct {
	Total                *int64                                                 `json:"total,omitempty"`
	Orders               []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner `json:"orders,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryHistoricalAlgoOrdersSpotAlgoResponse QueryHistoricalAlgoOrdersSpotAlgoResponse

// NewQueryHistoricalAlgoOrdersSpotAlgoResponse instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryHistoricalAlgoOrdersSpotAlgoResponse() *QueryHistoricalAlgoOrdersSpotAlgoResponse {
	this := QueryHistoricalAlgoOrdersSpotAlgoResponse{}
	return &this
}

// NewQueryHistoricalAlgoOrdersSpotAlgoResponseWithDefaults instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryHistoricalAlgoOrdersSpotAlgoResponseWithDefaults() *QueryHistoricalAlgoOrdersSpotAlgoResponse {
	this := QueryHistoricalAlgoOrdersSpotAlgoResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetOrders() []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner {
	if o == nil || common.IsNil(o.Orders) {
		var ret []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetOrdersOk() ([]QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner, bool) {
	if o == nil || common.IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) HasOrders() bool {
	if o != nil && !common.IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner and assigns it to the Orders field.
func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) SetOrders(v []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner) {
	o.Orders = v
}

func (o QueryHistoricalAlgoOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryHistoricalAlgoOrdersSpotAlgoResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryHistoricalAlgoOrdersSpotAlgoResponse := _QueryHistoricalAlgoOrdersSpotAlgoResponse{}

	err = json.Unmarshal(data, &varQueryHistoricalAlgoOrdersSpotAlgoResponse)

	if err != nil {
		return err
	}

	*o = QueryHistoricalAlgoOrdersSpotAlgoResponse(varQueryHistoricalAlgoOrdersSpotAlgoResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "orders")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryHistoricalAlgoOrdersSpotAlgoResponse struct {
	value *QueryHistoricalAlgoOrdersSpotAlgoResponse
	isSet bool
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) Get() *QueryHistoricalAlgoOrdersSpotAlgoResponse {
	return v.value
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) Set(val *QueryHistoricalAlgoOrdersSpotAlgoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryHistoricalAlgoOrdersSpotAlgoResponse(val *QueryHistoricalAlgoOrdersSpotAlgoResponse) *NullableQueryHistoricalAlgoOrdersSpotAlgoResponse {
	return &NullableQueryHistoricalAlgoOrdersSpotAlgoResponse{value: val, isSet: true}
}

func (v NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryHistoricalAlgoOrdersSpotAlgoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
