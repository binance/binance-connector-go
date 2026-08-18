# GetOtcBlocktradeEventsResponseEventsInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**BlocktradeId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**MarketId** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Fee** | Pointer to **string** |  | [optional] 
**EffectiveFee** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewGetOtcBlocktradeEventsResponseEventsInner

`func NewGetOtcBlocktradeEventsResponseEventsInner() *GetOtcBlocktradeEventsResponseEventsInner`

NewGetOtcBlocktradeEventsResponseEventsInner instantiates a new GetOtcBlocktradeEventsResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOtcBlocktradeEventsResponseEventsInnerWithDefaults

`func NewGetOtcBlocktradeEventsResponseEventsInnerWithDefaults() *GetOtcBlocktradeEventsResponseEventsInner`

NewGetOtcBlocktradeEventsResponseEventsInnerWithDefaults instantiates a new GetOtcBlocktradeEventsResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetBlocktradeId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetBlocktradeId() string`

GetBlocktradeId returns the BlocktradeId field if non-nil, zero value otherwise.

### GetBlocktradeIdOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetBlocktradeIdOk() (*string, bool)`

GetBlocktradeIdOk returns a tuple with the BlocktradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocktradeId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetBlocktradeId(v string)`

SetBlocktradeId sets BlocktradeId field to given value.

### HasBlocktradeId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasBlocktradeId() bool`

HasBlocktradeId returns a boolean if a field has been set.

### GetOrderId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMarketId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetMarketId() int32`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetMarketIdOk() (*int32, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetMarketId(v int32)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetAmount

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetFee(v string)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetEffectiveFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEffectiveFee() string`

GetEffectiveFee returns the EffectiveFee field if non-nil, zero value otherwise.

### GetEffectiveFeeOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetEffectiveFeeOk() (*string, bool)`

GetEffectiveFeeOk returns a tuple with the EffectiveFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetEffectiveFee(v string)`

SetEffectiveFee sets EffectiveFee field to given value.

### HasEffectiveFee

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasEffectiveFee() bool`

HasEffectiveFee returns a boolean if a field has been set.

### GetPrice

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTransactionHash

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### SetTransactionHashNil

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetTransactionHashNil(b bool)`

 SetTransactionHashNil sets the value for TransactionHash to be an explicit nil

### UnsetTransactionHash
`func (o *GetOtcBlocktradeEventsResponseEventsInner) UnsetTransactionHash()`

UnsetTransactionHash ensures that no value is present for TransactionHash, not even an explicit nil
### GetReason

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetOtcBlocktradeEventsResponseEventsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetIsMaker

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetOtcBlocktradeEventsResponseEventsInner) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetOtcBlocktradeEventsResponseEventsInner) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetOtcBlocktradeEventsResponseEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


