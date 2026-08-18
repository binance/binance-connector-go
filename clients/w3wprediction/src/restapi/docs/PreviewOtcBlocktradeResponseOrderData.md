# PreviewOtcBlocktradeResponseOrderData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **int32** | 0 &#x3D; BUY (Bid), 1 &#x3D; SELL (Ask) — see Create OTC Blocktrade for the side/quoteType mapping | [optional] 
**Maker** | Pointer to **string** |  | [optional] 
**Taker** | Pointer to **NullableString** |  | [optional] 
**MakerAmount** | Pointer to **string** |  | [optional] 
**TakerAmount** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 
**FilledAmount** | Pointer to **string** |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**IsNegRisk** | Pointer to **bool** |  | [optional] 
**IsYieldBearing** | Pointer to **bool** |  | [optional] 

## Methods

### NewPreviewOtcBlocktradeResponseOrderData

`func NewPreviewOtcBlocktradeResponseOrderData() *PreviewOtcBlocktradeResponseOrderData`

NewPreviewOtcBlocktradeResponseOrderData instantiates a new PreviewOtcBlocktradeResponseOrderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewOtcBlocktradeResponseOrderDataWithDefaults

`func NewPreviewOtcBlocktradeResponseOrderDataWithDefaults() *PreviewOtcBlocktradeResponseOrderData`

NewPreviewOtcBlocktradeResponseOrderDataWithDefaults instantiates a new PreviewOtcBlocktradeResponseOrderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PreviewOtcBlocktradeResponseOrderData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreviewOtcBlocktradeResponseOrderData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreviewOtcBlocktradeResponseOrderData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarketId

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *PreviewOtcBlocktradeResponseOrderData) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *PreviewOtcBlocktradeResponseOrderData) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetTokenId

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *PreviewOtcBlocktradeResponseOrderData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *PreviewOtcBlocktradeResponseOrderData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSide

`func (o *PreviewOtcBlocktradeResponseOrderData) GetSide() int32`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetSideOk() (*int32, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PreviewOtcBlocktradeResponseOrderData) SetSide(v int32)`

SetSide sets Side field to given value.

### HasSide

`func (o *PreviewOtcBlocktradeResponseOrderData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMaker

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMaker() string`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMakerOk() (*string, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *PreviewOtcBlocktradeResponseOrderData) SetMaker(v string)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *PreviewOtcBlocktradeResponseOrderData) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetTaker

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *PreviewOtcBlocktradeResponseOrderData) SetTaker(v string)`

SetTaker sets Taker field to given value.

### HasTaker

`func (o *PreviewOtcBlocktradeResponseOrderData) HasTaker() bool`

HasTaker returns a boolean if a field has been set.

### SetTakerNil

`func (o *PreviewOtcBlocktradeResponseOrderData) SetTakerNil(b bool)`

 SetTakerNil sets the value for Taker to be an explicit nil

### UnsetTaker
`func (o *PreviewOtcBlocktradeResponseOrderData) UnsetTaker()`

UnsetTaker ensures that no value is present for Taker, not even an explicit nil
### GetMakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMakerAmount() string`

GetMakerAmount returns the MakerAmount field if non-nil, zero value otherwise.

### GetMakerAmountOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetMakerAmountOk() (*string, bool)`

GetMakerAmountOk returns a tuple with the MakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) SetMakerAmount(v string)`

SetMakerAmount sets MakerAmount field to given value.

### HasMakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) HasMakerAmount() bool`

HasMakerAmount returns a boolean if a field has been set.

### GetTakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTakerAmount() string`

GetTakerAmount returns the TakerAmount field if non-nil, zero value otherwise.

### GetTakerAmountOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTakerAmountOk() (*string, bool)`

GetTakerAmountOk returns a tuple with the TakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) SetTakerAmount(v string)`

SetTakerAmount sets TakerAmount field to given value.

### HasTakerAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) HasTakerAmount() bool`

HasTakerAmount returns a boolean if a field has been set.

### GetPrice

`func (o *PreviewOtcBlocktradeResponseOrderData) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PreviewOtcBlocktradeResponseOrderData) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PreviewOtcBlocktradeResponseOrderData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *PreviewOtcBlocktradeResponseOrderData) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *PreviewOtcBlocktradeResponseOrderData) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *PreviewOtcBlocktradeResponseOrderData) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PreviewOtcBlocktradeResponseOrderData) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PreviewOtcBlocktradeResponseOrderData) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetExpiration

`func (o *PreviewOtcBlocktradeResponseOrderData) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *PreviewOtcBlocktradeResponseOrderData) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *PreviewOtcBlocktradeResponseOrderData) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFilledAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) GetFilledAmount() string`

GetFilledAmount returns the FilledAmount field if non-nil, zero value otherwise.

### GetFilledAmountOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetFilledAmountOk() (*string, bool)`

GetFilledAmountOk returns a tuple with the FilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) SetFilledAmount(v string)`

SetFilledAmount sets FilledAmount field to given value.

### HasFilledAmount

`func (o *PreviewOtcBlocktradeResponseOrderData) HasFilledAmount() bool`

HasFilledAmount returns a boolean if a field has been set.

### GetQuoteType

`func (o *PreviewOtcBlocktradeResponseOrderData) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *PreviewOtcBlocktradeResponseOrderData) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *PreviewOtcBlocktradeResponseOrderData) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PreviewOtcBlocktradeResponseOrderData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PreviewOtcBlocktradeResponseOrderData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PreviewOtcBlocktradeResponseOrderData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsNegRisk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetIsNegRisk() bool`

GetIsNegRisk returns the IsNegRisk field if non-nil, zero value otherwise.

### GetIsNegRiskOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetIsNegRiskOk() (*bool, bool)`

GetIsNegRiskOk returns a tuple with the IsNegRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegRisk

`func (o *PreviewOtcBlocktradeResponseOrderData) SetIsNegRisk(v bool)`

SetIsNegRisk sets IsNegRisk field to given value.

### HasIsNegRisk

`func (o *PreviewOtcBlocktradeResponseOrderData) HasIsNegRisk() bool`

HasIsNegRisk returns a boolean if a field has been set.

### GetIsYieldBearing

`func (o *PreviewOtcBlocktradeResponseOrderData) GetIsYieldBearing() bool`

GetIsYieldBearing returns the IsYieldBearing field if non-nil, zero value otherwise.

### GetIsYieldBearingOk

`func (o *PreviewOtcBlocktradeResponseOrderData) GetIsYieldBearingOk() (*bool, bool)`

GetIsYieldBearingOk returns a tuple with the IsYieldBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYieldBearing

`func (o *PreviewOtcBlocktradeResponseOrderData) SetIsYieldBearing(v bool)`

SetIsYieldBearing sets IsYieldBearing field to given value.

### HasIsYieldBearing

`func (o *PreviewOtcBlocktradeResponseOrderData) HasIsYieldBearing() bool`

HasIsYieldBearing returns a boolean if a field has been set.


[[Back to README]](../README.md)


