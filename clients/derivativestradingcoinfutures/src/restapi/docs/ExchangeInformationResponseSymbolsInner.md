# ExchangeInformationResponseSymbolsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]ExchangeInformationResponseSymbolsInnerFiltersInner**](ExchangeInformationResponseSymbolsInnerFiltersInner.md) |  | [optional] 
**OrderType** | Pointer to **[]string** |  | [optional] 
**TimeInForce** | Pointer to **[]string** |  | [optional] 
**LiquidationFee** | Pointer to **string** |  | [optional] 
**MarketTakeBound** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Pair** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** |  | [optional] 
**DeliveryDate** | Pointer to **int64** |  | [optional] 
**OnboardDate** | Pointer to **int64** |  | [optional] 
**ContractStatus** | Pointer to **string** |  | [optional] 
**ContractSize** | Pointer to **int64** |  | [optional] 
**QuoteAsset** | Pointer to **string** |  | [optional] 
**BaseAsset** | Pointer to **string** |  | [optional] 
**MarginAsset** | Pointer to **string** |  | [optional] 
**PricePrecision** | Pointer to **int64** |  | [optional] 
**QuantityPrecision** | Pointer to **int64** |  | [optional] 
**BaseAssetPrecision** | Pointer to **int64** |  | [optional] 
**QuotePrecision** | Pointer to **int64** |  | [optional] 
**EqualQtyPrecision** | Pointer to **int64** |  | [optional] 
**TriggerProtect** | Pointer to **string** |  | [optional] 
**MaintMarginPercent** | Pointer to **string** |  | [optional] 
**RequiredMarginPercent** | Pointer to **string** |  | [optional] 
**UnderlyingType** | Pointer to **string** |  | [optional] 
**UnderlyingSubType** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExchangeInformationResponseSymbolsInner

`func NewExchangeInformationResponseSymbolsInner() *ExchangeInformationResponseSymbolsInner`

NewExchangeInformationResponseSymbolsInner instantiates a new ExchangeInformationResponseSymbolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeInformationResponseSymbolsInnerWithDefaults

`func NewExchangeInformationResponseSymbolsInnerWithDefaults() *ExchangeInformationResponseSymbolsInner`

NewExchangeInformationResponseSymbolsInnerWithDefaults instantiates a new ExchangeInformationResponseSymbolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *ExchangeInformationResponseSymbolsInner) GetFilters() []ExchangeInformationResponseSymbolsInnerFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ExchangeInformationResponseSymbolsInner) GetFiltersOk() (*[]ExchangeInformationResponseSymbolsInnerFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ExchangeInformationResponseSymbolsInner) SetFilters(v []ExchangeInformationResponseSymbolsInnerFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ExchangeInformationResponseSymbolsInner) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrderType

`func (o *ExchangeInformationResponseSymbolsInner) GetOrderType() []string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetOrderTypeOk() (*[]string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *ExchangeInformationResponseSymbolsInner) SetOrderType(v []string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *ExchangeInformationResponseSymbolsInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *ExchangeInformationResponseSymbolsInner) GetTimeInForce() []string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *ExchangeInformationResponseSymbolsInner) GetTimeInForceOk() (*[]string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *ExchangeInformationResponseSymbolsInner) SetTimeInForce(v []string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *ExchangeInformationResponseSymbolsInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetLiquidationFee

`func (o *ExchangeInformationResponseSymbolsInner) GetLiquidationFee() string`

GetLiquidationFee returns the LiquidationFee field if non-nil, zero value otherwise.

### GetLiquidationFeeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetLiquidationFeeOk() (*string, bool)`

GetLiquidationFeeOk returns a tuple with the LiquidationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationFee

`func (o *ExchangeInformationResponseSymbolsInner) SetLiquidationFee(v string)`

SetLiquidationFee sets LiquidationFee field to given value.

### HasLiquidationFee

`func (o *ExchangeInformationResponseSymbolsInner) HasLiquidationFee() bool`

HasLiquidationFee returns a boolean if a field has been set.

### GetMarketTakeBound

`func (o *ExchangeInformationResponseSymbolsInner) GetMarketTakeBound() string`

GetMarketTakeBound returns the MarketTakeBound field if non-nil, zero value otherwise.

### GetMarketTakeBoundOk

`func (o *ExchangeInformationResponseSymbolsInner) GetMarketTakeBoundOk() (*string, bool)`

GetMarketTakeBoundOk returns a tuple with the MarketTakeBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTakeBound

`func (o *ExchangeInformationResponseSymbolsInner) SetMarketTakeBound(v string)`

SetMarketTakeBound sets MarketTakeBound field to given value.

### HasMarketTakeBound

`func (o *ExchangeInformationResponseSymbolsInner) HasMarketTakeBound() bool`

HasMarketTakeBound returns a boolean if a field has been set.

### GetSymbol

`func (o *ExchangeInformationResponseSymbolsInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ExchangeInformationResponseSymbolsInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ExchangeInformationResponseSymbolsInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ExchangeInformationResponseSymbolsInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetPair

`func (o *ExchangeInformationResponseSymbolsInner) GetPair() string`

GetPair returns the Pair field if non-nil, zero value otherwise.

### GetPairOk

`func (o *ExchangeInformationResponseSymbolsInner) GetPairOk() (*string, bool)`

GetPairOk returns a tuple with the Pair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPair

`func (o *ExchangeInformationResponseSymbolsInner) SetPair(v string)`

SetPair sets Pair field to given value.

### HasPair

`func (o *ExchangeInformationResponseSymbolsInner) HasPair() bool`

HasPair returns a boolean if a field has been set.

### GetContractType

`func (o *ExchangeInformationResponseSymbolsInner) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *ExchangeInformationResponseSymbolsInner) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *ExchangeInformationResponseSymbolsInner) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *ExchangeInformationResponseSymbolsInner) GetDeliveryDate() int64`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *ExchangeInformationResponseSymbolsInner) GetDeliveryDateOk() (*int64, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *ExchangeInformationResponseSymbolsInner) SetDeliveryDate(v int64)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *ExchangeInformationResponseSymbolsInner) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetOnboardDate

`func (o *ExchangeInformationResponseSymbolsInner) GetOnboardDate() int64`

GetOnboardDate returns the OnboardDate field if non-nil, zero value otherwise.

### GetOnboardDateOk

`func (o *ExchangeInformationResponseSymbolsInner) GetOnboardDateOk() (*int64, bool)`

GetOnboardDateOk returns a tuple with the OnboardDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardDate

`func (o *ExchangeInformationResponseSymbolsInner) SetOnboardDate(v int64)`

SetOnboardDate sets OnboardDate field to given value.

### HasOnboardDate

`func (o *ExchangeInformationResponseSymbolsInner) HasOnboardDate() bool`

HasOnboardDate returns a boolean if a field has been set.

### GetContractStatus

`func (o *ExchangeInformationResponseSymbolsInner) GetContractStatus() string`

GetContractStatus returns the ContractStatus field if non-nil, zero value otherwise.

### GetContractStatusOk

`func (o *ExchangeInformationResponseSymbolsInner) GetContractStatusOk() (*string, bool)`

GetContractStatusOk returns a tuple with the ContractStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractStatus

`func (o *ExchangeInformationResponseSymbolsInner) SetContractStatus(v string)`

SetContractStatus sets ContractStatus field to given value.

### HasContractStatus

`func (o *ExchangeInformationResponseSymbolsInner) HasContractStatus() bool`

HasContractStatus returns a boolean if a field has been set.

### GetContractSize

`func (o *ExchangeInformationResponseSymbolsInner) GetContractSize() int64`

GetContractSize returns the ContractSize field if non-nil, zero value otherwise.

### GetContractSizeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetContractSizeOk() (*int64, bool)`

GetContractSizeOk returns a tuple with the ContractSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractSize

`func (o *ExchangeInformationResponseSymbolsInner) SetContractSize(v int64)`

SetContractSize sets ContractSize field to given value.

### HasContractSize

`func (o *ExchangeInformationResponseSymbolsInner) HasContractSize() bool`

HasContractSize returns a boolean if a field has been set.

### GetQuoteAsset

`func (o *ExchangeInformationResponseSymbolsInner) GetQuoteAsset() string`

GetQuoteAsset returns the QuoteAsset field if non-nil, zero value otherwise.

### GetQuoteAssetOk

`func (o *ExchangeInformationResponseSymbolsInner) GetQuoteAssetOk() (*string, bool)`

GetQuoteAssetOk returns a tuple with the QuoteAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteAsset

`func (o *ExchangeInformationResponseSymbolsInner) SetQuoteAsset(v string)`

SetQuoteAsset sets QuoteAsset field to given value.

### HasQuoteAsset

`func (o *ExchangeInformationResponseSymbolsInner) HasQuoteAsset() bool`

HasQuoteAsset returns a boolean if a field has been set.

### GetBaseAsset

