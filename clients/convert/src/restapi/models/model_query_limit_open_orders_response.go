/*
Binance Convert REST API

OpenAPI Specification for the Binance Convert REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryLimitOpenOrdersResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryLimitOpenOrdersResponse{}

// QueryLimitOpenOrdersResponse struct for QueryLimitOpenOrdersResponse
type QueryLimitOpenOrdersResponse struct {
	List                 []QueryLimitOpenOrdersResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryLimitOpenOrdersResponse QueryLimitOpenOrdersResponse

// NewQueryLimitOpenOrdersResponse instantiates a new QueryLimitOpenOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryLimitOpenOrdersResponse() *QueryLimitOpenOrdersResponse {
	this := QueryLimitOpenOrdersResponse{}
	return &this
}

// NewQueryLimitOpenOrdersResponseWithDefaults instantiates a new QueryLimitOpenOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryLimitOpenOrdersResponseWithDefaults() *QueryLimitOpenOrdersResponse {
	this := QueryLimitOpenOrdersResponse{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *QueryLimitOpenOrdersResponse) GetList() []QueryLimitOpenOrdersResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []QueryLimitOpenOrdersResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryLimitOpenOrdersResponse) GetListOk() ([]QueryLimitOpenOrdersResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *QueryLimitOpenOrdersResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []QueryLimitOpenOrdersResponseListInner and assigns it to the List field.
func (o *QueryLimitOpenOrdersResponse) SetList(v []QueryLimitOpenOrdersResponseListInner) {
	o.List = v
}

func (o QueryLimitOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryLimitOpenOrdersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.List) {
		toSerialize["list"] = o.List
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryLimitOpenOrdersResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryLimitOpenOrdersResponse := _QueryLimitOpenOrdersResponse{}

	err = json.Unmarshal(data, &varQueryLimitOpenOrdersResponse)

	if err != nil {
		return err
	}

	*o = QueryLimitOpenOrdersResponse(varQueryLimitOpenOrdersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryLimitOpenOrdersResponse struct {
	value *QueryLimitOpenOrdersResponse
	isSet bool
}

func (v NullableQueryLimitOpenOrdersResponse) Get() *QueryLimitOpenOrdersResponse {
	return v.value
}

func (v *NullableQueryLimitOpenOrdersResponse) Set(val *QueryLimitOpenOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryLimitOpenOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryLimitOpenOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryLimitOpenOrdersResponse(val *QueryLimitOpenOrdersResponse) *NullableQueryLimitOpenOrdersResponse {
	return &NullableQueryLimitOpenOrdersResponse{value: val, isSet: true}
}

func (v NullableQueryLimitOpenOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryLimitOpenOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
