/*
Binance NFT REST API

OpenAPI Specification for the Binance NFT REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetNFTAssetResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetNFTAssetResponse{}

// GetNFTAssetResponse struct for GetNFTAssetResponse
type GetNFTAssetResponse struct {
	Total                *int64                         `json:"total,omitempty"`
	List                 []GetNFTAssetResponseListInner `json:"list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetNFTAssetResponse GetNFTAssetResponse

// NewGetNFTAssetResponse instantiates a new GetNFTAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNFTAssetResponse() *GetNFTAssetResponse {
	this := GetNFTAssetResponse{}
	return &this
}

// NewGetNFTAssetResponseWithDefaults instantiates a new GetNFTAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNFTAssetResponseWithDefaults() *GetNFTAssetResponse {
	this := GetNFTAssetResponse{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *GetNFTAssetResponse) GetTotal() int64 {
	if o == nil || common.IsNil(o.Total) {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTAssetResponse) GetTotalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *GetNFTAssetResponse) HasTotal() bool {
	if o != nil && !common.IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *GetNFTAssetResponse) SetTotal(v int64) {
	o.Total = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *GetNFTAssetResponse) GetList() []GetNFTAssetResponseListInner {
	if o == nil || common.IsNil(o.List) {
		var ret []GetNFTAssetResponseListInner
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNFTAssetResponse) GetListOk() ([]GetNFTAssetResponseListInner, bool) {
	if o == nil || common.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *GetNFTAssetResponse) HasList() bool {
	if o != nil && !common.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []GetNFTAssetResponseListInner and assigns it to the List field.
func (o *GetNFTAssetResponse) SetList(v []GetNFTAssetResponseListInner) {
	o.List = v
}

func (o GetNFTAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNFTAssetResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetNFTAssetResponse) UnmarshalJSON(data []byte) (err error) {
	varGetNFTAssetResponse := _GetNFTAssetResponse{}

	err = json.Unmarshal(data, &varGetNFTAssetResponse)

	if err != nil {
		return err
	}

	*o = GetNFTAssetResponse(varGetNFTAssetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetNFTAssetResponse struct {
	value *GetNFTAssetResponse
	isSet bool
}

func (v NullableGetNFTAssetResponse) Get() *GetNFTAssetResponse {
	return v.value
}

func (v *NullableGetNFTAssetResponse) Set(val *GetNFTAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNFTAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNFTAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNFTAssetResponse(val *GetNFTAssetResponse) *NullableGetNFTAssetResponse {
	return &NullableGetNFTAssetResponse{value: val, isSet: true}
}

func (v NullableGetNFTAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNFTAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
