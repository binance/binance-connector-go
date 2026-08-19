# QuoteStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;quote\&quot;&#x60;. | [optional] 
**E** | Pointer to **int64** | Event time — epoch milliseconds when the server pushed the message. | [optional] 
**S** | Pointer to **string** | Symbol (UPPERCASE ticker), e.g. &#x60;\&quot;AAPL\&quot;&#x60;. | [optional] 
**Bp** | Pointer to **string** | Best bid price. | [optional] 
**Ap** | Pointer to **string** | Best ask price. | [optional] 
**Bs** | Pointer to **int32** | Best bid size (shares). | [optional] 
**As** | Pointer to **int32** | Best ask size (shares). | [optional] 
**T** | Pointer to **NullableInt64** | Source quote timestamp (epoch milliseconds); may be null. | [optional] 

## Methods

### NewQuoteStreamResponse

`func NewQuoteStreamResponse() *QuoteStreamResponse`

NewQuoteStreamResponse instantiates a new QuoteStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteStreamResponseWithDefaults

`func NewQuoteStreamResponseWithDefaults() *QuoteStreamResponse`

NewQuoteStreamResponseWithDefaults instantiates a new QuoteStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *QuoteStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *QuoteStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *QuoteStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *QuoteStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetE

`func (o *QuoteStreamResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *QuoteStreamResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *QuoteStreamResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *QuoteStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetS

`func (o *QuoteStreamResponse) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *QuoteStreamResponse) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *QuoteStreamResponse) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *QuoteStreamResponse) HasS() bool`

HasS returns a boolean if a field has been set.

### GetBp

`func (o *QuoteStreamResponse) GetBp() string`

GetBp returns the Bp field if non-nil, zero value otherwise.

### GetBpOk

`func (o *QuoteStreamResponse) GetBpOk() (*string, bool)`

GetBpOk returns a tuple with the Bp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBp

`func (o *QuoteStreamResponse) SetBp(v string)`

SetBp sets Bp field to given value.

### HasBp

`func (o *QuoteStreamResponse) HasBp() bool`

HasBp returns a boolean if a field has been set.

### GetAp

`func (o *QuoteStreamResponse) GetAp() string`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *QuoteStreamResponse) GetApOk() (*string, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *QuoteStreamResponse) SetAp(v string)`

SetAp sets Ap field to given value.

### HasAp

`func (o *QuoteStreamResponse) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetBs

`func (o *QuoteStreamResponse) GetBs() int32`

GetBs returns the Bs field if non-nil, zero value otherwise.

### GetBsOk

`func (o *QuoteStreamResponse) GetBsOk() (*int32, bool)`

GetBsOk returns a tuple with the Bs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBs

`func (o *QuoteStreamResponse) SetBs(v int32)`

SetBs sets Bs field to given value.

### HasBs

`func (o *QuoteStreamResponse) HasBs() bool`

HasBs returns a boolean if a field has been set.

### GetAs

`func (o *QuoteStreamResponse) GetAs() int32`

GetAs returns the As field if non-nil, zero value otherwise.

### GetAsOk

`func (o *QuoteStreamResponse) GetAsOk() (*int32, bool)`

GetAsOk returns a tuple with the As field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAs

`func (o *QuoteStreamResponse) SetAs(v int32)`

SetAs sets As field to given value.

### HasAs

`func (o *QuoteStreamResponse) HasAs() bool`

HasAs returns a boolean if a field has been set.

### GetT

`func (o *QuoteStreamResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *QuoteStreamResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *QuoteStreamResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *QuoteStreamResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### SetTNil

`func (o *QuoteStreamResponse) SetTNil(b bool)`

 SetTNil sets the value for T to be an explicit nil

### UnsetT
`func (o *QuoteStreamResponse) UnsetT()`

UnsetT ensures that no value is present for T, not even an explicit nil

[[Back to README]](../README.md)


