package common

type TimeUnit string
type RateLimitType string
type Interval string
type WebsocketMode string
type WebsocketStatus string

const (
	MILLISECOND TimeUnit = "MILLISECOND"
	Millisecond TimeUnit = "millisecond"
	MICROSECOND TimeUnit = "MICROSECOND"
	Microsecond TimeUnit = "microsecond"
)

const (
	RequestWeight RateLimitType = "REQUEST_WEIGHT"
	Orders        RateLimitType = "ORDERS"
)

const (
	Second Interval = "SECOND"
	Minute Interval = "MINUTE"
	Day    Interval = "DAY"
)

const (
	CONNECTING WebsocketStatus = "CONNECTING"
	OPEN       WebsocketStatus = "OPEN"
	CLOSING    WebsocketStatus = "CLOSING"
	CLOSED     WebsocketStatus = "CLOSED"
)

const (
	SINGLE WebsocketMode = "single"
	POOL   WebsocketMode = "pool"
)

// Algo API URLs
const AlgoRestApiProdUrl = "https://api.binance.com"

// Alpha API URLs
const AlphaRestApiProdUrl = "https://www.binance.com"

// C2C API URLs
const C2CRestApiProdUrl = "https://api.binance.com"

// Convert API URLs
const ConvertRestApiProdUrl = "https://api.binance.com"

// Copy Trading API URLs
const CopyTradingRestApiProdUrl = "https://api.binance.com"

// Crypto Loan API URLs
const CryptoLoanRestApiProdUrl = "https://api.binance.com"

// Derivatives Trading Coin Futures API URLs
const DerivativesTradingCoinFuturesRestApiProdUrl = "https://dapi.binance.com"
const DerivativesTradingCoinFuturesRestApiTestnetUrl = "https://testnet.binancefuture.com"
const DerivativesTradingCoinFuturesWebsocketApiProdUrl = "wss://ws-dapi.binance.com/ws-dapi/v1"
const DerivativesTradingCoinFuturesWebsocketApiTestnetUrl = "wss://testnet.binancefuture.com/ws-dapi/v1"
const DerivativesTradingCoinFuturesWebsocketStreamsProdUrl = "wss://dstream.binance.com/stream"
const DerivativesTradingCoinFuturesWebsocketStreamsTestnetUrl = "wss://dstream.binancefuture.com/stream"

// Derivatives Trading USDS Futures API URLs
const DerivativesTradingUsdsFuturesRestApiProdUrl = "https://fapi.binance.com"
const DerivativesTradingUsdsFuturesRestApiTestnetUrl = "https://testnet.binancefuture.com"
const DerivativesTradingUsdsFuturesWebsocketApiProdUrl = "wss://ws-fapi.binance.com/ws-fapi/v1"
const DerivativesTradingUsdsFuturesWebsocketApiTestnetUrl = "wss://testnet.binancefuture.com/ws-fapi/v1"
const DerivativesTradingUsdsFuturesWebsocketStreamsProdUrl = "wss://fstream.binance.com/stream"
const DerivativesTradingUsdsFuturesWebsocketStreamsTestnetUrl = "wss://stream.binancefuture.com/stream"

// Derivatives Trading Options API URLs
const DerivativesTradingOptionsRestApiProdUrl = "https://eapi.binance.com"
const DerivativesTradingOptionsWebsocketStreamsProdUrl = "wss://fstream.binance.com"

// Derivatives Trading Portfolio Margin API URLs
const DerivativesTradingPortfolioMarginRestApiProdUrl = "https://papi.binance.com"
const DerivativesTradingPortfolioMarginRestApiTestnetUrl = "https://testnet.binancefuture.com"
const DerivativesTradingPortfolioMarginWebsocketStreamsProdUrl = "wss://fstream.binance.com/pm"
const DerivativesTradingPortfolioMarginWebsocketStreamsTestnetUrl = "wss://fstream.binancefuture.com/pm"

// Derivatives Trading Portfolio Margin Pro API URLs
const DerivativesTradingPortfolioMarginProRestApiProdUrl = "https://api.binance.com"
const DerivativesTradingPortfolioMarginProWebsocketStreamsProdUrl = "wss://fstream.binance.com/pm-classic"

// Dual Investment API URLs
const DualInvestmentRestApiProdUrl = "https://api.binance.com"

// Fiat API URLs
const FiatRestApiProdUrl = "https://api.binance.com"

// Gift Card API URLs
const GiftCardRestApiProdUrl = "https://api.binance.com"

// Margin Trading API URLs
const MarginTradingRestApiProdUrl = "https://api.binance.com"
const MarginTradingWebsocketStreamsProdUrl = "wss://stream.binance.com:9443"
const MarginTradingRiskWebsocketStreamsProdUrl = "wss://margin-stream.binance.com"

// Mining API URLs
const MiningRestApiProdUrl = "https://api.binance.com"

// NFT API URLs
const NFTRestApiProdUrl = "https://api.binance.com"

// Pay API URLs
const PayRestApiProdUrl = "https://api.binance.com"

// Rebate API URLs
const RebateRestApiProdUrl = "https://api.binance.com"

// Simple Earn API URLs
const SimpleEarnRestApiProdUrl = "https://api.binance.com"

// Spot API URLs
const SpotRestApiProdUrl = "https://api.binance.com"
const SpotRestApiTestnetUrl = "https://testnet.binance.vision"
const SpotWebsocketApiProdUrl = "wss://ws-api.binance.com/ws-api/v3"
const SpotWebsocketApiTestnetUrl = "wss://ws-api.testnet.binance.vision/ws-api/v3"
const SpotWebsocketStreamsProdUrl = "wss://stream.binance.com:9443/stream"
const SpotWebsocketStreamsTestnetUrl = "wss://stream.testnet.binance.vision/stream"
const SpotRestApiMarketUrl = "https://data-api.binance.vision"
const SpotWebsocketStreamsMarketUrl = "wss://data-stream.binance.vision/stream"

// Stake API URLs
const StakingRestApiProdUrl = "https://api.binance.com"

// Sub-Account API URLs
const SubAccountRestApiProdUrl = "https://api.binance.com"

// VIP Loan API URLs
const VipLoanRestApiProdUrl = "https://api.binance.com"

// Wallet API URLs
const WalletRestApiProdUrl = "https://api.binance.com"
