# GetNFTTransactionHistoryResponseListInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderNo** | Pointer to **string** |  | [optional] 
**Tokens** | Pointer to [**[]GetNFTTransactionHistoryResponseListInnerTokensInner**](GetNFTTransactionHistoryResponseListInnerTokensInner.md) |  | [optional] 
**TradeTime** | Pointer to **int64** |  | [optional] 
**TradeAmount** | Pointer to **string** |  | [optional] 
**TradeCurrency** | Pointer to **string** |  | [optional] 

## Methods

### NewGetNFTTransactionHistoryResponseListInner

`func NewGetNFTTransactionHistoryResponseListInner() *GetNFTTransactionHistoryResponseListInner`

NewGetNFTTransactionHistoryResponseListInner instantiates a new GetNFTTransactionHistoryResponseListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNFTTransactionHistoryResponseListInnerWithDefaults

`func NewGetNFTTransactionHistoryResponseListInnerWithDefaults() *GetNFTTransactionHistoryResponseListInner`

NewGetNFTTransactionHistoryResponseListInnerWithDefaults instantiates a new GetNFTTransactionHistoryResponseListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNo

`func (o *GetNFTTransactionHistoryResponseListInner) GetOrderNo() string`

GetOrderNo returns the OrderNo field if non-nil, zero value otherwise.

### GetOrderNoOk

`func (o *GetNFTTransactionHistoryResponseListInner) GetOrderNoOk() (*string, bool)`

GetOrderNoOk returns a tuple with the OrderNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNo

`func (o *GetNFTTransactionHistoryResponseListInner) SetOrderNo(v string)`

SetOrderNo sets OrderNo field to given value.

### HasOrderNo

`func (o *GetNFTTransactionHistoryResponseListInner) HasOrderNo() bool`

HasOrderNo returns a boolean if a field has been set.

### GetTokens

`func (o *GetNFTTransactionHistoryResponseListInner) GetTokens() []GetNFTTransactionHistoryResponseListInnerTokensInner`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GetNFTTransactionHistoryResponseListInner) GetTokensOk() (*[]GetNFTTransactionHistoryResponseListInnerTokensInner, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GetNFTTransactionHistoryResponseListInner) SetTokens(v []GetNFTTransactionHistoryResponseListInnerTokensInner)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *GetNFTTransactionHistoryResponseListInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTradeTime

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeTime() int64`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeTimeOk() (*int64, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *GetNFTTransactionHistoryResponseListInner) SetTradeTime(v int64)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *GetNFTTransactionHistoryResponseListInner) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.

### GetTradeAmount

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeAmount() string`

GetTradeAmount returns the TradeAmount field if non-nil, zero value otherwise.

### GetTradeAmountOk

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeAmountOk() (*string, bool)`

GetTradeAmountOk returns a tuple with the TradeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmount

`func (o *GetNFTTransactionHistoryResponseListInner) SetTradeAmount(v string)`

SetTradeAmount sets TradeAmount field to given value.

### HasTradeAmount

`func (o *GetNFTTransactionHistoryResponseListInner) HasTradeAmount() bool`

HasTradeAmount returns a boolean if a field has been set.

### GetTradeCurrency

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeCurrency() string`

GetTradeCurrency returns the TradeCurrency field if non-nil, zero value otherwise.

### GetTradeCurrencyOk

`func (o *GetNFTTransactionHistoryResponseListInner) GetTradeCurrencyOk() (*string, bool)`

GetTradeCurrencyOk returns a tuple with the TradeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeCurrency

`func (o *GetNFTTransactionHistoryResponseListInner) SetTradeCurrency(v string)`

SetTradeCurrency sets TradeCurrency field to given value.

### HasTradeCurrency

`func (o *GetNFTTransactionHistoryResponseListInner) HasTradeCurrency() bool`

HasTradeCurrency returns a boolean if a field has been set.


[[Back to README]](../README.md)


