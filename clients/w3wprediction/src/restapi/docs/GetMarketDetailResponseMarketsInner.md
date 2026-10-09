# GetMarketDetailResponseMarketsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**MarketId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Question** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConditionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TradingStatus** | Pointer to **string** |  | [optional] 
**TradeVolume** | Pointer to **string** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**DecimalPrecision** | Pointer to **int32** | Price precision of the market, in decimal places. The price tick is 10 to the power of minus &#x60;decimalPrecision&#x60;, for example &#x60;2&#x60; means a tick of 0.01. A limit price that is not on the tick is truncated down | [optional] 
**Outcomes** | Pointer to [**[]GetMarketDetailResponseMarketsInnerOutcomesInner**](GetMarketDetailResponseMarketsInnerOutcomesInner.md) |  | [optional] 

## Methods

### NewGetMarketDetailResponseMarketsInner

`func NewGetMarketDetailResponseMarketsInner() *GetMarketDetailResponseMarketsInner`

NewGetMarketDetailResponseMarketsInner instantiates a new GetMarketDetailResponseMarketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketDetailResponseMarketsInnerWithDefaults

`func NewGetMarketDetailResponseMarketsInnerWithDefaults() *GetMarketDetailResponseMarketsInner`

NewGetMarketDetailResponseMarketsInnerWithDefaults instantiates a new GetMarketDetailResponseMarketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketId

`func (o *GetMarketDetailResponseMarketsInner) GetMarketId() int64`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *GetMarketDetailResponseMarketsInner) GetMarketIdOk() (*int64, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *GetMarketDetailResponseMarketsInner) SetMarketId(v int64)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *GetMarketDetailResponseMarketsInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetMarketDetailResponseMarketsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetMarketDetailResponseMarketsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetMarketDetailResponseMarketsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetMarketDetailResponseMarketsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTitle

`func (o *GetMarketDetailResponseMarketsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetMarketDetailResponseMarketsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetMarketDetailResponseMarketsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetMarketDetailResponseMarketsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetQuestion

`func (o *GetMarketDetailResponseMarketsInner) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *GetMarketDetailResponseMarketsInner) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *GetMarketDetailResponseMarketsInner) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *GetMarketDetailResponseMarketsInner) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetDescription

`func (o *GetMarketDetailResponseMarketsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMarketDetailResponseMarketsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMarketDetailResponseMarketsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMarketDetailResponseMarketsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConditionId

`func (o *GetMarketDetailResponseMarketsInner) GetConditionId() string`

GetConditionId returns the ConditionId field if non-nil, zero value otherwise.

### GetConditionIdOk

`func (o *GetMarketDetailResponseMarketsInner) GetConditionIdOk() (*string, bool)`

GetConditionIdOk returns a tuple with the ConditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionId

`func (o *GetMarketDetailResponseMarketsInner) SetConditionId(v string)`

SetConditionId sets ConditionId field to given value.

### HasConditionId

`func (o *GetMarketDetailResponseMarketsInner) HasConditionId() bool`

HasConditionId returns a boolean if a field has been set.

### GetStatus

`func (o *GetMarketDetailResponseMarketsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMarketDetailResponseMarketsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMarketDetailResponseMarketsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMarketDetailResponseMarketsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTradingStatus

`func (o *GetMarketDetailResponseMarketsInner) GetTradingStatus() string`

GetTradingStatus returns the TradingStatus field if non-nil, zero value otherwise.

### GetTradingStatusOk

`func (o *GetMarketDetailResponseMarketsInner) GetTradingStatusOk() (*string, bool)`

GetTradingStatusOk returns a tuple with the TradingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingStatus

`func (o *GetMarketDetailResponseMarketsInner) SetTradingStatus(v string)`

SetTradingStatus sets TradingStatus field to given value.

### HasTradingStatus

`func (o *GetMarketDetailResponseMarketsInner) HasTradingStatus() bool`

HasTradingStatus returns a boolean if a field has been set.

### GetTradeVolume

`func (o *GetMarketDetailResponseMarketsInner) GetTradeVolume() string`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *GetMarketDetailResponseMarketsInner) GetTradeVolumeOk() (*string, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *GetMarketDetailResponseMarketsInner) SetTradeVolume(v string)`

SetTradeVolume sets TradeVolume field to given value.

### HasTradeVolume

`func (o *GetMarketDetailResponseMarketsInner) HasTradeVolume() bool`

HasTradeVolume returns a boolean if a field has been set.

### GetLiquidity

`func (o *GetMarketDetailResponseMarketsInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *GetMarketDetailResponseMarketsInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *GetMarketDetailResponseMarketsInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *GetMarketDetailResponseMarketsInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetDecimalPrecision

`func (o *GetMarketDetailResponseMarketsInner) GetDecimalPrecision() int32`

GetDecimalPrecision returns the DecimalPrecision field if non-nil, zero value otherwise.

### GetDecimalPrecisionOk

`func (o *GetMarketDetailResponseMarketsInner) GetDecimalPrecisionOk() (*int32, bool)`

GetDecimalPrecisionOk returns a tuple with the DecimalPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalPrecision

`func (o *GetMarketDetailResponseMarketsInner) SetDecimalPrecision(v int32)`

SetDecimalPrecision sets DecimalPrecision field to given value.

### HasDecimalPrecision

`func (o *GetMarketDetailResponseMarketsInner) HasDecimalPrecision() bool`

HasDecimalPrecision returns a boolean if a field has been set.

### GetOutcomes

`func (o *GetMarketDetailResponseMarketsInner) GetOutcomes() []GetMarketDetailResponseMarketsInnerOutcomesInner`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *GetMarketDetailResponseMarketsInner) GetOutcomesOk() (*[]GetMarketDetailResponseMarketsInnerOutcomesInner, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *GetMarketDetailResponseMarketsInner) SetOutcomes(v []GetMarketDetailResponseMarketsInnerOutcomesInner)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *GetMarketDetailResponseMarketsInner) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.


[[Back to README]](../README.md)


