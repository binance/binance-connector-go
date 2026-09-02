# GetQuoteResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Chance** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**MarketTitle** | Pointer to **string** |  | [optional] 
**MarketExtId** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**AmountIn** | Pointer to **string** |  | [optional] 
**AmountOut** | Pointer to **string** |  | [optional] 
**IsMinAmountOut** | Pointer to **bool** |  | [optional] 
**FeeAmount** | Pointer to **string** |  | [optional] 
**FeeDiscountBps** | Pointer to **string** |  | [optional] 
**AveragePrice** | Pointer to **float64** |  | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**PriceImpact** | Pointer to **float64** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**WalletAddress** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**SlippageBps** | Pointer to **int32** |  | [optional] 
**FeeRateBps** | Pointer to **int32** |  | [optional] 
**MinReceive** | Pointer to **string** |  | [optional] 
**ExpireAt** | Pointer to **int64** |  | [optional] 
**PriceLimit** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetQuoteResponse

`func NewGetQuoteResponse() *GetQuoteResponse`

NewGetQuoteResponse instantiates a new GetQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuoteResponseWithDefaults

`func NewGetQuoteResponseWithDefaults() *GetQuoteResponse`

NewGetQuoteResponseWithDefaults instantiates a new GetQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *GetQuoteResponse) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *GetQuoteResponse) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *GetQuoteResponse) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *GetQuoteResponse) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetTokenId

`func (o *GetQuoteResponse) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetQuoteResponse) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetQuoteResponse) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetQuoteResponse) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetChance

`func (o *GetQuoteResponse) GetChance() string`

GetChance returns the Chance field if non-nil, zero value otherwise.

### GetChanceOk

`func (o *GetQuoteResponse) GetChanceOk() (*string, bool)`

GetChanceOk returns a tuple with the Chance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChance

`func (o *GetQuoteResponse) SetChance(v string)`

SetChance sets Chance field to given value.

### HasChance

`func (o *GetQuoteResponse) HasChance() bool`

HasChance returns a boolean if a field has been set.

### GetVendor

`func (o *GetQuoteResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetQuoteResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetQuoteResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetQuoteResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetMarketTitle

`func (o *GetQuoteResponse) GetMarketTitle() string`

GetMarketTitle returns the MarketTitle field if non-nil, zero value otherwise.

### GetMarketTitleOk

`func (o *GetQuoteResponse) GetMarketTitleOk() (*string, bool)`

GetMarketTitleOk returns a tuple with the MarketTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTitle

`func (o *GetQuoteResponse) SetMarketTitle(v string)`

SetMarketTitle sets MarketTitle field to given value.

### HasMarketTitle

`func (o *GetQuoteResponse) HasMarketTitle() bool`

HasMarketTitle returns a boolean if a field has been set.

### GetMarketExtId

`func (o *GetQuoteResponse) GetMarketExtId() string`

GetMarketExtId returns the MarketExtId field if non-nil, zero value otherwise.

### GetMarketExtIdOk

`func (o *GetQuoteResponse) GetMarketExtIdOk() (*string, bool)`

GetMarketExtIdOk returns a tuple with the MarketExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketExtId

`func (o *GetQuoteResponse) SetMarketExtId(v string)`

SetMarketExtId sets MarketExtId field to given value.

### HasMarketExtId

`func (o *GetQuoteResponse) HasMarketExtId() bool`

HasMarketExtId returns a boolean if a field has been set.

### GetSide

`func (o *GetQuoteResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetQuoteResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetQuoteResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetQuoteResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetAmountIn

`func (o *GetQuoteResponse) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *GetQuoteResponse) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *GetQuoteResponse) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.

### HasAmountIn

`func (o *GetQuoteResponse) HasAmountIn() bool`

HasAmountIn returns a boolean if a field has been set.

### GetAmountOut

