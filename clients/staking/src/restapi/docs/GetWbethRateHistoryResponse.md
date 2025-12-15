# GetWbethRateHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetWbethRateHistoryResponseRowsInner**](GetWbethRateHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetWbethRateHistoryResponse

`func NewGetWbethRateHistoryResponse() *GetWbethRateHistoryResponse`

NewGetWbethRateHistoryResponse instantiates a new GetWbethRateHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWbethRateHistoryResponseWithDefaults

`func NewGetWbethRateHistoryResponseWithDefaults() *GetWbethRateHistoryResponse`

NewGetWbethRateHistoryResponseWithDefaults instantiates a new GetWbethRateHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetWbethRateHistoryResponse) GetRows() []GetWbethRateHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetWbethRateHistoryResponse) GetRowsOk() (*[]GetWbethRateHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetWbethRateHistoryResponse) SetRows(v []GetWbethRateHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetWbethRateHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetWbethRateHistoryResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetWbethRateHistoryResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetWbethRateHistoryResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetWbethRateHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


