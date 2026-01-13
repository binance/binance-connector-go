# OrderBookResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to [**[]OrderBookResponseResultBidsItem**](OrderBookResponseResultBidsItem.md) |  | [optional] 
**Asks** | Pointer to [**[]OrderBookResponseResultAsksItem**](OrderBookResponseResultAsksItem.md) |  | [optional] 

## Methods

### NewOrderBookResponseResult

`func NewOrderBookResponseResult() *OrderBookResponseResult`

NewOrderBookResponseResult instantiates a new OrderBookResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookResponseResultWithDefaults

`func NewOrderBookResponseResultWithDefaults() *OrderBookResponseResult`

NewOrderBookResponseResultWithDefaults instantiates a new OrderBookResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *OrderBookResponseResult) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *OrderBookResponseResult) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *OrderBookResponseResult) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *OrderBookResponseResult) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetE

`func (o *OrderBookResponseResult) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OrderBookResponseResult) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OrderBookResponseResult) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OrderBookResponseResult) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *OrderBookResponseResult) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *OrderBookResponseResult) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *OrderBookResponseResult) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *OrderBookResponseResult) HasT() bool`

HasT returns a boolean if a field has been set.

### GetBids

`func (o *OrderBookResponseResult) GetBids() []OrderBookResponseResultBidsItem`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *OrderBookResponseResult) GetBidsOk() (*[]OrderBookResponseResultBidsItem, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *OrderBookResponseResult) SetBids(v []OrderBookResponseResultBidsItem)`

SetBids sets Bids field to given value.

### HasBids

`func (o *OrderBookResponseResult) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *OrderBookResponseResult) GetAsks() []OrderBookResponseResultAsksItem`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *OrderBookResponseResult) GetAsksOk() (*[]OrderBookResponseResultAsksItem, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *OrderBookResponseResult) SetAsks(v []OrderBookResponseResultAsksItem)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *OrderBookResponseResult) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


