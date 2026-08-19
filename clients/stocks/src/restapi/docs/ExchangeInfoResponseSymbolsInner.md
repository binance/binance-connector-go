# ExchangeInfoResponseSymbolsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | US-equity ticker, e.g. &#x60;AAPL&#x60;. | [optional] 
**Tradability** | Pointer to **string** | Trading direction allowed — one of &#x60;BUY_SELL&#x60; / &#x60;BUY&#x60; / &#x60;SELL&#x60; / &#x60;NONE&#x60;. | [optional] 
**TradabilityUpdateTime** | Pointer to **int64** | Last time the &#x60;tradability&#x60; value was updated (ms epoch). | [optional] 
**OvernightSupported** | Pointer to **bool** | Whether the symbol supports overnight trading. | [optional] 
**Fractionable** | Pointer to **bool** | Whether fractional shares are supported during the regular session. | [optional] 
**FractionableEh** | Pointer to **bool** | Whether fractional shares are supported during extended hours. | [optional] 
**ExtendedSession** | Pointer to **bool** | Whether extended-session trading is enabled. | [optional] 
**MaxNumOrders** | Pointer to **int32** | Maximum number of open orders a user may have for this symbol. | [optional] 
**StepSize** | Pointer to **string** | Lot size — minimum increment for &#x60;quantity&#x60;. | [optional] 
**MultiplierUp** | Pointer to **string** | Upper price multiplier limit relative to reference. | [optional] 
**MultiplierDown** | Pointer to **string** | Lower price multiplier limit relative to reference. | [optional] 
**MinQty** | Pointer to **string** | Minimum allowed &#x60;quantity&#x60;. | [optional] 
**MaxQty** | Pointer to **string** | Maximum allowed &#x60;quantity&#x60;. | [optional] 
**MinNotional** | Pointer to **string** | Minimum order notional (USD). | [optional] 
**MaxNotional** | Pointer to **string** | Maximum order notional (USD). | [optional] 
**ListingTime** | Pointer to **int64** | Listing timestamp (ms epoch). | [optional] 
**DelistingTime** | Pointer to **NullableInt64** | Scheduled delisting timestamp (ms epoch); &#x60;null&#x60; if not scheduled for delisting. | [optional] 

## Methods

### NewExchangeInfoResponseSymbolsInner

`func NewExchangeInfoResponseSymbolsInner() *ExchangeInfoResponseSymbolsInner`

NewExchangeInfoResponseSymbolsInner instantiates a new ExchangeInfoResponseSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoResponseSymbolsInnerWithDefaults

`func NewExchangeInfoResponseSymbolsInnerWithDefaults() *ExchangeInfoResponseSymbolsInner`

NewExchangeInfoResponseSymbolsInnerWithDefaults instantiates a new ExchangeInfoResponseSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ExchangeInfoResponseSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ExchangeInfoResponseSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ExchangeInfoResponseSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ExchangeInfoResponseSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTradability

`func (o *ExchangeInfoResponseSymbolsInner) GetTradability() string`

GetTradability returns the Tradability field if non-nil, zero value otherwise.

### GetTradabilityOk

`func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityOk() (*string, bool)`

GetTradabilityOk returns a tuple with the Tradability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradability

`func (o *ExchangeInfoResponseSymbolsInner) SetTradability(v string)`

SetTradability sets Tradability field to given value.

### HasTradability

`func (o *ExchangeInfoResponseSymbolsInner) HasTradability() bool`

HasTradability returns a boolean if a field has been set.

### GetTradabilityUpdateTime

`func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityUpdateTime() int64`

GetTradabilityUpdateTime returns the TradabilityUpdateTime field if non-nil, zero value otherwise.

### GetTradabilityUpdateTimeOk

`func (o *ExchangeInfoResponseSymbolsInner) GetTradabilityUpdateTimeOk() (*int64, bool)`

GetTradabilityUpdateTimeOk returns a tuple with the TradabilityUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradabilityUpdateTime

`func (o *ExchangeInfoResponseSymbolsInner) SetTradabilityUpdateTime(v int64)`

SetTradabilityUpdateTime sets TradabilityUpdateTime field to given value.

### HasTradabilityUpdateTime

`func (o *ExchangeInfoResponseSymbolsInner) HasTradabilityUpdateTime() bool`

HasTradabilityUpdateTime returns a boolean if a field has been set.

### GetOvernightSupported

`func (o *ExchangeInfoResponseSymbolsInner) GetOvernightSupported() bool`

GetOvernightSupported returns the OvernightSupported field if non-nil, zero value otherwise.

### GetOvernightSupportedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetOvernightSupportedOk() (*bool, bool)`

GetOvernightSupportedOk returns a tuple with the OvernightSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvernightSupported

