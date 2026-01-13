# GetAllCrossMarginPairsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Base** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsBuyAllowed** | Pointer to **bool** |  | [optional] 
**IsMarginTrade** | Pointer to **bool** |  | [optional] 
**IsSellAllowed** | Pointer to **bool** |  | [optional] 
**Quote** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**DelistTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetAllCrossMarginPairsResponseInner

`func NewGetAllCrossMarginPairsResponseInner() *GetAllCrossMarginPairsResponseInner`

NewGetAllCrossMarginPairsResponseInner instantiates a new GetAllCrossMarginPairsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllCrossMarginPairsResponseInnerWithDefaults

`func NewGetAllCrossMarginPairsResponseInnerWithDefaults() *GetAllCrossMarginPairsResponseInner`

NewGetAllCrossMarginPairsResponseInnerWithDefaults instantiates a new GetAllCrossMarginPairsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase

`func (o *GetAllCrossMarginPairsResponseInner) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *GetAllCrossMarginPairsResponseInner) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *GetAllCrossMarginPairsResponseInner) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *GetAllCrossMarginPairsResponseInner) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetId

`func (o *GetAllCrossMarginPairsResponseInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllCrossMarginPairsResponseInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllCrossMarginPairsResponseInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetAllCrossMarginPairsResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBuyAllowed

`func (o *GetAllCrossMarginPairsResponseInner) GetIsBuyAllowed() bool`

GetIsBuyAllowed returns the IsBuyAllowed field if non-nil, zero value otherwise.

### GetIsBuyAllowedOk

`func (o *GetAllCrossMarginPairsResponseInner) GetIsBuyAllowedOk() (*bool, bool)`

GetIsBuyAllowedOk returns a tuple with the IsBuyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyAllowed

`func (o *GetAllCrossMarginPairsResponseInner) SetIsBuyAllowed(v bool)`

SetIsBuyAllowed sets IsBuyAllowed field to given value.

### HasIsBuyAllowed

`func (o *GetAllCrossMarginPairsResponseInner) HasIsBuyAllowed() bool`

HasIsBuyAllowed returns a boolean if a field has been set.

### GetIsMarginTrade

`func (o *GetAllCrossMarginPairsResponseInner) GetIsMarginTrade() bool`

GetIsMarginTrade returns the IsMarginTrade field if non-nil, zero value otherwise.

### GetIsMarginTradeOk

`func (o *GetAllCrossMarginPairsResponseInner) GetIsMarginTradeOk() (*bool, bool)`

GetIsMarginTradeOk returns a tuple with the IsMarginTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTrade

`func (o *GetAllCrossMarginPairsResponseInner) SetIsMarginTrade(v bool)`

SetIsMarginTrade sets IsMarginTrade field to given value.

### HasIsMarginTrade

`func (o *GetAllCrossMarginPairsResponseInner) HasIsMarginTrade() bool`

HasIsMarginTrade returns a boolean if a field has been set.

### GetIsSellAllowed

`func (o *GetAllCrossMarginPairsResponseInner) GetIsSellAllowed() bool`

GetIsSellAllowed returns the IsSellAllowed field if non-nil, zero value otherwise.

### GetIsSellAllowedOk

`func (o *GetAllCrossMarginPairsResponseInner) GetIsSellAllowedOk() (*bool, bool)`

GetIsSellAllowedOk returns a tuple with the IsSellAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSellAllowed

`func (o *GetAllCrossMarginPairsResponseInner) SetIsSellAllowed(v bool)`

SetIsSellAllowed sets IsSellAllowed field to given value.

### HasIsSellAllowed

`func (o *GetAllCrossMarginPairsResponseInner) HasIsSellAllowed() bool`

HasIsSellAllowed returns a boolean if a field has been set.

### GetQuote

`func (o *GetAllCrossMarginPairsResponseInner) GetQuote() string`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *GetAllCrossMarginPairsResponseInner) GetQuoteOk() (*string, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *GetAllCrossMarginPairsResponseInner) SetQuote(v string)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *GetAllCrossMarginPairsResponseInner) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetSymbol

`func (o *GetAllCrossMarginPairsResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetAllCrossMarginPairsResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetAllCrossMarginPairsResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetAllCrossMarginPairsResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDelistTime

`func (o *GetAllCrossMarginPairsResponseInner) GetDelistTime() int64`

GetDelistTime returns the DelistTime field if non-nil, zero value otherwise.

### GetDelistTimeOk

`func (o *GetAllCrossMarginPairsResponseInner) GetDelistTimeOk() (*int64, bool)`

GetDelistTimeOk returns a tuple with the DelistTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistTime

`func (o *GetAllCrossMarginPairsResponseInner) SetDelistTime(v int64)`

SetDelistTime sets DelistTime field to given value.

### HasDelistTime

`func (o *GetAllCrossMarginPairsResponseInner) HasDelistTime() bool`

HasDelistTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


