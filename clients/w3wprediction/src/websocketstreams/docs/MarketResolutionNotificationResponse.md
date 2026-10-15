# MarketResolutionNotificationResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is the &#x60;marketId&#x60;. Use it for client-side deduplication. | [optional] 
**Topic** | Pointer to **string** | Market title from &#x60;MarketTopic.title&#x60;, truncated to 12 chars with &#x60;..&#x60; appended when longer (max 14 chars); empty string when the title is missing. | [optional] 

## Methods

### NewMarketResolutionNotificationResponse

`func NewMarketResolutionNotificationResponse() *MarketResolutionNotificationResponse`

NewMarketResolutionNotificationResponse instantiates a new MarketResolutionNotificationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketResolutionNotificationResponseWithDefaults

`func NewMarketResolutionNotificationResponseWithDefaults() *MarketResolutionNotificationResponse`

NewMarketResolutionNotificationResponseWithDefaults instantiates a new MarketResolutionNotificationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *MarketResolutionNotificationResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *MarketResolutionNotificationResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *MarketResolutionNotificationResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *MarketResolutionNotificationResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetTopic

`func (o *MarketResolutionNotificationResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MarketResolutionNotificationResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MarketResolutionNotificationResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MarketResolutionNotificationResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to README]](../README.md)


