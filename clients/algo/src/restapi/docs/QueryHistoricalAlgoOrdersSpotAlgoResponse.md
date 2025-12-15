# QueryHistoricalAlgoOrdersSpotAlgoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner**](QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner.md) |  | [optional] 

## Methods

### NewQueryHistoricalAlgoOrdersSpotAlgoResponse

`func NewQueryHistoricalAlgoOrdersSpotAlgoResponse() *QueryHistoricalAlgoOrdersSpotAlgoResponse`

NewQueryHistoricalAlgoOrdersSpotAlgoResponse instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryHistoricalAlgoOrdersSpotAlgoResponseWithDefaults

`func NewQueryHistoricalAlgoOrdersSpotAlgoResponseWithDefaults() *QueryHistoricalAlgoOrdersSpotAlgoResponse`

NewQueryHistoricalAlgoOrdersSpotAlgoResponseWithDefaults instantiates a new QueryHistoricalAlgoOrdersSpotAlgoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOrders

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetOrders() []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) GetOrdersOk() (*[]QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) SetOrders(v []QueryHistoricalAlgoOrdersSpotAlgoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryHistoricalAlgoOrdersSpotAlgoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


