# AccountListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AccountListResponseDataInner**](AccountListResponseDataInner.md) |  | [optional] 

## Methods

### NewAccountListResponse

`func NewAccountListResponse() *AccountListResponse`

NewAccountListResponse instantiates a new AccountListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountListResponseWithDefaults

`func NewAccountListResponseWithDefaults() *AccountListResponse`

NewAccountListResponseWithDefaults instantiates a new AccountListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *AccountListResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AccountListResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AccountListResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *AccountListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *AccountListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AccountListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AccountListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AccountListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AccountListResponse) GetData() []AccountListResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccountListResponse) GetDataOk() (*[]AccountListResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccountListResponse) SetData(v []AccountListResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *AccountListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


