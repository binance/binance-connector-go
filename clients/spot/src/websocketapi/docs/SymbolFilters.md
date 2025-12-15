# SymbolFilters

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | Pointer to **string** |  | [optional] 
**PriceExponent** | Pointer to **int32** |  | [optional] 
**MinPrice** | Pointer to **string** |  | [optional] 
**MaxPrice** | Pointer to **string** |  | [optional] 
**TickSize** | Pointer to **string** |  | [optional] 
**MultiplierExponent** | Pointer to **int32** |  | [optional] 
**MultiplierUp** | Pointer to **string** |  | [optional] 
**MultiplierDown** | Pointer to **string** |  | [optional] 
**AvgPriceMins** | Pointer to **int32** |  | [optional] 
**BidMultiplierUp** | Pointer to **string** |  | [optional] 
**BidMultiplierDown** | Pointer to **string** |  | [optional] 
**AskMultiplierUp** | Pointer to **string** |  | [optional] 
**AskMultiplierDown** | Pointer to **string** |  | [optional] 
**QtyExponent** | Pointer to **int32** |  | [optional] 
**MinQty** | Pointer to **string** |  | [optional] 
**MaxQty** | Pointer to **string** |  | [optional] 
**StepSize** | Pointer to **string** |  | [optional] 
**MinNotional** | Pointer to **string** |  | [optional] 
**ApplyToMarket** | Pointer to **bool** |  | [optional] 
**ApplyMinToMarket** | Pointer to **bool** |  | [optional] 
**MaxNotional** | Pointer to **string** |  | [optional] 
**ApplyMaxToMarket** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**MaxNumOrders** | Pointer to **int64** |  | [optional] 
**MaxNumAlgoOrders** | Pointer to **int64** |  | [optional] 
**MaxNumIcebergOrders** | Pointer to **int64** |  | [optional] 
**MaxPosition** | Pointer to **string** |  | [optional] 
**MinTrailingAboveDelta** | Pointer to **int64** |  | [optional] 
**MaxTrailingAboveDelta** | Pointer to **int64** |  | [optional] 
**MinTrailingBelowDelta** | Pointer to **int64** |  | [optional] 
**MaxTrailingBelowDelta** | Pointer to **int64** |  | [optional] 
**EndTime** | Pointer to **int64** |  | [optional] 
**MaxNumOrderLists** | Pointer to **int64** |  | [optional] 
**MaxNumOrderAmends** | Pointer to **int64** |  | [optional] 

## Methods

### NewSymbolFilters

`func NewSymbolFilters() *SymbolFilters`

NewSymbolFilters instantiates a new SymbolFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolFiltersWithDefaults

`func NewSymbolFiltersWithDefaults() *SymbolFilters`

NewSymbolFiltersWithDefaults instantiates a new SymbolFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *SymbolFilters) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *SymbolFilters) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *SymbolFilters) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *SymbolFilters) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetPriceExponent

`func (o *SymbolFilters) GetPriceExponent() int32`

GetPriceExponent returns the PriceExponent field if non-nil, zero value otherwise.

### GetPriceExponentOk

`func (o *SymbolFilters) GetPriceExponentOk() (*int32, bool)`

GetPriceExponentOk returns a tuple with the PriceExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExponent

`func (o *SymbolFilters) SetPriceExponent(v int32)`

SetPriceExponent sets PriceExponent field to given value.

### HasPriceExponent

`func (o *SymbolFilters) HasPriceExponent() bool`

HasPriceExponent returns a boolean if a field has been set.

### GetMinPrice

`func (o *SymbolFilters) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *SymbolFilters) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *SymbolFilters) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *SymbolFilters) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *SymbolFilters) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *SymbolFilters) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *SymbolFilters) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *SymbolFilters) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetTickSize

`func (o *SymbolFilters) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *SymbolFilters) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *SymbolFilters) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *SymbolFilters) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.

### GetMultiplierExponent

`func (o *SymbolFilters) GetMultiplierExponent() int32`

GetMultiplierExponent returns the MultiplierExponent field if non-nil, zero value otherwise.

### GetMultiplierExponentOk

`func (o *SymbolFilters) GetMultiplierExponentOk() (*int32, bool)`

GetMultiplierExponentOk returns a tuple with the MultiplierExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierExponent

`func (o *SymbolFilters) SetMultiplierExponent(v int32)`

SetMultiplierExponent sets MultiplierExponent field to given value.