`func (o *ExchangeInformationResponseSymbolsInner) GetBaseAsset() string`

GetBaseAsset returns the BaseAsset field if non-nil, zero value otherwise.

### GetBaseAssetOk

`func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetOk() (*string, bool)`

GetBaseAssetOk returns a tuple with the BaseAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAsset

`func (o *ExchangeInformationResponseSymbolsInner) SetBaseAsset(v string)`

SetBaseAsset sets BaseAsset field to given value.

### HasBaseAsset

`func (o *ExchangeInformationResponseSymbolsInner) HasBaseAsset() bool`

HasBaseAsset returns a boolean if a field has been set.

### GetMarginAsset

`func (o *ExchangeInformationResponseSymbolsInner) GetMarginAsset() string`

GetMarginAsset returns the MarginAsset field if non-nil, zero value otherwise.

### GetMarginAssetOk

`func (o *ExchangeInformationResponseSymbolsInner) GetMarginAssetOk() (*string, bool)`

GetMarginAssetOk returns a tuple with the MarginAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginAsset

`func (o *ExchangeInformationResponseSymbolsInner) SetMarginAsset(v string)`

SetMarginAsset sets MarginAsset field to given value.

### HasMarginAsset

`func (o *ExchangeInformationResponseSymbolsInner) HasMarginAsset() bool`

HasMarginAsset returns a boolean if a field has been set.

### GetPricePrecision

`func (o *ExchangeInformationResponseSymbolsInner) GetPricePrecision() int64`

GetPricePrecision returns the PricePrecision field if non-nil, zero value otherwise.

### GetPricePrecisionOk

`func (o *ExchangeInformationResponseSymbolsInner) GetPricePrecisionOk() (*int64, bool)`

GetPricePrecisionOk returns a tuple with the PricePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePrecision

`func (o *ExchangeInformationResponseSymbolsInner) SetPricePrecision(v int64)`

SetPricePrecision sets PricePrecision field to given value.

### HasPricePrecision

`func (o *ExchangeInformationResponseSymbolsInner) HasPricePrecision() bool`

HasPricePrecision returns a boolean if a field has been set.

### GetQuantityPrecision

`func (o *ExchangeInformationResponseSymbolsInner) GetQuantityPrecision() int64`

GetQuantityPrecision returns the QuantityPrecision field if non-nil, zero value otherwise.

### GetQuantityPrecisionOk

`func (o *ExchangeInformationResponseSymbolsInner) GetQuantityPrecisionOk() (*int64, bool)`

GetQuantityPrecisionOk returns a tuple with the QuantityPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityPrecision

`func (o *ExchangeInformationResponseSymbolsInner) SetQuantityPrecision(v int64)`

SetQuantityPrecision sets QuantityPrecision field to given value.

### HasQuantityPrecision

`func (o *ExchangeInformationResponseSymbolsInner) HasQuantityPrecision() bool`

HasQuantityPrecision returns a boolean if a field has been set.

### GetBaseAssetPrecision

`func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetPrecision() int64`

GetBaseAssetPrecision returns the BaseAssetPrecision field if non-nil, zero value otherwise.

### GetBaseAssetPrecisionOk

`func (o *ExchangeInformationResponseSymbolsInner) GetBaseAssetPrecisionOk() (*int64, bool)`

GetBaseAssetPrecisionOk returns a tuple with the BaseAssetPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAssetPrecision

`func (o *ExchangeInformationResponseSymbolsInner) SetBaseAssetPrecision(v int64)`

SetBaseAssetPrecision sets BaseAssetPrecision field to given value.

### HasBaseAssetPrecision

`func (o *ExchangeInformationResponseSymbolsInner) HasBaseAssetPrecision() bool`

HasBaseAssetPrecision returns a boolean if a field has been set.

### GetQuotePrecision

`func (o *ExchangeInformationResponseSymbolsInner) GetQuotePrecision() int64`

GetQuotePrecision returns the QuotePrecision field if non-nil, zero value otherwise.

### GetQuotePrecisionOk

`func (o *ExchangeInformationResponseSymbolsInner) GetQuotePrecisionOk() (*int64, bool)`

GetQuotePrecisionOk returns a tuple with the QuotePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotePrecision

`func (o *ExchangeInformationResponseSymbolsInner) SetQuotePrecision(v int64)`

SetQuotePrecision sets QuotePrecision field to given value.

### HasQuotePrecision

`func (o *ExchangeInformationResponseSymbolsInner) HasQuotePrecision() bool`

