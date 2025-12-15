# PlaceMultipleOrdersOrdersParameterInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**PlaceMultipleOrdersOrdersParameterInnerSide**](PlaceMultipleOrdersOrdersParameterInnerSide.md) |  | [optional] 
**Type** | Pointer to [**PlaceMultipleOrdersOrdersParameterInnerType**](PlaceMultipleOrdersOrdersParameterInnerType.md) |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**PlaceMultipleOrdersOrdersParameterInnerTimeInForce**](PlaceMultipleOrdersOrdersParameterInnerTimeInForce.md) |  | [optional] 
**ReduceOnly** | Pointer to **string** |  | [optional] 
**PostOnly** | Pointer to **string** |  | [optional] 
**NewOrderRespType** | Pointer to [**PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType**](PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType.md) |  | [optional] 
**ClientOrderId** | Pointer to **string** |  | [optional] 
**IsMmp** | Pointer to **string** |  | [optional] 

## Methods

### NewPlaceMultipleOrdersOrdersParameterInner

`func NewPlaceMultipleOrdersOrdersParameterInner() *PlaceMultipleOrdersOrdersParameterInner`

NewPlaceMultipleOrdersOrdersParameterInner instantiates a new PlaceMultipleOrdersOrdersParameterInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults

`func NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults() *PlaceMultipleOrdersOrdersParameterInner`

NewPlaceMultipleOrdersOrdersParameterInnerWithDefaults instantiates a new PlaceMultipleOrdersOrdersParameterInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetSide() PlaceMultipleOrdersOrdersParameterInnerSide`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetSideOk() (*PlaceMultipleOrdersOrdersParameterInnerSide, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetSide(v PlaceMultipleOrdersOrdersParameterInnerSide)`

SetSide sets Side field to given value.

### HasSide

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetType

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetType() PlaceMultipleOrdersOrdersParameterInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetTypeOk() (*PlaceMultipleOrdersOrdersParameterInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetType(v PlaceMultipleOrdersOrdersParameterInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTimeInForce

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetTimeInForce() PlaceMultipleOrdersOrdersParameterInnerTimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetTimeInForceOk() (*PlaceMultipleOrdersOrdersParameterInnerTimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetTimeInForce(v PlaceMultipleOrdersOrdersParameterInnerTimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetReduceOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnly() string`

GetReduceOnly returns the ReduceOnly field if non-nil, zero value otherwise.

### GetReduceOnlyOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetReduceOnlyOk() (*string, bool)`

GetReduceOnlyOk returns a tuple with the ReduceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetReduceOnly(v string)`

SetReduceOnly sets ReduceOnly field to given value.

### HasReduceOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasReduceOnly() bool`

HasReduceOnly returns a boolean if a field has been set.

### GetPostOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnly() string`

GetPostOnly returns the PostOnly field if non-nil, zero value otherwise.

### GetPostOnlyOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetPostOnlyOk() (*string, bool)`

GetPostOnlyOk returns a tuple with the PostOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetPostOnly(v string)`

SetPostOnly sets PostOnly field to given value.

### HasPostOnly

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasPostOnly() bool`

HasPostOnly returns a boolean if a field has been set.

### GetNewOrderRespType

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetNewOrderRespType() PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType`

GetNewOrderRespType returns the NewOrderRespType field if non-nil, zero value otherwise.

### GetNewOrderRespTypeOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetNewOrderRespTypeOk() (*PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType, bool)`

GetNewOrderRespTypeOk returns a tuple with the NewOrderRespType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderRespType

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetNewOrderRespType(v PlaceMultipleOrdersOrdersParameterInnerNewOrderRespType)`

SetNewOrderRespType sets NewOrderRespType field to given value.

### HasNewOrderRespType

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasNewOrderRespType() bool`

HasNewOrderRespType returns a boolean if a field has been set.

### GetClientOrderId

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetIsMmp

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmp() string`

GetIsMmp returns the IsMmp field if non-nil, zero value otherwise.

### GetIsMmpOk

`func (o *PlaceMultipleOrdersOrdersParameterInner) GetIsMmpOk() (*string, bool)`

GetIsMmpOk returns a tuple with the IsMmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMmp

`func (o *PlaceMultipleOrdersOrdersParameterInner) SetIsMmp(v string)`

SetIsMmp sets IsMmp field to given value.

### HasIsMmp

`func (o *PlaceMultipleOrdersOrdersParameterInner) HasIsMmp() bool`

HasIsMmp returns a boolean if a field has been set.


[[Back to README]](../README.md)


