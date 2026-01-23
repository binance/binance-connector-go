# TickerResponseData

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**PriceChange** | Pointer to **string** |  | [optional] 
**PriceChangePercent** | Pointer to **string** |  | [optional] 
**WeightedAvgPrice** | Pointer to **string** |  | [optional] 
**LastPrice** | Pointer to **string** |  | [optional] 
**LastQty** | Pointer to **string** |  | [optional] 
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

### NewTickerResponseData

`func NewTickerResponseData() *TickerResponseData`

NewTickerResponseData instantiates a new TickerResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerResponseDataWithDefaults

`func NewTickerResponseDataWithDefaults() *TickerResponseData`

NewTickerResponseDataWithDefaults instantiates a new TickerResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *TickerResponseData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TickerResponseData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TickerResponseData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TickerResponseData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPriceChange

`func (o *TickerResponseData) GetPriceChange() string`

GetPriceChange returns the PriceChange field if non-nil, zero value otherwise.

### GetPriceChangeOk

`func (o *TickerResponseData) GetPriceChangeOk() (*string, bool)`

GetPriceChangeOk returns a tuple with the PriceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChange

`func (o *TickerResponseData) SetPriceChange(v string)`

SetPriceChange sets PriceChange field to given value.

### HasPriceChange

`func (o *TickerResponseData) HasPriceChange() bool`

HasPriceChange returns a boolean if a field has been set.

### GetPriceChangePercent

`func (o *TickerResponseData) GetPriceChangePercent() string`

GetPriceChangePercent returns the PriceChangePercent field if non-nil, zero value otherwise.

### GetPriceChangePercentOk

`func (o *TickerResponseData) GetPriceChangePercentOk() (*string, bool)`

GetPriceChangePercentOk returns a tuple with the PriceChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceChangePercent

`func (o *TickerResponseData) SetPriceChangePercent(v string)`

SetPriceChangePercent sets PriceChangePercent field to given value.

### HasPriceChangePercent

`func (o *TickerResponseData) HasPriceChangePercent() bool`

HasPriceChangePercent returns a boolean if a field has been set.

### GetWeightedAvgPrice

`func (o *TickerResponseData) GetWeightedAvgPrice() string`

GetWeightedAvgPrice returns the WeightedAvgPrice field if non-nil, zero value otherwise.

### GetWeightedAvgPriceOk

`func (o *TickerResponseData) GetWeightedAvgPriceOk() (*string, bool)`

GetWeightedAvgPriceOk returns a tuple with the WeightedAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightedAvgPrice

`func (o *TickerResponseData) SetWeightedAvgPrice(v string)`

SetWeightedAvgPrice sets WeightedAvgPrice field to given value.

### HasWeightedAvgPrice

`func (o *TickerResponseData) HasWeightedAvgPrice() bool`

HasWeightedAvgPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *TickerResponseData) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *TickerResponseData) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *TickerResponseData) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *TickerResponseData) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetLastQty

`func (o *TickerResponseData) GetLastQty() string`

GetLastQty returns the LastQty field if non-nil, zero value otherwise.

### GetLastQtyOk

`func (o *TickerResponseData) GetLastQtyOk() (*string, bool)`

GetLastQtyOk returns a tuple with the LastQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQty

`func (o *TickerResponseData) SetLastQty(v string)`

SetLastQty sets LastQty field to given value.

### HasLastQty

`func (o *TickerResponseData) HasLastQty() bool`

HasLastQty returns a boolean if a field has been set.

### GetOpenPrice

`func (o *TickerResponseData) GetOpenPrice() string`

GetOpenPrice returns the OpenPrice field if non-nil, zero value otherwise.

### GetOpenPriceOk

`func (o *TickerResponseData) GetOpenPriceOk() (*string, bool)`

GetOpenPriceOk returns a tuple with the OpenPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPrice

`func (o *TickerResponseData) SetOpenPrice(v string)`

SetOpenPrice sets OpenPrice field to given value.

### HasOpenPrice

`func (o *TickerResponseData) HasOpenPrice() bool`

HasOpenPrice returns a boolean if a field has been set.

### GetHighPrice

`func (o *TickerResponseData) GetHighPrice() string`

GetHighPrice returns the HighPrice field if non-nil, zero value otherwise.

### GetHighPriceOk

`func (o *TickerResponseData) GetHighPriceOk() (*string, bool)`

GetHighPriceOk returns a tuple with the HighPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighPrice

`func (o *TickerResponseData) SetHighPrice(v string)`

SetHighPrice sets HighPrice field to given value.

### HasHighPrice

`func (o *TickerResponseData) HasHighPrice() bool`

HasHighPrice returns a boolean if a field has been set.

### GetLowPrice

`func (o *TickerResponseData) GetLowPrice() string`

GetLowPrice returns the LowPrice field if non-nil, zero value otherwise.

### GetLowPriceOk

`func (o *TickerResponseData) GetLowPriceOk() (*string, bool)`

GetLowPriceOk returns a tuple with the LowPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPrice

`func (o *TickerResponseData) SetLowPrice(v string)`

SetLowPrice sets LowPrice field to given value.

### HasLowPrice

`func (o *TickerResponseData) HasLowPrice() bool`

HasLowPrice returns a boolean if a field has been set.

### GetVolume

`func (o *TickerResponseData) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *TickerResponseData) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *TickerResponseData) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *TickerResponseData) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetQuoteVolume

`func (o *TickerResponseData) GetQuoteVolume() string`

GetQuoteVolume returns the QuoteVolume field if non-nil, zero value otherwise.

### GetQuoteVolumeOk

`func (o *TickerResponseData) GetQuoteVolumeOk() (*string, bool)`

GetQuoteVolumeOk returns a tuple with the QuoteVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteVolume

`func (o *TickerResponseData) SetQuoteVolume(v string)`

SetQuoteVolume sets QuoteVolume field to given value.

### HasQuoteVolume

`func (o *TickerResponseData) HasQuoteVolume() bool`

HasQuoteVolume returns a boolean if a field has been set.

### GetOpenTime

`func (o *TickerResponseData) GetOpenTime() int64`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *TickerResponseData) GetOpenTimeOk() (*int64, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *TickerResponseData) SetOpenTime(v int64)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *TickerResponseData) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *TickerResponseData) GetCloseTime() int64`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *TickerResponseData) GetCloseTimeOk() (*int64, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *TickerResponseData) SetCloseTime(v int64)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *TickerResponseData) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetFirstId

`func (o *TickerResponseData) GetFirstId() int64`

GetFirstId returns the FirstId field if non-nil, zero value otherwise.

### GetFirstIdOk

`func (o *TickerResponseData) GetFirstIdOk() (*int64, bool)`

GetFirstIdOk returns a tuple with the FirstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstId

`func (o *TickerResponseData) SetFirstId(v int64)`

SetFirstId sets FirstId field to given value.

### HasFirstId

`func (o *TickerResponseData) HasFirstId() bool`

HasFirstId returns a boolean if a field has been set.

### GetLastId

`func (o *TickerResponseData) GetLastId() int64`

GetLastId returns the LastId field if non-nil, zero value otherwise.

### GetLastIdOk

`func (o *TickerResponseData) GetLastIdOk() (*int64, bool)`

GetLastIdOk returns a tuple with the LastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastId

`func (o *TickerResponseData) SetLastId(v int64)`

SetLastId sets LastId field to given value.

### HasLastId

`func (o *TickerResponseData) HasLastId() bool`

HasLastId returns a boolean if a field has been set.

### GetCount

`func (o *TickerResponseData) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TickerResponseData) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TickerResponseData) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TickerResponseData) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to README]](../README.md)


