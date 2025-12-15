# AccountListResponseDataInner

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**List** | Pointer to [**[]AccountListResponseDataInnerListInner**](AccountListResponseDataInnerListInner.md) |  | [optional] 

## Methods

### NewAccountListResponseDataInner

`func NewAccountListResponseDataInner() *AccountListResponseDataInner`

NewAccountListResponseDataInner instantiates a new AccountListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListResponseDataInnerWithDefaults

`func NewAccountListResponseDataInnerWithDefaults() *AccountListResponseDataInner`

NewAccountListResponseDataInnerWithDefaults instantiates a new AccountListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AccountListResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountListResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountListResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountListResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserName

`func (o *AccountListResponseDataInner) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AccountListResponseDataInner) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AccountListResponseDataInner) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AccountListResponseDataInner) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetList

`func (o *AccountListResponseDataInner) GetList() []AccountListResponseDataInnerListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *AccountListResponseDataInner) GetListOk() (*[]AccountListResponseDataInnerListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *AccountListResponseDataInner) SetList(v []AccountListResponseDataInnerListInner)`

SetList sets List field to given value.

### HasList

`func (o *AccountListResponseDataInner) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to README]](../README.md)


