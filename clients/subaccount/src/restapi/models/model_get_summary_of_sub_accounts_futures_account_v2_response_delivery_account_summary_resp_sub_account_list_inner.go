/*
Binance Sub Account REST API

OpenAPI Specification for the Binance Sub Account REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner{}

// GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner struct for GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner
type GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner struct {
	Email                 *string `json:"email,omitempty"`
	TotalMarginBalance    *string `json:"totalMarginBalance,omitempty"`
	TotalUnrealizedProfit *string `json:"totalUnrealizedProfit,omitempty"`
	TotalWalletBalance    *string `json:"totalWalletBalance,omitempty"`
	Asset                 *string `json:"asset,omitempty"`
	AdditionalProperties  map[string]interface{}
}

type _GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner{}
	return &this
}

// NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInnerWithDefaults instantiates a new GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInnerWithDefaults() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner {
	this := GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetEmail() string {
	if o == nil || common.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetEmailOk() (*string, bool) {
	if o == nil || common.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) HasEmail() bool {
	if o != nil && !common.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) SetEmail(v string) {
	o.Email = &v
}

// GetTotalMarginBalance returns the TotalMarginBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalMarginBalance() string {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		var ret string
		return ret
	}
	return *o.TotalMarginBalance
}

// GetTotalMarginBalanceOk returns a tuple with the TotalMarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalMarginBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalMarginBalance) {
		return nil, false
	}
	return o.TotalMarginBalance, true
}

// HasTotalMarginBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) HasTotalMarginBalance() bool {
	if o != nil && !common.IsNil(o.TotalMarginBalance) {
		return true
	}

	return false
}

// SetTotalMarginBalance gets a reference to the given string and assigns it to the TotalMarginBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) SetTotalMarginBalance(v string) {
	o.TotalMarginBalance = &v
}

// GetTotalUnrealizedProfit returns the TotalUnrealizedProfit field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalUnrealizedProfit() string {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		var ret string
		return ret
	}
	return *o.TotalUnrealizedProfit
}

// GetTotalUnrealizedProfitOk returns a tuple with the TotalUnrealizedProfit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalUnrealizedProfitOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalUnrealizedProfit) {
		return nil, false
	}
	return o.TotalUnrealizedProfit, true
}

// HasTotalUnrealizedProfit returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) HasTotalUnrealizedProfit() bool {
	if o != nil && !common.IsNil(o.TotalUnrealizedProfit) {
		return true
	}

	return false
}

// SetTotalUnrealizedProfit gets a reference to the given string and assigns it to the TotalUnrealizedProfit field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) SetTotalUnrealizedProfit(v string) {
	o.TotalUnrealizedProfit = &v
}

// GetTotalWalletBalance returns the TotalWalletBalance field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalWalletBalance() string {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		var ret string
		return ret
	}
	return *o.TotalWalletBalance
}

// GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetTotalWalletBalanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalWalletBalance) {
		return nil, false
	}
	return o.TotalWalletBalance, true
}

// HasTotalWalletBalance returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) HasTotalWalletBalance() bool {
	if o != nil && !common.IsNil(o.TotalWalletBalance) {
		return true
	}

	return false
}

// SetTotalWalletBalance gets a reference to the given string and assigns it to the TotalWalletBalance field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) SetTotalWalletBalance(v string) {
	o.TotalWalletBalance = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) SetAsset(v string) {
	o.Asset = &v
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !common.IsNil(o.TotalMarginBalance) {
		toSerialize["totalMarginBalance"] = o.TotalMarginBalance
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

func (o *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) UnmarshalJSON(data []byte) (err error) {
	varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner := _GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner{}

	err = json.Unmarshal(data, &varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner)

	if err != nil {
		return err
	}

	*o = GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner(varGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "totalMarginBalance")
		delete(additionalProperties, "totalUnrealizedProfit")
		delete(additionalProperties, "totalWalletBalance")
		delete(additionalProperties, "asset")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner struct {
	value *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner
	isSet bool
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) Get() *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner {
	return v.value
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) Set(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner(val *GetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner {
	return &NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner{value: val, isSet: true}
}

func (v NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSummaryOfSubAccountsFuturesAccountV2ResponseDeliveryAccountSummaryRespSubAccountListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
