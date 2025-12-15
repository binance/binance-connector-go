# GetBnsolRewardsHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**EstRewardsInSOL** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]GetBnsolRewardsHistoryResponseRowsInner**](GetBnsolRewardsHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetBnsolRewardsHistoryResponse

`func NewGetBnsolRewardsHistoryResponse() *GetBnsolRewardsHistoryResponse`

NewGetBnsolRewardsHistoryResponse instantiates a new GetBnsolRewardsHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBnsolRewardsHistoryResponseWithDefaults

`func NewGetBnsolRewardsHistoryResponseWithDefaults() *GetBnsolRewardsHistoryResponse`

NewGetBnsolRewardsHistoryResponseWithDefaults instantiates a new GetBnsolRewardsHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstRewardsInSOL

`func (o *GetBnsolRewardsHistoryResponse) GetEstRewardsInSOL() string`

GetEstRewardsInSOL returns the EstRewardsInSOL field if non-nil, zero value otherwise.

### GetEstRewardsInSOLOk

`func (o *GetBnsolRewardsHistoryResponse) GetEstRewardsInSOLOk() (*string, bool)`

GetEstRewardsInSOLOk returns a tuple with the EstRewardsInSOL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstRewardsInSOL

`func (o *GetBnsolRewardsHistoryResponse) SetEstRewardsInSOL(v string)`

SetEstRewardsInSOL sets EstRewardsInSOL field to given value.

### HasEstRewardsInSOL

`func (o *GetBnsolRewardsHistoryResponse) HasEstRewardsInSOL() bool`

HasEstRewardsInSOL returns a boolean if a field has been set.

### GetRows

`func (o *GetBnsolRewardsHistoryResponse) GetRows() []GetBnsolRewardsHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetBnsolRewardsHistoryResponse) GetRowsOk() (*[]GetBnsolRewardsHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetBnsolRewardsHistoryResponse) SetRows(v []GetBnsolRewardsHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetBnsolRewardsHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetBnsolRewardsHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBnsolRewardsHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBnsolRewardsHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBnsolRewardsHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


