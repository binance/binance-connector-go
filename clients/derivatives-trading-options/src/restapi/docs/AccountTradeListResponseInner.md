# AccountTradeListResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TradeId** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**RealizedProfit** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Volatility** | Pointer to **string** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**PriceScale** | Pointer to **int64** |  | [optional] 
**QuantityScale** | Pointer to **int64** |  | [optional] 
**OptionSide** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountTradeListResponseInner

`func NewAccountTradeListResponseInner() *AccountTradeListResponseInner`

NewAccountTradeListResponseInner instantiates a new AccountTradeListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTradeListResponseInnerWithDefaults

`func NewAccountTradeListResponseInnerWithDefaults() *AccountTradeListResponseInner`

NewAccountTradeListResponseInnerWithDefaults instantiates a new AccountTradeListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountTradeListResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountTradeListResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountTradeListResponseInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AccountTradeListResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTradeId

`func (o *AccountTradeListResponseInner) GetTradeId() int64`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *AccountTradeListResponseInner) GetTradeIdOk() (*int64, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *AccountTradeListResponseInner) SetTradeId(v int64)`

SetTradeId sets TradeId field to given value.

### HasTradeId

`func (o *AccountTradeListResponseInner) HasTradeId() bool`

HasTradeId returns a boolean if a field has been set.

### GetOrderId

`func (o *AccountTradeListResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AccountTradeListResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AccountTradeListResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AccountTradeListResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSymbol

`func (o *AccountTradeListResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountTradeListResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountTradeListResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountTradeListResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPrice

`func (o *AccountTradeListResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AccountTradeListResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AccountTradeListResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AccountTradeListResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *AccountTradeListResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AccountTradeListResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AccountTradeListResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AccountTradeListResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFee

`func (o *AccountTradeListResponseInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *AccountTradeListResponseInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *AccountTradeListResponseInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *AccountTradeListResponseInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetRealizedProfit

`func (o *AccountTradeListResponseInner) GetRealizedProfit() string`

GetRealizedProfit returns the RealizedProfit field if non-nil, zero value otherwise.

### GetRealizedProfitOk

`func (o *AccountTradeListResponseInner) GetRealizedProfitOk() (*string, bool)`

GetRealizedProfitOk returns a tuple with the RealizedProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedProfit

`func (o *AccountTradeListResponseInner) SetRealizedProfit(v string)`

SetRealizedProfit sets RealizedProfit field to given value.

### HasRealizedProfit

`func (o *AccountTradeListResponseInner) HasRealizedProfit() bool`

HasRealizedProfit returns a boolean if a field has been set.

### GetSide

`func (o *AccountTradeListResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *AccountTradeListResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *AccountTradeListResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *AccountTradeListResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *AccountTradeListResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountTradeListResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountTradeListResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountTradeListResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVolatility

`func (o *AccountTradeListResponseInner) GetVolatility() string`

GetVolatility returns the Volatility field if non-nil, zero value otherwise.

### GetVolatilityOk

`func (o *AccountTradeListResponseInner) GetVolatilityOk() (*string, bool)`

GetVolatilityOk returns a tuple with the Volatility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolatility

`func (o *AccountTradeListResponseInner) SetVolatility(v string)`

SetVolatility sets Volatility field to given value.

### HasVolatility

`func (o *AccountTradeListResponseInner) HasVolatility() bool`

HasVolatility returns a boolean if a field has been set.

### GetLiquidity

`func (o *AccountTradeListResponseInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *AccountTradeListResponseInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *AccountTradeListResponseInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *AccountTradeListResponseInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *AccountTradeListResponseInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *AccountTradeListResponseInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *AccountTradeListResponseInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *AccountTradeListResponseInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetTime

`func (o *AccountTradeListResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AccountTradeListResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AccountTradeListResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AccountTradeListResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetPriceScale

`func (o *AccountTradeListResponseInner) GetPriceScale() int64`

GetPriceScale returns the PriceScale field if non-nil, zero value otherwise.

### GetPriceScaleOk

`func (o *AccountTradeListResponseInner) GetPriceScaleOk() (*int64, bool)`

GetPriceScaleOk returns a tuple with the PriceScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceScale

`func (o *AccountTradeListResponseInner) SetPriceScale(v int64)`

SetPriceScale sets PriceScale field to given value.

### HasPriceScale

`func (o *AccountTradeListResponseInner) HasPriceScale() bool`

HasPriceScale returns a boolean if a field has been set.

### GetQuantityScale

`func (o *AccountTradeListResponseInner) GetQuantityScale() int64`

GetQuantityScale returns the QuantityScale field if non-nil, zero value otherwise.

### GetQuantityScaleOk

`func (o *AccountTradeListResponseInner) GetQuantityScaleOk() (*int64, bool)`

GetQuantityScaleOk returns a tuple with the QuantityScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityScale

`func (o *AccountTradeListResponseInner) SetQuantityScale(v int64)`

SetQuantityScale sets QuantityScale field to given value.

### HasQuantityScale

`func (o *AccountTradeListResponseInner) HasQuantityScale() bool`

HasQuantityScale returns a boolean if a field has been set.

### GetOptionSide

`func (o *AccountTradeListResponseInner) GetOptionSide() string`

GetOptionSide returns the OptionSide field if non-nil, zero value otherwise.

### GetOptionSideOk

`func (o *AccountTradeListResponseInner) GetOptionSideOk() (*string, bool)`

GetOptionSideOk returns a tuple with the OptionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSide

`func (o *AccountTradeListResponseInner) SetOptionSide(v string)`

SetOptionSide sets OptionSide field to given value.

### HasOptionSide

`func (o *AccountTradeListResponseInner) HasOptionSide() bool`

HasOptionSide returns a boolean if a field has been set.


[[Back to README]](../README.md)


