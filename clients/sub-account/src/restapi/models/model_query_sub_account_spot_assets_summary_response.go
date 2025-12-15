/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the QuerySubAccountSpotAssetsSummaryResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &QuerySubAccountSpotAssetsSummaryResponse{}

// QuerySubAccountSpotAssetsSummaryResponse struct for QuerySubAccountSpotAssetsSummaryResponse
type QuerySubAccountSpotAssetsSummaryResponse struct {
	TotalCount                *int64                                                                   `json:"totalCount,omitempty"`
	MasterAccountTotalAsset   *string                                                                  `json:"masterAccountTotalAsset,omitempty"`
	SpotSubUserAssetBtcVoList []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner `json:"spotSubUserAssetBtcVoList,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _QuerySubAccountSpotAssetsSummaryResponse QuerySubAccountSpotAssetsSummaryResponse

// NewQuerySubAccountSpotAssetsSummaryResponse instantiates a new QuerySubAccountSpotAssetsSummaryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySubAccountSpotAssetsSummaryResponse() *QuerySubAccountSpotAssetsSummaryResponse {
	this := QuerySubAccountSpotAssetsSummaryResponse{}
	return &this
}

// NewQuerySubAccountSpotAssetsSummaryResponseWithDefaults instantiates a new QuerySubAccountSpotAssetsSummaryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySubAccountSpotAssetsSummaryResponseWithDefaults() *QuerySubAccountSpotAssetsSummaryResponse {
	this := QuerySubAccountSpotAssetsSummaryResponse{}
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetTotalCount() int64 {
	if o == nil || common.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) HasTotalCount() bool {
	if o != nil && !common.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *QuerySubAccountSpotAssetsSummaryResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetMasterAccountTotalAsset returns the MasterAccountTotalAsset field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetMasterAccountTotalAsset() string {
	if o == nil || common.IsNil(o.MasterAccountTotalAsset) {
		var ret string
		return ret
	}
	return *o.MasterAccountTotalAsset
}

// GetMasterAccountTotalAssetOk returns a tuple with the MasterAccountTotalAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetMasterAccountTotalAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.MasterAccountTotalAsset) {
		return nil, false
	}
	return o.MasterAccountTotalAsset, true
}

// HasMasterAccountTotalAsset returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) HasMasterAccountTotalAsset() bool {
	if o != nil && !common.IsNil(o.MasterAccountTotalAsset) {
		return true
	}

	return false
}

// SetMasterAccountTotalAsset gets a reference to the given string and assigns it to the MasterAccountTotalAsset field.
func (o *QuerySubAccountSpotAssetsSummaryResponse) SetMasterAccountTotalAsset(v string) {
	o.MasterAccountTotalAsset = &v
}

// GetSpotSubUserAssetBtcVoList returns the SpotSubUserAssetBtcVoList field value if set, zero value otherwise.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetSpotSubUserAssetBtcVoList() []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner {
	if o == nil || common.IsNil(o.SpotSubUserAssetBtcVoList) {
		var ret []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner
		return ret
	}
	return o.SpotSubUserAssetBtcVoList
}

// GetSpotSubUserAssetBtcVoListOk returns a tuple with the SpotSubUserAssetBtcVoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) GetSpotSubUserAssetBtcVoListOk() ([]QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner, bool) {
	if o == nil || common.IsNil(o.SpotSubUserAssetBtcVoList) {
		return nil, false
	}
	return o.SpotSubUserAssetBtcVoList, true
}

// HasSpotSubUserAssetBtcVoList returns a boolean if a field has been set.
func (o *QuerySubAccountSpotAssetsSummaryResponse) HasSpotSubUserAssetBtcVoList() bool {
	if o != nil && !common.IsNil(o.SpotSubUserAssetBtcVoList) {
		return true
	}

	return false
}

// SetSpotSubUserAssetBtcVoList gets a reference to the given []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner and assigns it to the SpotSubUserAssetBtcVoList field.
func (o *QuerySubAccountSpotAssetsSummaryResponse) SetSpotSubUserAssetBtcVoList(v []QuerySubAccountSpotAssetsSummaryResponseSpotSubUserAssetBtcVoListInner) {
	o.SpotSubUserAssetBtcVoList = v
}

func (o QuerySubAccountSpotAssetsSummaryResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuerySubAccountSpotAssetsSummaryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	if !common.IsNil(o.MasterAccountTotalAsset) {
		toSerialize["masterAccountTotalAsset"] = o.MasterAccountTotalAsset
	}
	if !common.IsNil(o.SpotSubUserAssetBtcVoList) {
		toSerialize["spotSubUserAssetBtcVoList"] = o.SpotSubUserAssetBtcVoList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuerySubAccountSpotAssetsSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	varQuerySubAccountSpotAssetsSummaryResponse := _QuerySubAccountSpotAssetsSummaryResponse{}

	err = json.Unmarshal(data, &varQuerySubAccountSpotAssetsSummaryResponse)

	if err != nil {
		return err
	}

	*o = QuerySubAccountSpotAssetsSummaryResponse(varQuerySubAccountSpotAssetsSummaryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalCount")
		delete(additionalProperties, "masterAccountTotalAsset")
		delete(additionalProperties, "spotSubUserAssetBtcVoList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuerySubAccountSpotAssetsSummaryResponse struct {
	value *QuerySubAccountSpotAssetsSummaryResponse
	isSet bool
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponse) Get() *QuerySubAccountSpotAssetsSummaryResponse {
	return v.value
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponse) Set(val *QuerySubAccountSpotAssetsSummaryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuerySubAccountSpotAssetsSummaryResponse(val *QuerySubAccountSpotAssetsSummaryResponse) *NullableQuerySubAccountSpotAssetsSummaryResponse {
	return &NullableQuerySubAccountSpotAssetsSummaryResponse{value: val, isSet: true}
}

func (v NullableQuerySubAccountSpotAssetsSummaryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuerySubAccountSpotAssetsSummaryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
