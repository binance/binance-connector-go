# DustTransferResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalServiceCharge** | Pointer to **string** |  | [optional] 
**TotalTransfered** | Pointer to **string** |  | [optional] 
**TransferResult** | Pointer to [**[]DustTransferResponseTransferResultInner**](DustTransferResponseTransferResultInner.md) |  | [optional] 

## Methods

### NewDustTransferResponse

`func NewDustTransferResponse() *DustTransferResponse`

NewDustTransferResponse instantiates a new DustTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDustTransferResponseWithDefaults

`func NewDustTransferResponseWithDefaults() *DustTransferResponse`

NewDustTransferResponseWithDefaults instantiates a new DustTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalServiceCharge

`func (o *DustTransferResponse) GetTotalServiceCharge() string`

GetTotalServiceCharge returns the TotalServiceCharge field if non-nil, zero value otherwise.

### GetTotalServiceChargeOk

`func (o *DustTransferResponse) GetTotalServiceChargeOk() (*string, bool)`

GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceCharge

`func (o *DustTransferResponse) SetTotalServiceCharge(v string)`

SetTotalServiceCharge sets TotalServiceCharge field to given value.

### HasTotalServiceCharge

`func (o *DustTransferResponse) HasTotalServiceCharge() bool`

HasTotalServiceCharge returns a boolean if a field has been set.

### GetTotalTransfered

`func (o *DustTransferResponse) GetTotalTransfered() string`

GetTotalTransfered returns the TotalTransfered field if non-nil, zero value otherwise.

### GetTotalTransferedOk

`func (o *DustTransferResponse) GetTotalTransferedOk() (*string, bool)`

GetTotalTransferedOk returns a tuple with the TotalTransfered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransfered

`func (o *DustTransferResponse) SetTotalTransfered(v string)`

SetTotalTransfered sets TotalTransfered field to given value.

### HasTotalTransfered

`func (o *DustTransferResponse) HasTotalTransfered() bool`

HasTotalTransfered returns a boolean if a field has been set.

### GetTransferResult

`func (o *DustTransferResponse) GetTransferResult() []DustTransferResponseTransferResultInner`

GetTransferResult returns the TransferResult field if non-nil, zero value otherwise.

### GetTransferResultOk

`func (o *DustTransferResponse) GetTransferResultOk() (*[]DustTransferResponseTransferResultInner, bool)`

GetTransferResultOk returns a tuple with the TransferResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferResult

`func (o *DustTransferResponse) SetTransferResult(v []DustTransferResponseTransferResultInner)`

SetTransferResult sets TransferResult field to given value.

### HasTransferResult

`func (o *DustTransferResponse) HasTransferResult() bool`

HasTransferResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


