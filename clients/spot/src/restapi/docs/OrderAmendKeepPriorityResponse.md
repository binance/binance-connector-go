# OrderAmendKeepPriorityResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**TransactTime** | Pointer to **int64** |  | [optional] 
**ExecutionId** | Pointer to **int64** |  | [optional] 
**AmendedOrder** | Pointer to [**OrderAmendKeepPriorityResponseAmendedOrder**](OrderAmendKeepPriorityResponseAmendedOrder.md) |  | [optional] 
**ListStatus** | Pointer to [**OrderAmendKeepPriorityResponseListStatus**](OrderAmendKeepPriorityResponseListStatus.md) |  | [optional] 

## Methods

### NewOrderAmendKeepPriorityResponse

`func NewOrderAmendKeepPriorityResponse() *OrderAmendKeepPriorityResponse`

NewOrderAmendKeepPriorityResponse instantiates a new OrderAmendKeepPriorityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAmendKeepPriorityResponseWithDefaults

`func NewOrderAmendKeepPriorityResponseWithDefaults() *OrderAmendKeepPriorityResponse`

NewOrderAmendKeepPriorityResponseWithDefaults instantiates a new OrderAmendKeepPriorityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactTime

`func (o *OrderAmendKeepPriorityResponse) GetTransactTime() int64`

GetTransactTime returns the TransactTime field if non-nil, zero value otherwise.

### GetTransactTimeOk

`func (o *OrderAmendKeepPriorityResponse) GetTransactTimeOk() (*int64, bool)`

GetTransactTimeOk returns a tuple with the TransactTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactTime

`func (o *OrderAmendKeepPriorityResponse) SetTransactTime(v int64)`

SetTransactTime sets TransactTime field to given value.

### HasTransactTime

`func (o *OrderAmendKeepPriorityResponse) HasTransactTime() bool`

HasTransactTime returns a boolean if a field has been set.

### GetExecutionId

`func (o *OrderAmendKeepPriorityResponse) GetExecutionId() int64`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *OrderAmendKeepPriorityResponse) GetExecutionIdOk() (*int64, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *OrderAmendKeepPriorityResponse) SetExecutionId(v int64)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *OrderAmendKeepPriorityResponse) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetAmendedOrder

`func (o *OrderAmendKeepPriorityResponse) GetAmendedOrder() OrderAmendKeepPriorityResponseAmendedOrder`

GetAmendedOrder returns the AmendedOrder field if non-nil, zero value otherwise.

### GetAmendedOrderOk

`func (o *OrderAmendKeepPriorityResponse) GetAmendedOrderOk() (*OrderAmendKeepPriorityResponseAmendedOrder, bool)`

GetAmendedOrderOk returns a tuple with the AmendedOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendedOrder

`func (o *OrderAmendKeepPriorityResponse) SetAmendedOrder(v OrderAmendKeepPriorityResponseAmendedOrder)`

SetAmendedOrder sets AmendedOrder field to given value.

### HasAmendedOrder

`func (o *OrderAmendKeepPriorityResponse) HasAmendedOrder() bool`

HasAmendedOrder returns a boolean if a field has been set.

### GetListStatus

`func (o *OrderAmendKeepPriorityResponse) GetListStatus() OrderAmendKeepPriorityResponseListStatus`

GetListStatus returns the ListStatus field if non-nil, zero value otherwise.

### GetListStatusOk

`func (o *OrderAmendKeepPriorityResponse) GetListStatusOk() (*OrderAmendKeepPriorityResponseListStatus, bool)`

GetListStatusOk returns a tuple with the ListStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListStatus

`func (o *OrderAmendKeepPriorityResponse) SetListStatus(v OrderAmendKeepPriorityResponseListStatus)`

SetListStatus sets ListStatus field to given value.

### HasListStatus

`func (o *OrderAmendKeepPriorityResponse) HasListStatus() bool`

HasListStatus returns a boolean if a field has been set.


[[Back to README]](../README.md)


