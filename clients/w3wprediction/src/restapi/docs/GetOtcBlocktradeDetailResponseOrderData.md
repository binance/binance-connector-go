# GetOtcBlocktradeDetailResponseOrderData

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
**Price** | Pointer to **float64** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 
**FilledAmount** | Pointer to **string** |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**IsNegRisk** | Pointer to **bool** |  | [optional] 
**IsYieldBearing** | Pointer to **bool** |  | [optional] 
**SecretToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOtcBlocktradeDetailResponseOrderData

`func NewGetOtcBlocktradeDetailResponseOrderData() *GetOtcBlocktradeDetailResponseOrderData`

NewGetOtcBlocktradeDetailResponseOrderData instantiates a new GetOtcBlocktradeDetailResponseOrderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOtcBlocktradeDetailResponseOrderDataWithDefaults

`func NewGetOtcBlocktradeDetailResponseOrderDataWithDefaults() *GetOtcBlocktradeDetailResponseOrderData`

NewGetOtcBlocktradeDetailResponseOrderDataWithDefaults instantiates a new GetOtcBlocktradeDetailResponseOrderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarketId

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetTokenId

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSide

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetSide() int32`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetSideOk() (*int32, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetSide(v int32)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMaker() string`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerOk() (*string, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetMaker(v string)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetTaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTaker() string`

GetTaker returns the Taker field if non-nil, zero value otherwise.

### GetTakerOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerOk() (*string, bool)`

GetTakerOk returns a tuple with the Taker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetTaker(v string)`

SetTaker sets Taker field to given value.

### HasTaker

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasTaker() bool`

HasTaker returns a boolean if a field has been set.

### SetTakerNil

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetTakerNil(b bool)`

 SetTakerNil sets the value for Taker to be an explicit nil

### UnsetTaker
`func (o *GetOtcBlocktradeDetailResponseOrderData) UnsetTaker()`

UnsetTaker ensures that no value is present for Taker, not even an explicit nil
### GetMakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerAmount() string`

GetMakerAmount returns the MakerAmount field if non-nil, zero value otherwise.

### GetMakerAmountOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetMakerAmountOk() (*string, bool)`

GetMakerAmountOk returns a tuple with the MakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetMakerAmount(v string)`

SetMakerAmount sets MakerAmount field to given value.

### HasMakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasMakerAmount() bool`

HasMakerAmount returns a boolean if a field has been set.

### GetTakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerAmount() string`

GetTakerAmount returns the TakerAmount field if non-nil, zero value otherwise.

### GetTakerAmountOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTakerAmountOk() (*string, bool)`

GetTakerAmountOk returns a tuple with the TakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetTakerAmount(v string)`

SetTakerAmount sets TakerAmount field to given value.

### HasTakerAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasTakerAmount() bool`

HasTakerAmount returns a boolean if a field has been set.

### GetPrice

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrderType

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetExpiration

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetFilledAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetFilledAmount() string`

GetFilledAmount returns the FilledAmount field if non-nil, zero value otherwise.

### GetFilledAmountOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetFilledAmountOk() (*string, bool)`

GetFilledAmountOk returns a tuple with the FilledAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetFilledAmount(v string)`

SetFilledAmount sets FilledAmount field to given value.

### HasFilledAmount

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasFilledAmount() bool`

HasFilledAmount returns a boolean if a field has been set.

### GetQuoteType

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsNegRisk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsNegRisk() bool`

GetIsNegRisk returns the IsNegRisk field if non-nil, zero value otherwise.

### GetIsNegRiskOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsNegRiskOk() (*bool, bool)`

GetIsNegRiskOk returns a tuple with the IsNegRisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNegRisk

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetIsNegRisk(v bool)`

SetIsNegRisk sets IsNegRisk field to given value.

### HasIsNegRisk

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasIsNegRisk() bool`

HasIsNegRisk returns a boolean if a field has been set.

### GetIsYieldBearing

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsYieldBearing() bool`

GetIsYieldBearing returns the IsYieldBearing field if non-nil, zero value otherwise.

### GetIsYieldBearingOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetIsYieldBearingOk() (*bool, bool)`

GetIsYieldBearingOk returns a tuple with the IsYieldBearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYieldBearing

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetIsYieldBearing(v bool)`

SetIsYieldBearing sets IsYieldBearing field to given value.

### HasIsYieldBearing

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasIsYieldBearing() bool`

HasIsYieldBearing returns a boolean if a field has been set.

### GetSecretToken

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetSecretToken() string`

GetSecretToken returns the SecretToken field if non-nil, zero value otherwise.

### GetSecretTokenOk

`func (o *GetOtcBlocktradeDetailResponseOrderData) GetSecretTokenOk() (*string, bool)`

GetSecretTokenOk returns a tuple with the SecretToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretToken

`func (o *GetOtcBlocktradeDetailResponseOrderData) SetSecretToken(v string)`

SetSecretToken sets SecretToken field to given value.

### HasSecretToken

`func (o *GetOtcBlocktradeDetailResponseOrderData) HasSecretToken() bool`

HasSecretToken returns a boolean if a field has been set.


[[Back to README]](../README.md)


