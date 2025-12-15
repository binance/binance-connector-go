# QueryHistoricalAlgoOrdersFutureAlgoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner**](QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner.md) |  | [optional] 

## Methods

### NewQueryHistoricalAlgoOrdersFutureAlgoResponse

`func NewQueryHistoricalAlgoOrdersFutureAlgoResponse() *QueryHistoricalAlgoOrdersFutureAlgoResponse`

NewQueryHistoricalAlgoOrdersFutureAlgoResponse instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryHistoricalAlgoOrdersFutureAlgoResponseWithDefaults

`func NewQueryHistoricalAlgoOrdersFutureAlgoResponseWithDefaults() *QueryHistoricalAlgoOrdersFutureAlgoResponse`

NewQueryHistoricalAlgoOrdersFutureAlgoResponseWithDefaults instantiates a new QueryHistoricalAlgoOrdersFutureAlgoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOrders

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetOrders() []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) GetOrdersOk() (*[]QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) SetOrders(v []QueryHistoricalAlgoOrdersFutureAlgoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryHistoricalAlgoOrdersFutureAlgoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


