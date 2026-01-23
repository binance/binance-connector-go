# TokenListResponseDataInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **string** |  | [optional] 
**ChainIconUrl** | Pointer to **string** |  | [optional] 
**ChainName** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**PercentChange24h** | Pointer to **string** |  | [optional] 
**Volume24h** | Pointer to **string** |  | [optional] 
**MarketCap** | Pointer to **string** |  | [optional] 
**Fdv** | Pointer to **string** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**TotalSupply** | Pointer to **string** |  | [optional] 
**CirculatingSupply** | Pointer to **string** |  | [optional] 
**Holders** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int64** |  | [optional] 
**ListingCex** | Pointer to **bool** |  | [optional] 
**HotTag** | Pointer to **bool** |  | [optional] 
**CexCoinName** | Pointer to **string** |  | [optional] 
**CanTransfer** | Pointer to **bool** |  | [optional] 
**Denomination** | Pointer to **int64** |  | [optional] 
**Offline** | Pointer to **bool** |  | [optional] 
**TradeDecimal** | Pointer to **int64** |  | [optional] 
**AlphaId** | Pointer to **string** |  | [optional] 
**Offsell** | Pointer to **bool** |  | [optional] 
**PriceHigh24h** | Pointer to **string** |  | [optional] 
**PriceLow24h** | Pointer to **string** |  | [optional] 
**Count24h** | Pointer to **string** |  | [optional] 
**OnlineTge** | Pointer to **bool** |  | [optional] 
**OnlineAirdrop** | Pointer to **bool** |  | [optional] 
**Score** | Pointer to **int64** |  | [optional] 
**CexOffDisplay** | Pointer to **bool** |  | [optional] 
**StockState** | Pointer to **bool** |  | [optional] 
**ListingTime** | Pointer to **int64** |  | [optional] 
**MulPoint** | Pointer to **int64** |  | [optional] 
**BnExclusiveState** | Pointer to **bool** |  | [optional] 

## Methods

### NewTokenListResponseDataInner

`func NewTokenListResponseDataInner() *TokenListResponseDataInner`

NewTokenListResponseDataInner instantiates a new TokenListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListResponseDataInnerWithDefaults

`func NewTokenListResponseDataInnerWithDefaults() *TokenListResponseDataInner`

NewTokenListResponseDataInnerWithDefaults instantiates a new TokenListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *TokenListResponseDataInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenListResponseDataInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenListResponseDataInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenListResponseDataInner) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetChainId

`func (o *TokenListResponseDataInner) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *TokenListResponseDataInner) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *TokenListResponseDataInner) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *TokenListResponseDataInner) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetChainIconUrl

`func (o *TokenListResponseDataInner) GetChainIconUrl() string`

GetChainIconUrl returns the ChainIconUrl field if non-nil, zero value otherwise.

### GetChainIconUrlOk

`func (o *TokenListResponseDataInner) GetChainIconUrlOk() (*string, bool)`

GetChainIconUrlOk returns a tuple with the ChainIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainIconUrl

`func (o *TokenListResponseDataInner) SetChainIconUrl(v string)`

SetChainIconUrl sets ChainIconUrl field to given value.

### HasChainIconUrl

`func (o *TokenListResponseDataInner) HasChainIconUrl() bool`

HasChainIconUrl returns a boolean if a field has been set.

### GetChainName

`func (o *TokenListResponseDataInner) GetChainName() string`

GetChainName returns the ChainName field if non-nil, zero value otherwise.

### GetChainNameOk

`func (o *TokenListResponseDataInner) GetChainNameOk() (*string, bool)`

GetChainNameOk returns a tuple with the ChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainName

`func (o *TokenListResponseDataInner) SetChainName(v string)`

SetChainName sets ChainName field to given value.

### HasChainName

`func (o *TokenListResponseDataInner) HasChainName() bool`

HasChainName returns a boolean if a field has been set.

### GetContractAddress

`func (o *TokenListResponseDataInner) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *TokenListResponseDataInner) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *TokenListResponseDataInner) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *TokenListResponseDataInner) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetName

