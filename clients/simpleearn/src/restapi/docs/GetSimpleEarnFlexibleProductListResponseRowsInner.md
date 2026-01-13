# GetSimpleEarnFlexibleProductListResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**LatestAnnualPercentageRate** | Pointer to **string** |  | [optional] 
**TierAnnualPercentageRate** | Pointer to [**GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate**](GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate.md) |  | [optional] 
**AirDropPercentageRate** | Pointer to **string** |  | [optional] 
**CanPurchase** | Pointer to **bool** |  | [optional] 
**CanRedeem** | Pointer to **bool** |  | [optional] 
**IsSoldOut** | Pointer to **bool** |  | [optional] 
**Hot** | Pointer to **bool** |  | [optional] 
**MinPurchaseAmount** | Pointer to **string** |  | [optional] 
**ProductId** | Pointer to **string** |  | [optional] 
**SubscriptionStartTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetSimpleEarnFlexibleProductListResponseRowsInner

`func NewGetSimpleEarnFlexibleProductListResponseRowsInner() *GetSimpleEarnFlexibleProductListResponseRowsInner`

NewGetSimpleEarnFlexibleProductListResponseRowsInner instantiates a new GetSimpleEarnFlexibleProductListResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSimpleEarnFlexibleProductListResponseRowsInnerWithDefaults

`func NewGetSimpleEarnFlexibleProductListResponseRowsInnerWithDefaults() *GetSimpleEarnFlexibleProductListResponseRowsInner`

NewGetSimpleEarnFlexibleProductListResponseRowsInnerWithDefaults instantiates a new GetSimpleEarnFlexibleProductListResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetLatestAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetLatestAnnualPercentageRate() string`

GetLatestAnnualPercentageRate returns the LatestAnnualPercentageRate field if non-nil, zero value otherwise.

### GetLatestAnnualPercentageRateOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetLatestAnnualPercentageRateOk() (*string, bool)`

GetLatestAnnualPercentageRateOk returns a tuple with the LatestAnnualPercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetLatestAnnualPercentageRate(v string)`

SetLatestAnnualPercentageRate sets LatestAnnualPercentageRate field to given value.

### HasLatestAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasLatestAnnualPercentageRate() bool`

HasLatestAnnualPercentageRate returns a boolean if a field has been set.

### GetTierAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetTierAnnualPercentageRate() GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate`

GetTierAnnualPercentageRate returns the TierAnnualPercentageRate field if non-nil, zero value otherwise.

### GetTierAnnualPercentageRateOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetTierAnnualPercentageRateOk() (*GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate, bool)`

GetTierAnnualPercentageRateOk returns a tuple with the TierAnnualPercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetTierAnnualPercentageRate(v GetFlexibleProductPositionResponseRowsInnerTierAnnualPercentageRate)`

SetTierAnnualPercentageRate sets TierAnnualPercentageRate field to given value.

### HasTierAnnualPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasTierAnnualPercentageRate() bool`

HasTierAnnualPercentageRate returns a boolean if a field has been set.

### GetAirDropPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAirDropPercentageRate() string`

GetAirDropPercentageRate returns the AirDropPercentageRate field if non-nil, zero value otherwise.

### GetAirDropPercentageRateOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetAirDropPercentageRateOk() (*string, bool)`

GetAirDropPercentageRateOk returns a tuple with the AirDropPercentageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDropPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetAirDropPercentageRate(v string)`

SetAirDropPercentageRate sets AirDropPercentageRate field to given value.

### HasAirDropPercentageRate

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasAirDropPercentageRate() bool`

HasAirDropPercentageRate returns a boolean if a field has been set.

### GetCanPurchase

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanPurchase() bool`

GetCanPurchase returns the CanPurchase field if non-nil, zero value otherwise.

### GetCanPurchaseOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanPurchaseOk() (*bool, bool)`

GetCanPurchaseOk returns a tuple with the CanPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPurchase

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetCanPurchase(v bool)`

SetCanPurchase sets CanPurchase field to given value.

### HasCanPurchase

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasCanPurchase() bool`

HasCanPurchase returns a boolean if a field has been set.

### GetCanRedeem

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanRedeem() bool`

GetCanRedeem returns the CanRedeem field if non-nil, zero value otherwise.

### GetCanRedeemOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetCanRedeemOk() (*bool, bool)`

GetCanRedeemOk returns a tuple with the CanRedeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRedeem

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetCanRedeem(v bool)`

SetCanRedeem sets CanRedeem field to given value.

### HasCanRedeem

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasCanRedeem() bool`

HasCanRedeem returns a boolean if a field has been set.

### GetIsSoldOut

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetIsSoldOut() bool`

GetIsSoldOut returns the IsSoldOut field if non-nil, zero value otherwise.

### GetIsSoldOutOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetIsSoldOutOk() (*bool, bool)`

GetIsSoldOutOk returns a tuple with the IsSoldOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoldOut

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetIsSoldOut(v bool)`

SetIsSoldOut sets IsSoldOut field to given value.

### HasIsSoldOut

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasIsSoldOut() bool`

HasIsSoldOut returns a boolean if a field has been set.

### GetHot

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetHot() bool`

GetHot returns the Hot field if non-nil, zero value otherwise.

### GetHotOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetHotOk() (*bool, bool)`

GetHotOk returns a tuple with the Hot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHot

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetHot(v bool)`

SetHot sets Hot field to given value.

### HasHot

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasHot() bool`

HasHot returns a boolean if a field has been set.

### GetMinPurchaseAmount

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetMinPurchaseAmount() string`

GetMinPurchaseAmount returns the MinPurchaseAmount field if non-nil, zero value otherwise.

### GetMinPurchaseAmountOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetMinPurchaseAmountOk() (*string, bool)`

GetMinPurchaseAmountOk returns a tuple with the MinPurchaseAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPurchaseAmount

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetMinPurchaseAmount(v string)`

SetMinPurchaseAmount sets MinPurchaseAmount field to given value.

### HasMinPurchaseAmount

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasMinPurchaseAmount() bool`

HasMinPurchaseAmount returns a boolean if a field has been set.

### GetProductId

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSubscriptionStartTime

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetSubscriptionStartTime() int64`

GetSubscriptionStartTime returns the SubscriptionStartTime field if non-nil, zero value otherwise.

### GetSubscriptionStartTimeOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetSubscriptionStartTimeOk() (*int64, bool)`

GetSubscriptionStartTimeOk returns a tuple with the SubscriptionStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStartTime

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetSubscriptionStartTime(v int64)`

SetSubscriptionStartTime sets SubscriptionStartTime field to given value.

### HasSubscriptionStartTime

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasSubscriptionStartTime() bool`

HasSubscriptionStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSimpleEarnFlexibleProductListResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


