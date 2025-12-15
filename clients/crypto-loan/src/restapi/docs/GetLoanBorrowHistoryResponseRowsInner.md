# GetLoanBorrowHistoryResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**LoanCoin** | Pointer to **string** |  | [optional] 
**InitialLoanAmount** | Pointer to **string** |  | [optional] 
**HourlyInterestRate** | Pointer to **string** |  | [optional] 
**LoanTerm** | Pointer to **string** |  | [optional] 
**CollateralCoin** | Pointer to **string** |  | [optional] 
**InitialCollateralAmount** | Pointer to **string** |  | [optional] 
**BorrowTime** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewGetLoanBorrowHistoryResponseRowsInner

`func NewGetLoanBorrowHistoryResponseRowsInner() *GetLoanBorrowHistoryResponseRowsInner`

NewGetLoanBorrowHistoryResponseRowsInner instantiates a new GetLoanBorrowHistoryResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanBorrowHistoryResponseRowsInnerWithDefaults

`func NewGetLoanBorrowHistoryResponseRowsInnerWithDefaults() *GetLoanBorrowHistoryResponseRowsInner`

NewGetLoanBorrowHistoryResponseRowsInnerWithDefaults instantiates a new GetLoanBorrowHistoryResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetLoanCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanCoin() string`

GetLoanCoin returns the LoanCoin field if non-nil, zero value otherwise.

### GetLoanCoinOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool)`

GetLoanCoinOk returns a tuple with the LoanCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetLoanCoin(v string)`

SetLoanCoin sets LoanCoin field to given value.

### HasLoanCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasLoanCoin() bool`

HasLoanCoin returns a boolean if a field has been set.

### GetInitialLoanAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmount() string`

GetInitialLoanAmount returns the InitialLoanAmount field if non-nil, zero value otherwise.

### GetInitialLoanAmountOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialLoanAmountOk() (*string, bool)`

GetInitialLoanAmountOk returns a tuple with the InitialLoanAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialLoanAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetInitialLoanAmount(v string)`

SetInitialLoanAmount sets InitialLoanAmount field to given value.

### HasInitialLoanAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasInitialLoanAmount() bool`

HasInitialLoanAmount returns a boolean if a field has been set.

### GetHourlyInterestRate

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetHourlyInterestRate() string`

GetHourlyInterestRate returns the HourlyInterestRate field if non-nil, zero value otherwise.

### GetHourlyInterestRateOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetHourlyInterestRateOk() (*string, bool)`

GetHourlyInterestRateOk returns a tuple with the HourlyInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyInterestRate

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetHourlyInterestRate(v string)`

SetHourlyInterestRate sets HourlyInterestRate field to given value.

### HasHourlyInterestRate

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasHourlyInterestRate() bool`

HasHourlyInterestRate returns a boolean if a field has been set.

### GetLoanTerm

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanTerm() string`

GetLoanTerm returns the LoanTerm field if non-nil, zero value otherwise.

### GetLoanTermOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetLoanTermOk() (*string, bool)`

GetLoanTermOk returns a tuple with the LoanTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanTerm

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetLoanTerm(v string)`

SetLoanTerm sets LoanTerm field to given value.

### HasLoanTerm

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasLoanTerm() bool`

HasLoanTerm returns a boolean if a field has been set.

### GetCollateralCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetCollateralCoin() string`

GetCollateralCoin returns the CollateralCoin field if non-nil, zero value otherwise.

### GetCollateralCoinOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool)`

GetCollateralCoinOk returns a tuple with the CollateralCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetCollateralCoin(v string)`

SetCollateralCoin sets CollateralCoin field to given value.

### HasCollateralCoin

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasCollateralCoin() bool`

HasCollateralCoin returns a boolean if a field has been set.

### GetInitialCollateralAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmount() string`

GetInitialCollateralAmount returns the InitialCollateralAmount field if non-nil, zero value otherwise.

### GetInitialCollateralAmountOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetInitialCollateralAmountOk() (*string, bool)`

GetInitialCollateralAmountOk returns a tuple with the InitialCollateralAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCollateralAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetInitialCollateralAmount(v string)`

SetInitialCollateralAmount sets InitialCollateralAmount field to given value.

### HasInitialCollateralAmount

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasInitialCollateralAmount() bool`

HasInitialCollateralAmount returns a boolean if a field has been set.

### GetBorrowTime

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetBorrowTime() int64`

GetBorrowTime returns the BorrowTime field if non-nil, zero value otherwise.

### GetBorrowTimeOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetBorrowTimeOk() (*int64, bool)`

GetBorrowTimeOk returns a tuple with the BorrowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowTime

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetBorrowTime(v int64)`

SetBorrowTime sets BorrowTime field to given value.

### HasBorrowTime

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasBorrowTime() bool`

HasBorrowTime returns a boolean if a field has been set.

### GetStatus

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLoanBorrowHistoryResponseRowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLoanBorrowHistoryResponseRowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLoanBorrowHistoryResponseRowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


