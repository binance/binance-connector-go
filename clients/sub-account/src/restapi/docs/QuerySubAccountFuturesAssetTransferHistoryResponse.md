# QuerySubAccountFuturesAssetTransferHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**FuturesType** | Pointer to **int64** |  | [optional] 
**Transfers** | Pointer to [**[]QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner**](QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner.md) |  | [optional] 

## Methods

### NewQuerySubAccountFuturesAssetTransferHistoryResponse

`func NewQuerySubAccountFuturesAssetTransferHistoryResponse() *QuerySubAccountFuturesAssetTransferHistoryResponse`

NewQuerySubAccountFuturesAssetTransferHistoryResponse instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubAccountFuturesAssetTransferHistoryResponseWithDefaults

`func NewQuerySubAccountFuturesAssetTransferHistoryResponseWithDefaults() *QuerySubAccountFuturesAssetTransferHistoryResponse`

NewQuerySubAccountFuturesAssetTransferHistoryResponseWithDefaults instantiates a new QuerySubAccountFuturesAssetTransferHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFuturesType

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetFuturesType() int64`

GetFuturesType returns the FuturesType field if non-nil, zero value otherwise.

### GetFuturesTypeOk

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetFuturesTypeOk() (*int64, bool)`

GetFuturesTypeOk returns a tuple with the FuturesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturesType

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetFuturesType(v int64)`

SetFuturesType sets FuturesType field to given value.

### HasFuturesType

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasFuturesType() bool`

HasFuturesType returns a boolean if a field has been set.

### GetTransfers

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetTransfers() []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) GetTransfersOk() (*[]QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) SetTransfers(v []QuerySubAccountFuturesAssetTransferHistoryResponseTransfersInner)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *QuerySubAccountFuturesAssetTransferHistoryResponse) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.


[[Back to README]](../README.md)


