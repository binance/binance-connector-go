# QueryAllAlgoOrdersResponseInner

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

### NewQueryAllAlgoOrdersResponseInner

`func NewQueryAllAlgoOrdersResponseInner() *QueryAllAlgoOrdersResponseInner`

NewQueryAllAlgoOrdersResponseInner instantiates a new QueryAllAlgoOrdersResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAllAlgoOrdersResponseInnerWithDefaults

`func NewQueryAllAlgoOrdersResponseInnerWithDefaults() *QueryAllAlgoOrdersResponseInner`

NewQueryAllAlgoOrdersResponseInnerWithDefaults instantiates a new QueryAllAlgoOrdersResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoId() int64`

GetAlgoId returns the AlgoId field if non-nil, zero value otherwise.

### GetAlgoIdOk

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoIdOk() (*int64, bool)`

GetAlgoIdOk returns a tuple with the AlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) SetAlgoId(v int64)`

SetAlgoId sets AlgoId field to given value.

### HasAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) HasAlgoId() bool`

HasAlgoId returns a boolean if a field has been set.

### GetClientAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) GetClientAlgoId() string`

GetClientAlgoId returns the ClientAlgoId field if non-nil, zero value otherwise.

### GetClientAlgoIdOk

`func (o *QueryAllAlgoOrdersResponseInner) GetClientAlgoIdOk() (*string, bool)`

GetClientAlgoIdOk returns a tuple with the ClientAlgoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) SetClientAlgoId(v string)`

SetClientAlgoId sets ClientAlgoId field to given value.

### HasClientAlgoId

`func (o *QueryAllAlgoOrdersResponseInner) HasClientAlgoId() bool`

HasClientAlgoId returns a boolean if a field has been set.

### GetAlgoType

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoType() string`

GetAlgoType returns the AlgoType field if non-nil, zero value otherwise.

### GetAlgoTypeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoTypeOk() (*string, bool)`

GetAlgoTypeOk returns a tuple with the AlgoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoType

`func (o *QueryAllAlgoOrdersResponseInner) SetAlgoType(v string)`

SetAlgoType sets AlgoType field to given value.

### HasAlgoType

`func (o *QueryAllAlgoOrdersResponseInner) HasAlgoType() bool`

HasAlgoType returns a boolean if a field has been set.

### GetOrderType

`func (o *QueryAllAlgoOrdersResponseInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *QueryAllAlgoOrdersResponseInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *QueryAllAlgoOrdersResponseInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetSymbol

`func (o *QueryAllAlgoOrdersResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryAllAlgoOrdersResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryAllAlgoOrdersResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryAllAlgoOrdersResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *QueryAllAlgoOrdersResponseInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *QueryAllAlgoOrdersResponseInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *QueryAllAlgoOrdersResponseInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *QueryAllAlgoOrdersResponseInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *QueryAllAlgoOrdersResponseInner) GetPositionSide() string`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *QueryAllAlgoOrdersResponseInner) GetPositionSideOk() (*string, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *QueryAllAlgoOrdersResponseInner) SetPositionSide(v string)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *QueryAllAlgoOrdersResponseInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *QueryAllAlgoOrdersResponseInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *QueryAllAlgoOrdersResponseInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *QueryAllAlgoOrdersResponseInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *QueryAllAlgoOrdersResponseInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QueryAllAlgoOrdersResponseInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QueryAllAlgoOrdersResponseInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QueryAllAlgoOrdersResponseInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetAlgoStatus

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoStatus() string`

GetAlgoStatus returns the AlgoStatus field if non-nil, zero value otherwise.

### GetAlgoStatusOk

`func (o *QueryAllAlgoOrdersResponseInner) GetAlgoStatusOk() (*string, bool)`

GetAlgoStatusOk returns a tuple with the AlgoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoStatus

`func (o *QueryAllAlgoOrdersResponseInner) SetAlgoStatus(v string)`

SetAlgoStatus sets AlgoStatus field to given value.

### HasAlgoStatus

`func (o *QueryAllAlgoOrdersResponseInner) HasAlgoStatus() bool`

HasAlgoStatus returns a boolean if a field has been set.

### GetActualOrderId

`func (o *QueryAllAlgoOrdersResponseInner) GetActualOrderId() string`

GetActualOrderId returns the ActualOrderId field if non-nil, zero value otherwise.

### GetActualOrderIdOk

`func (o *QueryAllAlgoOrdersResponseInner) GetActualOrderIdOk() (*string, bool)`

GetActualOrderIdOk returns a tuple with the ActualOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualOrderId

