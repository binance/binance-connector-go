# TradabilityStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;tradability\&quot;&#x60;. | [optional] 
**Symbol** | Pointer to **string** | Symbol (UPPERCASE ticker). | [optional] 
**AssetCode** | Pointer to **NullableString** | Internal asset code &#x60;EQ_{symbol}&#x60;; reference only. May be null for symbols not yet in the symbol dictionary. | [optional] 
**Tradability** | Pointer to **string** | The new tradability value. &#x60;OFFMARKET&#x60; and &#x60;DELISTING&#x60; are protected states — trading-status events will not transition a symbol out of them. | [optional] 
**T** | Pointer to **int64** | Push time (epoch milliseconds UTC). | [optional] 

## Methods

### NewTradabilityStreamResponse

`func NewTradabilityStreamResponse() *TradabilityStreamResponse`

NewTradabilityStreamResponse instantiates a new TradabilityStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradabilityStreamResponseWithDefaults

`func NewTradabilityStreamResponseWithDefaults() *TradabilityStreamResponse`

NewTradabilityStreamResponseWithDefaults instantiates a new TradabilityStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *TradabilityStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *TradabilityStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *TradabilityStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *TradabilityStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetSymbol

`func (o *TradabilityStreamResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TradabilityStreamResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TradabilityStreamResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TradabilityStreamResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAssetCode

`func (o *TradabilityStreamResponse) GetAssetCode() string`

GetAssetCode returns the AssetCode field if non-nil, zero value otherwise.

### GetAssetCodeOk

`func (o *TradabilityStreamResponse) GetAssetCodeOk() (*string, bool)`

GetAssetCodeOk returns a tuple with the AssetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCode

`func (o *TradabilityStreamResponse) SetAssetCode(v string)`

SetAssetCode sets AssetCode field to given value.

### HasAssetCode

`func (o *TradabilityStreamResponse) HasAssetCode() bool`

HasAssetCode returns a boolean if a field has been set.

### SetAssetCodeNil

`func (o *TradabilityStreamResponse) SetAssetCodeNil(b bool)`

 SetAssetCodeNil sets the value for AssetCode to be an explicit nil

### UnsetAssetCode
`func (o *TradabilityStreamResponse) UnsetAssetCode()`

UnsetAssetCode ensures that no value is present for AssetCode, not even an explicit nil
### GetTradability

`func (o *TradabilityStreamResponse) GetTradability() string`

GetTradability returns the Tradability field if non-nil, zero value otherwise.

### GetTradabilityOk

`func (o *TradabilityStreamResponse) GetTradabilityOk() (*string, bool)`

GetTradabilityOk returns a tuple with the Tradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradability

`func (o *TradabilityStreamResponse) SetTradability(v string)`

SetTradability sets Tradability field to given value.

### HasTradability

`func (o *TradabilityStreamResponse) HasTradability() bool`

HasTradability returns a boolean if a field has been set.

### GetT

`func (o *TradabilityStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TradabilityStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TradabilityStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *TradabilityStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.


[[Back to README]](../README.md)


