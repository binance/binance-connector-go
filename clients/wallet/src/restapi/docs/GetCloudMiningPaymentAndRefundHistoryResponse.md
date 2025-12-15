# GetCloudMiningPaymentAndRefundHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Rows** | Pointer to [**[]GetCloudMiningPaymentAndRefundHistoryResponseRowsInner**](GetCloudMiningPaymentAndRefundHistoryResponseRowsInner.md) |  | [optional] 

## Methods

### NewGetCloudMiningPaymentAndRefundHistoryResponse

`func NewGetCloudMiningPaymentAndRefundHistoryResponse() *GetCloudMiningPaymentAndRefundHistoryResponse`

NewGetCloudMiningPaymentAndRefundHistoryResponse instantiates a new GetCloudMiningPaymentAndRefundHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudMiningPaymentAndRefundHistoryResponseWithDefaults

`func NewGetCloudMiningPaymentAndRefundHistoryResponseWithDefaults() *GetCloudMiningPaymentAndRefundHistoryResponse`

NewGetCloudMiningPaymentAndRefundHistoryResponseWithDefaults instantiates a new GetCloudMiningPaymentAndRefundHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetRows

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetRows() []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) GetRowsOk() (*[]GetCloudMiningPaymentAndRefundHistoryResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) SetRows(v []GetCloudMiningPaymentAndRefundHistoryResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetCloudMiningPaymentAndRefundHistoryResponse) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to README]](../README.md)


