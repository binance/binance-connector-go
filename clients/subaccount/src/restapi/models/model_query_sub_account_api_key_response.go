/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the QuerySubAccountApiKeyResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountApiKeyResponse{}

// QuerySubAccountApiKeyResponse struct for QuerySubAccountApiKeyResponse
type QuerySubAccountApiKeyResponse struct {
	Total                *int64                                   `json:"total,omitempty"`
	List                 []QuerySubAccountApiKeyResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuerySubAccountApiKeyResponse QuerySubAccountApiKeyResponse

// NewQuerySubAccountApiKeyResponse instantiates a new QuerySubAccountApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountApiKeyResponse() *QuerySubAccountApiKeyResponse {
	this := QuerySubAccountApiKeyResponse{}
	return &this
}

// NewQuerySubAccountApiKeyResponseWithDefaults instantiates a new QuerySubAccountApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountApiKeyResponseWithDefaults() *QuerySubAccountApiKeyResponse {
	this := QuerySubAccountApiKeyResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *QuerySubAccountApiKeyResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *QuerySubAccountApiKeyResponse) GetList() []QuerySubAccountApiKeyResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []QuerySubAccountApiKeyResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountApiKeyResponse) GetListOk() ([]QuerySubAccountApiKeyResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *QuerySubAccountApiKeyResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []QuerySubAccountApiKeyResponseListInner and assigns it to the List field.
func (o *QuerySubAccountApiKeyResponse) SetList(v []QuerySubAccountApiKeyResponseListInner) {
	o.List = v
}

func (o QuerySubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountApiKeyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !common.IsNil(o.List) {
		toSerialize["list"] = o.List
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountApiKeyResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountApiKeyResponse := _QuerySubAccountApiKeyResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountApiKeyResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountApiKeyResponse(varQuerySubAccountApiKeyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountApiKeyResponse struct {
	value *QuerySubAccountApiKeyResponse
	isSet bool
}

func (v NullableQuerySubAccountApiKeyResponse) Get() *QuerySubAccountApiKeyResponse {
	return v.value
}

func (v *NullableQuerySubAccountApiKeyResponse) Set(val *QuerySubAccountApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountApiKeyResponse(val *QuerySubAccountApiKeyResponse) *NullableQuerySubAccountApiKeyResponse {
	return &NullableQuerySubAccountApiKeyResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
