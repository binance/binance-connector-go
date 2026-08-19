# LatestQuoteResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | US-equity ticker. | [optional] 
**BidPrice** | Pointer to **string** | Best bid price (USD). | [optional] 
**AskPrice** | Pointer to **string** | Best ask price (USD). | [optional] 
**BidSize** | Pointer to **int32** | Best bid size (shares). | [optional] 
**AskSize** | Pointer to **int32** | Best ask size (shares). | [optional] 

## Methods

### NewLatestQuoteResponse

`func NewLatestQuoteResponse() *LatestQuoteResponse`

NewLatestQuoteResponse instantiates a new LatestQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLatestQuoteResponseWithDefaults

`func NewLatestQuoteResponseWithDefaults() *LatestQuoteResponse`

NewLatestQuoteResponseWithDefaults instantiates a new LatestQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *LatestQuoteResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *LatestQuoteResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *LatestQuoteResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *LatestQuoteResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetBidPrice

`func (o *LatestQuoteResponse) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *LatestQuoteResponse) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *LatestQuoteResponse) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *LatestQuoteResponse) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetAskPrice

`func (o *LatestQuoteResponse) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *LatestQuoteResponse) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *LatestQuoteResponse) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *LatestQuoteResponse) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetBidSize

`func (o *LatestQuoteResponse) GetBidSize() int32`

GetBidSize returns the BidSize field if non-nil, zero value otherwise.

### GetBidSizeOk

`func (o *LatestQuoteResponse) GetBidSizeOk() (*int32, bool)`

GetBidSizeOk returns a tuple with the BidSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidSize

`func (o *LatestQuoteResponse) SetBidSize(v int32)`

SetBidSize sets BidSize field to given value.

### HasBidSize

`func (o *LatestQuoteResponse) HasBidSize() bool`

HasBidSize returns a boolean if a field has been set.

### GetAskSize

`func (o *LatestQuoteResponse) GetAskSize() int32`

GetAskSize returns the AskSize field if non-nil, zero value otherwise.

### GetAskSizeOk

`func (o *LatestQuoteResponse) GetAskSizeOk() (*int32, bool)`

GetAskSizeOk returns a tuple with the AskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskSize

`func (o *LatestQuoteResponse) SetAskSize(v int32)`

SetAskSize sets AskSize field to given value.

### HasAskSize

`func (o *LatestQuoteResponse) HasAskSize() bool`

HasAskSize returns a boolean if a field has been set.


[[Back to README]](../README.md)


