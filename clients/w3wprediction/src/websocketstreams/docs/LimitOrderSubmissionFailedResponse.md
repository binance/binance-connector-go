# LimitOrderSubmissionFailedResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is &#x60;OrderHistory.orderId&#x60;. Use it for client-side deduplication. | [optional] 
**Topic** | Pointer to **string** | Market title from &#x60;MarketTopic.title&#x60;, truncated to 12 chars with &#x60;..&#x60; appended when longer (max 14 chars); empty string when the title is missing. | [optional] 

## Methods

### NewLimitOrderSubmissionFailedResponse

`func NewLimitOrderSubmissionFailedResponse() *LimitOrderSubmissionFailedResponse`

NewLimitOrderSubmissionFailedResponse instantiates a new LimitOrderSubmissionFailedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitOrderSubmissionFailedResponseWithDefaults

`func NewLimitOrderSubmissionFailedResponseWithDefaults() *LimitOrderSubmissionFailedResponse`

NewLimitOrderSubmissionFailedResponseWithDefaults instantiates a new LimitOrderSubmissionFailedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *LimitOrderSubmissionFailedResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *LimitOrderSubmissionFailedResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *LimitOrderSubmissionFailedResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *LimitOrderSubmissionFailedResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetTopic

`func (o *LimitOrderSubmissionFailedResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LimitOrderSubmissionFailedResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LimitOrderSubmissionFailedResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LimitOrderSubmissionFailedResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to README]](../README.md)


