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

// checks if the GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp{}

// GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp struct for GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp
type GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp struct {
	TotalMarginBalanceOfBTC    *string                                                                                        `json:"totalMarginBalanceOfBTC,omitempty"`
	TotalUnrealizedProfitOfBTC *string                                                                                        `json:"totalUnrealizedProfitOfBTC,omitempty"`
	TotalWalletBalanceOfBTC    *string                                                                                        `json:"totalWalletBalanceOfBTC,omitempty"`
	Asset                      *string                                                                                        `json:"asset,omitempty"`
	SubAccountList             []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner `json:"subAccountList,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp{}
	return &this
}

// GetTotalMarginBalanceOfBTC returns the TotalMarginBalanceOfBTC field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTC() string {
	if o == nil || common.IsNil(o.TotalMarginBalanceOfBTC) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalanceOfBTC
}

// GetTotalMarginBalanceOfBTCOk returns a tuple with the TotalMarginBalanceOfBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalMarginBalanceOfBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalanceOfBTC) {
		return nil, false
	}
	return o.TotalMarginBalanceOfBTC, true
}

// HasTotalMarginBalanceOfBTC returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalMarginBalanceOfBTC() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalanceOfBTC) {
		return true
	}

	return false
}

// SetTotalMarginBalanceOfBTC gets a reference to the given string and assigns it to the TotalMarginBalanceOfBTC field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalMarginBalanceOfBTC(v string) {
	o.TotalMarginBalanceOfBTC = &v
}

// GetTotalUnrealizedProfitOfBTC returns the TotalUnrealizedProfitOfBTC field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTC() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfitOfBTC) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfitOfBTC
}

// GetTotalUnrealizedProfitOfBTCOk returns a tuple with the TotalUnrealizedProfitOfBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalUnrealizedProfitOfBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfitOfBTC) {
		return nil, false
	}
	return o.TotalUnrealizedProfitOfBTC, true
}

// HasTotalUnrealizedProfitOfBTC returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalUnrealizedProfitOfBTC() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfitOfBTC) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfitOfBTC gets a reference to the given string and assigns it to the TotalUnrealizedProfitOfBTC field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalUnrealizedProfitOfBTC(v string) {
	o.TotalUnrealizedProfitOfBTC = &v
}

// GetTotalWalletBalanceOfBTC returns the TotalWalletBalanceOfBTC field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTC() string {
	if o == nil || common.IsNil(o.TotalWalletBalanceOfBTC) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalanceOfBTC
}

// GetTotalWalletBalanceOfBTCOk returns a tuple with the TotalWalletBalanceOfBTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetTotalWalletBalanceOfBTCOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalanceOfBTC) {
		return nil, false
	}
	return o.TotalWalletBalanceOfBTC, true
}

// HasTotalWalletBalanceOfBTC returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasTotalWalletBalanceOfBTC() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalanceOfBTC) {
		return true
	}

	return false
}

// SetTotalWalletBalanceOfBTC gets a reference to the given string and assigns it to the TotalWalletBalanceOfBTC field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetTotalWalletBalanceOfBTC(v string) {
	o.TotalWalletBalanceOfBTC = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetAsset(v string) {
	o.Asset = &v
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetSubAccountList() []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner {
	if o == nil || common.IsNil(o.SubAccountList) {
		var ret []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) GetSubAccountListOk() ([]GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner, bool) {
	if o == nil || common.IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) HasSubAccountList() bool {
	if o != nil && !common.IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) SetSubAccountList(v []GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) {
	o.SubAccountList = v
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalMarginBalanceOfBTC) {
		toSerialize["totalMarginBalanceOfBTC"] = o.TotalMarginBalanceOfBTC
	}
	if !common.IsNil(o.TotalUnrealizedProfitOfBTC) {
		toSerialize["totalUnrealizedProfitOfBTC"] = o.TotalUnrealizedProfitOfBTC
	}
	if !common.IsNil(o.TotalWalletBalanceOfBTC) {
		toSerialize["totalWalletBalanceOfBTC"] = o.TotalWalletBalanceOfBTC
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.SubAccountList) {
		toSerialize["subAccountList"] = o.SubAccountList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp := _GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp(varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalMarginBalanceOfBTC")
		delete(additionalProperties, "totalUnrealizedProfitOfBTC")
		delete(additionalProperties, "totalWalletBalanceOfBTC")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "subAccountList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp struct {
	value *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) Get() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) Set(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp {
	return &NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
