# VipLoanFixedRateBorrowResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BorrowCoin** | Pointer to **string** | Echo of input parameter | [optional] 
**BorrowAmount** | Pointer to **string** | Actual total borrow amount (aggregated when multiple supplyRequest) | [optional] 
**ActualReceivedAmount** | Pointer to **string** | Actual received amount | [optional] 
**CollateralCoin** | Pointer to **string** | Echo of input parameter, comma-separated | [optional] 
**CollateralAccountId** | Pointer to **string** | Echo of input parameter, comma-separated | [optional] 
**BorrowInterestRate** | Pointer to **string** | Actual borrow interest rate (weighted average when multiple) | [optional] 
**Duration** | Pointer to **string** | &#x60;{loanTerm}Days&#x60;, e.g. \&quot;30Days\&quot; | [optional] 
**AutoRepay** | Pointer to **bool** | Echo of input parameter | [optional] 
**OrderId** | Pointer to **int64** | Order ID | [optional] 
**Status** | Pointer to **string** | &#x60;Succeeds&#x60; / &#x60;Failed&#x60; / &#x60;Processing&#x60; | [optional] 

## Methods

### NewVipLoanFixedRateBorrowResponse

`func NewVipLoanFixedRateBorrowResponse() *VipLoanFixedRateBorrowResponse`

NewVipLoanFixedRateBorrowResponse instantiates a new VipLoanFixedRateBorrowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVipLoanFixedRateBorrowResponseWithDefaults

`func NewVipLoanFixedRateBorrowResponseWithDefaults() *VipLoanFixedRateBorrowResponse`

NewVipLoanFixedRateBorrowResponseWithDefaults instantiates a new VipLoanFixedRateBorrowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBorrowCoin

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowCoin() string`

GetBorrowCoin returns the BorrowCoin field if non-nil, zero value otherwise.

### GetBorrowCoinOk

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowCoinOk() (*string, bool)`

GetBorrowCoinOk returns a tuple with the BorrowCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowCoin

`func (o *VipLoanFixedRateBorrowResponse) SetBorrowCoin(v string)`

SetBorrowCoin sets BorrowCoin field to given value.

### HasBorrowCoin

`func (o *VipLoanFixedRateBorrowResponse) HasBorrowCoin() bool`

HasBorrowCoin returns a boolean if a field has been set.

### GetBorrowAmount

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowAmount() string`

GetBorrowAmount returns the BorrowAmount field if non-nil, zero value otherwise.

### GetBorrowAmountOk

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowAmountOk() (*string, bool)`

GetBorrowAmountOk returns a tuple with the BorrowAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowAmount

`func (o *VipLoanFixedRateBorrowResponse) SetBorrowAmount(v string)`

SetBorrowAmount sets BorrowAmount field to given value.

### HasBorrowAmount

`func (o *VipLoanFixedRateBorrowResponse) HasBorrowAmount() bool`

HasBorrowAmount returns a boolean if a field has been set.

### GetActualReceivedAmount

`func (o *VipLoanFixedRateBorrowResponse) GetActualReceivedAmount() string`

GetActualReceivedAmount returns the ActualReceivedAmount field if non-nil, zero value otherwise.

### GetActualReceivedAmountOk

`func (o *VipLoanFixedRateBorrowResponse) GetActualReceivedAmountOk() (*string, bool)`

GetActualReceivedAmountOk returns a tuple with the ActualReceivedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualReceivedAmount

`func (o *VipLoanFixedRateBorrowResponse) SetActualReceivedAmount(v string)`

SetActualReceivedAmount sets ActualReceivedAmount field to given value.

### HasActualReceivedAmount

`func (o *VipLoanFixedRateBorrowResponse) HasActualReceivedAmount() bool`

HasActualReceivedAmount returns a boolean if a field has been set.

### GetCollateralCoin

`func (o *VipLoanFixedRateBorrowResponse) GetCollateralCoin() string`

GetCollateralCoin returns the CollateralCoin field if non-nil, zero value otherwise.

### GetCollateralCoinOk

`func (o *VipLoanFixedRateBorrowResponse) GetCollateralCoinOk() (*string, bool)`

GetCollateralCoinOk returns a tuple with the CollateralCoin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralCoin

`func (o *VipLoanFixedRateBorrowResponse) SetCollateralCoin(v string)`

SetCollateralCoin sets CollateralCoin field to given value.

### HasCollateralCoin

`func (o *VipLoanFixedRateBorrowResponse) HasCollateralCoin() bool`

HasCollateralCoin returns a boolean if a field has been set.

### GetCollateralAccountId

`func (o *VipLoanFixedRateBorrowResponse) GetCollateralAccountId() string`

GetCollateralAccountId returns the CollateralAccountId field if non-nil, zero value otherwise.

### GetCollateralAccountIdOk

`func (o *VipLoanFixedRateBorrowResponse) GetCollateralAccountIdOk() (*string, bool)`

GetCollateralAccountIdOk returns a tuple with the CollateralAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollateralAccountId

`func (o *VipLoanFixedRateBorrowResponse) SetCollateralAccountId(v string)`

SetCollateralAccountId sets CollateralAccountId field to given value.

### HasCollateralAccountId

`func (o *VipLoanFixedRateBorrowResponse) HasCollateralAccountId() bool`

HasCollateralAccountId returns a boolean if a field has been set.

### GetBorrowInterestRate

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowInterestRate() string`

GetBorrowInterestRate returns the BorrowInterestRate field if non-nil, zero value otherwise.

### GetBorrowInterestRateOk

`func (o *VipLoanFixedRateBorrowResponse) GetBorrowInterestRateOk() (*string, bool)`

GetBorrowInterestRateOk returns a tuple with the BorrowInterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowInterestRate

`func (o *VipLoanFixedRateBorrowResponse) SetBorrowInterestRate(v string)`

SetBorrowInterestRate sets BorrowInterestRate field to given value.

### HasBorrowInterestRate

`func (o *VipLoanFixedRateBorrowResponse) HasBorrowInterestRate() bool`

HasBorrowInterestRate returns a boolean if a field has been set.

### GetDuration

`func (o *VipLoanFixedRateBorrowResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VipLoanFixedRateBorrowResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VipLoanFixedRateBorrowResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VipLoanFixedRateBorrowResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAutoRepay

`func (o *VipLoanFixedRateBorrowResponse) GetAutoRepay() bool`

GetAutoRepay returns the AutoRepay field if non-nil, zero value otherwise.

### GetAutoRepayOk

`func (o *VipLoanFixedRateBorrowResponse) GetAutoRepayOk() (*bool, bool)`

GetAutoRepayOk returns a tuple with the AutoRepay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRepay

`func (o *VipLoanFixedRateBorrowResponse) SetAutoRepay(v bool)`

SetAutoRepay sets AutoRepay field to given value.

### HasAutoRepay

`func (o *VipLoanFixedRateBorrowResponse) HasAutoRepay() bool`

HasAutoRepay returns a boolean if a field has been set.

### GetOrderId

`func (o *VipLoanFixedRateBorrowResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *VipLoanFixedRateBorrowResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *VipLoanFixedRateBorrowResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *VipLoanFixedRateBorrowResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *VipLoanFixedRateBorrowResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VipLoanFixedRateBorrowResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VipLoanFixedRateBorrowResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VipLoanFixedRateBorrowResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