`func (o *QueryAllAlgoOrdersResponseInner) SetActualOrderId(v string)`

SetActualOrderId sets ActualOrderId field to given value.

### HasActualOrderId

`func (o *QueryAllAlgoOrdersResponseInner) HasActualOrderId() bool`

HasActualOrderId returns a boolean if a field has been set.

### GetActualPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetActualPrice() string`

GetActualPrice returns the ActualPrice field if non-nil, zero value otherwise.

### GetActualPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetActualPriceOk() (*string, bool)`

GetActualPriceOk returns a tuple with the ActualPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetActualPrice(v string)`

SetActualPrice sets ActualPrice field to given value.

### HasActualPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasActualPrice() bool`

HasActualPrice returns a boolean if a field has been set.

### GetTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetTriggerPrice() string`

GetTriggerPrice returns the TriggerPrice field if non-nil, zero value otherwise.

### GetTriggerPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTriggerPriceOk() (*string, bool)`

GetTriggerPriceOk returns a tuple with the TriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetTriggerPrice(v string)`

SetTriggerPrice sets TriggerPrice field to given value.

### HasTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasTriggerPrice() bool`

HasTriggerPrice returns a boolean if a field has been set.

### GetPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetIcebergQuantity

`func (o *QueryAllAlgoOrdersResponseInner) GetIcebergQuantity() string`

GetIcebergQuantity returns the IcebergQuantity field if non-nil, zero value otherwise.

### GetIcebergQuantityOk

`func (o *QueryAllAlgoOrdersResponseInner) GetIcebergQuantityOk() (*string, bool)`

GetIcebergQuantityOk returns a tuple with the IcebergQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcebergQuantity

`func (o *QueryAllAlgoOrdersResponseInner) SetIcebergQuantity(v string)`

SetIcebergQuantity sets IcebergQuantity field to given value.

### HasIcebergQuantity

`func (o *QueryAllAlgoOrdersResponseInner) HasIcebergQuantity() bool`

HasIcebergQuantity returns a boolean if a field has been set.

### GetTpTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetTpTriggerPrice() string`

GetTpTriggerPrice returns the TpTriggerPrice field if non-nil, zero value otherwise.

### GetTpTriggerPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTpTriggerPriceOk() (*string, bool)`

GetTpTriggerPriceOk returns a tuple with the TpTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetTpTriggerPrice(v string)`

SetTpTriggerPrice sets TpTriggerPrice field to given value.

### HasTpTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasTpTriggerPrice() bool`

HasTpTriggerPrice returns a boolean if a field has been set.

### GetTpPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetTpPrice() string`

GetTpPrice returns the TpPrice field if non-nil, zero value otherwise.

### GetTpPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTpPriceOk() (*string, bool)`

GetTpPriceOk returns a tuple with the TpPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetTpPrice(v string)`

SetTpPrice sets TpPrice field to given value.

### HasTpPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasTpPrice() bool`

HasTpPrice returns a boolean if a field has been set.

### GetSlTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetSlTriggerPrice() string`

GetSlTriggerPrice returns the SlTriggerPrice field if non-nil, zero value otherwise.

### GetSlTriggerPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetSlTriggerPriceOk() (*string, bool)`

GetSlTriggerPriceOk returns a tuple with the SlTriggerPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetSlTriggerPrice(v string)`

SetSlTriggerPrice sets SlTriggerPrice field to given value.

### HasSlTriggerPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasSlTriggerPrice() bool`

HasSlTriggerPrice returns a boolean if a field has been set.

### GetSlPrice

`func (o *QueryAllAlgoOrdersResponseInner) GetSlPrice() string`

GetSlPrice returns the SlPrice field if non-nil, zero value otherwise.

### GetSlPriceOk

`func (o *QueryAllAlgoOrdersResponseInner) GetSlPriceOk() (*string, bool)`

GetSlPriceOk returns a tuple with the SlPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlPrice

`func (o *QueryAllAlgoOrdersResponseInner) SetSlPrice(v string)`

SetSlPrice sets SlPrice field to given value.

### HasSlPrice

`func (o *QueryAllAlgoOrdersResponseInner) HasSlPrice() bool`

HasSlPrice returns a boolean if a field has been set.

### GetTpOrderType

`func (o *QueryAllAlgoOrdersResponseInner) GetTpOrderType() string`

GetTpOrderType returns the TpOrderType field if non-nil, zero value otherwise.

### GetTpOrderTypeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTpOrderTypeOk() (*string, bool)`

GetTpOrderTypeOk returns a tuple with the TpOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpOrderType

