# CreateASingleTokenGiftCardResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CreateADualTokenGiftCardResponseData**](CreateADualTokenGiftCardResponseData.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateASingleTokenGiftCardResponse

`func NewCreateASingleTokenGiftCardResponse() *CreateASingleTokenGiftCardResponse`

NewCreateASingleTokenGiftCardResponse instantiates a new CreateASingleTokenGiftCardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateASingleTokenGiftCardResponseWithDefaults

`func NewCreateASingleTokenGiftCardResponseWithDefaults() *CreateASingleTokenGiftCardResponse`

NewCreateASingleTokenGiftCardResponseWithDefaults instantiates a new CreateASingleTokenGiftCardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateASingleTokenGiftCardResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateASingleTokenGiftCardResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateASingleTokenGiftCardResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateASingleTokenGiftCardResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *CreateASingleTokenGiftCardResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateASingleTokenGiftCardResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateASingleTokenGiftCardResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateASingleTokenGiftCardResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *CreateASingleTokenGiftCardResponse) GetData() CreateADualTokenGiftCardResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateASingleTokenGiftCardResponse) GetDataOk() (*CreateADualTokenGiftCardResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateASingleTokenGiftCardResponse) SetData(v CreateADualTokenGiftCardResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateASingleTokenGiftCardResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateASingleTokenGiftCardResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateASingleTokenGiftCardResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateASingleTokenGiftCardResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateASingleTokenGiftCardResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


