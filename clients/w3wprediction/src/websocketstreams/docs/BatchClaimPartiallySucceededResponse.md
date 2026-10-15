# BatchClaimPartiallySucceededResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is the claim &#x60;batchId&#x60;. Use it for client-side deduplication. | [optional] 

## Methods

### NewBatchClaimPartiallySucceededResponse

`func NewBatchClaimPartiallySucceededResponse() *BatchClaimPartiallySucceededResponse`

NewBatchClaimPartiallySucceededResponse instantiates a new BatchClaimPartiallySucceededResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchClaimPartiallySucceededResponseWithDefaults

`func NewBatchClaimPartiallySucceededResponseWithDefaults() *BatchClaimPartiallySucceededResponse`

NewBatchClaimPartiallySucceededResponseWithDefaults instantiates a new BatchClaimPartiallySucceededResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *BatchClaimPartiallySucceededResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *BatchClaimPartiallySucceededResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *BatchClaimPartiallySucceededResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *BatchClaimPartiallySucceededResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.


[[Back to README]](../README.md)


