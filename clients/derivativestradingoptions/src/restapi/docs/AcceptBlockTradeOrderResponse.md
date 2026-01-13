# AcceptBlockTradeOrderResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 
**ExpireTime** | Pointer to **int64** |  | [optional] 
**Liquidity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Legs** | Pointer to [**[]AcceptBlockTradeOrderResponseLegsInner**](AcceptBlockTradeOrderResponseLegsInner.md) |  | [optional] 

## Methods

### NewAcceptBlockTradeOrderResponse

`func NewAcceptBlockTradeOrderResponse() *AcceptBlockTradeOrderResponse`

NewAcceptBlockTradeOrderResponse instantiates a new AcceptBlockTradeOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptBlockTradeOrderResponseWithDefaults

`func NewAcceptBlockTradeOrderResponseWithDefaults() *AcceptBlockTradeOrderResponse`

NewAcceptBlockTradeOrderResponseWithDefaults instantiates a new AcceptBlockTradeOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockTradeSettlementKey

`func (o *AcceptBlockTradeOrderResponse) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *AcceptBlockTradeOrderResponse) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *AcceptBlockTradeOrderResponse) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *AcceptBlockTradeOrderResponse) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.

### GetExpireTime

`func (o *AcceptBlockTradeOrderResponse) GetExpireTime() int64`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *AcceptBlockTradeOrderResponse) GetExpireTimeOk() (*int64, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *AcceptBlockTradeOrderResponse) SetExpireTime(v int64)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *AcceptBlockTradeOrderResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetLiquidity

`func (o *AcceptBlockTradeOrderResponse) GetLiquidity() string`

GetLiquidity returns the Liquidity field if non-nil, zero value otherwise.

### GetLiquidityOk

`func (o *AcceptBlockTradeOrderResponse) GetLiquidityOk() (*string, bool)`

GetLiquidityOk returns a tuple with the Liquidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidity

`func (o *AcceptBlockTradeOrderResponse) SetLiquidity(v string)`

SetLiquidity sets Liquidity field to given value.

### HasLiquidity

`func (o *AcceptBlockTradeOrderResponse) HasLiquidity() bool`

HasLiquidity returns a boolean if a field has been set.

### GetStatus

`func (o *AcceptBlockTradeOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcceptBlockTradeOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcceptBlockTradeOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcceptBlockTradeOrderResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreateTime

`func (o *AcceptBlockTradeOrderResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AcceptBlockTradeOrderResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AcceptBlockTradeOrderResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AcceptBlockTradeOrderResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetLegs

`func (o *AcceptBlockTradeOrderResponse) GetLegs() []AcceptBlockTradeOrderResponseLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *AcceptBlockTradeOrderResponse) GetLegsOk() (*[]AcceptBlockTradeOrderResponseLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *AcceptBlockTradeOrderResponse) SetLegs(v []AcceptBlockTradeOrderResponseLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *AcceptBlockTradeOrderResponse) HasLegs() bool`

HasLegs returns a boolean if a field has been set.


[[Back to README]](../README.md)


