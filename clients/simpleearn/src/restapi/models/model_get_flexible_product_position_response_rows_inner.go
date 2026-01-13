/*
Binance Simple Earn REST API

OpenAPI Specification for the Binance Simple Earn REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the GetFlexibleProductPositionResponseRowsInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetFlexibleProductPositionResponseRowsInner{}

// GetFlexibleProductPositionResponseRowsInner struct for GetFlexibleProductPositionResponseRowsInner
type GetFlexibleProductPositionResponseRowsInner struct {
	TotalAmount                    *string                                                              `json:"totalAmount,omitempty"`
	TierAnnualPercentageRate       *GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate `json:"tierAnnualPercentageRate,omitempty"`
	LatestAnnualPercentageRate     *string                                                              `json:"latestAnnualPercentageRate,omitempty"`
	YesterdayAirdropPercentageRate *string                                                              `json:"yesterdayAirdropPercentageRate,omitempty"`
	Asset                          *string                                                              `json:"asset,omitempty"`
	AirDropAsset                   *string                                                              `json:"airDropAsset,omitempty"`
	CanRedeem                      *bool                                                                `json:"canRedeem,omitempty"`
	CollateralAmount               *string                                                              `json:"collateralAmount,omitempty"`
	ProductId                      *string                                                              `json:"productId,omitempty"`
	YesterdayRealTimeRewards       *string                                                              `json:"yesterdayRealTimeRewards,omitempty"`
	CumulativeBonusRewards         *string                                                              `json:"cumulativeBonusRewards,omitempty"`
	CumulativeRealTimeRewards      *string                                                              `json:"cumulativeRealTimeRewards,omitempty"`
	CumulativeTotalRewards         *string                                                              `json:"cumulativeTotalRewards,omitempty"`
	AutoSubscribe                  *bool                                                                `json:"autoSubscribe,omitempty"`
	AdditionalProperties           map[string]interface{}
}

type _GetFlexibleProductPositionResponseRowsInner GetFlexibleProductPositionResponseRowsInner

// NewGetFlexibleProductPositionResponseRowsInner instantiates a new GetFlexibleProductPositionResponseRowsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFlexibleProductPositionResponseRowsInner() *GetFlexibleProductPositionResponseRowsInner {
	this := GetFlexibleProductPositionResponseRowsInner{}
	return &this
}

// NewGetFlexibleProductPositionResponseRowsInnerWithDefaults instantiates a new GetFlexibleProductPositionResponseRowsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFlexibleProductPositionResponseRowsInnerWithDefaults() *GetFlexibleProductPositionResponseRowsInner {
	this := GetFlexibleProductPositionResponseRowsInner{}
	return &this
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetTotalAmount() string {
	if o == nil || common.IsNil(o.TotalAmount) {
		var ret string
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetTotalAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasTotalAmount() bool {
	if o != nil && !common.IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given string and assigns it to the TotalAmount field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetTotalAmount(v string) {
	o.TotalAmount = &v
}

// GetTierAnnualPercentageRate returns the TierAnnualPercentageRate field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetTierAnnualPercentageRate() GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate {
	if o == nil || common.IsNil(o.TierAnnualPercentageRate) {
		var ret GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate
		return ret
	}
	return *o.TierAnnualPercentageRate
}

// GetTierAnnualPercentageRateOk returns a tuple with the TierAnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetTierAnnualPercentageRateOk() (*GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate, bool) {
	if o == nil || common.IsNil(o.TierAnnualPercentageRate) {
		return nil, false
	}
	return o.TierAnnualPercentageRate, true
}

// HasTierAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasTierAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.TierAnnualPercentageRate) {
		return true
	}

	return false
}

// SetTierAnnualPercentageRate gets a reference to the given GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate and assigns it to the TierAnnualPercentageRate field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetTierAnnualPercentageRate(v GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate) {
	o.TierAnnualPercentageRate = &v
}

// GetLatestAnnualPercentageRate returns the LatestAnnualPercentageRate field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetLatestAnnualPercentageRate() string {
	if o == nil || common.IsNil(o.LatestAnnualPercentageRate) {
		var ret string
		return ret
	}
	return *o.LatestAnnualPercentageRate
}

// GetLatestAnnualPercentageRateOk returns a tuple with the LatestAnnualPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetLatestAnnualPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.LatestAnnualPercentageRate) {
		return nil, false
	}
	return o.LatestAnnualPercentageRate, true
}

// HasLatestAnnualPercentageRate returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasLatestAnnualPercentageRate() bool {
	if o != nil && !common.IsNil(o.LatestAnnualPercentageRate) {
		return true
	}

	return false
}

// SetLatestAnnualPercentageRate gets a reference to the given string and assigns it to the LatestAnnualPercentageRate field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetLatestAnnualPercentageRate(v string) {
	o.LatestAnnualPercentageRate = &v
}

// GetYesterdayAirdropPercentageRate returns the YesterdayAirdropPercentageRate field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetYesterdayAirdropPercentageRate() string {
	if o == nil || common.IsNil(o.YesterdayAirdropPercentageRate) {
		var ret string
		return ret
	}
	return *o.YesterdayAirdropPercentageRate
}

// GetYesterdayAirdropPercentageRateOk returns a tuple with the YesterdayAirdropPercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetYesterdayAirdropPercentageRateOk() (*string, bool) {
	if o == nil || common.IsNil(o.YesterdayAirdropPercentageRate) {
		return nil, false
	}
	return o.YesterdayAirdropPercentageRate, true
}

// HasYesterdayAirdropPercentageRate returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasYesterdayAirdropPercentageRate() bool {
	if o != nil && !common.IsNil(o.YesterdayAirdropPercentageRate) {
		return true
	}

	return false
}

// SetYesterdayAirdropPercentageRate gets a reference to the given string and assigns it to the YesterdayAirdropPercentageRate field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetYesterdayAirdropPercentageRate(v string) {
	o.YesterdayAirdropPercentageRate = &v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAsset() string {
	if o == nil || common.IsNil(o.Asset) {
		var ret string
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.Asset) {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasAsset() bool {
	if o != nil && !common.IsNil(o.Asset) {
		return true
	}

	return false
}

// SetAsset gets a reference to the given string and assigns it to the Asset field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetAsset(v string) {
	o.Asset = &v
}

// GetAirDropAsset returns the AirDropAsset field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAirDropAsset() string {
	if o == nil || common.IsNil(o.AirDropAsset) {
		var ret string
		return ret
	}
	return *o.AirDropAsset
}

// GetAirDropAssetOk returns a tuple with the AirDropAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAirDropAssetOk() (*string, bool) {
	if o == nil || common.IsNil(o.AirDropAsset) {
		return nil, false
	}
	return o.AirDropAsset, true
}

// HasAirDropAsset returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasAirDropAsset() bool {
	if o != nil && !common.IsNil(o.AirDropAsset) {
		return true
	}

	return false
}

// SetAirDropAsset gets a reference to the given string and assigns it to the AirDropAsset field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetAirDropAsset(v string) {
	o.AirDropAsset = &v
}

// GetCanRedeem returns the CanRedeem field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCanRedeem() bool {
	if o == nil || common.IsNil(o.CanRedeem) {
		var ret bool
		return ret
	}
	return *o.CanRedeem
}

// GetCanRedeemOk returns a tuple with the CanRedeem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCanRedeemOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanRedeem) {
		return nil, false
	}
	return o.CanRedeem, true
}

// HasCanRedeem returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasCanRedeem() bool {
	if o != nil && !common.IsNil(o.CanRedeem) {
		return true
	}

	return false
}

// SetCanRedeem gets a reference to the given bool and assigns it to the CanRedeem field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetCanRedeem(v bool) {
	o.CanRedeem = &v
}

// GetCollateralAmount returns the CollateralAmount field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCollateralAmount() string {
	if o == nil || common.IsNil(o.CollateralAmount) {
		var ret string
		return ret
	}
	return *o.CollateralAmount
}

// GetCollateralAmountOk returns a tuple with the CollateralAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCollateralAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.CollateralAmount) {
		return nil, false
	}
	return o.CollateralAmount, true
}

// HasCollateralAmount returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasCollateralAmount() bool {
	if o != nil && !common.IsNil(o.CollateralAmount) {
		return true
	}

	return false
}

// SetCollateralAmount gets a reference to the given string and assigns it to the CollateralAmount field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetCollateralAmount(v string) {
	o.CollateralAmount = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetProductId() string {
	if o == nil || common.IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetProductIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasProductId() bool {
	if o != nil && !common.IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetProductId(v string) {
	o.ProductId = &v
}

// GetYesterdayRealTimeRewards returns the YesterdayRealTimeRewards field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetYesterdayRealTimeRewards() string {
	if o == nil || common.IsNil(o.YesterdayRealTimeRewards) {
		var ret string
		return ret
	}
	return *o.YesterdayRealTimeRewards
}

// GetYesterdayRealTimeRewardsOk returns a tuple with the YesterdayRealTimeRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetYesterdayRealTimeRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.YesterdayRealTimeRewards) {
		return nil, false
	}
	return o.YesterdayRealTimeRewards, true
}

// HasYesterdayRealTimeRewards returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasYesterdayRealTimeRewards() bool {
	if o != nil && !common.IsNil(o.YesterdayRealTimeRewards) {
		return true
	}

	return false
}

// SetYesterdayRealTimeRewards gets a reference to the given string and assigns it to the YesterdayRealTimeRewards field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetYesterdayRealTimeRewards(v string) {
	o.YesterdayRealTimeRewards = &v
}

// GetCumulativeBonusRewards returns the CumulativeBonusRewards field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeBonusRewards() string {
	if o == nil || common.IsNil(o.CumulativeBonusRewards) {
		var ret string
		return ret
	}
	return *o.CumulativeBonusRewards
}

// GetCumulativeBonusRewardsOk returns a tuple with the CumulativeBonusRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeBonusRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumulativeBonusRewards) {
		return nil, false
	}
	return o.CumulativeBonusRewards, true
}

// HasCumulativeBonusRewards returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasCumulativeBonusRewards() bool {
	if o != nil && !common.IsNil(o.CumulativeBonusRewards) {
		return true
	}

	return false
}

// SetCumulativeBonusRewards gets a reference to the given string and assigns it to the CumulativeBonusRewards field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetCumulativeBonusRewards(v string) {
	o.CumulativeBonusRewards = &v
}

// GetCumulativeRealTimeRewards returns the CumulativeRealTimeRewards field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeRealTimeRewards() string {
	if o == nil || common.IsNil(o.CumulativeRealTimeRewards) {
		var ret string
		return ret
	}
	return *o.CumulativeRealTimeRewards
}

// GetCumulativeRealTimeRewardsOk returns a tuple with the CumulativeRealTimeRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeRealTimeRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumulativeRealTimeRewards) {
		return nil, false
	}
	return o.CumulativeRealTimeRewards, true
}

// HasCumulativeRealTimeRewards returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasCumulativeRealTimeRewards() bool {
	if o != nil && !common.IsNil(o.CumulativeRealTimeRewards) {
		return true
	}

	return false
}

// SetCumulativeRealTimeRewards gets a reference to the given string and assigns it to the CumulativeRealTimeRewards field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetCumulativeRealTimeRewards(v string) {
	o.CumulativeRealTimeRewards = &v
}

// GetCumulativeTotalRewards returns the CumulativeTotalRewards field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeTotalRewards() string {
	if o == nil || common.IsNil(o.CumulativeTotalRewards) {
		var ret string
		return ret
	}
	return *o.CumulativeTotalRewards
}

// GetCumulativeTotalRewardsOk returns a tuple with the CumulativeTotalRewards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetCumulativeTotalRewardsOk() (*string, bool) {
	if o == nil || common.IsNil(o.CumulativeTotalRewards) {
		return nil, false
	}
	return o.CumulativeTotalRewards, true
}

// HasCumulativeTotalRewards returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasCumulativeTotalRewards() bool {
	if o != nil && !common.IsNil(o.CumulativeTotalRewards) {
		return true
	}

	return false
}

// SetCumulativeTotalRewards gets a reference to the given string and assigns it to the CumulativeTotalRewards field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetCumulativeTotalRewards(v string) {
	o.CumulativeTotalRewards = &v
}

// GetAutoSubscribe returns the AutoSubscribe field value if set, zero value otherwise.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAutoSubscribe() bool {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		var ret bool
		return ret
	}
	return *o.AutoSubscribe
}

// GetAutoSubscribeOk returns a tuple with the AutoSubscribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) GetAutoSubscribeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.AutoSubscribe) {
		return nil, false
	}
	return o.AutoSubscribe, true
}

// HasAutoSubscribe returns a boolean if a field has been set.
func (o *GetFlexibleProductPositionResponseRowsInner) HasAutoSubscribe() bool {
	if o != nil && !common.IsNil(o.AutoSubscribe) {
		return true
	}

	return false
}

// SetAutoSubscribe gets a reference to the given bool and assigns it to the AutoSubscribe field.
func (o *GetFlexibleProductPositionResponseRowsInner) SetAutoSubscribe(v bool) {
	o.AutoSubscribe = &v
}

func (o GetFlexibleProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFlexibleProductPositionResponseRowsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !common.IsNil(o.TierAnnualPercentageRate) {
		toSerialize["tierAnnualPercentageRate"] = o.TierAnnualPercentageRate
	}
	if !common.IsNil(o.LatestAnnualPercentageRate) {
		toSerialize["latestAnnualPercentageRate"] = o.LatestAnnualPercentageRate
	}
	if !common.IsNil(o.YesterdayAirdropPercentageRate) {
		toSerialize["yesterdayAirdropPercentageRate"] = o.YesterdayAirdropPercentageRate
	}
	if !common.IsNil(o.Asset) {
		toSerialize["asset"] = o.Asset
	}
	if !common.IsNil(o.AirDropAsset) {
		toSerialize["airDropAsset"] = o.AirDropAsset
	}
	if !common.IsNil(o.CanRedeem) {
		toSerialize["canRedeem"] = o.CanRedeem
	}
	if !common.IsNil(o.CollateralAmount) {
		toSerialize["collateralAmount"] = o.CollateralAmount
	}
	if !common.IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !common.IsNil(o.YesterdayRealTimeRewards) {
		toSerialize["yesterdayRealTimeRewards"] = o.YesterdayRealTimeRewards
	}
	if !common.IsNil(o.CumulativeBonusRewards) {
		toSerialize["cumulativeBonusRewards"] = o.CumulativeBonusRewards
	}
	if !common.IsNil(o.CumulativeRealTimeRewards) {
		toSerialize["cumulativeRealTimeRewards"] = o.CumulativeRealTimeRewards
	}
	if !common.IsNil(o.CumulativeTotalRewards) {
		toSerialize["cumulativeTotalRewards"] = o.CumulativeTotalRewards
	}
	if !common.IsNil(o.AutoSubscribe) {
		toSerialize["autoSubscribe"] = o.AutoSubscribe
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetFlexibleProductPositionResponseRowsInner) UnmarshalJSON(data []byte) (err error) {
	varGetFlexibleProductPositionResponseRowsInner := _GetFlexibleProductPositionResponseRowsInner{}

	err = json.Unmarshal(data, &varGetFlexibleProductPositionResponseRowsInner)

	if err != nil {
		return err
	}

	*o = GetFlexibleProductPositionResponseRowsInner(varGetFlexibleProductPositionResponseRowsInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "totalAmount")
		delete(additionalProperties, "tierAnnualPercentageRate")
		delete(additionalProperties, "latestAnnualPercentageRate")
		delete(additionalProperties, "yesterdayAirdropPercentageRate")
		delete(additionalProperties, "asset")
		delete(additionalProperties, "airDropAsset")
		delete(additionalProperties, "canRedeem")
		delete(additionalProperties, "collateralAmount")
		delete(additionalProperties, "productId")
		delete(additionalProperties, "yesterdayRealTimeRewards")
		delete(additionalProperties, "cumulativeBonusRewards")
		delete(additionalProperties, "cumulativeRealTimeRewards")
		delete(additionalProperties, "cumulativeTotalRewards")
		delete(additionalProperties, "autoSubscribe")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetFlexibleProductPositionResponseRowsInner struct {
	value *GetFlexibleProductPositionResponseRowsInner
	isSet bool
}

func (v NullableGetFlexibleProductPositionResponseRowsInner) Get() *GetFlexibleProductPositionResponseRowsInner {
	return v.value
}

func (v *NullableGetFlexibleProductPositionResponseRowsInner) Set(val *GetFlexibleProductPositionResponseRowsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFlexibleProductPositionResponseRowsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFlexibleProductPositionResponseRowsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFlexibleProductPositionResponseRowsInner(val *GetFlexibleProductPositionResponseRowsInner) *NullableGetFlexibleProductPositionResponseRowsInner {
	return &NullableGetFlexibleProductPositionResponseRowsInner{value: val, isSet: true}
}

func (v NullableGetFlexibleProductPositionResponseRowsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFlexibleProductPositionResponseRowsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
