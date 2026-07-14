# GetVIPLoanRepaymentHistoryResponseRowsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LoanCoin** | Pointer to **string** |  | [optional] 
**RepayAmount** | Pointer to **string** |  | [optional] 
**CollateralCoin** | Pointer to **string** |  | [optional] 
**RepayStatus** | Pointer to **string** | Repayment status (&#x60;Repaid&#x60;, &#x60;Repaying&#x60;, &#x60;Failed&#x60;). | [optional] 
**LoanDate** | Pointer to **string** |  | [optional] 
**RepayTime** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetVIPLoanRepaymentHistoryResponseRowsInner

`func NewGetVIPLoanRepaymentHistoryResponseRowsInner() *GetVIPLoanRepaymentHistoryResponseRowsInner`

NewGetVIPLoanRepaymentHistoryResponseRowsInner instantiates a new GetVIPLoanRepaymentHistoryResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVIPLoanRepaymentHistoryResponseRowsInnerWithDefaults

`func NewGetVIPLoanRepaymentHistoryResponseRowsInnerWithDefaults() *GetVIPLoanRepaymentHistoryResponseRowsInner`

NewGetVIPLoanRepaymentHistoryResponseRowsInnerWithDefaults instantiates a new GetVIPLoanRepaymentHistoryResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoanCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanCoin() string`

GetLoanCoin returns the LoanCoin field if non-nil, zero value otherwise.

### GetLoanCoinOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanCoinOk() (*string, bool)`

GetLoanCoinOk returns a tuple with the LoanCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetLoanCoin(v string)`

SetLoanCoin sets LoanCoin field to given value.

### HasLoanCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasLoanCoin() bool`

HasLoanCoin returns a boolean if a field has been set.

### GetRepayAmount

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayAmount() string`

GetRepayAmount returns the RepayAmount field if non-nil, zero value otherwise.

### GetRepayAmountOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayAmountOk() (*string, bool)`

GetRepayAmountOk returns a tuple with the RepayAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayAmount

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayAmount(v string)`

SetRepayAmount sets RepayAmount field to given value.

### HasRepayAmount

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayAmount() bool`

HasRepayAmount returns a boolean if a field has been set.

### GetCollateralCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetCollateralCoin() string`

GetCollateralCoin returns the CollateralCoin field if non-nil, zero value otherwise.

### GetCollateralCoinOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetCollateralCoinOk() (*string, bool)`

GetCollateralCoinOk returns a tuple with the CollateralCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetCollateralCoin(v string)`

SetCollateralCoin sets CollateralCoin field to given value.

### HasCollateralCoin

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasCollateralCoin() bool`

HasCollateralCoin returns a boolean if a field has been set.

### GetRepayStatus

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayStatus() string`

GetRepayStatus returns the RepayStatus field if non-nil, zero value otherwise.

### GetRepayStatusOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayStatusOk() (*string, bool)`

GetRepayStatusOk returns a tuple with the RepayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayStatus

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayStatus(v string)`

SetRepayStatus sets RepayStatus field to given value.

### HasRepayStatus

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayStatus() bool`

HasRepayStatus returns a boolean if a field has been set.

### GetLoanDate

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanDate() string`

GetLoanDate returns the LoanDate field if non-nil, zero value otherwise.

### GetLoanDateOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetLoanDateOk() (*string, bool)`

GetLoanDateOk returns a tuple with the LoanDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanDate

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetLoanDate(v string)`

SetLoanDate sets LoanDate field to given value.

### HasLoanDate

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasLoanDate() bool`

HasLoanDate returns a boolean if a field has been set.

### GetRepayTime

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayTime() string`

GetRepayTime returns the RepayTime field if non-nil, zero value otherwise.

### GetRepayTimeOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetRepayTimeOk() (*string, bool)`

GetRepayTimeOk returns a tuple with the RepayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayTime

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetRepayTime(v string)`

SetRepayTime sets RepayTime field to given value.

### HasRepayTime

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasRepayTime() bool`

HasRepayTime returns a boolean if a field has been set.

### GetOrderId

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetVIPLoanRepaymentHistoryResponseRowsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to README]](../README.md)


