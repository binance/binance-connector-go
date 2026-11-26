# BstocksCorporateActionRecordsResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CorporateActionRecordsResponseItemsInner**](CorporateActionRecordsResponseItemsInner.md) | Corporate action asset change records on this page. | [optional] 
**NextTime** | Pointer to **int64** | Pass this value as &#x60;lastTime&#x60; on the next request. Not returned on the last page. | [optional] 
**NextId** | Pointer to **int64** | Pass this value as &#x60;lastId&#x60; on the next request. Not returned on the last page. | [optional] 

## Methods

### NewBstocksCorporateActionRecordsResponse

`func NewBstocksCorporateActionRecordsResponse() *BstocksCorporateActionRecordsResponse`

NewBstocksCorporateActionRecordsResponse instantiates a new BstocksCorporateActionRecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBstocksCorporateActionRecordsResponseWithDefaults

`func NewBstocksCorporateActionRecordsResponseWithDefaults() *BstocksCorporateActionRecordsResponse`

NewBstocksCorporateActionRecordsResponseWithDefaults instantiates a new BstocksCorporateActionRecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BstocksCorporateActionRecordsResponse) GetItems() []CorporateActionRecordsResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BstocksCorporateActionRecordsResponse) GetItemsOk() (*[]CorporateActionRecordsResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BstocksCorporateActionRecordsResponse) SetItems(v []CorporateActionRecordsResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *BstocksCorporateActionRecordsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextTime

`func (o *BstocksCorporateActionRecordsResponse) GetNextTime() int64`

GetNextTime returns the NextTime field if non-nil, zero value otherwise.

### GetNextTimeOk

`func (o *BstocksCorporateActionRecordsResponse) GetNextTimeOk() (*int64, bool)`

GetNextTimeOk returns a tuple with the NextTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTime

`func (o *BstocksCorporateActionRecordsResponse) SetNextTime(v int64)`

SetNextTime sets NextTime field to given value.

### HasNextTime

`func (o *BstocksCorporateActionRecordsResponse) HasNextTime() bool`

HasNextTime returns a boolean if a field has been set.

### GetNextId

`func (o *BstocksCorporateActionRecordsResponse) GetNextId() int64`

GetNextId returns the NextId field if non-nil, zero value otherwise.

### GetNextIdOk

`func (o *BstocksCorporateActionRecordsResponse) GetNextIdOk() (*int64, bool)`

GetNextIdOk returns a tuple with the NextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextId

`func (o *BstocksCorporateActionRecordsResponse) SetNextId(v int64)`

SetNextId sets NextId field to given value.

### HasNextId

`func (o *BstocksCorporateActionRecordsResponse) HasNextId() bool`

HasNextId returns a boolean if a field has been set.


[[Back to README]](../README.md)


