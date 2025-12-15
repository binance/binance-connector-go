# GetAccountResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MakerCommission** | Pointer to **int64** |  | [optional] 
**TakerCommission** | Pointer to **int64** |  | [optional] 
**BuyerCommission** | Pointer to **int64** |  | [optional] 
**SellerCommission** | Pointer to **int64** |  | [optional] 
**CommissionRates** | Pointer to [**GetAccountResponseCommissionRates**](GetAccountResponseCommissionRates.md) |  | [optional] 
**CanTrade** | Pointer to **bool** |  | [optional] 
**CanWithdraw** | Pointer to **bool** |  | [optional] 
**CanDeposit** | Pointer to **bool** |  | [optional] 
**Brokered** | Pointer to **bool** |  | [optional] 
**RequireSelfTradePrevention** | Pointer to **bool** |  | [optional] 
**PreventSor** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Balances** | Pointer to [**[]GetAccountResponseBalancesInner**](GetAccountResponseBalancesInner.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Uid** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAccountResponse

`func NewGetAccountResponse() *GetAccountResponse`

NewGetAccountResponse instantiates a new GetAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountResponseWithDefaults

`func NewGetAccountResponseWithDefaults() *GetAccountResponse`

NewGetAccountResponseWithDefaults instantiates a new GetAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMakerCommission

`func (o *GetAccountResponse) GetMakerCommission() int64`

GetMakerCommission returns the MakerCommission field if non-nil, zero value otherwise.

### GetMakerCommissionOk

`func (o *GetAccountResponse) GetMakerCommissionOk() (*int64, bool)`

GetMakerCommissionOk returns a tuple with the MakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerCommission

`func (o *GetAccountResponse) SetMakerCommission(v int64)`

SetMakerCommission sets MakerCommission field to given value.

### HasMakerCommission

`func (o *GetAccountResponse) HasMakerCommission() bool`

HasMakerCommission returns a boolean if a field has been set.

### GetTakerCommission

`func (o *GetAccountResponse) GetTakerCommission() int64`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *GetAccountResponse) GetTakerCommissionOk() (*int64, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *GetAccountResponse) SetTakerCommission(v int64)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *GetAccountResponse) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.

### GetBuyerCommission

`func (o *GetAccountResponse) GetBuyerCommission() int64`

GetBuyerCommission returns the BuyerCommission field if non-nil, zero value otherwise.

### GetBuyerCommissionOk

`func (o *GetAccountResponse) GetBuyerCommissionOk() (*int64, bool)`

GetBuyerCommissionOk returns a tuple with the BuyerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCommission

`func (o *GetAccountResponse) SetBuyerCommission(v int64)`

SetBuyerCommission sets BuyerCommission field to given value.

### HasBuyerCommission

`func (o *GetAccountResponse) HasBuyerCommission() bool`

HasBuyerCommission returns a boolean if a field has been set.

### GetSellerCommission

`func (o *GetAccountResponse) GetSellerCommission() int64`

GetSellerCommission returns the SellerCommission field if non-nil, zero value otherwise.

### GetSellerCommissionOk

`func (o *GetAccountResponse) GetSellerCommissionOk() (*int64, bool)`

GetSellerCommissionOk returns a tuple with the SellerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerCommission

`func (o *GetAccountResponse) SetSellerCommission(v int64)`

SetSellerCommission sets SellerCommission field to given value.

### HasSellerCommission

`func (o *GetAccountResponse) HasSellerCommission() bool`

HasSellerCommission returns a boolean if a field has been set.

### GetCommissionRates

`func (o *GetAccountResponse) GetCommissionRates() GetAccountResponseCommissionRates`

GetCommissionRates returns the CommissionRates field if non-nil, zero value otherwise.

### GetCommissionRatesOk

`func (o *GetAccountResponse) GetCommissionRatesOk() (*GetAccountResponseCommissionRates, bool)`

GetCommissionRatesOk returns a tuple with the CommissionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionRates

`func (o *GetAccountResponse) SetCommissionRates(v GetAccountResponseCommissionRates)`

SetCommissionRates sets CommissionRates field to given value.

### HasCommissionRates

`func (o *GetAccountResponse) HasCommissionRates() bool`

HasCommissionRates returns a boolean if a field has been set.

### GetCanTrade

`func (o *GetAccountResponse) GetCanTrade() bool`

GetCanTrade returns the CanTrade field if non-nil, zero value otherwise.

### GetCanTradeOk

`func (o *GetAccountResponse) GetCanTradeOk() (*bool, bool)`

