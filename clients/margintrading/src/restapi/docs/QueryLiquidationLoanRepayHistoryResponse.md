# QueryLiquidationLoanRepayHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** | Total number of repayment records | [optional] 
**Rows** | Pointer to [**[]QueryLiquidationLoanRepayHistoryResponseRowsInner**](QueryLiquidationLoanRepayHistoryResponseRowsInner.md) |  | [optional] 

## Methods

### NewQueryLiquidationLoanRepayHistoryResponse

`func NewQueryLiquidationLoanRepayHistoryResponse() *QueryLiquidationLoanRepayHistoryResponse`

NewQueryLiquidationLoanRepayHistoryResponse instantiates a new QueryLiquidationLoanRepayHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLiquidationLoanRepayHistoryResponseWithDefaults

`func NewQueryLiquidationLoanRepayHistoryResponseWithDefaults() *QueryLiquidationLoanRepayHistoryResponse`

NewQueryLiquidationLoanRepayHistoryResponseWithDefaults instantiates a new QueryLiquidationLoanRepayHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryLiquidationLoanRepayHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryLiquidationLoanRepayHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryLiquidationLoanRepayHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryLiquidationLoanRepayHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *QueryLiquidationLoanRepayHistoryResponse) GetRows() []QueryLiquidationLoanRepayHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *QueryLiquidationLoanRepayHistoryResponse) GetRowsOk() (*[]QueryLiquidationLoanRepayHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *QueryLiquidationLoanRepayHistoryResponse) SetRows(v []QueryLiquidationLoanRepayHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *QueryLiquidationLoanRepayHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


