# KlinesResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | API response code. \&quot;000000\&quot; indicates success. | [optional] 
**Message** | Pointer to **string** | Response message. | [optional] 
**MessageDetail** | Pointer to **string** | Detailed response message. | [optional] 
**Success** | Pointer to **bool** | Whether request is successful. | [optional] 
**Data** | Pointer to [**[][]KlinesResponseDataInnerInner**]([]KlinesResponseDataInnerInner.md) | Array of kline rows. Each row contains: open time, open, high, low, close, volume, close time, quote volume, trade count, taker buy base volume, taker buy quote volume, and a static 0. | [optional] 

## Methods

### NewKlinesResponse

`func NewKlinesResponse() *KlinesResponse`

NewKlinesResponse instantiates a new KlinesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKlinesResponseWithDefaults

`func NewKlinesResponseWithDefaults() *KlinesResponse`

NewKlinesResponseWithDefaults instantiates a new KlinesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *KlinesResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *KlinesResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *KlinesResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *KlinesResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *KlinesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KlinesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KlinesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KlinesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *KlinesResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *KlinesResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *KlinesResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *KlinesResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetSuccess

`func (o *KlinesResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *KlinesResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *KlinesResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *KlinesResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *KlinesResponse) GetData() [][]KlinesResponseDataInnerInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KlinesResponse) GetDataOk() (*[][]KlinesResponseDataInnerInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KlinesResponse) SetData(v [][]KlinesResponseDataInnerInner)`

SetData sets Data field to given value.

### HasData

`func (o *KlinesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


