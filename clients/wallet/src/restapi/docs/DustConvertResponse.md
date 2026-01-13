# DustConvertResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TotalTransfered** | Pointer to **string** |  | [optional] 
**TotalServiceCharge** | Pointer to **string** |  | [optional] 
**TransferResult** | Pointer to [**[]DustConvertResponseTransferResultInner**](DustConvertResponseTransferResultInner.md) |  | [optional] 

## Methods

### NewDustConvertResponse

`func NewDustConvertResponse() *DustConvertResponse`

NewDustConvertResponse instantiates a new DustConvertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDustConvertResponseWithDefaults

`func NewDustConvertResponseWithDefaults() *DustConvertResponse`

NewDustConvertResponseWithDefaults instantiates a new DustConvertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalTransfered

`func (o *DustConvertResponse) GetTotalTransfered() string`

GetTotalTransfered returns the TotalTransfered field if non-nil, zero value otherwise.

### GetTotalTransferedOk

`func (o *DustConvertResponse) GetTotalTransferedOk() (*string, bool)`

GetTotalTransferedOk returns a tuple with the TotalTransfered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransfered

`func (o *DustConvertResponse) SetTotalTransfered(v string)`

SetTotalTransfered sets TotalTransfered field to given value.

### HasTotalTransfered

`func (o *DustConvertResponse) HasTotalTransfered() bool`

HasTotalTransfered returns a boolean if a field has been set.

### GetTotalServiceCharge

`func (o *DustConvertResponse) GetTotalServiceCharge() string`

GetTotalServiceCharge returns the TotalServiceCharge field if non-nil, zero value otherwise.

### GetTotalServiceChargeOk

`func (o *DustConvertResponse) GetTotalServiceChargeOk() (*string, bool)`

GetTotalServiceChargeOk returns a tuple with the TotalServiceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceCharge

`func (o *DustConvertResponse) SetTotalServiceCharge(v string)`

SetTotalServiceCharge sets TotalServiceCharge field to given value.

### HasTotalServiceCharge

`func (o *DustConvertResponse) HasTotalServiceCharge() bool`

HasTotalServiceCharge returns a boolean if a field has been set.

### GetTransferResult

`func (o *DustConvertResponse) GetTransferResult() []DustConvertResponseTransferResultInner`

GetTransferResult returns the TransferResult field if non-nil, zero value otherwise.

### GetTransferResultOk

`func (o *DustConvertResponse) GetTransferResultOk() (*[]DustConvertResponseTransferResultInner, bool)`

GetTransferResultOk returns a tuple with the TransferResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferResult

`func (o *DustConvertResponse) SetTransferResult(v []DustConvertResponseTransferResultInner)`

SetTransferResult sets TransferResult field to given value.

### HasTransferResult

`func (o *DustConvertResponse) HasTransferResult() bool`

HasTransferResult returns a boolean if a field has been set.


[[Back to README]](../README.md)


