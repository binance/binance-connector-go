# GetFiatPaymentsHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetFiatPaymentsHistoryResponseDataInner**](GetFiatPaymentsHistoryResponseDataInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetFiatPaymentsHistoryResponse

`func NewGetFiatPaymentsHistoryResponse() *GetFiatPaymentsHistoryResponse`

NewGetFiatPaymentsHistoryResponse instantiates a new GetFiatPaymentsHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatPaymentsHistoryResponseWithDefaults

`func NewGetFiatPaymentsHistoryResponseWithDefaults() *GetFiatPaymentsHistoryResponse`

NewGetFiatPaymentsHistoryResponseWithDefaults instantiates a new GetFiatPaymentsHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetFiatPaymentsHistoryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFiatPaymentsHistoryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFiatPaymentsHistoryResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFiatPaymentsHistoryResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetFiatPaymentsHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFiatPaymentsHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFiatPaymentsHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFiatPaymentsHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetFiatPaymentsHistoryResponse) GetData() []GetFiatPaymentsHistoryResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFiatPaymentsHistoryResponse) GetDataOk() (*[]GetFiatPaymentsHistoryResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFiatPaymentsHistoryResponse) SetData(v []GetFiatPaymentsHistoryResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetFiatPaymentsHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetFiatPaymentsHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFiatPaymentsHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFiatPaymentsHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFiatPaymentsHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *GetFiatPaymentsHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFiatPaymentsHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFiatPaymentsHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetFiatPaymentsHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


