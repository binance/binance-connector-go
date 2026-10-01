# QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | symbol. | [optional] 
**OrderId** | Pointer to **NullableInt64** | order Id. Returns &#x60;null&#x60; if the working order is not filled yet. | [optional] 
**Status** | Pointer to **string** | status. | [optional] 
**ClientOrderId** | Pointer to **string** | client Order Id. | [optional] 

## Methods

### NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner

`func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner() *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner`

NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInnerWithDefaults

`func NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInnerWithDefaults() *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner`

NewQueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInnerWithDefaults instantiates a new QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *QueryMarginAccountsOpenOtootocoOrderListsResponseInnerOrdersInner) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.


[[Back to README]](../README.md)


