# MyAllocationsResponseInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**AllocationId** | Pointer to **int64** |  | [optional] 
**AllocationType** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderListId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Qty** | Pointer to **string** |  | [optional] 
**QuoteQty** | Pointer to **string** |  | [optional] 
**Commission** | Pointer to **string** |  | [optional] 
**CommissionAsset** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 
**IsBuyer** | Pointer to **bool** |  | [optional] 
**IsMaker** | Pointer to **bool** |  | [optional] 
**IsAllocator** | Pointer to **bool** |  | [optional] 

## Methods

### NewMyAllocationsResponseInner

`func NewMyAllocationsResponseInner() *MyAllocationsResponseInner`

NewMyAllocationsResponseInner instantiates a new MyAllocationsResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyAllocationsResponseInnerWithDefaults

`func NewMyAllocationsResponseInnerWithDefaults() *MyAllocationsResponseInner`

NewMyAllocationsResponseInnerWithDefaults instantiates a new MyAllocationsResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MyAllocationsResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MyAllocationsResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MyAllocationsResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MyAllocationsResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAllocationId

`func (o *MyAllocationsResponseInner) GetAllocationId() int64`

GetAllocationId returns the AllocationId field if non-nil, zero value otherwise.

### GetAllocationIdOk

`func (o *MyAllocationsResponseInner) GetAllocationIdOk() (*int64, bool)`

GetAllocationIdOk returns a tuple with the AllocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationId

`func (o *MyAllocationsResponseInner) SetAllocationId(v int64)`

SetAllocationId sets AllocationId field to given value.

### HasAllocationId

`func (o *MyAllocationsResponseInner) HasAllocationId() bool`

HasAllocationId returns a boolean if a field has been set.

### GetAllocationType

`func (o *MyAllocationsResponseInner) GetAllocationType() string`

GetAllocationType returns the AllocationType field if non-nil, zero value otherwise.

### GetAllocationTypeOk

`func (o *MyAllocationsResponseInner) GetAllocationTypeOk() (*string, bool)`

GetAllocationTypeOk returns a tuple with the AllocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationType

`func (o *MyAllocationsResponseInner) SetAllocationType(v string)`

SetAllocationType sets AllocationType field to given value.

### HasAllocationType

`func (o *MyAllocationsResponseInner) HasAllocationType() bool`

HasAllocationType returns a boolean if a field has been set.

### GetOrderId

`func (o *MyAllocationsResponseInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MyAllocationsResponseInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MyAllocationsResponseInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MyAllocationsResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderListId

`func (o *MyAllocationsResponseInner) GetOrderListId() int64`

GetOrderListId returns the OrderListId field if non-nil, zero value otherwise.

### GetOrderListIdOk

`func (o *MyAllocationsResponseInner) GetOrderListIdOk() (*int64, bool)`

GetOrderListIdOk returns a tuple with the OrderListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderListId

`func (o *MyAllocationsResponseInner) SetOrderListId(v int64)`

SetOrderListId sets OrderListId field to given value.

### HasOrderListId

`func (o *MyAllocationsResponseInner) HasOrderListId() bool`

HasOrderListId returns a boolean if a field has been set.

### GetPrice

`func (o *MyAllocationsResponseInner) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *MyAllocationsResponseInner) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *MyAllocationsResponseInner) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *MyAllocationsResponseInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQty

`func (o *MyAllocationsResponseInner) GetQty() string`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *MyAllocationsResponseInner) GetQtyOk() (*string, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *MyAllocationsResponseInner) SetQty(v string)`

SetQty sets Qty field to given value.

### HasQty

`func (o *MyAllocationsResponseInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetQuoteQty

`func (o *MyAllocationsResponseInner) GetQuoteQty() string`

GetQuoteQty returns the QuoteQty field if non-nil, zero value otherwise.

### GetQuoteQtyOk

`func (o *MyAllocationsResponseInner) GetQuoteQtyOk() (*string, bool)`

GetQuoteQtyOk returns a tuple with the QuoteQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteQty

`func (o *MyAllocationsResponseInner) SetQuoteQty(v string)`

SetQuoteQty sets QuoteQty field to given value.

### HasQuoteQty

`func (o *MyAllocationsResponseInner) HasQuoteQty() bool`

HasQuoteQty returns a boolean if a field has been set.

### GetCommission

`func (o *MyAllocationsResponseInner) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *MyAllocationsResponseInner) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *MyAllocationsResponseInner) SetCommission(v string)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *MyAllocationsResponseInner) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetCommissionAsset

`func (o *MyAllocationsResponseInner) GetCommissionAsset() string`

GetCommissionAsset returns the CommissionAsset field if non-nil, zero value otherwise.

### GetCommissionAssetOk

`func (o *MyAllocationsResponseInner) GetCommissionAssetOk() (*string, bool)`

GetCommissionAssetOk returns a tuple with the CommissionAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAsset

`func (o *MyAllocationsResponseInner) SetCommissionAsset(v string)`

SetCommissionAsset sets CommissionAsset field to given value.

### HasCommissionAsset

`func (o *MyAllocationsResponseInner) HasCommissionAsset() bool`

HasCommissionAsset returns a boolean if a field has been set.

### GetTime

`func (o *MyAllocationsResponseInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *MyAllocationsResponseInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *MyAllocationsResponseInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *MyAllocationsResponseInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetIsBuyer

`func (o *MyAllocationsResponseInner) GetIsBuyer() bool`

GetIsBuyer returns the IsBuyer field if non-nil, zero value otherwise.

### GetIsBuyerOk

`func (o *MyAllocationsResponseInner) GetIsBuyerOk() (*bool, bool)`

GetIsBuyerOk returns a tuple with the IsBuyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuyer

`func (o *MyAllocationsResponseInner) SetIsBuyer(v bool)`

SetIsBuyer sets IsBuyer field to given value.

### HasIsBuyer

`func (o *MyAllocationsResponseInner) HasIsBuyer() bool`

HasIsBuyer returns a boolean if a field has been set.

### GetIsMaker

`func (o *MyAllocationsResponseInner) GetIsMaker() bool`

GetIsMaker returns the IsMaker field if non-nil, zero value otherwise.

### GetIsMakerOk

`func (o *MyAllocationsResponseInner) GetIsMakerOk() (*bool, bool)`

GetIsMakerOk returns a tuple with the IsMaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaker

`func (o *MyAllocationsResponseInner) SetIsMaker(v bool)`

SetIsMaker sets IsMaker field to given value.

### HasIsMaker

`func (o *MyAllocationsResponseInner) HasIsMaker() bool`

HasIsMaker returns a boolean if a field has been set.

### GetIsAllocator

`func (o *MyAllocationsResponseInner) GetIsAllocator() bool`

GetIsAllocator returns the IsAllocator field if non-nil, zero value otherwise.

### GetIsAllocatorOk

`func (o *MyAllocationsResponseInner) GetIsAllocatorOk() (*bool, bool)`

GetIsAllocatorOk returns a tuple with the IsAllocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocator

`func (o *MyAllocationsResponseInner) SetIsAllocator(v bool)`

SetIsAllocator sets IsAllocator field to given value.

### HasIsAllocator

`func (o *MyAllocationsResponseInner) HasIsAllocator() bool`

HasIsAllocator returns a boolean if a field has been set.


[[Back to README]](../README.md)


