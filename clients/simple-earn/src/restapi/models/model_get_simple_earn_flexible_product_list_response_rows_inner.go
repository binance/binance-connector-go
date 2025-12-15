/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API

API version: 1.0.0
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetSimpleEarnFlexibleProductListResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetSimpleEarnFlexibleProductListResponseRowsInner{}

// GetSimpleEarnFlexibleProductListResponseRowsInner struct for GetSimpleEarnFlexibleProductListResponseRowsInner
type GetSimpleEarnFlexibleProductListResponseRowsInner struct {
	Asset                      *string                                                              `json:"asset,omitempty"`
	LatestAnnualPercentageRate *string                                                              `json:"latestAnnualPercentageRate,omitempty"`
	TierAnnualPercentageRate   *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate `json:"tierAnnualPercentageRate,omitempty"`
	AirDropPercentageRate      *string                                                              `json:"airDropPercentageRate,omitempty"`
	CanPurchase                *bool                                                                `json:"canPurchase,omitempty"`
	CanRedeem                  *bool                                                                `json:"canRedeem,omitempty"`
	IsSoldOut                  *bool                                                                `json:"isSoldOut,omitempty"`
	Hot                        *bool                                                                `json:"hot,omitempty"`
	MinPurchaseAmount          *string                                                              `json:"minPurchaseAmount,omitempty"`
	ProductId                  *string                                                              `json:"productId,omitempty"`
	SubscriptionStartTime      *int64                                                               `json:"subscriptionStartTime,omitempty"`
	Status                     *string                                                              `json:"status,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _GetSimpleEarnFlexibleProductListResponseRowsInner GetSimpleEarnFlexibleProductListResponseRowsInner

// NewGetSimpleEarnFlexibleProductListResponseRowsInner instantiates a new GetSimpleEarnFlexibleProductListResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSimpleEarnFlexibleProductListResponseRowsInner() *GetSimpleEarnFlexibleProductListResponseRowsInner {
	this := GetSimpleEarnFlexibleProductListResponseRowsInner{}
	return &this
}

// NewGetSimpleEarnFlexibleProductListResponseRowsInnerWithDefaults instantiates a new GetSimpleEarnFlexibleProductListResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSimpleEarnFlexibleProductListResponseRowsInnerWithDefaults() *GetSimpleEarnFlexibleProductListResponseRowsInner {
	this := GetSimpleEarnFlexibleProductListResponseRowsInner{}
	return &this
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetLatestAnnualPercentageRate returns the LatestAnnualPercentageRate field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetLatestAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.LatestAnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.LatestAnnualPercentageRate
}

// GetLatestAnnualPercentageRateOk returns a tuple with the LatestAnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetLatestAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LatestAnnualPercentageRate) {
		return nil, false
	}
	return o.LatestAnnualPercentageRate, true
}

// HasLatestAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasLatestAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.LatestAnnualPercentageRate) {
		return true
	}

	return false
}

// SetLatestAnnualPercentageRate gets a reference to the given string and assigns it to the LatestAnnualPercentageRate field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetLatestAnnualPercentageRate(v string) {
	o.LatestAnnualPercentageRate = &v
}

// GetTierAnnualPercentageRate returns the TierAnnualPercentageRate field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetTierAnnualPercentageRate() GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	if o == nil || common.IsNil(o.TierAnnualPercentageRate) {
		var ret GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate
		return ret
	}
	return *o.TierAnnualPercentageRate
}

// GetTierAnnualPercentageRateOk returns a tuple with the TierAnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetTierAnnualPercentageRateOk() (*GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate, bool) {
	if o == nil || common.IsNil(o.TierAnnualPercentageRate) {
		return nil, false
	}
	return o.TierAnnualPercentageRate, true
}

// HasTierAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasTierAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.TierAnnualPercentageRate) {
		return true
	}

	return false
}

// SetTierAnnualPercentageRate gets a reference to the given GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate and assigns it to the TierAnnualPercentageRate field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetTierAnnualPercentageRate(v GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) {
	o.TierAnnualPercentageRate = &v
}

// GetAirDropPercentageRate returns the AirDropPercentageRate field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAirDropPercentageRate() string {
	if o == nil || common.IsNil(o.AirDropPercentageRate) {
		var ret string
		return ret
	}
	return *o.AirDropPercentageRate
}

// GetAirDropPercentageRateOk returns a tuple with the AirDropPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAirDropPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.AirDropPercentageRate) {
		return nil, false
	}
	return o.AirDropPercentageRate, true
}

// HasAirDropPercentageRate returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasAirDropPercentageRate() bool {
	if o != nil && !common.IsNil(o.AirDropPercentageRate) {
		return true
	}

	return false
}

// SetAirDropPercentageRate gets a reference to the given string and assigns it to the AirDropPercentageRate field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetAirDropPercentageRate(v string) {
	o.AirDropPercentageRate = &v
}

// GetCanPurchase returns the CanPurchase field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanPurchase() bool {
	if o == nil || common.IsNil(o.CanPurchase) {
		var ret bool
		return ret
	}
	return *o.CanPurchase
}

// GetCanPurchaseOk returns a tuple with the CanPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanPurchaseOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanPurchase) {
		return nil, false
	}
	return o.CanPurchase, true
}

// HasCanPurchase returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasCanPurchase() bool {
	if o != nil && !common.IsNil(o.CanPurchase) {
		return true
	}

	return false
}

// SetCanPurchase gets a reference to the given bool and assigns it to the CanPurchase field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetCanPurchase(v bool) {
	o.CanPurchase = &v
}

// GetCanRedeem returns the CanRedeem field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanRedeem() bool {
	if o == nil || common.IsNil(o.CanRedeem) {
		var ret bool
		return ret
	}
	return *o.CanRedeem
}

// GetCanRedeemOk returns a tuple with the CanRedeem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanRedeemOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanRedeem) {
		return nil, false
	}
	return o.CanRedeem, true
}

// HasCanRedeem returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasCanRedeem() bool {
	if o != nil && !common.IsNil(o.CanRedeem) {
		return true
	}

	return false
}

// SetCanRedeem gets a reference to the given bool and assigns it to the CanRedeem field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetCanRedeem(v bool) {
	o.CanRedeem = &v
}

// GetIsSoldOut returns the IsSoldOut field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetIsSoldOut() bool {
	if o == nil || common.IsNil(o.IsSoldOut) {
		var ret bool
		return ret
	}
	return *o.IsSoldOut
}

// GetIsSoldOutOk returns a tuple with the IsSoldOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetIsSoldOutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsSoldOut) {
		return nil, false
	}
	return o.IsSoldOut, true
}

// HasIsSoldOut returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasIsSoldOut() bool {
	if o != nil && !common.IsNil(o.IsSoldOut) {
		return true
	}

	return false
}

// SetIsSoldOut gets a reference to the given bool and assigns it to the IsSoldOut field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetIsSoldOut(v bool) {
	o.IsSoldOut = &v
}

// GetHot returns the Hot field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetHot() bool {
	if o == nil || common.IsNil(o.Hot) {
		var ret bool
		return ret
	}
	return *o.Hot
}

// GetHotOk returns a tuple with the Hot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetHotOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Hot) {
		return nil, false
	}
	return o.Hot, true
}

// HasHot returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasHot() bool {
	if o != nil && !common.IsNil(o.Hot) {
		return true
	}

	return false
}

// SetHot gets a reference to the given bool and assigns it to the Hot field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetHot(v bool) {
	o.Hot = &v
}

// GetMinPurchaseAmount returns the MinPurchaseAmount field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetMinPurchaseAmount() string {
	if o == nil || common.IsNil(o.MinPurchaseAmount) {
		var ret string
		return ret
	}
	return *o.MinPurchaseAmount
}

// GetMinPurchaseAmountOk returns a tuple with the MinPurchaseAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetMinPurchaseAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinPurchaseAmount) {
		return nil, false
	}
	return o.MinPurchaseAmount, true
}

// HasMinPurchaseAmount returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasMinPurchaseAmount() bool {
	if o != nil && !common.IsNil(o.MinPurchaseAmount) {
		return true
	}

	return false
}

// SetMinPurchaseAmount gets a reference to the given string and assigns it to the MinPurchaseAmount field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetMinPurchaseAmount(v string) {
	o.MinPurchaseAmount = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetProductId() string {
	if o == nil || common.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetProductIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasProductId() bool {
	if o != nil && !common.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetSubscriptionStartTime returns the SubscriptionStartTime field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetSubscriptionStartTime() int64 {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		var ret int64
		return ret
	}
	return *o.SubscriptionStartTime
}

// GetSubscriptionStartTimeOk returns a tuple with the SubscriptionStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetSubscriptionStartTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.SubscriptionStartTime) {
		return nil, false
	}
	return o.SubscriptionStartTime, true
}

// HasSubscriptionStartTime returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasSubscriptionStartTime() bool {
	if o != nil && !common.IsNil(o.SubscriptionStartTime) {
		return true
	}

	return false
}

// SetSubscriptionStartTime gets a reference to the given int64 and assigns it to the SubscriptionStartTime field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetSubscriptionStartTime(v int64) {
	o.SubscriptionStartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetStatus() string {
	if o == nil || common.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetStatusOk() (*string, bool) {
	if o == nil || common.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasStatus() bool {
	if o != nil && !common.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetStatus(v string) {
	o.Status = &v
}

func (o GetSimpleEarnFlexibleProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSimpleEarnFlexibleProductListResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.LatestAnnualPercentageRate) {
		toSerialize["latestAnnualPercentageRate"] = o.LatestAnnualPercentageRate
	}
	if !common.IsNil(o.TierAnnualPercentageRate) {
		toSerialize["tierAnnualPercentageRate"] = o.TierAnnualPercentageRate
	}
	if !common.IsNil(o.AirDropPercentageRate) {
		toSerialize["airDropPercentageRate"] = o.AirDropPercentageRate
	}
	if !common.IsNil(o.CanPurchase) {
		toSerialize["canPurchase"] = o.CanPurchase
	}
	if !common.IsNil(o.CanRedeem) {
		toSerialize["canRedeem"] = o.CanRedeem
	}
	if !common.IsNil(o.IsSoldOut) {
		toSerialize["isSoldOut"] = o.IsSoldOut
	}
	if !common.IsNil(o.Hot) {
		toSerialize["hot"] = o.Hot
	}
	if !common.IsNil(o.MinPurchaseAmount) {
		toSerialize["minPurchaseAmount"] = o.MinPurchaseAmount
	}
	if !common.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !common.IsNil(o.SubscriptionStartTime) {
		toSerialize["subscriptionStartTime"] = o.SubscriptionStartTime
	}
	if !common.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetSimpleEarnFlexibleProductListResponseRowsInner := _GetSimpleEarnFlexibleProductListResponseRowsInner{}

	err = json.Unmarshal(data, &varGetSimpleEarnFlexibleProductListResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetSimpleEarnFlexibleProductListResponseRowsInner(varGetSimpleEarnFlexibleProductListResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asset")
		delete(additionalProperties, "latestAnnualPercentageRate")
		delete(additionalProperties, "tierAnnualPercentageRate")
		delete(additionalProperties, "airDropPercentageRate")
		delete(additionalProperties, "canPurchase")
		delete(additionalProperties, "canRedeem")
		delete(additionalProperties, "isSoldOut")
		delete(additionalProperties, "hot")
		delete(additionalProperties, "minPurchaseAmount")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "subscriptionStartTime")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetSimpleEarnFlexibleProductListResponseRowsInner struct {
	value *GetSimpleEarnFlexibleProductListResponseRowsInner
	isSet bool
}

func (v NullableGetSimpleEarnFlexibleProductListResponseRowsInner) Get() *GetSimpleEarnFlexibleProductListResponseRowsInner {
	return v.value
}

func (v *NullableGetSimpleEarnFlexibleProductListResponseRowsInner) Set(val *GetSimpleEarnFlexibleProductListResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSimpleEarnFlexibleProductListResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSimpleEarnFlexibleProductListResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSimpleEarnFlexibleProductListResponseRowsInner(val *GetSimpleEarnFlexibleProductListResponseRowsInner) *NullableGetSimpleEarnFlexibleProductListResponseRowsInner {
	return &NullableGetSimpleEarnFlexibleProductListResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetSimpleEarnFlexibleProductListResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSimpleEarnFlexibleProductListResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
