/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountTransferLogMasterAccountTradingResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountTransferLogMasterAccountTradingResponse{}

// QueryManagedSubAccountTransferLogMasterAccountTradingResponse struct for QueryManagedSubAccountTransferLogMasterAccountTradingResponse
type QueryManagedSubAccountTransferLogMasterAccountTradingResponse struct {
	ManagerSubTransferHistoryVos []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner `json:"managerSubTransferHistoryVos,omitempty"`
	Count                        *int64                                                                                            `json:"count,omitempty"`
	AdditionalProperties         map[string]interface{}
}

type _QueryManagedSubAccountTransferLogMasterAccountTradingResponse QueryManagedSubAccountTransferLogMasterAccountTradingResponse

// NewQueryManagedSubAccountTransferLogMasterAccountTradingResponse instantiates a new QueryManagedSubAccountTransferLogMasterAccountTradingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountTransferLogMasterAccountTradingResponse() *QueryManagedSubAccountTransferLogMasterAccountTradingResponse {
	this := QueryManagedSubAccountTransferLogMasterAccountTradingResponse{}
	return &this
}

// NewQueryManagedSubAccountTransferLogMasterAccountTradingResponseWithDefaults instantiates a new QueryManagedSubAccountTransferLogMasterAccountTradingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountTransferLogMasterAccountTradingResponseWithDefaults() *QueryManagedSubAccountTransferLogMasterAccountTradingResponse {
	this := QueryManagedSubAccountTransferLogMasterAccountTradingResponse{}
	return &this
}

// GetManagerSubTransferHistoryVos returns the ManagerSubTransferHistoryVos field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) GetManagerSubTransferHistoryVos() []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		var ret []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner
		return ret
	}
	return o.ManagerSubTransferHistoryVos
}

// GetManagerSubTransferHistoryVosOk returns a tuple with the ManagerSubTransferHistoryVos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) GetManagerSubTransferHistoryVosOk() ([]QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner, bool) {
	if o == nil || common.IsNil(o.ManagerSubTransferHistoryVos) {
		return nil, false
	}
	return o.ManagerSubTransferHistoryVos, true
}

// HasManagerSubTransferHistoryVos returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) HasManagerSubTransferHistoryVos() bool {
	if o != nil && !common.IsNil(o.ManagerSubTransferHistoryVos) {
		return true
	}

	return false
}

// SetManagerSubTransferHistoryVos gets a reference to the given []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner and assigns it to the ManagerSubTransferHistoryVos field.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) SetManagerSubTransferHistoryVos(v []QueryManagedSubAccountTransferLogMasterAccountInvestorResponseManagerSubTransferHistoryVosInner) {
	o.ManagerSubTransferHistoryVos = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) GetCount() int64 {
	if o == nil || common.IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) GetCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) HasCount() bool {
	if o != nil && !common.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) SetCount(v int64) {
	o.Count = &v
}

func (o QueryManagedSubAccountTransferLogMasterAccountTradingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountTransferLogMasterAccountTradingResponse) ToMap() (map[string]interface{}, error) {
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

func (o *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountTransferLogMasterAccountTradingResponse := _QueryManagedSubAccountTransferLogMasterAccountTradingResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountTransferLogMasterAccountTradingResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountTransferLogMasterAccountTradingResponse(varQueryManagedSubAccountTransferLogMasterAccountTradingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "managerSubTransferHistoryVos")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse struct {
	value *QueryManagedSubAccountTransferLogMasterAccountTradingResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) Get() *QueryManagedSubAccountTransferLogMasterAccountTradingResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) Set(val *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse(val *QueryManagedSubAccountTransferLogMasterAccountTradingResponse) *NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse {
	return &NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountTransferLogMasterAccountTradingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
