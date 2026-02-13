/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountListResponse{}

// QuerySubAccountListResponse struct for QuerySubAccountListResponse
type QuerySubAccountListResponse struct {
	SubAccounts          []QuerySubAccountListResponseSubAccountsInner `json:"subAccounts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountListResponse QuerySubAccountListResponse

// NewQuerySubAccountListResponse instantiates a new QuerySubAccountListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountListResponse() *QuerySubAccountListResponse {
	this := QuerySubAccountListResponse{}
	return &this
}

// NewQuerySubAccountListResponseWithDefaults instantiates a new QuerySubAccountListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountListResponseWithDefaults() *QuerySubAccountListResponse {
	this := QuerySubAccountListResponse{}
	return &this
}

// GetSubAccounts returns the SubAccounts field value if set, zero value otherwise.
func (o *QuerySubAccountListResponse) GetSubAccounts() []QuerySubAccountListResponseSubAccountsInner {
	if o == nil || common.IsNil(o.SubAccounts) {
		var ret []QuerySubAccountListResponseSubAccountsInner
		return ret
	}
	return o.SubAccounts
}

// GetSubAccountsOk returns a tuple with the SubAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountListResponse) GetSubAccountsOk() ([]QuerySubAccountListResponseSubAccountsInner, bool) {
	if o == nil || common.IsNil(o.SubAccounts) {
		return nil, false
	}
	return o.SubAccounts, true
}

// HasSubAccounts returns a boolean if a field has been set.
func (o *QuerySubAccountListResponse) HasSubAccounts() bool {
	if o != nil && !common.IsNil(o.SubAccounts) {
		return true
	}

	return false
}

// SetSubAccounts gets a reference to the given []QuerySubAccountListResponseSubAccountsInner and assigns it to the SubAccounts field.
func (o *QuerySubAccountListResponse) SetSubAccounts(v []QuerySubAccountListResponseSubAccountsInner) {
	o.SubAccounts = v
}

func (o QuerySubAccountListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.SubAccounts) {
		toSerialize["subAccounts"] = o.SubAccounts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountListResponse := _QuerySubAccountListResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountListResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountListResponse(varQuerySubAccountListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subAccounts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountListResponse struct {
	value *QuerySubAccountListResponse
	isSet bool
}

func (v NullableQuerySubAccountListResponse) Get() *QuerySubAccountListResponse {
	return v.value
}

func (v *NullableQuerySubAccountListResponse) Set(val *QuerySubAccountListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountListResponse(val *QuerySubAccountListResponse) *NullableQuerySubAccountListResponse {
	return &NullableQuerySubAccountListResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
