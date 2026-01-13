# QueryUniversalTransferHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]QueryUniversalTransferHistoryResponseResultInner**](QueryUniversalTransferHistoryResponseResultInner.md) |  | [optional] 
**TotalCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryUniversalTransferHistoryResponse

`func NewQueryUniversalTransferHistoryResponse() *QueryUniversalTransferHistoryResponse`

NewQueryUniversalTransferHistoryResponse instantiates a new QueryUniversalTransferHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUniversalTransferHistoryResponseWithDefaults

`func NewQueryUniversalTransferHistoryResponseWithDefaults() *QueryUniversalTransferHistoryResponse`

NewQueryUniversalTransferHistoryResponseWithDefaults instantiates a new QueryUniversalTransferHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *QueryUniversalTransferHistoryResponse) GetResult() []QueryUniversalTransferHistoryResponseResultInner`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QueryUniversalTransferHistoryResponse) GetResultOk() (*[]QueryUniversalTransferHistoryResponseResultInner, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QueryUniversalTransferHistoryResponse) SetResult(v []QueryUniversalTransferHistoryResponseResultInner)`

SetResult sets Result field to given value.

### HasResult

`func (o *QueryUniversalTransferHistoryResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetTotalCount

`func (o *QueryUniversalTransferHistoryResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *QueryUniversalTransferHistoryResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *QueryUniversalTransferHistoryResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *QueryUniversalTransferHistoryResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to README]](../README.md)


