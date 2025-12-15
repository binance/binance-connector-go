# GetFiatDepositWithdrawHistoryResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]GetFiatDepositWithdrawHistoryResponseDataInner**](GetFiatDepositWithdrawHistoryResponseDataInner.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetFiatDepositWithdrawHistoryResponse

`func NewGetFiatDepositWithdrawHistoryResponse() *GetFiatDepositWithdrawHistoryResponse`

NewGetFiatDepositWithdrawHistoryResponse instantiates a new GetFiatDepositWithdrawHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFiatDepositWithdrawHistoryResponseWithDefaults

`func NewGetFiatDepositWithdrawHistoryResponseWithDefaults() *GetFiatDepositWithdrawHistoryResponse`

NewGetFiatDepositWithdrawHistoryResponseWithDefaults instantiates a new GetFiatDepositWithdrawHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetFiatDepositWithdrawHistoryResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFiatDepositWithdrawHistoryResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFiatDepositWithdrawHistoryResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFiatDepositWithdrawHistoryResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetFiatDepositWithdrawHistoryResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFiatDepositWithdrawHistoryResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFiatDepositWithdrawHistoryResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFiatDepositWithdrawHistoryResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetFiatDepositWithdrawHistoryResponse) GetData() []GetFiatDepositWithdrawHistoryResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFiatDepositWithdrawHistoryResponse) GetDataOk() (*[]GetFiatDepositWithdrawHistoryResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFiatDepositWithdrawHistoryResponse) SetData(v []GetFiatDepositWithdrawHistoryResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetFiatDepositWithdrawHistoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotal

`func (o *GetFiatDepositWithdrawHistoryResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetFiatDepositWithdrawHistoryResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetFiatDepositWithdrawHistoryResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetFiatDepositWithdrawHistoryResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSuccess

`func (o *GetFiatDepositWithdrawHistoryResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFiatDepositWithdrawHistoryResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFiatDepositWithdrawHistoryResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetFiatDepositWithdrawHistoryResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


