# TickerResponse1Result

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**WeightedAvgPrice** | Pointer to **string** |  | [optional] 
**OpenPrice** | Pointer to **string** |  | [optional] 
**HighPrice** | Pointer to **string** |  | [optional] 
**LowPrice** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 
**QuoteVolume** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **int64** |  | [optional] 
**CloseTime** | Pointer to **int64** |  | [optional] 
**FirstId** | Pointer to **int64** |  | [optional] 
**LastId** | Pointer to **int64** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewTickerResponse1Result

`func NewTickerResponse1Result() *TickerResponse1Result`

NewTickerResponse1Result instantiates a new TickerResponse1Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerResponse1ResultWithDefaults

`func NewTickerResponse1ResultWithDefaults() *TickerResponse1Result`

NewTickerResponse1ResultWithDefaults instantiates a new TickerResponse1Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *TickerResponse1Result) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TickerResponse1Result) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TickerResponse1Result) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TickerResponse1Result) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPriceChange

`func (o *TickerResponse1Result) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *TickerResponse1Result) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *TickerResponse1Result) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *TickerResponse1Result) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *TickerResponse1Result) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *TickerResponse1Result) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *TickerResponse1Result) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *TickerResponse1Result) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetWeightedAvgPrice

`func (o *TickerResponse1Result) GetWeightedAvgPrice() string`

GetWeightedAvgPrice returns the WeightedAvgPrice field if non-nil, zero value otherwise.

### GetWeightedAvgPriceOk

`func (o *TickerResponse1Result) GetWeightedAvgPriceOk() (*string, bool)`

GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvgPrice

`func (o *TickerResponse1Result) SetWeightedAvgPrice(v string)`

SetWeightedAvgPrice sets WeightedAvgPrice field to given value.

### HasWeightedAvgPrice

`func (o *TickerResponse1Result) HasWeightedAvgPrice() bool`

HasWeightedAvgPrice returns a boolean if a field has been set.

### GetOpenPrice

`func (o *TickerResponse1Result) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *TickerResponse1Result) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *TickerResponse1Result) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *TickerResponse1Result) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *TickerResponse1Result) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *TickerResponse1Result) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *TickerResponse1Result) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *TickerResponse1Result) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *TickerResponse1Result) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *TickerResponse1Result) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *TickerResponse1Result) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *TickerResponse1Result) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *TickerResponse1Result) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *TickerResponse1Result) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *TickerResponse1Result) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *TickerResponse1Result) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetVolume

`func (o *TickerResponse1Result) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TickerResponse1Result) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TickerResponse1Result) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *TickerResponse1Result) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *TickerResponse1Result) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *TickerResponse1Result) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *TickerResponse1Result) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *TickerResponse1Result) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetOpenTime

`func (o *TickerResponse1Result) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *TickerResponse1Result) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *TickerResponse1Result) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *TickerResponse1Result) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *TickerResponse1Result) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *TickerResponse1Result) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *TickerResponse1Result) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *TickerResponse1Result) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFirstId

`func (o *TickerResponse1Result) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *TickerResponse1Result) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *TickerResponse1Result) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *TickerResponse1Result) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetLastId

`func (o *TickerResponse1Result) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *TickerResponse1Result) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *TickerResponse1Result) SetLastId(v int64)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *TickerResponse1Result) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetCount

`func (o *TickerResponse1Result) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TickerResponse1Result) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TickerResponse1Result) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TickerResponse1Result) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to README]](../README.md)


