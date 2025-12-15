# DepositResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**DepositResponseData**](DepositResponseData.md) |  | [optional] 

## Methods

### NewDepositResponse

`func NewDepositResponse() *DepositResponse`

NewDepositResponse instantiates a new DepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositResponseWithDefaults

`func NewDepositResponseWithDefaults() *DepositResponse`

NewDepositResponseWithDefaults instantiates a new DepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DepositResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DepositResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DepositResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *DepositResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *DepositResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DepositResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DepositResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DepositResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *DepositResponse) GetData() DepositResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DepositResponse) GetDataOk() (*DepositResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DepositResponse) SetData(v DepositResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *DepositResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