HasQuotePrecision returns a boolean if a field has been set.

### GetEqualQtyPrecision

`func (o *ExchangeInformationResponseSymbolsInner) GetEqualQtyPrecision() int64`

GetEqualQtyPrecision returns the EqualQtyPrecision field if non-nil, zero value otherwise.

### GetEqualQtyPrecisionOk

`func (o *ExchangeInformationResponseSymbolsInner) GetEqualQtyPrecisionOk() (*int64, bool)`

GetEqualQtyPrecisionOk returns a tuple with the EqualQtyPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEqualQtyPrecision

`func (o *ExchangeInformationResponseSymbolsInner) SetEqualQtyPrecision(v int64)`

SetEqualQtyPrecision sets EqualQtyPrecision field to given value.

### HasEqualQtyPrecision

`func (o *ExchangeInformationResponseSymbolsInner) HasEqualQtyPrecision() bool`

HasEqualQtyPrecision returns a boolean if a field has been set.

### GetTriggerProtect

`func (o *ExchangeInformationResponseSymbolsInner) GetTriggerProtect() string`

GetTriggerProtect returns the TriggerProtect field if non-nil, zero value otherwise.

### GetTriggerProtectOk

`func (o *ExchangeInformationResponseSymbolsInner) GetTriggerProtectOk() (*string, bool)`

GetTriggerProtectOk returns a tuple with the TriggerProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerProtect

`func (o *ExchangeInformationResponseSymbolsInner) SetTriggerProtect(v string)`

SetTriggerProtect sets TriggerProtect field to given value.

### HasTriggerProtect

`func (o *ExchangeInformationResponseSymbolsInner) HasTriggerProtect() bool`

HasTriggerProtect returns a boolean if a field has been set.

### GetMaintMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) GetMaintMarginPercent() string`

GetMaintMarginPercent returns the MaintMarginPercent field if non-nil, zero value otherwise.

### GetMaintMarginPercentOk

`func (o *ExchangeInformationResponseSymbolsInner) GetMaintMarginPercentOk() (*string, bool)`

GetMaintMarginPercentOk returns a tuple with the MaintMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) SetMaintMarginPercent(v string)`

SetMaintMarginPercent sets MaintMarginPercent field to given value.

### HasMaintMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) HasMaintMarginPercent() bool`

HasMaintMarginPercent returns a boolean if a field has been set.

### GetRequiredMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) GetRequiredMarginPercent() string`

GetRequiredMarginPercent returns the RequiredMarginPercent field if non-nil, zero value otherwise.

### GetRequiredMarginPercentOk

`func (o *ExchangeInformationResponseSymbolsInner) GetRequiredMarginPercentOk() (*string, bool)`

GetRequiredMarginPercentOk returns a tuple with the RequiredMarginPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) SetRequiredMarginPercent(v string)`

SetRequiredMarginPercent sets RequiredMarginPercent field to given value.

### HasRequiredMarginPercent

`func (o *ExchangeInformationResponseSymbolsInner) HasRequiredMarginPercent() bool`

HasRequiredMarginPercent returns a boolean if a field has been set.

### GetUnderlyingType

`func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingType() string`

GetUnderlyingType returns the UnderlyingType field if non-nil, zero value otherwise.

### GetUnderlyingTypeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingTypeOk() (*string, bool)`

GetUnderlyingTypeOk returns a tuple with the UnderlyingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingType

`func (o *ExchangeInformationResponseSymbolsInner) SetUnderlyingType(v string)`

SetUnderlyingType sets UnderlyingType field to given value.

### HasUnderlyingType

`func (o *ExchangeInformationResponseSymbolsInner) HasUnderlyingType() bool`

HasUnderlyingType returns a boolean if a field has been set.

### GetUnderlyingSubType

`func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingSubType() []string`

GetUnderlyingSubType returns the UnderlyingSubType field if non-nil, zero value otherwise.

### GetUnderlyingSubTypeOk

`func (o *ExchangeInformationResponseSymbolsInner) GetUnderlyingSubTypeOk() (*[]string, bool)`

GetUnderlyingSubTypeOk returns a tuple with the UnderlyingSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSubType

`func (o *ExchangeInformationResponseSymbolsInner) SetUnderlyingSubType(v []string)`

SetUnderlyingSubType sets UnderlyingSubType field to given value.

### HasUnderlyingSubType

`func (o *ExchangeInformationResponseSymbolsInner) HasUnderlyingSubType() bool`

HasUnderlyingSubType returns a boolean if a field has been set.


[[Back to README]](../README.md)


