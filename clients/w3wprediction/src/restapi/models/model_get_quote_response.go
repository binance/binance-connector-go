/*
Prediction Trading REST API

Place and manage prediction market orders, query positions, and transfer funds via the Prediction Trading REST API.
*/

package models

import (
	"encoding/json"

	"github.com/binance/binance-connector-go/common/v2/common"
)

// checks if the GetQuoteResponse type satisfies the MappedNullable interface at compile time
var _ common.MappedNullable = &GetQuoteResponse{}

// GetQuoteResponse struct for GetQuoteResponse
type GetQuoteResponse struct {
	QuoteId              *string  `json:"quoteId,omitempty"`
	TokenId              *string  `json:"tokenId,omitempty"`
	Chance               *string  `json:"chance,omitempty"`
	Vendor               *string  `json:"vendor,omitempty"`
	MarketTitle          *string  `json:"marketTitle,omitempty"`
	MarketExtId          *string  `json:"marketExtId,omitempty"`
	Side                 *string  `json:"side,omitempty"`
	AmountIn             *string  `json:"amountIn,omitempty"`
	AmountOut            *string  `json:"amountOut,omitempty"`
	IsMinAmountOut       *bool    `json:"isMinAmountOut,omitempty"`
	FeeAmount            *string  `json:"feeAmount,omitempty"`
	FeeDiscountBps       *string  `json:"feeDiscountBps,omitempty"`
	AveragePrice         *float64 `json:"averagePrice,omitempty"`
	LastPrice            *float64 `json:"lastPrice,omitempty"`
	PriceImpact          *float64 `json:"priceImpact,omitempty"`
	Timestamp            *int64   `json:"timestamp,omitempty"`
	ChainId              *string  `json:"chainId,omitempty"`
	UserId               *int64   `json:"userId,omitempty"`
	WalletAddress        *string  `json:"walletAddress,omitempty"`
	OrderType            *string  `json:"orderType,omitempty"`
	SlippageBps          *int32   `json:"slippageBps,omitempty"`
	FeeRateBps           *int32   `json:"feeRateBps,omitempty"`
	MinReceive           *string  `json:"minReceive,omitempty"`
	ExpireAt             *int64   `json:"expireAt,omitempty"`
	PriceLimit           *string  `json:"priceLimit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetQuoteResponse GetQuoteResponse

// NewGetQuoteResponse instantiates a new GetQuoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetQuoteResponse() *GetQuoteResponse {
	this := GetQuoteResponse{}
	return &this
}

// NewGetQuoteResponseWithDefaults instantiates a new GetQuoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetQuoteResponseWithDefaults() *GetQuoteResponse {
	this := GetQuoteResponse{}
	return &this
}

// GetQuoteId returns the QuoteId field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetQuoteId() string {
	if o == nil || common.IsNil(o.QuoteId) {
		var ret string
		return ret
	}
	return *o.QuoteId
}

// GetQuoteIdOk returns a tuple with the QuoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetQuoteIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.QuoteId) {
		return nil, false
	}
	return o.QuoteId, true
}

// HasQuoteId returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasQuoteId() bool {
	if o != nil && !common.IsNil(o.QuoteId) {
		return true
	}

	return false
}

// SetQuoteId gets a reference to the given string and assigns it to the QuoteId field.
func (o *GetQuoteResponse) SetQuoteId(v string) {
	o.QuoteId = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetTokenId() string {
	if o == nil || common.IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasTokenId() bool {
	if o != nil && !common.IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *GetQuoteResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetChance returns the Chance field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetChance() string {
	if o == nil || common.IsNil(o.Chance) {
		var ret string
		return ret
	}
	return *o.Chance
}

// GetChanceOk returns a tuple with the Chance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetChanceOk() (*string, bool) {
	if o == nil || common.IsNil(o.Chance) {
		return nil, false
	}
	return o.Chance, true
}

// HasChance returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasChance() bool {
	if o != nil && !common.IsNil(o.Chance) {
		return true
	}

	return false
}

// SetChance gets a reference to the given string and assigns it to the Chance field.
func (o *GetQuoteResponse) SetChance(v string) {
	o.Chance = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetVendor() string {
	if o == nil || common.IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetVendorOk() (*string, bool) {
	if o == nil || common.IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasVendor() bool {
	if o != nil && !common.IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *GetQuoteResponse) SetVendor(v string) {
	o.Vendor = &v
}

// GetMarketTitle returns the MarketTitle field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetMarketTitle() string {
	if o == nil || common.IsNil(o.MarketTitle) {
		var ret string
		return ret
	}
	return *o.MarketTitle
}

// GetMarketTitleOk returns a tuple with the MarketTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetMarketTitleOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketTitle) {
		return nil, false
	}
	return o.MarketTitle, true
}

// HasMarketTitle returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasMarketTitle() bool {
	if o != nil && !common.IsNil(o.MarketTitle) {
		return true
	}

	return false
}

// SetMarketTitle gets a reference to the given string and assigns it to the MarketTitle field.
func (o *GetQuoteResponse) SetMarketTitle(v string) {
	o.MarketTitle = &v
}

// GetMarketExtId returns the MarketExtId field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetMarketExtId() string {
	if o == nil || common.IsNil(o.MarketExtId) {
		var ret string
		return ret
	}
	return *o.MarketExtId
}

// GetMarketExtIdOk returns a tuple with the MarketExtId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetMarketExtIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.MarketExtId) {
		return nil, false
	}
	return o.MarketExtId, true
}

// HasMarketExtId returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasMarketExtId() bool {
	if o != nil && !common.IsNil(o.MarketExtId) {
		return true
	}

	return false
}

// SetMarketExtId gets a reference to the given string and assigns it to the MarketExtId field.
func (o *GetQuoteResponse) SetMarketExtId(v string) {
	o.MarketExtId = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetSide() string {
	if o == nil || common.IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetSideOk() (*string, bool) {
	if o == nil || common.IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasSide() bool {
	if o != nil && !common.IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *GetQuoteResponse) SetSide(v string) {
	o.Side = &v
}

// GetAmountIn returns the AmountIn field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetAmountIn() string {
	if o == nil || common.IsNil(o.AmountIn) {
		var ret string
		return ret
	}
	return *o.AmountIn
}

// GetAmountInOk returns a tuple with the AmountIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetAmountInOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountIn) {
		return nil, false
	}
	return o.AmountIn, true
}

// HasAmountIn returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasAmountIn() bool {
	if o != nil && !common.IsNil(o.AmountIn) {
		return true
	}

	return false
}

// SetAmountIn gets a reference to the given string and assigns it to the AmountIn field.
func (o *GetQuoteResponse) SetAmountIn(v string) {
	o.AmountIn = &v
}

// GetAmountOut returns the AmountOut field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetAmountOut() string {
	if o == nil || common.IsNil(o.AmountOut) {
		var ret string
		return ret
	}
	return *o.AmountOut
}

// GetAmountOutOk returns a tuple with the AmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetAmountOutOk() (*string, bool) {
	if o == nil || common.IsNil(o.AmountOut) {
		return nil, false
	}
	return o.AmountOut, true
}

// HasAmountOut returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasAmountOut() bool {
	if o != nil && !common.IsNil(o.AmountOut) {
		return true
	}

	return false
}

// SetAmountOut gets a reference to the given string and assigns it to the AmountOut field.
func (o *GetQuoteResponse) SetAmountOut(v string) {
	o.AmountOut = &v
}

// GetIsMinAmountOut returns the IsMinAmountOut field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetIsMinAmountOut() bool {
	if o == nil || common.IsNil(o.IsMinAmountOut) {
		var ret bool
		return ret
	}
	return *o.IsMinAmountOut
}

// GetIsMinAmountOutOk returns a tuple with the IsMinAmountOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetIsMinAmountOutOk() (*bool, bool) {
	if o == nil || common.IsNil(o.IsMinAmountOut) {
		return nil, false
	}
	return o.IsMinAmountOut, true
}

// HasIsMinAmountOut returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasIsMinAmountOut() bool {
	if o != nil && !common.IsNil(o.IsMinAmountOut) {
		return true
	}

	return false
}

// SetIsMinAmountOut gets a reference to the given bool and assigns it to the IsMinAmountOut field.
func (o *GetQuoteResponse) SetIsMinAmountOut(v bool) {
	o.IsMinAmountOut = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetFeeAmount() string {
	if o == nil || common.IsNil(o.FeeAmount) {
		var ret string
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetFeeAmountOk() (*string, bool) {
	if o == nil || common.IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasFeeAmount() bool {
	if o != nil && !common.IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given string and assigns it to the FeeAmount field.
func (o *GetQuoteResponse) SetFeeAmount(v string) {
	o.FeeAmount = &v
}

// GetFeeDiscountBps returns the FeeDiscountBps field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetFeeDiscountBps() string {
	if o == nil || common.IsNil(o.FeeDiscountBps) {
		var ret string
		return ret
	}
	return *o.FeeDiscountBps
}

// GetFeeDiscountBpsOk returns a tuple with the FeeDiscountBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetFeeDiscountBpsOk() (*string, bool) {
	if o == nil || common.IsNil(o.FeeDiscountBps) {
		return nil, false
	}
	return o.FeeDiscountBps, true
}

// HasFeeDiscountBps returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasFeeDiscountBps() bool {
	if o != nil && !common.IsNil(o.FeeDiscountBps) {
		return true
	}

	return false
}

// SetFeeDiscountBps gets a reference to the given string and assigns it to the FeeDiscountBps field.
func (o *GetQuoteResponse) SetFeeDiscountBps(v string) {
	o.FeeDiscountBps = &v
}

// GetAveragePrice returns the AveragePrice field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetAveragePrice() float64 {
	if o == nil || common.IsNil(o.AveragePrice) {
		var ret float64
		return ret
	}
	return *o.AveragePrice
}

// GetAveragePriceOk returns a tuple with the AveragePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetAveragePriceOk() (*float64, bool) {
	if o == nil || common.IsNil(o.AveragePrice) {
		return nil, false
	}
	return o.AveragePrice, true
}

// HasAveragePrice returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasAveragePrice() bool {
	if o != nil && !common.IsNil(o.AveragePrice) {
		return true
	}

	return false
}

// SetAveragePrice gets a reference to the given float64 and assigns it to the AveragePrice field.
func (o *GetQuoteResponse) SetAveragePrice(v float64) {
	o.AveragePrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetLastPrice() float64 {
	if o == nil || common.IsNil(o.LastPrice) {
		var ret float64
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetLastPriceOk() (*float64, bool) {
	if o == nil || common.IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasLastPrice() bool {
	if o != nil && !common.IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given float64 and assigns it to the LastPrice field.
func (o *GetQuoteResponse) SetLastPrice(v float64) {
	o.LastPrice = &v
}

// GetPriceImpact returns the PriceImpact field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetPriceImpact() float64 {
	if o == nil || common.IsNil(o.PriceImpact) {
		var ret float64
		return ret
	}
	return *o.PriceImpact
}

// GetPriceImpactOk returns a tuple with the PriceImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetPriceImpactOk() (*float64, bool) {
	if o == nil || common.IsNil(o.PriceImpact) {
		return nil, false
	}
	return o.PriceImpact, true
}

// HasPriceImpact returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasPriceImpact() bool {
	if o != nil && !common.IsNil(o.PriceImpact) {
		return true
	}

	return false
}

// SetPriceImpact gets a reference to the given float64 and assigns it to the PriceImpact field.
func (o *GetQuoteResponse) SetPriceImpact(v float64) {
	o.PriceImpact = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetTimestamp() int64 {
	if o == nil || common.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetTimestampOk() (*int64, bool) {
	if o == nil || common.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasTimestamp() bool {
	if o != nil && !common.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *GetQuoteResponse) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetChainId() string {
	if o == nil || common.IsNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetChainIdOk() (*string, bool) {
	if o == nil || common.IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasChainId() bool {
	if o != nil && !common.IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *GetQuoteResponse) SetChainId(v string) {
	o.ChainId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetUserId() int64 {
	if o == nil || common.IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetUserIdOk() (*int64, bool) {
	if o == nil || common.IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasUserId() bool {
	if o != nil && !common.IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *GetQuoteResponse) SetUserId(v int64) {
	o.UserId = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetWalletAddress() string {
	if o == nil || common.IsNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetWalletAddressOk() (*string, bool) {
	if o == nil || common.IsNil(o.WalletAddress) {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasWalletAddress() bool {
	if o != nil && !common.IsNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *GetQuoteResponse) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetOrderType() string {
	if o == nil || common.IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetOrderTypeOk() (*string, bool) {
	if o == nil || common.IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasOrderType() bool {
	if o != nil && !common.IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *GetQuoteResponse) SetOrderType(v string) {
	o.OrderType = &v
}

// GetSlippageBps returns the SlippageBps field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetSlippageBps() int32 {
	if o == nil || common.IsNil(o.SlippageBps) {
		var ret int32
		return ret
	}
	return *o.SlippageBps
}

// GetSlippageBpsOk returns a tuple with the SlippageBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetSlippageBpsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.SlippageBps) {
		return nil, false
	}
	return o.SlippageBps, true
}

// HasSlippageBps returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasSlippageBps() bool {
	if o != nil && !common.IsNil(o.SlippageBps) {
		return true
	}

	return false
}

// SetSlippageBps gets a reference to the given int32 and assigns it to the SlippageBps field.
func (o *GetQuoteResponse) SetSlippageBps(v int32) {
	o.SlippageBps = &v
}

// GetFeeRateBps returns the FeeRateBps field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetFeeRateBps() int32 {
	if o == nil || common.IsNil(o.FeeRateBps) {
		var ret int32
		return ret
	}
	return *o.FeeRateBps
}

// GetFeeRateBpsOk returns a tuple with the FeeRateBps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetFeeRateBpsOk() (*int32, bool) {
	if o == nil || common.IsNil(o.FeeRateBps) {
		return nil, false
	}
	return o.FeeRateBps, true
}

// HasFeeRateBps returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasFeeRateBps() bool {
	if o != nil && !common.IsNil(o.FeeRateBps) {
		return true
	}

	return false
}

// SetFeeRateBps gets a reference to the given int32 and assigns it to the FeeRateBps field.
func (o *GetQuoteResponse) SetFeeRateBps(v int32) {
	o.FeeRateBps = &v
}

// GetMinReceive returns the MinReceive field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetMinReceive() string {
	if o == nil || common.IsNil(o.MinReceive) {
		var ret string
		return ret
	}
	return *o.MinReceive
}

// GetMinReceiveOk returns a tuple with the MinReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetMinReceiveOk() (*string, bool) {
	if o == nil || common.IsNil(o.MinReceive) {
		return nil, false
	}
	return o.MinReceive, true
}

// HasMinReceive returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasMinReceive() bool {
	if o != nil && !common.IsNil(o.MinReceive) {
		return true
	}

	return false
}

// SetMinReceive gets a reference to the given string and assigns it to the MinReceive field.
func (o *GetQuoteResponse) SetMinReceive(v string) {
	o.MinReceive = &v
}

// GetExpireAt returns the ExpireAt field value if set, zero value otherwise.
func (o *GetQuoteResponse) GetExpireAt() int64 {
	if o == nil || common.IsNil(o.ExpireAt) {
		var ret int64
		return ret
	}
	return *o.ExpireAt
}

// GetExpireAtOk returns a tuple with the ExpireAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetQuoteResponse) GetExpireAtOk() (*int64, bool) {
	if o == nil || common.IsNil(o.ExpireAt) {
		return nil, false
	}
	return o.ExpireAt, true
}

// HasExpireAt returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasExpireAt() bool {
	if o != nil && !common.IsNil(o.ExpireAt) {
		return true
	}

	return false
}

// SetExpireAt gets a reference to the given int64 and assigns it to the ExpireAt field.
func (o *GetQuoteResponse) SetExpireAt(v int64) {
	o.ExpireAt = &v
}

// GetPriceLimit returns the PriceLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetQuoteResponse) GetPriceLimit() string {
	if o == nil || common.IsNil(o.PriceLimit) {
		var ret string
		return ret
	}
	return *o.PriceLimit
}

// GetPriceLimitOk returns a tuple with the PriceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetQuoteResponse) GetPriceLimitOk() (*string, bool) {
	if o == nil || common.IsNil(o.PriceLimit) {
		return nil, false
	}
	return o.PriceLimit, true
}

// HasPriceLimit returns a boolean if a field has been set.
func (o *GetQuoteResponse) HasPriceLimit() bool {
	if o != nil && !common.IsNil(o.PriceLimit) {
		return true
	}

	return false
}

// SetPriceLimit gets a reference to the given NullableString and assigns it to the PriceLimit field.
func (o *GetQuoteResponse) SetPriceLimit(v string) {
	o.PriceLimit = &v
}

func (o GetQuoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetQuoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !common.IsNil(o.QuoteId) {
		toSerialize["quoteId"] = o.QuoteId
	}
	if !common.IsNil(o.TokenId) {
		toSerialize["tokenId"] = o.TokenId
	}
	if !common.IsNil(o.Chance) {
		toSerialize["chance"] = o.Chance
	}
	if !common.IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !common.IsNil(o.MarketTitle) {
		toSerialize["marketTitle"] = o.MarketTitle
	}
	if !common.IsNil(o.MarketExtId) {
		toSerialize["marketExtId"] = o.MarketExtId
	}
	if !common.IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !common.IsNil(o.AmountIn) {
		toSerialize["amountIn"] = o.AmountIn
	}
	if !common.IsNil(o.AmountOut) {
		toSerialize["amountOut"] = o.AmountOut
	}
	if !common.IsNil(o.IsMinAmountOut) {
		toSerialize["isMinAmountOut"] = o.IsMinAmountOut
	}
	if !common.IsNil(o.FeeAmount) {
		toSerialize["feeAmount"] = o.FeeAmount
	}
	if !common.IsNil(o.FeeDiscountBps) {
		toSerialize["feeDiscountBps"] = o.FeeDiscountBps
	}
	if !common.IsNil(o.AveragePrice) {
		toSerialize["averagePrice"] = o.AveragePrice
	}
	if !common.IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !common.IsNil(o.PriceImpact) {
		toSerialize["priceImpact"] = o.PriceImpact
	}
	if !common.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !common.IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !common.IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !common.IsNil(o.WalletAddress) {
		toSerialize["walletAddress"] = o.WalletAddress
	}
	if !common.IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !common.IsNil(o.SlippageBps) {
		toSerialize["slippageBps"] = o.SlippageBps
	}
	if !common.IsNil(o.FeeRateBps) {
		toSerialize["feeRateBps"] = o.FeeRateBps
	}
	if !common.IsNil(o.MinReceive) {
		toSerialize["minReceive"] = o.MinReceive
	}
	if !common.IsNil(o.ExpireAt) {
		toSerialize["expireAt"] = o.ExpireAt
	}
	if !common.IsNil(o.PriceLimit) {
		toSerialize["priceLimit"] = o.PriceLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetQuoteResponse) UnmarshalJSON(data []byte) (err error) {
	varGetQuoteResponse := _GetQuoteResponse{}

	err = json.Unmarshal(data, &varGetQuoteResponse)

	if err != nil {
		return err
	}

	*o = GetQuoteResponse(varGetQuoteResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "quoteId")
		delete(additionalProperties, "tokenId")
		delete(additionalProperties, "chance")
		delete(additionalProperties, "vendor")
		delete(additionalProperties, "marketTitle")
		delete(additionalProperties, "marketExtId")
		delete(additionalProperties, "side")
		delete(additionalProperties, "amountIn")
		delete(additionalProperties, "amountOut")
		delete(additionalProperties, "isMinAmountOut")
		delete(additionalProperties, "feeAmount")
		delete(additionalProperties, "feeDiscountBps")
		delete(additionalProperties, "averagePrice")
		delete(additionalProperties, "lastPrice")
		delete(additionalProperties, "priceImpact")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "walletAddress")
		delete(additionalProperties, "orderType")
		delete(additionalProperties, "slippageBps")
		delete(additionalProperties, "feeRateBps")
		delete(additionalProperties, "minReceive")
		delete(additionalProperties, "expireAt")
		delete(additionalProperties, "priceLimit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetQuoteResponse struct {
	value *GetQuoteResponse
	isSet bool
}

func (v NullableGetQuoteResponse) Get() *GetQuoteResponse {
	return v.value
}

func (v *NullableGetQuoteResponse) Set(val *GetQuoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetQuoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetQuoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetQuoteResponse(val *GetQuoteResponse) *NullableGetQuoteResponse {
	return &NullableGetQuoteResponse{value: val, isSet: true}
}

func (v NullableGetQuoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetQuoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
