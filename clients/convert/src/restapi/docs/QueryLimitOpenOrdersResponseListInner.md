# QueryLimitOpenOrdersResponseListInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**QuoteId** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**FromAsset** | Pointer to **string** |  | [optional] 
**FromAmount** | Pointer to **string** |  | [optional] 
**ToAsset** | Pointer to **string** |  | [optional] 
**ToAmount** | Pointer to **string** |  | [optional] 
**Ratio** | Pointer to **string** |  | [optional] 
**InverseRatio** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**ExpiredTimestamp** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueryLimitOpenOrdersResponseListInner

`func NewQueryLimitOpenOrdersResponseListInner() *QueryLimitOpenOrdersResponseListInner`

NewQueryLimitOpenOrdersResponseListInner instantiates a new QueryLimitOpenOrdersResponseListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLimitOpenOrdersResponseListInnerWithDefaults

`func NewQueryLimitOpenOrdersResponseListInnerWithDefaults() *QueryLimitOpenOrdersResponseListInner`

NewQueryLimitOpenOrdersResponseListInnerWithDefaults instantiates a new QueryLimitOpenOrdersResponseListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuoteId

`func (o *QueryLimitOpenOrdersResponseListInner) GetQuoteId() string`

GetQuoteId returns the QuoteId field if non-nil, zero value otherwise.

### GetQuoteIdOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetQuoteIdOk() (*string, bool)`

GetQuoteIdOk returns a tuple with the QuoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteId

`func (o *QueryLimitOpenOrdersResponseListInner) SetQuoteId(v string)`

SetQuoteId sets QuoteId field to given value.

### HasQuoteId

`func (o *QueryLimitOpenOrdersResponseListInner) HasQuoteId() bool`

HasQuoteId returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryLimitOpenOrdersResponseListInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryLimitOpenOrdersResponseListInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryLimitOpenOrdersResponseListInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *QueryLimitOpenOrdersResponseListInner) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *QueryLimitOpenOrdersResponseListInner) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *QueryLimitOpenOrdersResponseListInner) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetFromAsset

`func (o *QueryLimitOpenOrdersResponseListInner) GetFromAsset() string`

GetFromAsset returns the FromAsset field if non-nil, zero value otherwise.

### GetFromAssetOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetFromAssetOk() (*string, bool)`

GetFromAssetOk returns a tuple with the FromAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAsset

`func (o *QueryLimitOpenOrdersResponseListInner) SetFromAsset(v string)`

SetFromAsset sets FromAsset field to given value.

### HasFromAsset

`func (o *QueryLimitOpenOrdersResponseListInner) HasFromAsset() bool`

HasFromAsset returns a boolean if a field has been set.

### GetFromAmount

`func (o *QueryLimitOpenOrdersResponseListInner) GetFromAmount() string`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetFromAmountOk() (*string, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *QueryLimitOpenOrdersResponseListInner) SetFromAmount(v string)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *QueryLimitOpenOrdersResponseListInner) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetToAsset

`func (o *QueryLimitOpenOrdersResponseListInner) GetToAsset() string`

GetToAsset returns the ToAsset field if non-nil, zero value otherwise.

### GetToAssetOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetToAssetOk() (*string, bool)`

GetToAssetOk returns a tuple with the ToAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAsset

`func (o *QueryLimitOpenOrdersResponseListInner) SetToAsset(v string)`

SetToAsset sets ToAsset field to given value.

### HasToAsset

`func (o *QueryLimitOpenOrdersResponseListInner) HasToAsset() bool`

HasToAsset returns a boolean if a field has been set.

### GetToAmount

`func (o *QueryLimitOpenOrdersResponseListInner) GetToAmount() string`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetToAmountOk() (*string, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *QueryLimitOpenOrdersResponseListInner) SetToAmount(v string)`

SetToAmount sets ToAmount field to given value.

### HasToAmount

`func (o *QueryLimitOpenOrdersResponseListInner) HasToAmount() bool`

HasToAmount returns a boolean if a field has been set.

### GetRatio

`func (o *QueryLimitOpenOrdersResponseListInner) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *QueryLimitOpenOrdersResponseListInner) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *QueryLimitOpenOrdersResponseListInner) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetInverseRatio

`func (o *QueryLimitOpenOrdersResponseListInner) GetInverseRatio() string`

GetInverseRatio returns the InverseRatio field if non-nil, zero value otherwise.

### GetInverseRatioOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetInverseRatioOk() (*string, bool)`

GetInverseRatioOk returns a tuple with the InverseRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseRatio

`func (o *QueryLimitOpenOrdersResponseListInner) SetInverseRatio(v string)`

SetInverseRatio sets InverseRatio field to given value.

### HasInverseRatio

`func (o *QueryLimitOpenOrdersResponseListInner) HasInverseRatio() bool`

HasInverseRatio returns a boolean if a field has been set.

### GetCreateTime

`func (o *QueryLimitOpenOrdersResponseListInner) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *QueryLimitOpenOrdersResponseListInner) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *QueryLimitOpenOrdersResponseListInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpiredTimestamp

`func (o *QueryLimitOpenOrdersResponseListInner) GetExpiredTimestamp() int64`

GetExpiredTimestamp returns the ExpiredTimestamp field if non-nil, zero value otherwise.

### GetExpiredTimestampOk

`func (o *QueryLimitOpenOrdersResponseListInner) GetExpiredTimestampOk() (*int64, bool)`

GetExpiredTimestampOk returns a tuple with the ExpiredTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredTimestamp

`func (o *QueryLimitOpenOrdersResponseListInner) SetExpiredTimestamp(v int64)`

SetExpiredTimestamp sets ExpiredTimestamp field to given value.

### HasExpiredTimestamp

`func (o *QueryLimitOpenOrdersResponseListInner) HasExpiredTimestamp() bool`

HasExpiredTimestamp returns a boolean if a field has been set.


[[Back to README]](../README.md)


