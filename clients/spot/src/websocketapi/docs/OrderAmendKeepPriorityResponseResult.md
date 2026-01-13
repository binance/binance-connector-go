# OrderAmendKeepPriorityResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TransactTime** | Pointer to **int64** |  | [optional] 
**ExecutionId** | Pointer to **int64** |  | [optional] 
**AmendedOrder** | Pointer to [**OrderAmendKeepPriorityResponseResultAmendedOrder**](OrderAmendKeepPriorityResponseResultAmendedOrder.md) |  | [optional] 
**ListStatus** | Pointer to [**OrderAmendKeepPriorityResponseResultListStatus**](OrderAmendKeepPriorityResponseResultListStatus.md) |  | [optional] 

## Methods

### NewOrderAmendKeepPriorityResponseResult

`func NewOrderAmendKeepPriorityResponseResult() *OrderAmendKeepPriorityResponseResult`

NewOrderAmendKeepPriorityResponseResult instantiates a new OrderAmendKeepPriorityResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseResultWithDefaults

`func NewOrderAmendKeepPriorityResponseResultWithDefaults() *OrderAmendKeepPriorityResponseResult`

NewOrderAmendKeepPriorityResponseResultWithDefaults instantiates a new OrderAmendKeepPriorityResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactTime

`func (o *OrderAmendKeepPriorityResponseResult) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderAmendKeepPriorityResponseResult) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderAmendKeepPriorityResponseResult) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderAmendKeepPriorityResponseResult) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetExecutionId

`func (o *OrderAmendKeepPriorityResponseResult) GetExecutionId() int64`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *OrderAmendKeepPriorityResponseResult) GetExecutionIdOk() (*int64, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *OrderAmendKeepPriorityResponseResult) SetExecutionId(v int64)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *OrderAmendKeepPriorityResponseResult) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetAmendedOrder

`func (o *OrderAmendKeepPriorityResponseResult) GetAmendedOrder() OrderAmendKeepPriorityResponseResultAmendedOrder`

GetAmendedOrder returns the AmendedOrder field if non-nil, zero value otherwise.

### GetAmendedOrderOk

`func (o *OrderAmendKeepPriorityResponseResult) GetAmendedOrderOk() (*OrderAmendKeepPriorityResponseResultAmendedOrder, bool)`

GetAmendedOrderOk returns a tuple with the AmendedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedOrder

`func (o *OrderAmendKeepPriorityResponseResult) SetAmendedOrder(v OrderAmendKeepPriorityResponseResultAmendedOrder)`

SetAmendedOrder sets AmendedOrder field to given value.

### HasAmendedOrder

`func (o *OrderAmendKeepPriorityResponseResult) HasAmendedOrder() bool`

HasAmendedOrder returns a boolean if a field has been set.

### GetListStatus

`func (o *OrderAmendKeepPriorityResponseResult) GetListStatus() OrderAmendKeepPriorityResponseResultListStatus`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *OrderAmendKeepPriorityResponseResult) GetListStatusOk() (*OrderAmendKeepPriorityResponseResultListStatus, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *OrderAmendKeepPriorityResponseResult) SetListStatus(v OrderAmendKeepPriorityResponseResultListStatus)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *OrderAmendKeepPriorityResponseResult) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