### HasMultiplierExponent

`func (o *SymbolFilters) HasMultiplierExponent() bool`

HasMultiplierExponent returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *SymbolFilters) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *SymbolFilters) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *SymbolFilters) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *SymbolFilters) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *SymbolFilters) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *SymbolFilters) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *SymbolFilters) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *SymbolFilters) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *SymbolFilters) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *SymbolFilters) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *SymbolFilters) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *SymbolFilters) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.

### GetBidMultiplierUp

`func (o *SymbolFilters) GetBidMultiplierUp() string`

GetBidMultiplierUp returns the BidMultiplierUp field if non-nil, zero value otherwise.

### GetBidMultiplierUpOk

`func (o *SymbolFilters) GetBidMultiplierUpOk() (*string, bool)`

GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierUp

`func (o *SymbolFilters) SetBidMultiplierUp(v string)`

SetBidMultiplierUp sets BidMultiplierUp field to given value.

### HasBidMultiplierUp

`func (o *SymbolFilters) HasBidMultiplierUp() bool`

HasBidMultiplierUp returns a boolean if a field has been set.

### GetBidMultiplierDown

`func (o *SymbolFilters) GetBidMultiplierDown() string`

GetBidMultiplierDown returns the BidMultiplierDown field if non-nil, zero value otherwise.

### GetBidMultiplierDownOk

`func (o *SymbolFilters) GetBidMultiplierDownOk() (*string, bool)`

GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierDown

`func (o *SymbolFilters) SetBidMultiplierDown(v string)`

SetBidMultiplierDown sets BidMultiplierDown field to given value.

### HasBidMultiplierDown

`func (o *SymbolFilters) HasBidMultiplierDown() bool`

HasBidMultiplierDown returns a boolean if a field has been set.

### GetAskMultiplierUp

`func (o *SymbolFilters) GetAskMultiplierUp() string`

GetAskMultiplierUp returns the AskMultiplierUp field if non-nil, zero value otherwise.

### GetAskMultiplierUpOk

`func (o *SymbolFilters) GetAskMultiplierUpOk() (*string, bool)`

GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierUp

`func (o *SymbolFilters) SetAskMultiplierUp(v string)`

SetAskMultiplierUp sets AskMultiplierUp field to given value.

### HasAskMultiplierUp

`func (o *SymbolFilters) HasAskMultiplierUp() bool`

HasAskMultiplierUp returns a boolean if a field has been set.

### GetAskMultiplierDown

`func (o *SymbolFilters) GetAskMultiplierDown() string`

GetAskMultiplierDown returns the AskMultiplierDown field if non-nil, zero value otherwise.

### GetAskMultiplierDownOk

`func (o *SymbolFilters) GetAskMultiplierDownOk() (*string, bool)`

GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierDown

`func (o *SymbolFilters) SetAskMultiplierDown(v string)`

SetAskMultiplierDown sets AskMultiplierDown field to given value.

### HasAskMultiplierDown

`func (o *SymbolFilters) HasAskMultiplierDown() bool`

HasAskMultiplierDown returns a boolean if a field has been set.

### GetQtyExponent

`func (o *SymbolFilters) GetQtyExponent() int32`

GetQtyExponent returns the QtyExponent field if non-nil, zero value otherwise.

### GetQtyExponentOk

`func (o *SymbolFilters) GetQtyExponentOk() (*int32, bool)`

GetQtyExponentOk returns a tuple with the QtyExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyExponent

`func (o *SymbolFilters) SetQtyExponent(v int32)`

SetQtyExponent sets QtyExponent field to given value.

### HasQtyExponent

`func (o *SymbolFilters) HasQtyExponent() bool`

HasQtyExponent returns a boolean if a field has been set.

### GetMinQty

`func (o *SymbolFilters) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *SymbolFilters) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *SymbolFilters) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *SymbolFilters) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMaxQty

`func (o *SymbolFilters) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *SymbolFilters) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *SymbolFilters) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *SymbolFilters) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetStepSize

`func (o *SymbolFilters) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *SymbolFilters) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *SymbolFilters) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *SymbolFilters) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetMinNotional

`func (o *SymbolFilters) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *SymbolFilters) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *SymbolFilters) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *SymbolFilters) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetApplyToMarket

`func (o *SymbolFilters) GetApplyToMarket() bool`

GetApplyToMarket returns the ApplyToMarket field if non-nil, zero value otherwise.