`func (o *ExchangeInfoResponseSymbolsInner) SetOvernightSupported(v bool)`

SetOvernightSupported sets OvernightSupported field to given value.

### HasOvernightSupported

`func (o *ExchangeInfoResponseSymbolsInner) HasOvernightSupported() bool`

HasOvernightSupported returns a boolean if a field has been set.

### GetFractionable

`func (o *ExchangeInfoResponseSymbolsInner) GetFractionable() bool`

GetFractionable returns the Fractionable field if non-nil, zero value otherwise.

### GetFractionableOk

`func (o *ExchangeInfoResponseSymbolsInner) GetFractionableOk() (*bool, bool)`

GetFractionableOk returns a tuple with the Fractionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionable

`func (o *ExchangeInfoResponseSymbolsInner) SetFractionable(v bool)`

SetFractionable sets Fractionable field to given value.

### HasFractionable

`func (o *ExchangeInfoResponseSymbolsInner) HasFractionable() bool`

HasFractionable returns a boolean if a field has been set.

### GetFractionableEh

`func (o *ExchangeInfoResponseSymbolsInner) GetFractionableEh() bool`

GetFractionableEh returns the FractionableEh field if non-nil, zero value otherwise.

### GetFractionableEhOk

`func (o *ExchangeInfoResponseSymbolsInner) GetFractionableEhOk() (*bool, bool)`

GetFractionableEhOk returns a tuple with the FractionableEh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFractionableEh

`func (o *ExchangeInfoResponseSymbolsInner) SetFractionableEh(v bool)`

SetFractionableEh sets FractionableEh field to given value.

### HasFractionableEh

`func (o *ExchangeInfoResponseSymbolsInner) HasFractionableEh() bool`

HasFractionableEh returns a boolean if a field has been set.

### GetExtendedSession

`func (o *ExchangeInfoResponseSymbolsInner) GetExtendedSession() bool`

GetExtendedSession returns the ExtendedSession field if non-nil, zero value otherwise.

### GetExtendedSessionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetExtendedSessionOk() (*bool, bool)`

GetExtendedSessionOk returns a tuple with the ExtendedSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSession

`func (o *ExchangeInfoResponseSymbolsInner) SetExtendedSession(v bool)`

SetExtendedSession sets ExtendedSession field to given value.

### HasExtendedSession

`func (o *ExchangeInfoResponseSymbolsInner) HasExtendedSession() bool`

HasExtendedSession returns a boolean if a field has been set.

### GetMaxNumOrders

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxNumOrders() int32`

GetMaxNumOrders returns the MaxNumOrders field if non-nil, zero value otherwise.

### GetMaxNumOrdersOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxNumOrdersOk() (*int32, bool)`

GetMaxNumOrdersOk returns a tuple with the MaxNumOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOrders

`func (o *ExchangeInfoResponseSymbolsInner) SetMaxNumOrders(v int32)`

SetMaxNumOrders sets MaxNumOrders field to given value.

### HasMaxNumOrders

`func (o *ExchangeInfoResponseSymbolsInner) HasMaxNumOrders() bool`

HasMaxNumOrders returns a boolean if a field has been set.

### GetStepSize

`func (o *ExchangeInfoResponseSymbolsInner) GetStepSize() string`

GetStepSize returns the StepSize field if non-nil, zero value otherwise.

### GetStepSizeOk

`func (o *ExchangeInfoResponseSymbolsInner) GetStepSizeOk() (*string, bool)`

GetStepSizeOk returns a tuple with the StepSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepSize

`func (o *ExchangeInfoResponseSymbolsInner) SetStepSize(v string)`

SetStepSize sets StepSize field to given value.

### HasStepSize

`func (o *ExchangeInfoResponseSymbolsInner) HasStepSize() bool`

HasStepSize returns a boolean if a field has been set.

### GetMultiplierUp

`func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierUp() string`

GetMultiplierUp returns the MultiplierUp field if non-nil, zero value otherwise.

### GetMultiplierUpOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierUpOk() (*string, bool)`

GetMultiplierUpOk returns a tuple with the MultiplierUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierUp

`func (o *ExchangeInfoResponseSymbolsInner) SetMultiplierUp(v string)`

SetMultiplierUp sets MultiplierUp field to given value.

### HasMultiplierUp

`func (o *ExchangeInfoResponseSymbolsInner) HasMultiplierUp() bool`

HasMultiplierUp returns a boolean if a field has been set.

### GetMultiplierDown

`func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierDown() string`

GetMultiplierDown returns the MultiplierDown field if non-nil, zero value otherwise.

### GetMultiplierDownOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMultiplierDownOk() (*string, bool)`

GetMultiplierDownOk returns a tuple with the MultiplierDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplierDown

