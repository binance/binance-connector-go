# FundTransferSucceededResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique ID of this push. Use it to remove duplicates. | [optional] 
**Amount** | Pointer to **string** | Amount transferred. Rounded down to 2 decimals; shown as &#x60;&lt;0.01&#x60; if smaller. | [optional] 

## Methods

### NewFundTransferSucceededResponse

`func NewFundTransferSucceededResponse() *FundTransferSucceededResponse`

NewFundTransferSucceededResponse instantiates a new FundTransferSucceededResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFundTransferSucceededResponseWithDefaults

`func NewFundTransferSucceededResponseWithDefaults() *FundTransferSucceededResponse`

NewFundTransferSucceededResponseWithDefaults instantiates a new FundTransferSucceededResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *FundTransferSucceededResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *FundTransferSucceededResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *FundTransferSucceededResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *FundTransferSucceededResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *FundTransferSucceededResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *FundTransferSucceededResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *FundTransferSucceededResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *FundTransferSucceededResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to README]](../README.md)


