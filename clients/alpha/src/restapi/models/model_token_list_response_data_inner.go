/*
Binance Alpha REST API

OpenAPI Specification for the Binance Alpha REST API
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/common"
)

// checks if the TokenListResponseDataInner type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &TokenListResponseDataInner{}

// TokenListResponseDataInner struct for TokenListResponseDataInner
type TokenListResponseDataInner struct {
	TokenId              *string `json:"tokenId,omitempty"`
	ChainId              *string `json:"chainId,omitempty"`
	ChainIconUrl         *string `json:"chainIconUrl,omitempty"`
	ChainName            *string `json:"chainName,omitempty"`
	ContractAddress      *string `json:"contractAddress,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Symbol               *string `json:"symbol,omitempty"`
	IconUrl              *string `json:"iconUrl,omitempty"`
	Price                *string `json:"price,omitempty"`
	PercentChange24h     *string `json:"percentChange24h,omitempty"`
	Volume24h            *string `json:"volume24h,omitempty"`
	MarketCap            *string `json:"marketCap,omitempty"`
	Fdv                  *string `json:"fdv,omitempty"`
	Liquidity            *string `json:"liquidity,omitempty"`
	TotalSupply          *string `json:"totalSupply,omitempty"`
	CirculatingSupply    *string `json:"circulatingSupply,omitempty"`
	Holders              *string `json:"holders,omitempty"`
	Decimals             *int64  `json:"decimals,omitempty"`
	ListingCex           *bool   `json:"listingCex,omitempty"`
	HotTag               *bool   `json:"hotTag,omitempty"`
	CexCoinName          *string `json:"cexCoinName,omitempty"`
	CanTransfer          *bool   `json:"canTransfer,omitempty"`
	Denomination         *int64  `json:"denomination,omitempty"`
	Offline              *bool   `json:"offline,omitempty"`
	TradeDecimal         *int64  `json:"tradeDecimal,omitempty"`
	AlphaId              *string `json:"alphaId,omitempty"`
	Offsell              *bool   `json:"offsell,omitempty"`
	PriceHigh24h         *string `json:"priceHigh24h,omitempty"`
	PriceLow24h          *string `json:"priceLow24h,omitempty"`
	Count24h             *string `json:"count24h,omitempty"`
	OnlineTge            *bool   `json:"onlineTge,omitempty"`
	OnlineAirdrop        *bool   `json:"onlineAirdrop,omitempty"`
	Score                *int64  `json:"score,omitempty"`
	CexOffDisplay        *bool   `json:"cexOffDisplay,omitempty"`
	StockState           *bool   `json:"stockState,omitempty"`
	ListingTime          *int64  `json:"listingTime,omitempty"`
	MulPoint             *int64  `json:"mulPoint,omitempty"`
	BnExclusiveState     *bool   `json:"bnExclusiveState,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenListResponseDataInner TokenListResponseDataInner

// NewTokenListResponseDataInner instantiates a new TokenListResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenListResponseDataInner() *TokenListResponseDataInner {
	this := TokenListResponseDataInner{}
	return &this
}

// NewTokenListResponseDataInnerWithDefaults instantiates a new TokenListResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenListResponseDataInnerWithDefaults() *TokenListResponseDataInner {
	this := TokenListResponseDataInner{}
	return &this
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TokenListResponseDataInner) SetTokenId(v string) {
	o.TokenId = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *TokenListResponseDataInner) SetChainId(v string) {
	o.ChainId = &v
}

// GetChainIconUrl returns the ChainIconUrl field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetChainIconUrl() string {
	if o == nil || common.IsNil(o.ChainIconUrl) {
		var ret string
		return ret
	}
	return *o.ChainIconUrl
}

// GetChainIconUrlOk returns a tuple with the ChainIconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetChainIconUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainIconUrl) {
		return nil, false
	}
	return o.ChainIconUrl, true
}

// HasChainIconUrl returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasChainIconUrl() bool {
	if o != nil && !common.IsNil(o.ChainIconUrl) {
		return true
	}

	return false
}

// SetChainIconUrl gets a reference to the given string and assigns it to the ChainIconUrl field.
func (o *TokenListResponseDataInner) SetChainIconUrl(v string) {
	o.ChainIconUrl = &v
}

// GetChainName returns the ChainName field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetChainName() string {
	if o == nil || common.IsNil(o.ChainName) {
		var ret string
		return ret
	}
	return *o.ChainName
}

// GetChainNameOk returns a tuple with the ChainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetChainNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainName) {
		return nil, false
	}
	return o.ChainName, true
}

// HasChainName returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasChainName() bool {
	if o != nil && !common.IsNil(o.ChainName) {
		return true
	}

	return false
}

// SetChainName gets a reference to the given string and assigns it to the ChainName field.
func (o *TokenListResponseDataInner) SetChainName(v string) {
	o.ChainName = &v
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetContractAddress() string {
	if o == nil || common.IsNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetContractAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasContractAddress() bool {
	if o != nil && !common.IsNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *TokenListResponseDataInner) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetName() string {
	if o == nil || common.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasName() bool {
	if o != nil && !common.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TokenListResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetSymbol() string {
	if o == nil || common.IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetSymbolOk() (*string, bool) {
	if o == nil || common.IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasSymbol() bool {
	if o != nil && !common.IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TokenListResponseDataInner) SetSymbol(v string) {
	o.Symbol = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetIconUrl() string {
	if o == nil || common.IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetIconUrlOk() (*string, bool) {
	if o == nil || common.IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasIconUrl() bool {
	if o != nil && !common.IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *TokenListResponseDataInner) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetPrice() string {
	if o == nil || common.IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetPriceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasPrice() bool {
	if o != nil && !common.IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *TokenListResponseDataInner) SetPrice(v string) {
	o.Price = &v
}

// GetPercentChange24h returns the PercentChange24h field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetPercentChange24h() string {
	if o == nil || common.IsNil(o.PercentChange24h) {
		var ret string
		return ret
	}
	return *o.PercentChange24h
}

// GetPercentChange24hOk returns a tuple with the PercentChange24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetPercentChange24hOk() (*string, bool) {
	if o == nil || common.IsNil(o.PercentChange24h) {
		return nil, false
	}
	return o.PercentChange24h, true
}

// HasPercentChange24h returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasPercentChange24h() bool {
	if o != nil && !common.IsNil(o.PercentChange24h) {
		return true
	}

	return false
}

// SetPercentChange24h gets a reference to the given string and assigns it to the PercentChange24h field.
func (o *TokenListResponseDataInner) SetPercentChange24h(v string) {
	o.PercentChange24h = &v
}

// GetVolume24h returns the Volume24h field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetVolume24h() string {
	if o == nil || common.IsNil(o.Volume24h) {
		var ret string
		return ret
	}
	return *o.Volume24h
}

// GetVolume24hOk returns a tuple with the Volume24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetVolume24hOk() (*string, bool) {
	if o == nil || common.IsNil(o.Volume24h) {
		return nil, false
	}
	return o.Volume24h, true
}

// HasVolume24h returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasVolume24h() bool {
	if o != nil && !common.IsNil(o.Volume24h) {
		return true
	}

	return false
}

// SetVolume24h gets a reference to the given string and assigns it to the Volume24h field.
func (o *TokenListResponseDataInner) SetVolume24h(v string) {
	o.Volume24h = &v
}

// GetMarketCap returns the MarketCap field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetMarketCap() string {
	if o == nil || common.IsNil(o.MarketCap) {
		var ret string
		return ret
	}
	return *o.MarketCap
}

// GetMarketCapOk returns a tuple with the MarketCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetMarketCapOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketCap) {
		return nil, false
	}
	return o.MarketCap, true
}

// HasMarketCap returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasMarketCap() bool {
	if o != nil && !common.IsNil(o.MarketCap) {
		return true
	}

	return false
}

// SetMarketCap gets a reference to the given string and assigns it to the MarketCap field.
func (o *TokenListResponseDataInner) SetMarketCap(v string) {
	o.MarketCap = &v
}

// GetFdv returns the Fdv field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetFdv() string {
	if o == nil || common.IsNil(o.Fdv) {
		var ret string
		return ret
	}
	return *o.Fdv
}

// GetFdvOk returns a tuple with the Fdv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetFdvOk() (*string, bool) {
	if o == nil || common.IsNil(o.Fdv) {
		return nil, false
	}
	return o.Fdv, true
}

// HasFdv returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasFdv() bool {
	if o != nil && !common.IsNil(o.Fdv) {
		return true
	}

	return false
}

// SetFdv gets a reference to the given string and assigns it to the Fdv field.
func (o *TokenListResponseDataInner) SetFdv(v string) {
	o.Fdv = &v
}

// GetLiquidity returns the Liquidity field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetLiquidity() string {
	if o == nil || common.IsNil(o.Liquidity) {
		var ret string
		return ret
	}
	return *o.Liquidity
}

// GetLiquidityOk returns a tuple with the Liquidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetLiquidityOk() (*string, bool) {
	if o == nil || common.IsNil(o.Liquidity) {
		return nil, false
	}
	return o.Liquidity, true
}

// HasLiquidity returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasLiquidity() bool {
	if o != nil && !common.IsNil(o.Liquidity) {
		return true
	}

	return false
}

// SetLiquidity gets a reference to the given string and assigns it to the Liquidity field.
func (o *TokenListResponseDataInner) SetLiquidity(v string) {
	o.Liquidity = &v
}

// GetTotalSupply returns the TotalSupply field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetTotalSupply() string {
	if o == nil || common.IsNil(o.TotalSupply) {
		var ret string
		return ret
	}
	return *o.TotalSupply
}

// GetTotalSupplyOk returns a tuple with the TotalSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetTotalSupplyOk() (*string, bool) {
	if o == nil || common.IsNil(o.TotalSupply) {
		return nil, false
	}
	return o.TotalSupply, true
}

// HasTotalSupply returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasTotalSupply() bool {
	if o != nil && !common.IsNil(o.TotalSupply) {
		return true
	}

	return false
}

// SetTotalSupply gets a reference to the given string and assigns it to the TotalSupply field.
func (o *TokenListResponseDataInner) SetTotalSupply(v string) {
	o.TotalSupply = &v
}

// GetCirculatingSupply returns the CirculatingSupply field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetCirculatingSupply() string {
	if o == nil || common.IsNil(o.CirculatingSupply) {
		var ret string
		return ret
	}
	return *o.CirculatingSupply
}

// GetCirculatingSupplyOk returns a tuple with the CirculatingSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetCirculatingSupplyOk() (*string, bool) {
	if o == nil || common.IsNil(o.CirculatingSupply) {
		return nil, false
	}
	return o.CirculatingSupply, true
}

// HasCirculatingSupply returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasCirculatingSupply() bool {
	if o != nil && !common.IsNil(o.CirculatingSupply) {
		return true
	}

	return false
}

// SetCirculatingSupply gets a reference to the given string and assigns it to the CirculatingSupply field.
func (o *TokenListResponseDataInner) SetCirculatingSupply(v string) {
	o.CirculatingSupply = &v
}

// GetHolders returns the Holders field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetHolders() string {
	if o == nil || common.IsNil(o.Holders) {
		var ret string
		return ret
	}
	return *o.Holders
}

// GetHoldersOk returns a tuple with the Holders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetHoldersOk() (*string, bool) {
	if o == nil || common.IsNil(o.Holders) {
		return nil, false
	}
	return o.Holders, true
}

// HasHolders returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasHolders() bool {
	if o != nil && !common.IsNil(o.Holders) {
		return true
	}

	return false
}

// SetHolders gets a reference to the given string and assigns it to the Holders field.
func (o *TokenListResponseDataInner) SetHolders(v string) {
	o.Holders = &v
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetDecimals() int64 {
	if o == nil || common.IsNil(o.Decimals) {
		var ret int64
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetDecimalsOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Decimals) {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasDecimals() bool {
	if o != nil && !common.IsNil(o.Decimals) {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int64 and assigns it to the Decimals field.
func (o *TokenListResponseDataInner) SetDecimals(v int64) {
	o.Decimals = &v
}

// GetListingCex returns the ListingCex field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetListingCex() bool {
	if o == nil || common.IsNil(o.ListingCex) {
		var ret bool
		return ret
	}
	return *o.ListingCex
}

// GetListingCexOk returns a tuple with the ListingCex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetListingCexOk() (*bool, bool) {
	if o == nil || common.IsNil(o.ListingCex) {
		return nil, false
	}
	return o.ListingCex, true
}

// HasListingCex returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasListingCex() bool {
	if o != nil && !common.IsNil(o.ListingCex) {
		return true
	}

	return false
}

// SetListingCex gets a reference to the given bool and assigns it to the ListingCex field.
func (o *TokenListResponseDataInner) SetListingCex(v bool) {
	o.ListingCex = &v
}

// GetHotTag returns the HotTag field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetHotTag() bool {
	if o == nil || common.IsNil(o.HotTag) {
		var ret bool
		return ret
	}
	return *o.HotTag
}

// GetHotTagOk returns a tuple with the HotTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetHotTagOk() (*bool, bool) {
	if o == nil || common.IsNil(o.HotTag) {
		return nil, false
	}
	return o.HotTag, true
}

// HasHotTag returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasHotTag() bool {
	if o != nil && !common.IsNil(o.HotTag) {
		return true
	}

	return false
}

// SetHotTag gets a reference to the given bool and assigns it to the HotTag field.
func (o *TokenListResponseDataInner) SetHotTag(v bool) {
	o.HotTag = &v
}

// GetCexCoinName returns the CexCoinName field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetCexCoinName() string {
	if o == nil || common.IsNil(o.CexCoinName) {
		var ret string
		return ret
	}
	return *o.CexCoinName
}

// GetCexCoinNameOk returns a tuple with the CexCoinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetCexCoinNameOk() (*string, bool) {
	if o == nil || common.IsNil(o.CexCoinName) {
		return nil, false
	}
	return o.CexCoinName, true
}

// HasCexCoinName returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasCexCoinName() bool {
	if o != nil && !common.IsNil(o.CexCoinName) {
		return true
	}

	return false
}

// SetCexCoinName gets a reference to the given string and assigns it to the CexCoinName field.
func (o *TokenListResponseDataInner) SetCexCoinName(v string) {
	o.CexCoinName = &v
}

// GetCanTransfer returns the CanTransfer field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetCanTransfer() bool {
	if o == nil || common.IsNil(o.CanTransfer) {
		var ret bool
		return ret
	}
	return *o.CanTransfer
}

// GetCanTransferOk returns a tuple with the CanTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetCanTransferOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CanTransfer) {
		return nil, false
	}
	return o.CanTransfer, true
}

// HasCanTransfer returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasCanTransfer() bool {
	if o != nil && !common.IsNil(o.CanTransfer) {
		return true
	}

	return false
}

// SetCanTransfer gets a reference to the given bool and assigns it to the CanTransfer field.
func (o *TokenListResponseDataInner) SetCanTransfer(v bool) {
	o.CanTransfer = &v
}

// GetDenomination returns the Denomination field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetDenomination() int64 {
	if o == nil || common.IsNil(o.Denomination) {
		var ret int64
		return ret
	}
	return *o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetDenominationOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Denomination) {
		return nil, false
	}
	return o.Denomination, true
}

// HasDenomination returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasDenomination() bool {
	if o != nil && !common.IsNil(o.Denomination) {
		return true
	}

	return false
}

// SetDenomination gets a reference to the given int64 and assigns it to the Denomination field.
func (o *TokenListResponseDataInner) SetDenomination(v int64) {
	o.Denomination = &v
}

// GetOffline returns the Offline field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetOffline() bool {
	if o == nil || common.IsNil(o.Offline) {
		var ret bool
		return ret
	}
	return *o.Offline
}

// GetOfflineOk returns a tuple with the Offline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetOfflineOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Offline) {
		return nil, false
	}
	return o.Offline, true
}

// HasOffline returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasOffline() bool {
	if o != nil && !common.IsNil(o.Offline) {
		return true
	}

	return false
}

// SetOffline gets a reference to the given bool and assigns it to the Offline field.
func (o *TokenListResponseDataInner) SetOffline(v bool) {
	o.Offline = &v
}

// GetTradeDecimal returns the TradeDecimal field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetTradeDecimal() int64 {
	if o == nil || common.IsNil(o.TradeDecimal) {
		var ret int64
		return ret
	}
	return *o.TradeDecimal
}

// GetTradeDecimalOk returns a tuple with the TradeDecimal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetTradeDecimalOk() (*int64, bool) {
	if o == nil || common.IsNil(o.TradeDecimal) {
		return nil, false
	}
	return o.TradeDecimal, true
}

// HasTradeDecimal returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasTradeDecimal() bool {
	if o != nil && !common.IsNil(o.TradeDecimal) {
		return true
	}

	return false
}

// SetTradeDecimal gets a reference to the given int64 and assigns it to the TradeDecimal field.
func (o *TokenListResponseDataInner) SetTradeDecimal(v int64) {
	o.TradeDecimal = &v
}

// GetAlphaId returns the AlphaId field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetAlphaId() string {
	if o == nil || common.IsNil(o.AlphaId) {
		var ret string
		return ret
	}
	return *o.AlphaId
}

// GetAlphaIdOk returns a tuple with the AlphaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetAlphaIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.AlphaId) {
		return nil, false
	}
	return o.AlphaId, true
}

// HasAlphaId returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasAlphaId() bool {
	if o != nil && !common.IsNil(o.AlphaId) {
		return true
	}

	return false
}

// SetAlphaId gets a reference to the given string and assigns it to the AlphaId field.
func (o *TokenListResponseDataInner) SetAlphaId(v string) {
	o.AlphaId = &v
}

// GetOffsell returns the Offsell field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetOffsell() bool {
	if o == nil || common.IsNil(o.Offsell) {
		var ret bool
		return ret
	}
	return *o.Offsell
}

// GetOffsellOk returns a tuple with the Offsell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetOffsellOk() (*bool, bool) {
	if o == nil || common.IsNil(o.Offsell) {
		return nil, false
	}
	return o.Offsell, true
}

// HasOffsell returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasOffsell() bool {
	if o != nil && !common.IsNil(o.Offsell) {
		return true
	}

	return false
}

// SetOffsell gets a reference to the given bool and assigns it to the Offsell field.
func (o *TokenListResponseDataInner) SetOffsell(v bool) {
	o.Offsell = &v
}

// GetPriceHigh24h returns the PriceHigh24h field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetPriceHigh24h() string {
	if o == nil || common.IsNil(o.PriceHigh24h) {
		var ret string
		return ret
	}
	return *o.PriceHigh24h
}

// GetPriceHigh24hOk returns a tuple with the PriceHigh24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetPriceHigh24hOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceHigh24h) {
		return nil, false
	}
	return o.PriceHigh24h, true
}

// HasPriceHigh24h returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasPriceHigh24h() bool {
	if o != nil && !common.IsNil(o.PriceHigh24h) {
		return true
	}

	return false
}

// SetPriceHigh24h gets a reference to the given string and assigns it to the PriceHigh24h field.
func (o *TokenListResponseDataInner) SetPriceHigh24h(v string) {
	o.PriceHigh24h = &v
}

// GetPriceLow24h returns the PriceLow24h field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetPriceLow24h() string {
	if o == nil || common.IsNil(o.PriceLow24h) {
		var ret string
		return ret
	}
	return *o.PriceLow24h
}

// GetPriceLow24hOk returns a tuple with the PriceLow24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetPriceLow24hOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceLow24h) {
		return nil, false
	}
	return o.PriceLow24h, true
}

// HasPriceLow24h returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasPriceLow24h() bool {
	if o != nil && !common.IsNil(o.PriceLow24h) {
		return true
	}

	return false
}

// SetPriceLow24h gets a reference to the given string and assigns it to the PriceLow24h field.
func (o *TokenListResponseDataInner) SetPriceLow24h(v string) {
	o.PriceLow24h = &v
}

// GetCount24h returns the Count24h field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetCount24h() string {
	if o == nil || common.IsNil(o.Count24h) {
		var ret string
		return ret
	}
	return *o.Count24h
}

// GetCount24hOk returns a tuple with the Count24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetCount24hOk() (*string, bool) {
	if o == nil || common.IsNil(o.Count24h) {
		return nil, false
	}
	return o.Count24h, true
}

// HasCount24h returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasCount24h() bool {
	if o != nil && !common.IsNil(o.Count24h) {
		return true
	}

	return false
}

// SetCount24h gets a reference to the given string and assigns it to the Count24h field.
func (o *TokenListResponseDataInner) SetCount24h(v string) {
	o.Count24h = &v
}

// GetOnlineTge returns the OnlineTge field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetOnlineTge() bool {
	if o == nil || common.IsNil(o.OnlineTge) {
		var ret bool
		return ret
	}
	return *o.OnlineTge
}

// GetOnlineTgeOk returns a tuple with the OnlineTge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetOnlineTgeOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OnlineTge) {
		return nil, false
	}
	return o.OnlineTge, true
}

// HasOnlineTge returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasOnlineTge() bool {
	if o != nil && !common.IsNil(o.OnlineTge) {
		return true
	}

	return false
}

// SetOnlineTge gets a reference to the given bool and assigns it to the OnlineTge field.
func (o *TokenListResponseDataInner) SetOnlineTge(v bool) {
	o.OnlineTge = &v
}

// GetOnlineAirdrop returns the OnlineAirdrop field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetOnlineAirdrop() bool {
	if o == nil || common.IsNil(o.OnlineAirdrop) {
		var ret bool
		return ret
	}
	return *o.OnlineAirdrop
}

// GetOnlineAirdropOk returns a tuple with the OnlineAirdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetOnlineAirdropOk() (*bool, bool) {
	if o == nil || common.IsNil(o.OnlineAirdrop) {
		return nil, false
	}
	return o.OnlineAirdrop, true
}

// HasOnlineAirdrop returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasOnlineAirdrop() bool {
	if o != nil && !common.IsNil(o.OnlineAirdrop) {
		return true
	}

	return false
}

// SetOnlineAirdrop gets a reference to the given bool and assigns it to the OnlineAirdrop field.
func (o *TokenListResponseDataInner) SetOnlineAirdrop(v bool) {
	o.OnlineAirdrop = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetScore() int64 {
	if o == nil || common.IsNil(o.Score) {
		var ret int64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetScoreOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasScore() bool {
	if o != nil && !common.IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int64 and assigns it to the Score field.
func (o *TokenListResponseDataInner) SetScore(v int64) {
	o.Score = &v
}

// GetCexOffDisplay returns the CexOffDisplay field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetCexOffDisplay() bool {
	if o == nil || common.IsNil(o.CexOffDisplay) {
		var ret bool
		return ret
	}
	return *o.CexOffDisplay
}

// GetCexOffDisplayOk returns a tuple with the CexOffDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetCexOffDisplayOk() (*bool, bool) {
	if o == nil || common.IsNil(o.CexOffDisplay) {
		return nil, false
	}
	return o.CexOffDisplay, true
}

// HasCexOffDisplay returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasCexOffDisplay() bool {
	if o != nil && !common.IsNil(o.CexOffDisplay) {
		return true
	}

	return false
}

// SetCexOffDisplay gets a reference to the given bool and assigns it to the CexOffDisplay field.
func (o *TokenListResponseDataInner) SetCexOffDisplay(v bool) {
	o.CexOffDisplay = &v
}

// GetStockState returns the StockState field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetStockState() bool {
	if o == nil || common.IsNil(o.StockState) {
		var ret bool
		return ret
	}
	return *o.StockState
}

// GetStockStateOk returns a tuple with the StockState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetStockStateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.StockState) {
		return nil, false
	}
	return o.StockState, true
}

// HasStockState returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasStockState() bool {
	if o != nil && !common.IsNil(o.StockState) {
		return true
	}

	return false
}

// SetStockState gets a reference to the given bool and assigns it to the StockState field.
func (o *TokenListResponseDataInner) SetStockState(v bool) {
	o.StockState = &v
}

// GetListingTime returns the ListingTime field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetListingTime() int64 {
	if o == nil || common.IsNil(o.ListingTime) {
		var ret int64
		return ret
	}
	return *o.ListingTime
}

// GetListingTimeOk returns a tuple with the ListingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetListingTimeOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ListingTime) {
		return nil, false
	}
	return o.ListingTime, true
}

// HasListingTime returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasListingTime() bool {
	if o != nil && !common.IsNil(o.ListingTime) {
		return true
	}

	return false
}

// SetListingTime gets a reference to the given int64 and assigns it to the ListingTime field.
func (o *TokenListResponseDataInner) SetListingTime(v int64) {
	o.ListingTime = &v
}

// GetMulPoint returns the MulPoint field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetMulPoint() int64 {
	if o == nil || common.IsNil(o.MulPoint) {
		var ret int64
		return ret
	}
	return *o.MulPoint
}

// GetMulPointOk returns a tuple with the MulPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetMulPointOk() (*int64, bool) {
	if o == nil || common.IsNil(o.MulPoint) {
		return nil, false
	}
	return o.MulPoint, true
}

// HasMulPoint returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasMulPoint() bool {
	if o != nil && !common.IsNil(o.MulPoint) {
		return true
	}

	return false
}

// SetMulPoint gets a reference to the given int64 and assigns it to the MulPoint field.
func (o *TokenListResponseDataInner) SetMulPoint(v int64) {
	o.MulPoint = &v
}

// GetBnExclusiveState returns the BnExclusiveState field value if set, zero value otherwise.
func (o *TokenListResponseDataInner) GetBnExclusiveState() bool {
	if o == nil || common.IsNil(o.BnExclusiveState) {
		var ret bool
		return ret
	}
	return *o.BnExclusiveState
}

// GetBnExclusiveStateOk returns a tuple with the BnExclusiveState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenListResponseDataInner) GetBnExclusiveStateOk() (*bool, bool) {
	if o == nil || common.IsNil(o.BnExclusiveState) {
		return nil, false
	}
	return o.BnExclusiveState, true
}

// HasBnExclusiveState returns a boolean if a field has been set.
func (o *TokenListResponseDataInner) HasBnExclusiveState() bool {
	if o != nil && !common.IsNil(o.BnExclusiveState) {
		return true
	}

	return false
}

// SetBnExclusiveState gets a reference to the given bool and assigns it to the BnExclusiveState field.
func (o *TokenListResponseDataInner) SetBnExclusiveState(v bool) {
	o.BnExclusiveState = &v
}

func (o TokenListResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenListResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.ChainIconUrl) {
		toSerialize["chainIconUrl"] = o.ChainIconUrl
	}
	if !common.IsNil(o.ChainName) {
		toSerialize["chainName"] = o.ChainName
	}
	if !common.IsNil(o.ContractAddress) {
		toSerialize["contractAddress"] = o.ContractAddress
	}
	if !common.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !common.IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !common.IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !common.IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !common.IsNil(o.PercentChange24h) {
		toSerialize["percentChange24h"] = o.PercentChange24h
	}
	if !common.IsNil(o.Volume24h) {
		toSerialize["volume24h"] = o.Volume24h
	}
	if !common.IsNil(o.MarketCap) {
		toSerialize["marketCap"] = o.MarketCap
	}
	if !common.IsNil(o.Fdv) {
		toSerialize["fdv"] = o.Fdv
	}
	if !common.IsNil(o.Liquidity) {
		toSerialize["liquidity"] = o.Liquidity
	}
	if !common.IsNil(o.TotalSupply) {
		toSerialize["totalSupply"] = o.TotalSupply
	}
	if !common.IsNil(o.CirculatingSupply) {
		toSerialize["circulatingSupply"] = o.CirculatingSupply
	}
	if !common.IsNil(o.Holders) {
		toSerialize["holders"] = o.Holders
	}
	if !common.IsNil(o.Decimals) {
		toSerialize["decimals"] = o.Decimals
	}
	if !common.IsNil(o.ListingCex) {
		toSerialize["listingCex"] = o.ListingCex
	}
	if !common.IsNil(o.HotTag) {
		toSerialize["hotTag"] = o.HotTag
	}
	if !common.IsNil(o.CexCoinName) {
		toSerialize["cexCoinName"] = o.CexCoinName
	}
	if !common.IsNil(o.CanTransfer) {
		toSerialize["canTransfer"] = o.CanTransfer
	}
	if !common.IsNil(o.Denomination) {
		toSerialize["denomination"] = o.Denomination
	}
	if !common.IsNil(o.Offline) {
		toSerialize["offline"] = o.Offline
	}
	if !common.IsNil(o.TradeDecimal) {
		toSerialize["tradeDecimal"] = o.TradeDecimal
	}
	if !common.IsNil(o.AlphaId) {
		toSerialize["alphaId"] = o.AlphaId
	}
	if !common.IsNil(o.Offsell) {
		toSerialize["offsell"] = o.Offsell
	}
	if !common.IsNil(o.PriceHigh24h) {
		toSerialize["priceHigh24h"] = o.PriceHigh24h
	}
	if !common.IsNil(o.PriceLow24h) {
		toSerialize["priceLow24h"] = o.PriceLow24h
	}
	if !common.IsNil(o.Count24h) {
		toSerialize["count24h"] = o.Count24h
	}
	if !common.IsNil(o.OnlineTge) {
		toSerialize["onlineTge"] = o.OnlineTge
	}
	if !common.IsNil(o.OnlineAirdrop) {
		toSerialize["onlineAirdrop"] = o.OnlineAirdrop
	}
	if !common.IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !common.IsNil(o.CexOffDisplay) {
		toSerialize["cexOffDisplay"] = o.CexOffDisplay
	}
	if !common.IsNil(o.StockState) {
		toSerialize["stockState"] = o.StockState
	}
	if !common.IsNil(o.ListingTime) {
		toSerialize["listingTime"] = o.ListingTime
	}
	if !common.IsNil(o.MulPoint) {
		toSerialize["mulPoint"] = o.MulPoint
	}
	if !common.IsNil(o.BnExclusiveState) {
		toSerialize["bnExclusiveState"] = o.BnExclusiveState
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TokenListResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	varTokenListResponseDataInner := _TokenListResponseDataInner{}

	err = json.Unmarshal(data, &varTokenListResponseDataInner)

	if err != nil {
		return err
	}

	*o = TokenListResponseDataInner(varTokenListResponseDataInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "chainIconUrl")
		delete(additionalProperties, "chainName")
		delete(additionalProperties, "contractAddress")
		delete(additionalProperties, "name")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "iconUrl")
		delete(additionalProperties, "price")
		delete(additionalProperties, "percentChange24h")
		delete(additionalProperties, "volume24h")
		delete(additionalProperties, "marketCap")
		delete(additionalProperties, "fdv")
		delete(additionalProperties, "liquidity")
		delete(additionalProperties, "totalSupply")
		delete(additionalProperties, "circulatingSupply")
		delete(additionalProperties, "holders")
		delete(additionalProperties, "decimals")
		delete(additionalProperties, "listingCex")
		delete(additionalProperties, "hotTag")
		delete(additionalProperties, "cexCoinName")
		delete(additionalProperties, "canTransfer")
		delete(additionalProperties, "denomination")
		delete(additionalProperties, "offline")
		delete(additionalProperties, "tradeDecimal")
		delete(additionalProperties, "alphaId")
		delete(additionalProperties, "offsell")
		delete(additionalProperties, "priceHigh24h")
		delete(additionalProperties, "priceLow24h")
		delete(additionalProperties, "count24h")
		delete(additionalProperties, "onlineTge")
		delete(additionalProperties, "onlineAirdrop")
		delete(additionalProperties, "score")
		delete(additionalProperties, "cexOffDisplay")
		delete(additionalProperties, "stockState")
		delete(additionalProperties, "listingTime")
		delete(additionalProperties, "mulPoint")
		delete(additionalProperties, "bnExclusiveState")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTokenListResponseDataInner struct {
	value *TokenListResponseDataInner
	isSet bool
}

func (v NullableTokenListResponseDataInner) Get() *TokenListResponseDataInner {
	return v.value
}

func (v *NullableTokenListResponseDataInner) Set(val *TokenListResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenListResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenListResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenListResponseDataInner(val *TokenListResponseDataInner) *NullableTokenListResponseDataInner {
	return &NullableTokenListResponseDataInner{value: val, isSet: true}
}

func (v NullableTokenListResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenListResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
