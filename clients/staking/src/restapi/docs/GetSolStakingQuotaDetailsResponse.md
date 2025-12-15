# GetSolStakingQuotaDetailsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LeftStakingPersonalQuota** | Pointer to **string** |  | [optional] 
**LeftRedemptionPersonalQuota** | Pointer to **string** |  | [optional] 
**MinStakeAmount** | Pointer to **string** |  | [optional] 
**MinRedeemAmount** | Pointer to **string** |  | [optional] 
**RedeemPeriod** | Pointer to **int64** |  | [optional] 
**Stakeable** | Pointer to **bool** |  | [optional] 
**Redeemable** | Pointer to **bool** |  | [optional] 
**SoldOut** | Pointer to **bool** |  | [optional] 
**CommissionFee** | Pointer to **string** |  | [optional] 
**NextEpochTime** | Pointer to **int64** |  | [optional] 
**Calculating** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetSolStakingQuotaDetailsResponse

`func NewGetSolStakingQuotaDetailsResponse() *GetSolStakingQuotaDetailsResponse`

NewGetSolStakingQuotaDetailsResponse instantiates a new GetSolStakingQuotaDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSolStakingQuotaDetailsResponseWithDefaults

`func NewGetSolStakingQuotaDetailsResponseWithDefaults() *GetSolStakingQuotaDetailsResponse`

NewGetSolStakingQuotaDetailsResponseWithDefaults instantiates a new GetSolStakingQuotaDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeftStakingPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) GetLeftStakingPersonalQuota() string`

GetLeftStakingPersonalQuota returns the LeftStakingPersonalQuota field if non-nil, zero value otherwise.

### GetLeftStakingPersonalQuotaOk

`func (o *GetSolStakingQuotaDetailsResponse) GetLeftStakingPersonalQuotaOk() (*string, bool)`

GetLeftStakingPersonalQuotaOk returns a tuple with the LeftStakingPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftStakingPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) SetLeftStakingPersonalQuota(v string)`

SetLeftStakingPersonalQuota sets LeftStakingPersonalQuota field to given value.

### HasLeftStakingPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) HasLeftStakingPersonalQuota() bool`

HasLeftStakingPersonalQuota returns a boolean if a field has been set.

### GetLeftRedemptionPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) GetLeftRedemptionPersonalQuota() string`

GetLeftRedemptionPersonalQuota returns the LeftRedemptionPersonalQuota field if non-nil, zero value otherwise.

### GetLeftRedemptionPersonalQuotaOk

`func (o *GetSolStakingQuotaDetailsResponse) GetLeftRedemptionPersonalQuotaOk() (*string, bool)`

GetLeftRedemptionPersonalQuotaOk returns a tuple with the LeftRedemptionPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftRedemptionPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) SetLeftRedemptionPersonalQuota(v string)`

SetLeftRedemptionPersonalQuota sets LeftRedemptionPersonalQuota field to given value.

### HasLeftRedemptionPersonalQuota

`func (o *GetSolStakingQuotaDetailsResponse) HasLeftRedemptionPersonalQuota() bool`

HasLeftRedemptionPersonalQuota returns a boolean if a field has been set.

### GetMinStakeAmount

`func (o *GetSolStakingQuotaDetailsResponse) GetMinStakeAmount() string`

GetMinStakeAmount returns the MinStakeAmount field if non-nil, zero value otherwise.

### GetMinStakeAmountOk

`func (o *GetSolStakingQuotaDetailsResponse) GetMinStakeAmountOk() (*string, bool)`

GetMinStakeAmountOk returns a tuple with the MinStakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakeAmount

`func (o *GetSolStakingQuotaDetailsResponse) SetMinStakeAmount(v string)`

SetMinStakeAmount sets MinStakeAmount field to given value.

### HasMinStakeAmount

`func (o *GetSolStakingQuotaDetailsResponse) HasMinStakeAmount() bool`

HasMinStakeAmount returns a boolean if a field has been set.

### GetMinRedeemAmount

`func (o *GetSolStakingQuotaDetailsResponse) GetMinRedeemAmount() string`

GetMinRedeemAmount returns the MinRedeemAmount field if non-nil, zero value otherwise.

### GetMinRedeemAmountOk

`func (o *GetSolStakingQuotaDetailsResponse) GetMinRedeemAmountOk() (*string, bool)`

GetMinRedeemAmountOk returns a tuple with the MinRedeemAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedeemAmount

`func (o *GetSolStakingQuotaDetailsResponse) SetMinRedeemAmount(v string)`

SetMinRedeemAmount sets MinRedeemAmount field to given value.

### HasMinRedeemAmount

`func (o *GetSolStakingQuotaDetailsResponse) HasMinRedeemAmount() bool`

HasMinRedeemAmount returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetSolStakingQuotaDetailsResponse) GetRedeemPeriod() int64`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetSolStakingQuotaDetailsResponse) GetRedeemPeriodOk() (*int64, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetSolStakingQuotaDetailsResponse) SetRedeemPeriod(v int64)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetSolStakingQuotaDetailsResponse) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetStakeable

`func (o *GetSolStakingQuotaDetailsResponse) GetStakeable() bool`

GetStakeable returns the Stakeable field if non-nil, zero value otherwise.

### GetStakeableOk

`func (o *GetSolStakingQuotaDetailsResponse) GetStakeableOk() (*bool, bool)`

GetStakeableOk returns a tuple with the Stakeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeable

`func (o *GetSolStakingQuotaDetailsResponse) SetStakeable(v bool)`

SetStakeable sets Stakeable field to given value.

### HasStakeable

`func (o *GetSolStakingQuotaDetailsResponse) HasStakeable() bool`

HasStakeable returns a boolean if a field has been set.

### GetRedeemable

`func (o *GetSolStakingQuotaDetailsResponse) GetRedeemable() bool`

GetRedeemable returns the Redeemable field if non-nil, zero value otherwise.

### GetRedeemableOk

`func (o *GetSolStakingQuotaDetailsResponse) GetRedeemableOk() (*bool, bool)`

GetRedeemableOk returns a tuple with the Redeemable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemable

`func (o *GetSolStakingQuotaDetailsResponse) SetRedeemable(v bool)`

SetRedeemable sets Redeemable field to given value.

### HasRedeemable

`func (o *GetSolStakingQuotaDetailsResponse) HasRedeemable() bool`

HasRedeemable returns a boolean if a field has been set.

### GetSoldOut

`func (o *GetSolStakingQuotaDetailsResponse) GetSoldOut() bool`

GetSoldOut returns the SoldOut field if non-nil, zero value otherwise.

### GetSoldOutOk

`func (o *GetSolStakingQuotaDetailsResponse) GetSoldOutOk() (*bool, bool)`

GetSoldOutOk returns a tuple with the SoldOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldOut

`func (o *GetSolStakingQuotaDetailsResponse) SetSoldOut(v bool)`

SetSoldOut sets SoldOut field to given value.

### HasSoldOut

`func (o *GetSolStakingQuotaDetailsResponse) HasSoldOut() bool`

HasSoldOut returns a boolean if a field has been set.

### GetCommissionFee

`func (o *GetSolStakingQuotaDetailsResponse) GetCommissionFee() string`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *GetSolStakingQuotaDetailsResponse) GetCommissionFeeOk() (*string, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *GetSolStakingQuotaDetailsResponse) SetCommissionFee(v string)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *GetSolStakingQuotaDetailsResponse) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.

### GetNextEpochTime

`func (o *GetSolStakingQuotaDetailsResponse) GetNextEpochTime() int64`

GetNextEpochTime returns the NextEpochTime field if non-nil, zero value otherwise.

### GetNextEpochTimeOk

`func (o *GetSolStakingQuotaDetailsResponse) GetNextEpochTimeOk() (*int64, bool)`

GetNextEpochTimeOk returns a tuple with the NextEpochTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEpochTime

`func (o *GetSolStakingQuotaDetailsResponse) SetNextEpochTime(v int64)`

SetNextEpochTime sets NextEpochTime field to given value.

### HasNextEpochTime

`func (o *GetSolStakingQuotaDetailsResponse) HasNextEpochTime() bool`

HasNextEpochTime returns a boolean if a field has been set.

### GetCalculating

`func (o *GetSolStakingQuotaDetailsResponse) GetCalculating() bool`

GetCalculating returns the Calculating field if non-nil, zero value otherwise.

### GetCalculatingOk

`func (o *GetSolStakingQuotaDetailsResponse) GetCalculatingOk() (*bool, bool)`

GetCalculatingOk returns a tuple with the Calculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculating

`func (o *GetSolStakingQuotaDetailsResponse) SetCalculating(v bool)`

SetCalculating sets Calculating field to given value.

### HasCalculating

`func (o *GetSolStakingQuotaDetailsResponse) HasCalculating() bool`

HasCalculating returns a boolean if a field has been set.


[[Back to README]](../README.md)


