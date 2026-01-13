# QueryMarginRepayRecordResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]QueryMarginRepayRecordResponseRowsInner**](QueryMarginRepayRecordResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryMarginRepayRecordResponse

`func NewQueryMarginRepayRecordResponse() *QueryMarginRepayRecordResponse`

NewQueryMarginRepayRecordResponse instantiates a new QueryMarginRepayRecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginRepayRecordResponseWithDefaults

`func NewQueryMarginRepayRecordResponseWithDefaults() *QueryMarginRepayRecordResponse`

NewQueryMarginRepayRecordResponseWithDefaults instantiates a new QueryMarginRepayRecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *QueryMarginRepayRecordResponse) GetRows() []QueryMarginRepayRecordResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryMarginRepayRecordResponse) GetRowsOk() (*[]QueryMarginRepayRecordResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryMarginRepayRecordResponse) SetRows(v []QueryMarginRepayRecordResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryMarginRepayRecordResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *QueryMarginRepayRecordResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryMarginRepayRecordResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryMarginRepayRecordResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryMarginRepayRecordResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


