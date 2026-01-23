# StartUserDataStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ListenKey** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int64** |  | [optional] 

## Methods

### NewStartUserDataStreamResponse

`func NewStartUserDataStreamResponse() *StartUserDataStreamResponse`

NewStartUserDataStreamResponse instantiates a new StartUserDataStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartUserDataStreamResponseWithDefaults

`func NewStartUserDataStreamResponseWithDefaults() *StartUserDataStreamResponse`

NewStartUserDataStreamResponseWithDefaults instantiates a new StartUserDataStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenKey

`func (o *StartUserDataStreamResponse) GetListenKey() string`

GetListenKey returns the ListenKey field if non-nil, zero value otherwise.

### GetListenKeyOk

`func (o *StartUserDataStreamResponse) GetListenKeyOk() (*string, bool)`

GetListenKeyOk returns a tuple with the ListenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenKey

`func (o *StartUserDataStreamResponse) SetListenKey(v string)`

SetListenKey sets ListenKey field to given value.

### HasListenKey

`func (o *StartUserDataStreamResponse) HasListenKey() bool`

HasListenKey returns a boolean if a field has been set.

### GetExpiration

`func (o *StartUserDataStreamResponse) GetExpiration() int64`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *StartUserDataStreamResponse) GetExpirationOk() (*int64, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *StartUserDataStreamResponse) SetExpiration(v int64)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *StartUserDataStreamResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to README]](../README.md)


