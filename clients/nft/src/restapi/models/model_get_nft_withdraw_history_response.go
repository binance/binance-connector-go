/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetNFTWithdrawHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTWithdrawHistoryResponse{}

// GetNFTWithdrawHistoryResponse struct for GetNFTWithdrawHistoryResponse
type GetNFTWithdrawHistoryResponse struct {
	Total                *int64                                   `json:"total,omitempty"`
	List                 []GetNFTWithdrawHistoryResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTWithdrawHistoryResponse GetNFTWithdrawHistoryResponse

// NewGetNFTWithdrawHistoryResponse instantiates a new GetNFTWithdrawHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTWithdrawHistoryResponse() *GetNFTWithdrawHistoryResponse {
	this := GetNFTWithdrawHistoryResponse{}
	return &this
}

// NewGetNFTWithdrawHistoryResponseWithDefaults instantiates a new GetNFTWithdrawHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTWithdrawHistoryResponseWithDefaults() *GetNFTWithdrawHistoryResponse {
	this := GetNFTWithdrawHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetNFTWithdrawHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetNFTWithdrawHistoryResponse) GetList() []GetNFTWithdrawHistoryResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetNFTWithdrawHistoryResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTWithdrawHistoryResponse) GetListOk() ([]GetNFTWithdrawHistoryResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetNFTWithdrawHistoryResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetNFTWithdrawHistoryResponseListInner and assigns it to the List field.
func (o *GetNFTWithdrawHistoryResponse) SetList(v []GetNFTWithdrawHistoryResponseListInner) {
	o.List = v
}

func (o GetNFTWithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTWithdrawHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetNFTWithdrawHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetNFTWithdrawHistoryResponse := _GetNFTWithdrawHistoryResponse{}

	err = json.Unmarshal(data, &varGetNFTWithdrawHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetNFTWithdrawHistoryResponse(varGetNFTWithdrawHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTWithdrawHistoryResponse struct {
	value *GetNFTWithdrawHistoryResponse
	isSet bool
}

func (v NullableGetNFTWithdrawHistoryResponse) Get() *GetNFTWithdrawHistoryResponse {
	return v.value
}

func (v *NullableGetNFTWithdrawHistoryResponse) Set(val *GetNFTWithdrawHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTWithdrawHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTWithdrawHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTWithdrawHistoryResponse(val *GetNFTWithdrawHistoryResponse) *NullableGetNFTWithdrawHistoryResponse {
	return &NullableGetNFTWithdrawHistoryResponse{value: val, isSet: true}
}

func (v NullableGetNFTWithdrawHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTWithdrawHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
