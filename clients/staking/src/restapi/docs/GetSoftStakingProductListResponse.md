# GetSoftStakingProductListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **bool** |  | [optional] 
**TotalRewardsUsdt** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]GetSoftStakingProductListResponseRowsInner**](GetSoftStakingProductListResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetSoftStakingProductListResponse

`func NewGetSoftStakingProductListResponse() *GetSoftStakingProductListResponse`

NewGetSoftStakingProductListResponse instantiates a new GetSoftStakingProductListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSoftStakingProductListResponseWithDefaults

`func NewGetSoftStakingProductListResponseWithDefaults() *GetSoftStakingProductListResponse`

NewGetSoftStakingProductListResponseWithDefaults instantiates a new GetSoftStakingProductListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetSoftStakingProductListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSoftStakingProductListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSoftStakingProductListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSoftStakingProductListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalRewardsUsdt

`func (o *GetSoftStakingProductListResponse) GetTotalRewardsUsdt() string`

GetTotalRewardsUsdt returns the TotalRewardsUsdt field if non-nil, zero value otherwise.

### GetTotalRewardsUsdtOk

`func (o *GetSoftStakingProductListResponse) GetTotalRewardsUsdtOk() (*string, bool)`

GetTotalRewardsUsdtOk returns a tuple with the TotalRewardsUsdt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRewardsUsdt

`func (o *GetSoftStakingProductListResponse) SetTotalRewardsUsdt(v string)`

SetTotalRewardsUsdt sets TotalRewardsUsdt field to given value.

### HasTotalRewardsUsdt

`func (o *GetSoftStakingProductListResponse) HasTotalRewardsUsdt() bool`

HasTotalRewardsUsdt returns a boolean if a field has been set.

### GetRows

`func (o *GetSoftStakingProductListResponse) GetRows() []GetSoftStakingProductListResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetSoftStakingProductListResponse) GetRowsOk() (*[]GetSoftStakingProductListResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetSoftStakingProductListResponse) SetRows(v []GetSoftStakingProductListResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetSoftStakingProductListResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetSoftStakingProductListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSoftStakingProductListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSoftStakingProductListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetSoftStakingProductListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