`func (o *TokenListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenListResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *TokenListResponseDataInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenListResponseDataInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenListResponseDataInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenListResponseDataInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetIconUrl

`func (o *TokenListResponseDataInner) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *TokenListResponseDataInner) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *TokenListResponseDataInner) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *TokenListResponseDataInner) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetPrice

`func (o *TokenListResponseDataInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TokenListResponseDataInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TokenListResponseDataInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TokenListResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPercentChange24h

`func (o *TokenListResponseDataInner) GetPercentChange24h() string`

GetPercentChange24h returns the PercentChange24h field if non-nil, zero value otherwise.

### GetPercentChange24hOk

`func (o *TokenListResponseDataInner) GetPercentChange24hOk() (*string, bool)`

GetPercentChange24hOk returns a tuple with the PercentChange24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentChange24h

`func (o *TokenListResponseDataInner) SetPercentChange24h(v string)`

SetPercentChange24h sets PercentChange24h field to given value.

### HasPercentChange24h

`func (o *TokenListResponseDataInner) HasPercentChange24h() bool`

HasPercentChange24h returns a boolean if a field has been set.

### GetVolume24h

`func (o *TokenListResponseDataInner) GetVolume24h() string`

GetVolume24h returns the Volume24h field if non-nil, zero value otherwise.

### GetVolume24hOk

`func (o *TokenListResponseDataInner) GetVolume24hOk() (*string, bool)`

GetVolume24hOk returns a tuple with the Volume24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume24h

`func (o *TokenListResponseDataInner) SetVolume24h(v string)`

SetVolume24h sets Volume24h field to given value.

### HasVolume24h

`func (o *TokenListResponseDataInner) HasVolume24h() bool`

HasVolume24h returns a boolean if a field has been set.

### GetMarketCap

`func (o *TokenListResponseDataInner) GetMarketCap() string`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *TokenListResponseDataInner) GetMarketCapOk() (*string, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *TokenListResponseDataInner) SetMarketCap(v string)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *TokenListResponseDataInner) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetFdv

`func (o *TokenListResponseDataInner) GetFdv() string`

GetFdv returns the Fdv field if non-nil, zero value otherwise.

### GetFdvOk

`func (o *TokenListResponseDataInner) GetFdvOk() (*string, bool)`

GetFdvOk returns a tuple with the Fdv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdv

`func (o *TokenListResponseDataInner) SetFdv(v string)`

SetFdv sets Fdv field to given value.

### HasFdv

`func (o *TokenListResponseDataInner) HasFdv() bool`

HasFdv returns a boolean if a field has been set.

### GetLiquidity

`func (o *TokenListResponseDataInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *TokenListResponseDataInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *TokenListResponseDataInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *TokenListResponseDataInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetTotalSupply

`func (o *TokenListResponseDataInner) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *TokenListResponseDataInner) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *TokenListResponseDataInner) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.

### HasTotalSupply

`func (o *TokenListResponseDataInner) HasTotalSupply() bool`

HasTotalSupply returns a boolean if a field has been set.

### GetCirculatingSupply

`func (o *TokenListResponseDataInner) GetCirculatingSupply() string`

GetCirculatingSupply returns the CirculatingSupply field if non-nil, zero value otherwise.

### GetCirculatingSupplyOk

`func (o *TokenListResponseDataInner) GetCirculatingSupplyOk() (*string, bool)`

GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCirculatingSupply

`func (o *TokenListResponseDataInner) SetCirculatingSupply(v string)`

SetCirculatingSupply sets CirculatingSupply field to given value.

### HasCirculatingSupply

`func (o *TokenListResponseDataInner) HasCirculatingSupply() bool`

HasCirculatingSupply returns a boolean if a field has been set.

### GetHolders

`func (o *TokenListResponseDataInner) GetHolders() string`

GetHolders returns the Holders field if non-nil, zero value otherwise.

### GetHoldersOk

`func (o *TokenListResponseDataInner) GetHoldersOk() (*string, bool)`

GetHoldersOk returns a tuple with the Holders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolders

`func (o *TokenListResponseDataInner) SetHolders(v string)`

SetHolders sets Holders field to given value.

### HasHolders

`func (o *TokenListResponseDataInner) HasHolders() bool`

HasHolders returns a boolean if a field has been set.

### GetDecimals

`func (o *TokenListResponseDataInner) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenListResponseDataInner) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenListResponseDataInner) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenListResponseDataInner) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetListingCex

`func (o *TokenListResponseDataInner) GetListingCex() bool`

GetListingCex returns the ListingCex field if non-nil, zero value otherwise.

### GetListingCexOk

`func (o *TokenListResponseDataInner) GetListingCexOk() (*bool, bool)`

GetListingCexOk returns a tuple with the ListingCex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingCex

`func (o *TokenListResponseDataInner) SetListingCex(v bool)`

SetListingCex sets ListingCex field to given value.

### HasListingCex

`func (o *TokenListResponseDataInner) HasListingCex() bool`

HasListingCex returns a boolean if a field has been set.

### GetHotTag

`func (o *TokenListResponseDataInner) GetHotTag() bool`

GetHotTag returns the HotTag field if non-nil, zero value otherwise.

### GetHotTagOk

`func (o *TokenListResponseDataInner) GetHotTagOk() (*bool, bool)`

GetHotTagOk returns a tuple with the HotTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotTag

`func (o *TokenListResponseDataInner) SetHotTag(v bool)`

SetHotTag sets HotTag field to given value.

### HasHotTag

`func (o *TokenListResponseDataInner) HasHotTag() bool`

HasHotTag returns a boolean if a field has been set.

### GetCexCoinName

`func (o *TokenListResponseDataInner) GetCexCoinName() string`

GetCexCoinName returns the CexCoinName field if non-nil, zero value otherwise.

### GetCexCoinNameOk

`func (o *TokenListResponseDataInner) GetCexCoinNameOk() (*string, bool)`

GetCexCoinNameOk returns a tuple with the CexCoinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCexCoinName

`func (o *TokenListResponseDataInner) SetCexCoinName(v string)`

SetCexCoinName sets CexCoinName field to given value.

### HasCexCoinName

`func (o *TokenListResponseDataInner) HasCexCoinName() bool`

HasCexCoinName returns a boolean if a field has been set.

### GetCanTransfer

`func (o *TokenListResponseDataInner) GetCanTransfer() bool`

GetCanTransfer returns the CanTransfer field if non-nil, zero value otherwise.

### GetCanTransferOk

`func (o *TokenListResponseDataInner) GetCanTransferOk() (*bool, bool)`

GetCanTransferOk returns a tuple with the CanTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTransfer

`func (o *TokenListResponseDataInner) SetCanTransfer(v bool)`

SetCanTransfer sets CanTransfer field to given value.

### HasCanTransfer

`func (o *TokenListResponseDataInner) HasCanTransfer() bool`

HasCanTransfer returns a boolean if a field has been set.

### GetDenomination

