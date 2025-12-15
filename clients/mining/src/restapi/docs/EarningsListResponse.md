# EarningsListResponse

## Properties

Name         | Type          | Description.  | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**EarningsListResponseData**](EarningsListResponseData.md) |  | [optional] 

## Methods

### NewEarningsListResponse

`func NewEarningsListResponse() *EarningsListResponse`

NewEarningsListResponse instantiates a new EarningsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEarningsListResponseWithDefaults

`func NewEarningsListResponseWithDefaults() *EarningsListResponse`

NewEarningsListResponseWithDefaults instantiates a new EarningsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *EarningsListResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EarningsListResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EarningsListResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *EarningsListResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *EarningsListResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *EarningsListResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *EarningsListResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *EarningsListResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *EarningsListResponse) GetData() EarningsListResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EarningsListResponse) GetDataOk() (*EarningsListResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EarningsListResponse) SetData(v EarningsListResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *EarningsListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to README]](../README.md)


