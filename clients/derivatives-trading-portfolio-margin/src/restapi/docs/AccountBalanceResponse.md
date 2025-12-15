# AccountBalanceResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**TotalWalletBalance** | Pointer to **string** |  | [optional] 
**CrossMarginBorrowed** | Pointer to **string** |  | [optional] 
**CrossMarginFree** | Pointer to **string** |  | [optional] 
**CrossMarginInterest** | Pointer to **string** |  | [optional] 
**CrossMarginLocked** | Pointer to **string** |  | [optional] 
**UmWalletBalance** | Pointer to **string** |  | [optional] 
**UmUnrealizedPNL** | Pointer to **string** |  | [optional] 
**CmWalletBalance** | Pointer to **string** |  | [optional] 
**CmUnrealizedPNL** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**NegativeBalance** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountBalanceResponse

`func NewAccountBalanceResponse() *AccountBalanceResponse`

NewAccountBalanceResponse instantiates a new AccountBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceResponseWithDefaults

`func NewAccountBalanceResponseWithDefaults() *AccountBalanceResponse`

NewAccountBalanceResponseWithDefaults instantiates a new AccountBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *AccountBalanceResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *AccountBalanceResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *AccountBalanceResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *AccountBalanceResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetTotalWalletBalance

`func (o *AccountBalanceResponse) GetTotalWalletBalance() string`

GetTotalWalletBalance returns the TotalWalletBalance field if non-nil, zero value otherwise.

### GetTotalWalletBalanceOk

`func (o *AccountBalanceResponse) GetTotalWalletBalanceOk() (*string, bool)`

GetTotalWalletBalanceOk returns a tuple with the TotalWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWalletBalance

`func (o *AccountBalanceResponse) SetTotalWalletBalance(v string)`

SetTotalWalletBalance sets TotalWalletBalance field to given value.

### HasTotalWalletBalance

`func (o *AccountBalanceResponse) HasTotalWalletBalance() bool`

HasTotalWalletBalance returns a boolean if a field has been set.

### GetCrossMarginBorrowed

`func (o *AccountBalanceResponse) GetCrossMarginBorrowed() string`

GetCrossMarginBorrowed returns the CrossMarginBorrowed field if non-nil, zero value otherwise.

### GetCrossMarginBorrowedOk

`func (o *AccountBalanceResponse) GetCrossMarginBorrowedOk() (*string, bool)`

GetCrossMarginBorrowedOk returns a tuple with the CrossMarginBorrowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginBorrowed

`func (o *AccountBalanceResponse) SetCrossMarginBorrowed(v string)`

SetCrossMarginBorrowed sets CrossMarginBorrowed field to given value.

### HasCrossMarginBorrowed

`func (o *AccountBalanceResponse) HasCrossMarginBorrowed() bool`

HasCrossMarginBorrowed returns a boolean if a field has been set.

### GetCrossMarginFree

`func (o *AccountBalanceResponse) GetCrossMarginFree() string`

GetCrossMarginFree returns the CrossMarginFree field if non-nil, zero value otherwise.

### GetCrossMarginFreeOk

`func (o *AccountBalanceResponse) GetCrossMarginFreeOk() (*string, bool)`

GetCrossMarginFreeOk returns a tuple with the CrossMarginFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginFree

`func (o *AccountBalanceResponse) SetCrossMarginFree(v string)`

SetCrossMarginFree sets CrossMarginFree field to given value.

### HasCrossMarginFree

`func (o *AccountBalanceResponse) HasCrossMarginFree() bool`

HasCrossMarginFree returns a boolean if a field has been set.

### GetCrossMarginInterest

`func (o *AccountBalanceResponse) GetCrossMarginInterest() string`

GetCrossMarginInterest returns the CrossMarginInterest field if non-nil, zero value otherwise.

### GetCrossMarginInterestOk

`func (o *AccountBalanceResponse) GetCrossMarginInterestOk() (*string, bool)`

GetCrossMarginInterestOk returns a tuple with the CrossMarginInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginInterest

`func (o *AccountBalanceResponse) SetCrossMarginInterest(v string)`

SetCrossMarginInterest sets CrossMarginInterest field to given value.

### HasCrossMarginInterest

`func (o *AccountBalanceResponse) HasCrossMarginInterest() bool`

HasCrossMarginInterest returns a boolean if a field has been set.

### GetCrossMarginLocked

`func (o *AccountBalanceResponse) GetCrossMarginLocked() string`

GetCrossMarginLocked returns the CrossMarginLocked field if non-nil, zero value otherwise.

