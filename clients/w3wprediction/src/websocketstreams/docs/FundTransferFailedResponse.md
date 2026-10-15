# FundTransferFailedResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is the &#x60;transferId&#x60;. Use it for client-side deduplication. | [optional] 
**Amount** | Pointer to **string** | Original transfer amount, from &#x60;FundTransferOrder.fromTokenAmount&#x60;. Truncated to 2 decimals (towards zero); rendered as &#x60;&lt;0.01&#x60; when the value is below 0.01. | [optional] 

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


