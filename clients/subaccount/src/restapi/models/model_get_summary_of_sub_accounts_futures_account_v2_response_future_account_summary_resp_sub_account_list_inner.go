/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner{}

// GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner struct for GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner
type GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner struct {
	Email                       *string `json:"email,omitempty"`
	TotalInitialMargin          *string `json:"totalInitialMargin,omitempty"`
	TotalMaintenanceMargin      *string `json:"totalMaintenanceMargin,omitempty"`
	TotalMarginBalance          *string `json:"totalMarginBalance,omitempty"`
	TotalOpenOrderInitialMargin *string `json:"totalOpenOrderInitialMargin,omitempty"`
	TotalPositionInitialMargin  *string `json:"totalPositionInitialMargin,omitempty"`
	TotalUnrealizedProfit       *string `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance          *string `json:"totalWalletBalance,omitempty"`
	Asset                       *string `json:"asset,omitempty"`
	AdditionalProperties        map[string]interface{}
}

type _GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInnerWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInnerWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetEmail(v string) {
	o.Email = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) SetAsset(v string) {
	o.Asset = &v
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner := _GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner(varGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "totalInitialMargin")
		delete(additionalProperties, "totalMaintenanceMargin")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalOpenOrderInitialMargin")
		delete(additionalProperties, "totalPositionInitialMargin")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "asset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner struct {
	value *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) Get() *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) Set(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner {
	return &NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseFutureAccountSummaryRespSubAccountListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
