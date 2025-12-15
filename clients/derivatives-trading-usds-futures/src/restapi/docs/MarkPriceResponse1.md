# MarkPriceResponse1

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**MarkPrice** | Pointer to **string** |  | [optional] 
**IndexPrice** | Pointer to **string** |  | [optional] 
**EstimatedSettlePrice** | Pointer to **string** |  | [optional] 
**LastFundingRate** | Pointer to **string** |  | [optional] 
**InterestRate** | Pointer to **string** |  | [optional] 
**NextFundingTime** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewMarkPriceResponse1

`func NewMarkPriceResponse1() *MarkPriceResponse1`

NewMarkPriceResponse1 instantiates a new MarkPriceResponse1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkPriceResponse1WithDefaults

`func NewMarkPriceResponse1WithDefaults() *MarkPriceResponse1`

NewMarkPriceResponse1WithDefaults instantiates a new MarkPriceResponse1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarkPriceResponse1) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarkPriceResponse1) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarkPriceResponse1) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarkPriceResponse1) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMarkPrice

`func (o *MarkPriceResponse1) GetMarkPrice() string`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *MarkPriceResponse1) GetMarkPriceOk() (*string, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *MarkPriceResponse1) SetMarkPrice(v string)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *MarkPriceResponse1) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetIndexPrice

`func (o *MarkPriceResponse1) GetIndexPrice() string`

GetIndexPrice returns the IndexPrice field if non-nil, zero value otherwise.

### GetIndexPriceOk

`func (o *MarkPriceResponse1) GetIndexPriceOk() (*string, bool)`

GetIndexPriceOk returns a tuple with the IndexPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPrice

`func (o *MarkPriceResponse1) SetIndexPrice(v string)`

SetIndexPrice sets IndexPrice field to given value.

### HasIndexPrice

`func (o *MarkPriceResponse1) HasIndexPrice() bool`

HasIndexPrice returns a boolean if a field has been set.

### GetEstimatedSettlePrice

`func (o *MarkPriceResponse1) GetEstimatedSettlePrice() string`

GetEstimatedSettlePrice returns the EstimatedSettlePrice field if non-nil, zero value otherwise.

### GetEstimatedSettlePriceOk

`func (o *MarkPriceResponse1) GetEstimatedSettlePriceOk() (*string, bool)`

GetEstimatedSettlePriceOk returns a tuple with the EstimatedSettlePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSettlePrice

`func (o *MarkPriceResponse1) SetEstimatedSettlePrice(v string)`

SetEstimatedSettlePrice sets EstimatedSettlePrice field to given value.

### HasEstimatedSettlePrice

`func (o *MarkPriceResponse1) HasEstimatedSettlePrice() bool`

HasEstimatedSettlePrice returns a boolean if a field has been set.

### GetLastFundingRate

`func (o *MarkPriceResponse1) GetLastFundingRate() string`

GetLastFundingRate returns the LastFundingRate field if non-nil, zero value otherwise.

### GetLastFundingRateOk

`func (o *MarkPriceResponse1) GetLastFundingRateOk() (*string, bool)`

GetLastFundingRateOk returns a tuple with the LastFundingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFundingRate

`func (o *MarkPriceResponse1) SetLastFundingRate(v string)`

SetLastFundingRate sets LastFundingRate field to given value.

### HasLastFundingRate

`func (o *MarkPriceResponse1) HasLastFundingRate() bool`

HasLastFundingRate returns a boolean if a field has been set.

### GetInterestRate

`func (o *MarkPriceResponse1) GetInterestRate() string`

GetInterestRate returns the InterestRate field if non-nil, zero value otherwise.

### GetInterestRateOk

`func (o *MarkPriceResponse1) GetInterestRateOk() (*string, bool)`

GetInterestRateOk returns a tuple with the InterestRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRate

`func (o *MarkPriceResponse1) SetInterestRate(v string)`

SetInterestRate sets InterestRate field to given value.

### HasInterestRate

`func (o *MarkPriceResponse1) HasInterestRate() bool`

HasInterestRate returns a boolean if a field has been set.

### GetNextFundingTime

`func (o *MarkPriceResponse1) GetNextFundingTime() int64`

GetNextFundingTime returns the NextFundingTime field if non-nil, zero value otherwise.

### GetNextFundingTimeOk

`func (o *MarkPriceResponse1) GetNextFundingTimeOk() (*int64, bool)`

GetNextFundingTimeOk returns a tuple with the NextFundingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFundingTime

`func (o *MarkPriceResponse1) SetNextFundingTime(v int64)`

SetNextFundingTime sets NextFundingTime field to given value.

### HasNextFundingTime

`func (o *MarkPriceResponse1) HasNextFundingTime() bool`

HasNextFundingTime returns a boolean if a field has been set.

### GetTime

`func (o *MarkPriceResponse1) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MarkPriceResponse1) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MarkPriceResponse1) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *MarkPriceResponse1) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


