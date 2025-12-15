# QueryUserDelegationHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Rows** | Pointer to [**[]QueryUserDelegationHistoryResponseRowsInner**](QueryUserDelegationHistoryResponseRowsInner.md) |  | [optional] 

## Methods

### NewQueryUserDelegationHistoryResponse

`func NewQueryUserDelegationHistoryResponse() *QueryUserDelegationHistoryResponse`

NewQueryUserDelegationHistoryResponse instantiates a new QueryUserDelegationHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryUserDelegationHistoryResponseWithDefaults

`func NewQueryUserDelegationHistoryResponseWithDefaults() *QueryUserDelegationHistoryResponse`

NewQueryUserDelegationHistoryResponseWithDefaults instantiates a new QueryUserDelegationHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryUserDelegationHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryUserDelegationHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryUserDelegationHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryUserDelegationHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *QueryUserDelegationHistoryResponse) GetRows() []QueryUserDelegationHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryUserDelegationHistoryResponse) GetRowsOk() (*[]QueryUserDelegationHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryUserDelegationHistoryResponse) SetRows(v []QueryUserDelegationHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryUserDelegationHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


