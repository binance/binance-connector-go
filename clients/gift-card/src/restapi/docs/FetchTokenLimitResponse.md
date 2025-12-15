# FetchTokenLimitResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]FetchTokenLimitResponseDataInner**](FetchTokenLimitResponseDataInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewFetchTokenLimitResponse

`func NewFetchTokenLimitResponse() *FetchTokenLimitResponse`

NewFetchTokenLimitResponse instantiates a new FetchTokenLimitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchTokenLimitResponseWithDefaults

`func NewFetchTokenLimitResponseWithDefaults() *FetchTokenLimitResponse`

NewFetchTokenLimitResponseWithDefaults instantiates a new FetchTokenLimitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FetchTokenLimitResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FetchTokenLimitResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FetchTokenLimitResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FetchTokenLimitResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *FetchTokenLimitResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FetchTokenLimitResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FetchTokenLimitResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FetchTokenLimitResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *FetchTokenLimitResponse) GetData() []FetchTokenLimitResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FetchTokenLimitResponse) GetDataOk() (*[]FetchTokenLimitResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FetchTokenLimitResponse) SetData(v []FetchTokenLimitResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *FetchTokenLimitResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *FetchTokenLimitResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FetchTokenLimitResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FetchTokenLimitResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *FetchTokenLimitResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