`func (o *ExchangeInfoResponseSymbolsInner) SetMultiplierDown(v string)`

SetMultiplierDown sets MultiplierDown field to given value.

### HasMultiplierDown

`func (o *ExchangeInfoResponseSymbolsInner) HasMultiplierDown() bool`

HasMultiplierDown returns a boolean if a field has been set.

### GetMinQty

`func (o *ExchangeInfoResponseSymbolsInner) GetMinQty() string`

GetMinQty returns the MinQty field if non-nil, zero value otherwise.

### GetMinQtyOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMinQtyOk() (*string, bool)`

GetMinQtyOk returns a tuple with the MinQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinQty

`func (o *ExchangeInfoResponseSymbolsInner) SetMinQty(v string)`

SetMinQty sets MinQty field to given value.

### HasMinQty

`func (o *ExchangeInfoResponseSymbolsInner) HasMinQty() bool`

HasMinQty returns a boolean if a field has been set.

### GetMaxQty

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxQty() string`

GetMaxQty returns the MaxQty field if non-nil, zero value otherwise.

### GetMaxQtyOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxQtyOk() (*string, bool)`

GetMaxQtyOk returns a tuple with the MaxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQty

`func (o *ExchangeInfoResponseSymbolsInner) SetMaxQty(v string)`

SetMaxQty sets MaxQty field to given value.

### HasMaxQty

`func (o *ExchangeInfoResponseSymbolsInner) HasMaxQty() bool`

HasMaxQty returns a boolean if a field has been set.

### GetMinNotional

`func (o *ExchangeInfoResponseSymbolsInner) GetMinNotional() string`

GetMinNotional returns the MinNotional field if non-nil, zero value otherwise.

### GetMinNotionalOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMinNotionalOk() (*string, bool)`

GetMinNotionalOk returns a tuple with the MinNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNotional

`func (o *ExchangeInfoResponseSymbolsInner) SetMinNotional(v string)`

SetMinNotional sets MinNotional field to given value.

### HasMinNotional

`func (o *ExchangeInfoResponseSymbolsInner) HasMinNotional() bool`

HasMinNotional returns a boolean if a field has been set.

### GetMaxNotional

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxNotional() string`

GetMaxNotional returns the MaxNotional field if non-nil, zero value otherwise.

### GetMaxNotionalOk

`func (o *ExchangeInfoResponseSymbolsInner) GetMaxNotionalOk() (*string, bool)`

GetMaxNotionalOk returns a tuple with the MaxNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNotional

`func (o *ExchangeInfoResponseSymbolsInner) SetMaxNotional(v string)`

SetMaxNotional sets MaxNotional field to given value.

### HasMaxNotional

`func (o *ExchangeInfoResponseSymbolsInner) HasMaxNotional() bool`

HasMaxNotional returns a boolean if a field has been set.

### GetListingTime

`func (o *ExchangeInfoResponseSymbolsInner) GetListingTime() int64`

GetListingTime returns the ListingTime field if non-nil, zero value otherwise.

### GetListingTimeOk

`func (o *ExchangeInfoResponseSymbolsInner) GetListingTimeOk() (*int64, bool)`

GetListingTimeOk returns a tuple with the ListingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingTime

`func (o *ExchangeInfoResponseSymbolsInner) SetListingTime(v int64)`

SetListingTime sets ListingTime field to given value.

### HasListingTime

`func (o *ExchangeInfoResponseSymbolsInner) HasListingTime() bool`

HasListingTime returns a boolean if a field has been set.

### GetDelistingTime

`func (o *ExchangeInfoResponseSymbolsInner) GetDelistingTime() int64`

GetDelistingTime returns the DelistingTime field if non-nil, zero value otherwise.

### GetDelistingTimeOk

`func (o *ExchangeInfoResponseSymbolsInner) GetDelistingTimeOk() (*int64, bool)`

GetDelistingTimeOk returns a tuple with the DelistingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistingTime

`func (o *ExchangeInfoResponseSymbolsInner) SetDelistingTime(v int64)`

SetDelistingTime sets DelistingTime field to given value.

### HasDelistingTime

`func (o *ExchangeInfoResponseSymbolsInner) HasDelistingTime() bool`

HasDelistingTime returns a boolean if a field has been set.

### SetDelistingTimeNil

`func (o *ExchangeInfoResponseSymbolsInner) SetDelistingTimeNil(b bool)`

 SetDelistingTimeNil sets the value for DelistingTime to be an explicit nil

### UnsetDelistingTime
`func (o *ExchangeInfoResponseSymbolsInner) UnsetDelistingTime()`

UnsetDelistingTime ensures that no value is present for DelistingTime, not even an explicit nil

[[Back to README]](../README.md)


