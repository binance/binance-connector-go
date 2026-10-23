# LimitOrderFullyFilledResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique ID of this push. Use it to remove duplicates. | [optional] 
**Amount** | Pointer to **string** | Filled amount in USDT — paid for a buy, received after fees for a sell. Rounded down to 2 decimals; shown as &#x60;&lt;0.01&#x60; if smaller. | [optional] 
**Topic** | Pointer to **string** | Market title, cut to 12 characters with &#x60;..&#x60; added when longer; empty if unavailable. | [optional] 

## Methods

### NewLimitOrderFullyFilledResponse

`func NewLimitOrderFullyFilledResponse() *LimitOrderFullyFilledResponse`

NewLimitOrderFullyFilledResponse instantiates a new LimitOrderFullyFilledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitOrderFullyFilledResponseWithDefaults

`func NewLimitOrderFullyFilledResponseWithDefaults() *LimitOrderFullyFilledResponse`

NewLimitOrderFullyFilledResponseWithDefaults instantiates a new LimitOrderFullyFilledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *LimitOrderFullyFilledResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *LimitOrderFullyFilledResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *LimitOrderFullyFilledResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *LimitOrderFullyFilledResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *LimitOrderFullyFilledResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LimitOrderFullyFilledResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LimitOrderFullyFilledResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *LimitOrderFullyFilledResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTopic

`func (o *LimitOrderFullyFilledResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LimitOrderFullyFilledResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LimitOrderFullyFilledResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LimitOrderFullyFilledResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to README]](../README.md)


