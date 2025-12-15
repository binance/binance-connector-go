# OrderAmendmentsResponseResultInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**ExecutionId** | Pointer to **int64** |  | [optional] 
**OrigClientOrderId** | Pointer to **string** |  | [optional] 
**NewClientOrderId** | Pointer to **string** |  | [optional] 
**OrigQty** | Pointer to **string** |  | [optional] 
**NewQty** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrderAmendmentsResponseResultInner

`func NewOrderAmendmentsResponseResultInner() *OrderAmendmentsResponseResultInner`

NewOrderAmendmentsResponseResultInner instantiates a new OrderAmendmentsResponseResultInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendmentsResponseResultInnerWithDefaults

`func NewOrderAmendmentsResponseResultInnerWithDefaults() *OrderAmendmentsResponseResultInner`

NewOrderAmendmentsResponseResultInnerWithDefaults instantiates a new OrderAmendmentsResponseResultInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *OrderAmendmentsResponseResultInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderAmendmentsResponseResultInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderAmendmentsResponseResultInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderAmendmentsResponseResultInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderAmendmentsResponseResultInner) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderAmendmentsResponseResultInner) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderAmendmentsResponseResultInner) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderAmendmentsResponseResultInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetExecutionId

`func (o *OrderAmendmentsResponseResultInner) GetExecutionId() int64`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *OrderAmendmentsResponseResultInner) GetExecutionIdOk() (*int64, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *OrderAmendmentsResponseResultInner) SetExecutionId(v int64)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *OrderAmendmentsResponseResultInner) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetOrigClientOrderId

`func (o *OrderAmendmentsResponseResultInner) GetOrigClientOrderId() string`

GetOrigClientOrderId returns the OrigClientOrderId field if non-nil, zero value otherwise.

### GetOrigClientOrderIdOk

`func (o *OrderAmendmentsResponseResultInner) GetOrigClientOrderIdOk() (*string, bool)`

GetOrigClientOrderIdOk returns a tuple with the OrigClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigClientOrderId

`func (o *OrderAmendmentsResponseResultInner) SetOrigClientOrderId(v string)`

SetOrigClientOrderId sets OrigClientOrderId field to given value.

### HasOrigClientOrderId

`func (o *OrderAmendmentsResponseResultInner) HasOrigClientOrderId() bool`

HasOrigClientOrderId returns a boolean if a field has been set.

### GetNewClientOrderId

`func (o *OrderAmendmentsResponseResultInner) GetNewClientOrderId() string`

GetNewClientOrderId returns the NewClientOrderId field if non-nil, zero value otherwise.

### GetNewClientOrderIdOk

`func (o *OrderAmendmentsResponseResultInner) GetNewClientOrderIdOk() (*string, bool)`

GetNewClientOrderIdOk returns a tuple with the NewClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientOrderId

`func (o *OrderAmendmentsResponseResultInner) SetNewClientOrderId(v string)`

SetNewClientOrderId sets NewClientOrderId field to given value.

### HasNewClientOrderId

`func (o *OrderAmendmentsResponseResultInner) HasNewClientOrderId() bool`

HasNewClientOrderId returns a boolean if a field has been set.

### GetOrigQty

`func (o *OrderAmendmentsResponseResultInner) GetOrigQty() string`

GetOrigQty returns the OrigQty field if non-nil, zero value otherwise.

### GetOrigQtyOk

`func (o *OrderAmendmentsResponseResultInner) GetOrigQtyOk() (*string, bool)`

GetOrigQtyOk returns a tuple with the OrigQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigQty

`func (o *OrderAmendmentsResponseResultInner) SetOrigQty(v string)`

SetOrigQty sets OrigQty field to given value.

### HasOrigQty

`func (o *OrderAmendmentsResponseResultInner) HasOrigQty() bool`

HasOrigQty returns a boolean if a field has been set.

### GetNewQty

`func (o *OrderAmendmentsResponseResultInner) GetNewQty() string`

GetNewQty returns the NewQty field if non-nil, zero value otherwise.

### GetNewQtyOk

`func (o *OrderAmendmentsResponseResultInner) GetNewQtyOk() (*string, bool)`

GetNewQtyOk returns a tuple with the NewQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewQty

`func (o *OrderAmendmentsResponseResultInner) SetNewQty(v string)`

SetNewQty sets NewQty field to given value.

### HasNewQty

`func (o *OrderAmendmentsResponseResultInner) HasNewQty() bool`

HasNewQty returns a boolean if a field has been set.

### GetTime

`func (o *OrderAmendmentsResponseResultInner) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *OrderAmendmentsResponseResultInner) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *OrderAmendmentsResponseResultInner) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *OrderAmendmentsResponseResultInner) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to README]](../README.md)


