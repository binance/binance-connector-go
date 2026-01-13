# NewBlockTradeOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Legs** | Pointer to [**[]ExtendBlockTradeOrderResponseLegsInner**](ExtendBlockTradeOrderResponseLegsInner.md) |  | [optional] 

## Methods

### NewNewBlockTradeOrderResponse

`func NewNewBlockTradeOrderResponse() *NewBlockTradeOrderResponse`

NewNewBlockTradeOrderResponse instantiates a new NewBlockTradeOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewBlockTradeOrderResponseWithDefaults

`func NewNewBlockTradeOrderResponseWithDefaults() *NewBlockTradeOrderResponse`

NewNewBlockTradeOrderResponseWithDefaults instantiates a new NewBlockTradeOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *NewBlockTradeOrderResponse) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *NewBlockTradeOrderResponse) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *NewBlockTradeOrderResponse) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *NewBlockTradeOrderResponse) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetExpireTime

`func (o *NewBlockTradeOrderResponse) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *NewBlockTradeOrderResponse) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *NewBlockTradeOrderResponse) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *NewBlockTradeOrderResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLiquidity

`func (o *NewBlockTradeOrderResponse) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *NewBlockTradeOrderResponse) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *NewBlockTradeOrderResponse) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *NewBlockTradeOrderResponse) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *NewBlockTradeOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NewBlockTradeOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NewBlockTradeOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NewBlockTradeOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLegs

`func (o *NewBlockTradeOrderResponse) GetLegs() []ExtendBlockTradeOrderResponseLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *NewBlockTradeOrderResponse) GetLegsOk() (*[]ExtendBlockTradeOrderResponseLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *NewBlockTradeOrderResponse) SetLegs(v []ExtendBlockTradeOrderResponseLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *NewBlockTradeOrderResponse) HasLegs() bool`

HasLegs returns a boolean if a field has been set.


[[Back to README]](../README.md)


