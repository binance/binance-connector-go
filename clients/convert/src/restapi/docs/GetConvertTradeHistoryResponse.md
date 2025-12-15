# GetConvertTradeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]GetConvertTradeHistoryResponseListInner**](GetConvertTradeHistoryResponseListInner.md) |  | [optional] 
**StartTime** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**MoreData** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetConvertTradeHistoryResponse

`func NewGetConvertTradeHistoryResponse() *GetConvertTradeHistoryResponse`

NewGetConvertTradeHistoryResponse instantiates a new GetConvertTradeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConvertTradeHistoryResponseWithDefaults

`func NewGetConvertTradeHistoryResponseWithDefaults() *GetConvertTradeHistoryResponse`

NewGetConvertTradeHistoryResponseWithDefaults instantiates a new GetConvertTradeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *GetConvertTradeHistoryResponse) GetList() []GetConvertTradeHistoryResponseListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *GetConvertTradeHistoryResponse) GetListOk() (*[]GetConvertTradeHistoryResponseListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *GetConvertTradeHistoryResponse) SetList(v []GetConvertTradeHistoryResponseListInner)`

SetList sets List field to given value.

### HasList

`func (o *GetConvertTradeHistoryResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### GetStartTime

`func (o *GetConvertTradeHistoryResponse) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GetConvertTradeHistoryResponse) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GetConvertTradeHistoryResponse) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GetConvertTradeHistoryResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GetConvertTradeHistoryResponse) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GetConvertTradeHistoryResponse) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GetConvertTradeHistoryResponse) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GetConvertTradeHistoryResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLimit

`func (o *GetConvertTradeHistoryResponse) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetConvertTradeHistoryResponse) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetConvertTradeHistoryResponse) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetConvertTradeHistoryResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMoreData

`func (o *GetConvertTradeHistoryResponse) GetMoreData() bool`

GetMoreData returns the MoreData field if non-nil, zero value otherwise.

### GetMoreDataOk

`func (o *GetConvertTradeHistoryResponse) GetMoreDataOk() (*bool, bool)`

GetMoreDataOk returns a tuple with the MoreData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreData

`func (o *GetConvertTradeHistoryResponse) SetMoreData(v bool)`

SetMoreData sets MoreData field to given value.

### HasMoreData

`func (o *GetConvertTradeHistoryResponse) HasMoreData() bool`

HasMoreData returns a boolean if a field has been set.


[[Back to README]](../README.md)


