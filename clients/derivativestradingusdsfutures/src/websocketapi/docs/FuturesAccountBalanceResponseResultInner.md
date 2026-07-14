# FuturesAccountBalanceResponseResultInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AccountAlias** | Pointer to **string** | unique account code | [optional] 
**Asset** | Pointer to **string** | asset name | [optional] 
**Balance** | Pointer to **string** | wallet balance | [optional] 
**CrossWalletBalance** | Pointer to **string** | crossed wallet balance | [optional] 
**CrossUnPnl** | Pointer to **string** | unrealized profit of crossed positions | [optional] 
**AvailableBalance** | Pointer to **string** | available balance | [optional] 
**MaxWithdrawAmount** | Pointer to **string** | maximum amount for transfer out | [optional] 
**MarginAvailable** | Pointer to **bool** | whether the asset can be used as margin in Multi-Assets mode | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFuturesAccountBalanceResponseResultInner

`func NewFuturesAccountBalanceResponseResultInner() *FuturesAccountBalanceResponseResultInner`

NewFuturesAccountBalanceResponseResultInner instantiates a new FuturesAccountBalanceResponseResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuturesAccountBalanceResponseResultInnerWithDefaults

`func NewFuturesAccountBalanceResponseResultInnerWithDefaults() *FuturesAccountBalanceResponseResultInner`

NewFuturesAccountBalanceResponseResultInnerWithDefaults instantiates a new FuturesAccountBalanceResponseResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAlias

`func (o *FuturesAccountBalanceResponseResultInner) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *FuturesAccountBalanceResponseResultInner) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *FuturesAccountBalanceResponseResultInner) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *FuturesAccountBalanceResponseResultInner) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAsset

`func (o *FuturesAccountBalanceResponseResultInner) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *FuturesAccountBalanceResponseResultInner) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *FuturesAccountBalanceResponseResultInner) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *FuturesAccountBalanceResponseResultInner) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetBalance

`func (o *FuturesAccountBalanceResponseResultInner) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *FuturesAccountBalanceResponseResultInner) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *FuturesAccountBalanceResponseResultInner) SetBalance(v string)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *FuturesAccountBalanceResponseResultInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCrossWalletBalance

`func (o *FuturesAccountBalanceResponseResultInner) GetCrossWalletBalance() string`

GetCrossWalletBalance returns the CrossWalletBalance field if non-nil, zero value otherwise.

### GetCrossWalletBalanceOk

`func (o *FuturesAccountBalanceResponseResultInner) GetCrossWalletBalanceOk() (*string, bool)`

GetCrossWalletBalanceOk returns a tuple with the CrossWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossWalletBalance

`func (o *FuturesAccountBalanceResponseResultInner) SetCrossWalletBalance(v string)`

SetCrossWalletBalance sets CrossWalletBalance field to given value.

### HasCrossWalletBalance

`func (o *FuturesAccountBalanceResponseResultInner) HasCrossWalletBalance() bool`

HasCrossWalletBalance returns a boolean if a field has been set.

### GetCrossUnPnl

`func (o *FuturesAccountBalanceResponseResultInner) GetCrossUnPnl() string`

GetCrossUnPnl returns the CrossUnPnl field if non-nil, zero value otherwise.

### GetCrossUnPnlOk

`func (o *FuturesAccountBalanceResponseResultInner) GetCrossUnPnlOk() (*string, bool)`

GetCrossUnPnlOk returns a tuple with the CrossUnPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossUnPnl

`func (o *FuturesAccountBalanceResponseResultInner) SetCrossUnPnl(v string)`

SetCrossUnPnl sets CrossUnPnl field to given value.

### HasCrossUnPnl

`func (o *FuturesAccountBalanceResponseResultInner) HasCrossUnPnl() bool`

HasCrossUnPnl returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *FuturesAccountBalanceResponseResultInner) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *FuturesAccountBalanceResponseResultInner) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *FuturesAccountBalanceResponseResultInner) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *FuturesAccountBalanceResponseResultInner) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetMaxWithdrawAmount

`func (o *FuturesAccountBalanceResponseResultInner) GetMaxWithdrawAmount() string`

GetMaxWithdrawAmount returns the MaxWithdrawAmount field if non-nil, zero value otherwise.

### GetMaxWithdrawAmountOk

`func (o *FuturesAccountBalanceResponseResultInner) GetMaxWithdrawAmountOk() (*string, bool)`

GetMaxWithdrawAmountOk returns a tuple with the MaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWithdrawAmount

`func (o *FuturesAccountBalanceResponseResultInner) SetMaxWithdrawAmount(v string)`

SetMaxWithdrawAmount sets MaxWithdrawAmount field to given value.

### HasMaxWithdrawAmount

`func (o *FuturesAccountBalanceResponseResultInner) HasMaxWithdrawAmount() bool`

HasMaxWithdrawAmount returns a boolean if a field has been set.

### GetMarginAvailable

`func (o *FuturesAccountBalanceResponseResultInner) GetMarginAvailable() bool`

GetMarginAvailable returns the MarginAvailable field if non-nil, zero value otherwise.

### GetMarginAvailableOk

`func (o *FuturesAccountBalanceResponseResultInner) GetMarginAvailableOk() (*bool, bool)`

GetMarginAvailableOk returns a tuple with the MarginAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAvailable

`func (o *FuturesAccountBalanceResponseResultInner) SetMarginAvailable(v bool)`

SetMarginAvailable sets MarginAvailable field to given value.

### HasMarginAvailable

`func (o *FuturesAccountBalanceResponseResultInner) HasMarginAvailable() bool`

HasMarginAvailable returns a boolean if a field has been set.

### GetUpdateTime

`func (o *FuturesAccountBalanceResponseResultInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *FuturesAccountBalanceResponseResultInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *FuturesAccountBalanceResponseResultInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *FuturesAccountBalanceResponseResultInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


