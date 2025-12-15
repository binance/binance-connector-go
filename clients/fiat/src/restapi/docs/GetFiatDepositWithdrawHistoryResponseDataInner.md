# GetFiatDepositWithdrawHistoryResponseDataInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderNo** | Pointer to **string** |  | [optional] 
**FiatCurrency** | Pointer to **string** |  | [optional] 
**IndicatedAmount** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**TotalFee** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFiatDepositWithdrawHistoryResponseDataInner

`func NewGetFiatDepositWithdrawHistoryResponseDataInner() *GetFiatDepositWithdrawHistoryResponseDataInner`

NewGetFiatDepositWithdrawHistoryResponseDataInner instantiates a new GetFiatDepositWithdrawHistoryResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositWithdrawHistoryResponseDataInnerWithDefaults

`func NewGetFiatDepositWithdrawHistoryResponseDataInnerWithDefaults() *GetFiatDepositWithdrawHistoryResponseDataInner`

NewGetFiatDepositWithdrawHistoryResponseDataInnerWithDefaults instantiates a new GetFiatDepositWithdrawHistoryResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNo

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.

### HasOrderNo

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasOrderNo() bool`

HasOrderNo returns a boolean if a field has been set.

### GetFiatCurrency

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.

### HasFiatCurrency

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasFiatCurrency() bool`

HasFiatCurrency returns a boolean if a field has been set.

### GetIndicatedAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetIndicatedAmount() string`

GetIndicatedAmount returns the IndicatedAmount field if non-nil, zero value otherwise.

### GetIndicatedAmountOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetIndicatedAmountOk() (*string, bool)`

GetIndicatedAmountOk returns a tuple with the IndicatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicatedAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetIndicatedAmount(v string)`

SetIndicatedAmount sets IndicatedAmount field to given value.

### HasIndicatedAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasIndicatedAmount() bool`

HasIndicatedAmount returns a boolean if a field has been set.

### GetAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotalFee

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetTotalFee() string`

GetTotalFee returns the TotalFee field if non-nil, zero value otherwise.

### GetTotalFeeOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetTotalFeeOk() (*string, bool)`

GetTotalFeeOk returns a tuple with the TotalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFee

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetTotalFee(v string)`

SetTotalFee sets TotalFee field to given value.

### HasTotalFee

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasTotalFee() bool`

HasTotalFee returns a boolean if a field has been set.

### GetMethod

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetStatus

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetFiatDepositWithdrawHistoryResponseDataInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


