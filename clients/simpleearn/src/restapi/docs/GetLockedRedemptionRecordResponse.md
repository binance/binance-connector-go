# GetLockedRedemptionRecordResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLockedRedemptionRecordResponseRowsInner**](GetLockedRedemptionRecordResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetLockedRedemptionRecordResponse

`func NewGetLockedRedemptionRecordResponse() *GetLockedRedemptionRecordResponse`

NewGetLockedRedemptionRecordResponse instantiates a new GetLockedRedemptionRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLockedRedemptionRecordResponseWithDefaults

`func NewGetLockedRedemptionRecordResponseWithDefaults() *GetLockedRedemptionRecordResponse`

NewGetLockedRedemptionRecordResponseWithDefaults instantiates a new GetLockedRedemptionRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLockedRedemptionRecordResponse) GetRows() []GetLockedRedemptionRecordResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLockedRedemptionRecordResponse) GetRowsOk() (*[]GetLockedRedemptionRecordResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLockedRedemptionRecordResponse) SetRows(v []GetLockedRedemptionRecordResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLockedRedemptionRecordResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLockedRedemptionRecordResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLockedRedemptionRecordResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLockedRedemptionRecordResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLockedRedemptionRecordResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


