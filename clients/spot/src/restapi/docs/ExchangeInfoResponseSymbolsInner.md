# ExchangeInfoResponseSymbolsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int64** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**QuotePrecision** | Pointer to **int64** |  | [optional] 
**QuoteAssetPrecision** | Pointer to **int64** |  | [optional] 
**BaseCommissionPrecision** | Pointer to **int64** |  | [optional] 
**QuoteCommissionPrecision** | Pointer to **int64** |  | [optional] 
**OrderTypes** | Pointer to **[]string** |  | [optional] 
**IcebergAllowed** | Pointer to **bool** |  | [optional] 
**OcoAllowed** | Pointer to **bool** |  | [optional] 
**OtoAllowed** | Pointer to **bool** |  | [optional] 
**OpoAllowed** | Pointer to **bool** |  | [optional] 
**QuoteOrderQtyMarketAllowed** | Pointer to **bool** |  | [optional] 
**AllowTrailingStop** | Pointer to **bool** |  | [optional] 
**CancelReplaceAllowed** | Pointer to **bool** |  | [optional] 
**AmendAllowed** | Pointer to **bool** |  | [optional] 
**PegInstructionsAllowed** | Pointer to **bool** |  | [optional] 
**IsSpotTradingAllowed** | Pointer to **bool** |  | [optional] 
**IsMarginTradingAllowed** | Pointer to **bool** |  | [optional] 
**Filters** | Pointer to [**[]SymbolFilters**](SymbolFilters.md) |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**PermissionSets** | Pointer to **[][]string** |  | [optional] 
**DefaultSelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**AllowedSelfTradePreventionModes** | Pointer to **[]string** |  | [optional] 

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

### GetStatus

`func (o *ExchangeInfoResponseSymbolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangeInfoResponseSymbolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangeInfoResponseSymbolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExchangeInfoResponseSymbolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseAsset

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *ExchangeInfoResponseSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *ExchangeInfoResponseSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetPrecision() int64`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) SetBaseAssetPrecision(v int64)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *ExchangeInfoResponseSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *ExchangeInfoResponseSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *ExchangeInfoResponseSymbolsInner) GetQuotePrecision() int64`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetQuotePrecisionOk() (*int64, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *ExchangeInfoResponseSymbolsInner) SetQuotePrecision(v int64)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *ExchangeInfoResponseSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetQuoteAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetPrecision() int64`

GetQuoteAssetPrecision returns the QuoteAssetPrecision field if non-nil, zero value otherwise.

### GetQuoteAssetPrecisionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteAssetPrecisionOk() (*int64, bool)`

GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) SetQuoteAssetPrecision(v int64)`

SetQuoteAssetPrecision sets QuoteAssetPrecision field to given value.

### HasQuoteAssetPrecision

`func (o *ExchangeInfoResponseSymbolsInner) HasQuoteAssetPrecision() bool`

HasQuoteAssetPrecision returns a boolean if a field has been set.

### GetBaseCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseCommissionPrecision() int64`

GetBaseCommissionPrecision returns the BaseCommissionPrecision field if non-nil, zero value otherwise.

### GetBaseCommissionPrecisionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetBaseCommissionPrecisionOk() (*int64, bool)`

GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) SetBaseCommissionPrecision(v int64)`

SetBaseCommissionPrecision sets BaseCommissionPrecision field to given value.

### HasBaseCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) HasBaseCommissionPrecision() bool`

HasBaseCommissionPrecision returns a boolean if a field has been set.

### GetQuoteCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteCommissionPrecision() int64`

GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field if non-nil, zero value otherwise.

### GetQuoteCommissionPrecisionOk

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteCommissionPrecisionOk() (*int64, bool)`

GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) SetQuoteCommissionPrecision(v int64)`

SetQuoteCommissionPrecision sets QuoteCommissionPrecision field to given value.

### HasQuoteCommissionPrecision

`func (o *ExchangeInfoResponseSymbolsInner) HasQuoteCommissionPrecision() bool`

HasQuoteCommissionPrecision returns a boolean if a field has been set.

### GetOrderTypes

`func (o *ExchangeInfoResponseSymbolsInner) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *ExchangeInfoResponseSymbolsInner) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *ExchangeInfoResponseSymbolsInner) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *ExchangeInfoResponseSymbolsInner) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetIcebergAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetIcebergAllowed() bool`

GetIcebergAllowed returns the IcebergAllowed field if non-nil, zero value otherwise.

### GetIcebergAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetIcebergAllowedOk() (*bool, bool)`

