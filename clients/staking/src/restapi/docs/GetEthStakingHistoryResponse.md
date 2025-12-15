# GetEthStakingHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetEthStakingHistoryResponseRowsInner**](GetEthStakingHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetEthStakingHistoryResponse

`func NewGetEthStakingHistoryResponse() *GetEthStakingHistoryResponse`

NewGetEthStakingHistoryResponse instantiates a new GetEthStakingHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthStakingHistoryResponseWithDefaults

`func NewGetEthStakingHistoryResponseWithDefaults() *GetEthStakingHistoryResponse`

NewGetEthStakingHistoryResponseWithDefaults instantiates a new GetEthStakingHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetEthStakingHistoryResponse) GetRows() []GetEthStakingHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetEthStakingHistoryResponse) GetRowsOk() (*[]GetEthStakingHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetEthStakingHistoryResponse) SetRows(v []GetEthStakingHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetEthStakingHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetEthStakingHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEthStakingHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEthStakingHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetEthStakingHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


