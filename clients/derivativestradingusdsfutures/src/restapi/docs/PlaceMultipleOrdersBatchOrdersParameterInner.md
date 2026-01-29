# PlaceMultipleOrdersBatchOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**NewAlgoOrderSideParameter**](NewAlgoOrderSideParameter.md) |  | [optional] 
**PositionSide** | Pointer to [**NewAlgoOrderPositionSideParameter**](NewAlgoOrderPositionSideParameter.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**NewAlgoOrderTimeInForceParameter**](NewAlgoOrderTimeInForceParameter.md) |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**ReduceOnly** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**NewClientOrderId** | Pointer to **string** |  | [optional] 
**NewOrderRespType** | Pointer to [**NewAlgoOrderNewOrderRespTypeParameter**](NewAlgoOrderNewOrderRespTypeParameter.md) |  | [optional] 
**PriceMatch** | Pointer to [**NewAlgoOrderPriceMatchParameter**](NewAlgoOrderPriceMatchParameter.md) |  | [optional] 
**SelfTradePreventionMode** | Pointer to [**NewAlgoOrderSelfTradePreventionModeParameter**](NewAlgoOrderSelfTradePreventionModeParameter.md) |  | [optional] 
**GoodTillDate** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaceMultipleOrdersBatchOrdersParameterInner

`func NewPlaceMultipleOrdersBatchOrdersParameterInner() *PlaceMultipleOrdersBatchOrdersParameterInner`

NewPlaceMultipleOrdersBatchOrdersParameterInner instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults

`func NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersBatchOrdersParameterInner`

NewPlaceMultipleOrdersBatchOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersBatchOrdersParameterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSide() NewAlgoOrderSideParameter`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSideOk() (*NewAlgoOrderSideParameter, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSide(v NewAlgoOrderSideParameter)`

SetSide sets Side field to given value.

### HasSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSide() NewAlgoOrderPositionSideParameter`

GetPositionSide returns the PositionSide field if non-nil, zero value otherwise.

### GetPositionSideOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPositionSideOk() (*NewAlgoOrderPositionSideParameter, bool)`

GetPositionSideOk returns a tuple with the PositionSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPositionSide(v NewAlgoOrderPositionSideParameter)`

SetPositionSide sets PositionSide field to given value.

### HasPositionSide

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPositionSide() bool`

HasPositionSide returns a boolean if a field has been set.

### GetType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForce() NewAlgoOrderTimeInForceParameter`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetTimeInForceOk() (*NewAlgoOrderTimeInForceParameter, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetTimeInForce(v NewAlgoOrderTimeInForceParameter)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespType() NewAlgoOrderNewOrderRespTypeParameter`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetNewOrderRespTypeOk() (*NewAlgoOrderNewOrderRespTypeParameter, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetNewOrderRespType(v NewAlgoOrderNewOrderRespTypeParameter)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatch() NewAlgoOrderPriceMatchParameter`

GetPriceMatch returns the PriceMatch field if non-nil, zero value otherwise.

### GetPriceMatchOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetPriceMatchOk() (*NewAlgoOrderPriceMatchParameter, bool)`

GetPriceMatchOk returns a tuple with the PriceMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetPriceMatch(v NewAlgoOrderPriceMatchParameter)`

SetPriceMatch sets PriceMatch field to given value.

### HasPriceMatch

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasPriceMatch() bool`

HasPriceMatch returns a boolean if a field has been set.

### GetSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionMode() NewAlgoOrderSelfTradePreventionModeParameter`

GetSelfTradePreventionMode returns the SelfTradePreventionMode field if non-nil, zero value otherwise.

### GetSelfTradePreventionModeOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetSelfTradePreventionModeOk() (*NewAlgoOrderSelfTradePreventionModeParameter, bool)`

GetSelfTradePreventionModeOk returns a tuple with the SelfTradePreventionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetSelfTradePreventionMode(v NewAlgoOrderSelfTradePreventionModeParameter)`

SetSelfTradePreventionMode sets SelfTradePreventionMode field to given value.

### HasSelfTradePreventionMode

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasSelfTradePreventionMode() bool`

HasSelfTradePreventionMode returns a boolean if a field has been set.

### GetGoodTillDate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetGoodTillDate() string`

GetGoodTillDate returns the GoodTillDate field if non-nil, zero value otherwise.

### GetGoodTillDateOk

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) GetGoodTillDateOk() (*string, bool)`

GetGoodTillDateOk returns a tuple with the GoodTillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoodTillDate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) SetGoodTillDate(v string)`

SetGoodTillDate sets GoodTillDate field to given value.

### HasGoodTillDate

`func (o *PlaceMultipleOrdersBatchOrdersParameterInner) HasGoodTillDate() bool`

HasGoodTillDate returns a boolean if a field has been set.


[[Back to README]](../README.md)