### GetApplyToMarketOk

`func (o *SymbolFilters) GetApplyToMarketOk() (*bool, bool)`

GetApplyToMarketOk returns a tuple with the ApplyToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToMarket

`func (o *SymbolFilters) SetApplyToMarket(v bool)`

SetApplyToMarket sets ApplyToMarket field to given value.

### HasApplyToMarket

`func (o *SymbolFilters) HasApplyToMarket() bool`

HasApplyToMarket returns a boolean if a field has been set.

### GetApplyMinToMarket

`func (o *SymbolFilters) GetApplyMinToMarket() bool`

GetApplyMinToMarket returns the ApplyMinToMarket field if non-nil, zero value otherwise.

### GetApplyMinToMarketOk

`func (o *SymbolFilters) GetApplyMinToMarketOk() (*bool, bool)`

GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMinToMarket

`func (o *SymbolFilters) SetApplyMinToMarket(v bool)`

SetApplyMinToMarket sets ApplyMinToMarket field to given value.

### HasApplyMinToMarket

`func (o *SymbolFilters) HasApplyMinToMarket() bool`

HasApplyMinToMarket returns a boolean if a field has been set.

### GetMaxNotional

`func (o *SymbolFilters) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *SymbolFilters) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *SymbolFilters) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *SymbolFilters) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetApplyMaxToMarket

`func (o *SymbolFilters) GetApplyMaxToMarket() bool`

GetApplyMaxToMarket returns the ApplyMaxToMarket field if non-nil, zero value otherwise.

### GetApplyMaxToMarketOk

`func (o *SymbolFilters) GetApplyMaxToMarketOk() (*bool, bool)`

GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMaxToMarket

`func (o *SymbolFilters) SetApplyMaxToMarket(v bool)`

SetApplyMaxToMarket sets ApplyMaxToMarket field to given value.

### HasApplyMaxToMarket

`func (o *SymbolFilters) HasApplyMaxToMarket() bool`

HasApplyMaxToMarket returns a boolean if a field has been set.

### GetLimit

`func (o *SymbolFilters) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SymbolFilters) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SymbolFilters) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SymbolFilters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *SymbolFilters) GetMaxNumOrders() int64`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *SymbolFilters) GetMaxNumOrdersOk() (*int64, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *SymbolFilters) SetMaxNumOrders(v int64)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *SymbolFilters) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetMaxNumAlgoOrders

`func (o *SymbolFilters) GetMaxNumAlgoOrders() int64`

GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field if non-nil, zero value otherwise.

### GetMaxNumAlgoOrdersOk

`func (o *SymbolFilters) GetMaxNumAlgoOrdersOk() (*int64, bool)`

GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumAlgoOrders

`func (o *SymbolFilters) SetMaxNumAlgoOrders(v int64)`

SetMaxNumAlgoOrders sets MaxNumAlgoOrders field to given value.

### HasMaxNumAlgoOrders

`func (o *SymbolFilters) HasMaxNumAlgoOrders() bool`

HasMaxNumAlgoOrders returns a boolean if a field has been set.

### GetMaxNumIcebergOrders

`func (o *SymbolFilters) GetMaxNumIcebergOrders() int64`

GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field if non-nil, zero value otherwise.

### GetMaxNumIcebergOrdersOk

`func (o *SymbolFilters) GetMaxNumIcebergOrdersOk() (*int64, bool)`

GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumIcebergOrders

`func (o *SymbolFilters) SetMaxNumIcebergOrders(v int64)`

SetMaxNumIcebergOrders sets MaxNumIcebergOrders field to given value.

### HasMaxNumIcebergOrders

`func (o *SymbolFilters) HasMaxNumIcebergOrders() bool`

HasMaxNumIcebergOrders returns a boolean if a field has been set.

### GetMaxPosition

`func (o *SymbolFilters) GetMaxPosition() string`

GetMaxPosition returns the MaxPosition field if non-nil, zero value otherwise.

### GetMaxPositionOk

`func (o *SymbolFilters) GetMaxPositionOk() (*string, bool)`

GetMaxPositionOk returns a tuple with the MaxPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPosition

`func (o *SymbolFilters) SetMaxPosition(v string)`

SetMaxPosition sets MaxPosition field to given value.

### HasMaxPosition

`func (o *SymbolFilters) HasMaxPosition() bool`

HasMaxPosition returns a boolean if a field has been set.

### GetMinTrailingAboveDelta

`func (o *SymbolFilters) GetMinTrailingAboveDelta() int64`

GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMinTrailingAboveDeltaOk

`func (o *SymbolFilters) GetMinTrailingAboveDeltaOk() (*int64, bool)`

GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingAboveDelta

`func (o *SymbolFilters) SetMinTrailingAboveDelta(v int64)`

SetMinTrailingAboveDelta sets MinTrailingAboveDelta field to given value.

### HasMinTrailingAboveDelta

`func (o *SymbolFilters) HasMinTrailingAboveDelta() bool`

HasMinTrailingAboveDelta returns a boolean if a field has been set.

### GetMaxTrailingAboveDelta

`func (o *SymbolFilters) GetMaxTrailingAboveDelta() int64`

GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMaxTrailingAboveDeltaOk

`func (o *SymbolFilters) GetMaxTrailingAboveDeltaOk() (*int64, bool)`

GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingAboveDelta

`func (o *SymbolFilters) SetMaxTrailingAboveDelta(v int64)`

SetMaxTrailingAboveDelta sets MaxTrailingAboveDelta field to given value.

### HasMaxTrailingAboveDelta

`func (o *SymbolFilters) HasMaxTrailingAboveDelta() bool`

HasMaxTrailingAboveDelta returns a boolean if a field has been set.

### GetMinTrailingBelowDelta

`func (o *SymbolFilters) GetMinTrailingBelowDelta() int64`

GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMinTrailingBelowDeltaOk

`func (o *SymbolFilters) GetMinTrailingBelowDeltaOk() (*int64, bool)`

GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingBelowDelta

`func (o *SymbolFilters) SetMinTrailingBelowDelta(v int64)`

SetMinTrailingBelowDelta sets MinTrailingBelowDelta field to given value.

### HasMinTrailingBelowDelta

`func (o *SymbolFilters) HasMinTrailingBelowDelta() bool`

HasMinTrailingBelowDelta returns a boolean if a field has been set.

### GetMaxTrailingBelowDelta

`func (o *SymbolFilters) GetMaxTrailingBelowDelta() int64`

GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMaxTrailingBelowDeltaOk

`func (o *SymbolFilters) GetMaxTrailingBelowDeltaOk() (*int64, bool)`

GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingBelowDelta

`func (o *SymbolFilters) SetMaxTrailingBelowDelta(v int64)`

SetMaxTrailingBelowDelta sets MaxTrailingBelowDelta field to given value.

### HasMaxTrailingBelowDelta

`func (o *SymbolFilters) HasMaxTrailingBelowDelta() bool`

HasMaxTrailingBelowDelta returns a boolean if a field has been set.

### GetEndTime

`func (o *SymbolFilters) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SymbolFilters) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SymbolFilters) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SymbolFilters) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMaxNumOrderLists

