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

// checks if the QueryManagedSubAccountTransferLogMasterAccountInvestorResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountTransferLogMasterAccountInvestorResponse{}

// QueryManagedSubAccountTransferLogMasterAccountInvestorResponse struct for QueryManagedSubAccountTransferLogMasterAccountInvestorResponse
type QueryManagedSubAccountTransferLogMasterAccountInvestorResponse struct {
	ManagerSubTransferHistoryVos []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner `json:"managerSubTransferHistoryVos,omitempty"`
	Count                        *int64                                                                                            `json:"count,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _QueryManagedSubAccountTransferLogMasterAccountInvestorResponse QueryManagedSubAccountTransferLogMasterAccountInvestorResponse

// NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponse instantiates a new QueryManagedSubAccountTransferLogMasterAccountInvestorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponse() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse {
	this := QueryManagedSubAccountTransferLogMasterAccountInvestorResponse{}
	return &this
}

// NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseWithDefaults instantiates a new QueryManagedSubAccountTransferLogMasterAccountInvestorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountTransferLogMasterAccountInvestorResponseWithDefaults() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse {
	this := QueryManagedSubAccountTransferLogMasterAccountInvestorResponse{}
	return &this
}

// GetManagerSubTransferHistoryVos returns the ManagerSubTransferHistoryVos field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) GetManagerSubTransferHistoryVos() []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		var ret []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner
		return ret
	}
	return o.ManagerSubTransferHistoryVos
}

// GetManagerSubTransferHistoryVosOk returns a tuple with the ManagerSubTransferHistoryVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) GetManagerSubTransferHistoryVosOk() ([]QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner, bool) {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		return nil, false
	}
	return o.ManagerSubTransferHistoryVos, true
}

// HasManagerSubTransferHistoryVos returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) HasManagerSubTransferHistoryVos() bool {
	if o != nil && !common.IsNil(o.ManagerSubTransferHistoryVos) {
		return true
	}

	return false
}

// SetManagerSubTransferHistoryVos gets a reference to the given []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner and assigns it to the ManagerSubTransferHistoryVos field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) SetManagerSubTransferHistoryVos(v []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) {
	o.ManagerSubTransferHistoryVos = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) SetCount(v int64) {
	o.Count = &v
}

func (o QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountTransferLogMasterAccountInvestorResponse := _QueryManagedSubAccountTransferLogMasterAccountInvestorResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountTransferLogMasterAccountInvestorResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountTransferLogMasterAccountInvestorResponse(varQueryManagedSubAccountTransferLogMasterAccountInvestorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managerSubTransferHistoryVos")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse struct {
	value *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) Get() *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) Set(val *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse(val *QueryManagedSubAccountTransferLogMasterAccountInvestorResponse) *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse {
	return &NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountInvestorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
