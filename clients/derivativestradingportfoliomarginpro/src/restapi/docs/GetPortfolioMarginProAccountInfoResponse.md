# GetPortfolioMarginProAccountInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**UniMMR** | Pointer to **string** | Classic Portfolio margin account maintenance margin rate | [optional] 
**AccountEquity** | Pointer to **string** | Account equity, unit：USD | [optional] 
**ActualEquity** | Pointer to **string** | Actual equity, unit：USD | [optional] 
**AccountMaintMargin** | Pointer to **string** | Classic Portfolio margin account maintenance margin, unit：USD | [optional] 
**AccountInitialMargin** | Pointer to **string** | Ignored for PM PRO and PM PRO SPAN | [optional] 
**TotalAvailableBalance** | Pointer to **string** | Ignored for PM PRO and PM PRO SPAN | [optional] 
**AccountStatus** | Pointer to **string** | Classic Portfolio margin account status:\&quot;NORMAL\&quot;, \&quot;MARGIN_CALL\&quot;, \&quot;SUPPLY_MARGIN\&quot;, \&quot;REDUCE_ONLY\&quot;, \&quot;ACTIVE_LIQUIDATION\&quot;, \&quot;FORCE_LIQUIDATION\&quot;, \&quot;BANKRUPTED\&quot; | [optional] 
**AccountType** | Pointer to **string** | PM_1 for PM PRO, PM_2 for PM, PM_3 for PM PRO SPAN | [optional] 

## Methods

### NewGetPortfolioMarginProAccountInfoResponse

`func NewGetPortfolioMarginProAccountInfoResponse() *GetPortfolioMarginProAccountInfoResponse`

NewGetPortfolioMarginProAccountInfoResponse instantiates a new GetPortfolioMarginProAccountInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPortfolioMarginProAccountInfoResponseWithDefaults

`func NewGetPortfolioMarginProAccountInfoResponseWithDefaults() *GetPortfolioMarginProAccountInfoResponse`

NewGetPortfolioMarginProAccountInfoResponseWithDefaults instantiates a new GetPortfolioMarginProAccountInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUniMMR

`func (o *GetPortfolioMarginProAccountInfoResponse) GetUniMMR() string`

GetUniMMR returns the UniMMR field if non-nil, zero value otherwise.

### GetUniMMROk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetUniMMROk() (*string, bool)`

GetUniMMROk returns a tuple with the UniMMR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniMMR

`func (o *GetPortfolioMarginProAccountInfoResponse) SetUniMMR(v string)`

SetUniMMR sets UniMMR field to given value.

### HasUniMMR

`func (o *GetPortfolioMarginProAccountInfoResponse) HasUniMMR() bool`

HasUniMMR returns a boolean if a field has been set.

### GetAccountEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountEquity() string`

GetAccountEquity returns the AccountEquity field if non-nil, zero value otherwise.

### GetAccountEquityOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountEquityOk() (*string, bool)`

GetAccountEquityOk returns a tuple with the AccountEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountEquity(v string)`

SetAccountEquity sets AccountEquity field to given value.

### HasAccountEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountEquity() bool`

HasAccountEquity returns a boolean if a field has been set.

### GetActualEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) GetActualEquity() string`

GetActualEquity returns the ActualEquity field if non-nil, zero value otherwise.

### GetActualEquityOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetActualEquityOk() (*string, bool)`

GetActualEquityOk returns a tuple with the ActualEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) SetActualEquity(v string)`

SetActualEquity sets ActualEquity field to given value.

### HasActualEquity

`func (o *GetPortfolioMarginProAccountInfoResponse) HasActualEquity() bool`

HasActualEquity returns a boolean if a field has been set.

### GetAccountMaintMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountMaintMargin() string`

GetAccountMaintMargin returns the AccountMaintMargin field if non-nil, zero value otherwise.

### GetAccountMaintMarginOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountMaintMarginOk() (*string, bool)`

GetAccountMaintMarginOk returns a tuple with the AccountMaintMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMaintMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountMaintMargin(v string)`

SetAccountMaintMargin sets AccountMaintMargin field to given value.

### HasAccountMaintMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountMaintMargin() bool`

HasAccountMaintMargin returns a boolean if a field has been set.

### GetAccountInitialMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountInitialMargin() string`

GetAccountInitialMargin returns the AccountInitialMargin field if non-nil, zero value otherwise.

### GetAccountInitialMarginOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountInitialMarginOk() (*string, bool)`

GetAccountInitialMarginOk returns a tuple with the AccountInitialMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInitialMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountInitialMargin(v string)`

SetAccountInitialMargin sets AccountInitialMargin field to given value.

### HasAccountInitialMargin

`func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountInitialMargin() bool`

HasAccountInitialMargin returns a boolean if a field has been set.

### GetTotalAvailableBalance

`func (o *GetPortfolioMarginProAccountInfoResponse) GetTotalAvailableBalance() string`

GetTotalAvailableBalance returns the TotalAvailableBalance field if non-nil, zero value otherwise.

### GetTotalAvailableBalanceOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetTotalAvailableBalanceOk() (*string, bool)`

GetTotalAvailableBalanceOk returns a tuple with the TotalAvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAvailableBalance

`func (o *GetPortfolioMarginProAccountInfoResponse) SetTotalAvailableBalance(v string)`

SetTotalAvailableBalance sets TotalAvailableBalance field to given value.

### HasTotalAvailableBalance

`func (o *GetPortfolioMarginProAccountInfoResponse) HasTotalAvailableBalance() bool`

HasTotalAvailableBalance returns a boolean if a field has been set.

### GetAccountStatus

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetAccountType

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetPortfolioMarginProAccountInfoResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetPortfolioMarginProAccountInfoResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetPortfolioMarginProAccountInfoResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to README]](../README.md)


