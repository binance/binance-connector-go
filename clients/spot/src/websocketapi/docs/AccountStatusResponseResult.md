# AccountStatusResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MakerCommission** | Pointer to **int64** |  | [optional] 
**TakerCommission** | Pointer to **int64** |  | [optional] 
**BuyerCommission** | Pointer to **int64** |  | [optional] 
**SellerCommission** | Pointer to **int64** |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**CommissionRates** | Pointer to [**AccountStatusResponseResultCommissionRates**](AccountStatusResponseResultCommissionRates.md) |  | [optional] 
**Brokered** | Pointer to **bool** |  | [optional] 
**RequireSelfTradePrevention** | Pointer to **bool** |  | [optional] 
**PreventSor** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Balances** | Pointer to [**[]AccountStatusResponseResultBalancesInner**](AccountStatusResponseResultBalancesInner.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 

## Methods

### NewAccountStatusResponseResult

`func NewAccountStatusResponseResult() *AccountStatusResponseResult`

NewAccountStatusResponseResult instantiates a new AccountStatusResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusResponseResultWithDefaults

`func NewAccountStatusResponseResultWithDefaults() *AccountStatusResponseResult`

NewAccountStatusResponseResultWithDefaults instantiates a new AccountStatusResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakerCommission

`func (o *AccountStatusResponseResult) GetMakerCommission() int64`

GetMakerCommission returns the MakerCommission field if non-nil, zero value otherwise.

### GetMakerCommissionOk

`func (o *AccountStatusResponseResult) GetMakerCommissionOk() (*int64, bool)`

GetMakerCommissionOk returns a tuple with the MakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommission

`func (o *AccountStatusResponseResult) SetMakerCommission(v int64)`

SetMakerCommission sets MakerCommission field to given value.

### HasMakerCommission

`func (o *AccountStatusResponseResult) HasMakerCommission() bool`

HasMakerCommission returns a boolean if a field has been set.

### GetTakerCommission

`func (o *AccountStatusResponseResult) GetTakerCommission() int64`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *AccountStatusResponseResult) GetTakerCommissionOk() (*int64, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *AccountStatusResponseResult) SetTakerCommission(v int64)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *AccountStatusResponseResult) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.

### GetBuyerCommission

`func (o *AccountStatusResponseResult) GetBuyerCommission() int64`

GetBuyerCommission returns the BuyerCommission field if non-nil, zero value otherwise.

### GetBuyerCommissionOk

`func (o *AccountStatusResponseResult) GetBuyerCommissionOk() (*int64, bool)`

GetBuyerCommissionOk returns a tuple with the BuyerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCommission

`func (o *AccountStatusResponseResult) SetBuyerCommission(v int64)`

SetBuyerCommission sets BuyerCommission field to given value.

### HasBuyerCommission

`func (o *AccountStatusResponseResult) HasBuyerCommission() bool`

HasBuyerCommission returns a boolean if a field has been set.

### GetSellerCommission

`func (o *AccountStatusResponseResult) GetSellerCommission() int64`

GetSellerCommission returns the SellerCommission field if non-nil, zero value otherwise.

### GetSellerCommissionOk

`func (o *AccountStatusResponseResult) GetSellerCommissionOk() (*int64, bool)`

GetSellerCommissionOk returns a tuple with the SellerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerCommission

`func (o *AccountStatusResponseResult) SetSellerCommission(v int64)`

SetSellerCommission sets SellerCommission field to given value.

### HasSellerCommission

`func (o *AccountStatusResponseResult) HasSellerCommission() bool`

HasSellerCommission returns a boolean if a field has been set.

### GetCanTrade

`func (o *AccountStatusResponseResult) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *AccountStatusResponseResult) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *AccountStatusResponseResult) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *AccountStatusResponseResult) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *AccountStatusResponseResult) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *AccountStatusResponseResult) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *AccountStatusResponseResult) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *AccountStatusResponseResult) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetCanDeposit

`func (o *AccountStatusResponseResult) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *AccountStatusResponseResult) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *AccountStatusResponseResult) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *AccountStatusResponseResult) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetCommissionRates

