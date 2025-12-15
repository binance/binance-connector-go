# GetPayTradeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetPayTradeHistoryResponseDataInner**](GetPayTradeHistoryResponseDataInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetPayTradeHistoryResponse

`func NewGetPayTradeHistoryResponse() *GetPayTradeHistoryResponse`

NewGetPayTradeHistoryResponse instantiates a new GetPayTradeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayTradeHistoryResponseWithDefaults

`func NewGetPayTradeHistoryResponseWithDefaults() *GetPayTradeHistoryResponse`

NewGetPayTradeHistoryResponseWithDefaults instantiates a new GetPayTradeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetPayTradeHistoryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPayTradeHistoryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPayTradeHistoryResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPayTradeHistoryResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetPayTradeHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetPayTradeHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetPayTradeHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetPayTradeHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetPayTradeHistoryResponse) GetData() []GetPayTradeHistoryResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPayTradeHistoryResponse) GetDataOk() (*[]GetPayTradeHistoryResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPayTradeHistoryResponse) SetData(v []GetPayTradeHistoryResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetPayTradeHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *GetPayTradeHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetPayTradeHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetPayTradeHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetPayTradeHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


