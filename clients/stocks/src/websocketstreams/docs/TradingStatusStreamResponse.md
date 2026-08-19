# TradingStatusStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;tradingStatus\&quot;&#x60;. | [optional] 
**Symbol** | Pointer to **string** | Symbol (UPPERCASE ticker), e.g. &#x60;\&quot;AAPL\&quot;&#x60;. | [optional] 
**AssetCode** | Pointer to **string** | Internal asset code &#x60;EQ_{symbol}&#x60;; reference only. | [optional] 
**Status** | Pointer to **string** | Trading status. | [optional] 
**Msg** | Pointer to **string** | Reason code. | [optional] 
**Time** | Pointer to **int64** | Status-effective time (epoch milliseconds UTC). May be in the future. | [optional] 
**Z** | Pointer to **string** | Tape designation: &#x60;C&#x60; &#x3D; CTA (NYSE / AMEX), &#x60;N&#x60; &#x3D; UTP (Nasdaq). | [optional] 
**Tradability** | Pointer to **string** | Tradability after the status change: &#x60;BUY_SELL&#x60; / &#x60;BUY&#x60; / &#x60;SELL&#x60; / &#x60;NONE&#x60;. | [optional] 

## Methods

### NewTradingStatusStreamResponse

`func NewTradingStatusStreamResponse() *TradingStatusStreamResponse`

NewTradingStatusStreamResponse instantiates a new TradingStatusStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradingStatusStreamResponseWithDefaults

`func NewTradingStatusStreamResponseWithDefaults() *TradingStatusStreamResponse`

NewTradingStatusStreamResponseWithDefaults instantiates a new TradingStatusStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *TradingStatusStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *TradingStatusStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *TradingStatusStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *TradingStatusStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetSymbol

`func (o *TradingStatusStreamResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TradingStatusStreamResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TradingStatusStreamResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TradingStatusStreamResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAssetCode

`func (o *TradingStatusStreamResponse) GetAssetCode() string`

GetAssetCode returns the AssetCode field if non-nil, zero value otherwise.

### GetAssetCodeOk

`func (o *TradingStatusStreamResponse) GetAssetCodeOk() (*string, bool)`

GetAssetCodeOk returns a tuple with the AssetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCode

`func (o *TradingStatusStreamResponse) SetAssetCode(v string)`

SetAssetCode sets AssetCode field to given value.

### HasAssetCode

`func (o *TradingStatusStreamResponse) HasAssetCode() bool`

HasAssetCode returns a boolean if a field has been set.

### GetStatus

`func (o *TradingStatusStreamResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TradingStatusStreamResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TradingStatusStreamResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TradingStatusStreamResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *TradingStatusStreamResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TradingStatusStreamResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TradingStatusStreamResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *TradingStatusStreamResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTime

`func (o *TradingStatusStreamResponse) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *TradingStatusStreamResponse) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *TradingStatusStreamResponse) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *TradingStatusStreamResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetZ

`func (o *TradingStatusStreamResponse) GetZ() string`

GetZ returns the Z field if non-nil, zero value otherwise.

### GetZOk

`func (o *TradingStatusStreamResponse) GetZOk() (*string, bool)`

GetZOk returns a tuple with the Z field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZ

`func (o *TradingStatusStreamResponse) SetZ(v string)`

SetZ sets Z field to given value.

### HasZ

`func (o *TradingStatusStreamResponse) HasZ() bool`

HasZ returns a boolean if a field has been set.

### GetTradability

`func (o *TradingStatusStreamResponse) GetTradability() string`

GetTradability returns the Tradability field if non-nil, zero value otherwise.

### GetTradabilityOk

`func (o *TradingStatusStreamResponse) GetTradabilityOk() (*string, bool)`

GetTradabilityOk returns a tuple with the Tradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradability

`func (o *TradingStatusStreamResponse) SetTradability(v string)`

SetTradability sets Tradability field to given value.

### HasTradability

`func (o *TradingStatusStreamResponse) HasTradability() bool`

HasTradability returns a boolean if a field has been set.


[[Back to README]](../README.md)


