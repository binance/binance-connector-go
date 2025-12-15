# DepthResponseResult

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**LastUpdateId** | Pointer to **int64** |  | [optional] 
**Bids** | Pointer to **[][]string** |  | [optional] 
**Asks** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewDepthResponseResult

`func NewDepthResponseResult() *DepthResponseResult`

NewDepthResponseResult instantiates a new DepthResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepthResponseResultWithDefaults

`func NewDepthResponseResultWithDefaults() *DepthResponseResult`

NewDepthResponseResultWithDefaults instantiates a new DepthResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdateId

`func (o *DepthResponseResult) GetLastUpdateId() int64`

GetLastUpdateId returns the LastUpdateId field if non-nil, zero value otherwise.

### GetLastUpdateIdOk

`func (o *DepthResponseResult) GetLastUpdateIdOk() (*int64, bool)`

GetLastUpdateIdOk returns a tuple with the LastUpdateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateId

`func (o *DepthResponseResult) SetLastUpdateId(v int64)`

SetLastUpdateId sets LastUpdateId field to given value.

### HasLastUpdateId

`func (o *DepthResponseResult) HasLastUpdateId() bool`

HasLastUpdateId returns a boolean if a field has been set.

### GetBids

`func (o *DepthResponseResult) GetBids() [][]string`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *DepthResponseResult) GetBidsOk() (*[][]string, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *DepthResponseResult) SetBids(v [][]string)`

SetBids sets Bids field to given value.

### HasBids

`func (o *DepthResponseResult) HasBids() bool`

HasBids returns a boolean if a field has been set.

### GetAsks

`func (o *DepthResponseResult) GetAsks() [][]string`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *DepthResponseResult) GetAsksOk() (*[][]string, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *DepthResponseResult) SetAsks(v [][]string)`

SetAsks sets Asks field to given value.

### HasAsks

`func (o *DepthResponseResult) HasAsks() bool`

HasAsks returns a boolean if a field has been set.


[[Back to README]](../README.md)


