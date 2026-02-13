/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryUniversalTransferHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryUniversalTransferHistoryResponse{}

// QueryUniversalTransferHistoryResponse struct for QueryUniversalTransferHistoryResponse
type QueryUniversalTransferHistoryResponse struct {
	Result               []QueryUniversalTransferHistoryResponseResultInner `json:"result,omitempty"`
	TotalCount           *int64                                             `json:"totalCount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QueryUniversalTransferHistoryResponse QueryUniversalTransferHistoryResponse

// NewQueryUniversalTransferHistoryResponse instantiates a new QueryUniversalTransferHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryUniversalTransferHistoryResponse() *QueryUniversalTransferHistoryResponse {
	this := QueryUniversalTransferHistoryResponse{}
	return &this
}

// NewQueryUniversalTransferHistoryResponseWithDefaults instantiates a new QueryUniversalTransferHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryUniversalTransferHistoryResponseWithDefaults() *QueryUniversalTransferHistoryResponse {
	this := QueryUniversalTransferHistoryResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponse) GetResult() []QueryUniversalTransferHistoryResponseResultInner {
	if o == nil || common.IsNil(o.Result) {
		var ret []QueryUniversalTransferHistoryResponseResultInner
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponse) GetResultOk() ([]QueryUniversalTransferHistoryResponseResultInner, bool) {
	if o == nil || common.IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponse) HasResult() bool {
	if o != nil && !common.IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []QueryUniversalTransferHistoryResponseResultInner and assigns it to the Result field.
func (o *QueryUniversalTransferHistoryResponse) SetResult(v []QueryUniversalTransferHistoryResponseResultInner) {
	o.Result = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QueryUniversalTransferHistoryResponse) GetTotalCount() int64 {
	if o == nil || common.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUniversalTransferHistoryResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QueryUniversalTransferHistoryResponse) HasTotalCount() bool {
	if o != nil && !common.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *QueryUniversalTransferHistoryResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o QueryUniversalTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryUniversalTransferHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !common.IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryUniversalTransferHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryUniversalTransferHistoryResponse := _QueryUniversalTransferHistoryResponse{}

	err = json.Unmarshal(data, &varQueryUniversalTransferHistoryResponse)

	if err != nil {
		return err
	}

	*o = QueryUniversalTransferHistoryResponse(varQueryUniversalTransferHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		delete(additionalProperties, "totalCount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryUniversalTransferHistoryResponse struct {
	value *QueryUniversalTransferHistoryResponse
	isSet bool
}

func (v NullableQueryUniversalTransferHistoryResponse) Get() *QueryUniversalTransferHistoryResponse {
	return v.value
}

func (v *NullableQueryUniversalTransferHistoryResponse) Set(val *QueryUniversalTransferHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryUniversalTransferHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryUniversalTransferHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryUniversalTransferHistoryResponse(val *QueryUniversalTransferHistoryResponse) *NullableQueryUniversalTransferHistoryResponse {
	return &NullableQueryUniversalTransferHistoryResponse{value: val, isSet: true}
}

func (v NullableQueryUniversalTransferHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryUniversalTransferHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
