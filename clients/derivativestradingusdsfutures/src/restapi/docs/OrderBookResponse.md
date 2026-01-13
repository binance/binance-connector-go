# OrderBookResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to [**[]OrderBookResponseBidsItem**](OrderBookResponseBidsItem.md) |  | [optional] 
**Asks** | Pointer to [**[]OrderBookResponseAsksItem**](OrderBookResponseAsksItem.md) |  | [optional] 

## Methods

### NewOrderBookResponse

`func NewOrderBookResponse() *OrderBookResponse`

NewOrderBookResponse instantiates a new OrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBookResponseWithDefaults

`func NewOrderBookResponseWithDefaults() *OrderBookResponse`

NewOrderBookResponseWithDefaults instantiates a new OrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *OrderBookResponse) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *OrderBookResponse) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *OrderBookResponse) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *OrderBookResponse) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetE

`func (o *OrderBookResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OrderBookResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OrderBookResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *OrderBookResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *OrderBookResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *OrderBookResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *OrderBookResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *OrderBookResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetBids

`func (o *OrderBookResponse) GetBids() []OrderBookResponseBidsItem`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *OrderBookResponse) GetBidsOk() (*[]OrderBookResponseBidsItem, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *OrderBookResponse) SetBids(v []OrderBookResponseBidsItem)`

SetBids sets Bids field to given value.

### HasBids

`func (o *OrderBookResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *OrderBookResponse) GetAsks() []OrderBookResponseAsksItem`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *OrderBookResponse) GetAsksOk() (*[]OrderBookResponseAsksItem, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *OrderBookResponse) SetAsks(v []OrderBookResponseAsksItem)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *OrderBookResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


