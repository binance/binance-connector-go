# ListOtcBlocktradesResponseBlocktradesInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**MarketId** | Pointer to **string** |  | [optional] 
**TokenId** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **int32** | 0 &#x3D; BUY (Bid), 1 &#x3D; SELL (Ask) — see Create OTC Blocktrade for the side/quoteType mapping | [optional] 
**Maker** | Pointer to **string** |  | [optional] 
**MakerAmount** | Pointer to **string** |  | [optional] 
**TakerAmount** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**SecretToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListOtcBlocktradesResponseBlocktradesInner

`func NewListOtcBlocktradesResponseBlocktradesInner() *ListOtcBlocktradesResponseBlocktradesInner`

NewListOtcBlocktradesResponseBlocktradesInner instantiates a new ListOtcBlocktradesResponseBlocktradesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOtcBlocktradesResponseBlocktradesInnerWithDefaults

`func NewListOtcBlocktradesResponseBlocktradesInnerWithDefaults() *ListOtcBlocktradesResponseBlocktradesInner`

NewListOtcBlocktradesResponseBlocktradesInnerWithDefaults instantiates a new ListOtcBlocktradesResponseBlocktradesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarketId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMarketId() string`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMarketIdOk() (*string, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMarketId(v string)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetTokenId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetSide

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSide() int32`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSideOk() (*int32, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetSide(v int32)`

SetSide sets Side field to given value.

### HasSide

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetMaker

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMaker() string`

GetMaker returns the Maker field if non-nil, zero value otherwise.

### GetMakerOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerOk() (*string, bool)`

GetMakerOk returns a tuple with the Maker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaker

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMaker(v string)`

SetMaker sets Maker field to given value.

### HasMaker

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMaker() bool`

HasMaker returns a boolean if a field has been set.

### GetMakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerAmount() string`

GetMakerAmount returns the MakerAmount field if non-nil, zero value otherwise.

### GetMakerAmountOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetMakerAmountOk() (*string, bool)`

GetMakerAmountOk returns a tuple with the MakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetMakerAmount(v string)`

SetMakerAmount sets MakerAmount field to given value.

### HasMakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasMakerAmount() bool`

HasMakerAmount returns a boolean if a field has been set.

### GetTakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTakerAmount() string`

GetTakerAmount returns the TakerAmount field if non-nil, zero value otherwise.

### GetTakerAmountOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetTakerAmountOk() (*string, bool)`

GetTakerAmountOk returns a tuple with the TakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetTakerAmount(v string)`

SetTakerAmount sets TakerAmount field to given value.

### HasTakerAmount

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasTakerAmount() bool`

HasTakerAmount returns a boolean if a field has been set.

### GetPrice

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuoteType

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSecretToken

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSecretToken() string`

GetSecretToken returns the SecretToken field if non-nil, zero value otherwise.

### GetSecretTokenOk

`func (o *ListOtcBlocktradesResponseBlocktradesInner) GetSecretTokenOk() (*string, bool)`

GetSecretTokenOk returns a tuple with the SecretToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretToken

`func (o *ListOtcBlocktradesResponseBlocktradesInner) SetSecretToken(v string)`

SetSecretToken sets SecretToken field to given value.

### HasSecretToken

`func (o *ListOtcBlocktradesResponseBlocktradesInner) HasSecretToken() bool`

HasSecretToken returns a boolean if a field has been set.


[[Back to README]](../README.md)


