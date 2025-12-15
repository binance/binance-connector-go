/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetNFTTransactionHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTTransactionHistoryResponse{}

// GetNFTTransactionHistoryResponse struct for GetNFTTransactionHistoryResponse
type GetNFTTransactionHistoryResponse struct {
	Total                *int64                                      `json:"total,omitempty"`
	List                 []GetNFTTransactionHistoryResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTTransactionHistoryResponse GetNFTTransactionHistoryResponse

// NewGetNFTTransactionHistoryResponse instantiates a new GetNFTTransactionHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTTransactionHistoryResponse() *GetNFTTransactionHistoryResponse {
	this := GetNFTTransactionHistoryResponse{}
	return &this
}

// NewGetNFTTransactionHistoryResponseWithDefaults instantiates a new GetNFTTransactionHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTTransactionHistoryResponseWithDefaults() *GetNFTTransactionHistoryResponse {
	this := GetNFTTransactionHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetNFTTransactionHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetNFTTransactionHistoryResponse) GetList() []GetNFTTransactionHistoryResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetNFTTransactionHistoryResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTTransactionHistoryResponse) GetListOk() ([]GetNFTTransactionHistoryResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetNFTTransactionHistoryResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetNFTTransactionHistoryResponseListInner and assigns it to the List field.
func (o *GetNFTTransactionHistoryResponse) SetList(v []GetNFTTransactionHistoryResponseListInner) {
	o.List = v
}

func (o GetNFTTransactionHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTTransactionHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetNFTTransactionHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetNFTTransactionHistoryResponse := _GetNFTTransactionHistoryResponse{}

	err = json.Unmarshal(data, &varGetNFTTransactionHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetNFTTransactionHistoryResponse(varGetNFTTransactionHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTTransactionHistoryResponse struct {
	value *GetNFTTransactionHistoryResponse
	isSet bool
}

func (v NullableGetNFTTransactionHistoryResponse) Get() *GetNFTTransactionHistoryResponse {
	return v.value
}

func (v *NullableGetNFTTransactionHistoryResponse) Set(val *GetNFTTransactionHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTTransactionHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTTransactionHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTTransactionHistoryResponse(val *GetNFTTransactionHistoryResponse) *NullableGetNFTTransactionHistoryResponse {
	return &NullableGetNFTTransactionHistoryResponse{value: val, isSet: true}
}

func (v NullableGetNFTTransactionHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTTransactionHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
