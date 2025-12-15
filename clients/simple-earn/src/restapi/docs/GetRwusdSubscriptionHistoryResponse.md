# GetRwusdSubscriptionHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetRwusdSubscriptionHistoryResponseRowsInner**](GetRwusdSubscriptionHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetRwusdSubscriptionHistoryResponse

`func NewGetRwusdSubscriptionHistoryResponse() *GetRwusdSubscriptionHistoryResponse`

NewGetRwusdSubscriptionHistoryResponse instantiates a new GetRwusdSubscriptionHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRwusdSubscriptionHistoryResponseWithDefaults

`func NewGetRwusdSubscriptionHistoryResponseWithDefaults() *GetRwusdSubscriptionHistoryResponse`

NewGetRwusdSubscriptionHistoryResponseWithDefaults instantiates a new GetRwusdSubscriptionHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetRwusdSubscriptionHistoryResponse) GetRows() []GetRwusdSubscriptionHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetRwusdSubscriptionHistoryResponse) GetRowsOk() (*[]GetRwusdSubscriptionHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetRwusdSubscriptionHistoryResponse) SetRows(v []GetRwusdSubscriptionHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetRwusdSubscriptionHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetRwusdSubscriptionHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRwusdSubscriptionHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRwusdSubscriptionHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetRwusdSubscriptionHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


