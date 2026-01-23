# TickerResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**MessageDetail** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**TickerResponseData**](TickerResponseData.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewTickerResponse

`func NewTickerResponse() *TickerResponse`

NewTickerResponse instantiates a new TickerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTickerResponseWithDefaults

`func NewTickerResponseWithDefaults() *TickerResponse`

NewTickerResponseWithDefaults instantiates a new TickerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TickerResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TickerResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TickerResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TickerResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *TickerResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TickerResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TickerResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TickerResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageDetail

`func (o *TickerResponse) GetMessageDetail() string`

GetMessageDetail returns the MessageDetail field if non-nil, zero value otherwise.

### GetMessageDetailOk

`func (o *TickerResponse) GetMessageDetailOk() (*string, bool)`

GetMessageDetailOk returns a tuple with the MessageDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDetail

`func (o *TickerResponse) SetMessageDetail(v string)`

SetMessageDetail sets MessageDetail field to given value.

### HasMessageDetail

`func (o *TickerResponse) HasMessageDetail() bool`

HasMessageDetail returns a boolean if a field has been set.

### GetData

`func (o *TickerResponse) GetData() TickerResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TickerResponse) GetDataOk() (*TickerResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TickerResponse) SetData(v TickerResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *TickerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *TickerResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *TickerResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *TickerResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *TickerResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


