# QueryAlgoOrderResponse

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

### NewQueryAlgoOrderResponse

`func NewQueryAlgoOrderResponse() *QueryAlgoOrderResponse`

NewQueryAlgoOrderResponse instantiates a new QueryAlgoOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAlgoOrderResponseWithDefaults

`func NewQueryAlgoOrderResponseWithDefaults() *QueryAlgoOrderResponse`

NewQueryAlgoOrderResponseWithDefaults instantiates a new QueryAlgoOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *QueryAlgoOrderResponse) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *QueryAlgoOrderResponse) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *QueryAlgoOrderResponse) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *QueryAlgoOrderResponse) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *QueryAlgoOrderResponse) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *QueryAlgoOrderResponse) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *QueryAlgoOrderResponse) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *QueryAlgoOrderResponse) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetAlgoType

`func (o *QueryAlgoOrderResponse) GetAlgoType() string`

GetAlgoType returns the AlgoType field if non-nil, zero value otherwise.

### GetAlgoTypeOk

`func (o *QueryAlgoOrderResponse) GetAlgoTypeOk() (*string, bool)`

GetAlgoTypeOk returns a tuple with the AlgoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoType

`func (o *QueryAlgoOrderResponse) SetAlgoType(v string)`

SetAlgoType sets AlgoType field to given value.

### HasAlgoType

`func (o *QueryAlgoOrderResponse) HasAlgoType() bool`

HasAlgoType returns a boolean if a field has been set.

### GetOrderType

