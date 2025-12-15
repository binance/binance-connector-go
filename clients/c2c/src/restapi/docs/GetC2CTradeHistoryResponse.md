# GetC2CTradeHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetC2CTradeHistoryResponseDataInner**](GetC2CTradeHistoryResponseDataInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetC2CTradeHistoryResponse

`func NewGetC2CTradeHistoryResponse() *GetC2CTradeHistoryResponse`

NewGetC2CTradeHistoryResponse instantiates a new GetC2CTradeHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetC2CTradeHistoryResponseWithDefaults

`func NewGetC2CTradeHistoryResponseWithDefaults() *GetC2CTradeHistoryResponse`

NewGetC2CTradeHistoryResponseWithDefaults instantiates a new GetC2CTradeHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetC2CTradeHistoryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetC2CTradeHistoryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetC2CTradeHistoryResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetC2CTradeHistoryResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetC2CTradeHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetC2CTradeHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetC2CTradeHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetC2CTradeHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetC2CTradeHistoryResponse) GetData() []GetC2CTradeHistoryResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetC2CTradeHistoryResponse) GetDataOk() (*[]GetC2CTradeHistoryResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetC2CTradeHistoryResponse) SetData(v []GetC2CTradeHistoryResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetC2CTradeHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetC2CTradeHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetC2CTradeHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetC2CTradeHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetC2CTradeHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *GetC2CTradeHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetC2CTradeHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetC2CTradeHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetC2CTradeHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


