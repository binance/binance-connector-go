# GetLoanRepaymentHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetLoanRepaymentHistoryResponseRowsInner**](GetLoanRepaymentHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetLoanRepaymentHistoryResponse

`func NewGetLoanRepaymentHistoryResponse() *GetLoanRepaymentHistoryResponse`

NewGetLoanRepaymentHistoryResponse instantiates a new GetLoanRepaymentHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoanRepaymentHistoryResponseWithDefaults

`func NewGetLoanRepaymentHistoryResponseWithDefaults() *GetLoanRepaymentHistoryResponse`

NewGetLoanRepaymentHistoryResponseWithDefaults instantiates a new GetLoanRepaymentHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetLoanRepaymentHistoryResponse) GetRows() []GetLoanRepaymentHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetLoanRepaymentHistoryResponse) GetRowsOk() (*[]GetLoanRepaymentHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetLoanRepaymentHistoryResponse) SetRows(v []GetLoanRepaymentHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetLoanRepaymentHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetLoanRepaymentHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLoanRepaymentHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLoanRepaymentHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


