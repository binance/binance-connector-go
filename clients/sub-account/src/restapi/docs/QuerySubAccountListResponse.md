# QuerySubAccountListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**SubAccounts** | Pointer to [**[]QuerySubAccountListResponseSubAccountsInner**](QuerySubAccountListResponseSubAccountsInner.md) |  | [optional] 

## Methods

### NewQuerySubAccountListResponse

`func NewQuerySubAccountListResponse() *QuerySubAccountListResponse`

NewQuerySubAccountListResponse instantiates a new QuerySubAccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySubAccountListResponseWithDefaults

`func NewQuerySubAccountListResponseWithDefaults() *QuerySubAccountListResponse`

NewQuerySubAccountListResponseWithDefaults instantiates a new QuerySubAccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubAccounts

`func (o *QuerySubAccountListResponse) GetSubAccounts() []QuerySubAccountListResponseSubAccountsInner`

GetSubAccounts returns the SubAccounts field if non-nil, zero value otherwise.

### GetSubAccountsOk

`func (o *QuerySubAccountListResponse) GetSubAccountsOk() (*[]QuerySubAccountListResponseSubAccountsInner, bool)`

GetSubAccountsOk returns a tuple with the SubAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccounts

`func (o *QuerySubAccountListResponse) SetSubAccounts(v []QuerySubAccountListResponseSubAccountsInner)`

SetSubAccounts sets SubAccounts field to given value.

### HasSubAccounts

`func (o *QuerySubAccountListResponse) HasSubAccounts() bool`

HasSubAccounts returns a boolean if a field has been set.


[[Back to README]](../README.md)


