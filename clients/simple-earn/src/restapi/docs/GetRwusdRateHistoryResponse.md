# GetRwusdRateHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetBfusdRateHistoryResponseRowsInner**](GetBfusdRateHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **string** |  | [optional] 

## Methods

### NewGetRwusdRateHistoryResponse

`func NewGetRwusdRateHistoryResponse() *GetRwusdRateHistoryResponse`

NewGetRwusdRateHistoryResponse instantiates a new GetRwusdRateHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRwusdRateHistoryResponseWithDefaults

`func NewGetRwusdRateHistoryResponseWithDefaults() *GetRwusdRateHistoryResponse`

NewGetRwusdRateHistoryResponseWithDefaults instantiates a new GetRwusdRateHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetRwusdRateHistoryResponse) GetRows() []GetBfusdRateHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetRwusdRateHistoryResponse) GetRowsOk() (*[]GetBfusdRateHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetRwusdRateHistoryResponse) SetRows(v []GetBfusdRateHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetRwusdRateHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetRwusdRateHistoryResponse) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetRwusdRateHistoryResponse) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetRwusdRateHistoryResponse) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetRwusdRateHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


