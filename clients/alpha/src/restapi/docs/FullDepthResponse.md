# FullDepthResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | API response code. \&quot;000000\&quot; indicates success. | [optional] 
**Message** | Pointer to **string** | Response message. | [optional] 
**MessageDetail** | Pointer to **string** | Detailed response message. | [optional] 
**Success** | Pointer to **bool** | Whether request is successful. | [optional] 
**Data** | Pointer to [**FullDepthResponseData**](FullDepthResponseData.md) |  | [optional] 

## Methods

### NewFullDepthResponse

`func NewFullDepthResponse() *FullDepthResponse`

NewFullDepthResponse instantiates a new FullDepthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullDepthResponseWithDefaults

`func NewFullDepthResponseWithDefaults() *FullDepthResponse`

NewFullDepthResponseWithDefaults instantiates a new FullDepthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FullDepthResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FullDepthResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FullDepthResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FullDepthResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *FullDepthResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FullDepthResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FullDepthResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FullDepthResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *FullDepthResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *FullDepthResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *FullDepthResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *FullDepthResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetSuccess

`func (o *FullDepthResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FullDepthResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FullDepthResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *FullDepthResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *FullDepthResponse) GetData() FullDepthResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FullDepthResponse) GetDataOk() (*FullDepthResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FullDepthResponse) SetData(v FullDepthResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *FullDepthResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


