/*
Binance Derivatives Trading USDS Futures WebSocket API

OpenAPI Specification for the Binance Derivatives Trading USDS Futures WebSocket API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryOrderResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryOrderResponse{}

// QueryOrderResponse struct for QueryOrderResponse
type QueryOrderResponse struct {
	Id                   *string                   `json:"id,omitempty"`
	Status               *int64                    `json:"status,omitempty"`
	Result               *QueryOrderResponseResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryOrderResponse QueryOrderResponse

// NewQueryOrderResponse instantiates a new QueryOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryOrderResponse() *QueryOrderResponse {
	this := QueryOrderResponse{}
	return &this
}

// NewQueryOrderResponseWithDefaults instantiates a new QueryOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryOrderResponseWithDefaults() *QueryOrderResponse {
	this := QueryOrderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QueryOrderResponse) GetId() string {
	if o == nil || common.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QueryOrderResponse) HasId() bool {
	if o != nil && !common.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *QueryOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *QueryOrderResponse) GetStatus() int64 {
	if o == nil || common.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponse) GetStatusOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *QueryOrderResponse) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *QueryOrderResponse) SetStatus(v int64) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *QueryOrderResponse) GetResult() QueryOrderResponseResult {
	if o == nil || common.IsNil(o.Result) {
		var ret QueryOrderResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryOrderResponse) GetResultOk() (*QueryOrderResponseResult, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *QueryOrderResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given QueryOrderResponseResult and assigns it to the Result field.
func (o *QueryOrderResponse) SetResult(v QueryOrderResponseResult) {
	o.Result = &v
}

func (o QueryOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryOrderResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryOrderResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryOrderResponse := _QueryOrderResponse{}

	err = json.Unmarshal(data, &varQueryOrderResponse)

	if err != nil {
		return err
	}

	*o = QueryOrderResponse(varQueryOrderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryOrderResponse struct {
	value *QueryOrderResponse
	isSet bool
}

func (v NullableQueryOrderResponse) Get() *QueryOrderResponse {
	return v.value
}

func (v *NullableQueryOrderResponse) Set(val *QueryOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryOrderResponse(val *QueryOrderResponse) *NullableQueryOrderResponse {
	return &NullableQueryOrderResponse{value: val, isSet: true}
}

func (v NullableQueryOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
