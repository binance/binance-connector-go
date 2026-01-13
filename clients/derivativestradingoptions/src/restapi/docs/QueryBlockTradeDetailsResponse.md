# QueryBlockTradeDetailsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]QueryBlockTradeDetailsResponseLegsInner**](QueryBlockTradeDetailsResponseLegsInner.md) |  | [optional] 

## Methods

### NewQueryBlockTradeDetailsResponse

`func NewQueryBlockTradeDetailsResponse() *QueryBlockTradeDetailsResponse`

NewQueryBlockTradeDetailsResponse instantiates a new QueryBlockTradeDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryBlockTradeDetailsResponseWithDefaults

`func NewQueryBlockTradeDetailsResponseWithDefaults() *QueryBlockTradeDetailsResponse`

NewQueryBlockTradeDetailsResponseWithDefaults instantiates a new QueryBlockTradeDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *QueryBlockTradeDetailsResponse) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *QueryBlockTradeDetailsResponse) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *QueryBlockTradeDetailsResponse) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *QueryBlockTradeDetailsResponse) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetExpireTime

`func (o *QueryBlockTradeDetailsResponse) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *QueryBlockTradeDetailsResponse) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *QueryBlockTradeDetailsResponse) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *QueryBlockTradeDetailsResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLiquidity

`func (o *QueryBlockTradeDetailsResponse) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *QueryBlockTradeDetailsResponse) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *QueryBlockTradeDetailsResponse) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *QueryBlockTradeDetailsResponse) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *QueryBlockTradeDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryBlockTradeDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryBlockTradeDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryBlockTradeDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryBlockTradeDetailsResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryBlockTradeDetailsResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryBlockTradeDetailsResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryBlockTradeDetailsResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLegs

`func (o *QueryBlockTradeDetailsResponse) GetLegs() []QueryBlockTradeDetailsResponseLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *QueryBlockTradeDetailsResponse) GetLegsOk() (*[]QueryBlockTradeDetailsResponseLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *QueryBlockTradeDetailsResponse) SetLegs(v []QueryBlockTradeDetailsResponseLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *QueryBlockTradeDetailsResponse) HasLegs() bool`

HasLegs returns a boolean if a field has been set.


[[Back to README]](../README.md)


