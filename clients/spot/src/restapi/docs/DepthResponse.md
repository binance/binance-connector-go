# DepthResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**Asks** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewDepthResponse

`func NewDepthResponse() *DepthResponse`

NewDepthResponse instantiates a new DepthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepthResponseWithDefaults

`func NewDepthResponseWithDefaults() *DepthResponse`

NewDepthResponseWithDefaults instantiates a new DepthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *DepthResponse) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *DepthResponse) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *DepthResponse) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *DepthResponse) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetBids

`func (o *DepthResponse) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *DepthResponse) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *DepthResponse) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *DepthResponse) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *DepthResponse) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *DepthResponse) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *DepthResponse) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *DepthResponse) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


