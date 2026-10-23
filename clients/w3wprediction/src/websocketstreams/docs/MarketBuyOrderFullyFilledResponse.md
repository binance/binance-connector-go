# MarketBuyOrderFullyFilledResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**PushId** | Pointer to **string** | Unique ID of this push. Use it to remove duplicates. | [optional] 
**Amount** | Pointer to **string** | USDT the user paid. Rounded down to 2 decimals; shown as &#x60;&lt;0.01&#x60; if smaller. | [optional] 
**Topic** | Pointer to **string** | Market title, cut to 12 characters with &#x60;..&#x60; added when longer; empty if unavailable. | [optional] 

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


