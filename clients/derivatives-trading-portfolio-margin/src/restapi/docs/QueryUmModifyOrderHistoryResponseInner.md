# QueryUmModifyOrderHistoryResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Amendment** | Pointer to [**QueryCmModifyOrderHistoryResponseInnerAmendment**](QueryCmModifyOrderHistoryResponseInnerAmendment.md) |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 

## Methods

### NewQueryUmModifyOrderHistoryResponseInner

`func NewQueryUmModifyOrderHistoryResponseInner() *QueryUmModifyOrderHistoryResponseInner`

NewQueryUmModifyOrderHistoryResponseInner instantiates a new QueryUmModifyOrderHistoryResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUmModifyOrderHistoryResponseInnerWithDefaults

`func NewQueryUmModifyOrderHistoryResponseInnerWithDefaults() *QueryUmModifyOrderHistoryResponseInner`

NewQueryUmModifyOrderHistoryResponseInnerWithDefaults instantiates a new QueryUmModifyOrderHistoryResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendmentId

`func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentId() int64`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentIdOk() (*int64, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *QueryUmModifyOrderHistoryResponseInner) SetAmendmentId(v int64)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *QueryUmModifyOrderHistoryResponseInner) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryUmModifyOrderHistoryResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryUmModifyOrderHistoryResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryUmModifyOrderHistoryResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPair

`func (o *QueryUmModifyOrderHistoryResponseInner) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *QueryUmModifyOrderHistoryResponseInner) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *QueryUmModifyOrderHistoryResponseInner) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryUmModifyOrderHistoryResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTime

`func (o *QueryUmModifyOrderHistoryResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueryUmModifyOrderHistoryResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueryUmModifyOrderHistoryResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAmendment

`func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendment() QueryCmModifyOrderHistoryResponseInnerAmendment`

GetAmendment returns the Amendment field if non-nil, zero value otherwise.

### GetAmendmentOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetAmendmentOk() (*QueryCmModifyOrderHistoryResponseInnerAmendment, bool)`

GetAmendmentOk returns a tuple with the Amendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendment

`func (o *QueryUmModifyOrderHistoryResponseInner) SetAmendment(v QueryCmModifyOrderHistoryResponseInnerAmendment)`

SetAmendment sets Amendment field to given value.

### HasAmendment

`func (o *QueryUmModifyOrderHistoryResponseInner) HasAmendment() bool`

HasAmendment returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryUmModifyOrderHistoryResponseInner) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryUmModifyOrderHistoryResponseInner) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryUmModifyOrderHistoryResponseInner) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryUmModifyOrderHistoryResponseInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.


[[Back to README]](../README.md)


