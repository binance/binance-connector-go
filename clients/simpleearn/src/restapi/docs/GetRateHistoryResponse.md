# GetRateHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetRateHistoryResponseRowsInner**](GetRateHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRateHistoryResponse

`func NewGetRateHistoryResponse() *GetRateHistoryResponse`

NewGetRateHistoryResponse instantiates a new GetRateHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRateHistoryResponseWithDefaults

`func NewGetRateHistoryResponseWithDefaults() *GetRateHistoryResponse`

NewGetRateHistoryResponseWithDefaults instantiates a new GetRateHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetRateHistoryResponse) GetRows() []GetRateHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetRateHistoryResponse) GetRowsOk() (*[]GetRateHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetRateHistoryResponse) SetRows(v []GetRateHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetRateHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetRateHistoryResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRateHistoryResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRateHistoryResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetRateHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