GetIcebergAllowedOk returns a tuple with the IcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetIcebergAllowed(v bool)`

SetIcebergAllowed sets IcebergAllowed field to given value.

### HasIcebergAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasIcebergAllowed() bool`

HasIcebergAllowed returns a boolean if a field has been set.

### GetOcoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetOcoAllowed() bool`

GetOcoAllowed returns the OcoAllowed field if non-nil, zero value otherwise.

### GetOcoAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetOcoAllowedOk() (*bool, bool)`

GetOcoAllowedOk returns a tuple with the OcoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetOcoAllowed(v bool)`

SetOcoAllowed sets OcoAllowed field to given value.

### HasOcoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasOcoAllowed() bool`

HasOcoAllowed returns a boolean if a field has been set.

### GetOtoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetOtoAllowed() bool`

GetOtoAllowed returns the OtoAllowed field if non-nil, zero value otherwise.

### GetOtoAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetOtoAllowedOk() (*bool, bool)`

GetOtoAllowedOk returns a tuple with the OtoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetOtoAllowed(v bool)`

SetOtoAllowed sets OtoAllowed field to given value.

### HasOtoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasOtoAllowed() bool`

HasOtoAllowed returns a boolean if a field has been set.

### GetOpoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetOpoAllowed() bool`

GetOpoAllowed returns the OpoAllowed field if non-nil, zero value otherwise.

### GetOpoAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetOpoAllowedOk() (*bool, bool)`

GetOpoAllowedOk returns a tuple with the OpoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetOpoAllowed(v bool)`

SetOpoAllowed sets OpoAllowed field to given value.

### HasOpoAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasOpoAllowed() bool`

HasOpoAllowed returns a boolean if a field has been set.

### GetQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteOrderQtyMarketAllowed() bool`

GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field if non-nil, zero value otherwise.

### GetQuoteOrderQtyMarketAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool)`

GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetQuoteOrderQtyMarketAllowed(v bool)`

SetQuoteOrderQtyMarketAllowed sets QuoteOrderQtyMarketAllowed field to given value.

### HasQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasQuoteOrderQtyMarketAllowed() bool`

HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.

### GetAllowTrailingStop

`func (o *ExchangeInfoResponseSymbolsInner) GetAllowTrailingStop() bool`

GetAllowTrailingStop returns the AllowTrailingStop field if non-nil, zero value otherwise.

### GetAllowTrailingStopOk

`func (o *ExchangeInfoResponseSymbolsInner) GetAllowTrailingStopOk() (*bool, bool)`

GetAllowTrailingStopOk returns a tuple with the AllowTrailingStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrailingStop

`func (o *ExchangeInfoResponseSymbolsInner) SetAllowTrailingStop(v bool)`

SetAllowTrailingStop sets AllowTrailingStop field to given value.

### HasAllowTrailingStop

`func (o *ExchangeInfoResponseSymbolsInner) HasAllowTrailingStop() bool`

HasAllowTrailingStop returns a boolean if a field has been set.

### GetCancelReplaceAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetCancelReplaceAllowed() bool`

GetCancelReplaceAllowed returns the CancelReplaceAllowed field if non-nil, zero value otherwise.

### GetCancelReplaceAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetCancelReplaceAllowedOk() (*bool, bool)`

GetCancelReplaceAllowedOk returns a tuple with the CancelReplaceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReplaceAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetCancelReplaceAllowed(v bool)`

SetCancelReplaceAllowed sets CancelReplaceAllowed field to given value.

### HasCancelReplaceAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasCancelReplaceAllowed() bool`

HasCancelReplaceAllowed returns a boolean if a field has been set.

### GetAmendAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetAmendAllowed() bool`

GetAmendAllowed returns the AmendAllowed field if non-nil, zero value otherwise.

### GetAmendAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetAmendAllowedOk() (*bool, bool)`

GetAmendAllowedOk returns a tuple with the AmendAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetAmendAllowed(v bool)`

SetAmendAllowed sets AmendAllowed field to given value.

### HasAmendAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasAmendAllowed() bool`

HasAmendAllowed returns a boolean if a field has been set.

### GetPegInstructionsAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetPegInstructionsAllowed() bool`

GetPegInstructionsAllowed returns the PegInstructionsAllowed field if non-nil, zero value otherwise.

### GetPegInstructionsAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetPegInstructionsAllowedOk() (*bool, bool)`

