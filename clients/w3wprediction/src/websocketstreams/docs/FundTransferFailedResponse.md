# FundTransferFailedResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique ID of this push. Use it to remove duplicates. | [optional] 
**Amount** | Pointer to **string** | Amount of the attempted transfer. Rounded down to 2 decimals; shown as &#x60;&lt;0.01&#x60; if smaller. | [optional] 

## Methods

### NewFundTransferFailedResponse

`func NewFundTransferFailedResponse() *FundTransferFailedResponse`

NewFundTransferFailedResponse instantiates a new FundTransferFailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundTransferFailedResponseWithDefaults

`func NewFundTransferFailedResponseWithDefaults() *FundTransferFailedResponse`

NewFundTransferFailedResponseWithDefaults instantiates a new FundTransferFailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *FundTransferFailedResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *FundTransferFailedResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *FundTransferFailedResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *FundTransferFailedResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *FundTransferFailedResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundTransferFailedResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundTransferFailedResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FundTransferFailedResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to README]](../README.md)