`func (o *AccountStatusResponseResult) GetCommissionRates() AccountStatusResponseResultCommissionRates`

GetCommissionRates returns the CommissionRates field if non-nil, zero value otherwise.

### GetCommissionRatesOk

`func (o *AccountStatusResponseResult) GetCommissionRatesOk() (*AccountStatusResponseResultCommissionRates, bool)`

GetCommissionRatesOk returns a tuple with the CommissionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRates

`func (o *AccountStatusResponseResult) SetCommissionRates(v AccountStatusResponseResultCommissionRates)`

SetCommissionRates sets CommissionRates field to given value.

### HasCommissionRates

`func (o *AccountStatusResponseResult) HasCommissionRates() bool`

HasCommissionRates returns a boolean if a field has been set.

### GetBrokered

`func (o *AccountStatusResponseResult) GetBrokered() bool`

GetBrokered returns the Brokered field if non-nil, zero value otherwise.

### GetBrokeredOk

`func (o *AccountStatusResponseResult) GetBrokeredOk() (*bool, bool)`

GetBrokeredOk returns a tuple with the Brokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokered

`func (o *AccountStatusResponseResult) SetBrokered(v bool)`

SetBrokered sets Brokered field to given value.

### HasBrokered

`func (o *AccountStatusResponseResult) HasBrokered() bool`

HasBrokered returns a boolean if a field has been set.

### GetRequireSelfTradePrevention

`func (o *AccountStatusResponseResult) GetRequireSelfTradePrevention() bool`

GetRequireSelfTradePrevention returns the RequireSelfTradePrevention field if non-nil, zero value otherwise.

### GetRequireSelfTradePreventionOk

`func (o *AccountStatusResponseResult) GetRequireSelfTradePreventionOk() (*bool, bool)`

GetRequireSelfTradePreventionOk returns a tuple with the RequireSelfTradePrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSelfTradePrevention

`func (o *AccountStatusResponseResult) SetRequireSelfTradePrevention(v bool)`

SetRequireSelfTradePrevention sets RequireSelfTradePrevention field to given value.

### HasRequireSelfTradePrevention

`func (o *AccountStatusResponseResult) HasRequireSelfTradePrevention() bool`

HasRequireSelfTradePrevention returns a boolean if a field has been set.

### GetPreventSor

`func (o *AccountStatusResponseResult) GetPreventSor() bool`

GetPreventSor returns the PreventSor field if non-nil, zero value otherwise.

### GetPreventSorOk

`func (o *AccountStatusResponseResult) GetPreventSorOk() (*bool, bool)`

GetPreventSorOk returns a tuple with the PreventSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventSor

`func (o *AccountStatusResponseResult) SetPreventSor(v bool)`

SetPreventSor sets PreventSor field to given value.

### HasPreventSor

`func (o *AccountStatusResponseResult) HasPreventSor() bool`

HasPreventSor returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AccountStatusResponseResult) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AccountStatusResponseResult) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AccountStatusResponseResult) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AccountStatusResponseResult) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountStatusResponseResult) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountStatusResponseResult) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountStatusResponseResult) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountStatusResponseResult) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *AccountStatusResponseResult) GetBalances() []AccountStatusResponseResultBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountStatusResponseResult) GetBalancesOk() (*[]AccountStatusResponseResultBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountStatusResponseResult) SetBalances(v []AccountStatusResponseResultBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *AccountStatusResponseResult) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetPermissions

`func (o *AccountStatusResponseResult) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AccountStatusResponseResult) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AccountStatusResponseResult) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AccountStatusResponseResult) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUid

`func (o *AccountStatusResponseResult) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *AccountStatusResponseResult) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *AccountStatusResponseResult) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *AccountStatusResponseResult) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to README]](../README.md)


