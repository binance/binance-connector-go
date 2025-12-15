# GetEthRedemptionHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Rows** | Pointer to [**[]GetEthRedemptionHistoryResponseRowsInner**](GetEthRedemptionHistoryResponseRowsInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetEthRedemptionHistoryResponse

`func NewGetEthRedemptionHistoryResponse() *GetEthRedemptionHistoryResponse`

NewGetEthRedemptionHistoryResponse instantiates a new GetEthRedemptionHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEthRedemptionHistoryResponseWithDefaults

`func NewGetEthRedemptionHistoryResponseWithDefaults() *GetEthRedemptionHistoryResponse`

NewGetEthRedemptionHistoryResponseWithDefaults instantiates a new GetEthRedemptionHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRows

`func (o *GetEthRedemptionHistoryResponse) GetRows() []GetEthRedemptionHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetEthRedemptionHistoryResponse) GetRowsOk() (*[]GetEthRedemptionHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetEthRedemptionHistoryResponse) SetRows(v []GetEthRedemptionHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetEthRedemptionHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.

### GetTotal

`func (o *GetEthRedemptionHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetEthRedemptionHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetEthRedemptionHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetEthRedemptionHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to README]](../README.md)