`func (o *QueryAlgoOrderResponse) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *QueryAlgoOrderResponse) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *QueryAlgoOrderResponse) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *QueryAlgoOrderResponse) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryAlgoOrderResponse) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryAlgoOrderResponse) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryAlgoOrderResponse) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryAlgoOrderResponse) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *QueryAlgoOrderResponse) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryAlgoOrderResponse) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryAlgoOrderResponse) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryAlgoOrderResponse) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryAlgoOrderResponse) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryAlgoOrderResponse) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryAlgoOrderResponse) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryAlgoOrderResponse) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryAlgoOrderResponse) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryAlgoOrderResponse) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryAlgoOrderResponse) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryAlgoOrderResponse) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *QueryAlgoOrderResponse) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QueryAlgoOrderResponse) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QueryAlgoOrderResponse) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QueryAlgoOrderResponse) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAlgoStatus

`func (o *QueryAlgoOrderResponse) GetAlgoStatus() string`

GetAlgoStatus returns the AlgoStatus field if non-nil, zero value otherwise.

### GetAlgoStatusOk

`func (o *QueryAlgoOrderResponse) GetAlgoStatusOk() (*string, bool)`

GetAlgoStatusOk returns a tuple with the AlgoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoStatus

`func (o *QueryAlgoOrderResponse) SetAlgoStatus(v string)`

SetAlgoStatus sets AlgoStatus field to given value.

### HasAlgoStatus

`func (o *QueryAlgoOrderResponse) HasAlgoStatus() bool`

HasAlgoStatus returns a boolean if a field has been set.

### GetActualOrderId

`func (o *QueryAlgoOrderResponse) GetActualOrderId() string`

GetActualOrderId returns the ActualOrderId field if non-nil, zero value otherwise.

### GetActualOrderIdOk

`func (o *QueryAlgoOrderResponse) GetActualOrderIdOk() (*string, bool)`

GetActualOrderIdOk returns a tuple with the ActualOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualOrderId

`func (o *QueryAlgoOrderResponse) SetActualOrderId(v string)`

SetActualOrderId sets ActualOrderId field to given value.

### HasActualOrderId

`func (o *QueryAlgoOrderResponse) HasActualOrderId() bool`

HasActualOrderId returns a boolean if a field has been set.

### GetActualPrice

`func (o *QueryAlgoOrderResponse) GetActualPrice() string`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *QueryAlgoOrderResponse) GetActualPriceOk() (*string, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *QueryAlgoOrderResponse) SetActualPrice(v string)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *QueryAlgoOrderResponse) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetTriggerPrice

`func (o *QueryAlgoOrderResponse) GetTriggerPrice() string`

GetTriggerPrice returns the TriggerPrice field if non-nil, zero value otherwise.

### GetTriggerPriceOk

`func (o *QueryAlgoOrderResponse) GetTriggerPriceOk() (*string, bool)`

GetTriggerPriceOk returns a tuple with the TriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPrice

`func (o *QueryAlgoOrderResponse) SetTriggerPrice(v string)`

SetTriggerPrice sets TriggerPrice field to given value.

### HasTriggerPrice

`func (o *QueryAlgoOrderResponse) HasTriggerPrice() bool`

HasTriggerPrice returns a boolean if a field has been set.

### GetPrice

`func (o *QueryAlgoOrderResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryAlgoOrderResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryAlgoOrderResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryAlgoOrderResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIcebergQuantity

`func (o *QueryAlgoOrderResponse) GetIcebergQuantity() string`

GetIcebergQuantity returns the IcebergQuantity field if non-nil, zero value otherwise.

### GetIcebergQuantityOk

`func (o *QueryAlgoOrderResponse) GetIcebergQuantityOk() (*string, bool)`

GetIcebergQuantityOk returns a tuple with the IcebergQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQuantity

`func (o *QueryAlgoOrderResponse) SetIcebergQuantity(v string)`

SetIcebergQuantity sets IcebergQuantity field to given value.

### HasIcebergQuantity

`func (o *QueryAlgoOrderResponse) HasIcebergQuantity() bool`

HasIcebergQuantity returns a boolean if a field has been set.

### GetTpTriggerPrice

`func (o *QueryAlgoOrderResponse) GetTpTriggerPrice() string`

GetTpTriggerPrice returns the TpTriggerPrice field if non-nil, zero value otherwise.

### GetTpTriggerPriceOk

`func (o *QueryAlgoOrderResponse) GetTpTriggerPriceOk() (*string, bool)`

GetTpTriggerPriceOk returns a tuple with the TpTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPrice

`func (o *QueryAlgoOrderResponse) SetTpTriggerPrice(v string)`

SetTpTriggerPrice sets TpTriggerPrice field to given value.

### HasTpTriggerPrice

`func (o *QueryAlgoOrderResponse) HasTpTriggerPrice() bool`

HasTpTriggerPrice returns a boolean if a field has been set.

### GetTpPrice

`func (o *QueryAlgoOrderResponse) GetTpPrice() string`

GetTpPrice returns the TpPrice field if non-nil, zero value otherwise.

### GetTpPriceOk

`func (o *QueryAlgoOrderResponse) GetTpPriceOk() (*string, bool)`

GetTpPriceOk returns a tuple with the TpPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpPrice

`func (o *QueryAlgoOrderResponse) SetTpPrice(v string)`

SetTpPrice sets TpPrice field to given value.

### HasTpPrice

`func (o *QueryAlgoOrderResponse) HasTpPrice() bool`

HasTpPrice returns a boolean if a field has been set.

### GetSlTriggerPrice

`func (o *QueryAlgoOrderResponse) GetSlTriggerPrice() string`

GetSlTriggerPrice returns the SlTriggerPrice field if non-nil, zero value otherwise.

### GetSlTriggerPriceOk

`func (o *QueryAlgoOrderResponse) GetSlTriggerPriceOk() (*string, bool)`

GetSlTriggerPriceOk returns a tuple with the SlTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPrice

`func (o *QueryAlgoOrderResponse) SetSlTriggerPrice(v string)`

SetSlTriggerPrice sets SlTriggerPrice field to given value.

### HasSlTriggerPrice

`func (o *QueryAlgoOrderResponse) HasSlTriggerPrice() bool`

HasSlTriggerPrice returns a boolean if a field has been set.

### GetSlPrice

`func (o *QueryAlgoOrderResponse) GetSlPrice() string`

GetSlPrice returns the SlPrice field if non-nil, zero value otherwise.

### GetSlPriceOk

`func (o *QueryAlgoOrderResponse) GetSlPriceOk() (*string, bool)`

GetSlPriceOk returns a tuple with the SlPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlPrice

`func (o *QueryAlgoOrderResponse) SetSlPrice(v string)`

SetSlPrice sets SlPrice field to given value.

### HasSlPrice

`func (o *QueryAlgoOrderResponse) HasSlPrice() bool`

HasSlPrice returns a boolean if a field has been set.

### GetTpOrderType

`func (o *QueryAlgoOrderResponse) GetTpOrderType() string`

GetTpOrderType returns the TpOrderType field if non-nil, zero value otherwise.

### GetTpOrderTypeOk

`func (o *QueryAlgoOrderResponse) GetTpOrderTypeOk() (*string, bool)`

GetTpOrderTypeOk returns a tuple with the TpOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrderType

`func (o *QueryAlgoOrderResponse) SetTpOrderType(v string)`

SetTpOrderType sets TpOrderType field to given value.

### HasTpOrderType

`func (o *QueryAlgoOrderResponse) HasTpOrderType() bool`

HasTpOrderType returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *QueryAlgoOrderResponse) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *QueryAlgoOrderResponse) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *QueryAlgoOrderResponse) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *QueryAlgoOrderResponse) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetWorkingType

`func (o *QueryAlgoOrderResponse) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *QueryAlgoOrderResponse) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *QueryAlgoOrderResponse) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *QueryAlgoOrderResponse) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryAlgoOrderResponse) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryAlgoOrderResponse) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryAlgoOrderResponse) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryAlgoOrderResponse) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetClosePosition

`func (o *QueryAlgoOrderResponse) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *QueryAlgoOrderResponse) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *QueryAlgoOrderResponse) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *QueryAlgoOrderResponse) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetPriceProtect

`func (o *QueryAlgoOrderResponse) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *QueryAlgoOrderResponse) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *QueryAlgoOrderResponse) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *QueryAlgoOrderResponse) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryAlgoOrderResponse) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryAlgoOrderResponse) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryAlgoOrderResponse) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryAlgoOrderResponse) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryAlgoOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryAlgoOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryAlgoOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryAlgoOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryAlgoOrderResponse) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryAlgoOrderResponse) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryAlgoOrderResponse) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryAlgoOrderResponse) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *QueryAlgoOrderResponse) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *QueryAlgoOrderResponse) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *QueryAlgoOrderResponse) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *QueryAlgoOrderResponse) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *QueryAlgoOrderResponse) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *QueryAlgoOrderResponse) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *QueryAlgoOrderResponse) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *QueryAlgoOrderResponse) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


