# GetOrderModifyHistoryResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AmendmentId** | Pointer to **int64** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**Amendment** | Pointer to [**GetOrderModifyHistoryResponseInnerAmendment**](GetOrderModifyHistoryResponseInnerAmendment.md) |  | [optional] 

## Methods

### NewGetOrderModifyHistoryResponseInner

`func NewGetOrderModifyHistoryResponseInner() *GetOrderModifyHistoryResponseInner`

NewGetOrderModifyHistoryResponseInner instantiates a new GetOrderModifyHistoryResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderModifyHistoryResponseInnerWithDefaults

`func NewGetOrderModifyHistoryResponseInnerWithDefaults() *GetOrderModifyHistoryResponseInner`

NewGetOrderModifyHistoryResponseInnerWithDefaults instantiates a new GetOrderModifyHistoryResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmendmentId

`func (o *GetOrderModifyHistoryResponseInner) GetAmendmentId() int64`

GetAmendmentId returns the AmendmentId field if non-nil, zero value otherwise.

### GetAmendmentIdOk

`func (o *GetOrderModifyHistoryResponseInner) GetAmendmentIdOk() (*int64, bool)`

GetAmendmentIdOk returns a tuple with the AmendmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentId

`func (o *GetOrderModifyHistoryResponseInner) SetAmendmentId(v int64)`

SetAmendmentId sets AmendmentId field to given value.

### HasAmendmentId

`func (o *GetOrderModifyHistoryResponseInner) HasAmendmentId() bool`

HasAmendmentId returns a boolean if a field has been set.

### GetSymbol

`func (o *GetOrderModifyHistoryResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetOrderModifyHistoryResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetOrderModifyHistoryResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetOrderModifyHistoryResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPair

`func (o *GetOrderModifyHistoryResponseInner) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *GetOrderModifyHistoryResponseInner) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *GetOrderModifyHistoryResponseInner) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *GetOrderModifyHistoryResponseInner) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOrderModifyHistoryResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderModifyHistoryResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderModifyHistoryResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOrderModifyHistoryResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetClientOrderId

`func (o *GetOrderModifyHistoryResponseInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *GetOrderModifyHistoryResponseInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *GetOrderModifyHistoryResponseInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *GetOrderModifyHistoryResponseInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetTime

`func (o *GetOrderModifyHistoryResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetOrderModifyHistoryResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetOrderModifyHistoryResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetOrderModifyHistoryResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetAmendment

`func (o *GetOrderModifyHistoryResponseInner) GetAmendment() GetOrderModifyHistoryResponseInnerAmendment`

GetAmendment returns the Amendment field if non-nil, zero value otherwise.

### GetAmendmentOk

`func (o *GetOrderModifyHistoryResponseInner) GetAmendmentOk() (*GetOrderModifyHistoryResponseInnerAmendment, bool)`

GetAmendmentOk returns a tuple with the Amendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendment

`func (o *GetOrderModifyHistoryResponseInner) SetAmendment(v GetOrderModifyHistoryResponseInnerAmendment)`

SetAmendment sets Amendment field to given value.

### HasAmendment

`func (o *GetOrderModifyHistoryResponseInner) HasAmendment() bool`

HasAmendment returns a boolean if a field has been set.


[[Back to README]](../README.md)


