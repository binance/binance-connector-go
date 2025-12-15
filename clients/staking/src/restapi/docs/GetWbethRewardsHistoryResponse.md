# GetWbethRewardsHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**EstRewardsInETH** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]GetWbethRewardsHistoryResponseRowsInner**](GetWbethRewardsHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetWbethRewardsHistoryResponse

`func NewGetWbethRewardsHistoryResponse() *GetWbethRewardsHistoryResponse`

NewGetWbethRewardsHistoryResponse instantiates a new GetWbethRewardsHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWbethRewardsHistoryResponseWithDefaults

`func NewGetWbethRewardsHistoryResponseWithDefaults() *GetWbethRewardsHistoryResponse`

NewGetWbethRewardsHistoryResponseWithDefaults instantiates a new GetWbethRewardsHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstRewardsInETH

`func (o *GetWbethRewardsHistoryResponse) GetEstRewardsInETH() string`

GetEstRewardsInETH returns the EstRewardsInETH field if non-nil, zero value otherwise.

### GetEstRewardsInETHOk

`func (o *GetWbethRewardsHistoryResponse) GetEstRewardsInETHOk() (*string, bool)`

GetEstRewardsInETHOk returns a tuple with the EstRewardsInETH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstRewardsInETH

`func (o *GetWbethRewardsHistoryResponse) SetEstRewardsInETH(v string)`

SetEstRewardsInETH sets EstRewardsInETH field to given value.

### HasEstRewardsInETH

`func (o *GetWbethRewardsHistoryResponse) HasEstRewardsInETH() bool`

HasEstRewardsInETH returns a boolean if a field has been set.

### GetRows

`func (o *GetWbethRewardsHistoryResponse) GetRows() []GetWbethRewardsHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetWbethRewardsHistoryResponse) GetRowsOk() (*[]GetWbethRewardsHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetWbethRewardsHistoryResponse) SetRows(v []GetWbethRewardsHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetWbethRewardsHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetWbethRewardsHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetWbethRewardsHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetWbethRewardsHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetWbethRewardsHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


