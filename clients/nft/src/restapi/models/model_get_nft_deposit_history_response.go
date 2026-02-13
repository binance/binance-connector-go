/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetNFTDepositHistoryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTDepositHistoryResponse{}

// GetNFTDepositHistoryResponse struct for GetNFTDepositHistoryResponse
type GetNFTDepositHistoryResponse struct {
	Total                *int64                                  `json:"total,omitempty"`
	List                 []GetNFTDepositHistoryResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTDepositHistoryResponse GetNFTDepositHistoryResponse

// NewGetNFTDepositHistoryResponse instantiates a new GetNFTDepositHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTDepositHistoryResponse() *GetNFTDepositHistoryResponse {
	this := GetNFTDepositHistoryResponse{}
	return &this
}

// NewGetNFTDepositHistoryResponseWithDefaults instantiates a new GetNFTDepositHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTDepositHistoryResponseWithDefaults() *GetNFTDepositHistoryResponse {
	this := GetNFTDepositHistoryResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetNFTDepositHistoryResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetNFTDepositHistoryResponse) GetList() []GetNFTDepositHistoryResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetNFTDepositHistoryResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTDepositHistoryResponse) GetListOk() ([]GetNFTDepositHistoryResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetNFTDepositHistoryResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetNFTDepositHistoryResponseListInner and assigns it to the List field.
func (o *GetNFTDepositHistoryResponse) SetList(v []GetNFTDepositHistoryResponseListInner) {
	o.List = v
}

func (o GetNFTDepositHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTDepositHistoryResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetNFTDepositHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varGetNFTDepositHistoryResponse := _GetNFTDepositHistoryResponse{}

	err = json.Unmarshal(data, &varGetNFTDepositHistoryResponse)

	if err != nil {
		return err
	}

	*o = GetNFTDepositHistoryResponse(varGetNFTDepositHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTDepositHistoryResponse struct {
	value *GetNFTDepositHistoryResponse
	isSet bool
}

func (v NullableGetNFTDepositHistoryResponse) Get() *GetNFTDepositHistoryResponse {
	return v.value
}

func (v *NullableGetNFTDepositHistoryResponse) Set(val *GetNFTDepositHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTDepositHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTDepositHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTDepositHistoryResponse(val *GetNFTDepositHistoryResponse) *NullableGetNFTDepositHistoryResponse {
	return &NullableGetNFTDepositHistoryResponse{value: val, isSet: true}
}

func (v NullableGetNFTDepositHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTDepositHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
