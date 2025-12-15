# OrderStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**FromAsset** | Pointer to **string** |  | [optional] 
**FromAmount** | Pointer to **string** |  | [optional] 
**ToAsset** | Pointer to **string** |  | [optional] 
**ToAmount** | Pointer to **string** |  | [optional] 
**Ratio** | Pointer to **string** |  | [optional] 
**InverseRatio** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrderStatusResponse

`func NewOrderStatusResponse() *OrderStatusResponse`

NewOrderStatusResponse instantiates a new OrderStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusResponseWithDefaults

`func NewOrderStatusResponseWithDefaults() *OrderStatusResponse`

NewOrderStatusResponseWithDefaults instantiates a new OrderStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *OrderStatusResponse) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderStatusResponse) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderStatusResponse) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderStatusResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderStatusResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderStatusResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderStatusResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderStatusResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetFromAsset

`func (o *OrderStatusResponse) GetFromAsset() string`

GetFromAsset returns the FromAsset field if non-nil, zero value otherwise.

### GetFromAssetOk

`func (o *OrderStatusResponse) GetFromAssetOk() (*string, bool)`

GetFromAssetOk returns a tuple with the FromAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAsset

`func (o *OrderStatusResponse) SetFromAsset(v string)`

SetFromAsset sets FromAsset field to given value.

### HasFromAsset

`func (o *OrderStatusResponse) HasFromAsset() bool`

HasFromAsset returns a boolean if a field has been set.

### GetFromAmount

`func (o *OrderStatusResponse) GetFromAmount() string`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *OrderStatusResponse) GetFromAmountOk() (*string, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *OrderStatusResponse) SetFromAmount(v string)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *OrderStatusResponse) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetToAsset

`func (o *OrderStatusResponse) GetToAsset() string`

GetToAsset returns the ToAsset field if non-nil, zero value otherwise.

### GetToAssetOk

`func (o *OrderStatusResponse) GetToAssetOk() (*string, bool)`

GetToAssetOk returns a tuple with the ToAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAsset

`func (o *OrderStatusResponse) SetToAsset(v string)`

SetToAsset sets ToAsset field to given value.

### HasToAsset

`func (o *OrderStatusResponse) HasToAsset() bool`

HasToAsset returns a boolean if a field has been set.

### GetToAmount

`func (o *OrderStatusResponse) GetToAmount() string`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *OrderStatusResponse) GetToAmountOk() (*string, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *OrderStatusResponse) SetToAmount(v string)`

SetToAmount sets ToAmount field to given value.

### HasToAmount

`func (o *OrderStatusResponse) HasToAmount() bool`

HasToAmount returns a boolean if a field has been set.

### GetRatio

`func (o *OrderStatusResponse) GetRatio() string`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *OrderStatusResponse) GetRatioOk() (*string, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *OrderStatusResponse) SetRatio(v string)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *OrderStatusResponse) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetInverseRatio

`func (o *OrderStatusResponse) GetInverseRatio() string`

GetInverseRatio returns the InverseRatio field if non-nil, zero value otherwise.

### GetInverseRatioOk

`func (o *OrderStatusResponse) GetInverseRatioOk() (*string, bool)`

GetInverseRatioOk returns a tuple with the InverseRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInverseRatio

`func (o *OrderStatusResponse) SetInverseRatio(v string)`

SetInverseRatio sets InverseRatio field to given value.

### HasInverseRatio

`func (o *OrderStatusResponse) HasInverseRatio() bool`

HasInverseRatio returns a boolean if a field has been set.

### GetCreateTime

`func (o *OrderStatusResponse) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OrderStatusResponse) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OrderStatusResponse) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OrderStatusResponse) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