`func (o *TokenListResponseDataInner) GetDenomination() int64`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *TokenListResponseDataInner) GetDenominationOk() (*int64, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *TokenListResponseDataInner) SetDenomination(v int64)`

SetDenomination sets Denomination field to given value.

### HasDenomination

`func (o *TokenListResponseDataInner) HasDenomination() bool`

HasDenomination returns a boolean if a field has been set.

### GetOffline

`func (o *TokenListResponseDataInner) GetOffline() bool`

GetOffline returns the Offline field if non-nil, zero value otherwise.

### GetOfflineOk

`func (o *TokenListResponseDataInner) GetOfflineOk() (*bool, bool)`

GetOfflineOk returns a tuple with the Offline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffline

`func (o *TokenListResponseDataInner) SetOffline(v bool)`

SetOffline sets Offline field to given value.

### HasOffline

`func (o *TokenListResponseDataInner) HasOffline() bool`

HasOffline returns a boolean if a field has been set.

### GetTradeDecimal

`func (o *TokenListResponseDataInner) GetTradeDecimal() int64`

GetTradeDecimal returns the TradeDecimal field if non-nil, zero value otherwise.

### GetTradeDecimalOk

`func (o *TokenListResponseDataInner) GetTradeDecimalOk() (*int64, bool)`

GetTradeDecimalOk returns a tuple with the TradeDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDecimal

`func (o *TokenListResponseDataInner) SetTradeDecimal(v int64)`

SetTradeDecimal sets TradeDecimal field to given value.

### HasTradeDecimal

`func (o *TokenListResponseDataInner) HasTradeDecimal() bool`

HasTradeDecimal returns a boolean if a field has been set.

### GetAlphaId

`func (o *TokenListResponseDataInner) GetAlphaId() string`

GetAlphaId returns the AlphaId field if non-nil, zero value otherwise.

### GetAlphaIdOk

`func (o *TokenListResponseDataInner) GetAlphaIdOk() (*string, bool)`

GetAlphaIdOk returns a tuple with the AlphaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphaId

`func (o *TokenListResponseDataInner) SetAlphaId(v string)`

SetAlphaId sets AlphaId field to given value.

### HasAlphaId

`func (o *TokenListResponseDataInner) HasAlphaId() bool`

HasAlphaId returns a boolean if a field has been set.

### GetOffsell

`func (o *TokenListResponseDataInner) GetOffsell() bool`

GetOffsell returns the Offsell field if non-nil, zero value otherwise.

### GetOffsellOk

`func (o *TokenListResponseDataInner) GetOffsellOk() (*bool, bool)`

GetOffsellOk returns a tuple with the Offsell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsell

`func (o *TokenListResponseDataInner) SetOffsell(v bool)`

SetOffsell sets Offsell field to given value.

### HasOffsell

`func (o *TokenListResponseDataInner) HasOffsell() bool`

HasOffsell returns a boolean if a field has been set.

### GetPriceHigh24h

`func (o *TokenListResponseDataInner) GetPriceHigh24h() string`

GetPriceHigh24h returns the PriceHigh24h field if non-nil, zero value otherwise.

### GetPriceHigh24hOk

`func (o *TokenListResponseDataInner) GetPriceHigh24hOk() (*string, bool)`

GetPriceHigh24hOk returns a tuple with the PriceHigh24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHigh24h

`func (o *TokenListResponseDataInner) SetPriceHigh24h(v string)`

SetPriceHigh24h sets PriceHigh24h field to given value.

### HasPriceHigh24h

`func (o *TokenListResponseDataInner) HasPriceHigh24h() bool`

HasPriceHigh24h returns a boolean if a field has been set.

### GetPriceLow24h

`func (o *TokenListResponseDataInner) GetPriceLow24h() string`

GetPriceLow24h returns the PriceLow24h field if non-nil, zero value otherwise.

### GetPriceLow24hOk

`func (o *TokenListResponseDataInner) GetPriceLow24hOk() (*string, bool)`

GetPriceLow24hOk returns a tuple with the PriceLow24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLow24h

`func (o *TokenListResponseDataInner) SetPriceLow24h(v string)`

SetPriceLow24h sets PriceLow24h field to given value.

### HasPriceLow24h

`func (o *TokenListResponseDataInner) HasPriceLow24h() bool`

HasPriceLow24h returns a boolean if a field has been set.

### GetCount24h

`func (o *TokenListResponseDataInner) GetCount24h() string`

GetCount24h returns the Count24h field if non-nil, zero value otherwise.

### GetCount24hOk

`func (o *TokenListResponseDataInner) GetCount24hOk() (*string, bool)`

GetCount24hOk returns a tuple with the Count24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount24h

`func (o *TokenListResponseDataInner) SetCount24h(v string)`

SetCount24h sets Count24h field to given value.

### HasCount24h

`func (o *TokenListResponseDataInner) HasCount24h() bool`

HasCount24h returns a boolean if a field has been set.

### GetOnlineTge

`func (o *TokenListResponseDataInner) GetOnlineTge() bool`

GetOnlineTge returns the OnlineTge field if non-nil, zero value otherwise.

### GetOnlineTgeOk

`func (o *TokenListResponseDataInner) GetOnlineTgeOk() (*bool, bool)`

GetOnlineTgeOk returns a tuple with the OnlineTge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineTge

`func (o *TokenListResponseDataInner) SetOnlineTge(v bool)`

SetOnlineTge sets OnlineTge field to given value.

### HasOnlineTge

`func (o *TokenListResponseDataInner) HasOnlineTge() bool`

HasOnlineTge returns a boolean if a field has been set.

### GetOnlineAirdrop

`func (o *TokenListResponseDataInner) GetOnlineAirdrop() bool`

GetOnlineAirdrop returns the OnlineAirdrop field if non-nil, zero value otherwise.

### GetOnlineAirdropOk

`func (o *TokenListResponseDataInner) GetOnlineAirdropOk() (*bool, bool)`

GetOnlineAirdropOk returns a tuple with the OnlineAirdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineAirdrop

`func (o *TokenListResponseDataInner) SetOnlineAirdrop(v bool)`

SetOnlineAirdrop sets OnlineAirdrop field to given value.

### HasOnlineAirdrop

`func (o *TokenListResponseDataInner) HasOnlineAirdrop() bool`

HasOnlineAirdrop returns a boolean if a field has been set.

### GetScore

`func (o *TokenListResponseDataInner) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TokenListResponseDataInner) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TokenListResponseDataInner) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *TokenListResponseDataInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetCexOffDisplay

