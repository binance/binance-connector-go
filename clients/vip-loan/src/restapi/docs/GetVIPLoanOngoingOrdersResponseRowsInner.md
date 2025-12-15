# GetVIPLoanOngoingOrdersResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**LoanCoin** | Pointer to **string** |  | [optional] 
**TotalDebt** | Pointer to **string** |  | [optional] 
**ResidualInterest** | Pointer to **string** |  | [optional] 
**CollateralAccountId** | Pointer to **string** |  | [optional] 
**CollateralCoin** | Pointer to **string** |  | [optional] 
**TotalCollateralValueAfterHaircut** | Pointer to **string** |  | [optional] 
**LockedCollateralValue** | Pointer to **string** |  | [optional] 
**CurrentLTV** | Pointer to **string** |  | [optional] 
**ExpirationTime** | Pointer to **int64** |  | [optional] 
**LoanDate** | Pointer to **string** |  | [optional] 
**LoanTerm** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVIPLoanOngoingOrdersResponseRowsInner

`func NewGetVIPLoanOngoingOrdersResponseRowsInner() *GetVIPLoanOngoingOrdersResponseRowsInner`

NewGetVIPLoanOngoingOrdersResponseRowsInner instantiates a new GetVIPLoanOngoingOrdersResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVIPLoanOngoingOrdersResponseRowsInnerWithDefaults

`func NewGetVIPLoanOngoingOrdersResponseRowsInnerWithDefaults() *GetVIPLoanOngoingOrdersResponseRowsInner`

NewGetVIPLoanOngoingOrdersResponseRowsInnerWithDefaults instantiates a new GetVIPLoanOngoingOrdersResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetLoanCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanCoin() string`

GetLoanCoin returns the LoanCoin field if non-nil, zero value otherwise.

### GetLoanCoinOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanCoinOk() (*string, bool)`

GetLoanCoinOk returns a tuple with the LoanCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanCoin(v string)`

SetLoanCoin sets LoanCoin field to given value.

### HasLoanCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanCoin() bool`

HasLoanCoin returns a boolean if a field has been set.

### GetTotalDebt

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalDebt() string`

GetTotalDebt returns the TotalDebt field if non-nil, zero value otherwise.

### GetTotalDebtOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalDebtOk() (*string, bool)`

GetTotalDebtOk returns a tuple with the TotalDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebt

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetTotalDebt(v string)`

SetTotalDebt sets TotalDebt field to given value.

### HasTotalDebt

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasTotalDebt() bool`

HasTotalDebt returns a boolean if a field has been set.

### GetResidualInterest

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetResidualInterest() string`

GetResidualInterest returns the ResidualInterest field if non-nil, zero value otherwise.

### GetResidualInterestOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetResidualInterestOk() (*string, bool)`

GetResidualInterestOk returns a tuple with the ResidualInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidualInterest

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetResidualInterest(v string)`

SetResidualInterest sets ResidualInterest field to given value.

### HasResidualInterest

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasResidualInterest() bool`

HasResidualInterest returns a boolean if a field has been set.

### GetCollateralAccountId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralAccountId() string`

GetCollateralAccountId returns the CollateralAccountId field if non-nil, zero value otherwise.

### GetCollateralAccountIdOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralAccountIdOk() (*string, bool)`

GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralAccountId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCollateralAccountId(v string)`

SetCollateralAccountId sets CollateralAccountId field to given value.

### HasCollateralAccountId

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCollateralAccountId() bool`

HasCollateralAccountId returns a boolean if a field has been set.

### GetCollateralCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralCoin() string`

GetCollateralCoin returns the CollateralCoin field if non-nil, zero value otherwise.

### GetCollateralCoinOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCollateralCoinOk() (*string, bool)`

GetCollateralCoinOk returns a tuple with the CollateralCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCollateralCoin(v string)`

SetCollateralCoin sets CollateralCoin field to given value.

### HasCollateralCoin

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCollateralCoin() bool`

HasCollateralCoin returns a boolean if a field has been set.

### GetTotalCollateralValueAfterHaircut

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalCollateralValueAfterHaircut() string`

GetTotalCollateralValueAfterHaircut returns the TotalCollateralValueAfterHaircut field if non-nil, zero value otherwise.

### GetTotalCollateralValueAfterHaircutOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetTotalCollateralValueAfterHaircutOk() (*string, bool)`

GetTotalCollateralValueAfterHaircutOk returns a tuple with the TotalCollateralValueAfterHaircut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCollateralValueAfterHaircut

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetTotalCollateralValueAfterHaircut(v string)`

SetTotalCollateralValueAfterHaircut sets TotalCollateralValueAfterHaircut field to given value.

### HasTotalCollateralValueAfterHaircut

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasTotalCollateralValueAfterHaircut() bool`

HasTotalCollateralValueAfterHaircut returns a boolean if a field has been set.

### GetLockedCollateralValue

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLockedCollateralValue() string`

GetLockedCollateralValue returns the LockedCollateralValue field if non-nil, zero value otherwise.

### GetLockedCollateralValueOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLockedCollateralValueOk() (*string, bool)`

GetLockedCollateralValueOk returns a tuple with the LockedCollateralValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedCollateralValue

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLockedCollateralValue(v string)`

SetLockedCollateralValue sets LockedCollateralValue field to given value.

### HasLockedCollateralValue

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLockedCollateralValue() bool`

HasLockedCollateralValue returns a boolean if a field has been set.

### GetCurrentLTV

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCurrentLTV() string`

GetCurrentLTV returns the CurrentLTV field if non-nil, zero value otherwise.

### GetCurrentLTVOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetCurrentLTVOk() (*string, bool)`

GetCurrentLTVOk returns a tuple with the CurrentLTV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLTV

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetCurrentLTV(v string)`

SetCurrentLTV sets CurrentLTV field to given value.

### HasCurrentLTV

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasCurrentLTV() bool`

HasCurrentLTV returns a boolean if a field has been set.

### GetExpirationTime

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLoanDate

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanDate() string`

GetLoanDate returns the LoanDate field if non-nil, zero value otherwise.

### GetLoanDateOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanDateOk() (*string, bool)`

GetLoanDateOk returns a tuple with the LoanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanDate

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanDate(v string)`

SetLoanDate sets LoanDate field to given value.

### HasLoanDate

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanDate() bool`

HasLoanDate returns a boolean if a field has been set.

### GetLoanTerm

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanTerm() string`

GetLoanTerm returns the LoanTerm field if non-nil, zero value otherwise.

### GetLoanTermOk

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) GetLoanTermOk() (*string, bool)`

GetLoanTermOk returns a tuple with the LoanTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTerm

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) SetLoanTerm(v string)`

SetLoanTerm sets LoanTerm field to given value.

### HasLoanTerm

`func (o *GetVIPLoanOngoingOrdersResponseRowsInner) HasLoanTerm() bool`

HasLoanTerm returns a boolean if a field has been set.


[[Back to README]](../README.md)


