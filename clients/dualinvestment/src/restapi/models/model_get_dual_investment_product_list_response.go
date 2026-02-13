/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetDualInvestmentProductListResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDualInvestmentProductListResponse{}

// GetDualInvestmentProductListResponse struct for GetDualInvestmentProductListResponse
type GetDualInvestmentProductListResponse struct {
	Total                *int64                                          `json:"total,omitempty"`
	List                 []GetDualInvestmentProductListResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDualInvestmentProductListResponse GetDualInvestmentProductListResponse

// NewGetDualInvestmentProductListResponse instantiates a new GetDualInvestmentProductListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDualInvestmentProductListResponse() *GetDualInvestmentProductListResponse {
	this := GetDualInvestmentProductListResponse{}
	return &this
}

// NewGetDualInvestmentProductListResponseWithDefaults instantiates a new GetDualInvestmentProductListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDualInvestmentProductListResponseWithDefaults() *GetDualInvestmentProductListResponse {
	this := GetDualInvestmentProductListResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetDualInvestmentProductListResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetDualInvestmentProductListResponse) GetList() []GetDualInvestmentProductListResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetDualInvestmentProductListResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentProductListResponse) GetListOk() ([]GetDualInvestmentProductListResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetDualInvestmentProductListResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetDualInvestmentProductListResponseListInner and assigns it to the List field.
func (o *GetDualInvestmentProductListResponse) SetList(v []GetDualInvestmentProductListResponseListInner) {
	o.List = v
}

func (o GetDualInvestmentProductListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDualInvestmentProductListResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetDualInvestmentProductListResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDualInvestmentProductListResponse := _GetDualInvestmentProductListResponse{}

	err = json.Unmarshal(data, &varGetDualInvestmentProductListResponse)

	if err != nil {
		return err
	}

	*o = GetDualInvestmentProductListResponse(varGetDualInvestmentProductListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDualInvestmentProductListResponse struct {
	value *GetDualInvestmentProductListResponse
	isSet bool
}

func (v NullableGetDualInvestmentProductListResponse) Get() *GetDualInvestmentProductListResponse {
	return v.value
}

func (v *NullableGetDualInvestmentProductListResponse) Set(val *GetDualInvestmentProductListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentProductListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentProductListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentProductListResponse(val *GetDualInvestmentProductListResponse) *NullableGetDualInvestmentProductListResponse {
	return &NullableGetDualInvestmentProductListResponse{value: val, isSet: true}
}

func (v NullableGetDualInvestmentProductListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentProductListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
