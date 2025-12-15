# ExchangeInfoResponseResultSymbolsInner

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

### NewExchangeInfoResponseResultSymbolsInner

`func NewExchangeInfoResponseResultSymbolsInner() *ExchangeInfoResponseResultSymbolsInner`

NewExchangeInfoResponseResultSymbolsInner instantiates a new ExchangeInfoResponseResultSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInfoResponseResultSymbolsInnerWithDefaults

`func NewExchangeInfoResponseResultSymbolsInnerWithDefaults() *ExchangeInfoResponseResultSymbolsInner`

NewExchangeInfoResponseResultSymbolsInnerWithDefaults instantiates a new ExchangeInfoResponseResultSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ExchangeInfoResponseResultSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ExchangeInfoResponseResultSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ExchangeInfoResponseResultSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetStatus

`func (o *ExchangeInfoResponseResultSymbolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExchangeInfoResponseResultSymbolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExchangeInfoResponseResultSymbolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBaseAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetPrecision() int64`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseAssetPrecision(v int64)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuotePrecision() int64`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuotePrecisionOk() (*int64, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) SetQuotePrecision(v int64)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetQuoteAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetPrecision() int64`

GetQuoteAssetPrecision returns the QuoteAssetPrecision field if non-nil, zero value otherwise.

### GetQuoteAssetPrecisionOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteAssetPrecisionOk() (*int64, bool)`

GetQuoteAssetPrecisionOk returns a tuple with the QuoteAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteAssetPrecision(v int64)`

SetQuoteAssetPrecision sets QuoteAssetPrecision field to given value.

### HasQuoteAssetPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteAssetPrecision() bool`

HasQuoteAssetPrecision returns a boolean if a field has been set.

### GetBaseCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseCommissionPrecision() int64`

GetBaseCommissionPrecision returns the BaseCommissionPrecision field if non-nil, zero value otherwise.

### GetBaseCommissionPrecisionOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetBaseCommissionPrecisionOk() (*int64, bool)`

GetBaseCommissionPrecisionOk returns a tuple with the BaseCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) SetBaseCommissionPrecision(v int64)`

SetBaseCommissionPrecision sets BaseCommissionPrecision field to given value.

### HasBaseCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) HasBaseCommissionPrecision() bool`

HasBaseCommissionPrecision returns a boolean if a field has been set.

### GetQuoteCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteCommissionPrecision() int64`

GetQuoteCommissionPrecision returns the QuoteCommissionPrecision field if non-nil, zero value otherwise.

### GetQuoteCommissionPrecisionOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteCommissionPrecisionOk() (*int64, bool)`

GetQuoteCommissionPrecisionOk returns a tuple with the QuoteCommissionPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteCommissionPrecision(v int64)`

SetQuoteCommissionPrecision sets QuoteCommissionPrecision field to given value.

### HasQuoteCommissionPrecision

`func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteCommissionPrecision() bool`

HasQuoteCommissionPrecision returns a boolean if a field has been set.

### GetOrderTypes

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOrderTypes() []string`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOrderTypesOk() (*[]string, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *ExchangeInfoResponseResultSymbolsInner) SetOrderTypes(v []string)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *ExchangeInfoResponseResultSymbolsInner) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetIcebergAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIcebergAllowed() bool`

GetIcebergAllowed returns the IcebergAllowed field if non-nil, zero value otherwise.

### GetIcebergAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIcebergAllowedOk() (*bool, bool)`

GetIcebergAllowedOk returns a tuple with the IcebergAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetIcebergAllowed(v bool)`

SetIcebergAllowed sets IcebergAllowed field to given value.

### HasIcebergAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasIcebergAllowed() bool`

HasIcebergAllowed returns a boolean if a field has been set.

### GetOcoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOcoAllowed() bool`

GetOcoAllowed returns the OcoAllowed field if non-nil, zero value otherwise.

### GetOcoAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOcoAllowedOk() (*bool, bool)`

GetOcoAllowedOk returns a tuple with the OcoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetOcoAllowed(v bool)`

SetOcoAllowed sets OcoAllowed field to given value.

### HasOcoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasOcoAllowed() bool`

HasOcoAllowed returns a boolean if a field has been set.

### GetOtoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOtoAllowed() bool`

