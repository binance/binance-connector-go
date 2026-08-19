# PriceStreamResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | Event type, always &#x60;\&quot;price\&quot;&#x60;. | [optional] 
**Rates** | Pointer to [**[]PriceStreamResponseRatesInner**](PriceStreamResponseRatesInner.md) | One entry per symbol. | [optional] 

## Methods

### NewPriceStreamResponse

`func NewPriceStreamResponse() *PriceStreamResponse`

NewPriceStreamResponse instantiates a new PriceStreamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceStreamResponseWithDefaults

`func NewPriceStreamResponseWithDefaults() *PriceStreamResponse`

NewPriceStreamResponseWithDefaults instantiates a new PriceStreamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *PriceStreamResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *PriceStreamResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *PriceStreamResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *PriceStreamResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetRates

`func (o *PriceStreamResponse) GetRates() []PriceStreamResponseRatesInner`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *PriceStreamResponse) GetRatesOk() (*[]PriceStreamResponseRatesInner, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *PriceStreamResponse) SetRates(v []PriceStreamResponseRatesInner)`

SetRates sets Rates field to given value.

### HasRates

`func (o *PriceStreamResponse) HasRates() bool`

HasRates returns a boolean if a field has been set.


[[Back to README]](../README.md)


