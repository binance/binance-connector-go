/*
Sub Account REST API

Create and manage sub-accounts, control permissions, and transfer assets via the Sub Account API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner{}

// GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner struct for GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner
type GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner struct {
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

type _GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner

// NewGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner instantiates a new GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner() *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInnerWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInnerWithDefaults() *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetEmail(v string) {
	o.Email = &v
}

// GetTotalInitialMargin returns the TotalInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalInitialMargin() string {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalInitialMargin
}

// GetTotalInitialMarginOk returns a tuple with the TotalInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalInitialMargin) {
		return nil, false
	}
	return o.TotalInitialMargin, true
}

// HasTotalInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalInitialMargin) {
		return true
	}

	return false
}

// SetTotalInitialMargin gets a reference to the given string and assigns it to the TotalInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalInitialMargin(v string) {
	o.TotalInitialMargin = &v
}

// GetTotalMaintenanceMargin returns the TotalMaintenanceMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalMaintenanceMargin() string {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		var ret string
		return ret
	}
	return *o.TotalMaintenanceMargin
}

// GetTotalMaintenanceMarginOk returns a tuple with the TotalMaintenanceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalMaintenanceMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMaintenanceMargin) {
		return nil, false
	}
	return o.TotalMaintenanceMargin, true
}

// HasTotalMaintenanceMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalMaintenanceMargin() bool {
	if o != nil && !common.IsNil(o.TotalMaintenanceMargin) {
		return true
	}

	return false
}

// SetTotalMaintenanceMargin gets a reference to the given string and assigns it to the TotalMaintenanceMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalMaintenanceMargin(v string) {
	o.TotalMaintenanceMargin = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalOpenOrderInitialMargin returns the TotalOpenOrderInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalOpenOrderInitialMargin() string {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalOpenOrderInitialMargin
}

// GetTotalOpenOrderInitialMarginOk returns a tuple with the TotalOpenOrderInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalOpenOrderInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalOpenOrderInitialMargin) {
		return nil, false
	}
	return o.TotalOpenOrderInitialMargin, true
}

// HasTotalOpenOrderInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalOpenOrderInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalOpenOrderInitialMargin) {
		return true
	}

	return false
}

// SetTotalOpenOrderInitialMargin gets a reference to the given string and assigns it to the TotalOpenOrderInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalOpenOrderInitialMargin(v string) {
	o.TotalOpenOrderInitialMargin = &v
}

// GetTotalPositionInitialMargin returns the TotalPositionInitialMargin field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalPositionInitialMargin() string {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		var ret string
		return ret
	}
	return *o.TotalPositionInitialMargin
}

// GetTotalPositionInitialMarginOk returns a tuple with the TotalPositionInitialMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalPositionInitialMarginOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalPositionInitialMargin) {
		return nil, false
	}
	return o.TotalPositionInitialMargin, true
}

// HasTotalPositionInitialMargin returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalPositionInitialMargin() bool {
	if o != nil && !common.IsNil(o.TotalPositionInitialMargin) {
		return true
	}

	return false
}

// SetTotalPositionInitialMargin gets a reference to the given string and assigns it to the TotalPositionInitialMargin field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalPositionInitialMargin(v string) {
	o.TotalPositionInitialMargin = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) SetAsset(v string) {
	o.Asset = &v
}

func (o GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) ToMap() (map[string]interface{}, error) {
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

func (o *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner := _GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner(varGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner)

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

type NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner struct {
	value *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) Get() *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) Set(val *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner(val *GetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) *NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner {
	return &NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountResponseSubAccountListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
