# QueryCrossMarginFeeDataResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**VipLevel** | Pointer to **int64** |  | [optional] 
**Coin** | Pointer to **string** |  | [optional] 
**TransferIn** | Pointer to **bool** |  | [optional] 
**Borrowable** | Pointer to **bool** |  | [optional] 
**DailyInterest** | Pointer to **string** |  | [optional] 
**YearlyInterest** | Pointer to **string** |  | [optional] 
**BorrowLimit** | Pointer to **string** |  | [optional] 
**MarginablePairs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewQueryCrossMarginFeeDataResponseInner

`func NewQueryCrossMarginFeeDataResponseInner() *QueryCrossMarginFeeDataResponseInner`

NewQueryCrossMarginFeeDataResponseInner instantiates a new QueryCrossMarginFeeDataResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCrossMarginFeeDataResponseInnerWithDefaults

`func NewQueryCrossMarginFeeDataResponseInnerWithDefaults() *QueryCrossMarginFeeDataResponseInner`

NewQueryCrossMarginFeeDataResponseInnerWithDefaults instantiates a new QueryCrossMarginFeeDataResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVipLevel

`func (o *QueryCrossMarginFeeDataResponseInner) GetVipLevel() int64`

GetVipLevel returns the VipLevel field if non-nil, zero value otherwise.

### GetVipLevelOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetVipLevelOk() (*int64, bool)`

GetVipLevelOk returns a tuple with the VipLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipLevel

`func (o *QueryCrossMarginFeeDataResponseInner) SetVipLevel(v int64)`

SetVipLevel sets VipLevel field to given value.

### HasVipLevel

`func (o *QueryCrossMarginFeeDataResponseInner) HasVipLevel() bool`

HasVipLevel returns a boolean if a field has been set.

### GetCoin

`func (o *QueryCrossMarginFeeDataResponseInner) GetCoin() string`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetCoinOk() (*string, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *QueryCrossMarginFeeDataResponseInner) SetCoin(v string)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *QueryCrossMarginFeeDataResponseInner) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetTransferIn

`func (o *QueryCrossMarginFeeDataResponseInner) GetTransferIn() bool`

GetTransferIn returns the TransferIn field if non-nil, zero value otherwise.

### GetTransferInOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetTransferInOk() (*bool, bool)`

GetTransferInOk returns a tuple with the TransferIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferIn

`func (o *QueryCrossMarginFeeDataResponseInner) SetTransferIn(v bool)`

SetTransferIn sets TransferIn field to given value.

### HasTransferIn

`func (o *QueryCrossMarginFeeDataResponseInner) HasTransferIn() bool`

HasTransferIn returns a boolean if a field has been set.

### GetBorrowable

`func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowable() bool`

GetBorrowable returns the Borrowable field if non-nil, zero value otherwise.

### GetBorrowableOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowableOk() (*bool, bool)`

GetBorrowableOk returns a tuple with the Borrowable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowable

`func (o *QueryCrossMarginFeeDataResponseInner) SetBorrowable(v bool)`

SetBorrowable sets Borrowable field to given value.

### HasBorrowable

`func (o *QueryCrossMarginFeeDataResponseInner) HasBorrowable() bool`

HasBorrowable returns a boolean if a field has been set.

### GetDailyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) GetDailyInterest() string`

GetDailyInterest returns the DailyInterest field if non-nil, zero value otherwise.

### GetDailyInterestOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetDailyInterestOk() (*string, bool)`

GetDailyInterestOk returns a tuple with the DailyInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) SetDailyInterest(v string)`

SetDailyInterest sets DailyInterest field to given value.

### HasDailyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) HasDailyInterest() bool`

HasDailyInterest returns a boolean if a field has been set.

### GetYearlyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) GetYearlyInterest() string`

GetYearlyInterest returns the YearlyInterest field if non-nil, zero value otherwise.

### GetYearlyInterestOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetYearlyInterestOk() (*string, bool)`

GetYearlyInterestOk returns a tuple with the YearlyInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYearlyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) SetYearlyInterest(v string)`

SetYearlyInterest sets YearlyInterest field to given value.

### HasYearlyInterest

`func (o *QueryCrossMarginFeeDataResponseInner) HasYearlyInterest() bool`

HasYearlyInterest returns a boolean if a field has been set.

### GetBorrowLimit

`func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowLimit() string`

GetBorrowLimit returns the BorrowLimit field if non-nil, zero value otherwise.

### GetBorrowLimitOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetBorrowLimitOk() (*string, bool)`

GetBorrowLimitOk returns a tuple with the BorrowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorrowLimit

`func (o *QueryCrossMarginFeeDataResponseInner) SetBorrowLimit(v string)`

SetBorrowLimit sets BorrowLimit field to given value.

### HasBorrowLimit

`func (o *QueryCrossMarginFeeDataResponseInner) HasBorrowLimit() bool`

HasBorrowLimit returns a boolean if a field has been set.

### GetMarginablePairs

`func (o *QueryCrossMarginFeeDataResponseInner) GetMarginablePairs() []string`

GetMarginablePairs returns the MarginablePairs field if non-nil, zero value otherwise.

### GetMarginablePairsOk

`func (o *QueryCrossMarginFeeDataResponseInner) GetMarginablePairsOk() (*[]string, bool)`

GetMarginablePairsOk returns a tuple with the MarginablePairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginablePairs

`func (o *QueryCrossMarginFeeDataResponseInner) SetMarginablePairs(v []string)`

SetMarginablePairs sets MarginablePairs field to given value.

### HasMarginablePairs

`func (o *QueryCrossMarginFeeDataResponseInner) HasMarginablePairs() bool`

HasMarginablePairs returns a boolean if a field has been set.


[[Back to README]](../README.md)


