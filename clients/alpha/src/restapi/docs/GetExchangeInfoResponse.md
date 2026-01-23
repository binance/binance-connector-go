# GetExchangeInfoResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageDetail** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**GetExchangeInfoResponseData**](GetExchangeInfoResponseData.md) |  | [optional] 

## Methods

### NewGetExchangeInfoResponse

`func NewGetExchangeInfoResponse() *GetExchangeInfoResponse`

NewGetExchangeInfoResponse instantiates a new GetExchangeInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExchangeInfoResponseWithDefaults

`func NewGetExchangeInfoResponseWithDefaults() *GetExchangeInfoResponse`

NewGetExchangeInfoResponseWithDefaults instantiates a new GetExchangeInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetExchangeInfoResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetExchangeInfoResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetExchangeInfoResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetExchangeInfoResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetExchangeInfoResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetExchangeInfoResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetExchangeInfoResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetExchangeInfoResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *GetExchangeInfoResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *GetExchangeInfoResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *GetExchangeInfoResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *GetExchangeInfoResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetSuccess

`func (o *GetExchangeInfoResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetExchangeInfoResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetExchangeInfoResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetExchangeInfoResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *GetExchangeInfoResponse) GetData() GetExchangeInfoResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExchangeInfoResponse) GetDataOk() (*GetExchangeInfoResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExchangeInfoResponse) SetData(v GetExchangeInfoResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetExchangeInfoResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