GetOtoAllowed returns the OtoAllowed field if non-nil, zero value otherwise.

### GetOtoAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOtoAllowedOk() (*bool, bool)`

GetOtoAllowedOk returns a tuple with the OtoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetOtoAllowed(v bool)`

SetOtoAllowed sets OtoAllowed field to given value.

### HasOtoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasOtoAllowed() bool`

HasOtoAllowed returns a boolean if a field has been set.

### GetOpoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOpoAllowed() bool`

GetOpoAllowed returns the OpoAllowed field if non-nil, zero value otherwise.

### GetOpoAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetOpoAllowedOk() (*bool, bool)`

GetOpoAllowedOk returns a tuple with the OpoAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetOpoAllowed(v bool)`

SetOpoAllowed sets OpoAllowed field to given value.

### HasOpoAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasOpoAllowed() bool`

HasOpoAllowed returns a boolean if a field has been set.

### GetQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteOrderQtyMarketAllowed() bool`

GetQuoteOrderQtyMarketAllowed returns the QuoteOrderQtyMarketAllowed field if non-nil, zero value otherwise.

### GetQuoteOrderQtyMarketAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetQuoteOrderQtyMarketAllowedOk() (*bool, bool)`

GetQuoteOrderQtyMarketAllowedOk returns a tuple with the QuoteOrderQtyMarketAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetQuoteOrderQtyMarketAllowed(v bool)`

SetQuoteOrderQtyMarketAllowed sets QuoteOrderQtyMarketAllowed field to given value.

### HasQuoteOrderQtyMarketAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasQuoteOrderQtyMarketAllowed() bool`

HasQuoteOrderQtyMarketAllowed returns a boolean if a field has been set.

### GetAllowTrailingStop

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowTrailingStop() bool`

GetAllowTrailingStop returns the AllowTrailingStop field if non-nil, zero value otherwise.

### GetAllowTrailingStopOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowTrailingStopOk() (*bool, bool)`

GetAllowTrailingStopOk returns a tuple with the AllowTrailingStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTrailingStop

`func (o *ExchangeInfoResponseResultSymbolsInner) SetAllowTrailingStop(v bool)`

SetAllowTrailingStop sets AllowTrailingStop field to given value.

### HasAllowTrailingStop

`func (o *ExchangeInfoResponseResultSymbolsInner) HasAllowTrailingStop() bool`

HasAllowTrailingStop returns a boolean if a field has been set.

### GetCancelReplaceAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetCancelReplaceAllowed() bool`

GetCancelReplaceAllowed returns the CancelReplaceAllowed field if non-nil, zero value otherwise.

### GetCancelReplaceAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetCancelReplaceAllowedOk() (*bool, bool)`

GetCancelReplaceAllowedOk returns a tuple with the CancelReplaceAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReplaceAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetCancelReplaceAllowed(v bool)`

SetCancelReplaceAllowed sets CancelReplaceAllowed field to given value.

### HasCancelReplaceAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasCancelReplaceAllowed() bool`

HasCancelReplaceAllowed returns a boolean if a field has been set.

### GetAmendAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAmendAllowed() bool`

GetAmendAllowed returns the AmendAllowed field if non-nil, zero value otherwise.

### GetAmendAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAmendAllowedOk() (*bool, bool)`

GetAmendAllowedOk returns a tuple with the AmendAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetAmendAllowed(v bool)`

SetAmendAllowed sets AmendAllowed field to given value.

### HasAmendAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasAmendAllowed() bool`

HasAmendAllowed returns a boolean if a field has been set.

### GetPegInstructionsAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPegInstructionsAllowed() bool`

GetPegInstructionsAllowed returns the PegInstructionsAllowed field if non-nil, zero value otherwise.

### GetPegInstructionsAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPegInstructionsAllowedOk() (*bool, bool)`

GetPegInstructionsAllowedOk returns a tuple with the PegInstructionsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPegInstructionsAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetPegInstructionsAllowed(v bool)`

SetPegInstructionsAllowed sets PegInstructionsAllowed field to given value.

### HasPegInstructionsAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasPegInstructionsAllowed() bool`

HasPegInstructionsAllowed returns a boolean if a field has been set.

### GetIsSpotTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIsSpotTradingAllowed() bool`

GetIsSpotTradingAllowed returns the IsSpotTradingAllowed field if non-nil, zero value otherwise.

### GetIsSpotTradingAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIsSpotTradingAllowedOk() (*bool, bool)`

GetIsSpotTradingAllowedOk returns a tuple with the IsSpotTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpotTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetIsSpotTradingAllowed(v bool)`

SetIsSpotTradingAllowed sets IsSpotTradingAllowed field to given value.

### HasIsSpotTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasIsSpotTradingAllowed() bool`

HasIsSpotTradingAllowed returns a boolean if a field has been set.

### GetIsMarginTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIsMarginTradingAllowed() bool`

GetIsMarginTradingAllowed returns the IsMarginTradingAllowed field if non-nil, zero value otherwise.

### GetIsMarginTradingAllowedOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetIsMarginTradingAllowedOk() (*bool, bool)`

GetIsMarginTradingAllowedOk returns a tuple with the IsMarginTradingAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarginTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) SetIsMarginTradingAllowed(v bool)`

SetIsMarginTradingAllowed sets IsMarginTradingAllowed field to given value.

### HasIsMarginTradingAllowed

`func (o *ExchangeInfoResponseResultSymbolsInner) HasIsMarginTradingAllowed() bool`

HasIsMarginTradingAllowed returns a boolean if a field has been set.

### GetFilters

`func (o *ExchangeInfoResponseResultSymbolsInner) GetFilters() []SymbolFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetFiltersOk() (*[]SymbolFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExchangeInfoResponseResultSymbolsInner) SetFilters(v []SymbolFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ExchangeInfoResponseResultSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPermissions

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ExchangeInfoResponseResultSymbolsInner) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ExchangeInfoResponseResultSymbolsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetPermissionSets

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionSets() [][]string`

GetPermissionSets returns the PermissionSets field if non-nil, zero value otherwise.

### GetPermissionSetsOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetPermissionSetsOk() (*[][]string, bool)`

GetPermissionSetsOk returns a tuple with the PermissionSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSets

`func (o *ExchangeInfoResponseResultSymbolsInner) SetPermissionSets(v [][]string)`

SetPermissionSets sets PermissionSets field to given value.

### HasPermissionSets

`func (o *ExchangeInfoResponseResultSymbolsInner) HasPermissionSets() bool`

HasPermissionSets returns a boolean if a field has been set.

### GetDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseResultSymbolsInner) GetDefaultSelfTradePreventionMode() string`

GetDefaultSelfTradePreventionMode returns the DefaultSelfTradePreventionMode field if non-nil, zero value otherwise.

### GetDefaultSelfTradePreventionModeOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetDefaultSelfTradePreventionModeOk() (*string, bool)`

GetDefaultSelfTradePreventionModeOk returns a tuple with the DefaultSelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseResultSymbolsInner) SetDefaultSelfTradePreventionMode(v string)`

SetDefaultSelfTradePreventionMode sets DefaultSelfTradePreventionMode field to given value.

### HasDefaultSelfTradePreventionMode

`func (o *ExchangeInfoResponseResultSymbolsInner) HasDefaultSelfTradePreventionMode() bool`

HasDefaultSelfTradePreventionMode returns a boolean if a field has been set.

### GetAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowedSelfTradePreventionModes() []string`

GetAllowedSelfTradePreventionModes returns the AllowedSelfTradePreventionModes field if non-nil, zero value otherwise.

### GetAllowedSelfTradePreventionModesOk

`func (o *ExchangeInfoResponseResultSymbolsInner) GetAllowedSelfTradePreventionModesOk() (*[]string, bool)`

GetAllowedSelfTradePreventionModesOk returns a tuple with the AllowedSelfTradePreventionModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseResultSymbolsInner) SetAllowedSelfTradePreventionModes(v []string)`

SetAllowedSelfTradePreventionModes sets AllowedSelfTradePreventionModes field to given value.

### HasAllowedSelfTradePreventionModes

`func (o *ExchangeInfoResponseResultSymbolsInner) HasAllowedSelfTradePreventionModes() bool`

HasAllowedSelfTradePreventionModes returns a boolean if a field has been set.


[[Back to README]](../README.md)


