# AccountInformationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UniMMR** | Pointer to **string** |  | [optional] 
**AccountEquity** | Pointer to **string** |  | [optional] 
**ActualEquity** | Pointer to **string** |  | [optional] 
**AccountInitialMargin** | Pointer to **string** |  | [optional] 
**AccountMaintMargin** | Pointer to **string** |  | [optional] 
**AccountStatus** | Pointer to **string** |  | [optional] 
**VirtualMaxWithdrawAmount** | Pointer to **string** |  | [optional] 
**TotalAvailableBalance** | Pointer to **string** |  | [optional] 
**TotalMarginOpenLoss** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountInformationResponse

`func NewAccountInformationResponse() *AccountInformationResponse`

NewAccountInformationResponse instantiates a new AccountInformationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInformationResponseWithDefaults

`func NewAccountInformationResponseWithDefaults() *AccountInformationResponse`

NewAccountInformationResponseWithDefaults instantiates a new AccountInformationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniMMR

`func (o *AccountInformationResponse) GetUniMMR() string`

GetUniMMR returns the UniMMR field if non-nil, zero value otherwise.

### GetUniMMROk

`func (o *AccountInformationResponse) GetUniMMROk() (*string, bool)`

GetUniMMROk returns a tuple with the UniMMR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniMMR

`func (o *AccountInformationResponse) SetUniMMR(v string)`

SetUniMMR sets UniMMR field to given value.

### HasUniMMR

`func (o *AccountInformationResponse) HasUniMMR() bool`

HasUniMMR returns a boolean if a field has been set.

### GetAccountEquity

`func (o *AccountInformationResponse) GetAccountEquity() string`

GetAccountEquity returns the AccountEquity field if non-nil, zero value otherwise.

### GetAccountEquityOk

`func (o *AccountInformationResponse) GetAccountEquityOk() (*string, bool)`

GetAccountEquityOk returns a tuple with the AccountEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEquity

`func (o *AccountInformationResponse) SetAccountEquity(v string)`

SetAccountEquity sets AccountEquity field to given value.

### HasAccountEquity

`func (o *AccountInformationResponse) HasAccountEquity() bool`

HasAccountEquity returns a boolean if a field has been set.

### GetActualEquity

`func (o *AccountInformationResponse) GetActualEquity() string`

GetActualEquity returns the ActualEquity field if non-nil, zero value otherwise.

### GetActualEquityOk

`func (o *AccountInformationResponse) GetActualEquityOk() (*string, bool)`

GetActualEquityOk returns a tuple with the ActualEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEquity

`func (o *AccountInformationResponse) SetActualEquity(v string)`

SetActualEquity sets ActualEquity field to given value.

### HasActualEquity

`func (o *AccountInformationResponse) HasActualEquity() bool`

HasActualEquity returns a boolean if a field has been set.

### GetAccountInitialMargin

`func (o *AccountInformationResponse) GetAccountInitialMargin() string`

GetAccountInitialMargin returns the AccountInitialMargin field if non-nil, zero value otherwise.

### GetAccountInitialMarginOk

`func (o *AccountInformationResponse) GetAccountInitialMarginOk() (*string, bool)`

GetAccountInitialMarginOk returns a tuple with the AccountInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInitialMargin

`func (o *AccountInformationResponse) SetAccountInitialMargin(v string)`

SetAccountInitialMargin sets AccountInitialMargin field to given value.

### HasAccountInitialMargin

`func (o *AccountInformationResponse) HasAccountInitialMargin() bool`

HasAccountInitialMargin returns a boolean if a field has been set.

### GetAccountMaintMargin

`func (o *AccountInformationResponse) GetAccountMaintMargin() string`

GetAccountMaintMargin returns the AccountMaintMargin field if non-nil, zero value otherwise.

### GetAccountMaintMarginOk

`func (o *AccountInformationResponse) GetAccountMaintMarginOk() (*string, bool)`

GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMaintMargin

`func (o *AccountInformationResponse) SetAccountMaintMargin(v string)`

SetAccountMaintMargin sets AccountMaintMargin field to given value.

### HasAccountMaintMargin

`func (o *AccountInformationResponse) HasAccountMaintMargin() bool`

HasAccountMaintMargin returns a boolean if a field has been set.

### GetAccountStatus

`func (o *AccountInformationResponse) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *AccountInformationResponse) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *AccountInformationResponse) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *AccountInformationResponse) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetVirtualMaxWithdrawAmount

`func (o *AccountInformationResponse) GetVirtualMaxWithdrawAmount() string`

GetVirtualMaxWithdrawAmount returns the VirtualMaxWithdrawAmount field if non-nil, zero value otherwise.

### GetVirtualMaxWithdrawAmountOk

`func (o *AccountInformationResponse) GetVirtualMaxWithdrawAmountOk() (*string, bool)`

GetVirtualMaxWithdrawAmountOk returns a tuple with the VirtualMaxWithdrawAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMaxWithdrawAmount

`func (o *AccountInformationResponse) SetVirtualMaxWithdrawAmount(v string)`

SetVirtualMaxWithdrawAmount sets VirtualMaxWithdrawAmount field to given value.

### HasVirtualMaxWithdrawAmount

`func (o *AccountInformationResponse) HasVirtualMaxWithdrawAmount() bool`

HasVirtualMaxWithdrawAmount returns a boolean if a field has been set.

### GetTotalAvailableBalance

`func (o *AccountInformationResponse) GetTotalAvailableBalance() string`

GetTotalAvailableBalance returns the TotalAvailableBalance field if non-nil, zero value otherwise.

### GetTotalAvailableBalanceOk

`func (o *AccountInformationResponse) GetTotalAvailableBalanceOk() (*string, bool)`

GetTotalAvailableBalanceOk returns a tuple with the TotalAvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailableBalance

`func (o *AccountInformationResponse) SetTotalAvailableBalance(v string)`

SetTotalAvailableBalance sets TotalAvailableBalance field to given value.

### HasTotalAvailableBalance

`func (o *AccountInformationResponse) HasTotalAvailableBalance() bool`

HasTotalAvailableBalance returns a boolean if a field has been set.

### GetTotalMarginOpenLoss

`func (o *AccountInformationResponse) GetTotalMarginOpenLoss() string`

GetTotalMarginOpenLoss returns the TotalMarginOpenLoss field if non-nil, zero value otherwise.

### GetTotalMarginOpenLossOk

`func (o *AccountInformationResponse) GetTotalMarginOpenLossOk() (*string, bool)`

GetTotalMarginOpenLossOk returns a tuple with the TotalMarginOpenLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMarginOpenLoss

`func (o *AccountInformationResponse) SetTotalMarginOpenLoss(v string)`

SetTotalMarginOpenLoss sets TotalMarginOpenLoss field to given value.

### HasTotalMarginOpenLoss

`func (o *AccountInformationResponse) HasTotalMarginOpenLoss() bool`

HasTotalMarginOpenLoss returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountInformationResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountInformationResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountInformationResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountInformationResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


