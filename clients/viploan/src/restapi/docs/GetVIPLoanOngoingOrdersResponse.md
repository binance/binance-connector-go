# GetVIPLoanOngoingOrdersResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetVIPLoanOngoingOrdersResponseRowsInner**](GetVIPLoanOngoingOrdersResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetVIPLoanOngoingOrdersResponse

`func NewGetVIPLoanOngoingOrdersResponse() *GetVIPLoanOngoingOrdersResponse`

NewGetVIPLoanOngoingOrdersResponse instantiates a new GetVIPLoanOngoingOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVIPLoanOngoingOrdersResponseWithDefaults

`func NewGetVIPLoanOngoingOrdersResponseWithDefaults() *GetVIPLoanOngoingOrdersResponse`

NewGetVIPLoanOngoingOrdersResponseWithDefaults instantiates a new GetVIPLoanOngoingOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetVIPLoanOngoingOrdersResponse) GetRows() []GetVIPLoanOngoingOrdersResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetVIPLoanOngoingOrdersResponse) GetRowsOk() (*[]GetVIPLoanOngoingOrdersResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetVIPLoanOngoingOrdersResponse) SetRows(v []GetVIPLoanOngoingOrdersResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetVIPLoanOngoingOrdersResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetVIPLoanOngoingOrdersResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetVIPLoanOngoingOrdersResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetVIPLoanOngoingOrdersResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetVIPLoanOngoingOrdersResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


