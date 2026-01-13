# GetCollateralRecordResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetCollateralRecordResponseRowsInner**](GetCollateralRecordResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetCollateralRecordResponse

`func NewGetCollateralRecordResponse() *GetCollateralRecordResponse`

NewGetCollateralRecordResponse instantiates a new GetCollateralRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollateralRecordResponseWithDefaults

`func NewGetCollateralRecordResponseWithDefaults() *GetCollateralRecordResponse`

NewGetCollateralRecordResponseWithDefaults instantiates a new GetCollateralRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetCollateralRecordResponse) GetRows() []GetCollateralRecordResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetCollateralRecordResponse) GetRowsOk() (*[]GetCollateralRecordResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetCollateralRecordResponse) SetRows(v []GetCollateralRecordResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetCollateralRecordResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetCollateralRecordResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetCollateralRecordResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetCollateralRecordResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetCollateralRecordResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


