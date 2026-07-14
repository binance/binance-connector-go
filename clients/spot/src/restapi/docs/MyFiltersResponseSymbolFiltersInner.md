# MyFiltersResponseSymbolFiltersInner

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

### NewMyFiltersResponseSymbolFiltersInner

`func NewMyFiltersResponseSymbolFiltersInner() *MyFiltersResponseSymbolFiltersInner`

NewMyFiltersResponseSymbolFiltersInner instantiates a new MyFiltersResponseSymbolFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyFiltersResponseSymbolFiltersInnerWithDefaults

`func NewMyFiltersResponseSymbolFiltersInnerWithDefaults() *MyFiltersResponseSymbolFiltersInner`

NewMyFiltersResponseSymbolFiltersInnerWithDefaults instantiates a new MyFiltersResponseSymbolFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *MyFiltersResponseSymbolFiltersInner) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *MyFiltersResponseSymbolFiltersInner) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *MyFiltersResponseSymbolFiltersInner) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetPriceExponent

`func (o *MyFiltersResponseSymbolFiltersInner) GetPriceExponent() int32`

GetPriceExponent returns the PriceExponent field if non-nil, zero value otherwise.

### GetPriceExponentOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetPriceExponentOk() (*int32, bool)`

GetPriceExponentOk returns a tuple with the PriceExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExponent

`func (o *MyFiltersResponseSymbolFiltersInner) SetPriceExponent(v int32)`

SetPriceExponent sets PriceExponent field to given value.

### HasPriceExponent

`func (o *MyFiltersResponseSymbolFiltersInner) HasPriceExponent() bool`

HasPriceExponent returns a boolean if a field has been set.

### GetMinPrice

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *MyFiltersResponseSymbolFiltersInner) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *MyFiltersResponseSymbolFiltersInner) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### GetMaxPrice

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### GetTickSize

`func (o *MyFiltersResponseSymbolFiltersInner) GetTickSize() string`

GetTickSize returns the TickSize field if non-nil, zero value otherwise.

### GetTickSizeOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetTickSizeOk() (*string, bool)`

GetTickSizeOk returns a tuple with the TickSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickSize

`func (o *MyFiltersResponseSymbolFiltersInner) SetTickSize(v string)`

SetTickSize sets TickSize field to given value.

### HasTickSize

`func (o *MyFiltersResponseSymbolFiltersInner) HasTickSize() bool`

HasTickSize returns a boolean if a field has been set.

### GetMultiplierExponent

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierExponent() int32`

GetMultiplierExponent returns the MultiplierExponent field if non-nil, zero value otherwise.

### GetMultiplierExponentOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierExponentOk() (*int32, bool)`

GetMultiplierExponentOk returns a tuple with the MultiplierExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierExponent

`func (o *MyFiltersResponseSymbolFiltersInner) SetMultiplierExponent(v int32)`

SetMultiplierExponent sets MultiplierExponent field to given value.

### HasMultiplierExponent

`func (o *MyFiltersResponseSymbolFiltersInner) HasMultiplierExponent() bool`

HasMultiplierExponent returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetAvgPriceMins

`func (o *MyFiltersResponseSymbolFiltersInner) GetAvgPriceMins() int32`

GetAvgPriceMins returns the AvgPriceMins field if non-nil, zero value otherwise.

### GetAvgPriceMinsOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetAvgPriceMinsOk() (*int32, bool)`

GetAvgPriceMinsOk returns a tuple with the AvgPriceMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPriceMins

`func (o *MyFiltersResponseSymbolFiltersInner) SetAvgPriceMins(v int32)`

SetAvgPriceMins sets AvgPriceMins field to given value.

### HasAvgPriceMins

`func (o *MyFiltersResponseSymbolFiltersInner) HasAvgPriceMins() bool`

HasAvgPriceMins returns a boolean if a field has been set.

### GetBidMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) GetBidMultiplierUp() string`

GetBidMultiplierUp returns the BidMultiplierUp field if non-nil, zero value otherwise.

### GetBidMultiplierUpOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetBidMultiplierUpOk() (*string, bool)`

GetBidMultiplierUpOk returns a tuple with the BidMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) SetBidMultiplierUp(v string)`

SetBidMultiplierUp sets BidMultiplierUp field to given value.

### HasBidMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) HasBidMultiplierUp() bool`

HasBidMultiplierUp returns a boolean if a field has been set.

### GetBidMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) GetBidMultiplierDown() string`

GetBidMultiplierDown returns the BidMultiplierDown field if non-nil, zero value otherwise.

### GetBidMultiplierDownOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetBidMultiplierDownOk() (*string, bool)`

GetBidMultiplierDownOk returns a tuple with the BidMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) SetBidMultiplierDown(v string)`

SetBidMultiplierDown sets BidMultiplierDown field to given value.

### HasBidMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) HasBidMultiplierDown() bool`

HasBidMultiplierDown returns a boolean if a field has been set.

### GetAskMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) GetAskMultiplierUp() string`

GetAskMultiplierUp returns the AskMultiplierUp field if non-nil, zero value otherwise.

### GetAskMultiplierUpOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetAskMultiplierUpOk() (*string, bool)`

GetAskMultiplierUpOk returns a tuple with the AskMultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) SetAskMultiplierUp(v string)`

SetAskMultiplierUp sets AskMultiplierUp field to given value.

### HasAskMultiplierUp

`func (o *MyFiltersResponseSymbolFiltersInner) HasAskMultiplierUp() bool`

HasAskMultiplierUp returns a boolean if a field has been set.

### GetAskMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) GetAskMultiplierDown() string`

GetAskMultiplierDown returns the AskMultiplierDown field if non-nil, zero value otherwise.

### GetAskMultiplierDownOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetAskMultiplierDownOk() (*string, bool)`

GetAskMultiplierDownOk returns a tuple with the AskMultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) SetAskMultiplierDown(v string)`

SetAskMultiplierDown sets AskMultiplierDown field to given value.

### HasAskMultiplierDown

`func (o *MyFiltersResponseSymbolFiltersInner) HasAskMultiplierDown() bool`

HasAskMultiplierDown returns a boolean if a field has been set.

### GetQtyExponent

`func (o *MyFiltersResponseSymbolFiltersInner) GetQtyExponent() int32`

GetQtyExponent returns the QtyExponent field if non-nil, zero value otherwise.

### GetQtyExponentOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetQtyExponentOk() (*int32, bool)`

GetQtyExponentOk returns a tuple with the QtyExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtyExponent

`func (o *MyFiltersResponseSymbolFiltersInner) SetQtyExponent(v int32)`

SetQtyExponent sets QtyExponent field to given value.

### HasQtyExponent

`func (o *MyFiltersResponseSymbolFiltersInner) HasQtyExponent() bool`

HasQtyExponent returns a boolean if a field has been set.

### GetMinQty

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *MyFiltersResponseSymbolFiltersInner) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *MyFiltersResponseSymbolFiltersInner) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMaxQty

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetStepSize

`func (o *MyFiltersResponseSymbolFiltersInner) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *MyFiltersResponseSymbolFiltersInner) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *MyFiltersResponseSymbolFiltersInner) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetMinNotional

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *MyFiltersResponseSymbolFiltersInner) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *MyFiltersResponseSymbolFiltersInner) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetApplyToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyToMarket() bool`

GetApplyToMarket returns the ApplyToMarket field if non-nil, zero value otherwise.

### GetApplyToMarketOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyToMarketOk() (*bool, bool)`

GetApplyToMarketOk returns a tuple with the ApplyToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) SetApplyToMarket(v bool)`

SetApplyToMarket sets ApplyToMarket field to given value.

### HasApplyToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) HasApplyToMarket() bool`

HasApplyToMarket returns a boolean if a field has been set.

### GetApplyMinToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyMinToMarket() bool`

GetApplyMinToMarket returns the ApplyMinToMarket field if non-nil, zero value otherwise.

### GetApplyMinToMarketOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyMinToMarketOk() (*bool, bool)`

GetApplyMinToMarketOk returns a tuple with the ApplyMinToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMinToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) SetApplyMinToMarket(v bool)`

SetApplyMinToMarket sets ApplyMinToMarket field to given value.

### HasApplyMinToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) HasApplyMinToMarket() bool`

HasApplyMinToMarket returns a boolean if a field has been set.

### GetMaxNotional

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetApplyMaxToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyMaxToMarket() bool`

GetApplyMaxToMarket returns the ApplyMaxToMarket field if non-nil, zero value otherwise.

### GetApplyMaxToMarketOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetApplyMaxToMarketOk() (*bool, bool)`

GetApplyMaxToMarketOk returns a tuple with the ApplyMaxToMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyMaxToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) SetApplyMaxToMarket(v bool)`

SetApplyMaxToMarket sets ApplyMaxToMarket field to given value.

### HasApplyMaxToMarket

`func (o *MyFiltersResponseSymbolFiltersInner) HasApplyMaxToMarket() bool`

HasApplyMaxToMarket returns a boolean if a field has been set.

### GetLimit

`func (o *MyFiltersResponseSymbolFiltersInner) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MyFiltersResponseSymbolFiltersInner) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MyFiltersResponseSymbolFiltersInner) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrders() int64`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrdersOk() (*int64, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNumOrders(v int64)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetMaxNumAlgoOrders

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumAlgoOrders() int64`

GetMaxNumAlgoOrders returns the MaxNumAlgoOrders field if non-nil, zero value otherwise.

### GetMaxNumAlgoOrdersOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumAlgoOrdersOk() (*int64, bool)`

GetMaxNumAlgoOrdersOk returns a tuple with the MaxNumAlgoOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumAlgoOrders

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNumAlgoOrders(v int64)`

SetMaxNumAlgoOrders sets MaxNumAlgoOrders field to given value.

### HasMaxNumAlgoOrders

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNumAlgoOrders() bool`

HasMaxNumAlgoOrders returns a boolean if a field has been set.

### GetMaxNumIcebergOrders

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumIcebergOrders() int64`

GetMaxNumIcebergOrders returns the MaxNumIcebergOrders field if non-nil, zero value otherwise.

### GetMaxNumIcebergOrdersOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumIcebergOrdersOk() (*int64, bool)`

