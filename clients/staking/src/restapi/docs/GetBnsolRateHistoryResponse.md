# GetBnsolRateHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetBnsolRateHistoryResponseRowsInner**](GetBnsolRateHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetBnsolRateHistoryResponse

`func NewGetBnsolRateHistoryResponse() *GetBnsolRateHistoryResponse`

NewGetBnsolRateHistoryResponse instantiates a new GetBnsolRateHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBnsolRateHistoryResponseWithDefaults

`func NewGetBnsolRateHistoryResponseWithDefaults() *GetBnsolRateHistoryResponse`

NewGetBnsolRateHistoryResponseWithDefaults instantiates a new GetBnsolRateHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetBnsolRateHistoryResponse) GetRows() []GetBnsolRateHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetBnsolRateHistoryResponse) GetRowsOk() (*[]GetBnsolRateHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetBnsolRateHistoryResponse) SetRows(v []GetBnsolRateHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetBnsolRateHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetBnsolRateHistoryResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetBnsolRateHistoryResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetBnsolRateHistoryResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetBnsolRateHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


