# GetC2CTradeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderNumber** | Pointer to **string** |  | [optional] 
**AdvNo** | Pointer to **string** |  | [optional] 
**TradeType** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
**Fiat** | Pointer to **string** |  | [optional] 
**FiatSymbol** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**TotalPrice** | Pointer to **string** |  | [optional] 
**UnitPrice** | Pointer to **string** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**TakerCommissionRate** | Pointer to **string** |  | [optional] 
**TakerCommission** | Pointer to **string** |  | [optional] 
**TakerAmount** | Pointer to **string** |  | [optional] 
**CounterPartNickName** | Pointer to **string** |  | [optional] 
**PayMethodName** | Pointer to **string** |  | [optional] 
**AdditionalKycVerify** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetC2CTradeHistoryResponse

`func NewGetC2CTradeHistoryResponse() *GetC2CTradeHistoryResponse`

NewGetC2CTradeHistoryResponse instantiates a new GetC2CTradeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetC2CTradeHistoryResponseWithDefaults

`func NewGetC2CTradeHistoryResponseWithDefaults() *GetC2CTradeHistoryResponse`

NewGetC2CTradeHistoryResponseWithDefaults instantiates a new GetC2CTradeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderNumber

`func (o *GetC2CTradeHistoryResponse) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *GetC2CTradeHistoryResponse) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *GetC2CTradeHistoryResponse) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *GetC2CTradeHistoryResponse) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetAdvNo

`func (o *GetC2CTradeHistoryResponse) GetAdvNo() string`

GetAdvNo returns the AdvNo field if non-nil, zero value otherwise.

### GetAdvNoOk

`func (o *GetC2CTradeHistoryResponse) GetAdvNoOk() (*string, bool)`

GetAdvNoOk returns a tuple with the AdvNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvNo

`func (o *GetC2CTradeHistoryResponse) SetAdvNo(v string)`

SetAdvNo sets AdvNo field to given value.

### HasAdvNo

`func (o *GetC2CTradeHistoryResponse) HasAdvNo() bool`

HasAdvNo returns a boolean if a field has been set.

### GetTradeType

`func (o *GetC2CTradeHistoryResponse) GetTradeType() string`

GetTradeType returns the TradeType field if non-nil, zero value otherwise.

### GetTradeTypeOk

`func (o *GetC2CTradeHistoryResponse) GetTradeTypeOk() (*string, bool)`

GetTradeTypeOk returns a tuple with the TradeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeType

`func (o *GetC2CTradeHistoryResponse) SetTradeType(v string)`

SetTradeType sets TradeType field to given value.

### HasTradeType

`func (o *GetC2CTradeHistoryResponse) HasTradeType() bool`

HasTradeType returns a boolean if a field has been set.

### GetAsset

`func (o *GetC2CTradeHistoryResponse) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *GetC2CTradeHistoryResponse) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *GetC2CTradeHistoryResponse) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *GetC2CTradeHistoryResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetFiat

`func (o *GetC2CTradeHistoryResponse) GetFiat() string`

GetFiat returns the Fiat field if non-nil, zero value otherwise.

### GetFiatOk

`func (o *GetC2CTradeHistoryResponse) GetFiatOk() (*string, bool)`

GetFiatOk returns a tuple with the Fiat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiat

`func (o *GetC2CTradeHistoryResponse) SetFiat(v string)`

SetFiat sets Fiat field to given value.

### HasFiat

`func (o *GetC2CTradeHistoryResponse) HasFiat() bool`

HasFiat returns a boolean if a field has been set.

### GetFiatSymbol

`func (o *GetC2CTradeHistoryResponse) GetFiatSymbol() string`

GetFiatSymbol returns the FiatSymbol field if non-nil, zero value otherwise.

### GetFiatSymbolOk

`func (o *GetC2CTradeHistoryResponse) GetFiatSymbolOk() (*string, bool)`

GetFiatSymbolOk returns a tuple with the FiatSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatSymbol

`func (o *GetC2CTradeHistoryResponse) SetFiatSymbol(v string)`

SetFiatSymbol sets FiatSymbol field to given value.

### HasFiatSymbol

`func (o *GetC2CTradeHistoryResponse) HasFiatSymbol() bool`

HasFiatSymbol returns a boolean if a field has been set.

### GetAmount

`func (o *GetC2CTradeHistoryResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetC2CTradeHistoryResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetC2CTradeHistoryResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetC2CTradeHistoryResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotalPrice

`func (o *GetC2CTradeHistoryResponse) GetTotalPrice() string`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *GetC2CTradeHistoryResponse) GetTotalPriceOk() (*string, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *GetC2CTradeHistoryResponse) SetTotalPrice(v string)`

SetTotalPrice sets TotalPrice field to given value.

### HasTotalPrice

`func (o *GetC2CTradeHistoryResponse) HasTotalPrice() bool`

HasTotalPrice returns a boolean if a field has been set.

### GetUnitPrice

