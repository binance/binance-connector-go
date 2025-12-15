# QueryCurrentAlgoOpenOrdersSpotAlgoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Orders** | Pointer to [**[]QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner**](QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner.md) |  | [optional] 

## Methods

### NewQueryCurrentAlgoOpenOrdersSpotAlgoResponse

`func NewQueryCurrentAlgoOpenOrdersSpotAlgoResponse() *QueryCurrentAlgoOpenOrdersSpotAlgoResponse`

NewQueryCurrentAlgoOpenOrdersSpotAlgoResponse instantiates a new QueryCurrentAlgoOpenOrdersSpotAlgoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryCurrentAlgoOpenOrdersSpotAlgoResponseWithDefaults

`func NewQueryCurrentAlgoOpenOrdersSpotAlgoResponseWithDefaults() *QueryCurrentAlgoOpenOrdersSpotAlgoResponse`

NewQueryCurrentAlgoOpenOrdersSpotAlgoResponseWithDefaults instantiates a new QueryCurrentAlgoOpenOrdersSpotAlgoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOrders

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetOrders() []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) GetOrdersOk() (*[]QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) SetOrders(v []QueryCurrentAlgoOpenOrdersSpotAlgoResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *QueryCurrentAlgoOpenOrdersSpotAlgoResponse) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to README]](../README.md)