### GetCrossMarginLockedOk

`func (o *AccountBalanceResponse) GetCrossMarginLockedOk() (*string, bool)`

GetCrossMarginLockedOk returns a tuple with the CrossMarginLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossMarginLocked

`func (o *AccountBalanceResponse) SetCrossMarginLocked(v string)`

SetCrossMarginLocked sets CrossMarginLocked field to given value.

### HasCrossMarginLocked

`func (o *AccountBalanceResponse) HasCrossMarginLocked() bool`

HasCrossMarginLocked returns a boolean if a field has been set.

### GetUmWalletBalance

`func (o *AccountBalanceResponse) GetUmWalletBalance() string`

GetUmWalletBalance returns the UmWalletBalance field if non-nil, zero value otherwise.

### GetUmWalletBalanceOk

`func (o *AccountBalanceResponse) GetUmWalletBalanceOk() (*string, bool)`

GetUmWalletBalanceOk returns a tuple with the UmWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmWalletBalance

`func (o *AccountBalanceResponse) SetUmWalletBalance(v string)`

SetUmWalletBalance sets UmWalletBalance field to given value.

### HasUmWalletBalance

`func (o *AccountBalanceResponse) HasUmWalletBalance() bool`

HasUmWalletBalance returns a boolean if a field has been set.

### GetUmUnrealizedPNL

`func (o *AccountBalanceResponse) GetUmUnrealizedPNL() string`

GetUmUnrealizedPNL returns the UmUnrealizedPNL field if non-nil, zero value otherwise.

### GetUmUnrealizedPNLOk

`func (o *AccountBalanceResponse) GetUmUnrealizedPNLOk() (*string, bool)`

GetUmUnrealizedPNLOk returns a tuple with the UmUnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUmUnrealizedPNL

`func (o *AccountBalanceResponse) SetUmUnrealizedPNL(v string)`

SetUmUnrealizedPNL sets UmUnrealizedPNL field to given value.

### HasUmUnrealizedPNL

`func (o *AccountBalanceResponse) HasUmUnrealizedPNL() bool`

HasUmUnrealizedPNL returns a boolean if a field has been set.

### GetCmWalletBalance

`func (o *AccountBalanceResponse) GetCmWalletBalance() string`

GetCmWalletBalance returns the CmWalletBalance field if non-nil, zero value otherwise.

### GetCmWalletBalanceOk

`func (o *AccountBalanceResponse) GetCmWalletBalanceOk() (*string, bool)`

GetCmWalletBalanceOk returns a tuple with the CmWalletBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmWalletBalance

`func (o *AccountBalanceResponse) SetCmWalletBalance(v string)`

SetCmWalletBalance sets CmWalletBalance field to given value.

### HasCmWalletBalance

`func (o *AccountBalanceResponse) HasCmWalletBalance() bool`

HasCmWalletBalance returns a boolean if a field has been set.

### GetCmUnrealizedPNL

`func (o *AccountBalanceResponse) GetCmUnrealizedPNL() string`

GetCmUnrealizedPNL returns the CmUnrealizedPNL field if non-nil, zero value otherwise.

### GetCmUnrealizedPNLOk

`func (o *AccountBalanceResponse) GetCmUnrealizedPNLOk() (*string, bool)`

GetCmUnrealizedPNLOk returns a tuple with the CmUnrealizedPNL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmUnrealizedPNL

`func (o *AccountBalanceResponse) SetCmUnrealizedPNL(v string)`

SetCmUnrealizedPNL sets CmUnrealizedPNL field to given value.

### HasCmUnrealizedPNL

`func (o *AccountBalanceResponse) HasCmUnrealizedPNL() bool`

HasCmUnrealizedPNL returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountBalanceResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountBalanceResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountBalanceResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountBalanceResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetNegativeBalance

`func (o *AccountBalanceResponse) GetNegativeBalance() string`

GetNegativeBalance returns the NegativeBalance field if non-nil, zero value otherwise.

### GetNegativeBalanceOk

`func (o *AccountBalanceResponse) GetNegativeBalanceOk() (*string, bool)`

GetNegativeBalanceOk returns a tuple with the NegativeBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeBalance

`func (o *AccountBalanceResponse) SetNegativeBalance(v string)`

SetNegativeBalance sets NegativeBalance field to given value.

### HasNegativeBalance

`func (o *AccountBalanceResponse) HasNegativeBalance() bool`

HasNegativeBalance returns a boolean if a field has been set.


[[Back to README]](../README.md)


