# TokenizedConvertHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]TokenizedConvertHistoryResponseRowsInner**](TokenizedConvertHistoryResponseRowsInner.md) | Convert history rows on this page. Empty array if nothing matches. | [optional] 
**HasMore** | Pointer to **bool** | &#x60;true&#x60; when more pages exist — pass &#x60;nextLastId&#x60; as &#x60;lastId&#x60; on the next request. | [optional] 
**NextLastId** | Pointer to **int64** | Pass this value as &#x60;lastId&#x60; on the next request to get the following page. | [optional] 

## Methods

### NewTokenizedConvertHistoryResponse

`func NewTokenizedConvertHistoryResponse() *TokenizedConvertHistoryResponse`

NewTokenizedConvertHistoryResponse instantiates a new TokenizedConvertHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenizedConvertHistoryResponseWithDefaults

`func NewTokenizedConvertHistoryResponseWithDefaults() *TokenizedConvertHistoryResponse`

NewTokenizedConvertHistoryResponseWithDefaults instantiates a new TokenizedConvertHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *TokenizedConvertHistoryResponse) GetRows() []TokenizedConvertHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *TokenizedConvertHistoryResponse) GetRowsOk() (*[]TokenizedConvertHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *TokenizedConvertHistoryResponse) SetRows(v []TokenizedConvertHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *TokenizedConvertHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetHasMore

`func (o *TokenizedConvertHistoryResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TokenizedConvertHistoryResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TokenizedConvertHistoryResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *TokenizedConvertHistoryResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextLastId

`func (o *TokenizedConvertHistoryResponse) GetNextLastId() int64`

GetNextLastId returns the NextLastId field if non-nil, zero value otherwise.

### GetNextLastIdOk

`func (o *TokenizedConvertHistoryResponse) GetNextLastIdOk() (*int64, bool)`

GetNextLastIdOk returns a tuple with the NextLastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLastId

`func (o *TokenizedConvertHistoryResponse) SetNextLastId(v int64)`

SetNextLastId sets NextLastId field to given value.

### HasNextLastId

`func (o *TokenizedConvertHistoryResponse) HasNextLastId() bool`

HasNextLastId returns a boolean if a field has been set.


[[Back to README]](../README.md)


