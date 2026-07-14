# AllTokens24hTickerStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type | [optional] 
**D** | Pointer to [**[]AllTokens24hTickerStreamResponseDInner**](AllTokens24hTickerStreamResponseDInner.md) | Per-token 24-hour ticker metrics | [optional] 

## Methods

### NewAllTokens24hTickerStreamResponse

`func NewAllTokens24hTickerStreamResponse() *AllTokens24hTickerStreamResponse`

NewAllTokens24hTickerStreamResponse instantiates a new AllTokens24hTickerStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllTokens24hTickerStreamResponseWithDefaults

`func NewAllTokens24hTickerStreamResponseWithDefaults() *AllTokens24hTickerStreamResponse`

NewAllTokens24hTickerStreamResponseWithDefaults instantiates a new AllTokens24hTickerStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *AllTokens24hTickerStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AllTokens24hTickerStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AllTokens24hTickerStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AllTokens24hTickerStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetD

`func (o *AllTokens24hTickerStreamResponse) GetD() []AllTokens24hTickerStreamResponseDInner`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *AllTokens24hTickerStreamResponse) GetDOk() (*[]AllTokens24hTickerStreamResponseDInner, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *AllTokens24hTickerStreamResponse) SetD(v []AllTokens24hTickerStreamResponseDInner)`

SetD sets D field to given value.

### HasD

`func (o *AllTokens24hTickerStreamResponse) HasD() bool`

HasD returns a boolean if a field has been set.


[[Back to README]](../README.md)


