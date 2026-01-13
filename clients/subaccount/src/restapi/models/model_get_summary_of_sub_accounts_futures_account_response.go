/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSummaryOfSubAccountsFuturesAccountResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountResponse{}

// GetSummaryOfSubAccountsFuturesAccountResponse struct for GetSummaryOfSubAccountsFuturesAccountResponse
type GetSummaryOfSubAccountsFuturesAccountResponse struct {
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

type _GetSummaryOfSubAccountsFuturesAccountResponse GetSummaryOfSubAccountsFuturesAccountResponse

// NewGetSummaryOfSubAccountsFuturesAccountResponse instantiates a new GetSummaryOfSubAccountsFuturesAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountResponse() *GetSummaryOfSubAccountsFuturesAccountResponse {
	this := GetSummaryOfSubAccountsFuturesAccountResponse{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountResponseWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountResponseWithDefaults() *GetSummaryOfSubAccountsFuturesAccountResponse {
	this := GetSummaryOfSubAccountsFuturesAccountResponse{}
	return &this
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetAsset(v string) {
	o.Asset = &v
}

// GetSubAccountList returns the SubAccountList field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetSubAccountList() []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	if o == nil || common.IsNil(o.SubAccountList) {
		var ret []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner
		return ret
	}
	return o.SubAccountList
}

// GetSubAccountListOk returns a tuple with the SubAccountList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) GetSubAccountListOk() ([]GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner, bool) {
	if o == nil || common.IsNil(o.SubAccountList) {
		return nil, false
	}
	return o.SubAccountList, true
}

// HasSubAccountList returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) HasSubAccountList() bool {
	if o != nil && !common.IsNil(o.SubAccountList) {
		return true
	}

	return false
}

// SetSubAccountList gets a reference to the given []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner and assigns it to the SubAccountList field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponse) SetSubAccountList(v []GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) {
	o.SubAccountList = v
}

func (o GetSummaryOfSubAccountsFuturesAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetSummaryOfSubAccountsFuturesAccountResponse) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountResponse := _GetSummaryOfSubAccountsFuturesAccountResponse{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountResponse)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountResponse(varGetSummaryOfSubAccountsFuturesAccountResponse)

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

type NullableGetSummaryOfSubAccountsFuturesAccountResponse struct {
	value *GetSummaryOfSubAccountsFuturesAccountResponse
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponse) Get() *GetSummaryOfSubAccountsFuturesAccountResponse {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponse) Set(val *GetSummaryOfSubAccountsFuturesAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountResponse(val *GetSummaryOfSubAccountsFuturesAccountResponse) *NullableGetSummaryOfSubAccountsFuturesAccountResponse {
	return &NullableGetSummaryOfSubAccountsFuturesAccountResponse{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
