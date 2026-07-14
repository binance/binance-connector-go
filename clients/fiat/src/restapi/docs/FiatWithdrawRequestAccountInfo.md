# FiatWithdrawRequestAccountInfo

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Your destination bank account number is required to receive the withdrawal. In Argentina, this will be your CBU/CVU; in Mexico, it will be your CLABE. | 
**Agency** | Pointer to **string** | Bank agency code. If contains a hyphen (e.g. &#x60;123-4&#x60;), enter &#x60;123&#x60; only. | [optional] 
**BankCodeForPix** | Pointer to **string** | Bank code used for PIX routing. | [optional] 
**AccountType** | Pointer to **string** | Account type, e.g. &#x60;current&#x60; (Checking Account), &#x60;saving&#x60; (Savings Account), etc. | [optional] 

## Methods

### NewFiatWithdrawRequestAccountInfo

`func NewFiatWithdrawRequestAccountInfo(accountNumber string, ) *FiatWithdrawRequestAccountInfo`

NewFiatWithdrawRequestAccountInfo instantiates a new FiatWithdrawRequestAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatWithdrawRequestAccountInfoWithDefaults

`func NewFiatWithdrawRequestAccountInfoWithDefaults() *FiatWithdrawRequestAccountInfo`

NewFiatWithdrawRequestAccountInfoWithDefaults instantiates a new FiatWithdrawRequestAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *FiatWithdrawRequestAccountInfo) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *FiatWithdrawRequestAccountInfo) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *FiatWithdrawRequestAccountInfo) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetAgency

`func (o *FiatWithdrawRequestAccountInfo) GetAgency() string`

GetAgency returns the Agency field if non-nil, zero value otherwise.

### GetAgencyOk

`func (o *FiatWithdrawRequestAccountInfo) GetAgencyOk() (*string, bool)`

GetAgencyOk returns a tuple with the Agency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgency

`func (o *FiatWithdrawRequestAccountInfo) SetAgency(v string)`

SetAgency sets Agency field to given value.

### HasAgency

`func (o *FiatWithdrawRequestAccountInfo) HasAgency() bool`

HasAgency returns a boolean if a field has been set.

### GetBankCodeForPix

`func (o *FiatWithdrawRequestAccountInfo) GetBankCodeForPix() string`

GetBankCodeForPix returns the BankCodeForPix field if non-nil, zero value otherwise.

### GetBankCodeForPixOk

`func (o *FiatWithdrawRequestAccountInfo) GetBankCodeForPixOk() (*string, bool)`

GetBankCodeForPixOk returns a tuple with the BankCodeForPix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankCodeForPix

`func (o *FiatWithdrawRequestAccountInfo) SetBankCodeForPix(v string)`

SetBankCodeForPix sets BankCodeForPix field to given value.

### HasBankCodeForPix

`func (o *FiatWithdrawRequestAccountInfo) HasBankCodeForPix() bool`

HasBankCodeForPix returns a boolean if a field has been set.

### GetAccountType

`func (o *FiatWithdrawRequestAccountInfo) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *FiatWithdrawRequestAccountInfo) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *FiatWithdrawRequestAccountInfo) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *FiatWithdrawRequestAccountInfo) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.


[[Back to README]](../README.md)


