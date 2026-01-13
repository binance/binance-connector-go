# GetBfusdQuotaDetailsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FastRedemptionQuota** | Pointer to [**GetBfusdQuotaDetailsResponseFastRedemptionQuota**](GetBfusdQuotaDetailsResponseFastRedemptionQuota.md) |  | [optional] 
**StandardRedemptionQuota** | Pointer to [**GetBfusdQuotaDetailsResponseStandardRedemptionQuota**](GetBfusdQuotaDetailsResponseStandardRedemptionQuota.md) |  | [optional] 
**SubscribeEnable** | Pointer to **bool** |  | [optional] 
**RedeemEnable** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetBfusdQuotaDetailsResponse

`func NewGetBfusdQuotaDetailsResponse() *GetBfusdQuotaDetailsResponse`

NewGetBfusdQuotaDetailsResponse instantiates a new GetBfusdQuotaDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBfusdQuotaDetailsResponseWithDefaults

`func NewGetBfusdQuotaDetailsResponseWithDefaults() *GetBfusdQuotaDetailsResponse`

NewGetBfusdQuotaDetailsResponseWithDefaults instantiates a new GetBfusdQuotaDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFastRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) GetFastRedemptionQuota() GetBfusdQuotaDetailsResponseFastRedemptionQuota`

GetFastRedemptionQuota returns the FastRedemptionQuota field if non-nil, zero value otherwise.

### GetFastRedemptionQuotaOk

`func (o *GetBfusdQuotaDetailsResponse) GetFastRedemptionQuotaOk() (*GetBfusdQuotaDetailsResponseFastRedemptionQuota, bool)`

GetFastRedemptionQuotaOk returns a tuple with the FastRedemptionQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) SetFastRedemptionQuota(v GetBfusdQuotaDetailsResponseFastRedemptionQuota)`

SetFastRedemptionQuota sets FastRedemptionQuota field to given value.

### HasFastRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) HasFastRedemptionQuota() bool`

HasFastRedemptionQuota returns a boolean if a field has been set.

### GetStandardRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) GetStandardRedemptionQuota() GetBfusdQuotaDetailsResponseStandardRedemptionQuota`

GetStandardRedemptionQuota returns the StandardRedemptionQuota field if non-nil, zero value otherwise.

### GetStandardRedemptionQuotaOk

`func (o *GetBfusdQuotaDetailsResponse) GetStandardRedemptionQuotaOk() (*GetBfusdQuotaDetailsResponseStandardRedemptionQuota, bool)`

GetStandardRedemptionQuotaOk returns a tuple with the StandardRedemptionQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) SetStandardRedemptionQuota(v GetBfusdQuotaDetailsResponseStandardRedemptionQuota)`

SetStandardRedemptionQuota sets StandardRedemptionQuota field to given value.

### HasStandardRedemptionQuota

`func (o *GetBfusdQuotaDetailsResponse) HasStandardRedemptionQuota() bool`

HasStandardRedemptionQuota returns a boolean if a field has been set.

### GetSubscribeEnable

`func (o *GetBfusdQuotaDetailsResponse) GetSubscribeEnable() bool`

GetSubscribeEnable returns the SubscribeEnable field if non-nil, zero value otherwise.

### GetSubscribeEnableOk

`func (o *GetBfusdQuotaDetailsResponse) GetSubscribeEnableOk() (*bool, bool)`

GetSubscribeEnableOk returns a tuple with the SubscribeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeEnable

`func (o *GetBfusdQuotaDetailsResponse) SetSubscribeEnable(v bool)`

SetSubscribeEnable sets SubscribeEnable field to given value.

### HasSubscribeEnable

`func (o *GetBfusdQuotaDetailsResponse) HasSubscribeEnable() bool`

HasSubscribeEnable returns a boolean if a field has been set.

### GetRedeemEnable

`func (o *GetBfusdQuotaDetailsResponse) GetRedeemEnable() bool`

GetRedeemEnable returns the RedeemEnable field if non-nil, zero value otherwise.

### GetRedeemEnableOk

`func (o *GetBfusdQuotaDetailsResponse) GetRedeemEnableOk() (*bool, bool)`

GetRedeemEnableOk returns a tuple with the RedeemEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemEnable

`func (o *GetBfusdQuotaDetailsResponse) SetRedeemEnable(v bool)`

SetRedeemEnable sets RedeemEnable field to given value.

### HasRedeemEnable

`func (o *GetBfusdQuotaDetailsResponse) HasRedeemEnable() bool`

HasRedeemEnable returns a boolean if a field has been set.


[[Back to README]](../README.md)


