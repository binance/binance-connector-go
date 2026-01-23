# GetExchangeInfoResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]GetExchangeInfoResponseDataAssetsInner**](GetExchangeInfoResponseDataAssetsInner.md) |  | [optional] 
**Symbols** | Pointer to [**[]GetExchangeInfoResponseDataSymbolsInner**](GetExchangeInfoResponseDataSymbolsInner.md) |  | [optional] 
**OrderTypes** | Pointer to **string** |  | [optional] 

## Methods

### NewGetExchangeInfoResponseData

`func NewGetExchangeInfoResponseData() *GetExchangeInfoResponseData`

NewGetExchangeInfoResponseData instantiates a new GetExchangeInfoResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeInfoResponseDataWithDefaults

`func NewGetExchangeInfoResponseDataWithDefaults() *GetExchangeInfoResponseData`

NewGetExchangeInfoResponseDataWithDefaults instantiates a new GetExchangeInfoResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *GetExchangeInfoResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetExchangeInfoResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetExchangeInfoResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetExchangeInfoResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAssets

`func (o *GetExchangeInfoResponseData) GetAssets() []GetExchangeInfoResponseDataAssetsInner`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *GetExchangeInfoResponseData) GetAssetsOk() (*[]GetExchangeInfoResponseDataAssetsInner, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *GetExchangeInfoResponseData) SetAssets(v []GetExchangeInfoResponseDataAssetsInner)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *GetExchangeInfoResponseData) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetSymbols

`func (o *GetExchangeInfoResponseData) GetSymbols() []GetExchangeInfoResponseDataSymbolsInner`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *GetExchangeInfoResponseData) GetSymbolsOk() (*[]GetExchangeInfoResponseDataSymbolsInner, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *GetExchangeInfoResponseData) SetSymbols(v []GetExchangeInfoResponseDataSymbolsInner)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *GetExchangeInfoResponseData) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetOrderTypes

`func (o *GetExchangeInfoResponseData) GetOrderTypes() string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *GetExchangeInfoResponseData) GetOrderTypesOk() (*string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *GetExchangeInfoResponseData) SetOrderTypes(v string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *GetExchangeInfoResponseData) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.


[[Back to README]](../README.md)


