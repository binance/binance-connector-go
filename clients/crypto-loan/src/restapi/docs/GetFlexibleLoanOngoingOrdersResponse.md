# GetFlexibleLoanOngoingOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetFlexibleLoanOngoingOrdersResponseRowsInner**](GetFlexibleLoanOngoingOrdersResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetFlexibleLoanOngoingOrdersResponse

`func NewGetFlexibleLoanOngoingOrdersResponse() *GetFlexibleLoanOngoingOrdersResponse`

NewGetFlexibleLoanOngoingOrdersResponse instantiates a new GetFlexibleLoanOngoingOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFlexibleLoanOngoingOrdersResponseWithDefaults

`func NewGetFlexibleLoanOngoingOrdersResponseWithDefaults() *GetFlexibleLoanOngoingOrdersResponse`

NewGetFlexibleLoanOngoingOrdersResponseWithDefaults instantiates a new GetFlexibleLoanOngoingOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetFlexibleLoanOngoingOrdersResponse) GetRows() []GetFlexibleLoanOngoingOrdersResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetFlexibleLoanOngoingOrdersResponse) GetRowsOk() (*[]GetFlexibleLoanOngoingOrdersResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetFlexibleLoanOngoingOrdersResponse) SetRows(v []GetFlexibleLoanOngoingOrdersResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetFlexibleLoanOngoingOrdersResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetFlexibleLoanOngoingOrdersResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFlexibleLoanOngoingOrdersResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFlexibleLoanOngoingOrdersResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFlexibleLoanOngoingOrdersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