`func (o *SymbolFilters) GetMaxNumOrderLists() int64`

GetMaxNumOrderLists returns the MaxNumOrderLists field if non-nil, zero value otherwise.

### GetMaxNumOrderListsOk

`func (o *SymbolFilters) GetMaxNumOrderListsOk() (*int64, bool)`

GetMaxNumOrderListsOk returns a tuple with the MaxNumOrderLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrderLists

`func (o *SymbolFilters) SetMaxNumOrderLists(v int64)`

SetMaxNumOrderLists sets MaxNumOrderLists field to given value.

### HasMaxNumOrderLists

`func (o *SymbolFilters) HasMaxNumOrderLists() bool`

HasMaxNumOrderLists returns a boolean if a field has been set.

### GetMaxNumOrderAmends

`func (o *SymbolFilters) GetMaxNumOrderAmends() int64`

GetMaxNumOrderAmends returns the MaxNumOrderAmends field if non-nil, zero value otherwise.

### GetMaxNumOrderAmendsOk

`func (o *SymbolFilters) GetMaxNumOrderAmendsOk() (*int64, bool)`

GetMaxNumOrderAmendsOk returns a tuple with the MaxNumOrderAmends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrderAmends

`func (o *SymbolFilters) SetMaxNumOrderAmends(v int64)`

SetMaxNumOrderAmends sets MaxNumOrderAmends field to given value.

### HasMaxNumOrderAmends

`func (o *SymbolFilters) HasMaxNumOrderAmends() bool`

HasMaxNumOrderAmends returns a boolean if a field has been set.


[[Back to README]](../README.md)