`func (o *QueryAllAlgoOrdersResponseInner) SetTpOrderType(v string)`

SetTpOrderType sets TpOrderType field to given value.

### HasTpOrderType

`func (o *QueryAllAlgoOrdersResponseInner) HasTpOrderType() bool`

HasTpOrderType returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *QueryAllAlgoOrdersResponseInner) GetSelfTradePreventionMode() string`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetSelfTradePreventionModeOk() (*string, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *QueryAllAlgoOrdersResponseInner) SetSelfTradePreventionMode(v string)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *QueryAllAlgoOrdersResponseInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetWorkingType

`func (o *QueryAllAlgoOrdersResponseInner) GetWorkingType() string`

GetWorkingType returns the WorkingType field if non-nil, zero value otherwise.

### GetWorkingTypeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetWorkingTypeOk() (*string, bool)`

GetWorkingTypeOk returns a tuple with the WorkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingType

`func (o *QueryAllAlgoOrdersResponseInner) SetWorkingType(v string)`

SetWorkingType sets WorkingType field to given value.

### HasWorkingType

`func (o *QueryAllAlgoOrdersResponseInner) HasWorkingType() bool`

HasWorkingType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *QueryAllAlgoOrdersResponseInner) GetPriceMatch() string`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *QueryAllAlgoOrdersResponseInner) GetPriceMatchOk() (*string, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *QueryAllAlgoOrdersResponseInner) SetPriceMatch(v string)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *QueryAllAlgoOrdersResponseInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetClosePosition

`func (o *QueryAllAlgoOrdersResponseInner) GetClosePosition() bool`

GetClosePosition returns the ClosePosition field if non-nil, zero value otherwise.

### GetClosePositionOk

`func (o *QueryAllAlgoOrdersResponseInner) GetClosePositionOk() (*bool, bool)`

GetClosePositionOk returns a tuple with the ClosePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePosition

`func (o *QueryAllAlgoOrdersResponseInner) SetClosePosition(v bool)`

SetClosePosition sets ClosePosition field to given value.

### HasClosePosition

`func (o *QueryAllAlgoOrdersResponseInner) HasClosePosition() bool`

HasClosePosition returns a boolean if a field has been set.

### GetPriceProtect

`func (o *QueryAllAlgoOrdersResponseInner) GetPriceProtect() bool`

GetPriceProtect returns the PriceProtect field if non-nil, zero value otherwise.

### GetPriceProtectOk

`func (o *QueryAllAlgoOrdersResponseInner) GetPriceProtectOk() (*bool, bool)`

GetPriceProtectOk returns a tuple with the PriceProtect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceProtect

`func (o *QueryAllAlgoOrdersResponseInner) SetPriceProtect(v bool)`

SetPriceProtect sets PriceProtect field to given value.

### HasPriceProtect

`func (o *QueryAllAlgoOrdersResponseInner) HasPriceProtect() bool`

HasPriceProtect returns a boolean if a field has been set.

### GetReduceOnly

`func (o *QueryAllAlgoOrdersResponseInner) GetReduceOnly() bool`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *QueryAllAlgoOrdersResponseInner) GetReduceOnlyOk() (*bool, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *QueryAllAlgoOrdersResponseInner) SetReduceOnly(v bool)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *QueryAllAlgoOrdersResponseInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryAllAlgoOrdersResponseInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryAllAlgoOrdersResponseInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryAllAlgoOrdersResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *QueryAllAlgoOrdersResponseInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *QueryAllAlgoOrdersResponseInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *QueryAllAlgoOrdersResponseInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetTriggerTime

`func (o *QueryAllAlgoOrdersResponseInner) GetTriggerTime() int64`

GetTriggerTime returns the TriggerTime field if non-nil, zero value otherwise.

### GetTriggerTimeOk

`func (o *QueryAllAlgoOrdersResponseInner) GetTriggerTimeOk() (*int64, bool)`

GetTriggerTimeOk returns a tuple with the TriggerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTime

`func (o *QueryAllAlgoOrdersResponseInner) SetTriggerTime(v int64)`

SetTriggerTime sets TriggerTime field to given value.

### HasTriggerTime

`func (o *QueryAllAlgoOrdersResponseInner) HasTriggerTime() bool`

HasTriggerTime returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *QueryAllAlgoOrdersResponseInner) GetGoodTillDate() int64`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *QueryAllAlgoOrdersResponseInner) GetGoodTillDateOk() (*int64, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *QueryAllAlgoOrdersResponseInner) SetGoodTillDate(v int64)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *QueryAllAlgoOrdersResponseInner) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