GetCanTradeOk returns a tuple with the CanTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTrade

`func (o *GetAccountResponse) SetCanTrade(v bool)`

SetCanTrade sets CanTrade field to given value.

### HasCanTrade

`func (o *GetAccountResponse) HasCanTrade() bool`

HasCanTrade returns a boolean if a field has been set.

### GetCanWithdraw

`func (o *GetAccountResponse) GetCanWithdraw() bool`

GetCanWithdraw returns the CanWithdraw field if non-nil, zero value otherwise.

### GetCanWithdrawOk

`func (o *GetAccountResponse) GetCanWithdrawOk() (*bool, bool)`

GetCanWithdrawOk returns a tuple with the CanWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanWithdraw

`func (o *GetAccountResponse) SetCanWithdraw(v bool)`

SetCanWithdraw sets CanWithdraw field to given value.

### HasCanWithdraw

`func (o *GetAccountResponse) HasCanWithdraw() bool`

HasCanWithdraw returns a boolean if a field has been set.

### GetCanDeposit

`func (o *GetAccountResponse) GetCanDeposit() bool`

GetCanDeposit returns the CanDeposit field if non-nil, zero value otherwise.

### GetCanDepositOk

`func (o *GetAccountResponse) GetCanDepositOk() (*bool, bool)`

GetCanDepositOk returns a tuple with the CanDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDeposit

`func (o *GetAccountResponse) SetCanDeposit(v bool)`

SetCanDeposit sets CanDeposit field to given value.

### HasCanDeposit

`func (o *GetAccountResponse) HasCanDeposit() bool`

HasCanDeposit returns a boolean if a field has been set.

### GetBrokered

`func (o *GetAccountResponse) GetBrokered() bool`

GetBrokered returns the Brokered field if non-nil, zero value otherwise.

### GetBrokeredOk

`func (o *GetAccountResponse) GetBrokeredOk() (*bool, bool)`

GetBrokeredOk returns a tuple with the Brokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokered

`func (o *GetAccountResponse) SetBrokered(v bool)`

SetBrokered sets Brokered field to given value.

### HasBrokered

`func (o *GetAccountResponse) HasBrokered() bool`

HasBrokered returns a boolean if a field has been set.

### GetRequireSelfTradePrevention

`func (o *GetAccountResponse) GetRequireSelfTradePrevention() bool`

GetRequireSelfTradePrevention returns the RequireSelfTradePrevention field if non-nil, zero value otherwise.

### GetRequireSelfTradePreventionOk

`func (o *GetAccountResponse) GetRequireSelfTradePreventionOk() (*bool, bool)`

GetRequireSelfTradePreventionOk returns a tuple with the RequireSelfTradePrevention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSelfTradePrevention

`func (o *GetAccountResponse) SetRequireSelfTradePrevention(v bool)`

SetRequireSelfTradePrevention sets RequireSelfTradePrevention field to given value.

### HasRequireSelfTradePrevention

`func (o *GetAccountResponse) HasRequireSelfTradePrevention() bool`

HasRequireSelfTradePrevention returns a boolean if a field has been set.

### GetPreventSor

`func (o *GetAccountResponse) GetPreventSor() bool`

GetPreventSor returns the PreventSor field if non-nil, zero value otherwise.

### GetPreventSorOk

`func (o *GetAccountResponse) GetPreventSorOk() (*bool, bool)`

GetPreventSorOk returns a tuple with the PreventSor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventSor

`func (o *GetAccountResponse) SetPreventSor(v bool)`

SetPreventSor sets PreventSor field to given value.

### HasPreventSor

`func (o *GetAccountResponse) HasPreventSor() bool`

HasPreventSor returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetAccountResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetAccountResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetAccountResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetAccountResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAccountType

`func (o *GetAccountResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetAccountResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetAccountResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetAccountResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetBalances

`func (o *GetAccountResponse) GetBalances() []GetAccountResponseBalancesInner`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *GetAccountResponse) GetBalancesOk() (*[]GetAccountResponseBalancesInner, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *GetAccountResponse) SetBalances(v []GetAccountResponseBalancesInner)`

SetBalances sets Balances field to given value.

### HasBalances

`func (o *GetAccountResponse) HasBalances() bool`

HasBalances returns a boolean if a field has been set.

### GetPermissions

`func (o *GetAccountResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetAccountResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetAccountResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetAccountResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetUid

`func (o *GetAccountResponse) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetAccountResponse) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetAccountResponse) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetAccountResponse) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to README]](../README.md)


