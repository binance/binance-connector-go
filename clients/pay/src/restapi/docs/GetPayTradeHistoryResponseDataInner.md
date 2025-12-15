# GetPayTradeHistoryResponseDataInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderType** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TransactionTime** | Pointer to **int64** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**WalletType** | Pointer to **int64** |  | [optional] 
**WalletTypes** | Pointer to **[]int64** |  | [optional] 
**FundsDetail** | Pointer to [**[]GetPayTradeHistoryResponseDataInnerFundsDetailInner**](GetPayTradeHistoryResponseDataInnerFundsDetailInner.md) |  | [optional] 
**PayerInfo** | Pointer to [**GetPayTradeHistoryResponseDataInnerPayerInfo**](GetPayTradeHistoryResponseDataInnerPayerInfo.md) |  | [optional] 
**ReceiverInfo** | Pointer to [**GetPayTradeHistoryResponseDataInnerReceiverInfo**](GetPayTradeHistoryResponseDataInnerReceiverInfo.md) |  | [optional] 

## Methods

### NewGetPayTradeHistoryResponseDataInner

`func NewGetPayTradeHistoryResponseDataInner() *GetPayTradeHistoryResponseDataInner`

NewGetPayTradeHistoryResponseDataInner instantiates a new GetPayTradeHistoryResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayTradeHistoryResponseDataInnerWithDefaults

`func NewGetPayTradeHistoryResponseDataInnerWithDefaults() *GetPayTradeHistoryResponseDataInner`

NewGetPayTradeHistoryResponseDataInnerWithDefaults instantiates a new GetPayTradeHistoryResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderType

`func (o *GetPayTradeHistoryResponseDataInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetPayTradeHistoryResponseDataInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetPayTradeHistoryResponseDataInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetPayTradeHistoryResponseDataInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTransactionId

`func (o *GetPayTradeHistoryResponseDataInner) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *GetPayTradeHistoryResponseDataInner) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *GetPayTradeHistoryResponseDataInner) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *GetPayTradeHistoryResponseDataInner) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransactionTime

`func (o *GetPayTradeHistoryResponseDataInner) GetTransactionTime() int64`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *GetPayTradeHistoryResponseDataInner) GetTransactionTimeOk() (*int64, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *GetPayTradeHistoryResponseDataInner) SetTransactionTime(v int64)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *GetPayTradeHistoryResponseDataInner) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### GetAmount

`func (o *GetPayTradeHistoryResponseDataInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPayTradeHistoryResponseDataInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPayTradeHistoryResponseDataInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetPayTradeHistoryResponseDataInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *GetPayTradeHistoryResponseDataInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPayTradeHistoryResponseDataInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPayTradeHistoryResponseDataInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetPayTradeHistoryResponseDataInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetWalletType

`func (o *GetPayTradeHistoryResponseDataInner) GetWalletType() int64`

GetWalletType returns the WalletType field if non-nil, zero value otherwise.

### GetWalletTypeOk

`func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypeOk() (*int64, bool)`

GetWalletTypeOk returns a tuple with the WalletType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletType

`func (o *GetPayTradeHistoryResponseDataInner) SetWalletType(v int64)`

SetWalletType sets WalletType field to given value.

### HasWalletType

`func (o *GetPayTradeHistoryResponseDataInner) HasWalletType() bool`

HasWalletType returns a boolean if a field has been set.

### GetWalletTypes

`func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypes() []int64`

GetWalletTypes returns the WalletTypes field if non-nil, zero value otherwise.

### GetWalletTypesOk

`func (o *GetPayTradeHistoryResponseDataInner) GetWalletTypesOk() (*[]int64, bool)`

GetWalletTypesOk returns a tuple with the WalletTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletTypes

`func (o *GetPayTradeHistoryResponseDataInner) SetWalletTypes(v []int64)`

SetWalletTypes sets WalletTypes field to given value.

### HasWalletTypes

`func (o *GetPayTradeHistoryResponseDataInner) HasWalletTypes() bool`

HasWalletTypes returns a boolean if a field has been set.

### GetFundsDetail

`func (o *GetPayTradeHistoryResponseDataInner) GetFundsDetail() []GetPayTradeHistoryResponseDataInnerFundsDetailInner`

GetFundsDetail returns the FundsDetail field if non-nil, zero value otherwise.

### GetFundsDetailOk

`func (o *GetPayTradeHistoryResponseDataInner) GetFundsDetailOk() (*[]GetPayTradeHistoryResponseDataInnerFundsDetailInner, bool)`

GetFundsDetailOk returns a tuple with the FundsDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundsDetail

`func (o *GetPayTradeHistoryResponseDataInner) SetFundsDetail(v []GetPayTradeHistoryResponseDataInnerFundsDetailInner)`

SetFundsDetail sets FundsDetail field to given value.

### HasFundsDetail

`func (o *GetPayTradeHistoryResponseDataInner) HasFundsDetail() bool`

HasFundsDetail returns a boolean if a field has been set.

### GetPayerInfo

`func (o *GetPayTradeHistoryResponseDataInner) GetPayerInfo() GetPayTradeHistoryResponseDataInnerPayerInfo`

GetPayerInfo returns the PayerInfo field if non-nil, zero value otherwise.

### GetPayerInfoOk

`func (o *GetPayTradeHistoryResponseDataInner) GetPayerInfoOk() (*GetPayTradeHistoryResponseDataInnerPayerInfo, bool)`

GetPayerInfoOk returns a tuple with the PayerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerInfo

`func (o *GetPayTradeHistoryResponseDataInner) SetPayerInfo(v GetPayTradeHistoryResponseDataInnerPayerInfo)`

SetPayerInfo sets PayerInfo field to given value.

### HasPayerInfo

`func (o *GetPayTradeHistoryResponseDataInner) HasPayerInfo() bool`

HasPayerInfo returns a boolean if a field has been set.

### GetReceiverInfo

`func (o *GetPayTradeHistoryResponseDataInner) GetReceiverInfo() GetPayTradeHistoryResponseDataInnerReceiverInfo`

GetReceiverInfo returns the ReceiverInfo field if non-nil, zero value otherwise.

### GetReceiverInfoOk

`func (o *GetPayTradeHistoryResponseDataInner) GetReceiverInfoOk() (*GetPayTradeHistoryResponseDataInnerReceiverInfo, bool)`

GetReceiverInfoOk returns a tuple with the ReceiverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverInfo

`func (o *GetPayTradeHistoryResponseDataInner) SetReceiverInfo(v GetPayTradeHistoryResponseDataInnerReceiverInfo)`

SetReceiverInfo sets ReceiverInfo field to given value.

### HasReceiverInfo

`func (o *GetPayTradeHistoryResponseDataInner) HasReceiverInfo() bool`

HasReceiverInfo returns a boolean if a field has been set.


[[Back to README]](../README.md)


