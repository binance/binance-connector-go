# GetFuturesLeadTraderStatusResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**GetFuturesLeadTraderStatusResponseData**](GetFuturesLeadTraderStatusResponseData.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetFuturesLeadTraderStatusResponse

`func NewGetFuturesLeadTraderStatusResponse() *GetFuturesLeadTraderStatusResponse`

NewGetFuturesLeadTraderStatusResponse instantiates a new GetFuturesLeadTraderStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFuturesLeadTraderStatusResponseWithDefaults

`func NewGetFuturesLeadTraderStatusResponseWithDefaults() *GetFuturesLeadTraderStatusResponse`

NewGetFuturesLeadTraderStatusResponseWithDefaults instantiates a new GetFuturesLeadTraderStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetFuturesLeadTraderStatusResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetFuturesLeadTraderStatusResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetFuturesLeadTraderStatusResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetFuturesLeadTraderStatusResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetFuturesLeadTraderStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetFuturesLeadTraderStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetFuturesLeadTraderStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetFuturesLeadTraderStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *GetFuturesLeadTraderStatusResponse) GetData() GetFuturesLeadTraderStatusResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFuturesLeadTraderStatusResponse) GetDataOk() (*GetFuturesLeadTraderStatusResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFuturesLeadTraderStatusResponse) SetData(v GetFuturesLeadTraderStatusResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetFuturesLeadTraderStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSuccess

`func (o *GetFuturesLeadTraderStatusResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetFuturesLeadTraderStatusResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetFuturesLeadTraderStatusResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetFuturesLeadTraderStatusResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to README]](../README.md)