`func (o *TokenListResponseDataInner) GetCexOffDisplay() bool`

GetCexOffDisplay returns the CexOffDisplay field if non-nil, zero value otherwise.

### GetCexOffDisplayOk

`func (o *TokenListResponseDataInner) GetCexOffDisplayOk() (*bool, bool)`

GetCexOffDisplayOk returns a tuple with the CexOffDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCexOffDisplay

`func (o *TokenListResponseDataInner) SetCexOffDisplay(v bool)`

SetCexOffDisplay sets CexOffDisplay field to given value.

### HasCexOffDisplay

`func (o *TokenListResponseDataInner) HasCexOffDisplay() bool`

HasCexOffDisplay returns a boolean if a field has been set.

### GetStockState

`func (o *TokenListResponseDataInner) GetStockState() bool`

GetStockState returns the StockState field if non-nil, zero value otherwise.

### GetStockStateOk

`func (o *TokenListResponseDataInner) GetStockStateOk() (*bool, bool)`

GetStockStateOk returns a tuple with the StockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockState

`func (o *TokenListResponseDataInner) SetStockState(v bool)`

SetStockState sets StockState field to given value.

### HasStockState

`func (o *TokenListResponseDataInner) HasStockState() bool`

HasStockState returns a boolean if a field has been set.

### GetListingTime

`func (o *TokenListResponseDataInner) GetListingTime() int64`

GetListingTime returns the ListingTime field if non-nil, zero value otherwise.

### GetListingTimeOk

`func (o *TokenListResponseDataInner) GetListingTimeOk() (*int64, bool)`

GetListingTimeOk returns a tuple with the ListingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingTime

`func (o *TokenListResponseDataInner) SetListingTime(v int64)`

SetListingTime sets ListingTime field to given value.

### HasListingTime

`func (o *TokenListResponseDataInner) HasListingTime() bool`

HasListingTime returns a boolean if a field has been set.

### GetMulPoint

`func (o *TokenListResponseDataInner) GetMulPoint() int64`

GetMulPoint returns the MulPoint field if non-nil, zero value otherwise.

### GetMulPointOk

`func (o *TokenListResponseDataInner) GetMulPointOk() (*int64, bool)`

GetMulPointOk returns a tuple with the MulPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulPoint

`func (o *TokenListResponseDataInner) SetMulPoint(v int64)`

SetMulPoint sets MulPoint field to given value.

### HasMulPoint

`func (o *TokenListResponseDataInner) HasMulPoint() bool`

HasMulPoint returns a boolean if a field has been set.

### GetBnExclusiveState

`func (o *TokenListResponseDataInner) GetBnExclusiveState() bool`

GetBnExclusiveState returns the BnExclusiveState field if non-nil, zero value otherwise.

### GetBnExclusiveStateOk

`func (o *TokenListResponseDataInner) GetBnExclusiveStateOk() (*bool, bool)`

GetBnExclusiveStateOk returns a tuple with the BnExclusiveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBnExclusiveState

`func (o *TokenListResponseDataInner) SetBnExclusiveState(v bool)`

SetBnExclusiveState sets BnExclusiveState field to given value.

### HasBnExclusiveState

`func (o *TokenListResponseDataInner) HasBnExclusiveState() bool`

HasBnExclusiveState returns a boolean if a field has been set.


[[Back to README]](../README.md)


