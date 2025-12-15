# FiatWithdrawResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**DepositResponseData**](DepositResponseData.md) |  | [optional] 

## Methods

### NewFiatWithdrawResponse

`func NewFiatWithdrawResponse() *FiatWithdrawResponse`

NewFiatWithdrawResponse instantiates a new FiatWithdrawResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiatWithdrawResponseWithDefaults

`func NewFiatWithdrawResponseWithDefaults() *FiatWithdrawResponse`

NewFiatWithdrawResponseWithDefaults instantiates a new FiatWithdrawResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *FiatWithdrawResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FiatWithdrawResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FiatWithdrawResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FiatWithdrawResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *FiatWithdrawResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FiatWithdrawResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FiatWithdrawResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FiatWithdrawResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *FiatWithdrawResponse) GetData() DepositResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FiatWithdrawResponse) GetDataOk() (*DepositResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FiatWithdrawResponse) SetData(v DepositResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *FiatWithdrawResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


