/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSummaryOfSubAccountsMarginAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsMarginAccountResponse{}

// GetSummaryOfSubAccountsMarginAccountResponse struct for GetSummaryOfSubAccountsMarginAccountResponse
type GetSummaryOfSubAccountsMarginAccountResponse struct {
	TotalAssetOfBtc      *string                                                           `json:"totalAssetOfBtc,omitempty"`
	TotalLiabilityOfBtc  *string                                                           `json:"totalLiabilityOfBtc,omitempty"`
	TotalNetAssetOfBtc   *string                                                           `json:"totalNetAssetOfBtc,omitempty"`
	SubAccountList       []GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner `json:"subAccountList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetSummaryOfSubAccountsMarginAccountResponse GetSummaryOfSubAccountsMarginAccountResponse

// NewGetSummaryOfSubAccountsMarginAccountResponse instantiates a new GetSummaryOfSubAccountsMarginAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsMarginAccountResponse() *GetSummaryOfSubAccountsMarginAccountResponse {
	this := GetSummaryOfSubAccountsMarginAccountResponse{}
	return &this
}

// NewGetSummaryOfSubAccountsMarginAccountResponseWithDefaults instantiates a new GetSummaryOfSubAccountsMarginAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsMarginAccountResponseWithDefaults() *GetSummaryOfSubAccountsMarginAccountResponse {
	this := GetSummaryOfSubAccountsMarginAccountResponse{}
	return &this
}

// GetTotalAssetOfBtc returns the TotalAssetOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalAssetOfBtc
}

// GetTotalAssetOfBtcOk returns a tuple with the TotalAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAssetOfBtc) {
		return nil, false
	}
	return o.TotalAssetOfBtc, true
}

// HasTotalAssetOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) HasTotalAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalAssetOfBtc gets a reference to the given string and assigns it to the TotalAssetOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) SetTotalAssetOfBtc(v string) {
	o.TotalAssetOfBtc = &v
}

// GetTotalLiabilityOfBtc returns the TotalLiabilityOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtc() string {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalLiabilityOfBtc
}

// GetTotalLiabilityOfBtcOk returns a tuple with the TotalLiabilityOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalLiabilityOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalLiabilityOfBtc) {
		return nil, false
	}
	return o.TotalLiabilityOfBtc, true
}

// HasTotalLiabilityOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) HasTotalLiabilityOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalLiabilityOfBtc) {
		return true
	}

	return false
}

// SetTotalLiabilityOfBtc gets a reference to the given string and assigns it to the TotalLiabilityOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) SetTotalLiabilityOfBtc(v string) {
	o.TotalLiabilityOfBtc = &v
}

// GetTotalNetAssetOfBtc returns the TotalNetAssetOfBtc field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtc() string {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		var ret string
		return ret
	}
	return *o.TotalNetAssetOfBtc
}

// GetTotalNetAssetOfBtcOk returns a tuple with the TotalNetAssetOfBtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetTotalNetAssetOfBtcOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalNetAssetOfBtc) {
		return nil, false
	}
	return o.TotalNetAssetOfBtc, true
}

// HasTotalNetAssetOfBtc returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) HasTotalNetAssetOfBtc() bool {
	if o != nil && !common.IsNil(o.TotalNetAssetOfBtc) {
		return true
	}

	return false
}

// SetTotalNetAssetOfBtc gets a reference to the given string and assigns it to the TotalNetAssetOfBtc field.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) SetTotalNetAssetOfBtc(v string) {
	o.TotalNetAssetOfBtc = &v
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetSubAccountList() []GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner {
	if o == nil || common.IsNil(o.SubAccountList) {
		var ret []GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) GetSubAccountListOk() ([]GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner, bool) {
	if o == nil || common.IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) HasSubAccountList() bool {
	if o != nil && !common.IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSummaryOfSubAccountsMarginAccountResponse) SetSubAccountList(v []GetSummaryOfSubAccountsMarginAccountResponseSubAccountListInner) {
	o.SubAccountList = v
}

func (o GetSummaryOfSubAccountsMarginAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsMarginAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalAssetOfBtc) {
		toSerialize["totalAssetOfBtc"] = o.TotalAssetOfBtc
	}
	if !common.IsNil(o.TotalLiabilityOfBtc) {
		toSerialize["totalLiabilityOfBtc"] = o.TotalLiabilityOfBtc
	}
	if !common.IsNil(o.TotalNetAssetOfBtc) {
		toSerialize["totalNetAssetOfBtc"] = o.TotalNetAssetOfBtc
	}
	if !common.IsNil(o.SubAccountList) {
		toSerialize["subAccountList"] = o.SubAccountList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfSubAccountsMarginAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsMarginAccountResponse := _GetSummaryOfSubAccountsMarginAccountResponse{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsMarginAccountResponse)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsMarginAccountResponse(varGetSummaryOfSubAccountsMarginAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalAssetOfBtc")
		delete(additionalProperties, "totalLiabilityOfBtc")
		delete(additionalProperties, "totalNetAssetOfBtc")
		delete(additionalProperties, "subAccountList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsMarginAccountResponse struct {
	value *GetSummaryOfSubAccountsMarginAccountResponse
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponse) Get() *GetSummaryOfSubAccountsMarginAccountResponse {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponse) Set(val *GetSummaryOfSubAccountsMarginAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsMarginAccountResponse(val *GetSummaryOfSubAccountsMarginAccountResponse) *NullableGetSummaryOfSubAccountsMarginAccountResponse {
	return &NullableGetSummaryOfSubAccountsMarginAccountResponse{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsMarginAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsMarginAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
