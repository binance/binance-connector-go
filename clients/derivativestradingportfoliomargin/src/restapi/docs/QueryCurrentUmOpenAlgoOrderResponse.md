# QueryCurrentUmOpenAlgoOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**AlgoId** | Pointer to **int64** |  | [optional] 
**ClientAlgoId** | Pointer to **string** |  | [optional] 
**AlgoType** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**PositionSide** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**AlgoStatus** | Pointer to **string** |  | [optional] 
**ActualOrderId** | Pointer to **string** |  | [optional] 
**ActualPrice** | Pointer to **string** |  | [optional] 
**TriggerPrice** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**IcebergQuantity** | Pointer to **string** |  | [optional] 
**TpTriggerPrice** | Pointer to **string** |  | [optional] 
**TpPrice** | Pointer to **string** |  | [optional] 
**SlTriggerPrice** | Pointer to **string** |  | [optional] 
**SlPrice** | Pointer to **string** |  | [optional] 
**TpOrderType** | Pointer to **string** |  | [optional] 
**SelfTradePreventionMode** | Pointer to **string** |  | [optional] 
**WorkingType** | Pointer to **string** |  | [optional] 
**PriceMatch** | Pointer to **string** |  | [optional] 
**ClosePosition** | Pointer to **bool** |  | [optional] 
**PriceProtect** | Pointer to **bool** |  | [optional] 
**ReduceOnly** | Pointer to **bool** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**UpdateTime** | Pointer to **int64** |  | [optional] 
**TriggerTime** | Pointer to **int64** |  | [optional] 
**GoodTillDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryCurrentUmOpenAlgoOrderResponse

`func NewQueryCurrentUmOpenAlgoOrderResponse() *QueryCurrentUmOpenAlgoOrderResponse`

NewQueryCurrentUmOpenAlgoOrderResponse instantiates a new QueryCurrentUmOpenAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCurrentUmOpenAlgoOrderResponseWithDefaults

`func NewQueryCurrentUmOpenAlgoOrderResponseWithDefaults() *QueryCurrentUmOpenAlgoOrderResponse`

NewQueryCurrentUmOpenAlgoOrderResponseWithDefaults instantiates a new QueryCurrentUmOpenAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetAlgoType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoType() string`

GetAlgoType returns the AlgoType field if non-nil, zero value otherwise.

### GetAlgoTypeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoTypeOk() (*string, bool)`

GetAlgoTypeOk returns a tuple with the AlgoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetAlgoType(v string)`

SetAlgoType sets AlgoType field to given value.

### HasAlgoType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasAlgoType() bool`

HasAlgoType returns a boolean if a field has been set.

### GetOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAlgoStatus

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoStatus() string`

GetAlgoStatus returns the AlgoStatus field if non-nil, zero value otherwise.

### GetAlgoStatusOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetAlgoStatusOk() (*string, bool)`

GetAlgoStatusOk returns a tuple with the AlgoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoStatus

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetAlgoStatus(v string)`

SetAlgoStatus sets AlgoStatus field to given value.

### HasAlgoStatus

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasAlgoStatus() bool`

HasAlgoStatus returns a boolean if a field has been set.

### GetActualOrderId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetActualOrderId() string`

GetActualOrderId returns the ActualOrderId field if non-nil, zero value otherwise.

### GetActualOrderIdOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetActualOrderIdOk() (*string, bool)`

GetActualOrderIdOk returns a tuple with the ActualOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualOrderId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetActualOrderId(v string)`

SetActualOrderId sets ActualOrderId field to given value.

### HasActualOrderId

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasActualOrderId() bool`

HasActualOrderId returns a boolean if a field has been set.

### GetActualPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetActualPrice() string`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetActualPriceOk() (*string, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetActualPrice(v string)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTriggerPrice() string`

GetTriggerPrice returns the TriggerPrice field if non-nil, zero value otherwise.

### GetTriggerPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTriggerPriceOk() (*string, bool)`

GetTriggerPriceOk returns a tuple with the TriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTriggerPrice(v string)`

SetTriggerPrice sets TriggerPrice field to given value.

### HasTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTriggerPrice() bool`

HasTriggerPrice returns a boolean if a field has been set.

### GetPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIcebergQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetIcebergQuantity() string`

GetIcebergQuantity returns the IcebergQuantity field if non-nil, zero value otherwise.

### GetIcebergQuantityOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetIcebergQuantityOk() (*string, bool)`

GetIcebergQuantityOk returns a tuple with the IcebergQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetIcebergQuantity(v string)`

SetIcebergQuantity sets IcebergQuantity field to given value.

### HasIcebergQuantity

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasIcebergQuantity() bool`

HasIcebergQuantity returns a boolean if a field has been set.

### GetTpTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpTriggerPrice() string`

GetTpTriggerPrice returns the TpTriggerPrice field if non-nil, zero value otherwise.

### GetTpTriggerPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpTriggerPriceOk() (*string, bool)`

GetTpTriggerPriceOk returns a tuple with the TpTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTpTriggerPrice(v string)`

SetTpTriggerPrice sets TpTriggerPrice field to given value.

### HasTpTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTpTriggerPrice() bool`

HasTpTriggerPrice returns a boolean if a field has been set.

### GetTpPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpPrice() string`

GetTpPrice returns the TpPrice field if non-nil, zero value otherwise.

### GetTpPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpPriceOk() (*string, bool)`

GetTpPriceOk returns a tuple with the TpPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTpPrice(v string)`

SetTpPrice sets TpPrice field to given value.

### HasTpPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTpPrice() bool`

HasTpPrice returns a boolean if a field has been set.

### GetSlTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSlTriggerPrice() string`

GetSlTriggerPrice returns the SlTriggerPrice field if non-nil, zero value otherwise.

### GetSlTriggerPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSlTriggerPriceOk() (*string, bool)`

GetSlTriggerPriceOk returns a tuple with the SlTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetSlTriggerPrice(v string)`

SetSlTriggerPrice sets SlTriggerPrice field to given value.

### HasSlTriggerPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasSlTriggerPrice() bool`

HasSlTriggerPrice returns a boolean if a field has been set.

### GetSlPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSlPrice() string`

GetSlPrice returns the SlPrice field if non-nil, zero value otherwise.

### GetSlPriceOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSlPriceOk() (*string, bool)`

GetSlPriceOk returns a tuple with the SlPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetSlPrice(v string)`

SetSlPrice sets SlPrice field to given value.

### HasSlPrice

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasSlPrice() bool`

HasSlPrice returns a boolean if a field has been set.

### GetTpOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpOrderType() string`

GetTpOrderType returns the TpOrderType field if non-nil, zero value otherwise.

### GetTpOrderTypeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTpOrderTypeOk() (*string, bool)`

GetTpOrderTypeOk returns a tuple with the TpOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTpOrderType(v string)`

SetTpOrderType sets TpOrderType field to given value.

### HasTpOrderType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTpOrderType() bool`

HasTpOrderType returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetWorkingType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetClosePosition

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetPriceProtect

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *QueryCurrentUmOpenAlgoOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *QueryCurrentUmOpenAlgoOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *QueryCurrentUmOpenAlgoOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


