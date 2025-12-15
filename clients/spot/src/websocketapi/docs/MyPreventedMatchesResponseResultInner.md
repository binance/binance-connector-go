# MyPreventedMatchesResponseResultInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PreventedMatchId** | Pointer to **int64** |  | [optional] 
**TakerOrderId** | Pointer to **int64** |  | [optional] 
**MakerSymbol** | Pointer to **string** |  | [optional] 
**MakerOrderId** | Pointer to **int64** |  | [optional] 
**TradeGroupId** | Pointer to **int64** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**MakerPreventedQuantity** | Pointer to **string** |  | [optional] 
**TransactTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewMyPreventedMatchesResponseResultInner

`func NewMyPreventedMatchesResponseResultInner() *MyPreventedMatchesResponseResultInner`

NewMyPreventedMatchesResponseResultInner instantiates a new MyPreventedMatchesResponseResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyPreventedMatchesResponseResultInnerWithDefaults

`func NewMyPreventedMatchesResponseResultInnerWithDefaults() *MyPreventedMatchesResponseResultInner`

NewMyPreventedMatchesResponseResultInnerWithDefaults instantiates a new MyPreventedMatchesResponseResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MyPreventedMatchesResponseResultInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MyPreventedMatchesResponseResultInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MyPreventedMatchesResponseResultInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MyPreventedMatchesResponseResultInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPreventedMatchId

`func (o *MyPreventedMatchesResponseResultInner) GetPreventedMatchId() int64`

GetPreventedMatchId returns the PreventedMatchId field if non-nil, zero value otherwise.

### GetPreventedMatchIdOk

`func (o *MyPreventedMatchesResponseResultInner) GetPreventedMatchIdOk() (*int64, bool)`

GetPreventedMatchIdOk returns a tuple with the PreventedMatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventedMatchId

`func (o *MyPreventedMatchesResponseResultInner) SetPreventedMatchId(v int64)`

SetPreventedMatchId sets PreventedMatchId field to given value.

### HasPreventedMatchId

`func (o *MyPreventedMatchesResponseResultInner) HasPreventedMatchId() bool`

HasPreventedMatchId returns a boolean if a field has been set.

### GetTakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) GetTakerOrderId() int64`

GetTakerOrderId returns the TakerOrderId field if non-nil, zero value otherwise.

### GetTakerOrderIdOk

`func (o *MyPreventedMatchesResponseResultInner) GetTakerOrderIdOk() (*int64, bool)`

GetTakerOrderIdOk returns a tuple with the TakerOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) SetTakerOrderId(v int64)`

SetTakerOrderId sets TakerOrderId field to given value.

### HasTakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) HasTakerOrderId() bool`

HasTakerOrderId returns a boolean if a field has been set.

### GetMakerSymbol

`func (o *MyPreventedMatchesResponseResultInner) GetMakerSymbol() string`

GetMakerSymbol returns the MakerSymbol field if non-nil, zero value otherwise.

### GetMakerSymbolOk

`func (o *MyPreventedMatchesResponseResultInner) GetMakerSymbolOk() (*string, bool)`

GetMakerSymbolOk returns a tuple with the MakerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerSymbol

`func (o *MyPreventedMatchesResponseResultInner) SetMakerSymbol(v string)`

SetMakerSymbol sets MakerSymbol field to given value.

### HasMakerSymbol

`func (o *MyPreventedMatchesResponseResultInner) HasMakerSymbol() bool`

HasMakerSymbol returns a boolean if a field has been set.

### GetMakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) GetMakerOrderId() int64`

GetMakerOrderId returns the MakerOrderId field if non-nil, zero value otherwise.

### GetMakerOrderIdOk

`func (o *MyPreventedMatchesResponseResultInner) GetMakerOrderIdOk() (*int64, bool)`

GetMakerOrderIdOk returns a tuple with the MakerOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) SetMakerOrderId(v int64)`

SetMakerOrderId sets MakerOrderId field to given value.

### HasMakerOrderId

`func (o *MyPreventedMatchesResponseResultInner) HasMakerOrderId() bool`

HasMakerOrderId returns a boolean if a field has been set.

### GetTradeGroupId

`func (o *MyPreventedMatchesResponseResultInner) GetTradeGroupId() int64`

GetTradeGroupId returns the TradeGroupId field if non-nil, zero value otherwise.

### GetTradeGroupIdOk

`func (o *MyPreventedMatchesResponseResultInner) GetTradeGroupIdOk() (*int64, bool)`

GetTradeGroupIdOk returns a tuple with the TradeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGroupId

`func (o *MyPreventedMatchesResponseResultInner) SetTradeGroupId(v int64)`

SetTradeGroupId sets TradeGroupId field to given value.

### HasTradeGroupId

`func (o *MyPreventedMatchesResponseResultInner) HasTradeGroupId() bool`

HasTradeGroupId returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *MyPreventedMatchesResponseResultInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *MyPreventedMatchesResponseResultInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *MyPreventedMatchesResponseResultInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *MyPreventedMatchesResponseResultInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetPrice

`func (o *MyPreventedMatchesResponseResultInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MyPreventedMatchesResponseResultInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MyPreventedMatchesResponseResultInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MyPreventedMatchesResponseResultInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetMakerPreventedQuantity

`func (o *MyPreventedMatchesResponseResultInner) GetMakerPreventedQuantity() string`

GetMakerPreventedQuantity returns the MakerPreventedQuantity field if non-nil, zero value otherwise.

### GetMakerPreventedQuantityOk

`func (o *MyPreventedMatchesResponseResultInner) GetMakerPreventedQuantityOk() (*string, bool)`

GetMakerPreventedQuantityOk returns a tuple with the MakerPreventedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMakerPreventedQuantity

`func (o *MyPreventedMatchesResponseResultInner) SetMakerPreventedQuantity(v string)`

SetMakerPreventedQuantity sets MakerPreventedQuantity field to given value.

### HasMakerPreventedQuantity

`func (o *MyPreventedMatchesResponseResultInner) HasMakerPreventedQuantity() bool`

HasMakerPreventedQuantity returns a boolean if a field has been set.

### GetTransactTime

`func (o *MyPreventedMatchesResponseResultInner) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *MyPreventedMatchesResponseResultInner) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *MyPreventedMatchesResponseResultInner) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *MyPreventedMatchesResponseResultInner) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


