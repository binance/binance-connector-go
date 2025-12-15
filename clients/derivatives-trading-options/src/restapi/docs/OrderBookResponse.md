# OrderBookResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**T** | Pointer to **int64** |  | [optional] 
**U** | Pointer to **int64** |  | [optional] 
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

### GetU

`func (o *OrderBookResponse) GetU() int64`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *OrderBookResponse) GetUOk() (*int64, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *OrderBookResponse) SetU(v int64)`

SetU sets U field to given value.

### HasU

`func (o *OrderBookResponse) HasU() bool`

HasU returns a boolean if a field has been set.

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


