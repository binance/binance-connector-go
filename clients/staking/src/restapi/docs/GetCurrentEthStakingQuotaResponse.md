# GetCurrentEthStakingQuotaResponse

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
**CommissionFee** | Pointer to **string** |  | [optional] 
**Calculating** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetCurrentEthStakingQuotaResponse

`func NewGetCurrentEthStakingQuotaResponse() *GetCurrentEthStakingQuotaResponse`

NewGetCurrentEthStakingQuotaResponse instantiates a new GetCurrentEthStakingQuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCurrentEthStakingQuotaResponseWithDefaults

`func NewGetCurrentEthStakingQuotaResponseWithDefaults() *GetCurrentEthStakingQuotaResponse`

NewGetCurrentEthStakingQuotaResponseWithDefaults instantiates a new GetCurrentEthStakingQuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeftStakingPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) GetLeftStakingPersonalQuota() string`

GetLeftStakingPersonalQuota returns the LeftStakingPersonalQuota field if non-nil, zero value otherwise.

### GetLeftStakingPersonalQuotaOk

`func (o *GetCurrentEthStakingQuotaResponse) GetLeftStakingPersonalQuotaOk() (*string, bool)`

GetLeftStakingPersonalQuotaOk returns a tuple with the LeftStakingPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftStakingPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) SetLeftStakingPersonalQuota(v string)`

SetLeftStakingPersonalQuota sets LeftStakingPersonalQuota field to given value.

### HasLeftStakingPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) HasLeftStakingPersonalQuota() bool`

HasLeftStakingPersonalQuota returns a boolean if a field has been set.

### GetLeftRedemptionPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) GetLeftRedemptionPersonalQuota() string`

GetLeftRedemptionPersonalQuota returns the LeftRedemptionPersonalQuota field if non-nil, zero value otherwise.

### GetLeftRedemptionPersonalQuotaOk

`func (o *GetCurrentEthStakingQuotaResponse) GetLeftRedemptionPersonalQuotaOk() (*string, bool)`

