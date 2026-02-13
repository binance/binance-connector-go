/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QueryManagedSubAccountListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QueryManagedSubAccountListResponse{}

// QueryManagedSubAccountListResponse struct for QueryManagedSubAccountListResponse
type QueryManagedSubAccountListResponse struct {
	Total                    *int64                                                            `json:"total,omitempty"`
	ManagerSubUserInfoVoList []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner `json:"managerSubUserInfoVoList,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _QueryManagedSubAccountListResponse QueryManagedSubAccountListResponse

// NewQueryManagedSubAccountListResponse instantiates a new QueryManagedSubAccountListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryManagedSubAccountListResponse() *QueryManagedSubAccountListResponse {
	this := QueryManagedSubAccountListResponse{}
	return &this
}

// NewQueryManagedSubAccountListResponseWithDefaults instantiates a new QueryManagedSubAccountListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryManagedSubAccountListResponseWithDefaults() *QueryManagedSubAccountListResponse {
	this := QueryManagedSubAccountListResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QueryManagedSubAccountListResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetManagerSubUserInfoVoList returns the ManagerSubUserInfoVoList field value if set, zero value otherwise.
func (o *QueryManagedSubAccountListResponse) GetManagerSubUserInfoVoList() []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner {
	if o == nil || common.IsNil(o.ManagerSubUserInfoVoList) {
		var ret []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner
		return ret
	}
	return o.ManagerSubUserInfoVoList
}

// GetManagerSubUserInfoVoListOk returns a tuple with the ManagerSubUserInfoVoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryManagedSubAccountListResponse) GetManagerSubUserInfoVoListOk() ([]QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner, bool) {
	if o == nil || common.IsNil(o.ManagerSubUserInfoVoList) {
		return nil, false
	}
	return o.ManagerSubUserInfoVoList, true
}

// HasManagerSubUserInfoVoList returns a boolean if a field has been set.
func (o *QueryManagedSubAccountListResponse) HasManagerSubUserInfoVoList() bool {
	if o != nil && !common.IsNil(o.ManagerSubUserInfoVoList) {
		return true
	}

	return false
}

// SetManagerSubUserInfoVoList gets a reference to the given []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner and assigns it to the ManagerSubUserInfoVoList field.
func (o *QueryManagedSubAccountListResponse) SetManagerSubUserInfoVoList(v []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner) {
	o.ManagerSubUserInfoVoList = v
}

func (o QueryManagedSubAccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryManagedSubAccountListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.ManagerSubUserInfoVoList) {
		toSerialize["managerSubUserInfoVoList"] = o.ManagerSubUserInfoVoList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QueryManagedSubAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	varQueryManagedSubAccountListResponse := _QueryManagedSubAccountListResponse{}

	err = json.Unmarshal(data, &varQueryManagedSubAccountListResponse)

	if err != nil {
		return err
	}

	*o = QueryManagedSubAccountListResponse(varQueryManagedSubAccountListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "managerSubUserInfoVoList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQueryManagedSubAccountListResponse struct {
	value *QueryManagedSubAccountListResponse
	isSet bool
}

func (v NullableQueryManagedSubAccountListResponse) Get() *QueryManagedSubAccountListResponse {
	return v.value
}

func (v *NullableQueryManagedSubAccountListResponse) Set(val *QueryManagedSubAccountListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryManagedSubAccountListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryManagedSubAccountListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryManagedSubAccountListResponse(val *QueryManagedSubAccountListResponse) *NullableQueryManagedSubAccountListResponse {
	return &NullableQueryManagedSubAccountListResponse{value: val, isSet: true}
}

func (v NullableQueryManagedSubAccountListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryManagedSubAccountListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
