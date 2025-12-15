# QueryManagedSubAccountListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**ManagerSubUserInfoVoList** | Pointer to [**[]QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner**](QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner.md) |  | [optional] 

## Methods

### NewQueryManagedSubAccountListResponse

`func NewQueryManagedSubAccountListResponse() *QueryManagedSubAccountListResponse`

NewQueryManagedSubAccountListResponse instantiates a new QueryManagedSubAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryManagedSubAccountListResponseWithDefaults

`func NewQueryManagedSubAccountListResponseWithDefaults() *QueryManagedSubAccountListResponse`

NewQueryManagedSubAccountListResponseWithDefaults instantiates a new QueryManagedSubAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *QueryManagedSubAccountListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueryManagedSubAccountListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueryManagedSubAccountListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueryManagedSubAccountListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetManagerSubUserInfoVoList

`func (o *QueryManagedSubAccountListResponse) GetManagerSubUserInfoVoList() []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner`

GetManagerSubUserInfoVoList returns the ManagerSubUserInfoVoList field if non-nil, zero value otherwise.

### GetManagerSubUserInfoVoListOk

`func (o *QueryManagedSubAccountListResponse) GetManagerSubUserInfoVoListOk() (*[]QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner, bool)`

GetManagerSubUserInfoVoListOk returns a tuple with the ManagerSubUserInfoVoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerSubUserInfoVoList

`func (o *QueryManagedSubAccountListResponse) SetManagerSubUserInfoVoList(v []QueryManagedSubAccountListResponseManagerSubUserInfoVoListInner)`

SetManagerSubUserInfoVoList sets ManagerSubUserInfoVoList field to given value.

### HasManagerSubUserInfoVoList

`func (o *QueryManagedSubAccountListResponse) HasManagerSubUserInfoVoList() bool`

HasManagerSubUserInfoVoList returns a boolean if a field has been set.


[[Back to README]](../README.md)