GetLeftRedemptionPersonalQuotaOk returns a tuple with the LeftRedemptionPersonalQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftRedemptionPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) SetLeftRedemptionPersonalQuota(v string)`

SetLeftRedemptionPersonalQuota sets LeftRedemptionPersonalQuota field to given value.

### HasLeftRedemptionPersonalQuota

`func (o *GetCurrentEthStakingQuotaResponse) HasLeftRedemptionPersonalQuota() bool`

HasLeftRedemptionPersonalQuota returns a boolean if a field has been set.

### GetMinStakeAmount

`func (o *GetCurrentEthStakingQuotaResponse) GetMinStakeAmount() string`

GetMinStakeAmount returns the MinStakeAmount field if non-nil, zero value otherwise.

### GetMinStakeAmountOk

`func (o *GetCurrentEthStakingQuotaResponse) GetMinStakeAmountOk() (*string, bool)`

GetMinStakeAmountOk returns a tuple with the MinStakeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStakeAmount

`func (o *GetCurrentEthStakingQuotaResponse) SetMinStakeAmount(v string)`

SetMinStakeAmount sets MinStakeAmount field to given value.

### HasMinStakeAmount

`func (o *GetCurrentEthStakingQuotaResponse) HasMinStakeAmount() bool`

HasMinStakeAmount returns a boolean if a field has been set.

### GetMinRedeemAmount

`func (o *GetCurrentEthStakingQuotaResponse) GetMinRedeemAmount() string`

GetMinRedeemAmount returns the MinRedeemAmount field if non-nil, zero value otherwise.

### GetMinRedeemAmountOk

`func (o *GetCurrentEthStakingQuotaResponse) GetMinRedeemAmountOk() (*string, bool)`

GetMinRedeemAmountOk returns a tuple with the MinRedeemAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRedeemAmount

`func (o *GetCurrentEthStakingQuotaResponse) SetMinRedeemAmount(v string)`

SetMinRedeemAmount sets MinRedeemAmount field to given value.

### HasMinRedeemAmount

`func (o *GetCurrentEthStakingQuotaResponse) HasMinRedeemAmount() bool`

HasMinRedeemAmount returns a boolean if a field has been set.

### GetRedeemPeriod

`func (o *GetCurrentEthStakingQuotaResponse) GetRedeemPeriod() int64`

GetRedeemPeriod returns the RedeemPeriod field if non-nil, zero value otherwise.

### GetRedeemPeriodOk

`func (o *GetCurrentEthStakingQuotaResponse) GetRedeemPeriodOk() (*int64, bool)`

GetRedeemPeriodOk returns a tuple with the RedeemPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemPeriod

`func (o *GetCurrentEthStakingQuotaResponse) SetRedeemPeriod(v int64)`

SetRedeemPeriod sets RedeemPeriod field to given value.

### HasRedeemPeriod

`func (o *GetCurrentEthStakingQuotaResponse) HasRedeemPeriod() bool`

HasRedeemPeriod returns a boolean if a field has been set.

### GetStakeable

`func (o *GetCurrentEthStakingQuotaResponse) GetStakeable() bool`

GetStakeable returns the Stakeable field if non-nil, zero value otherwise.

### GetStakeableOk

`func (o *GetCurrentEthStakingQuotaResponse) GetStakeableOk() (*bool, bool)`

GetStakeableOk returns a tuple with the Stakeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeable

`func (o *GetCurrentEthStakingQuotaResponse) SetStakeable(v bool)`

SetStakeable sets Stakeable field to given value.

### HasStakeable

`func (o *GetCurrentEthStakingQuotaResponse) HasStakeable() bool`

HasStakeable returns a boolean if a field has been set.

### GetRedeemable

`func (o *GetCurrentEthStakingQuotaResponse) GetRedeemable() bool`

GetRedeemable returns the Redeemable field if non-nil, zero value otherwise.

### GetRedeemableOk

`func (o *GetCurrentEthStakingQuotaResponse) GetRedeemableOk() (*bool, bool)`

GetRedeemableOk returns a tuple with the Redeemable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemable

`func (o *GetCurrentEthStakingQuotaResponse) SetRedeemable(v bool)`

SetRedeemable sets Redeemable field to given value.

### HasRedeemable

`func (o *GetCurrentEthStakingQuotaResponse) HasRedeemable() bool`

HasRedeemable returns a boolean if a field has been set.

### GetCommissionFee

`func (o *GetCurrentEthStakingQuotaResponse) GetCommissionFee() string`

GetCommissionFee returns the CommissionFee field if non-nil, zero value otherwise.

### GetCommissionFeeOk

`func (o *GetCurrentEthStakingQuotaResponse) GetCommissionFeeOk() (*string, bool)`

GetCommissionFeeOk returns a tuple with the CommissionFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionFee

`func (o *GetCurrentEthStakingQuotaResponse) SetCommissionFee(v string)`

SetCommissionFee sets CommissionFee field to given value.

### HasCommissionFee

`func (o *GetCurrentEthStakingQuotaResponse) HasCommissionFee() bool`

HasCommissionFee returns a boolean if a field has been set.

### GetCalculating

`func (o *GetCurrentEthStakingQuotaResponse) GetCalculating() bool`

GetCalculating returns the Calculating field if non-nil, zero value otherwise.

### GetCalculatingOk

`func (o *GetCurrentEthStakingQuotaResponse) GetCalculatingOk() (*bool, bool)`

GetCalculatingOk returns a tuple with the Calculating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculating

`func (o *GetCurrentEthStakingQuotaResponse) SetCalculating(v bool)`

SetCalculating sets Calculating field to given value.

### HasCalculating

`func (o *GetCurrentEthStakingQuotaResponse) HasCalculating() bool`

HasCalculating returns a boolean if a field has been set.


[[Back to README]](../README.md)


