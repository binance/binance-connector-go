# TokenListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageDetail** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]TokenListResponseDataInner**](TokenListResponseDataInner.md) |  | [optional] 

## Methods

### NewTokenListResponse

`func NewTokenListResponse() *TokenListResponse`

NewTokenListResponse instantiates a new TokenListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenListResponseWithDefaults

`func NewTokenListResponseWithDefaults() *TokenListResponse`

NewTokenListResponseWithDefaults instantiates a new TokenListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TokenListResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TokenListResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TokenListResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TokenListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *TokenListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TokenListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TokenListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TokenListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *TokenListResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *TokenListResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *TokenListResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *TokenListResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetSuccess

`func (o *TokenListResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TokenListResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TokenListResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TokenListResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *TokenListResponse) GetData() []TokenListResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenListResponse) GetDataOk() (*[]TokenListResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenListResponse) SetData(v []TokenListResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *TokenListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