`func (o *GetC2CTradeHistoryResponse) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *GetC2CTradeHistoryResponse) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *GetC2CTradeHistoryResponse) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *GetC2CTradeHistoryResponse) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetOrderStatus

`func (o *GetC2CTradeHistoryResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *GetC2CTradeHistoryResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *GetC2CTradeHistoryResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *GetC2CTradeHistoryResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetC2CTradeHistoryResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetC2CTradeHistoryResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetC2CTradeHistoryResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetC2CTradeHistoryResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCommission

`func (o *GetC2CTradeHistoryResponse) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *GetC2CTradeHistoryResponse) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *GetC2CTradeHistoryResponse) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *GetC2CTradeHistoryResponse) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetTakerCommissionRate

`func (o *GetC2CTradeHistoryResponse) GetTakerCommissionRate() string`

GetTakerCommissionRate returns the TakerCommissionRate field if non-nil, zero value otherwise.

### GetTakerCommissionRateOk

`func (o *GetC2CTradeHistoryResponse) GetTakerCommissionRateOk() (*string, bool)`

GetTakerCommissionRateOk returns a tuple with the TakerCommissionRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommissionRate

`func (o *GetC2CTradeHistoryResponse) SetTakerCommissionRate(v string)`

SetTakerCommissionRate sets TakerCommissionRate field to given value.

### HasTakerCommissionRate

`func (o *GetC2CTradeHistoryResponse) HasTakerCommissionRate() bool`

HasTakerCommissionRate returns a boolean if a field has been set.

### GetTakerCommission

`func (o *GetC2CTradeHistoryResponse) GetTakerCommission() string`

GetTakerCommission returns the TakerCommission field if non-nil, zero value otherwise.

### GetTakerCommissionOk

`func (o *GetC2CTradeHistoryResponse) GetTakerCommissionOk() (*string, bool)`

GetTakerCommissionOk returns a tuple with the TakerCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerCommission

`func (o *GetC2CTradeHistoryResponse) SetTakerCommission(v string)`

SetTakerCommission sets TakerCommission field to given value.

### HasTakerCommission

`func (o *GetC2CTradeHistoryResponse) HasTakerCommission() bool`

HasTakerCommission returns a boolean if a field has been set.

### GetTakerAmount

`func (o *GetC2CTradeHistoryResponse) GetTakerAmount() string`

GetTakerAmount returns the TakerAmount field if non-nil, zero value otherwise.

### GetTakerAmountOk

`func (o *GetC2CTradeHistoryResponse) GetTakerAmountOk() (*string, bool)`

GetTakerAmountOk returns a tuple with the TakerAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakerAmount

`func (o *GetC2CTradeHistoryResponse) SetTakerAmount(v string)`

SetTakerAmount sets TakerAmount field to given value.

### HasTakerAmount

`func (o *GetC2CTradeHistoryResponse) HasTakerAmount() bool`

HasTakerAmount returns a boolean if a field has been set.

### GetCounterPartNickName

`func (o *GetC2CTradeHistoryResponse) GetCounterPartNickName() string`

GetCounterPartNickName returns the CounterPartNickName field if non-nil, zero value otherwise.

### GetCounterPartNickNameOk

`func (o *GetC2CTradeHistoryResponse) GetCounterPartNickNameOk() (*string, bool)`

GetCounterPartNickNameOk returns a tuple with the CounterPartNickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterPartNickName

`func (o *GetC2CTradeHistoryResponse) SetCounterPartNickName(v string)`

SetCounterPartNickName sets CounterPartNickName field to given value.

### HasCounterPartNickName

`func (o *GetC2CTradeHistoryResponse) HasCounterPartNickName() bool`

HasCounterPartNickName returns a boolean if a field has been set.

### GetPayMethodName

`func (o *GetC2CTradeHistoryResponse) GetPayMethodName() string`

GetPayMethodName returns the PayMethodName field if non-nil, zero value otherwise.

### GetPayMethodNameOk

`func (o *GetC2CTradeHistoryResponse) GetPayMethodNameOk() (*string, bool)`

GetPayMethodNameOk returns a tuple with the PayMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethodName

`func (o *GetC2CTradeHistoryResponse) SetPayMethodName(v string)`

SetPayMethodName sets PayMethodName field to given value.

### HasPayMethodName

`func (o *GetC2CTradeHistoryResponse) HasPayMethodName() bool`

HasPayMethodName returns a boolean if a field has been set.

### GetAdditionalKycVerify

`func (o *GetC2CTradeHistoryResponse) GetAdditionalKycVerify() int64`

GetAdditionalKycVerify returns the AdditionalKycVerify field if non-nil, zero value otherwise.

### GetAdditionalKycVerifyOk

`func (o *GetC2CTradeHistoryResponse) GetAdditionalKycVerifyOk() (*int64, bool)`

GetAdditionalKycVerifyOk returns a tuple with the AdditionalKycVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalKycVerify

`func (o *GetC2CTradeHistoryResponse) SetAdditionalKycVerify(v int64)`

SetAdditionalKycVerify sets AdditionalKycVerify field to given value.

### HasAdditionalKycVerify

`func (o *GetC2CTradeHistoryResponse) HasAdditionalKycVerify() bool`

HasAdditionalKycVerify returns a boolean if a field has been set.


[[Back to README]](../README.md)


