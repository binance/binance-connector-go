/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QueryManagedSubAccountTransferLogSubAccountTradingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountTransferLogSubAccountTradingResponse{}

// QueryManagedSubAccountTransferLogSubAccountTradingResponse struct for QueryManagedSubAccountTransferLogSubAccountTradingResponse
type QueryManagedSubAccountTransferLogSubAccountTradingResponse struct {
	ManagerSubTransferHistoryVos []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner `json:"managerSubTransferHistoryVos,omitempty"`
	Count                        *int64                                                                                            `json:"count,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _QueryManagedSubAccountTransferLogSubAccountTradingResponse QueryManagedSubAccountTransferLogSubAccountTradingResponse

// NewQueryManagedSubAccountTransferLogSubAccountTradingResponse instantiates a new QueryManagedSubAccountTransferLogSubAccountTradingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountTransferLogSubAccountTradingResponse() *QueryManagedSubAccountTransferLogSubAccountTradingResponse {
	this := QueryManagedSubAccountTransferLogSubAccountTradingResponse{}
	return &this
}

// NewQueryManagedSubAccountTransferLogSubAccountTradingResponseWithDefaults instantiates a new QueryManagedSubAccountTransferLogSubAccountTradingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountTransferLogSubAccountTradingResponseWithDefaults() *QueryManagedSubAccountTransferLogSubAccountTradingResponse {
	this := QueryManagedSubAccountTransferLogSubAccountTradingResponse{}
	return &this
}

// GetManagerSubTransferHistoryVos returns the ManagerSubTransferHistoryVos field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) GetManagerSubTransferHistoryVos() []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		var ret []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner
		return ret
	}
	return o.ManagerSubTransferHistoryVos
}

// GetManagerSubTransferHistoryVosOk returns a tuple with the ManagerSubTransferHistoryVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) GetManagerSubTransferHistoryVosOk() ([]QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner, bool) {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		return nil, false
	}
	return o.ManagerSubTransferHistoryVos, true
}

// HasManagerSubTransferHistoryVos returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) HasManagerSubTransferHistoryVos() bool {
	if o != nil && !common.IsNil(o.ManagerSubTransferHistoryVos) {
		return true
	}

	return false
}

// SetManagerSubTransferHistoryVos gets a reference to the given []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner and assigns it to the ManagerSubTransferHistoryVos field.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) SetManagerSubTransferHistoryVos(v []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) {
	o.ManagerSubTransferHistoryVos = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) SetCount(v int64) {
	o.Count = &v
}

func (o QueryManagedSubAccountTransferLogSubAccountTradingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountTransferLogSubAccountTradingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.ManagerSubTransferHistoryVos) {
		toSerialize["managerSubTransferHistoryVos"] = o.ManagerSubTransferHistoryVos
	}
	if !common.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountTransferLogSubAccountTradingResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountTransferLogSubAccountTradingResponse := _QueryManagedSubAccountTransferLogSubAccountTradingResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountTransferLogSubAccountTradingResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountTransferLogSubAccountTradingResponse(varQueryManagedSubAccountTransferLogSubAccountTradingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managerSubTransferHistoryVos")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse struct {
	value *QueryManagedSubAccountTransferLogSubAccountTradingResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) Get() *QueryManagedSubAccountTransferLogSubAccountTradingResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) Set(val *QueryManagedSubAccountTransferLogSubAccountTradingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountTransferLogSubAccountTradingResponse(val *QueryManagedSubAccountTransferLogSubAccountTradingResponse) *NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse {
	return &NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountTransferLogSubAccountTradingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
