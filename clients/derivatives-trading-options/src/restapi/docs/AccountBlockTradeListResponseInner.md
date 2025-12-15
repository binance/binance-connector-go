# AccountBlockTradeListResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**ParentOrderId** | Pointer to **string** |  | [optional] 
**CrossType** | Pointer to **string** |  | [optional] 
**Legs** | Pointer to [**[]AccountBlockTradeListResponseInnerLegsInner**](AccountBlockTradeListResponseInnerLegsInner.md) |  | [optional] 
**BlockTradeSettlementKey** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountBlockTradeListResponseInner

`func NewAccountBlockTradeListResponseInner() *AccountBlockTradeListResponseInner`

NewAccountBlockTradeListResponseInner instantiates a new AccountBlockTradeListResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBlockTradeListResponseInnerWithDefaults

`func NewAccountBlockTradeListResponseInnerWithDefaults() *AccountBlockTradeListResponseInner`

NewAccountBlockTradeListResponseInnerWithDefaults instantiates a new AccountBlockTradeListResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentOrderId

`func (o *AccountBlockTradeListResponseInner) GetParentOrderId() string`

GetParentOrderId returns the ParentOrderId field if non-nil, zero value otherwise.

### GetParentOrderIdOk

`func (o *AccountBlockTradeListResponseInner) GetParentOrderIdOk() (*string, bool)`

GetParentOrderIdOk returns a tuple with the ParentOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrderId

`func (o *AccountBlockTradeListResponseInner) SetParentOrderId(v string)`

SetParentOrderId sets ParentOrderId field to given value.

### HasParentOrderId

`func (o *AccountBlockTradeListResponseInner) HasParentOrderId() bool`

HasParentOrderId returns a boolean if a field has been set.

### GetCrossType

`func (o *AccountBlockTradeListResponseInner) GetCrossType() string`

GetCrossType returns the CrossType field if non-nil, zero value otherwise.

### GetCrossTypeOk

`func (o *AccountBlockTradeListResponseInner) GetCrossTypeOk() (*string, bool)`

GetCrossTypeOk returns a tuple with the CrossType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossType

`func (o *AccountBlockTradeListResponseInner) SetCrossType(v string)`

SetCrossType sets CrossType field to given value.

### HasCrossType

`func (o *AccountBlockTradeListResponseInner) HasCrossType() bool`

HasCrossType returns a boolean if a field has been set.

### GetLegs

`func (o *AccountBlockTradeListResponseInner) GetLegs() []AccountBlockTradeListResponseInnerLegsInner`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *AccountBlockTradeListResponseInner) GetLegsOk() (*[]AccountBlockTradeListResponseInnerLegsInner, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *AccountBlockTradeListResponseInner) SetLegs(v []AccountBlockTradeListResponseInnerLegsInner)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *AccountBlockTradeListResponseInner) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetBlockTradeSettlementKey

`func (o *AccountBlockTradeListResponseInner) GetBlockTradeSettlementKey() string`

GetBlockTradeSettlementKey returns the BlockTradeSettlementKey field if non-nil, zero value otherwise.

### GetBlockTradeSettlementKeyOk

`func (o *AccountBlockTradeListResponseInner) GetBlockTradeSettlementKeyOk() (*string, bool)`

GetBlockTradeSettlementKeyOk returns a tuple with the BlockTradeSettlementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTradeSettlementKey

`func (o *AccountBlockTradeListResponseInner) SetBlockTradeSettlementKey(v string)`

SetBlockTradeSettlementKey sets BlockTradeSettlementKey field to given value.

### HasBlockTradeSettlementKey

`func (o *AccountBlockTradeListResponseInner) HasBlockTradeSettlementKey() bool`

HasBlockTradeSettlementKey returns a boolean if a field has been set.


[[Back to README]](../README.md)


