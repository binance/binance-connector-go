# Ticker24hrResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**WeightedAvgPrice** | Pointer to **string** |  | [optional] 
**PrevClosePrice** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**LastQty** | Pointer to **string** |  | [optional] 
**BidPrice** | Pointer to **string** |  | [optional] 
**BidQty** | Pointer to **string** |  | [optional] 
**AskPrice** | Pointer to **string** |  | [optional] 
**AskQty** | Pointer to **string** |  | [optional] 
**OpenPrice** | Pointer to **string** |  | [optional] 
**HighPrice** | Pointer to **string** |  | [optional] 
**LowPrice** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**QuoteVolume** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **int64** |  | [optional] 
**CloseTime** | Pointer to **int64** |  | [optional] 
**FirstId** | Pointer to **int64** |  | [optional] 
**LastId** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewTicker24hrResponse

`func NewTicker24hrResponse() *Ticker24hrResponse`

NewTicker24hrResponse instantiates a new Ticker24hrResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicker24hrResponseWithDefaults

`func NewTicker24hrResponseWithDefaults() *Ticker24hrResponse`

NewTicker24hrResponseWithDefaults instantiates a new Ticker24hrResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Ticker24hrResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ticker24hrResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ticker24hrResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Ticker24hrResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPriceChange

`func (o *Ticker24hrResponse) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *Ticker24hrResponse) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *Ticker24hrResponse) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *Ticker24hrResponse) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *Ticker24hrResponse) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *Ticker24hrResponse) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *Ticker24hrResponse) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *Ticker24hrResponse) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetWeightedAvgPrice

`func (o *Ticker24hrResponse) GetWeightedAvgPrice() string`

GetWeightedAvgPrice returns the WeightedAvgPrice field if non-nil, zero value otherwise.

### GetWeightedAvgPriceOk

`func (o *Ticker24hrResponse) GetWeightedAvgPriceOk() (*string, bool)`

GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvgPrice

`func (o *Ticker24hrResponse) SetWeightedAvgPrice(v string)`

SetWeightedAvgPrice sets WeightedAvgPrice field to given value.

### HasWeightedAvgPrice

`func (o *Ticker24hrResponse) HasWeightedAvgPrice() bool`

HasWeightedAvgPrice returns a boolean if a field has been set.

### GetPrevClosePrice

`func (o *Ticker24hrResponse) GetPrevClosePrice() string`

GetPrevClosePrice returns the PrevClosePrice field if non-nil, zero value otherwise.

### GetPrevClosePriceOk

`func (o *Ticker24hrResponse) GetPrevClosePriceOk() (*string, bool)`

GetPrevClosePriceOk returns a tuple with the PrevClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevClosePrice

`func (o *Ticker24hrResponse) SetPrevClosePrice(v string)`

SetPrevClosePrice sets PrevClosePrice field to given value.

### HasPrevClosePrice

`func (o *Ticker24hrResponse) HasPrevClosePrice() bool`

HasPrevClosePrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *Ticker24hrResponse) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *Ticker24hrResponse) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *Ticker24hrResponse) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *Ticker24hrResponse) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastQty

`func (o *Ticker24hrResponse) GetLastQty() string`

GetLastQty returns the LastQty field if non-nil, zero value otherwise.

### GetLastQtyOk

`func (o *Ticker24hrResponse) GetLastQtyOk() (*string, bool)`

GetLastQtyOk returns a tuple with the LastQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQty

`func (o *Ticker24hrResponse) SetLastQty(v string)`

SetLastQty sets LastQty field to given value.

### HasLastQty

`func (o *Ticker24hrResponse) HasLastQty() bool`

HasLastQty returns a boolean if a field has been set.

### GetBidPrice

`func (o *Ticker24hrResponse) GetBidPrice() string`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *Ticker24hrResponse) GetBidPriceOk() (*string, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *Ticker24hrResponse) SetBidPrice(v string)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *Ticker24hrResponse) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetBidQty

`func (o *Ticker24hrResponse) GetBidQty() string`

GetBidQty returns the BidQty field if non-nil, zero value otherwise.

### GetBidQtyOk

`func (o *Ticker24hrResponse) GetBidQtyOk() (*string, bool)`

GetBidQtyOk returns a tuple with the BidQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidQty

`func (o *Ticker24hrResponse) SetBidQty(v string)`

SetBidQty sets BidQty field to given value.

### HasBidQty

`func (o *Ticker24hrResponse) HasBidQty() bool`

HasBidQty returns a boolean if a field has been set.

### GetAskPrice

`func (o *Ticker24hrResponse) GetAskPrice() string`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *Ticker24hrResponse) GetAskPriceOk() (*string, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *Ticker24hrResponse) SetAskPrice(v string)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *Ticker24hrResponse) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetAskQty

`func (o *Ticker24hrResponse) GetAskQty() string`

GetAskQty returns the AskQty field if non-nil, zero value otherwise.

### GetAskQtyOk

`func (o *Ticker24hrResponse) GetAskQtyOk() (*string, bool)`

GetAskQtyOk returns a tuple with the AskQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskQty

`func (o *Ticker24hrResponse) SetAskQty(v string)`

SetAskQty sets AskQty field to given value.

### HasAskQty

`func (o *Ticker24hrResponse) HasAskQty() bool`

HasAskQty returns a boolean if a field has been set.

### GetOpenPrice

`func (o *Ticker24hrResponse) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *Ticker24hrResponse) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *Ticker24hrResponse) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *Ticker24hrResponse) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *Ticker24hrResponse) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *Ticker24hrResponse) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *Ticker24hrResponse) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *Ticker24hrResponse) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *Ticker24hrResponse) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *Ticker24hrResponse) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *Ticker24hrResponse) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *Ticker24hrResponse) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetVolume

`func (o *Ticker24hrResponse) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Ticker24hrResponse) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Ticker24hrResponse) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Ticker24hrResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *Ticker24hrResponse) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *Ticker24hrResponse) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *Ticker24hrResponse) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *Ticker24hrResponse) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetOpenTime

`func (o *Ticker24hrResponse) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *Ticker24hrResponse) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *Ticker24hrResponse) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *Ticker24hrResponse) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *Ticker24hrResponse) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *Ticker24hrResponse) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *Ticker24hrResponse) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *Ticker24hrResponse) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFirstId

`func (o *Ticker24hrResponse) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *Ticker24hrResponse) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *Ticker24hrResponse) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *Ticker24hrResponse) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetLastId

`func (o *Ticker24hrResponse) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *Ticker24hrResponse) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *Ticker24hrResponse) SetLastId(v int64)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *Ticker24hrResponse) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetCount

`func (o *Ticker24hrResponse) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Ticker24hrResponse) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Ticker24hrResponse) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Ticker24hrResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to README]](../README.md)