GetPegInstructionsAllowedOk returns a tuple with the PegInstructionsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegInstructionsAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetPegInstructionsAllowed(v bool)`

SetPegInstructionsAllowed sets PegInstructionsAllowed field to given value.

### HasPegInstructionsAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasPegInstructionsAllowed() bool`

HasPegInstructionsAllowed returns a boolean if a field has been set.

### GetIsSpotTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetIsSpotTradingAllowed() bool`

GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field if non-nil, zero value otherwise.

### GetIsSpotTradingAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetIsSpotTradingAllowedOk() (*bool, bool)`

GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpotTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetIsSpotTradingAllowed(v bool)`

SetIsSpotTradingAllowed sets IsSpotTradingAllowed field to given value.

### HasIsSpotTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasIsSpotTradingAllowed() bool`

HasIsSpotTradingAllowed returns a boolean if a field has been set.

### GetIsMarginTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) GetIsMarginTradingAllowed() bool`

GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field if non-nil, zero value otherwise.

### GetIsMarginTradingAllowedOk

`func (o *ExchangeInfoResponseSymbolsInner) GetIsMarginTradingAllowedOk() (*bool, bool)`

GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) SetIsMarginTradingAllowed(v bool)`

SetIsMarginTradingAllowed sets IsMarginTradingAllowed field to given value.

### HasIsMarginTradingAllowed

`func (o *ExchangeInfoResponseSymbolsInner) HasIsMarginTradingAllowed() bool`

HasIsMarginTradingAllowed returns a boolean if a field has been set.

### GetFilters

`func (o *ExchangeInfoResponseSymbolsInner) GetFilters() []SymbolFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExchangeInfoResponseSymbolsInner) GetFiltersOk() (*[]SymbolFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExchangeInfoResponseSymbolsInner) SetFilters(v []SymbolFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ExchangeInfoResponseSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPermissions

`func (o *ExchangeInfoResponseSymbolsInner) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ExchangeInfoResponseSymbolsInner) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ExchangeInfoResponseSymbolsInner) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ExchangeInfoResponseSymbolsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPermissionSets

`func (o *ExchangeInfoResponseSymbolsInner) GetPermissionSets() [][]string`

GetPermissionSets returns the PermissionSets field if non-nil, zero value otherwise.

### GetPermissionSetsOk

`func (o *ExchangeInfoResponseSymbolsInner) GetPermissionSetsOk() (*[][]string, bool)`

GetPermissionSetsOk returns a tuple with the PermissionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSets

`func (o *ExchangeInfoResponseSymbolsInner) SetPermissionSets(v [][]string)`

SetPermissionSets sets PermissionSets field to given value.

### HasPermissionSets

`func (o *ExchangeInfoResponseSymbolsInner) HasPermissionSets() bool`

HasPermissionSets returns a boolean if a field has been set.

### GetDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseSymbolsInner) GetDefaultSelfTradePreventionMode() string`

GetDefaultSelfTradePreventionMode returns the DefaultSelfTradePreventionMode field if non-nil, zero value otherwise.

### GetDefaultSelfTradePreventionModeOk

`func (o *ExchangeInfoResponseSymbolsInner) GetDefaultSelfTradePreventionModeOk() (*string, bool)`

GetDefaultSelfTradePreventionModeOk returns a tuple with the DefaultSelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseSymbolsInner) SetDefaultSelfTradePreventionMode(v string)`

SetDefaultSelfTradePreventionMode sets DefaultSelfTradePreventionMode field to given value.

### HasDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseSymbolsInner) HasDefaultSelfTradePreventionMode() bool`

HasDefaultSelfTradePreventionMode returns a boolean if a field has been set.

### GetAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseSymbolsInner) GetAllowedSelfTradePreventionModes() []string`

GetAllowedSelfTradePreventionModes returns the AllowedSelfTradePreventionModes field if non-nil, zero value otherwise.

### GetAllowedSelfTradePreventionModesOk

`func (o *ExchangeInfoResponseSymbolsInner) GetAllowedSelfTradePreventionModesOk() (*[]string, bool)`

GetAllowedSelfTradePreventionModesOk returns a tuple with the AllowedSelfTradePreventionModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseSymbolsInner) SetAllowedSelfTradePreventionModes(v []string)`

SetAllowedSelfTradePreventionModes sets AllowedSelfTradePreventionModes field to given value.

### HasAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseSymbolsInner) HasAllowedSelfTradePreventionModes() bool`

HasAllowedSelfTradePreventionModes returns a boolean if a field has been set.


[[Back to README]](../README.md)


