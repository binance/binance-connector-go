# GetBoostRewardsHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetBoostRewardsHistoryResponseRowsInner**](GetBoostRewardsHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBoostRewardsHistoryResponse

`func NewGetBoostRewardsHistoryResponse() *GetBoostRewardsHistoryResponse`

NewGetBoostRewardsHistoryResponse instantiates a new GetBoostRewardsHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBoostRewardsHistoryResponseWithDefaults

`func NewGetBoostRewardsHistoryResponseWithDefaults() *GetBoostRewardsHistoryResponse`

NewGetBoostRewardsHistoryResponseWithDefaults instantiates a new GetBoostRewardsHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetBoostRewardsHistoryResponse) GetRows() []GetBoostRewardsHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetBoostRewardsHistoryResponse) GetRowsOk() (*[]GetBoostRewardsHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetBoostRewardsHistoryResponse) SetRows(v []GetBoostRewardsHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetBoostRewardsHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetBoostRewardsHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBoostRewardsHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBoostRewardsHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBoostRewardsHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


