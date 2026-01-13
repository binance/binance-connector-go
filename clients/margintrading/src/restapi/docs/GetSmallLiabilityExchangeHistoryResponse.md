# GetSmallLiabilityExchangeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Rows** | Pointer to [**[]GetSmallLiabilityExchangeHistoryResponseRowsInner**](GetSmallLiabilityExchangeHistoryResponseRowsInner.md) |  | [optional] 

## Methods

### NewGetSmallLiabilityExchangeHistoryResponse

`func NewGetSmallLiabilityExchangeHistoryResponse() *GetSmallLiabilityExchangeHistoryResponse`

NewGetSmallLiabilityExchangeHistoryResponse instantiates a new GetSmallLiabilityExchangeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSmallLiabilityExchangeHistoryResponseWithDefaults

`func NewGetSmallLiabilityExchangeHistoryResponseWithDefaults() *GetSmallLiabilityExchangeHistoryResponse`

NewGetSmallLiabilityExchangeHistoryResponseWithDefaults instantiates a new GetSmallLiabilityExchangeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetSmallLiabilityExchangeHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSmallLiabilityExchangeHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSmallLiabilityExchangeHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetSmallLiabilityExchangeHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *GetSmallLiabilityExchangeHistoryResponse) GetRows() []GetSmallLiabilityExchangeHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetSmallLiabilityExchangeHistoryResponse) GetRowsOk() (*[]GetSmallLiabilityExchangeHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetSmallLiabilityExchangeHistoryResponse) SetRows(v []GetSmallLiabilityExchangeHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetSmallLiabilityExchangeHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


