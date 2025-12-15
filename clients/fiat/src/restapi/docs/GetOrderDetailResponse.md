# GetOrderDetailResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**GetOrderDetailResponseData**](GetOrderDetailResponseData.md) |  | [optional] 

## Methods

### NewGetOrderDetailResponse

`func NewGetOrderDetailResponse() *GetOrderDetailResponse`

NewGetOrderDetailResponse instantiates a new GetOrderDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderDetailResponseWithDefaults

`func NewGetOrderDetailResponseWithDefaults() *GetOrderDetailResponse`

NewGetOrderDetailResponseWithDefaults instantiates a new GetOrderDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetOrderDetailResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOrderDetailResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOrderDetailResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetOrderDetailResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetOrderDetailResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOrderDetailResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOrderDetailResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetOrderDetailResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetOrderDetailResponse) GetData() GetOrderDetailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrderDetailResponse) GetDataOk() (*GetOrderDetailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrderDetailResponse) SetData(v GetOrderDetailResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetOrderDetailResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