GetMaxNumIcebergOrdersOk returns a tuple with the MaxNumIcebergOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumIcebergOrders

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNumIcebergOrders(v int64)`

SetMaxNumIcebergOrders sets MaxNumIcebergOrders field to given value.

### HasMaxNumIcebergOrders

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNumIcebergOrders() bool`

HasMaxNumIcebergOrders returns a boolean if a field has been set.

### GetMaxPosition

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxPosition() string`

GetMaxPosition returns the MaxPosition field if non-nil, zero value otherwise.

### GetMaxPositionOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxPositionOk() (*string, bool)`

GetMaxPositionOk returns a tuple with the MaxPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPosition

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxPosition(v string)`

SetMaxPosition sets MaxPosition field to given value.

### HasMaxPosition

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxPosition() bool`

HasMaxPosition returns a boolean if a field has been set.

### GetMinTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinTrailingAboveDelta() int64`

GetMinTrailingAboveDelta returns the MinTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMinTrailingAboveDeltaOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinTrailingAboveDeltaOk() (*int64, bool)`

GetMinTrailingAboveDeltaOk returns a tuple with the MinTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) SetMinTrailingAboveDelta(v int64)`

SetMinTrailingAboveDelta sets MinTrailingAboveDelta field to given value.

### HasMinTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) HasMinTrailingAboveDelta() bool`

HasMinTrailingAboveDelta returns a boolean if a field has been set.

### GetMaxTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxTrailingAboveDelta() int64`

GetMaxTrailingAboveDelta returns the MaxTrailingAboveDelta field if non-nil, zero value otherwise.

### GetMaxTrailingAboveDeltaOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxTrailingAboveDeltaOk() (*int64, bool)`

GetMaxTrailingAboveDeltaOk returns a tuple with the MaxTrailingAboveDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxTrailingAboveDelta(v int64)`

SetMaxTrailingAboveDelta sets MaxTrailingAboveDelta field to given value.

### HasMaxTrailingAboveDelta

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxTrailingAboveDelta() bool`

HasMaxTrailingAboveDelta returns a boolean if a field has been set.

### GetMinTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinTrailingBelowDelta() int64`

GetMinTrailingBelowDelta returns the MinTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMinTrailingBelowDeltaOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMinTrailingBelowDeltaOk() (*int64, bool)`

GetMinTrailingBelowDeltaOk returns a tuple with the MinTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) SetMinTrailingBelowDelta(v int64)`

SetMinTrailingBelowDelta sets MinTrailingBelowDelta field to given value.

### HasMinTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) HasMinTrailingBelowDelta() bool`

HasMinTrailingBelowDelta returns a boolean if a field has been set.

### GetMaxTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxTrailingBelowDelta() int64`

GetMaxTrailingBelowDelta returns the MaxTrailingBelowDelta field if non-nil, zero value otherwise.

### GetMaxTrailingBelowDeltaOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxTrailingBelowDeltaOk() (*int64, bool)`

GetMaxTrailingBelowDeltaOk returns a tuple with the MaxTrailingBelowDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxTrailingBelowDelta(v int64)`

SetMaxTrailingBelowDelta sets MaxTrailingBelowDelta field to given value.

### HasMaxTrailingBelowDelta

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxTrailingBelowDelta() bool`

HasMaxTrailingBelowDelta returns a boolean if a field has been set.

### GetEndTime

`func (o *MyFiltersResponseSymbolFiltersInner) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *MyFiltersResponseSymbolFiltersInner) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *MyFiltersResponseSymbolFiltersInner) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMaxNumOrderLists

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrderLists() int64`

GetMaxNumOrderLists returns the MaxNumOrderLists field if non-nil, zero value otherwise.

### GetMaxNumOrderListsOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrderListsOk() (*int64, bool)`

GetMaxNumOrderListsOk returns a tuple with the MaxNumOrderLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrderLists

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNumOrderLists(v int64)`

SetMaxNumOrderLists sets MaxNumOrderLists field to given value.

### HasMaxNumOrderLists

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNumOrderLists() bool`

HasMaxNumOrderLists returns a boolean if a field has been set.

### GetMaxNumOrderAmends

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrderAmends() int64`

GetMaxNumOrderAmends returns the MaxNumOrderAmends field if non-nil, zero value otherwise.

### GetMaxNumOrderAmendsOk

`func (o *MyFiltersResponseSymbolFiltersInner) GetMaxNumOrderAmendsOk() (*int64, bool)`

GetMaxNumOrderAmendsOk returns a tuple with the MaxNumOrderAmends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrderAmends

`func (o *MyFiltersResponseSymbolFiltersInner) SetMaxNumOrderAmends(v int64)`

SetMaxNumOrderAmends sets MaxNumOrderAmends field to given value.

### HasMaxNumOrderAmends

`func (o *MyFiltersResponseSymbolFiltersInner) HasMaxNumOrderAmends() bool`

HasMaxNumOrderAmends returns a boolean if a field has been set.


[[Back to README]](../README.md)


