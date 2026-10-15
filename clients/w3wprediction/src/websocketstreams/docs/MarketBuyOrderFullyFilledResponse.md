# MarketBuyOrderFullyFilledResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique push ID in the form &#x60;pm_&lt;refId&gt;_&lt;scenarioCode&gt;_&lt;8hex&gt;&#x60;, present in every topic; &#x60;refId&#x60; is &#x60;OrderHistory.orderId&#x60;. Use it for client-side deduplication. | [optional] 
**Amount** | Pointer to **string** | USDT amount the user paid, from &#x60;OrderHistory.amountIn&#x60;. Truncated to 2 decimals (towards zero); rendered as &#x60;&lt;0.01&#x60; when the value is below 0.01. | [optional] 
**Topic** | Pointer to **string** | Market title from &#x60;MarketTopic.title&#x60;, truncated to 12 chars with &#x60;..&#x60; appended when longer (max 14 chars); empty string when the title is missing. | [optional] 

## Methods

### NewMarketBuyOrderFullyFilledResponse

`func NewMarketBuyOrderFullyFilledResponse() *MarketBuyOrderFullyFilledResponse`

NewMarketBuyOrderFullyFilledResponse instantiates a new MarketBuyOrderFullyFilledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketBuyOrderFullyFilledResponseWithDefaults

`func NewMarketBuyOrderFullyFilledResponseWithDefaults() *MarketBuyOrderFullyFilledResponse`

NewMarketBuyOrderFullyFilledResponseWithDefaults instantiates a new MarketBuyOrderFullyFilledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushId

`func (o *MarketBuyOrderFullyFilledResponse) GetPushId() string`

GetPushId returns the PushId field if non-nil, zero value otherwise.

### GetPushIdOk

`func (o *MarketBuyOrderFullyFilledResponse) GetPushIdOk() (*string, bool)`

GetPushIdOk returns a tuple with the PushId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushId

`func (o *MarketBuyOrderFullyFilledResponse) SetPushId(v string)`

SetPushId sets PushId field to given value.

### HasPushId

`func (o *MarketBuyOrderFullyFilledResponse) HasPushId() bool`

HasPushId returns a boolean if a field has been set.

### GetAmount

`func (o *MarketBuyOrderFullyFilledResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *MarketBuyOrderFullyFilledResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *MarketBuyOrderFullyFilledResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *MarketBuyOrderFullyFilledResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTopic

`func (o *MarketBuyOrderFullyFilledResponse) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MarketBuyOrderFullyFilledResponse) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MarketBuyOrderFullyFilledResponse) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MarketBuyOrderFullyFilledResponse) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to README]](../README.md)


