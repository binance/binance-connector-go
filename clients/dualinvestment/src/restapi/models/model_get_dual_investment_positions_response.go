/*
Binance Dual Investment REST API

OpenAPI Specification for the Binance Dual Investment REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetDualInvestmentPositionsResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetDualInvestmentPositionsResponse{}

// GetDualInvestmentPositionsResponse struct for GetDualInvestmentPositionsResponse
type GetDualInvestmentPositionsResponse struct {
	Total                *int64                                        `json:"total,omitempty"`
	List                 []GetDualInvestmentPositionsResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetDualInvestmentPositionsResponse GetDualInvestmentPositionsResponse

// NewGetDualInvestmentPositionsResponse instantiates a new GetDualInvestmentPositionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDualInvestmentPositionsResponse() *GetDualInvestmentPositionsResponse {
	this := GetDualInvestmentPositionsResponse{}
	return &this
}

// NewGetDualInvestmentPositionsResponseWithDefaults instantiates a new GetDualInvestmentPositionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDualInvestmentPositionsResponseWithDefaults() *GetDualInvestmentPositionsResponse {
	this := GetDualInvestmentPositionsResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetDualInvestmentPositionsResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetDualInvestmentPositionsResponse) GetList() []GetDualInvestmentPositionsResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetDualInvestmentPositionsResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDualInvestmentPositionsResponse) GetListOk() ([]GetDualInvestmentPositionsResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetDualInvestmentPositionsResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetDualInvestmentPositionsResponseListInner and assigns it to the List field.
func (o *GetDualInvestmentPositionsResponse) SetList(v []GetDualInvestmentPositionsResponseListInner) {
	o.List = v
}

func (o GetDualInvestmentPositionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDualInvestmentPositionsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetDualInvestmentPositionsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetDualInvestmentPositionsResponse := _GetDualInvestmentPositionsResponse{}

	err = json.Unmarshal(data, &varGetDualInvestmentPositionsResponse)

	if err != nil {
		return err
	}

	*o = GetDualInvestmentPositionsResponse(varGetDualInvestmentPositionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetDualInvestmentPositionsResponse struct {
	value *GetDualInvestmentPositionsResponse
	isSet bool
}

func (v NullableGetDualInvestmentPositionsResponse) Get() *GetDualInvestmentPositionsResponse {
	return v.value
}

func (v *NullableGetDualInvestmentPositionsResponse) Set(val *GetDualInvestmentPositionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDualInvestmentPositionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDualInvestmentPositionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDualInvestmentPositionsResponse(val *GetDualInvestmentPositionsResponse) *NullableGetDualInvestmentPositionsResponse {
	return &NullableGetDualInvestmentPositionsResponse{value: val, isSet: true}
}

func (v NullableGetDualInvestmentPositionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDualInvestmentPositionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
