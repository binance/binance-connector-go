# RpiOrderBookResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**E** | Pointer to **int64** |  | [optional] 
**T** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to [**[]RpiOrderBookResponseBidsItem**](RpiOrderBookResponseBidsItem.md) |  | [optional] 
**Asks** | Pointer to [**[]RpiOrderBookResponseAsksItem**](RpiOrderBookResponseAsksItem.md) |  | [optional] 

## Methods

### NewRpiOrderBookResponse

`func NewRpiOrderBookResponse() *RpiOrderBookResponse`

NewRpiOrderBookResponse instantiates a new RpiOrderBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRpiOrderBookResponseWithDefaults

`func NewRpiOrderBookResponseWithDefaults() *RpiOrderBookResponse`

NewRpiOrderBookResponseWithDefaults instantiates a new RpiOrderBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *RpiOrderBookResponse) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *RpiOrderBookResponse) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *RpiOrderBookResponse) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *RpiOrderBookResponse) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetE

`func (o *RpiOrderBookResponse) GetE() int64`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *RpiOrderBookResponse) GetEOk() (*int64, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *RpiOrderBookResponse) SetE(v int64)`

SetE sets E field to given value.

### HasE

`func (o *RpiOrderBookResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetT

`func (o *RpiOrderBookResponse) GetT() int64`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RpiOrderBookResponse) GetTOk() (*int64, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RpiOrderBookResponse) SetT(v int64)`

SetT sets T field to given value.

### HasT

`func (o *RpiOrderBookResponse) HasT() bool`

HasT returns a boolean if a field has been set.

### GetBids

`func (o *RpiOrderBookResponse) GetBids() []RpiOrderBookResponseBidsItem`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *RpiOrderBookResponse) GetBidsOk() (*[]RpiOrderBookResponseBidsItem, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *RpiOrderBookResponse) SetBids(v []RpiOrderBookResponseBidsItem)`

SetBids sets Bids field to given value.

### HasBids

`func (o *RpiOrderBookResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *RpiOrderBookResponse) GetAsks() []RpiOrderBookResponseAsksItem`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *RpiOrderBookResponse) GetAsksOk() (*[]RpiOrderBookResponseAsksItem, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *RpiOrderBookResponse) SetAsks(v []RpiOrderBookResponseAsksItem)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *RpiOrderBookResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