`func (o *GetQuoteResponse) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *GetQuoteResponse) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *GetQuoteResponse) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.

### HasAmountOut

`func (o *GetQuoteResponse) HasAmountOut() bool`

HasAmountOut returns a boolean if a field has been set.

### GetIsMinAmountOut

`func (o *GetQuoteResponse) GetIsMinAmountOut() bool`

GetIsMinAmountOut returns the IsMinAmountOut field if non-nil, zero value otherwise.

### GetIsMinAmountOutOk

`func (o *GetQuoteResponse) GetIsMinAmountOutOk() (*bool, bool)`

GetIsMinAmountOutOk returns a tuple with the IsMinAmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMinAmountOut

`func (o *GetQuoteResponse) SetIsMinAmountOut(v bool)`

SetIsMinAmountOut sets IsMinAmountOut field to given value.

### HasIsMinAmountOut

`func (o *GetQuoteResponse) HasIsMinAmountOut() bool`

HasIsMinAmountOut returns a boolean if a field has been set.

### GetFeeAmount

`func (o *GetQuoteResponse) GetFeeAmount() string`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *GetQuoteResponse) GetFeeAmountOk() (*string, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *GetQuoteResponse) SetFeeAmount(v string)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *GetQuoteResponse) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetFeeDiscountBps

`func (o *GetQuoteResponse) GetFeeDiscountBps() string`

GetFeeDiscountBps returns the FeeDiscountBps field if non-nil, zero value otherwise.

### GetFeeDiscountBpsOk

`func (o *GetQuoteResponse) GetFeeDiscountBpsOk() (*string, bool)`

GetFeeDiscountBpsOk returns a tuple with the FeeDiscountBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeDiscountBps

`func (o *GetQuoteResponse) SetFeeDiscountBps(v string)`

SetFeeDiscountBps sets FeeDiscountBps field to given value.

### HasFeeDiscountBps

`func (o *GetQuoteResponse) HasFeeDiscountBps() bool`

HasFeeDiscountBps returns a boolean if a field has been set.

### GetAveragePrice

`func (o *GetQuoteResponse) GetAveragePrice() float64`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *GetQuoteResponse) GetAveragePriceOk() (*float64, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *GetQuoteResponse) SetAveragePrice(v float64)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *GetQuoteResponse) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *GetQuoteResponse) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *GetQuoteResponse) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *GetQuoteResponse) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *GetQuoteResponse) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetPriceImpact

`func (o *GetQuoteResponse) GetPriceImpact() float64`

GetPriceImpact returns the PriceImpact field if non-nil, zero value otherwise.

### GetPriceImpactOk

`func (o *GetQuoteResponse) GetPriceImpactOk() (*float64, bool)`

GetPriceImpactOk returns a tuple with the PriceImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceImpact

`func (o *GetQuoteResponse) SetPriceImpact(v float64)`

SetPriceImpact sets PriceImpact field to given value.

### HasPriceImpact

`func (o *GetQuoteResponse) HasPriceImpact() bool`

HasPriceImpact returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetQuoteResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetQuoteResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetQuoteResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetQuoteResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetChainId

`func (o *GetQuoteResponse) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *GetQuoteResponse) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *GetQuoteResponse) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *GetQuoteResponse) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetUserId

`func (o *GetQuoteResponse) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetQuoteResponse) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetQuoteResponse) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetQuoteResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWalletAddress

`func (o *GetQuoteResponse) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GetQuoteResponse) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GetQuoteResponse) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *GetQuoteResponse) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetOrderType

`func (o *GetQuoteResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetQuoteResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetQuoteResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetQuoteResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSlippageBps

`func (o *GetQuoteResponse) GetSlippageBps() int32`

GetSlippageBps returns the SlippageBps field if non-nil, zero value otherwise.

### GetSlippageBpsOk

`func (o *GetQuoteResponse) GetSlippageBpsOk() (*int32, bool)`

GetSlippageBpsOk returns a tuple with the SlippageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlippageBps

`func (o *GetQuoteResponse) SetSlippageBps(v int32)`

SetSlippageBps sets SlippageBps field to given value.

### HasSlippageBps

`func (o *GetQuoteResponse) HasSlippageBps() bool`

HasSlippageBps returns a boolean if a field has been set.

### GetFeeRateBps

`func (o *GetQuoteResponse) GetFeeRateBps() int32`

GetFeeRateBps returns the FeeRateBps field if non-nil, zero value otherwise.

### GetFeeRateBpsOk

`func (o *GetQuoteResponse) GetFeeRateBpsOk() (*int32, bool)`

GetFeeRateBpsOk returns a tuple with the FeeRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeRateBps

`func (o *GetQuoteResponse) SetFeeRateBps(v int32)`

SetFeeRateBps sets FeeRateBps field to given value.

### HasFeeRateBps

`func (o *GetQuoteResponse) HasFeeRateBps() bool`

HasFeeRateBps returns a boolean if a field has been set.

### GetMinReceive

`func (o *GetQuoteResponse) GetMinReceive() string`

GetMinReceive returns the MinReceive field if non-nil, zero value otherwise.

### GetMinReceiveOk

`func (o *GetQuoteResponse) GetMinReceiveOk() (*string, bool)`

GetMinReceiveOk returns a tuple with the MinReceive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReceive

`func (o *GetQuoteResponse) SetMinReceive(v string)`

SetMinReceive sets MinReceive field to given value.

### HasMinReceive

`func (o *GetQuoteResponse) HasMinReceive() bool`

HasMinReceive returns a boolean if a field has been set.

### GetExpireAt

`func (o *GetQuoteResponse) GetExpireAt() int64`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *GetQuoteResponse) GetExpireAtOk() (*int64, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *GetQuoteResponse) SetExpireAt(v int64)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *GetQuoteResponse) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetPriceLimit

`func (o *GetQuoteResponse) GetPriceLimit() string`

GetPriceLimit returns the PriceLimit field if non-nil, zero value otherwise.

### GetPriceLimitOk

`func (o *GetQuoteResponse) GetPriceLimitOk() (*string, bool)`

GetPriceLimitOk returns a tuple with the PriceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLimit

`func (o *GetQuoteResponse) SetPriceLimit(v string)`

SetPriceLimit sets PriceLimit field to given value.

### HasPriceLimit

`func (o *GetQuoteResponse) HasPriceLimit() bool`

HasPriceLimit returns a boolean if a field has been set.

### SetPriceLimitNil

`func (o *GetQuoteResponse) SetPriceLimitNil(b bool)`

 SetPriceLimitNil sets the value for PriceLimit to be an explicit nil

### UnsetPriceLimit
`func (o *GetQuoteResponse) UnsetPriceLimit()`

UnsetPriceLimit ensures that no value is present for PriceLimit, not even an explicit nil

[[Back to README]](../README.md)


