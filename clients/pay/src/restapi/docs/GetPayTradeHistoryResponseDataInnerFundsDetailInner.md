# GetPayTradeHistoryResponseDataInnerFundsDetailInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Asset. | [optional] 
**Amount** | Pointer to **string** | Asset amount. | [optional] 
**WalletAssetCost** | Pointer to **map[string]string** | Asset cost details per wallet type. Keys are wallet type IDs (e.g. \&quot;1\&quot;, \&quot;2\&quot;), values are cost amounts. | [optional] 

## Methods

### NewGetPayTradeHistoryResponseDataInnerFundsDetailInner

`func NewGetPayTradeHistoryResponseDataInnerFundsDetailInner() *GetPayTradeHistoryResponseDataInnerFundsDetailInner`

NewGetPayTradeHistoryResponseDataInnerFundsDetailInner instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWithDefaults

`func NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWithDefaults() *GetPayTradeHistoryResponseDataInnerFundsDetailInner`

NewGetPayTradeHistoryResponseDataInnerFundsDetailInnerWithDefaults instantiates a new GetPayTradeHistoryResponseDataInnerFundsDetailInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAmount

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetWalletAssetCost

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetWalletAssetCost() map[string]string`

GetWalletAssetCost returns the WalletAssetCost field if non-nil, zero value otherwise.

### GetWalletAssetCostOk

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) GetWalletAssetCostOk() (*map[string]string, bool)`

GetWalletAssetCostOk returns a tuple with the WalletAssetCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAssetCost

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) SetWalletAssetCost(v map[string]string)`

SetWalletAssetCost sets WalletAssetCost field to given value.

### HasWalletAssetCost

`func (o *GetPayTradeHistoryResponseDataInnerFundsDetailInner) HasWalletAssetCost() bool`

HasWalletAssetCost returns a boolean if a field has been set.


[[Back to README]](../README.md)


