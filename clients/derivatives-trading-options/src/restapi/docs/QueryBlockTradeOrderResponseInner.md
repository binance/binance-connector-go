# QueryBlockTradeOrderResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]ExtendBlockTradeOrderResponseLegsInner**](ExtendBlockTradeOrderResponseLegsInner.md) |  | [optional] 

## Methods

### NewQueryBlockTradeOrderResponseInner

`func NewQueryBlockTradeOrderResponseInner() *QueryBlockTradeOrderResponseInner`

NewQueryBlockTradeOrderResponseInner instantiates a new QueryBlockTradeOrderResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBlockTradeOrderResponseInnerWithDefaults

`func NewQueryBlockTradeOrderResponseInnerWithDefaults() *QueryBlockTradeOrderResponseInner`

NewQueryBlockTradeOrderResponseInnerWithDefaults instantiates a new QueryBlockTradeOrderResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *QueryBlockTradeOrderResponseInner) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *QueryBlockTradeOrderResponseInner) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *QueryBlockTradeOrderResponseInner) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *QueryBlockTradeOrderResponseInner) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetExpireTime

`func (o *QueryBlockTradeOrderResponseInner) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *QueryBlockTradeOrderResponseInner) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *QueryBlockTradeOrderResponseInner) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *QueryBlockTradeOrderResponseInner) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLiquidity

`func (o *QueryBlockTradeOrderResponseInner) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *QueryBlockTradeOrderResponseInner) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *QueryBlockTradeOrderResponseInner) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *QueryBlockTradeOrderResponseInner) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *QueryBlockTradeOrderResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryBlockTradeOrderResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryBlockTradeOrderResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryBlockTradeOrderResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryBlockTradeOrderResponseInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryBlockTradeOrderResponseInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryBlockTradeOrderResponseInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryBlockTradeOrderResponseInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLegs

`func (o *QueryBlockTradeOrderResponseInner) GetLegs() []ExtendBlockTradeOrderResponseLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *QueryBlockTradeOrderResponseInner) GetLegsOk() (*[]ExtendBlockTradeOrderResponseLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *QueryBlockTradeOrderResponseInner) SetLegs(v []ExtendBlockTradeOrderResponseLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *QueryBlockTradeOrderResponseInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.


[[Back to README]](../README.md)


