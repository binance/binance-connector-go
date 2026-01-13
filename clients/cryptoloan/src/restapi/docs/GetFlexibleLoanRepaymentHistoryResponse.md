# GetFlexibleLoanRepaymentHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetFlexibleLoanRepaymentHistoryResponseRowsInner**](GetFlexibleLoanRepaymentHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFlexibleLoanRepaymentHistoryResponse

`func NewGetFlexibleLoanRepaymentHistoryResponse() *GetFlexibleLoanRepaymentHistoryResponse`

NewGetFlexibleLoanRepaymentHistoryResponse instantiates a new GetFlexibleLoanRepaymentHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlexibleLoanRepaymentHistoryResponseWithDefaults

`func NewGetFlexibleLoanRepaymentHistoryResponseWithDefaults() *GetFlexibleLoanRepaymentHistoryResponse`

NewGetFlexibleLoanRepaymentHistoryResponseWithDefaults instantiates a new GetFlexibleLoanRepaymentHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetFlexibleLoanRepaymentHistoryResponse) GetRows() []GetFlexibleLoanRepaymentHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetFlexibleLoanRepaymentHistoryResponse) GetRowsOk() (*[]GetFlexibleLoanRepaymentHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetFlexibleLoanRepaymentHistoryResponse) SetRows(v []GetFlexibleLoanRepaymentHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetFlexibleLoanRepaymentHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetFlexibleLoanRepaymentHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFlexibleLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFlexibleLoanRepaymentHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFlexibleLoanRepaymentHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


