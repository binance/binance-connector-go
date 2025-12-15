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

// checks if the GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp{}

// GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp struct for GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp
type GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp struct {
	TotalInitialMargin          *string                                                                                      `json:"totalInitialMargin,omitempty"`
	TotalMaintenanceMargin      *string                                                                                      `json:"totalMaintenanceMargin,omitempty"`
	TotalMarginBalance          *string                                                                                      `json:"totalMarginBalance,omitempty"`
	TotalOpenOrderInitialMargin *string                                                                                      `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalPositionInitialMargin  *string                                                                                      `json:"totalPositionInitialMargin,omitempty"`
	TotalUnrealizedProfit       *string                                                                                      `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance          *string                                                                                      `json:"totalWalletBalance,omitempty"`
	Asset                       *string                                                                                      `json:"asset,omitempty"`
	SubAccountList              []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner `json:"subAccountList,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp{}
	return &this
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetAsset(v string) {
	o.Asset = &v
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetSubAccountList() []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	if o == nil || common.IsNil(o.SubAccountList) {
		var ret []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) GetSubAccountListOk() ([]GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner, bool) {
	if o == nil || common.IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) HasSubAccountList() bool {
	if o != nil && !common.IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) SetSubAccountList(v []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) {
	o.SubAccountList = v
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalInitialMargin) {
		toSerialize["totalInitialMargin"] = o.TotalInitialMargin
	}
	if !common.IsNil(o.TotalMaintenanceMargin) {
		toSerialize["totalMaintenanceMargin"] = o.TotalMaintenanceMargin
	}
	if !common.IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
	}
	if !common.IsNil(o.TotalOpenOrderInitialMargin) {
		toSerialize["totalOpenOrderInitialMargin"] = o.TotalOpenOrderInitialMargin
	}
	if !common.IsNil(o.TotalPositionInitialMargin) {
		toSerialize["totalPositionInitialMargin"] = o.TotalPositionInitialMargin
	}
	if !common.IsNil(o.TotalUnrealizedProfit) {
		toSerialize["totalUnrealizedProfit"] = o.TotalUnrealizedProfit
	}
	if !common.IsNil(o.TotalWalletBalance) {
		toSerialize["totalWalletBalance"] = o.TotalWalletBalance
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

func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp := _GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp(varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalInitialMargin")
		delete(additionalProperties, "totalMaintenanceMargin")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalOpenOrderInitialMargin")
		delete(additionalProperties, "totalPositionInitialMargin")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "subAccountList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp struct {
	value *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) Get() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) Set(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp {
	return &NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
