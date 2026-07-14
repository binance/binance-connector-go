# GetVIPLoanRepaymentHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetVIPLoanRepaymentHistoryResponseRowsInner**](GetVIPLoanRepaymentHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetVIPLoanRepaymentHistoryResponse

`func NewGetVIPLoanRepaymentHistoryResponse() *GetVIPLoanRepaymentHistoryResponse`

NewGetVIPLoanRepaymentHistoryResponse instantiates a new GetVIPLoanRepaymentHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVIPLoanRepaymentHistoryResponseWithDefaults

`func NewGetVIPLoanRepaymentHistoryResponseWithDefaults() *GetVIPLoanRepaymentHistoryResponse`

NewGetVIPLoanRepaymentHistoryResponseWithDefaults instantiates a new GetVIPLoanRepaymentHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetVIPLoanRepaymentHistoryResponse) GetRows() []GetVIPLoanRepaymentHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetVIPLoanRepaymentHistoryResponse) GetRowsOk() (*[]GetVIPLoanRepaymentHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetVIPLoanRepaymentHistoryResponse) SetRows(v []GetVIPLoanRepaymentHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetVIPLoanRepaymentHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetVIPLoanRepaymentHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVIPLoanRepaymentHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVIPLoanRepaymentHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVIPLoanRepaymentHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


