# GetFlexibleLoanLtvAdjustmentHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner**](GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFlexibleLoanLtvAdjustmentHistoryResponse

`func NewGetFlexibleLoanLtvAdjustmentHistoryResponse() *GetFlexibleLoanLtvAdjustmentHistoryResponse`

NewGetFlexibleLoanLtvAdjustmentHistoryResponse instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlexibleLoanLtvAdjustmentHistoryResponseWithDefaults

`func NewGetFlexibleLoanLtvAdjustmentHistoryResponseWithDefaults() *GetFlexibleLoanLtvAdjustmentHistoryResponse`

NewGetFlexibleLoanLtvAdjustmentHistoryResponseWithDefaults instantiates a new GetFlexibleLoanLtvAdjustmentHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetRows() []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetRowsOk() (*[]GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) SetRows(v []GetFlexibleLoanLtvAdjustmentHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFlexibleLoanLtvAdjustmentHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


